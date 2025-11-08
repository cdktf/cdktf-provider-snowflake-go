// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package table


type TablePrimaryKey struct {
	// Columns to use in primary key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/table#keys Table#keys}
	Keys *[]*string `field:"required" json:"keys" yaml:"keys"`
	// Name of constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/table#name Table#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

