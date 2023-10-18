// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type SnowflakeProviderTokenAccessor struct {
	// The client ID for the OAuth provider when using a refresh token to renew access token.
	//
	// Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_CLIENT_ID` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs#client_id SnowflakeProvider#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret for the OAuth provider when using a refresh token to renew access token.
	//
	// Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_CLIENT_SECRET` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs#client_secret SnowflakeProvider#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The redirect URI for the OAuth provider when using a refresh token to renew access token.
	//
	// Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_REDIRECT_URI` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs#redirect_uri SnowflakeProvider#redirect_uri}
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
	// The refresh token for the OAuth provider when using a refresh token to renew access token.
	//
	// Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_REFRESH_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs#refresh_token SnowflakeProvider#refresh_token}
	RefreshToken *string `field:"required" json:"refreshToken" yaml:"refreshToken"`
	// The token endpoint for the OAuth provider e.g. https://{yourDomain}/oauth/token when using a refresh token to renew access token. Can also be sourced from the `SNOWFLAKE_TOKEN_ACCESSOR_TOKEN_ENDPOINT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs#token_endpoint SnowflakeProvider#token_endpoint}
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

