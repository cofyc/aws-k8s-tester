// Package fargate implements tester for Fargate.
package fargate

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-k8s-tester/eksconfig"
	awscfn "github.com/aws/aws-k8s-tester/pkg/aws/cloudformation"
	awsiam "github.com/aws/aws-k8s-tester/pkg/aws/iam"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/dustin/go-humanize"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/utils/exec"
)

// Config defines "Secrets" configuration.
// ref. https://aws.amazon.com/blogs/opensource/introducing-fine-grained-iam-roles-service-accounts/
type Config struct {
	Logger    *zap.Logger
	Stopc     chan struct{}
	Sig       chan os.Signal
	EKSConfig *eksconfig.Config
	K8SClient k8sClientSetGetter
	CFNAPI    cloudformationiface.CloudFormationAPI
	EKSAPI    eksiface.EKSAPI
	IAMAPI    iamiface.IAMAPI
}

type k8sClientSetGetter interface {
	KubernetesClientSet() *clientset.Clientset
}

// Tester defines Fargate tester.
type Tester interface {
	// Create creates Fargate pods.
	Create() error
	// Delete deletes Fargate pods.
	Delete() error
}

// New creates a new Job tester.
func New(cfg Config) (Tester, error) {
	return &tester{cfg: cfg}, nil
}

type tester struct {
	cfg Config
}

func (ts *tester) Create() error {
	if ts.cfg.EKSConfig.AddOnFargate.Created {
		ts.cfg.Logger.Info("skipping create AddOnFargate")
		return nil
	}

	ts.cfg.EKSConfig.AddOnFargate.Created = true
	ts.cfg.EKSConfig.Sync()

	createStart := time.Now()
	defer func() {
		ts.cfg.EKSConfig.AddOnFargate.CreateTook = time.Since(createStart)
		ts.cfg.EKSConfig.AddOnFargate.CreateTookString = ts.cfg.EKSConfig.AddOnFargate.CreateTook.String()
		ts.cfg.EKSConfig.Sync()
	}()

	if err := ts.createNamespace(); err != nil {
		return err
	}
	if err := ts.createRole(); err != nil {
		return err
	}
	if err := ts.createSecret(); err != nil {
		return err
	}
	if err := ts.createProfile(); err != nil {
		return err
	}
	if err := ts.createPod(); err != nil {
		return err
	}
	if err := ts.checkPod(); err != nil {
		return err
	}
	if err := ts.checkNode(); err != nil {
		return err
	}

	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) Delete() error {
	if !ts.cfg.EKSConfig.AddOnFargate.Created {
		ts.cfg.Logger.Info("skipping delete AddOnFargate")
		return nil
	}

	deleteStart := time.Now()
	defer func() {
		ts.cfg.EKSConfig.AddOnFargate.DeleteTook = time.Since(deleteStart)
		ts.cfg.EKSConfig.AddOnFargate.DeleteTookString = ts.cfg.EKSConfig.AddOnFargate.DeleteTook.String()
		ts.cfg.EKSConfig.Sync()
	}()

	var errs []string

	if err := ts.deletePod(); err != nil {
		errs = append(errs, fmt.Sprintf("failed to delete Fargate Pod (%v)", err))
	}
	ts.cfg.Logger.Info("wait after deleting Fargate Pod")

	if err := ts.deleteProfile(); err != nil {
		errs = append(errs, fmt.Sprintf("failed to delete Fargate profile (%v)", err))
	}
	ts.cfg.Logger.Info("wait after deleting Fargate profile")
	time.Sleep(10 * time.Second)

	if err := ts.deleteRole(); err != nil {
		errs = append(errs, fmt.Sprintf("failed to delete IAM Role (%v)", err))
	}
	ts.cfg.Logger.Info("wait after deleting IAM Role")
	time.Sleep(20 * time.Second)

	if err := ts.deleteSecret(); err != nil {
		return err
	}
	if err := ts.deleteNamespace(); err != nil {
		return err
	}

	ts.cfg.EKSConfig.AddOnFargate.Created = false
	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) createNamespace() error {
	ts.cfg.Logger.Info("creating namespace", zap.String("namespace", ts.cfg.EKSConfig.AddOnFargate.Namespace))
	_, err := ts.cfg.K8SClient.KubernetesClientSet().
		CoreV1().
		Namespaces().
		Create(&v1.Namespace{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind:       "Namespace",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: ts.cfg.EKSConfig.AddOnFargate.Namespace,
				Labels: map[string]string{
					"name": ts.cfg.EKSConfig.AddOnFargate.Namespace,
				},
			},
		})
	if err != nil {
		return err
	}
	ts.cfg.Logger.Info("created namespace", zap.String("namespace", ts.cfg.EKSConfig.AddOnFargate.Namespace))
	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) deleteNamespace() error {
	ts.cfg.Logger.Info("deleting namespace", zap.String("namespace", ts.cfg.EKSConfig.AddOnFargate.Namespace))
	foreground := metav1.DeletePropagationForeground
	err := ts.cfg.K8SClient.KubernetesClientSet().
		CoreV1().
		Namespaces().
		Delete(
			ts.cfg.EKSConfig.AddOnFargate.Namespace,
			&metav1.DeleteOptions{
				GracePeriodSeconds: aws.Int64(0),
				PropagationPolicy:  &foreground,
			},
		)
	if err != nil {
		// ref. https://github.com/aws/aws-k8s-tester/issues/79
		if !strings.Contains(err.Error(), ` not found`) {
			return err
		}
	}
	ts.cfg.Logger.Info("deleted namespace", zap.Error(err))
	return ts.cfg.EKSConfig.Sync()
}

