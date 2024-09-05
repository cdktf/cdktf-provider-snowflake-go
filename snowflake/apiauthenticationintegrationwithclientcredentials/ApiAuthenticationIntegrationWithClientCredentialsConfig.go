// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiauthenticationintegrationwithclientcredentials

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiAuthenticationIntegrationWithClientCredentialsConfig struct {
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
	// Specifies whether this security integration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#enabled ApiAuthenticationIntegrationWithClientCredentials#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the identifier (i.e. name) for the integration. This value must be unique in your account. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#name ApiAuthenticationIntegrationWithClientCredentials#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the client ID for the OAuth application in the external service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#oauth_client_id ApiAuthenticationIntegrationWithClientCredentials#oauth_client_id}
	OauthClientId *string `field:"required" json:"oauthClientId" yaml:"oauthClientId"`
	// Specifies the client secret for the OAuth application in the ServiceNow instance from the previous step.
	//
	// The connector uses this to request an access token from the ServiceNow instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#oauth_client_secret ApiAuthenticationIntegrationWithClientCredentials#oauth_client_secret}
	OauthClientSecret *string `field:"required" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// Specifies a comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#comment ApiAuthenticationIntegrationWithClientCredentials#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#id ApiAuthenticationIntegrationWithClientCredentials#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the default lifetime of the OAuth access token (in seconds) issued by an OAuth server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#oauth_access_token_validity ApiAuthenticationIntegrationWithClientCredentials#oauth_access_token_validity}
	OauthAccessTokenValidity *float64 `field:"optional" json:"oauthAccessTokenValidity" yaml:"oauthAccessTokenValidity"`
	// Specifies a list of scopes to use when making a request from the OAuth by a role with USAGE on the integration during the OAuth client credentials flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#oauth_allowed_scopes ApiAuthenticationIntegrationWithClientCredentials#oauth_allowed_scopes}
	OauthAllowedScopes *[]*string `field:"optional" json:"oauthAllowedScopes" yaml:"oauthAllowedScopes"`
	// Specifies that POST is used as the authentication method to the external service.
	//
	// If removed from the config, the resource is recreated. Valid values are (case-insensitive): `CLIENT_SECRET_POST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#oauth_client_auth_method ApiAuthenticationIntegrationWithClientCredentials#oauth_client_auth_method}
	OauthClientAuthMethod *string `field:"optional" json:"oauthClientAuthMethod" yaml:"oauthClientAuthMethod"`
	// Specifies the value to determine the validity of the refresh token obtained from the OAuth server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#oauth_refresh_token_validity ApiAuthenticationIntegrationWithClientCredentials#oauth_refresh_token_validity}
	OauthRefreshTokenValidity *float64 `field:"optional" json:"oauthRefreshTokenValidity" yaml:"oauthRefreshTokenValidity"`
	// Specifies the token endpoint used by the client to obtain an access token by presenting its authorization grant or refresh token.
	//
	// The token endpoint is used with every authorization grant except for the implicit grant type (since an access token is issued directly). If removed from the config, the resource is recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/api_authentication_integration_with_client_credentials#oauth_token_endpoint ApiAuthenticationIntegrationWithClientCredentials#oauth_token_endpoint}
	OauthTokenEndpoint *string `field:"optional" json:"oauthTokenEndpoint" yaml:"oauthTokenEndpoint"`
}

