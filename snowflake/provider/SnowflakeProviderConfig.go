package provider


type SnowflakeProviderConfig struct {
	// The name of the Snowflake account. Can also come from the `SNOWFLAKE_ACCOUNT` environment variable. Required unless using profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#account SnowflakeProvider#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#alias SnowflakeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_USE_BROWSER_AUTH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#browser_auth SnowflakeProvider#browser_auth}
	BrowserAuth interface{} `field:"optional" json:"browserAuth" yaml:"browserAuth"`
	// Supports passing in a custom host value to the snowflake go driver for use with privatelink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#host SnowflakeProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// If true, bypass the Online Certificate Status Protocol (OCSP) certificate revocation check.
	//
	// IMPORTANT: Change the default value for testing or emergency situations only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#insecure_mode SnowflakeProvider#insecure_mode}
	InsecureMode interface{} `field:"optional" json:"insecureMode" yaml:"insecureMode"`
	// Token for use with OAuth.
	//
	// Generating the token is left to other tools. Cannot be used with `browser_auth`, `private_key_path`, `oauth_refresh_token` or `password`. Can be sourced from `SNOWFLAKE_OAUTH_ACCESS_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#oauth_access_token SnowflakeProvider#oauth_access_token}
	OauthAccessToken *string `field:"optional" json:"oauthAccessToken" yaml:"oauthAccessToken"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_OAUTH_CLIENT_ID` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#oauth_client_id SnowflakeProvider#oauth_client_id}
	OauthClientId *string `field:"optional" json:"oauthClientId" yaml:"oauthClientId"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_OAUTH_CLIENT_SECRET` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#oauth_client_secret SnowflakeProvider#oauth_client_secret}
	OauthClientSecret *string `field:"optional" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_OAUTH_ENDPOINT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#oauth_endpoint SnowflakeProvider#oauth_endpoint}
	OauthEndpoint *string `field:"optional" json:"oauthEndpoint" yaml:"oauthEndpoint"`
	// Required when `oauth_refresh_token` is used. Can be sourced from `SNOWFLAKE_OAUTH_REDIRECT_URL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#oauth_redirect_url SnowflakeProvider#oauth_redirect_url}
	OauthRedirectUrl *string `field:"optional" json:"oauthRedirectUrl" yaml:"oauthRedirectUrl"`
	// Token for use with OAuth.
	//
	// Setup and generation of the token is left to other tools. Should be used in conjunction with `oauth_client_id`, `oauth_client_secret`, `oauth_endpoint`, `oauth_redirect_url`. Cannot be used with `browser_auth`, `private_key_path`, `oauth_access_token` or `password`. Can be sourced from `SNOWFLAKE_OAUTH_REFRESH_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#oauth_refresh_token SnowflakeProvider#oauth_refresh_token}
	OauthRefreshToken *string `field:"optional" json:"oauthRefreshToken" yaml:"oauthRefreshToken"`
	// Password for username+password auth. Cannot be used with `browser_auth` or `private_key_path`. Can be sourced from `SNOWFLAKE_PASSWORD` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#password SnowflakeProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Support custom port values to snowflake go driver for use with privatelink. Can be sourced from `SNOWFLAKE_PORT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#port SnowflakeProvider#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Private Key for username+private-key auth. Cannot be used with `browser_auth` or `password`. Can be sourced from `SNOWFLAKE_PRIVATE_KEY` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#private_key SnowflakeProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#private_key_passphrase SnowflakeProvider#private_key_passphrase}
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
	// Path to a private key for using keypair authentication.
	//
	// Cannot be used with `browser_auth`, `oauth_access_token` or `password`. Can be sourced from `SNOWFLAKE_PRIVATE_KEY_PATH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#private_key_path SnowflakeProvider#private_key_path}
	PrivateKeyPath *string `field:"optional" json:"privateKeyPath" yaml:"privateKeyPath"`
	// Sets the profile to read from ~/.snowflake/config file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#profile SnowflakeProvider#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Support custom protocols to snowflake go driver. Can be sourced from `SNOWFLAKE_PROTOCOL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#protocol SnowflakeProvider#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// [Snowflake region](https://docs.snowflake.com/en/user-guide/intro-regions.html) to use.  Required if using the [legacy format for the `account` identifier](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#format-2-legacy-account-locator-in-a-region) in the form of `<cloud_region_id>.<cloud>`. Can be sourced from the `SNOWFLAKE_REGION` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#region SnowflakeProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Snowflake role to use for operations.
	//
	// If left unset, default role for user will be used. Can be sourced from the `SNOWFLAKE_ROLE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#role SnowflakeProvider#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Username for username+password authentication. Can come from the `SNOWFLAKE_USER` environment variable. Required unless using profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#username SnowflakeProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Sets the default warehouse. Optional. Can be sourced from SNOWFLAKE_WAREHOUSE environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs#warehouse SnowflakeProvider#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}

