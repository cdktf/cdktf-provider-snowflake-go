// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rowaccesspolicy


type RowAccessPolicyArgument struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs/resources/row_access_policy#name RowAccessPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type. VECTOR data types are not yet supported. For more information about data types, check [Snowflake docs](https://docs.snowflake.com/en/sql-reference/intro-summary-data-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs/resources/row_access_policy#type RowAccessPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