// TemplateRole is the CloudFormation template for EKS Fargate role.
const TemplateRole = `
---
AWSTemplateFormatVersion: '2010-09-09'
Description: 'Amazon EKS Cluster Fargate Role'

Parameters:

  FargateRoleName:
    Type: String
    Description: The name of the Fargate role

  FargateRoleServicePrincipals:
    Type: CommaDelimitedList
    Default: 'eks.amazonaws.com,eks-fargate-pods.amazonaws.com'
    Description: EKS Fargate Role Service Principals

  FargateRoleManagedPolicyARNs:
    Type: CommaDelimitedList
    Default: 'arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy'
    Description: EKS Fargate policy ARNs

Resources:

  FargateRole:
    Type: AWS::IAM::Role
    Properties:
      RoleName: !Ref FargateRoleName
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
        - Effect: Allow
          Principal:
            Service: !Ref FargateRoleServicePrincipals
          Action:
          - sts:AssumeRole
      ManagedPolicyArns: !Ref FargateRoleManagedPolicyARNs
      Path: /

Outputs:

  FargateRoleARN:
    Value: !GetAtt FargateRole.Arn
    Description: The Fargate role ARN

`

func (ts *tester) createRole() error {
	if !ts.cfg.EKSConfig.AddOnFargate.RoleCreate {
		ts.cfg.Logger.Info("EKSConfig.AddOnFargate.RoleCreate false; skipping creation")
		return awsiam.Validate(
			ts.cfg.Logger,
			ts.cfg.IAMAPI,
			ts.cfg.EKSConfig.AddOnFargate.RoleName,
			[]string{
				"eks.amazonaws.com",
				"eks-fargate-pods.amazonaws.com",
			},
			[]string{
				"arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy",
			},
		)
	}
	if ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID != "" &&
		ts.cfg.EKSConfig.AddOnFargate.RoleARN != "" {
		ts.cfg.Logger.Info("role already created; no need to create a new one")
		return nil
	}
	if ts.cfg.EKSConfig.AddOnFargate.RoleName == "" {
		return errors.New("cannot create a cluster role with an empty AddOnFargate.RoleName")
	}

	ts.cfg.Logger.Info("creating a new Fargate role using CFN", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.RoleName))
	stackInput := &cloudformation.CreateStackInput{
		StackName:    aws.String(ts.cfg.EKSConfig.AddOnFargate.RoleName),
		Capabilities: aws.StringSlice([]string{"CAPABILITY_NAMED_IAM"}),
		OnFailure:    aws.String(cloudformation.OnFailureDelete),
		TemplateBody: aws.String(TemplateRole),
		Tags: awscfn.NewTags(map[string]string{
			"Kind": "aws-k8s-tester",
			"Name": ts.cfg.EKSConfig.Name,
		}),
		Parameters: []*cloudformation.Parameter{
			{
				ParameterKey:   aws.String("FargateRoleName"),
				ParameterValue: aws.String(ts.cfg.EKSConfig.AddOnFargate.RoleName),
			},
		},
	}
	if len(ts.cfg.EKSConfig.AddOnFargate.RoleServicePrincipals) > 0 {
		ts.cfg.Logger.Info("creating a new Fargate role with custom service principals",
			zap.Strings("service-principals", ts.cfg.EKSConfig.AddOnFargate.RoleServicePrincipals),
		)
		stackInput.Parameters = append(stackInput.Parameters, &cloudformation.Parameter{
			ParameterKey:   aws.String("FargateRoleServicePrincipals"),
			ParameterValue: aws.String(strings.Join(ts.cfg.EKSConfig.AddOnFargate.RoleServicePrincipals, ",")),
		})
	}
	if len(ts.cfg.EKSConfig.AddOnFargate.RoleManagedPolicyARNs) > 0 {
		ts.cfg.Logger.Info("creating a new Fargate role with custom managed role policies",
			zap.Strings("policy-arns", ts.cfg.EKSConfig.AddOnFargate.RoleManagedPolicyARNs),
		)
		stackInput.Parameters = append(stackInput.Parameters, &cloudformation.Parameter{
			ParameterKey:   aws.String("FargateRoleManagedPolicyARNs"),
			ParameterValue: aws.String(strings.Join(ts.cfg.EKSConfig.AddOnFargate.RoleManagedPolicyARNs, ",")),
		})
	}

	stackOutput, err := ts.cfg.CFNAPI.CreateStack(stackInput)
	if err != nil {
		return err
	}
	ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID = aws.StringValue(stackOutput.StackId)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	ch := awscfn.Poll(
		ctx,
		ts.cfg.Stopc,
		ts.cfg.Sig,
		ts.cfg.Logger,
		ts.cfg.CFNAPI,
		ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID,
		cloudformation.ResourceStatusCreateComplete,
		time.Minute,
		10*time.Second,
	)
	var st awscfn.StackStatus
	for st = range ch {
		if st.Error != nil {
			cancel()
			ts.cfg.EKSConfig.RecordStatus(fmt.Sprintf("failed to create Fargate role (%v)", st.Error))
			return st.Error
		}
	}
	cancel()

	for _, o := range st.Stack.Outputs {
		switch k := aws.StringValue(o.OutputKey); k {
		case "FargateRoleARN":
			ts.cfg.EKSConfig.AddOnFargate.RoleARN = aws.StringValue(o.OutputValue)
		default:
			return fmt.Errorf("unexpected OutputKey %q from %q", k, ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID)
		}
	}

	ts.cfg.Logger.Info("created a Fargate role",
		zap.String("cfn-stack-id", ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID),
		zap.String("role-name", ts.cfg.EKSConfig.AddOnFargate.RoleName),
		zap.String("role-arn", ts.cfg.EKSConfig.AddOnFargate.RoleARN),
	)
	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) deleteRole() error {
	if ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID == "" {
		ts.cfg.Logger.Info("empty Fargate role CFN stack ID; no need to delete Fargate")
		return nil
	}

	ts.cfg.Logger.Info("deleting Fargate role CFN stack",
		zap.String("role-cfn-stack-id", ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID),
	)
	_, err := ts.cfg.CFNAPI.DeleteStack(&cloudformation.DeleteStackInput{
		StackName: aws.String(ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID),
	})
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	ch := awscfn.Poll(
		ctx,
		make(chan struct{}),  // do not exit on stop
		make(chan os.Signal), // do not exit on stop
		ts.cfg.Logger,
		ts.cfg.CFNAPI,
		ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID,
		cloudformation.ResourceStatusDeleteComplete,
		time.Minute,
		10*time.Second,
	)
	var st awscfn.StackStatus
	for st = range ch {
		if st.Error != nil {
			cancel()
			ts.cfg.EKSConfig.RecordStatus(fmt.Sprintf("failed to delete Fargate role (%v)", st.Error))
			return st.Error
		}
	}
	cancel()
	ts.cfg.Logger.Info("deleted a Fargate role",
		zap.String("role-cfn-stack-id", ts.cfg.EKSConfig.AddOnFargate.RoleCFNStackID),
	)
	return ts.cfg.EKSConfig.Sync()
}

