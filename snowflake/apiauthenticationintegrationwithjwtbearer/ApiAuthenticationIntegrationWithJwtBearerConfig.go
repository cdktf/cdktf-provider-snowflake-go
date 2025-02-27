// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiauthenticationintegrationwithjwtbearer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiAuthenticationIntegrationWithJwtBearerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#enabled ApiAuthenticationIntegrationWithJwtBearer#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the identifier (i.e. name) for the integration. This value must be unique in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#name ApiAuthenticationIntegrationWithJwtBearer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#oauth_assertion_issuer ApiAuthenticationIntegrationWithJwtBearer#oauth_assertion_issuer}.
	OauthAssertionIssuer *string `field:"required" json:"oauthAssertionIssuer" yaml:"oauthAssertionIssuer"`
	// Specifies the client ID for the OAuth application in the external service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#oauth_client_id ApiAuthenticationIntegrationWithJwtBearer#oauth_client_id}
	OauthClientId *string `field:"required" json:"oauthClientId" yaml:"oauthClientId"`
	// Specifies the client secret for the OAuth application in the ServiceNow instance from the previous step.
	//
	// The connector uses this to request an access token from the ServiceNow instance. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#oauth_client_secret ApiAuthenticationIntegrationWithJwtBearer#oauth_client_secret}
	OauthClientSecret *string `field:"required" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// Specifies a comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#comment ApiAuthenticationIntegrationWithJwtBearer#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#id ApiAuthenticationIntegrationWithJwtBearer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the default lifetime of the OAuth access token (in seconds) issued by an OAuth server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#oauth_access_token_validity ApiAuthenticationIntegrationWithJwtBearer#oauth_access_token_validity}
	OauthAccessTokenValidity *float64 `field:"optional" json:"oauthAccessTokenValidity" yaml:"oauthAccessTokenValidity"`
	// Specifies the URL for authenticating to the external service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#oauth_authorization_endpoint ApiAuthenticationIntegrationWithJwtBearer#oauth_authorization_endpoint}
	OauthAuthorizationEndpoint *string `field:"optional" json:"oauthAuthorizationEndpoint" yaml:"oauthAuthorizationEndpoint"`
	// Specifies that POST is used as the authentication method to the external service.
	//
	// If removed from the config, the resource is recreated. Valid values are (case-insensitive): `CLIENT_SECRET_POST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#oauth_client_auth_method ApiAuthenticationIntegrationWithJwtBearer#oauth_client_auth_method}
	OauthClientAuthMethod *string `field:"optional" json:"oauthClientAuthMethod" yaml:"oauthClientAuthMethod"`
	// Specifies the value to determine the validity of the refresh token obtained from the OAuth server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#oauth_refresh_token_validity ApiAuthenticationIntegrationWithJwtBearer#oauth_refresh_token_validity}
	OauthRefreshTokenValidity *float64 `field:"optional" json:"oauthRefreshTokenValidity" yaml:"oauthRefreshTokenValidity"`
	// Specifies the token endpoint used by the client to obtain an access token by presenting its authorization grant or refresh token.
	//
	// The token endpoint is used with every authorization grant except for the implicit grant type (since an access token is issued directly). If removed from the config, the resource is recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/api_authentication_integration_with_jwt_bearer#oauth_token_endpoint ApiAuthenticationIntegrationWithJwtBearer#oauth_token_endpoint}
	OauthTokenEndpoint *string `field:"optional" json:"oauthTokenEndpoint" yaml:"oauthTokenEndpoint"`
}

