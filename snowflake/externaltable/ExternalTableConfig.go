// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externaltable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExternalTableConfig struct {
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
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#column ExternalTable#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The database in which to create the external table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#database ExternalTable#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the file format for the external table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#file_format ExternalTable#file_format}
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// Specifies a location for the external table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#location ExternalTable#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Specifies the identifier for the external table;
	//
	// must be unique for the database and schema in which the externalTable is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#name ExternalTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the external table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#schema ExternalTable#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies whether to automatically refresh the external table metadata once, immediately after the external table is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#auto_refresh ExternalTable#auto_refresh}
	AutoRefresh interface{} `field:"optional" json:"autoRefresh" yaml:"autoRefresh"`
	// Specifies the aws sns topic for the external table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#aws_sns_topic ExternalTable#aws_sns_topic}
	AwsSnsTopic *string `field:"optional" json:"awsSnsTopic" yaml:"awsSnsTopic"`
	// Specifies a comment for the external table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#comment ExternalTable#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies to retain the access permissions from the original table when an external table is recreated using the CREATE OR REPLACE TABLE variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#copy_grants ExternalTable#copy_grants}
	CopyGrants interface{} `field:"optional" json:"copyGrants" yaml:"copyGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#id ExternalTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies any partition columns to evaluate for the external table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#partition_by ExternalTable#partition_by}
	PartitionBy *[]*string `field:"optional" json:"partitionBy" yaml:"partitionBy"`
	// Specifies the file names and/or paths on the external stage to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#pattern ExternalTable#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Specifies weather to refresh when an external table is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#refresh_on_create ExternalTable#refresh_on_create}
	RefreshOnCreate interface{} `field:"optional" json:"refreshOnCreate" yaml:"refreshOnCreate"`
	// Identifies the external table table type. For now, only "delta" for Delta Lake table format is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#table_format ExternalTable#table_format}
	TableFormat *string `field:"optional" json:"tableFormat" yaml:"tableFormat"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/external_table#tag ExternalTable#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

