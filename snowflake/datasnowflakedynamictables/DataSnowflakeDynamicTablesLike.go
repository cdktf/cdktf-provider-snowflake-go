// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakedynamictables


type DataSnowflakeDynamicTablesLike struct {
	// Filters the command output by object name.
	//
	// The filter uses case-insensitive pattern matching with support for SQL wildcard characters (% and _).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/data-sources/dynamic_tables#pattern DataSnowflakeDynamicTables#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
}

