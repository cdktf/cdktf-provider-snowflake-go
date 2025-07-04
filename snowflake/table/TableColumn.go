// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package table


type TableColumn struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/table#name Table#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Column type, e.g. VARIANT. For a full list of column types, see [Summary of Data Types](https://docs.snowflake.com/en/sql-reference/intro-summary-data-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/table#type Table#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// (Default: ``) Column collation, e.g. utf8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/table#collate Table#collate}
	Collate *string `field:"optional" json:"collate" yaml:"collate"`
	// (Default: ``) Column comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/table#comment Table#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/table#default Table#default}
	Default *TableColumnDefault `field:"optional" json:"default" yaml:"default"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/table#identity Table#identity}
	Identity *TableColumnIdentity `field:"optional" json:"identity" yaml:"identity"`
	// (Default: ``) Masking policy to apply on column. It has to be a fully qualified name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/table#masking_policy Table#masking_policy}
	MaskingPolicy *string `field:"optional" json:"maskingPolicy" yaml:"maskingPolicy"`
	// (Default: `true`) Whether this column can contain null values.
	//
	// **Note**: Depending on your Snowflake version, the default value will not suffice if this column is used in a primary key constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/table#nullable Table#nullable}
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
}

