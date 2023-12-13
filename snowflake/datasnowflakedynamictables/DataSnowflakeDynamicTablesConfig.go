// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakedynamictables

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeDynamicTablesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/data-sources/dynamic_tables#id DataSnowflakeDynamicTables#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/data-sources/dynamic_tables#in DataSnowflakeDynamicTables#in}
	In *DataSnowflakeDynamicTablesIn `field:"optional" json:"in" yaml:"in"`
	// like block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/data-sources/dynamic_tables#like DataSnowflakeDynamicTables#like}
	Like *DataSnowflakeDynamicTablesLike `field:"optional" json:"like" yaml:"like"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/data-sources/dynamic_tables#limit DataSnowflakeDynamicTables#limit}
	Limit *DataSnowflakeDynamicTablesLimit `field:"optional" json:"limit" yaml:"limit"`
	// Optionally filters the command output based on the characters that appear at the beginning of the object name.
	//
	// The string is case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/data-sources/dynamic_tables#starts_with DataSnowflakeDynamicTables#starts_with}
	StartsWith *string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