const secretReadTxt = "HELLO-WORLD-SECRET-IN-FARGATE"

func (ts *tester) createSecret() error {
	ts.cfg.Logger.Info("creating secret", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.SecretName))

	secret := &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ts.cfg.EKSConfig.AddOnFargate.SecretName,
			Namespace: ts.cfg.EKSConfig.AddOnFargate.Namespace,
		},
		Type: v1.SecretTypeOpaque,
		Data: map[string][]byte{ts.cfg.EKSConfig.AddOnFargate.SecretName: []byte(secretReadTxt)},
	}
	_, err := ts.cfg.K8SClient.KubernetesClientSet().
		CoreV1().
		Secrets(ts.cfg.EKSConfig.AddOnFargate.Namespace).
		Create(secret)
	if err != nil {
		return err
	}

	ts.cfg.Logger.Info("created secret", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.SecretName))
	return ts.cfg.EKSConfig.Sync()
}

var propagationBackground = metav1.DeletePropagationBackground

func (ts *tester) deleteSecret() error {
	ts.cfg.Logger.Info("deleting Secret", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.SecretName))
	err := ts.cfg.
		K8SClient.KubernetesClientSet().
		CoreV1().
		Secrets(ts.cfg.EKSConfig.AddOnFargate.Namespace).
		Delete(
			ts.cfg.EKSConfig.AddOnFargate.SecretName,
			&metav1.DeleteOptions{
				GracePeriodSeconds: aws.Int64(0),
				PropagationPolicy:  &propagationBackground,
			},
		)
	if err != nil {
		return fmt.Errorf("failed to delete Secret %q (%v)", ts.cfg.EKSConfig.AddOnFargate.SecretName, err)
	}
	ts.cfg.Logger.Info("deleted Secret", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.SecretName))
	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) createProfile() error {
	if ts.cfg.EKSConfig.AddOnFargate.RoleARN == "" {
		return errors.New("empty AddOnFargate.RoleARN")
	}
	if len(ts.cfg.EKSConfig.Parameters.PrivateSubnetIDs) == 0 {
		return errors.New("empty Parameters.PrivateSubnetIDs")
	}
	ts.cfg.Logger.Info("creating fargate profile", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.ProfileName))

	req, _ := ts.cfg.EKSAPI.CreateFargateProfileRequest(&eks.CreateFargateProfileInput{
		ClusterName:         aws.String(ts.cfg.EKSConfig.Name),
		FargateProfileName:  aws.String(ts.cfg.EKSConfig.AddOnFargate.ProfileName),
		PodExecutionRoleArn: aws.String(ts.cfg.EKSConfig.AddOnFargate.RoleARN),
		Subnets:             aws.StringSlice(ts.cfg.EKSConfig.Parameters.PrivateSubnetIDs),
		Selectors: []*eks.FargateProfileSelector{
			{
				Namespace: aws.String(ts.cfg.EKSConfig.AddOnFargate.Namespace),
			},
		},
	})
	err := req.Send()
	if err != nil {
		return err
	}
	ts.cfg.Logger.Info("sent create fargate profile request")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	ch := Poll(
		ctx,
		ts.cfg.Stopc,
		ts.cfg.Logger,
		ts.cfg.EKSAPI,
		ts.cfg.EKSConfig.Name,
		ts.cfg.EKSConfig.AddOnFargate.ProfileName,
		eks.FargateProfileStatusActive,
		10*time.Second,
		7*time.Second,
	)
	for sv := range ch {
		if sv.Error != nil {
			cancel()
			return sv.Error
		}
	}
	cancel()

	ts.cfg.Logger.Info("created fargate profile", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.ProfileName))
	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) deleteProfile() error {
	ts.cfg.Logger.Info("deleting fargate profile", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.ProfileName))

	var err error
	retryStart := time.Now()
	for time.Now().Sub(retryStart) < time.Minute {
		_, err = ts.cfg.EKSAPI.DeleteFargateProfile(&eks.DeleteFargateProfileInput{
			ClusterName:        aws.String(ts.cfg.EKSConfig.Name),
			FargateProfileName: aws.String(ts.cfg.EKSConfig.AddOnFargate.ProfileName),
		})
		if err != nil {
			if IsProfileDeleted(err) {
				err = nil
				break
			}
			ts.cfg.Logger.Warn("failed to delete fargate profile; retrying", zap.Error(err))
			select {
			case <-ts.cfg.Stopc:
				ts.cfg.Logger.Warn("aborted")
				return nil
			case <-time.After(5 * time.Second):
			}
			continue
		}
		ts.cfg.Logger.Warn("requested to delete fargate profile")
		break
	}

	ch := Poll(
		context.Background(),
		ts.cfg.Stopc,
		ts.cfg.Logger,
		ts.cfg.EKSAPI,
		ts.cfg.EKSConfig.Name,
		ts.cfg.EKSConfig.AddOnFargate.ProfileName,
		FargateProfileStatusDELETEDORNOTEXIST,
		10*time.Second,
		7*time.Second,
	)
	for sv := range ch {
		if sv.Error != nil {
			return sv.Error
		}
	}

	ts.cfg.Logger.Info("deleted fargate profile", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.ProfileName))
	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) createPod() error {
	if err := ts.listPods(ts.cfg.EKSConfig.AddOnFargate.Namespace); err != nil {
		ts.cfg.Logger.Warn("listing pods failed", zap.Error(err))
	}

	ts.cfg.Logger.Info("creating Pod")

	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ts.cfg.EKSConfig.AddOnFargate.PodName,
			Namespace: ts.cfg.EKSConfig.AddOnFargate.Namespace,
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyOnFailure,
			Containers: []v1.Container{
				{
					Name:            ts.cfg.EKSConfig.AddOnFargate.ContainerName,
					Image:           "amazonlinux",
					ImagePullPolicy: v1.PullIfNotPresent,
					Command: []string{
						"/bin/sh",
						"-c",
					},
					Args: []string{
						fmt.Sprintf("cat /tmp/%s && sleep 10000", ts.cfg.EKSConfig.AddOnFargate.SecretName),
					},

					// ref. https://kubernetes.io/docs/concepts/cluster-administration/logging/
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "secret-volume",
							MountPath: "/tmp",
							ReadOnly:  true,
						},
					},
				},
			},

			// ref. https://kubernetes.io/docs/concepts/cluster-administration/logging/
			Volumes: []v1.Volume{
				{ // to read
					Name: "secret-volume",
					VolumeSource: v1.VolumeSource{
						Secret: &v1.SecretVolumeSource{
							SecretName: ts.cfg.EKSConfig.AddOnFargate.SecretName,
						},
					},
				},
			},
		},
	}
	_, err := ts.cfg.K8SClient.KubernetesClientSet().
		CoreV1().
		Pods(ts.cfg.EKSConfig.AddOnFargate.Namespace).
		Create(pod)
	if err != nil {
		return err
	}

	ts.cfg.Logger.Info("created Pod")
	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) deletePod() error {
	ts.cfg.Logger.Info("deleting Pod", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.PodName))
	err := ts.cfg.
		K8SClient.KubernetesClientSet().
		CoreV1().
		Pods(ts.cfg.EKSConfig.AddOnFargate.Namespace).
		Delete(
			ts.cfg.EKSConfig.AddOnFargate.PodName,
			&metav1.DeleteOptions{
				GracePeriodSeconds: aws.Int64(0),
				PropagationPolicy:  &propagationBackground,
			},
		)
	if err != nil {
		return fmt.Errorf("failed to delete Pod %q (%v)", ts.cfg.EKSConfig.AddOnFargate.PodName, err)
	}
	ts.cfg.Logger.Info("deleted Pod", zap.String("name", ts.cfg.EKSConfig.AddOnFargate.PodName))
	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) listPods(ns string) error {
	pods, err := ts.getPods(ns)
	if err != nil {
		return err
	}
	println()
	for _, v := range pods.Items {
		fmt.Printf("%q Pod using client-go: %q\n", ns, v.Name)
	}
	println()
	return nil
}

