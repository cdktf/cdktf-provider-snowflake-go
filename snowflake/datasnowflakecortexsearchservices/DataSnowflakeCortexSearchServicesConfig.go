// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakecortexsearchservices

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeCortexSearchServicesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/data-sources/cortex_search_services#id DataSnowflakeCortexSearchServices#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/data-sources/cortex_search_services#in DataSnowflakeCortexSearchServices#in}
	In *DataSnowflakeCortexSearchServicesIn `field:"optional" json:"in" yaml:"in"`
	// Filters the output with **case-insensitive** pattern, with support for SQL wildcard characters (`%` and `_`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/data-sources/cortex_search_services#like DataSnowflakeCortexSearchServices#like}
	Like *string `field:"optional" json:"like" yaml:"like"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/data-sources/cortex_search_services#limit DataSnowflakeCortexSearchServices#limit}
	Limit *DataSnowflakeCortexSearchServicesLimit `field:"optional" json:"limit" yaml:"limit"`
	// Filters the output with **case-sensitive** characters indicating the beginning of the object name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/data-sources/cortex_search_services#starts_with DataSnowflakeCortexSearchServices#starts_with}
	StartsWith *string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

