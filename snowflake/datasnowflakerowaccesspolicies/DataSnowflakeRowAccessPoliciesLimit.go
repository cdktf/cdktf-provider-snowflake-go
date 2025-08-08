// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakerowaccesspolicies


type DataSnowflakeRowAccessPoliciesLimit struct {
	// The maximum number of rows to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/data-sources/row_access_policies#rows DataSnowflakeRowAccessPolicies#rows}
	Rows *float64 `field:"required" json:"rows" yaml:"rows"`
	// Specifies a **case-sensitive** pattern that is used to match object name.
	//
	// After the first match, the limit on the number of rows will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/data-sources/row_access_policies#from DataSnowflakeRowAccessPolicies#from}
	From *string `field:"optional" json:"from" yaml:"from"`
}

