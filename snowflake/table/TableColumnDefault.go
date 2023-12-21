// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package table


type TableColumnDefault struct {
	// The default constant value for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.82.0/docs/resources/table#constant Table#constant}
	Constant *string `field:"optional" json:"constant" yaml:"constant"`
	// The default expression value for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.82.0/docs/resources/table#expression Table#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The default sequence to use for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.82.0/docs/resources/table#sequence Table#sequence}
	Sequence *string `field:"optional" json:"sequence" yaml:"sequence"`
}

