// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeusers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeUsersConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/data-sources/users#id DataSnowflakeUsers#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Filters the output with **case-insensitive** pattern, with support for SQL wildcard characters (`%` and `_`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/data-sources/users#like DataSnowflakeUsers#like}
	Like *string `field:"optional" json:"like" yaml:"like"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/data-sources/users#limit DataSnowflakeUsers#limit}
	Limit *DataSnowflakeUsersLimit `field:"optional" json:"limit" yaml:"limit"`
	// Filters the output with **case-sensitive** characters indicating the beginning of the object name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/data-sources/users#starts_with DataSnowflakeUsers#starts_with}
	StartsWith *string `field:"optional" json:"startsWith" yaml:"startsWith"`
	// (Default: `true`) Runs DESC USER for each user returned by SHOW USERS.
	//
	// The output of describe is saved to the description field. By default this value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/data-sources/users#with_describe DataSnowflakeUsers#with_describe}
	WithDescribe interface{} `field:"optional" json:"withDescribe" yaml:"withDescribe"`
	// (Default: `true`) Runs SHOW PARAMETERS FOR USER for each user returned by SHOW USERS.
	//
	// The output of describe is saved to the parameters field as a map. By default this value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/data-sources/users#with_parameters DataSnowflakeUsers#with_parameters}
	WithParameters interface{} `field:"optional" json:"withParameters" yaml:"withParameters"`
}

