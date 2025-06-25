// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view


type ViewColumn struct {
	// Specifies affected column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/view#column_name View#column_name}
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Specifies a comment for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/view#comment View#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// masking_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/view#masking_policy View#masking_policy}
	MaskingPolicy *ViewColumnMaskingPolicy `field:"optional" json:"maskingPolicy" yaml:"maskingPolicy"`
	// projection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/view#projection_policy View#projection_policy}
	ProjectionPolicy *ViewColumnProjectionPolicy `field:"optional" json:"projectionPolicy" yaml:"projectionPolicy"`
}

