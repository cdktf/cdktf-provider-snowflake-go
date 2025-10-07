// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tablecolumnmaskingpolicyapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TableColumnMaskingPolicyApplicationConfig struct {
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
	// The column to apply the masking policy to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/table_column_masking_policy_application#column TableColumnMaskingPolicyApplication#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// Fully qualified name (`database.schema.policyname`) of the policy to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/table_column_masking_policy_application#masking_policy TableColumnMaskingPolicyApplication#masking_policy}
	MaskingPolicy *string `field:"required" json:"maskingPolicy" yaml:"maskingPolicy"`
	// The fully qualified name (`database.schema.table`) of the table to apply the masking policy to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/table_column_masking_policy_application#table TableColumnMaskingPolicyApplication#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/table_column_masking_policy_application#id TableColumnMaskingPolicyApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/table_column_masking_policy_application#timeouts TableColumnMaskingPolicyApplication#timeouts}
	Timeouts *TableColumnMaskingPolicyApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

