package oauthintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OauthIntegrationConfig struct {
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
	// Specifies the name of the OAuth integration.
	//
	// This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#name OauthIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the OAuth client type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#oauth_client OauthIntegration#oauth_client}
	OauthClient *string `field:"required" json:"oauthClient" yaml:"oauthClient"`
	// List of roles that a user cannot explicitly consent to using after authenticating.
	//
	// Do not include ACCOUNTADMIN, ORGADMIN or SECURITYADMIN as they are already implicitly enforced and will cause in-place updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#blocked_roles_list OauthIntegration#blocked_roles_list}
	BlockedRolesList *[]*string `field:"optional" json:"blockedRolesList" yaml:"blockedRolesList"`
	// Specifies a comment for the OAuth integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#comment OauthIntegration#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether this OAuth integration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#enabled OauthIntegration#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#id OauthIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the type of client being registered. Snowflake supports both confidential and public clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#oauth_client_type OauthIntegration#oauth_client_type}
	OauthClientType *string `field:"optional" json:"oauthClientType" yaml:"oauthClientType"`
	// Specifies whether to allow the client to exchange a refresh token for an access token when the current access token has expired.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#oauth_issue_refresh_tokens OauthIntegration#oauth_issue_refresh_tokens}
	OauthIssueRefreshTokens interface{} `field:"optional" json:"oauthIssueRefreshTokens" yaml:"oauthIssueRefreshTokens"`
	// Specifies the client URI. After a user is authenticated, the web browser is redirected to this URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#oauth_redirect_uri OauthIntegration#oauth_redirect_uri}
	OauthRedirectUri *string `field:"optional" json:"oauthRedirectUri" yaml:"oauthRedirectUri"`
	// Specifies how long refresh tokens should be valid (in seconds). OAUTH_ISSUE_REFRESH_TOKENS must be set to TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#oauth_refresh_token_validity OauthIntegration#oauth_refresh_token_validity}
	OauthRefreshTokenValidity *float64 `field:"optional" json:"oauthRefreshTokenValidity" yaml:"oauthRefreshTokenValidity"`
	// Specifies whether default secondary roles set in the user properties are activated by default in the session being opened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/oauth_integration#oauth_use_secondary_roles OauthIntegration#oauth_use_secondary_roles}
	OauthUseSecondaryRoles *string `field:"optional" json:"oauthUseSecondaryRoles" yaml:"oauthUseSecondaryRoles"`
}

