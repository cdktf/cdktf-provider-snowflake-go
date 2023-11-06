// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sessionparameter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SessionParameterConfig struct {
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
	// Name of session parameter. Valid values are those in [session parameters](https://docs.snowflake.com/en/sql-reference/parameters.html#session-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/session_parameter#key SessionParameter#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of session parameter, as a string. Constraints are the same as those for the parameters in Snowflake documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/session_parameter#value SessionParameter#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/session_parameter#id SessionParameter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, the session parameter will be set on the account level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/session_parameter#on_account SessionParameter#on_account}
	OnAccount interface{} `field:"optional" json:"onAccount" yaml:"onAccount"`
	// The user to set the session parameter for. Required if on_account is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/session_parameter#user SessionParameter#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

