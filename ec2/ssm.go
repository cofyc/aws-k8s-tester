package ec2

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/aws/aws-k8s-tester/ec2config"
	awscfn "github.com/aws/aws-k8s-tester/pkg/aws/cloudformation"
	"github.com/aws/aws-k8s-tester/version"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/dustin/go-humanize"
	"go.uber.org/zap"
)

func (ts *Tester) createSSM() error {
	if err := ts.createSSMDocument(); err != nil {
		return err
	}
	if err := ts.sendSSMDocumentCommand(); err != nil {
		return err
	}
	return ts.cfg.Sync()
}

func (ts *Tester) deleteSSM() error {
	if err := ts.deleteSSMDocument(); err != nil {
		return err
	}
	return ts.cfg.Sync()
}

// TemplateSSMDocument is the CFN template for SSM Document.
const TemplateSSMDocument = `
---
AWSTemplateFormatVersion: '2010-09-09'
Description: 'Amazon SSM Document'

Parameters:

  Name:
    Type: String
    Default: aws-k8s-tester-ec2-ssm

  DocumentName:
    Type: String
    Default: aws-k8s-tester-ec2-ssm-document

  Commands:
    Type: String

Resources:

  SSMDocument:
    Type: AWS::SSM::Document
    Properties:
      DocumentType: Command
      Tags:
      - Key: Name
        Value: !Ref Name
      - Key: DocumentName
        Value: !Ref DocumentName
      - Key: DocumentVersion
        Value: v1
      Content:
        schemaVersion: '2.2'
        description: SSM document to bootstrap EC2.
        parameters:
          region: {type: String, description: 'AWS Region', default: { Ref: "AWS::Region" } }
          executionTimeoutSeconds: {type: String, description: 'timeout for script, in seconds'}
          moreCommands: {type: String, description: 'more commands', default: { Ref: Commands } }
        mainSteps:
          - action: aws:runShellScript
            name: !Ref DocumentName
            inputs:
              timeoutSeconds: '{{ executionTimeoutSeconds }}'
              runCommand:
                - |
                  set -xue
                  log() {
                    echo -e "[$(date -u +'%Y-%m-%dT%H:%M:%SZ')] $1"
                  }
                  AWS_DEFAULT_REGION={{region}}
                  log "running SSM with AWS_DEFAULT_REGION: ${AWS_DEFAULT_REGION}"

                  log "running more SSM command"
                  {{ moreCommands }}

Outputs:

  SSMDocumentName:
    Value: !Ref SSMDocument

`

func (ts *Tester) createSSMDocument() error {
	createStart := time.Now()

	for asgName, cur := range ts.cfg.ASGs {
		if !cur.SSMDocumentCreate {
			ts.lg.Info("skipping SSM document create",
				zap.String("asg-name", asgName),
				zap.String("ssm-document-name", cur.SSMDocumentName),
			)
			continue
		}
		ts.lg.Info("creating SSM document",
			zap.String("asg-name", asgName),
			zap.String("ssm-document-name", cur.SSMDocumentName),
		)

		stackInput := &cloudformation.CreateStackInput{
			StackName:    aws.String(cur.SSMDocumentName),
			Capabilities: aws.StringSlice([]string{"CAPABILITY_IAM"}),
			OnFailure:    aws.String(cloudformation.OnFailureDelete),
			TemplateBody: aws.String(TemplateSSMDocument),
			Tags: awscfn.NewTags(map[string]string{
				"Kind":                   "aws-k8s-tester",
				"Name":                   ts.cfg.Name,
				"aws-k8s-tester-version": version.ReleaseVersion,
			}),
			Parameters: []*cloudformation.Parameter{
				{
					ParameterKey:   aws.String("Name"),
					ParameterValue: aws.String(ts.cfg.Name),
				},
				{
					ParameterKey:   aws.String("DocumentName"),
					ParameterValue: aws.String(cur.SSMDocumentName),
				},
			},
		}
		if len(cur.SSMDocumentCommands) > 0 {
			ts.lg.Info("added SSM document commands", zap.String("commands", cur.SSMDocumentCommands))
			stackInput.Parameters = append(stackInput.Parameters, &cloudformation.Parameter{
				ParameterKey:   aws.String("Commands"),
				ParameterValue: aws.String(cur.SSMDocumentCommands),
			})
		}
		stackOutput, err := ts.cfnAPI.CreateStack(stackInput)
		if err != nil {
			return err
		}
		cur.SSMDocumentCFNStackID = aws.StringValue(stackOutput.StackId)
		ts.cfg.ASGs[asgName] = cur
		ts.cfg.Sync()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
		ch := awscfn.Poll(
			ctx,
			ts.stopCreationCh,
			ts.interruptSig,
			ts.lg,
			ts.cfnAPI,
			cur.SSMDocumentCFNStackID,
			cloudformation.ResourceStatusCreateComplete,
			time.Minute,
			30*time.Second,
		)
		var st awscfn.StackStatus
		for st = range ch {
			if st.Error != nil {
				ts.cfg.RecordStatus(fmt.Sprintf("failed to create ASG (%v)", st.Error))
				ts.lg.Warn("polling errror", zap.Error(st.Error))
			}
		}
		cancel()
		if st.Error != nil {
			return st.Error
		}
		// update status after creating a new ASG
		for _, o := range st.Stack.Outputs {
			switch k := aws.StringValue(o.OutputKey); k {
			case "SSMDocumentName":
				cur.SSMDocumentName = aws.StringValue(o.OutputValue)
				ts.lg.Info("found SSMDocumentName value from CFN", zap.String("value", cur.SSMDocumentName))
			default:
				return fmt.Errorf("unexpected OutputKey %q from %q", k, cur.SSMDocumentCFNStackID)
			}
		}

		ts.lg.Info("created SSM Document",
			zap.String("asg-name", cur.Name),
			zap.String("ssm-document-name", cur.SSMDocumentName),
			zap.String("cfn-stack-id", cur.SSMDocumentCFNStackID),
			zap.String("started", humanize.RelTime(createStart, time.Now(), "ago", "from now")),
		)
		ts.cfg.ASGs[asgName] = cur
		ts.cfg.Sync()
	}

	return ts.cfg.Sync()
}

