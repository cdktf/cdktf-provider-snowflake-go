// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package table


type TableColumnIdentity struct {
	// (Default: `1`) The number to start incrementing at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/table#start_num Table#start_num}
	StartNum *float64 `field:"optional" json:"startNum" yaml:"startNum"`
	// (Default: `1`) Step size to increment by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/table#step_num Table#step_num}
	StepNum *float64 `field:"optional" json:"stepNum" yaml:"stepNum"`
}

