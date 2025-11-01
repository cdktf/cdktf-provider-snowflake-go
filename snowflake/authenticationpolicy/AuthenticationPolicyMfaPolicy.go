// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticationpolicy


type AuthenticationPolicyMfaPolicy struct {
	// Specifies the allowed methods for the MFA policy.
	//
	// Valid values are: `ALL` | `PASSKEY` | `TOTP` | `DUO`. These values are case-sensitive due to Terraform limitations (it's a nested field). Prefer using uppercased values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#allowed_methods AuthenticationPolicy#allowed_methods}
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Determines whether multi-factor authentication (MFA) is enforced on external authentication. Valid values are (case-insensitive): `ALL` | `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#enforce_mfa_on_external_authentication AuthenticationPolicy#enforce_mfa_on_external_authentication}
	EnforceMfaOnExternalAuthentication *string `field:"optional" json:"enforceMfaOnExternalAuthentication" yaml:"enforceMfaOnExternalAuthentication"`
}

