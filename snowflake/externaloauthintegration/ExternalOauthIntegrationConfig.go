// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externaloauthintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExternalOauthIntegrationConfig struct {
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
	// Specifies whether to initiate operation of the integration or suspend it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#enabled ExternalOauthIntegration#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the URL to define the OAuth 2.0 authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#issuer ExternalOauthIntegration#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Specifies the name of the External Oath integration.
	//
	// This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#name ExternalOauthIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates which Snowflake user record attribute should be used to map the access token to a Snowflake user record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#snowflake_user_mapping_attribute ExternalOauthIntegration#snowflake_user_mapping_attribute}
	SnowflakeUserMappingAttribute *string `field:"required" json:"snowflakeUserMappingAttribute" yaml:"snowflakeUserMappingAttribute"`
	// Specifies the access token claim or claims that can be used to map the access token to a Snowflake user record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#token_user_mapping_claims ExternalOauthIntegration#token_user_mapping_claims}
	TokenUserMappingClaims *[]*string `field:"required" json:"tokenUserMappingClaims" yaml:"tokenUserMappingClaims"`
	// Specifies the OAuth 2.0 authorization server to be Okta, Microsoft Azure AD, Ping Identity PingFederate, or a Custom OAuth 2.0 authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#type ExternalOauthIntegration#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the list of roles that the client can set as the primary role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#allowed_roles ExternalOauthIntegration#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// Specifies whether the OAuth client or user can use a role that is not defined in the OAuth access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#any_role_mode ExternalOauthIntegration#any_role_mode}
	AnyRoleMode *string `field:"optional" json:"anyRoleMode" yaml:"anyRoleMode"`
	// Specifies additional values that can be used for the access token's audience validation on top of using the Customer's Snowflake Account URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#audience_urls ExternalOauthIntegration#audience_urls}
	AudienceUrls *[]*string `field:"optional" json:"audienceUrls" yaml:"audienceUrls"`
	// Specifies the list of roles that a client cannot set as the primary role.
	//
	// Do not include ACCOUNTADMIN, ORGADMIN or SECURITYADMIN as they are already implicitly enforced and will cause in-place updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#blocked_roles ExternalOauthIntegration#blocked_roles}
	BlockedRoles *[]*string `field:"optional" json:"blockedRoles" yaml:"blockedRoles"`
	// Specifies a comment for the OAuth integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#comment ExternalOauthIntegration#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#id ExternalOauthIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the endpoint or a list of endpoints from which to download public keys or certificates to validate an External OAuth access token.
	//
	// The maximum number of URLs that can be specified in the list is 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#jws_keys_urls ExternalOauthIntegration#jws_keys_urls}
	JwsKeysUrls *[]*string `field:"optional" json:"jwsKeysUrls" yaml:"jwsKeysUrls"`
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#rsa_public_key ExternalOauthIntegration#rsa_public_key}
	RsaPublicKey *string `field:"optional" json:"rsaPublicKey" yaml:"rsaPublicKey"`
	// Specifies a second RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	//
	// Used for key rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#rsa_public_key_2 ExternalOauthIntegration#rsa_public_key_2}
	RsaPublicKey2 *string `field:"optional" json:"rsaPublicKey2" yaml:"rsaPublicKey2"`
	// Specifies the scope delimiter in the authorization token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#scope_delimiter ExternalOauthIntegration#scope_delimiter}
	ScopeDelimiter *string `field:"optional" json:"scopeDelimiter" yaml:"scopeDelimiter"`
	// Specifies the access token claim to map the access token to an account role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_oauth_integration#scope_mapping_attribute ExternalOauthIntegration#scope_mapping_attribute}
	ScopeMappingAttribute *string `field:"optional" json:"scopeMappingAttribute" yaml:"scopeMappingAttribute"`
}

