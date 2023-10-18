// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package table


type TableColumn struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/table#name Table#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Column type, e.g. VARIANT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/table#type Table#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Column comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/table#comment Table#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/table#default Table#default}
	Default *TableColumnDefault `field:"optional" json:"default" yaml:"default"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/table#identity Table#identity}
	Identity *TableColumnIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Masking policy to apply on column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/table#masking_policy Table#masking_policy}
	MaskingPolicy *string `field:"optional" json:"maskingPolicy" yaml:"maskingPolicy"`
	// Whether this column can contain null values.
	//
	// **Note**: Depending on your Snowflake version, the default value will not suffice if this column is used in a primary key constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/table#nullable Table#nullable}
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
}

