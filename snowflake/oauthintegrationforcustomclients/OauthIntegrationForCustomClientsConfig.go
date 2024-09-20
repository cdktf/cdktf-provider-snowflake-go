// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforcustomclients

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OauthIntegrationForCustomClientsConfig struct {
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
	// A set of Snowflake roles that a user cannot explicitly consent to using after authenticating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#blocked_roles_list OauthIntegrationForCustomClients#blocked_roles_list}
	BlockedRolesList *[]*string `field:"required" json:"blockedRolesList" yaml:"blockedRolesList"`
	// Specifies the name of the OAuth integration.
	//
	// This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#name OauthIntegrationForCustomClients#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the type of client being registered.
	//
	// Snowflake supports both confidential and public clients. Valid options are: `PUBLIC` | `CONFIDENTIAL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_client_type OauthIntegrationForCustomClients#oauth_client_type}
	OauthClientType *string `field:"required" json:"oauthClientType" yaml:"oauthClientType"`
	// Specifies the client URI. After a user is authenticated, the web browser is redirected to this URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_redirect_uri OauthIntegrationForCustomClients#oauth_redirect_uri}
	OauthRedirectUri *string `field:"required" json:"oauthRedirectUri" yaml:"oauthRedirectUri"`
	// Specifies a comment for the OAuth integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#comment OauthIntegrationForCustomClients#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether this OAuth integration is enabled or disabled.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#enabled OauthIntegrationForCustomClients#enabled}
	Enabled *string `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#id OauthIntegrationForCustomClients#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies an existing network policy.
	//
	// This network policy controls network traffic that is attempting to exchange an authorization code for an access or refresh token or to use a refresh token to obtain a new access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#network_policy OauthIntegrationForCustomClients#network_policy}
	NetworkPolicy *string `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
	// If true, allows setting oauth_redirect_uri to a URI not protected by TLS.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_allow_non_tls_redirect_uri OauthIntegrationForCustomClients#oauth_allow_non_tls_redirect_uri}
	OauthAllowNonTlsRedirectUri *string `field:"optional" json:"oauthAllowNonTlsRedirectUri" yaml:"oauthAllowNonTlsRedirectUri"`
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource using `terraform taint`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_client_rsa_public_key OauthIntegrationForCustomClients#oauth_client_rsa_public_key}
	OauthClientRsaPublicKey *string `field:"optional" json:"oauthClientRsaPublicKey" yaml:"oauthClientRsaPublicKey"`
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource using `terraform taint`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_client_rsa_public_key_2 OauthIntegrationForCustomClients#oauth_client_rsa_public_key_2}
	OauthClientRsaPublicKey2 *string `field:"optional" json:"oauthClientRsaPublicKey2" yaml:"oauthClientRsaPublicKey2"`
	// Boolean that specifies whether Proof Key for Code Exchange (PKCE) should be required for the integration.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_enforce_pkce OauthIntegrationForCustomClients#oauth_enforce_pkce}
	OauthEnforcePkce *string `field:"optional" json:"oauthEnforcePkce" yaml:"oauthEnforcePkce"`
	// Specifies whether to allow the client to exchange a refresh token for an access token when the current access token has expired.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_issue_refresh_tokens OauthIntegrationForCustomClients#oauth_issue_refresh_tokens}
	OauthIssueRefreshTokens *string `field:"optional" json:"oauthIssueRefreshTokens" yaml:"oauthIssueRefreshTokens"`
	// Specifies how long refresh tokens should be valid (in seconds). OAUTH_ISSUE_REFRESH_TOKENS must be set to TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_refresh_token_validity OauthIntegrationForCustomClients#oauth_refresh_token_validity}
	OauthRefreshTokenValidity *float64 `field:"optional" json:"oauthRefreshTokenValidity" yaml:"oauthRefreshTokenValidity"`
	// Specifies whether default secondary roles set in the user properties are activated by default in the session being opened.
	//
	// Valid options are: `IMPLICIT` | `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#oauth_use_secondary_roles OauthIntegrationForCustomClients#oauth_use_secondary_roles}
	OauthUseSecondaryRoles *string `field:"optional" json:"oauthUseSecondaryRoles" yaml:"oauthUseSecondaryRoles"`
	// A set of Snowflake roles that a user does not need to explicitly consent to using after authenticating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/oauth_integration_for_custom_clients#pre_authorized_roles_list OauthIntegrationForCustomClients#pre_authorized_roles_list}
	PreAuthorizedRolesList *[]*string `field:"optional" json:"preAuthorizedRolesList" yaml:"preAuthorizedRolesList"`
}

