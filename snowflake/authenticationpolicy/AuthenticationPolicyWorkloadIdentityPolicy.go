// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticationpolicy


type AuthenticationPolicyWorkloadIdentityPolicy struct {
	// Specifies the list of AWS account IDs allowed by the authentication policy during workload identity authentication of type `AWS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#allowed_aws_accounts AuthenticationPolicy#allowed_aws_accounts}
	AllowedAwsAccounts *[]*string `field:"optional" json:"allowedAwsAccounts" yaml:"allowedAwsAccounts"`
	// Specifies the list of Azure Entra ID issuers allowed by the authentication policy during workload identity authentication of type `AZURE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#allowed_azure_issuers AuthenticationPolicy#allowed_azure_issuers}
	AllowedAzureIssuers *[]*string `field:"optional" json:"allowedAzureIssuers" yaml:"allowedAzureIssuers"`
	// Specifies the list of OIDC issuers allowed by the authentication policy during workload identity authentication of type `OIDC`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#allowed_oidc_issuers AuthenticationPolicy#allowed_oidc_issuers}
	AllowedOidcIssuers *[]*string `field:"optional" json:"allowedOidcIssuers" yaml:"allowedOidcIssuers"`
	// Specifies the allowed providers for the workload identity policy.
	//
	// Valid values are: `ALL` | `AWS` | `AZURE` | `GCP` | `OIDC`. These values are case-sensitive due to Terraform limitations (it's a nested field). Prefer using uppercased values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/authentication_policy#allowed_providers AuthenticationPolicy#allowed_providers}
	AllowedProviders *[]*string `field:"optional" json:"allowedProviders" yaml:"allowedProviders"`
}

