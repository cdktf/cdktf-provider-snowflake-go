// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view


type ViewColumnProjectionPolicy struct {
	// Specifies the projection policy to set on a column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#policy_name View#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