func (ts *tester) getPods(ns string) (*v1.PodList, error) {
	return ts.cfg.K8SClient.KubernetesClientSet().CoreV1().Pods(ns).List(metav1.ListOptions{})
}

func (ts *tester) checkPod() error {
	cmdFlags := []string{
		"--namespace=" + ts.cfg.EKSConfig.AddOnFargate.Namespace,
		"--kubeconfig=" + ts.cfg.EKSConfig.KubeConfigPath,
		"exec",
		"-it",
		ts.cfg.EKSConfig.AddOnFargate.PodName,
		"--",
		"cat",
		fmt.Sprintf("/tmp/%s", ts.cfg.EKSConfig.AddOnFargate.SecretName),
	}
	ts.cfg.Logger.Info("checking Pod exec",
		zap.String("container-name", ts.cfg.EKSConfig.AddOnFargate.ContainerName),
		zap.String("command", ts.cfg.EKSConfig.KubectlPath+" "+strings.Join(cmdFlags, " ")),
	)
	found := false
	retryStart, waitDur := time.Now(), 3*time.Minute
	for time.Now().Sub(retryStart) < waitDur {
		select {
		case <-ts.cfg.Stopc:
			ts.cfg.Logger.Warn("aborted")
			return nil
		case <-time.After(5 * time.Second):
		}

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		output, err := exec.New().CommandContext(
			ctx,
			ts.cfg.EKSConfig.KubectlPath,
			cmdFlags...,
		).CombinedOutput()
		cancel()
		out := string(output)
		if err != nil {
			ts.cfg.Logger.Warn("'kubectl exec' failed", zap.String("output", out), zap.Error(err))
			continue
		}
		if !strings.Contains(out, secretReadTxt) {
			ts.cfg.Logger.Warn("unexpected exec output", zap.String("output", out))
			continue
		}
		ts.cfg.Logger.Info("successfully checked Pod exec",
			zap.String("container-name", ts.cfg.EKSConfig.AddOnFargate.ContainerName),
			zap.String("output", out),
		)
		found = true
		break
	}

	if !found {
		ts.cfg.EKSConfig.Sync()
		return errors.New("failed to check Pod")
	}

	// TODO: not working...
	/*
		cmdFlags = []string{
			"--namespace=" + ts.cfg.EKSConfig.AddOnFargate.Namespace,
			"--kubeconfig=" + ts.cfg.EKSConfig.KubeConfigPath,
			"logs",
			ts.cfg.EKSConfig.AddOnFargate.PodName,
			"--timestamps",
		}
		ts.cfg.Logger.Info("checking Pod logs",
			zap.String("container-name", ts.cfg.EKSConfig.AddOnFargate.ContainerName),
			zap.String("command", ts.cfg.EKSConfig.KubectlPath+" "+strings.Join(cmdFlags, " ")),
		)
		retryStart, waitDur = time.Now(), 2*time.Minute
		for time.Now().Sub(retryStart) < waitDur {
			select {
			case <-ts.cfg.Stopc:
				ts.cfg.Logger.Warn("aborted")
				return nil
			case <-time.After(5 * time.Second):
			}

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			output, err := exec.New().CommandContext(
				ctx,
				ts.cfg.EKSConfig.KubectlPath,
				cmdFlags...,
			).CombinedOutput()
			cancel()
			out := string(output)
			if err != nil {
				ts.cfg.Logger.Warn("'kubectl logs' failed", zap.String("output", out), zap.Error(err))
				continue
			}
			if !strings.Contains(out, secretReadTxt) {
				ts.cfg.Logger.Warn("unexpected logs output", zap.String("output", out))
				continue
			}
			ts.cfg.Logger.Info("checked Pod logs",
				zap.String("container-name", ts.cfg.EKSConfig.AddOnFargate.ContainerName),
				zap.String("output", out),
			)
			break
		}
	*/

	return ts.cfg.EKSConfig.Sync()
}

