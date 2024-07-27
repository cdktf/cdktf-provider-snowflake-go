// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipe

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipeConfig struct {
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
	// Specifies the copy statement for the pipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#copy_statement Pipe#copy_statement}
	CopyStatement *string `field:"required" json:"copyStatement" yaml:"copyStatement"`
	// The database in which to create the pipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#database Pipe#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the pipe;
	//
	// must be unique for the database and schema in which the pipe is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#name Pipe#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the pipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#schema Pipe#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies a auto_ingest param for the pipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#auto_ingest Pipe#auto_ingest}
	AutoIngest interface{} `field:"optional" json:"autoIngest" yaml:"autoIngest"`
	// Specifies the Amazon Resource Name (ARN) for the SNS topic for your S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#aws_sns_topic_arn Pipe#aws_sns_topic_arn}
	AwsSnsTopicArn *string `field:"optional" json:"awsSnsTopicArn" yaml:"awsSnsTopicArn"`
	// Specifies a comment for the pipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#comment Pipe#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the name of the notification integration used for error notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#error_integration Pipe#error_integration}
	ErrorIntegration *string `field:"optional" json:"errorIntegration" yaml:"errorIntegration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#id Pipe#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies an integration for the pipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/pipe#integration Pipe#integration}
	Integration *string `field:"optional" json:"integration" yaml:"integration"`
}

