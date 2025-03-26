// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakedynamictables


type DataSnowflakeDynamicTablesLimit struct {
	// The optional FROM 'name_string' subclause effectively serves as a “cursor” for the results.
	//
	// This enables fetching the specified number of rows following the first row whose object name matches the specified string
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/data-sources/dynamic_tables#from DataSnowflakeDynamicTables#from}
	From *string `field:"optional" json:"from" yaml:"from"`
	// Specifies the maximum number of rows to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/data-sources/dynamic_tables#rows DataSnowflakeDynamicTables#rows}
	Rows *float64 `field:"optional" json:"rows" yaml:"rows"`
}

