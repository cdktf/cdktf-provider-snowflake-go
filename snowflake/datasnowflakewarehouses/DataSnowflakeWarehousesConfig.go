// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakewarehouses

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeWarehousesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/data-sources/warehouses#id DataSnowflakeWarehouses#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Filters the output with **case-insensitive** pattern, with support for SQL wildcard characters (`%` and `_`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/data-sources/warehouses#like DataSnowflakeWarehouses#like}
	Like *string `field:"optional" json:"like" yaml:"like"`
	// Runs DESC WAREHOUSE for each warehouse returned by SHOW WAREHOUSES.
	//
	// The output of describe is saved to the description field. By default this value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/data-sources/warehouses#with_describe DataSnowflakeWarehouses#with_describe}
	WithDescribe interface{} `field:"optional" json:"withDescribe" yaml:"withDescribe"`
	// Runs SHOW PARAMETERS FOR WAREHOUSE for each warehouse returned by SHOW WAREHOUSES.
	//
	// The output of describe is saved to the parameters field as a map. By default this value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/data-sources/warehouses#with_parameters DataSnowflakeWarehouses#with_parameters}
	WithParameters interface{} `field:"optional" json:"withParameters" yaml:"withParameters"`
}