func (ts *tester) checkNode() error {
	ts.cfg.Logger.Info("checking node")

	desired := 1
	retryStart, waitDur := time.Now(), 3*time.Minute
	for time.Now().Sub(retryStart) < waitDur {
		select {
		case <-ts.cfg.Stopc:
			ts.cfg.Logger.Warn("aborted")
			return nil
		case <-time.After(5 * time.Second):
		}

		nodes, err := ts.cfg.K8SClient.KubernetesClientSet().CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			ts.cfg.Logger.Warn("get nodes failed", zap.Error(err))
			continue
		}
		items := nodes.Items

		readies := 0
		for _, node := range items {
			for _, cond := range node.Status.Conditions {
				if cond.Type != v1.NodeReady {
					continue
				}
				name := node.GetName()
				ts.cfg.Logger.Info("node info",
					zap.String("name", name),
					zap.String("type", fmt.Sprintf("%s", cond.Type)),
					zap.String("status", fmt.Sprintf("%s", cond.Status)),
				)
				if cond.Status == v1.ConditionTrue && strings.HasPrefix(name, "fargate-") {
					readies++
				}
			}
		}
		ts.cfg.Logger.Info("nodes",
			zap.Int("current-ready-nodes", readies),
			zap.Int("desired-ready-nodes", desired),
		)
		if readies >= desired {
			break
		}
	}

	ts.cfg.Logger.Info("checked node")
	return ts.cfg.EKSConfig.Sync()
}

