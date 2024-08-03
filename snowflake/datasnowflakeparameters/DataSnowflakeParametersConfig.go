// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeparameters

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeParametersConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/data-sources/parameters#id DataSnowflakeParameters#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If parameter_type is set to "OBJECT" then object_name is the name of the object to display object parameters for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/data-sources/parameters#object_name DataSnowflakeParameters#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// If parameter_type is set to "OBJECT" then object_type is the type of object to display object parameters for.
	//
	// Valid values are any object supported by the IN clause of the [SHOW PARAMETERS](https://docs.snowflake.com/en/sql-reference/sql/show-parameters.html#parameters) statement, including: WAREHOUSE | DATABASE | SCHEMA | TASK | TABLE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/data-sources/parameters#object_type DataSnowflakeParameters#object_type}
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
	// The type of parameter to filter by. Valid values are: "ACCOUNT", "SESSION", "OBJECT".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/data-sources/parameters#parameter_type DataSnowflakeParameters#parameter_type}
	ParameterType *string `field:"optional" json:"parameterType" yaml:"parameterType"`
	// Allows limiting the list of parameters by name using LIKE clause.
	//
	// Refer to [Limiting the List of Parameters by Name](https://docs.snowflake.com/en/sql-reference/parameters.html#limiting-the-list-of-parameters-by-name)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/data-sources/parameters#pattern DataSnowflakeParameters#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// If parameter_type is set to "SESSION" then user is the name of the user to display session parameters for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/data-sources/parameters#user DataSnowflakeParameters#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

