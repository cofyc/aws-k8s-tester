
```
AWS_K8S_TESTER_EKS_CONFIG_PATH | *eksconfig.Config.ConfigPath | string | read-only "false"
AWS_K8S_TESTER_EKS_KUBECTL_COMMANDS_OUTPUT_PATH | *eksconfig.Config.KubectlCommandsOutputPath | string | read-only "false"
AWS_K8S_TESTER_EKS_REMOTE_ACCESS_COMMANDS_OUTPUT_PATH | *eksconfig.Config.RemoteAccessCommandsOutputPath | string | read-only "false"
AWS_K8S_TESTER_EKS_REGION | *eksconfig.Config.Region | string | read-only "false"
AWS_K8S_TESTER_EKS_NAME | *eksconfig.Config.Name | string | read-only "false"
AWS_K8S_TESTER_EKS_LOG_LEVEL | *eksconfig.Config.LogLevel | string | read-only "false"
AWS_K8S_TESTER_EKS_LOG_OUTPUTS | *eksconfig.Config.LogOutputs | []string | read-only "false"
AWS_K8S_TESTER_EKS_AWS_CLI_PATH | *eksconfig.Config.AWSCLIPath | string | read-only "false"
AWS_K8S_TESTER_EKS_KUBECTL_PATH | *eksconfig.Config.KubectlPath | string | read-only "false"
AWS_K8S_TESTER_EKS_KUBECTL_DOWNLOAD_URL | *eksconfig.Config.KubectlDownloadURL | string | read-only "false"
AWS_K8S_TESTER_EKS_KUBECONFIG_PATH | *eksconfig.Config.KubeConfigPath | string | read-only "false"
AWS_K8S_TESTER_EKS_AWS_IAM_AUTHENTICATOR_PATH | *eksconfig.Config.AWSIAMAuthenticatorPath | string | read-only "false"
AWS_K8S_TESTER_EKS_AWS_IAM_AUTHENTICATOR_DOWNLOAD_URL | *eksconfig.Config.AWSIAMAuthenticatorDownloadURL | string | read-only "false"
AWS_K8S_TESTER_EKS_ON_FAILURE_DELETE | *eksconfig.Config.OnFailureDelete | bool | read-only "false"
AWS_K8S_TESTER_EKS_ON_FAILURE_DELETE_WAIT_SECONDS | *eksconfig.Config.OnFailureDeleteWaitSeconds | uint64 | read-only "false"
AWS_K8S_TESTER_EKS_COMMAND_AFTER_CREATE_CLUSTER | *eksconfig.Config.CommandAfterCreateCluster | string | read-only "false"
AWS_K8S_TESTER_EKS_COMMAND_AFTER_CREATE_CLUSTER_OUTPUT_PATH | *eksconfig.Config.CommandAfterCreateClusterOutputPath | string | read-only "true"
AWS_K8S_TESTER_EKS_COMMAND_AFTER_CREATE_ADD_ONS | *eksconfig.Config.CommandAfterCreateAddOns | string | read-only "false"
AWS_K8S_TESTER_EKS_COMMAND_AFTER_CREATE_ADD_ONS_OUTPUT_PATH | *eksconfig.Config.CommandAfterCreateAddOnsOutputPath | string | read-only "true"
AWS_K8S_TESTER_EKS_S3_BUCKET_NAME | *eksconfig.Config.S3BucketName | string | read-only "false"
AWS_K8S_TESTER_EKS_S3_BUCKET_CREATE | *eksconfig.Config.S3BucketCreate | bool | read-only "false"
AWS_K8S_TESTER_EKS_S3_BUCKET_LIFECYCLE_EXPIRATION_DAYS | *eksconfig.Config.S3BucketLifecycleExpirationDays | int64 | read-only "false"
AWS_K8S_TESTER_EKS_REMOTE_ACCESS_KEY_CREATE | *eksconfig.Config.RemoteAccessKeyCreate | bool | read-only "false"
AWS_K8S_TESTER_EKS_REMOTE_ACCESS_KEY_NAME | *eksconfig.Config.RemoteAccessKeyName | string | read-only "false"
AWS_K8S_TESTER_EKS_REMOTE_ACCESS_PRIVATE_KEY_PATH | *eksconfig.Config.RemoteAccessPrivateKeyPath | string | read-only "false"

AWS_K8S_TESTER_EKS_PARAMETERS_ROLE_NAME | *eksconfig.Parameters.RoleName | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_ROLE_CREATE | *eksconfig.Parameters.RoleCreate | bool | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_ROLE_ARN | *eksconfig.Parameters.RoleARN | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_ROLE_SERVICE_PRINCIPALS | *eksconfig.Parameters.RoleServicePrincipals | []string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_ROLE_MANAGED_POLICY_ARNS | *eksconfig.Parameters.RoleManagedPolicyARNs | []string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_ROLE_CFN_STACK_ID | *eksconfig.Parameters.RoleCFNStackID | string | read-only "true"
AWS_K8S_TESTER_EKS_PARAMETERS_TAGS | *eksconfig.Parameters.Tags | map[string]string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_REQUEST_HEADER_KEY | *eksconfig.Parameters.RequestHeaderKey | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_REQUEST_HEADER_VALUE | *eksconfig.Parameters.RequestHeaderValue | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_RESOLVER_URL | *eksconfig.Parameters.ResolverURL | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_SIGNING_NAME | *eksconfig.Parameters.SigningName | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_VPC_CREATE | *eksconfig.Parameters.VPCCreate | bool | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_VPC_ID | *eksconfig.Parameters.VPCID | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_VPC_CFN_STACK_ID | *eksconfig.Parameters.VPCCFNStackID | string | read-only "true"
AWS_K8S_TESTER_EKS_PARAMETERS_VPC_CIDR | *eksconfig.Parameters.VPCCIDR | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_PUBLIC_SUBNET_CIDR_1 | *eksconfig.Parameters.PublicSubnetCIDR1 | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_PUBLIC_SUBNET_CIDR_2 | *eksconfig.Parameters.PublicSubnetCIDR2 | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_PUBLIC_SUBNET_CIDR_3 | *eksconfig.Parameters.PublicSubnetCIDR3 | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_PRIVATE_SUBNET_CIDR_1 | *eksconfig.Parameters.PrivateSubnetCIDR1 | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_PRIVATE_SUBNET_CIDR_2 | *eksconfig.Parameters.PrivateSubnetCIDR2 | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_PUBLIC_SUBNET_IDS | *eksconfig.Parameters.PublicSubnetIDs | []string | read-only "true"
AWS_K8S_TESTER_EKS_PARAMETERS_PRIVATE_SUBNET_IDS | *eksconfig.Parameters.PrivateSubnetIDs | []string | read-only "true"
AWS_K8S_TESTER_EKS_PARAMETERS_CONTROL_PLANE_SECURITY_GROUP_ID | *eksconfig.Parameters.ControlPlaneSecurityGroupID | string | read-only "true"
AWS_K8S_TESTER_EKS_PARAMETERS_VERSION | *eksconfig.Parameters.Version | string | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_VERSION_VALUE | *eksconfig.Parameters.VersionValue | float64 | read-only "true"
AWS_K8S_TESTER_EKS_PARAMETERS_ENCRYPTION_CMK_CREATE | *eksconfig.Parameters.EncryptionCMKCreate | bool | read-only "false"
AWS_K8S_TESTER_EKS_PARAMETERS_ENCRYPTION_CMK_ARN | *eksconfig.Parameters.EncryptionCMKARN | string | read-only "false"

AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ENABLE | *eksconfig.AddOnNodeGroups.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_CREATED | *eksconfig.AddOnNodeGroups.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_FETCH_LOGS | *eksconfig.AddOnNodeGroups.FetchLogs | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ROLE_NAME | *eksconfig.AddOnNodeGroups.RoleName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ROLE_CREATE | *eksconfig.AddOnNodeGroups.RoleCreate | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ROLE_ARN | *eksconfig.AddOnNodeGroups.RoleARN | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ROLE_SERVICE_PRINCIPALS | *eksconfig.AddOnNodeGroups.RoleServicePrincipals | []string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ROLE_MANAGED_POLICY_ARNS | *eksconfig.AddOnNodeGroups.RoleManagedPolicyARNs | []string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ROLE_CFN_STACK_ID | *eksconfig.AddOnNodeGroups.RoleCFNStackID | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_NODE_GROUP_SECURITY_GROUP_ID | *eksconfig.AddOnNodeGroups.NodeGroupSecurityGroupID | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_NODE_GROUP_SECURITY_GROUP_CFN_STACK_ID | *eksconfig.AddOnNodeGroups.NodeGroupSecurityGroupCFNStackID | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_LOGS_DIR | *eksconfig.AddOnNodeGroups.LogsDir | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ASGS | *eksconfig.AddOnNodeGroups.ASGs | map[string]ec2config.ASG | read-only "false"

AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_ENABLE | *eksconfig.AddOnManagedNodeGroups.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_CREATED | *eksconfig.AddOnManagedNodeGroups.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_CREATE_TOOK | *eksconfig.AddOnManagedNodeGroups.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_CREATE_TOOK_STRING | *eksconfig.AddOnManagedNodeGroups.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_DELETE_TOOK | *eksconfig.AddOnManagedNodeGroups.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_DELETE_TOOK_STRING | *eksconfig.AddOnManagedNodeGroups.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_FETCH_LOGS | *eksconfig.AddOnManagedNodeGroups.FetchLogs | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_ROLE_NAME | *eksconfig.AddOnManagedNodeGroups.RoleName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_ROLE_CREATE | *eksconfig.AddOnManagedNodeGroups.RoleCreate | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_ROLE_ARN | *eksconfig.AddOnManagedNodeGroups.RoleARN | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_ROLE_SERVICE_PRINCIPALS | *eksconfig.AddOnManagedNodeGroups.RoleServicePrincipals | []string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_ROLE_MANAGED_POLICY_ARNS | *eksconfig.AddOnManagedNodeGroups.RoleManagedPolicyARNs | []string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_ROLE_CFN_STACK_ID | *eksconfig.AddOnManagedNodeGroups.RoleCFNStackID | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_REQUEST_HEADER_KEY | *eksconfig.AddOnManagedNodeGroups.RequestHeaderKey | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_REQUEST_HEADER_VALUE | *eksconfig.AddOnManagedNodeGroups.RequestHeaderValue | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_RESOLVER_URL | *eksconfig.AddOnManagedNodeGroups.ResolverURL | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_SIGNING_NAME | *eksconfig.AddOnManagedNodeGroups.SigningName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_LOGS_DIR | *eksconfig.AddOnManagedNodeGroups.LogsDir | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_MANAGED_NODE_GROUPS_MNGS | *eksconfig.AddOnManagedNodeGroups.MNGs | map[string]eksconfig.MNG | read-only "false"

AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_ENABLE | *eksconfig.AddOnNLBHelloWorld.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_CREATED | *eksconfig.AddOnNLBHelloWorld.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_CREATE_TOOK | *eksconfig.AddOnNLBHelloWorld.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_CREATE_TOOK_STRING | *eksconfig.AddOnNLBHelloWorld.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_DELETE_TOOK | *eksconfig.AddOnNLBHelloWorld.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_DELETE_TOOK_STRING | *eksconfig.AddOnNLBHelloWorld.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_DEPLOYMENT_REPLICAS | *eksconfig.AddOnNLBHelloWorld.DeploymentReplicas | int32 | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_NAMESPACE | *eksconfig.AddOnNLBHelloWorld.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_NLB_ARN | *eksconfig.AddOnNLBHelloWorld.NLBARN | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_NLB_NAME | *eksconfig.AddOnNLBHelloWorld.NLBName | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_NLB_HELLO_WORLD_URL | *eksconfig.AddOnNLBHelloWorld.URL | string | read-only "true"

AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_ENABLE | *eksconfig.AddOnALB2048.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_CREATED | *eksconfig.AddOnALB2048.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_CREATE_TOOK | *eksconfig.AddOnALB2048.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_CREATE_TOOK_STRING | *eksconfig.AddOnALB2048.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_DELETE_TOOK | *eksconfig.AddOnALB2048.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_DELETE_TOOK_STRING | *eksconfig.AddOnALB2048.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_NAMESPACE | *eksconfig.AddOnALB2048.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_DEPLOYMENT_REPLICAS_ALB | *eksconfig.AddOnALB2048.DeploymentReplicasALB | int32 | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_DEPLOYMENT_REPLICAS_2048 | *eksconfig.AddOnALB2048.DeploymentReplicas2048 | int32 | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_ALB_ARN | *eksconfig.AddOnALB2048.ALBARN | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_ALB_NAME | *eksconfig.AddOnALB2048.ALBName | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_ALB_2048_URL | *eksconfig.AddOnALB2048.URL | string | read-only "true"

AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_ENABLE | *eksconfig.AddOnJobPi.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_CREATED | *eksconfig.AddOnJobPi.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_CREATE_TOOK | *eksconfig.AddOnJobPi.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_CREATE_TOOK_STRING | *eksconfig.AddOnJobPi.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_DELETE_TOOK | *eksconfig.AddOnJobPi.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_DELETE_TOOK_STRING | *eksconfig.AddOnJobPi.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_NAMESPACE | *eksconfig.AddOnJobPi.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_COMPLETES | *eksconfig.AddOnJobPi.Completes | int | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_PI_PARALLELS | *eksconfig.AddOnJobPi.Parallels | int | read-only "false"

AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_ENABLE | *eksconfig.AddOnJobEcho.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_CREATED | *eksconfig.AddOnJobEcho.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_CREATE_TOOK | *eksconfig.AddOnJobEcho.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_CREATE_TOOK_STRING | *eksconfig.AddOnJobEcho.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_DELETE_TOOK | *eksconfig.AddOnJobEcho.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_DELETE_TOOK_STRING | *eksconfig.AddOnJobEcho.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_NAMESPACE | *eksconfig.AddOnJobEcho.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_COMPLETES | *eksconfig.AddOnJobEcho.Completes | int | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_PARALLELS | *eksconfig.AddOnJobEcho.Parallels | int | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_JOB_ECHO_ECHO_SIZE | *eksconfig.AddOnJobEcho.EchoSize | int | read-only "false"

AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_ENABLE | *eksconfig.AddOnCronJob.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_CREATED | *eksconfig.AddOnCronJob.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_CREATE_TOOK | *eksconfig.AddOnCronJob.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_CREATE_TOOK_STRING | *eksconfig.AddOnCronJob.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_DELETE_TOOK | *eksconfig.AddOnCronJob.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_DELETE_TOOK_STRING | *eksconfig.AddOnCronJob.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_NAMESPACE | *eksconfig.AddOnCronJob.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_SCHEDULE | *eksconfig.AddOnCronJob.Schedule | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_COMPLETES | *eksconfig.AddOnCronJob.Completes | int | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_PARALLELS | *eksconfig.AddOnCronJob.Parallels | int | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_SUCCESSFUL_JOBS_HISTORY_LIMIT | *eksconfig.AddOnCronJob.SuccessfulJobsHistoryLimit | int32 | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_FAILED_JOBS_HISTORY_LIMIT | *eksconfig.AddOnCronJob.FailedJobsHistoryLimit | int32 | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_CRON_JOB_ECHO_SIZE | *eksconfig.AddOnCronJob.EchoSize | int | read-only "false"

AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_ENABLE | *eksconfig.AddOnSecrets.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_CREATED | *eksconfig.AddOnSecrets.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_CREATE_TOOK | *eksconfig.AddOnSecrets.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_CREATE_TOOK_STRING | *eksconfig.AddOnSecrets.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_DELETE_TOOK | *eksconfig.AddOnSecrets.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_DELETE_TOOK_STRING | *eksconfig.AddOnSecrets.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_NAMESPACE | *eksconfig.AddOnSecrets.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_OBJECTS | *eksconfig.AddOnSecrets.Objects | int | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_SIZE | *eksconfig.AddOnSecrets.Size | int | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_SECRET_QPS | *eksconfig.AddOnSecrets.SecretQPS | uint | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_SECRET_BURST | *eksconfig.AddOnSecrets.SecretBurst | uint | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_CREATED_SECRET_NAMES | *eksconfig.AddOnSecrets.CreatedSecretNames | []string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_POD_QPS | *eksconfig.AddOnSecrets.PodQPS | uint | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_POD_BURST | *eksconfig.AddOnSecrets.PodBurst | uint | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_CREATED_POD_NAMES | *eksconfig.AddOnSecrets.CreatedPodNames | []string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_WRITES_RESULT_PATH | *eksconfig.AddOnSecrets.WritesResultPath | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_SECRETS_READS_RESULT_PATH | *eksconfig.AddOnSecrets.ReadsResultPath | string | read-only "false"

AWS_K8S_TESTER_EKS_ADD_ON_IRSA_ENABLE | *eksconfig.AddOnIRSA.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_CREATED | *eksconfig.AddOnIRSA.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_CREATE_TOOK | *eksconfig.AddOnIRSA.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_CREATE_TOOK_STRING | *eksconfig.AddOnIRSA.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_DELETE_TOOK | *eksconfig.AddOnIRSA.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_DELETE_TOOK_STRING | *eksconfig.AddOnIRSA.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_NAMESPACE | *eksconfig.AddOnIRSA.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_ROLE_NAME | *eksconfig.AddOnIRSA.RoleName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_ROLE_ARN | *eksconfig.AddOnIRSA.RoleARN | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_ROLE_MANAGED_POLICY_ARNS | *eksconfig.AddOnIRSA.RoleManagedPolicyARNs | []string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_ROLE_CFN_STACK_ID | *eksconfig.AddOnIRSA.RoleCFNStackID | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_SERVICE_ACCOUNT_NAME | *eksconfig.AddOnIRSA.ServiceAccountName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_CONFIG_MAP_NAME | *eksconfig.AddOnIRSA.ConfigMapName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_CONFIG_MAP_SCRIPT_FILE_NAME | *eksconfig.AddOnIRSA.ConfigMapScriptFileName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_S3_KEY | *eksconfig.AddOnIRSA.S3Key | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_DEPLOYMENT_NAME | *eksconfig.AddOnIRSA.DeploymentName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_DEPLOYMENT_REPLICAS | *eksconfig.AddOnIRSA.DeploymentReplicas | int32 | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_DEPLOYMENT_RESULT_PATH | *eksconfig.AddOnIRSA.DeploymentResultPath | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_DEPLOYMENT_TOOK | *eksconfig.AddOnIRSA.DeploymentTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_IRSA_DEPLOYMENT_TOOK_STRING | *eksconfig.AddOnIRSA.DeploymentTookString | string | read-only "true"

AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_ENABLE | *eksconfig.AddOnFargate.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_CREATED | *eksconfig.AddOnFargate.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_CREATE_TOOK | *eksconfig.AddOnFargate.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_CREATE_TOOK_STRING | *eksconfig.AddOnFargate.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_DELETE_TOOK | *eksconfig.AddOnFargate.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_DELETE_TOOK_STRING | *eksconfig.AddOnFargate.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_NAMESPACE | *eksconfig.AddOnFargate.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_ROLE_NAME | *eksconfig.AddOnFargate.RoleName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_ROLE_CREATE | *eksconfig.AddOnFargate.RoleCreate | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_ROLE_ARN | *eksconfig.AddOnFargate.RoleARN | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_ROLE_SERVICE_PRINCIPALS | *eksconfig.AddOnFargate.RoleServicePrincipals | []string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_ROLE_MANAGED_POLICY_ARNS | *eksconfig.AddOnFargate.RoleManagedPolicyARNs | []string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_ROLE_CFN_STACK_ID | *eksconfig.AddOnFargate.RoleCFNStackID | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_PROFILE_NAME | *eksconfig.AddOnFargate.ProfileName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_SECRET_NAME | *eksconfig.AddOnFargate.SecretName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_POD_NAME | *eksconfig.AddOnFargate.PodName | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_FARGATE_CONTAINER_NAME | *eksconfig.AddOnFargate.ContainerName | string | read-only "false"

AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_ENABLE | *eksconfig.AddOnAppMesh.Enable | bool | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_NAMESPACE | *eksconfig.AddOnAppMesh.Namespace | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_CONTROLLER_IMAGE | *eksconfig.AddOnAppMesh.ControllerImage | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_INJECTOR_IMAGE | *eksconfig.AddOnAppMesh.InjectorImage | string | read-only "false"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_CREATED | *eksconfig.AddOnAppMesh.Created | bool | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_CREATE_TOOK | *eksconfig.AddOnAppMesh.CreateTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_CREATE_TOOK_STRING | *eksconfig.AddOnAppMesh.CreateTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_DELETE_TOOK | *eksconfig.AddOnAppMesh.DeleteTook | time.Duration | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_DELETE_TOOK_STRING | *eksconfig.AddOnAppMesh.DeleteTookString | string | read-only "true"
AWS_K8S_TESTER_EKS_ADD_ON_APP_MESH_ADD_ON_CFN_STACK_ARN | *eksconfig.AddOnAppMesh.AddOnCFNStackARN | string | read-only "true"
```
