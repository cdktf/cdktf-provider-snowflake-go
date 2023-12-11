// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maskingpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaskingPolicyConfig struct {
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
	// The database in which to create the masking policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#database MaskingPolicy#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the SQL expression that transforms the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#masking_expression MaskingPolicy#masking_expression}
	MaskingExpression *string `field:"required" json:"maskingExpression" yaml:"maskingExpression"`
	// Specifies the identifier for the masking policy;
	//
	// must be unique for the database and schema in which the masking policy is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#name MaskingPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the data type to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#return_data_type MaskingPolicy#return_data_type}
	ReturnDataType *string `field:"required" json:"returnDataType" yaml:"returnDataType"`
	// The schema in which to create the masking policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#schema MaskingPolicy#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// signature block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#signature MaskingPolicy#signature}
	Signature *MaskingPolicySignature `field:"required" json:"signature" yaml:"signature"`
	// Specifies a comment for the masking policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#comment MaskingPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether the row access policy or conditional masking policy can reference a column that is already protected by a masking policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#exempt_other_policies MaskingPolicy#exempt_other_policies}
	ExemptOtherPolicies interface{} `field:"optional" json:"exemptOtherPolicies" yaml:"exemptOtherPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#id MaskingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Prevent overwriting a previous masking policy with the same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#if_not_exists MaskingPolicy#if_not_exists}
	IfNotExists interface{} `field:"optional" json:"ifNotExists" yaml:"ifNotExists"`
	// Whether to override a previous masking policy with the same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.0/docs/resources/masking_policy#or_replace MaskingPolicy#or_replace}
	OrReplace interface{} `field:"optional" json:"orReplace" yaml:"orReplace"`
}

