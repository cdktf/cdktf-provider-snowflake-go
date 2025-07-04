// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StageConfig struct {
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
	// The database in which to create the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#database Stage#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the stage;
	//
	// must be unique for the database and schema in which the stage is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#name Stage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#schema Stage#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// A unique ID assigned to the specific stage. The ID has the following format: &lt;snowflakeAccount&gt;_SFCRole=&lt;snowflakeRoleId&gt;_&lt;randomId&gt;
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#aws_external_id Stage#aws_external_id}
	AwsExternalId *string `field:"optional" json:"awsExternalId" yaml:"awsExternalId"`
	// Specifies a comment for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#comment Stage#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the copy options for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#copy_options Stage#copy_options}
	CopyOptions *string `field:"optional" json:"copyOptions" yaml:"copyOptions"`
	// Specifies the credentials for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#credentials Stage#credentials}
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// Specifies the directory settings for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#directory Stage#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Specifies the encryption settings for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#encryption Stage#encryption}
	Encryption *string `field:"optional" json:"encryption" yaml:"encryption"`
	// Specifies the file format for the stage.
	//
	// Specifying the default Snowflake value (e.g. TYPE = CSV) will currently result in a permadiff (check [#2679](https://github.com/snowflakedb/terraform-provider-snowflake/issues/2679)). For now, omit the default values; it will be fixed in the upcoming provider versions. Examples of usage: <b>1. with hardcoding value:</b> `file_format="FORMAT_NAME = DB.SCHEMA.FORMATNAME"` <b>2. from dynamic value:</b> `file_format = "FORMAT_NAME = ${snowflake_file_format.myfileformat.fully_qualified_name}"` <b>3. from expression:</b> `file_format = format("FORMAT_NAME =%s.%s.MYFILEFORMAT", var.db_name, each.value.schema_name)`. Reference: [#265](https://github.com/snowflakedb/terraform-provider-snowflake/issues/265)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#file_format Stage#file_format}
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#id Stage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An AWS IAM user created for your Snowflake account.
	//
	// This user is the same for every external S3 stage created in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#snowflake_iam_user Stage#snowflake_iam_user}
	SnowflakeIamUser *string `field:"optional" json:"snowflakeIamUser" yaml:"snowflakeIamUser"`
	// Specifies the name of the storage integration used to delegate authentication responsibility for external cloud storage to a Snowflake identity and access management (IAM) entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#storage_integration Stage#storage_integration}
	StorageIntegration *string `field:"optional" json:"storageIntegration" yaml:"storageIntegration"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#tag Stage#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#timeouts Stage#timeouts}
	Timeouts *StageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Specifies the URL for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/stage#url Stage#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

