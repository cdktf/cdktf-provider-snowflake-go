// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationIntegrationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#name NotificationIntegration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// AWS IAM role ARN for notification integration to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#aws_sns_role_arn NotificationIntegration#aws_sns_role_arn}
	AwsSnsRoleArn *string `field:"optional" json:"awsSnsRoleArn" yaml:"awsSnsRoleArn"`
	// AWS SNS Topic ARN for notification integration to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#aws_sns_topic_arn NotificationIntegration#aws_sns_topic_arn}
	AwsSnsTopicArn *string `field:"optional" json:"awsSnsTopicArn" yaml:"awsSnsTopicArn"`
	// AWS SQS queue ARN for notification integration to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#aws_sqs_arn NotificationIntegration#aws_sqs_arn}
	AwsSqsArn *string `field:"optional" json:"awsSqsArn" yaml:"awsSqsArn"`
	// AWS IAM role ARN for notification integration to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#aws_sqs_role_arn NotificationIntegration#aws_sqs_role_arn}
	AwsSqsRoleArn *string `field:"optional" json:"awsSqsRoleArn" yaml:"awsSqsRoleArn"`
	// The queue ID for the Azure Queue Storage queue created for Event Grid notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#azure_storage_queue_primary_uri NotificationIntegration#azure_storage_queue_primary_uri}
	AzureStorageQueuePrimaryUri *string `field:"optional" json:"azureStorageQueuePrimaryUri" yaml:"azureStorageQueuePrimaryUri"`
	// The ID of the Azure Active Directory tenant used for identity management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#azure_tenant_id NotificationIntegration#azure_tenant_id}
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// A comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#comment NotificationIntegration#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Direction of the cloud messaging with respect to Snowflake (required only for error notifications).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#direction NotificationIntegration#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#enabled NotificationIntegration#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The subscription id that Snowflake will listen to when using the GCP_PUBSUB provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#gcp_pubsub_subscription_name NotificationIntegration#gcp_pubsub_subscription_name}
	GcpPubsubSubscriptionName *string `field:"optional" json:"gcpPubsubSubscriptionName" yaml:"gcpPubsubSubscriptionName"`
	// The topic id that Snowflake will use to push notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#gcp_pubsub_topic_name NotificationIntegration#gcp_pubsub_topic_name}
	GcpPubsubTopicName *string `field:"optional" json:"gcpPubsubTopicName" yaml:"gcpPubsubTopicName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#id NotificationIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The third-party cloud message queuing service (e.g. AZURE_STORAGE_QUEUE, AWS_SQS, AWS_SNS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#notification_provider NotificationIntegration#notification_provider}
	NotificationProvider *string `field:"optional" json:"notificationProvider" yaml:"notificationProvider"`
	// A type of integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/notification_integration#type NotificationIntegration#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

