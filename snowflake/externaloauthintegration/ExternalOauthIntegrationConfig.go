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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#enabled ExternalOauthIntegration#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the URL to define the OAuth 2.0 authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_issuer ExternalOauthIntegration#external_oauth_issuer}
	ExternalOauthIssuer *string `field:"required" json:"externalOauthIssuer" yaml:"externalOauthIssuer"`
	// Indicates which Snowflake user record attribute should be used to map the access token to a Snowflake user record.
	//
	// Valid values are (case-insensitive): `LOGIN_NAME` | `EMAIL_ADDRESS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_snowflake_user_mapping_attribute ExternalOauthIntegration#external_oauth_snowflake_user_mapping_attribute}
	ExternalOauthSnowflakeUserMappingAttribute *string `field:"required" json:"externalOauthSnowflakeUserMappingAttribute" yaml:"externalOauthSnowflakeUserMappingAttribute"`
	// Specifies the access token claim or claims that can be used to map the access token to a Snowflake user record.
	//
	// If removed from the config, the resource is recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_token_user_mapping_claim ExternalOauthIntegration#external_oauth_token_user_mapping_claim}
	ExternalOauthTokenUserMappingClaim *[]*string `field:"required" json:"externalOauthTokenUserMappingClaim" yaml:"externalOauthTokenUserMappingClaim"`
	// Specifies the OAuth 2.0 authorization server to be Okta, Microsoft Azure AD, Ping Identity PingFederate, or a Custom OAuth 2.0 authorization server. Valid values are (case-insensitive): `OKTA` | `AZURE` | `PING_FEDERATE` | `CUSTOM`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_type ExternalOauthIntegration#external_oauth_type}
	ExternalOauthType *string `field:"required" json:"externalOauthType" yaml:"externalOauthType"`
	// Specifies the name of the External Oath integration.
	//
	// This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#name ExternalOauthIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies a comment for the OAuth integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#comment ExternalOauthIntegration#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the list of roles that the client can set as the primary role.
	//
	// For more information about this resource, see [docs](./account_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_allowed_roles_list ExternalOauthIntegration#external_oauth_allowed_roles_list}
	ExternalOauthAllowedRolesList *[]*string `field:"optional" json:"externalOauthAllowedRolesList" yaml:"externalOauthAllowedRolesList"`
	// Specifies whether the OAuth client or user can use a role that is not defined in the OAuth access token.
	//
	// Valid values are (case-insensitive): `DISABLE` | `ENABLE` | `ENABLE_FOR_PRIVILEGE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_any_role_mode ExternalOauthIntegration#external_oauth_any_role_mode}
	ExternalOauthAnyRoleMode *string `field:"optional" json:"externalOauthAnyRoleMode" yaml:"externalOauthAnyRoleMode"`
	// Specifies additional values that can be used for the access token's audience validation on top of using the Customer's Snowflake Account URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_audience_list ExternalOauthIntegration#external_oauth_audience_list}
	ExternalOauthAudienceList *[]*string `field:"optional" json:"externalOauthAudienceList" yaml:"externalOauthAudienceList"`
	// Specifies the list of roles that a client cannot set as the primary role.
	//
	// By default, this list includes the ACCOUNTADMIN, ORGADMIN and SECURITYADMIN roles. To remove these privileged roles from the list, use the ALTER ACCOUNT command to set the EXTERNAL_OAUTH_ADD_PRIVILEGED_ROLES_TO_BLOCKED_LIST account parameter to FALSE. For more information about this resource, see [docs](./account_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_blocked_roles_list ExternalOauthIntegration#external_oauth_blocked_roles_list}
	ExternalOauthBlockedRolesList *[]*string `field:"optional" json:"externalOauthBlockedRolesList" yaml:"externalOauthBlockedRolesList"`
	// Specifies the endpoint or a list of endpoints from which to download public keys or certificates to validate an External OAuth access token.
	//
	// The maximum number of URLs that can be specified in the list is 3. If removed from the config, the resource is recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_jws_keys_url ExternalOauthIntegration#external_oauth_jws_keys_url}
	ExternalOauthJwsKeysUrl *[]*string `field:"optional" json:"externalOauthJwsKeysUrl" yaml:"externalOauthJwsKeysUrl"`
	// Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	//
	// If removed from the config, the resource is recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_rsa_public_key ExternalOauthIntegration#external_oauth_rsa_public_key}
	ExternalOauthRsaPublicKey *string `field:"optional" json:"externalOauthRsaPublicKey" yaml:"externalOauthRsaPublicKey"`
	// Specifies a second RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers.
	//
	// Used for key rotation. If removed from the config, the resource is recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_rsa_public_key_2 ExternalOauthIntegration#external_oauth_rsa_public_key_2}
	ExternalOauthRsaPublicKey2 *string `field:"optional" json:"externalOauthRsaPublicKey2" yaml:"externalOauthRsaPublicKey2"`
	// Specifies the scope delimiter in the authorization token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_scope_delimiter ExternalOauthIntegration#external_oauth_scope_delimiter}
	ExternalOauthScopeDelimiter *string `field:"optional" json:"externalOauthScopeDelimiter" yaml:"externalOauthScopeDelimiter"`
	// Specifies the access token claim to map the access token to an account role.
	//
	// If removed from the config, the resource is recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#external_oauth_scope_mapping_attribute ExternalOauthIntegration#external_oauth_scope_mapping_attribute}
	ExternalOauthScopeMappingAttribute *string `field:"optional" json:"externalOauthScopeMappingAttribute" yaml:"externalOauthScopeMappingAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#id ExternalOauthIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/external_oauth_integration#timeouts ExternalOauthIntegration#timeouts}
	Timeouts *ExternalOauthIntegrationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

