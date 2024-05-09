// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package unsafeexecute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UnsafeExecuteConfig struct {
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
	// SQL statement to execute. Forces recreation of resource on change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/unsafe_execute#execute UnsafeExecute#execute}
	Execute *string `field:"required" json:"execute" yaml:"execute"`
	// SQL statement to revert the execute statement. Invoked when resource is being destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/unsafe_execute#revert UnsafeExecute#revert}
	Revert *string `field:"required" json:"revert" yaml:"revert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/unsafe_execute#id UnsafeExecute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional SQL statement to do a read. Invoked after creation and every time it is changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/unsafe_execute#query UnsafeExecute#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

