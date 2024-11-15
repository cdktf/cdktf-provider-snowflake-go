// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type SnowflakeProviderConfig struct {
	// Use `account_name` and `organization_name` instead.
	//
	// Specifies your Snowflake account identifier assigned, by Snowflake. The [account locator](https://docs.snowflake.com/en/user-guide/admin-account-identifier#format-2-account-locator-in-a-region) format is not supported. For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ACCOUNT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#account SnowflakeProvider#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Specifies your Snowflake account name assigned by Snowflake.
	//
	// For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#account-name). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ACCOUNT_NAME` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#account_name SnowflakeProvider#account_name}
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#alias SnowflakeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Specifies the [authentication type](https://pkg.go.dev/github.com/snowflakedb/gosnowflake#AuthType) to use when connecting to Snowflake. Valid options are: `SNOWFLAKE` | `OAUTH` | `EXTERNALBROWSER` | `OKTA` | `JWT` | `SNOWFLAKE_JWT` | `TOKENACCESSOR` | `USERNAMEPASSWORDMFA`. Value `JWT` is deprecated and will be removed in future releases. Can also be sourced from the `SNOWFLAKE_AUTHENTICATOR` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#authenticator SnowflakeProvider#authenticator}
	Authenticator *string `field:"optional" json:"authenticator" yaml:"authenticator"`
	// Required when `oauth_refresh_token` is used. Can also be sourced from `SNOWFLAKE_USE_BROWSER_AUTH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#browser_auth SnowflakeProvider#browser_auth}
	BrowserAuth interface{} `field:"optional" json:"browserAuth" yaml:"browserAuth"`
	// IP address for network checks. Can also be sourced from the `SNOWFLAKE_CLIENT_IP` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#client_ip SnowflakeProvider#client_ip}
	ClientIp *string `field:"optional" json:"clientIp" yaml:"clientIp"`
	// When true the MFA token is cached in the credential manager.
	//
	// True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_REQUEST_MFA_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#client_request_mfa_token SnowflakeProvider#client_request_mfa_token}
	ClientRequestMfaToken *string `field:"optional" json:"clientRequestMfaToken" yaml:"clientRequestMfaToken"`
	// When true the ID token is cached in the credential manager.
	//
	// True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_STORE_TEMPORARY_CREDENTIAL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#client_store_temporary_credential SnowflakeProvider#client_store_temporary_credential}
	ClientStoreTemporaryCredential *string `field:"optional" json:"clientStoreTemporaryCredential" yaml:"clientStoreTemporaryCredential"`
	// The timeout in seconds for the client to complete the authentication.
	//
	// Can also be sourced from the `SNOWFLAKE_CLIENT_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#client_timeout SnowflakeProvider#client_timeout}
	ClientTimeout *float64 `field:"optional" json:"clientTimeout" yaml:"clientTimeout"`
	// Indicates whether console login should be disabled in the driver. Can also be sourced from the `SNOWFLAKE_DISABLE_CONSOLE_LOGIN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#disable_console_login SnowflakeProvider#disable_console_login}
	DisableConsoleLogin *string `field:"optional" json:"disableConsoleLogin" yaml:"disableConsoleLogin"`
	// Disables HTAP query context cache in the driver. Can also be sourced from the `SNOWFLAKE_DISABLE_QUERY_CONTEXT_CACHE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#disable_query_context_cache SnowflakeProvider#disable_query_context_cache}
	DisableQueryContextCache interface{} `field:"optional" json:"disableQueryContextCache" yaml:"disableQueryContextCache"`
	// Disables telemetry in the driver. Can also be sourced from the `DISABLE_TELEMETRY` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#disable_telemetry SnowflakeProvider#disable_telemetry}
	DisableTelemetry interface{} `field:"optional" json:"disableTelemetry" yaml:"disableTelemetry"`
	// Specifies the logging level to be used by the driver.
	//
	// Valid options are: `trace` | `debug` | `info` | `print` | `warning` | `error` | `fatal` | `panic`. Can also be sourced from the `SNOWFLAKE_DRIVER_TRACING` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#driver_tracing SnowflakeProvider#driver_tracing}
	DriverTracing *string `field:"optional" json:"driverTracing" yaml:"driverTracing"`
	// The timeout in seconds for the external browser to complete the authentication.
	//
	// Can also be sourced from the `SNOWFLAKE_EXTERNAL_BROWSER_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#external_browser_timeout SnowflakeProvider#external_browser_timeout}
	ExternalBrowserTimeout *float64 `field:"optional" json:"externalBrowserTimeout" yaml:"externalBrowserTimeout"`
	// Specifies a custom host value used by the driver for privatelink connections.
	//
	// Can also be sourced from the `SNOWFLAKE_HOST` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#host SnowflakeProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Should retried request contain retry reason. Can also be sourced from the `SNOWFLAKE_INCLUDE_RETRY_REASON` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#include_retry_reason SnowflakeProvider#include_retry_reason}
	IncludeRetryReason *string `field:"optional" json:"includeRetryReason" yaml:"includeRetryReason"`
	// If true, bypass the Online Certificate Status Protocol (OCSP) certificate revocation check.
	//
	// IMPORTANT: Change the default value for testing or emergency situations only. Can also be sourced from the `SNOWFLAKE_INSECURE_MODE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#insecure_mode SnowflakeProvider#insecure_mode}
	InsecureMode interface{} `field:"optional" json:"insecureMode" yaml:"insecureMode"`
	// The timeout in seconds for the JWT client to complete the authentication.
	//
	// Can also be sourced from the `SNOWFLAKE_JWT_CLIENT_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#jwt_client_timeout SnowflakeProvider#jwt_client_timeout}
	JwtClientTimeout *float64 `field:"optional" json:"jwtClientTimeout" yaml:"jwtClientTimeout"`
	// JWT expire after timeout in seconds. Can also be sourced from the `SNOWFLAKE_JWT_EXPIRE_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#jwt_expire_timeout SnowflakeProvider#jwt_expire_timeout}
	JwtExpireTimeout *float64 `field:"optional" json:"jwtExpireTimeout" yaml:"jwtExpireTimeout"`
	// Enables the session to persist even after the connection is closed.
	//
	// Can also be sourced from the `SNOWFLAKE_KEEP_SESSION_ALIVE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#keep_session_alive SnowflakeProvider#keep_session_alive}
	KeepSessionAlive interface{} `field:"optional" json:"keepSessionAlive" yaml:"keepSessionAlive"`
	// Login retry timeout in seconds EXCLUDING network roundtrip and read out http response.
	//
	// Can also be sourced from the `SNOWFLAKE_LOGIN_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#login_timeout SnowflakeProvider#login_timeout}
	LoginTimeout *float64 `field:"optional" json:"loginTimeout" yaml:"loginTimeout"`
	// Specifies how many times non-periodic HTTP request can be retried by the driver.
	//
	// Can also be sourced from the `SNOWFLAKE_MAX_RETRY_COUNT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#max_retry_count SnowflakeProvider#max_retry_count}
	MaxRetryCount *float64 `field:"optional" json:"maxRetryCount" yaml:"maxRetryCount"`
	// Token for use with OAuth.
	//
	// Generating the token is left to other tools. Cannot be used with `browser_auth`, `private_key_path`, `oauth_refresh_token` or `password`. Can also be sourced from `SNOWFLAKE_OAUTH_ACCESS_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#oauth_access_token SnowflakeProvider#oauth_access_token}
	OauthAccessToken *string `field:"optional" json:"oauthAccessToken" yaml:"oauthAccessToken"`
	// Required when `oauth_refresh_token` is used. Can also be sourced from `SNOWFLAKE_OAUTH_CLIENT_ID` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#oauth_client_id SnowflakeProvider#oauth_client_id}
	OauthClientId *string `field:"optional" json:"oauthClientId" yaml:"oauthClientId"`
	// Required when `oauth_refresh_token` is used. Can also be sourced from `SNOWFLAKE_OAUTH_CLIENT_SECRET` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#oauth_client_secret SnowflakeProvider#oauth_client_secret}
	OauthClientSecret *string `field:"optional" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// Required when `oauth_refresh_token` is used. Can also be sourced from `SNOWFLAKE_OAUTH_ENDPOINT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#oauth_endpoint SnowflakeProvider#oauth_endpoint}
	OauthEndpoint *string `field:"optional" json:"oauthEndpoint" yaml:"oauthEndpoint"`
	// Required when `oauth_refresh_token` is used. Can also be sourced from `SNOWFLAKE_OAUTH_REDIRECT_URL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#oauth_redirect_url SnowflakeProvider#oauth_redirect_url}
	OauthRedirectUrl *string `field:"optional" json:"oauthRedirectUrl" yaml:"oauthRedirectUrl"`
	// Token for use with OAuth.
	//
	// Setup and generation of the token is left to other tools. Should be used in conjunction with `oauth_client_id`, `oauth_client_secret`, `oauth_endpoint`, `oauth_redirect_url`. Cannot be used with `browser_auth`, `private_key_path`, `oauth_access_token` or `password`. Can also be sourced from `SNOWFLAKE_OAUTH_REFRESH_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#oauth_refresh_token SnowflakeProvider#oauth_refresh_token}
	OauthRefreshToken *string `field:"optional" json:"oauthRefreshToken" yaml:"oauthRefreshToken"`
	// True represents OCSP fail open mode.
	//
	// False represents OCSP fail closed mode. Fail open true by default. Can also be sourced from the `SNOWFLAKE_OCSP_FAIL_OPEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#ocsp_fail_open SnowflakeProvider#ocsp_fail_open}
	OcspFailOpen *string `field:"optional" json:"ocspFailOpen" yaml:"ocspFailOpen"`
	// The URL of the Okta server.
	//
	// e.g. https://example.okta.com. Okta URL host needs to to have a suffix `okta.com`. Read more in Snowflake [docs](https://docs.snowflake.com/en/user-guide/oauth-okta). Can also be sourced from the `SNOWFLAKE_OKTA_URL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#okta_url SnowflakeProvider#okta_url}
	OktaUrl *string `field:"optional" json:"oktaUrl" yaml:"oktaUrl"`
	// Specifies your Snowflake organization name assigned by Snowflake.
	//
	// For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#organization-name). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ORGANIZATION_NAME` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#organization_name SnowflakeProvider#organization_name}
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// Sets other connection (i.e. session) parameters. [Parameters](https://docs.snowflake.com/en/sql-reference/parameters). This field can not be set with environmental variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#params SnowflakeProvider#params}
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
	// Specifies the passcode provided by Duo when using multi-factor authentication (MFA) for login.
	//
	// Can also be sourced from the `SNOWFLAKE_PASSCODE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#passcode SnowflakeProvider#passcode}
	Passcode *string `field:"optional" json:"passcode" yaml:"passcode"`
	// False by default.
	//
	// Set to true if the MFA passcode is embedded to the configured password. Can also be sourced from the `SNOWFLAKE_PASSCODE_IN_PASSWORD` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#passcode_in_password SnowflakeProvider#passcode_in_password}
	PasscodeInPassword interface{} `field:"optional" json:"passcodeInPassword" yaml:"passcodeInPassword"`
	// Password for user + password auth.
	//
	// Cannot be used with `browser_auth` or `private_key_path`. Can also be sourced from the `SNOWFLAKE_PASSWORD` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#password SnowflakeProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Specifies a custom port value used by the driver for privatelink connections.
	//
	// Can also be sourced from the `SNOWFLAKE_PORT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#port SnowflakeProvider#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Private Key for username+private-key auth.
	//
	// Cannot be used with `browser_auth` or `password`. Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#private_key SnowflakeProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc.
	//
	// Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY_PASSPHRASE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#private_key_passphrase SnowflakeProvider#private_key_passphrase}
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
	// Path to a private key for using keypair authentication.
	//
	// Cannot be used with `browser_auth`, `oauth_access_token` or `password`. Can also be sourced from `SNOWFLAKE_PRIVATE_KEY_PATH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#private_key_path SnowflakeProvider#private_key_path}
	PrivateKeyPath *string `field:"optional" json:"privateKeyPath" yaml:"privateKeyPath"`
	// Sets the profile to read from ~/.snowflake/config file. Can also be sourced from the `SNOWFLAKE_PROFILE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#profile SnowflakeProvider#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// A protocol used in the connection.
	//
	// Valid options are: `http` | `https`. Can also be sourced from the `SNOWFLAKE_PROTOCOL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#protocol SnowflakeProvider#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Snowflake region, such as "eu-central-1", with this parameter.
	//
	// However, since this parameter is deprecated, it is best to specify the region as part of the account parameter. For details, see the description of the account parameter. [Snowflake region](https://docs.snowflake.com/en/user-guide/intro-regions.html) to use.  Required if using the [legacy format for the `account` identifier](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#format-2-legacy-account-locator-in-a-region) in the form of `<cloud_region_id>.<cloud>`. Can also be sourced from the `SNOWFLAKE_REGION` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#region SnowflakeProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// request retry timeout in seconds EXCLUDING network roundtrip and read out http response.
	//
	// Can also be sourced from the `SNOWFLAKE_REQUEST_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#request_timeout SnowflakeProvider#request_timeout}
	RequestTimeout *float64 `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// Specifies the role to use by default for accessing Snowflake objects in the client session.
	//
	// Can also be sourced from the `SNOWFLAKE_ROLE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#role SnowflakeProvider#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Sets session parameters. [Parameters](https://docs.snowflake.com/en/sql-reference/parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#session_params SnowflakeProvider#session_params}
	SessionParams *map[string]*string `field:"optional" json:"sessionParams" yaml:"sessionParams"`
	// Sets temporary directory used by the driver for operations like encrypting, compressing etc.
	//
	// Can also be sourced from the `SNOWFLAKE_TMP_DIRECTORY_PATH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#tmp_directory_path SnowflakeProvider#tmp_directory_path}
	TmpDirectoryPath *string `field:"optional" json:"tmpDirectoryPath" yaml:"tmpDirectoryPath"`
	// Token to use for OAuth and other forms of token based auth.
	//
	// Can also be sourced from the `SNOWFLAKE_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#token SnowflakeProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// token_accessor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#token_accessor SnowflakeProvider#token_accessor}
	TokenAccessor *SnowflakeProviderTokenAccessor `field:"optional" json:"tokenAccessor" yaml:"tokenAccessor"`
	// Username. Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_USER` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#user SnowflakeProvider#user}
	User *string `field:"optional" json:"user" yaml:"user"`
	// Username for user + password authentication. Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_USERNAME` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#username SnowflakeProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// True by default.
	//
	// If false, disables the validation checks for Database, Schema, Warehouse and Role at the time a connection is established. Can also be sourced from the `SNOWFLAKE_VALIDATE_DEFAULT_PARAMETERS` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#validate_default_parameters SnowflakeProvider#validate_default_parameters}
	ValidateDefaultParameters *string `field:"optional" json:"validateDefaultParameters" yaml:"validateDefaultParameters"`
	// Specifies the virtual warehouse to use by default for queries, loading, etc.
	//
	// in the client session. Can also be sourced from the `SNOWFLAKE_WAREHOUSE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs#warehouse SnowflakeProvider#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}

