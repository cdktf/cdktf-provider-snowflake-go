// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package objectparameter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObjectParameterConfig struct {
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
	// Name of object parameter. Valid values are those in [object parameters](https://docs.snowflake.com/en/sql-reference/parameters.html#object-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.89.0/docs/resources/object_parameter#key ObjectParameter#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of object parameter, as a string. Constraints are the same as those for the parameters in Snowflake documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.89.0/docs/resources/object_parameter#value ObjectParameter#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.89.0/docs/resources/object_parameter#id ObjectParameter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// object_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.89.0/docs/resources/object_parameter#object_identifier ObjectParameter#object_identifier}
	ObjectIdentifier interface{} `field:"optional" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Type of object to which the parameter applies.
	//
	// Valid values are those in [object types](https://docs.snowflake.com/en/sql-reference/parameters.html#object-types). If no value is provided, then the resource will default to setting the object parameter at account level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.89.0/docs/resources/object_parameter#object_type ObjectParameter#object_type}
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
	// If true, the object parameter will be set on the account level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.89.0/docs/resources/object_parameter#on_account ObjectParameter#on_account}
	OnAccount interface{} `field:"optional" json:"onAccount" yaml:"onAccount"`
}