func (ts *Tester) deleteSSMDocument() error {
	for asgName, cur := range ts.cfg.ASGs {
		if !cur.SSMDocumentCreate {
			ts.lg.Info("skipping SSM document delete",
				zap.String("asg-name", asgName),
				zap.String("ssm-document-name", cur.SSMDocumentName),
			)
			continue
		}
		ts.lg.Info("deleting SSM document",
			zap.String("asg-name", cur.Name),
			zap.String("ssm-document-name", cur.SSMDocumentName),
			zap.String("cfn-stack-id", cur.SSMDocumentCFNStackID),
		)
		_, err := ts.cfnAPI.DeleteStack(&cloudformation.DeleteStackInput{
			StackName: aws.String(cur.SSMDocumentCFNStackID),
		})
		if err != nil {
			ts.cfg.RecordStatus(fmt.Sprintf("failed to delete SSM Document (%v)", err))
			return err
		}
		ts.cfg.Sync()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
		ch := awscfn.Poll(
			ctx,
			make(chan struct{}),  // do not exit on stop
			make(chan os.Signal), // do not exit on stop
			ts.lg,
			ts.cfnAPI,
			cur.SSMDocumentCFNStackID,
			cloudformation.ResourceStatusDeleteComplete,
			time.Minute,
			20*time.Second,
		)
		var st awscfn.StackStatus
		for st = range ch {
			if st.Error != nil {
				cancel()
				ts.cfg.RecordStatus(fmt.Sprintf("failed to delete SSM Document (%v)", st.Error))
				ts.lg.Warn("polling errror", zap.Error(st.Error))
			}
		}
		cancel()
		if st.Error != nil {
			return st.Error
		}
		ts.cfg.RecordStatus(fmt.Sprintf("%q/%s", cur.SSMDocumentName, ec2config.StatusDELETEDORNOTEXIST))

		ts.lg.Info("deleted SSM document",
			zap.String("asg-name", cur.Name),
			zap.String("ssm-document-name", cur.SSMDocumentName),
		)
	}

	return ts.cfg.Sync()
}

func (ts *Tester) sendSSMDocumentCommand() error {
	for asgName, cur := range ts.cfg.ASGs {
		if cur.SSMDocumentName == "" {
			ts.lg.Info("skipping SSM document send",
				zap.String("asg-name", asgName),
				zap.String("ssm-document-name", cur.SSMDocumentName),
			)
			continue
		}
		ids := make([]string, 0)
		for id := range cur.Instances {
			ids = append(ids, id)
		}
		if len(ids) == 0 {
			return fmt.Errorf("no instance found for SSM document %q", cur.SSMDocumentName)
		}
		ts.lg.Info("sending SSM document",
			zap.String("asg-name", asgName),
			zap.String("ssm-document-name", cur.SSMDocumentName),
			zap.Strings("instance-ids", ids),
		)
		ssmInput := &ssm.SendCommandInput{
			DocumentName:   aws.String(cur.SSMDocumentName),
			Comment:        aws.String(cur.SSMDocumentName + "-" + randString(10)),
			InstanceIds:    aws.StringSlice(ids),
			MaxConcurrency: aws.String(fmt.Sprintf("%d", len(ids))),
			Parameters: map[string][]*string{
				"region":                  aws.StringSlice([]string{ts.cfg.Region}),
				"executionTimeoutSeconds": aws.StringSlice([]string{fmt.Sprintf("%d", cur.SSMDocumentExecutionTimeoutSeconds)}),
			},
			OutputS3BucketName: aws.String(ts.cfg.S3BucketName),
			OutputS3KeyPrefix:  aws.String(path.Join(ts.cfg.Name, "ssm-outputs")),
		}
		if len(cur.SSMDocumentCommands) > 0 {
			ssmInput.Parameters["moreCommands"] = aws.StringSlice([]string{cur.SSMDocumentCommands})
		}
		cmd, err := ts.ssmAPI.SendCommand(ssmInput)
		if err != nil {
			return err
		}
		docName := aws.StringValue(cmd.Command.DocumentName)
		if docName != cur.SSMDocumentName {
			return fmt.Errorf("SSM Document Name expected %q, got %q", cur.SSMDocumentName, docName)
		}
		cur.SSMDocumentCommandID = aws.StringValue(cmd.Command.CommandId)
		ts.lg.Info("sent SSM document",
			zap.String("asg-name", asgName),
			zap.String("ssm-document-name", cur.SSMDocumentName),
			zap.String("ssm-command-id", cur.SSMDocumentCommandID),
		)
		ts.cfg.ASGs[asgName] = cur
		ts.cfg.Sync()
	}

	return ts.cfg.Sync()
}
