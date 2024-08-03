// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamictable


type DynamicTableTargetLag struct {
	// Specifies whether the target lag time is downstream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/resources/dynamic_table#downstream DynamicTable#downstream}
	Downstream interface{} `field:"optional" json:"downstream" yaml:"downstream"`
	// Specifies the maximum target lag time for the dynamic table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/resources/dynamic_table#maximum_duration DynamicTable#maximum_duration}
	MaximumDuration *string `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
}

