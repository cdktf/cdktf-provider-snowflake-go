// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tablecolumnmaskingpolicyapplication


type TableColumnMaskingPolicyApplicationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/table_column_masking_policy_application#create TableColumnMaskingPolicyApplication#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/table_column_masking_policy_application#delete TableColumnMaskingPolicyApplication#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/table_column_masking_policy_application#read TableColumnMaskingPolicyApplication#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/table_column_masking_policy_application#update TableColumnMaskingPolicyApplication#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

