// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretwithauthorizationcodegrant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretWithAuthorizationCodeGrantConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Specifies the name value of the Snowflake security integration that connects Snowflake to an external service.
	//
	// For more information about this resource, see [docs](./api_authentication_integration_with_authorization_code_grant).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#api_authentication SecretWithAuthorizationCodeGrant#api_authentication}
	ApiAuthentication *string `field:"required" json:"apiAuthentication" yaml:"apiAuthentication"`
	// The database in which to create the secret Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#database SecretWithAuthorizationCodeGrant#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// String that specifies the identifier (i.e. name) for the secret, must be unique in your schema. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#name SecretWithAuthorizationCodeGrant#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the token as a string that is used to obtain a new access token from the OAuth authorization server when the access token expires.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#oauth_refresh_token SecretWithAuthorizationCodeGrant#oauth_refresh_token}
	OauthRefreshToken *string `field:"required" json:"oauthRefreshToken" yaml:"oauthRefreshToken"`
	// Specifies the timestamp as a string when the OAuth refresh token expires.
	//
	// Accepted string formats: YYYY-MM-DD, YYYY-MM-DD HH:MI, YYYY-MM-DD HH:MI:SS, YYYY-MM-DD HH:MI <timezone>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#oauth_refresh_token_expiry_time SecretWithAuthorizationCodeGrant#oauth_refresh_token_expiry_time}
	OauthRefreshTokenExpiryTime *string `field:"required" json:"oauthRefreshTokenExpiryTime" yaml:"oauthRefreshTokenExpiryTime"`
	// The schema in which to create the secret.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#schema SecretWithAuthorizationCodeGrant#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies a comment for the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#comment SecretWithAuthorizationCodeGrant#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#id SecretWithAuthorizationCodeGrant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/secret_with_authorization_code_grant#timeouts SecretWithAuthorizationCodeGrant#timeouts}
	Timeouts *SecretWithAuthorizationCodeGrantTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

