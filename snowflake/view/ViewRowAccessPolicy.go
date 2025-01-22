// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view


type ViewRowAccessPolicy struct {
	// Defines which columns are affected by the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/view#on View#on}
	On *[]*string `field:"required" json:"on" yaml:"on"`
	// Row access policy name. For more information about this resource, see [docs](./row_access_policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/view#policy_name View#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

