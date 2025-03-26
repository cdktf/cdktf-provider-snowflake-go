// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeschemas


type DataSnowflakeSchemasLimit struct {
	// The maximum number of rows to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/data-sources/schemas#rows DataSnowflakeSchemas#rows}
	Rows *float64 `field:"required" json:"rows" yaml:"rows"`
	// Specifies a **case-sensitive** pattern that is used to match object name.
	//
	// After the first match, the limit on the number of rows will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/data-sources/schemas#from DataSnowflakeSchemas#from}
	From *string `field:"optional" json:"from" yaml:"from"`
}

