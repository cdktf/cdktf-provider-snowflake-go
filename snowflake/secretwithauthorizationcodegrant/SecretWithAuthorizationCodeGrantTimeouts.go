// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretwithauthorizationcodegrant


type SecretWithAuthorizationCodeGrantTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/secret_with_authorization_code_grant#create SecretWithAuthorizationCodeGrant#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/secret_with_authorization_code_grant#delete SecretWithAuthorizationCodeGrant#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/secret_with_authorization_code_grant#read SecretWithAuthorizationCodeGrant#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/secret_with_authorization_code_grant#update SecretWithAuthorizationCodeGrant#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

