// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticationpolicy


type AuthenticationPolicyPatPolicy struct {
	// Specifies the default expiration time (in days) for a programmatic access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#default_expiry_in_days AuthenticationPolicy#default_expiry_in_days}
	DefaultExpiryInDays *float64 `field:"optional" json:"defaultExpiryInDays" yaml:"defaultExpiryInDays"`
	// Specifies the maximum number of days that can be set for the expiration time for a programmatic access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#max_expiry_in_days AuthenticationPolicy#max_expiry_in_days}
	MaxExpiryInDays *float64 `field:"optional" json:"maxExpiryInDays" yaml:"maxExpiryInDays"`
	// Specifies the network policy evaluation for the PAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#network_policy_evaluation AuthenticationPolicy#network_policy_evaluation}
	NetworkPolicyEvaluation *string `field:"optional" json:"networkPolicyEvaluation" yaml:"networkPolicyEvaluation"`
}

