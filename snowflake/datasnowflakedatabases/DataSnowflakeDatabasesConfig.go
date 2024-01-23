// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakedatabases

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeDatabasesConfig struct {
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
	// Optionally includes dropped databases that have not yet been purged The output also includes an additional `dropped_on` column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.1/docs/data-sources/databases#history DataSnowflakeDatabases#history}
	History interface{} `field:"optional" json:"history" yaml:"history"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.1/docs/data-sources/databases#id DataSnowflakeDatabases#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optionally filters the databases by a pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.1/docs/data-sources/databases#pattern DataSnowflakeDatabases#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Optionally filters the databases by a pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.1/docs/data-sources/databases#starts_with DataSnowflakeDatabases#starts_with}
	StartsWith *string `field:"optional" json:"startsWith" yaml:"startsWith"`
	// Optionally returns only the columns `created_on` and `name` in the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.1/docs/data-sources/databases#terse DataSnowflakeDatabases#terse}
	Terse interface{} `field:"optional" json:"terse" yaml:"terse"`
}

