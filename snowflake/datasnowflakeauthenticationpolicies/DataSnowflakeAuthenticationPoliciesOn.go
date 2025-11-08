// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeauthenticationpolicies


type DataSnowflakeAuthenticationPoliciesOn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/data-sources/authentication_policies#account DataSnowflakeAuthenticationPolicies#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the specified user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/data-sources/authentication_policies#user DataSnowflakeAuthenticationPolicies#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