// FargateProfileStatusDELETEDORNOTEXIST defines the cluster status when the cluster is not found.
//
// ref. https://docs.aws.amazon.com/eks/latest/APIReference/API_FargateProfile.html
//
//  CREATING
//  ACTIVE
//  DELETING
//  CREATE_FAILED
//  DELETE_FAILED
//
const FargateProfileStatusDELETEDORNOTEXIST = "DELETED/NOT-EXIST"

// FargateProfileStatus represents the CloudFormation status.
type FargateProfileStatus struct {
	FargateProfile *eks.FargateProfile
	Error          error
}

// Poll periodically fetches the fargate profile status
// until the node group becomes the desired state.
func Poll(
	ctx context.Context,
	stopc chan struct{},
	lg *zap.Logger,
	eksAPI eksiface.EKSAPI,
	clusterName string,
	profileName string,
	desiredStatus string,
	initialWait time.Duration,
	wait time.Duration,
) <-chan FargateProfileStatus {
	lg.Info("polling fargate profile",
		zap.String("cluster-name", clusterName),
		zap.String("profile-name", profileName),
		zap.String("desired-fargate-status", desiredStatus),
	)

	now := time.Now()

	ch := make(chan FargateProfileStatus, 10)
	go func() {
		// very first poll should be no-wait
		// in case stack has already reached desired status
		// wait from second interation
		waitDur := time.Duration(0)

		first := true
		for ctx.Err() == nil {
			select {
			case <-ctx.Done():
				lg.Warn("wait aborted", zap.Error(ctx.Err()))
				ch <- FargateProfileStatus{FargateProfile: nil, Error: ctx.Err()}
				close(ch)
				return

			case <-stopc:
				lg.Warn("wait stopped", zap.Error(ctx.Err()))
				ch <- FargateProfileStatus{FargateProfile: nil, Error: errors.New("wait stopped")}
				close(ch)
				return

			case <-time.After(waitDur):
				// very first poll should be no-wait
				// in case stack has already reached desired status
				// wait from second interation
				if waitDur == time.Duration(0) {
					waitDur = wait
				}
			}

			output, err := eksAPI.DescribeFargateProfile(&eks.DescribeFargateProfileInput{
				ClusterName:        aws.String(clusterName),
				FargateProfileName: aws.String(profileName),
			})
			if err != nil {
				if IsProfileDeleted(err) {
					if desiredStatus == FargateProfileStatusDELETEDORNOTEXIST {
						lg.Info("fargate profile is already deleted as desired; exiting", zap.Error(err))
						ch <- FargateProfileStatus{FargateProfile: nil, Error: nil}
						close(ch)
						return
					}

					lg.Warn("fargate profile does not exist", zap.Error(err))
					lg.Warn("aborting", zap.Error(ctx.Err()))
					ch <- FargateProfileStatus{FargateProfile: nil, Error: err}
					close(ch)
					return
				}

				lg.Warn("describe fargate profile failed; retrying", zap.Error(err))
				ch <- FargateProfileStatus{FargateProfile: nil, Error: err}
				continue
			}

			if output.FargateProfile == nil {
				lg.Warn("expected non-nil fargate profile; retrying")
				ch <- FargateProfileStatus{FargateProfile: nil, Error: fmt.Errorf("unexpected empty response %+v", output.GoString())}
				continue
			}

			fargateProfile := output.FargateProfile
			currentStatus := aws.StringValue(fargateProfile.Status)
			lg.Info("poll",
				zap.String("cluster-name", clusterName),
				zap.String("fargate-name", profileName),
				zap.String("fargate-status", currentStatus),
				zap.String("started", humanize.RelTime(now, time.Now(), "ago", "from now")),
			)
			switch currentStatus {
			case desiredStatus:
				ch <- FargateProfileStatus{FargateProfile: fargateProfile, Error: nil}
				lg.Info("became desired fargate profile status; exiting", zap.String("status", currentStatus))
				close(ch)
				return

			case eks.FargateProfileStatusCreateFailed,
				eks.FargateProfileStatusDeleteFailed:
				ch <- FargateProfileStatus{FargateProfile: fargateProfile, Error: fmt.Errorf("unexpected fargate status %q", currentStatus)}
				close(ch)
				return
			default:
				ch <- FargateProfileStatus{FargateProfile: fargateProfile, Error: nil}
			}

			if first {
				lg.Info("sleeping", zap.Duration("initial-wait", initialWait))
				select {
				case <-ctx.Done():
					lg.Warn("wait aborted", zap.Error(ctx.Err()))
					ch <- FargateProfileStatus{FargateProfile: nil, Error: ctx.Err()}
					close(ch)
					return
				case <-stopc:
					lg.Warn("wait stopped", zap.Error(ctx.Err()))
					ch <- FargateProfileStatus{FargateProfile: nil, Error: errors.New("wait stopped")}
					close(ch)
					return
				case <-time.After(initialWait):
				}
				first = false
			}
		}

		lg.Warn("wait aborted", zap.Error(ctx.Err()))
		ch <- FargateProfileStatus{FargateProfile: nil, Error: ctx.Err()}
		close(ch)
		return
	}()
	return ch
}

// IsProfileDeleted returns true if error from EKS API indicates that
// the EKS fargate profile has already been deleted.
func IsProfileDeleted(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := err.(awserr.Error)
	if ok && awsErr.Code() == "ResourceNotFoundException" {
		return true
	}

	return strings.Contains(err.Error(), " not found ")
}

const ll = "0123456789abcdefghijklmnopqrstuvwxyz"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = ll[rand.Intn(len(ll))]
	}
	return string(b)
}
