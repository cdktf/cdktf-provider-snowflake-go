// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package table


type TableColumnIdentity struct {
	// The number to start incrementing at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/table#start_num Table#start_num}
	StartNum *float64 `field:"optional" json:"startNum" yaml:"startNum"`
	// Step size to increment by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/table#step_num Table#step_num}
	StepNum *float64 `field:"optional" json:"stepNum" yaml:"stepNum"`
}

