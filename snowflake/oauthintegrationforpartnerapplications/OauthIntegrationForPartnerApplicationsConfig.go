// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforpartnerapplications

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OauthIntegrationForPartnerApplicationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#blocked_roles_list OauthIntegrationForPartnerApplications#blocked_roles_list}
	BlockedRolesList *[]*string `field:"required" json:"blockedRolesList" yaml:"blockedRolesList"`
	// Specifies the name of the OAuth integration.
	//
	// This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#name OauthIntegrationForPartnerApplications#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Creates an OAuth interface between Snowflake and a partner application. Valid options are: `LOOKER` | `TABLEAU_DESKTOP` | `TABLEAU_SERVER`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#oauth_client OauthIntegrationForPartnerApplications#oauth_client}
	OauthClient *string `field:"required" json:"oauthClient" yaml:"oauthClient"`
	// Specifies a comment for the OAuth integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#comment OauthIntegrationForPartnerApplications#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether this OAuth integration is enabled or disabled.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#enabled OauthIntegrationForPartnerApplications#enabled}
	Enabled *string `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#id OauthIntegrationForPartnerApplications#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether to allow the client to exchange a refresh token for an access token when the current access token has expired.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#oauth_issue_refresh_tokens OauthIntegrationForPartnerApplications#oauth_issue_refresh_tokens}
	OauthIssueRefreshTokens *string `field:"optional" json:"oauthIssueRefreshTokens" yaml:"oauthIssueRefreshTokens"`
	// Specifies the client URI.
	//
	// After a user is authenticated, the web browser is redirected to this URI. The field should be only set when OAUTH_CLIENT = LOOKER. In any other case the field should be left out empty. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#oauth_redirect_uri OauthIntegrationForPartnerApplications#oauth_redirect_uri}
	OauthRedirectUri *string `field:"optional" json:"oauthRedirectUri" yaml:"oauthRedirectUri"`
	// Specifies how long refresh tokens should be valid (in seconds). OAUTH_ISSUE_REFRESH_TOKENS must be set to TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#oauth_refresh_token_validity OauthIntegrationForPartnerApplications#oauth_refresh_token_validity}
	OauthRefreshTokenValidity *float64 `field:"optional" json:"oauthRefreshTokenValidity" yaml:"oauthRefreshTokenValidity"`
	// Specifies whether default secondary roles set in the user properties are activated by default in the session being opened.
	//
	// Valid options are: `IMPLICIT` | `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/oauth_integration_for_partner_applications#oauth_use_secondary_roles OauthIntegrationForPartnerApplications#oauth_use_secondary_roles}
	OauthUseSecondaryRoles *string `field:"optional" json:"oauthUseSecondaryRoles" yaml:"oauthUseSecondaryRoles"`
}

