// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake


type SnowflakeProviderConfig struct {
	// The name of the Snowflake account. Can also come from the `SNOWFLAKE_ACCOUNT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#account SnowflakeProvider#account}
	Account *string `field:"required" json:"account" yaml:"account"`
	// Username for username+password authentication. Can come from the `SNOWFLAKE_USER` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#username SnowflakeProvider#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#alias SnowflakeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_USE_BROWSER_AUTH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#browser_auth SnowflakeProvider#browser_auth}
	BrowserAuth interface{} `field:"optional" json:"browserAuth" yaml:"browserAuth"`
	// Supports passing in a custom host value to the snowflake go driver for use with privatelink.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#host SnowflakeProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Token for use with OAuth.
	//
	// Generating the token is left to other tools. Cannot be used with `browser_auth`, `private_key_path`, `oauth_refresh_token` or `password`. Can be sourced from `SNOWFLAKE_OAUTH_ACCESS_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#oauth_access_token SnowflakeProvider#oauth_access_token}
	OauthAccessToken *string `field:"optional" json:"oauthAccessToken" yaml:"oauthAccessToken"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_OAUTH_CLIENT_ID` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#oauth_client_id SnowflakeProvider#oauth_client_id}
	OauthClientId *string `field:"optional" json:"oauthClientId" yaml:"oauthClientId"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_OAUTH_CLIENT_SECRET` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#oauth_client_secret SnowflakeProvider#oauth_client_secret}
	OauthClientSecret *string `field:"optional" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_OAUTH_ENDPOINT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#oauth_endpoint SnowflakeProvider#oauth_endpoint}
	OauthEndpoint *string `field:"optional" json:"oauthEndpoint" yaml:"oauthEndpoint"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_OAUTH_REDIRECT_URL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#oauth_redirect_url SnowflakeProvider#oauth_redirect_url}
	OauthRedirectUrl *string `field:"optional" json:"oauthRedirectUrl" yaml:"oauthRedirectUrl"`
	// Token for use with OAuth.
	//
	// Setup and generation of the token is left to other tools. Should be used in conjunction with `oauth_client_id`, `oauth_client_secret`, `oauth_endpoint`, `oauth_redirect_url`. Cannot be used with `browser_auth`, `private_key_path`, `oauth_access_token` or `password`. Can be sourced from `SNOWFLAKE_OAUTH_REFRESH_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#oauth_refresh_token SnowflakeProvider#oauth_refresh_token}
	OauthRefreshToken *string `field:"optional" json:"oauthRefreshToken" yaml:"oauthRefreshToken"`
	// Password for username+password auth. Cannot be used with `browser_auth` or `private_key_path`. Can be source from `SNOWFLAKE_PASSWORD` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#password SnowflakeProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Private Key for username+private-key auth. Cannot be used with `browser_auth` or `password`. Can be source from `SNOWFLAKE_PRIVATE_KEY` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#private_key SnowflakeProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#private_key_passphrase SnowflakeProvider#private_key_passphrase}
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
	// Path to a private key for using keypair authentication.
	//
	// Cannot be used with `browser_auth`, `oauth_access_token` or `password`. Can be source from `SNOWFLAKE_PRIVATE_KEY_PATH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#private_key_path SnowflakeProvider#private_key_path}
	PrivateKeyPath *string `field:"optional" json:"privateKeyPath" yaml:"privateKeyPath"`
	// [Snowflake region](https://docs.snowflake.com/en/user-guide/intro-regions.html) to use. Can be source from the `SNOWFLAKE_REGION` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#region SnowflakeProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Snowflake role to use for operations.
	//
	// If left unset, default role for user will be used. Can come from the `SNOWFLAKE_ROLE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#role SnowflakeProvider#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Sets the default warehouse. Optional. Can be sourced from SNOWFLAKE_WAREHOUSE environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake#warehouse SnowflakeProvider#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}

