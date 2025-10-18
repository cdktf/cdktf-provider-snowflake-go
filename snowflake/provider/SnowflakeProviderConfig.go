// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type SnowflakeProviderConfig struct {
	// Specifies your Snowflake account name assigned by Snowflake.
	//
	// For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#account-name). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ACCOUNT_NAME` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#account_name SnowflakeProvider#account_name}
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#alias SnowflakeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Specifies the [authentication type](https://pkg.go.dev/github.com/snowflakedb/gosnowflake#AuthType) to use when connecting to Snowflake. Valid options are: `SNOWFLAKE` | `OAUTH` | `EXTERNALBROWSER` | `OKTA` | `SNOWFLAKE_JWT` | `TOKENACCESSOR` | `USERNAMEPASSWORDMFA` | `PROGRAMMATIC_ACCESS_TOKEN` | `OAUTH_CLIENT_CREDENTIALS` | `OAUTH_AUTHORIZATION_CODE`. Can also be sourced from the `SNOWFLAKE_AUTHENTICATOR` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#authenticator SnowflakeProvider#authenticator}
	Authenticator *string `field:"optional" json:"authenticator" yaml:"authenticator"`
	// IP address for network checks. Can also be sourced from the `SNOWFLAKE_CLIENT_IP` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#client_ip SnowflakeProvider#client_ip}
	ClientIp *string `field:"optional" json:"clientIp" yaml:"clientIp"`
	// When true the MFA token is cached in the credential manager.
	//
	// True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_REQUEST_MFA_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#client_request_mfa_token SnowflakeProvider#client_request_mfa_token}
	ClientRequestMfaToken *string `field:"optional" json:"clientRequestMfaToken" yaml:"clientRequestMfaToken"`
	// When true the ID token is cached in the credential manager.
	//
	// True by default in Windows/OSX. False for Linux. Can also be sourced from the `SNOWFLAKE_CLIENT_STORE_TEMPORARY_CREDENTIAL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#client_store_temporary_credential SnowflakeProvider#client_store_temporary_credential}
	ClientStoreTemporaryCredential *string `field:"optional" json:"clientStoreTemporaryCredential" yaml:"clientStoreTemporaryCredential"`
	// The timeout in seconds for the client to complete the authentication.
	//
	// Can also be sourced from the `SNOWFLAKE_CLIENT_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#client_timeout SnowflakeProvider#client_timeout}
	ClientTimeout *float64 `field:"optional" json:"clientTimeout" yaml:"clientTimeout"`
	// Indicates whether console login should be disabled in the driver. Can also be sourced from the `SNOWFLAKE_DISABLE_CONSOLE_LOGIN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#disable_console_login SnowflakeProvider#disable_console_login}
	DisableConsoleLogin *string `field:"optional" json:"disableConsoleLogin" yaml:"disableConsoleLogin"`
	// Disables HTAP query context cache in the driver. Can also be sourced from the `SNOWFLAKE_DISABLE_QUERY_CONTEXT_CACHE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#disable_query_context_cache SnowflakeProvider#disable_query_context_cache}
	DisableQueryContextCache interface{} `field:"optional" json:"disableQueryContextCache" yaml:"disableQueryContextCache"`
	// Disables telemetry in the driver. Can also be sourced from the `DISABLE_TELEMETRY` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#disable_telemetry SnowflakeProvider#disable_telemetry}
	DisableTelemetry interface{} `field:"optional" json:"disableTelemetry" yaml:"disableTelemetry"`
	// Specifies the logging level to be used by the driver.
	//
	// Valid options are: `trace` | `debug` | `info` | `print` | `warning` | `error` | `fatal` | `panic`. Can also be sourced from the `SNOWFLAKE_DRIVER_TRACING` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#driver_tracing SnowflakeProvider#driver_tracing}
	DriverTracing *string `field:"optional" json:"driverTracing" yaml:"driverTracing"`
	// Enables single use refresh tokens for Snowflake IdP. Can also be sourced from the `SNOWFLAKE_ENABLE_SINGLE_USE_REFRESH_TOKENS` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#enable_single_use_refresh_tokens SnowflakeProvider#enable_single_use_refresh_tokens}
	EnableSingleUseRefreshTokens interface{} `field:"optional" json:"enableSingleUseRefreshTokens" yaml:"enableSingleUseRefreshTokens"`
	// The timeout in seconds for the external browser to complete the authentication.
	//
	// Can also be sourced from the `SNOWFLAKE_EXTERNAL_BROWSER_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#external_browser_timeout SnowflakeProvider#external_browser_timeout}
	ExternalBrowserTimeout *float64 `field:"optional" json:"externalBrowserTimeout" yaml:"externalBrowserTimeout"`
	// Specifies a custom host value used by the driver for privatelink connections.
	//
	// Can also be sourced from the `SNOWFLAKE_HOST` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#host SnowflakeProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Should retried request contain retry reason. Can also be sourced from the `SNOWFLAKE_INCLUDE_RETRY_REASON` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#include_retry_reason SnowflakeProvider#include_retry_reason}
	IncludeRetryReason *string `field:"optional" json:"includeRetryReason" yaml:"includeRetryReason"`
	// If true, bypass the Online Certificate Status Protocol (OCSP) certificate revocation check.
	//
	// IMPORTANT: Change the default value for testing or emergency situations only. Can also be sourced from the `SNOWFLAKE_INSECURE_MODE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#insecure_mode SnowflakeProvider#insecure_mode}
	InsecureMode interface{} `field:"optional" json:"insecureMode" yaml:"insecureMode"`
	// The timeout in seconds for the JWT client to complete the authentication.
	//
	// Can also be sourced from the `SNOWFLAKE_JWT_CLIENT_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#jwt_client_timeout SnowflakeProvider#jwt_client_timeout}
	JwtClientTimeout *float64 `field:"optional" json:"jwtClientTimeout" yaml:"jwtClientTimeout"`
	// JWT expire after timeout in seconds. Can also be sourced from the `SNOWFLAKE_JWT_EXPIRE_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#jwt_expire_timeout SnowflakeProvider#jwt_expire_timeout}
	JwtExpireTimeout *float64 `field:"optional" json:"jwtExpireTimeout" yaml:"jwtExpireTimeout"`
	// Enables the session to persist even after the connection is closed.
	//
	// Can also be sourced from the `SNOWFLAKE_KEEP_SESSION_ALIVE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#keep_session_alive SnowflakeProvider#keep_session_alive}
	KeepSessionAlive interface{} `field:"optional" json:"keepSessionAlive" yaml:"keepSessionAlive"`
	// Login retry timeout in seconds EXCLUDING network roundtrip and read out http response.
	//
	// Can also be sourced from the `SNOWFLAKE_LOGIN_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#login_timeout SnowflakeProvider#login_timeout}
	LoginTimeout *float64 `field:"optional" json:"loginTimeout" yaml:"loginTimeout"`
	// Specifies how many times non-periodic HTTP request can be retried by the driver.
	//
	// Can also be sourced from the `SNOWFLAKE_MAX_RETRY_COUNT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#max_retry_count SnowflakeProvider#max_retry_count}
	MaxRetryCount *float64 `field:"optional" json:"maxRetryCount" yaml:"maxRetryCount"`
	// Authorization URL of OAuth2 external IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_AUTHORIZATION_URL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#oauth_authorization_url SnowflakeProvider#oauth_authorization_url}
	OauthAuthorizationUrl *string `field:"optional" json:"oauthAuthorizationUrl" yaml:"oauthAuthorizationUrl"`
	// Client id for OAuth2 external IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_CLIENT_ID` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#oauth_client_id SnowflakeProvider#oauth_client_id}
	OauthClientId *string `field:"optional" json:"oauthClientId" yaml:"oauthClientId"`
	// Client secret for OAuth2 external IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_CLIENT_SECRET` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#oauth_client_secret SnowflakeProvider#oauth_client_secret}
	OauthClientSecret *string `field:"optional" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// Redirect URI registered in IdP. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_REDIRECT_URI` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#oauth_redirect_uri SnowflakeProvider#oauth_redirect_uri}
	OauthRedirectUri *string `field:"optional" json:"oauthRedirectUri" yaml:"oauthRedirectUri"`
	// Comma separated list of scopes.
	//
	// If empty it is derived from role. See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_SCOPE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#oauth_scope SnowflakeProvider#oauth_scope}
	OauthScope *string `field:"optional" json:"oauthScope" yaml:"oauthScope"`
	// Token request URL of OAuth2 external IdP.
	//
	// See [Snowflake OAuth documentation](https://docs.snowflake.com/en/user-guide/oauth). Can also be sourced from the `SNOWFLAKE_OAUTH_TOKEN_REQUEST_URL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#oauth_token_request_url SnowflakeProvider#oauth_token_request_url}
	OauthTokenRequestUrl *string `field:"optional" json:"oauthTokenRequestUrl" yaml:"oauthTokenRequestUrl"`
	// True represents OCSP fail open mode.
	//
	// False represents OCSP fail closed mode. Fail open true by default. Can also be sourced from the `SNOWFLAKE_OCSP_FAIL_OPEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#ocsp_fail_open SnowflakeProvider#ocsp_fail_open}
	OcspFailOpen *string `field:"optional" json:"ocspFailOpen" yaml:"ocspFailOpen"`
	// The URL of the Okta server.
	//
	// e.g. https://example.okta.com. Okta URL host needs to to have a suffix `okta.com`. Read more in Snowflake [docs](https://docs.snowflake.com/en/user-guide/oauth-okta). Can also be sourced from the `SNOWFLAKE_OKTA_URL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#okta_url SnowflakeProvider#okta_url}
	OktaUrl *string `field:"optional" json:"oktaUrl" yaml:"oktaUrl"`
	// Specifies your Snowflake organization name assigned by Snowflake.
	//
	// For information about account identifiers, see the [Snowflake documentation](https://docs.snowflake.com/en/user-guide/admin-account-identifier#organization-name). Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_ORGANIZATION_NAME` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#organization_name SnowflakeProvider#organization_name}
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// Sets other connection (i.e. session) parameters. [Parameters](https://docs.snowflake.com/en/sql-reference/parameters). This field can not be set with environmental variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#params SnowflakeProvider#params}
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
	// Specifies the passcode provided by Duo when using multi-factor authentication (MFA) for login.
	//
	// Can also be sourced from the `SNOWFLAKE_PASSCODE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#passcode SnowflakeProvider#passcode}
	Passcode *string `field:"optional" json:"passcode" yaml:"passcode"`
	// False by default.
	//
	// Set to true if the MFA passcode is embedded to the configured password. Can also be sourced from the `SNOWFLAKE_PASSCODE_IN_PASSWORD` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#passcode_in_password SnowflakeProvider#passcode_in_password}
	PasscodeInPassword interface{} `field:"optional" json:"passcodeInPassword" yaml:"passcodeInPassword"`
	// Password for user + password or [token](https://docs.snowflake.com/en/user-guide/programmatic-access-tokens#generating-a-programmatic-access-token) for [PAT auth](https://docs.snowflake.com/en/user-guide/programmatic-access-tokens). Cannot be used with `private_key` and `private_key_passphrase`. Can also be sourced from the `SNOWFLAKE_PASSWORD` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#password SnowflakeProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Specifies a custom port value used by the driver for privatelink connections.
	//
	// Can also be sourced from the `SNOWFLAKE_PORT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#port SnowflakeProvider#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A list of preview features that are handled by the provider.
	//
	// See [preview features list](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/v1-preparations/LIST_OF_PREVIEW_FEATURES_FOR_V1.md). Preview features may have breaking changes in future releases, even without raising the major version. This field can not be set with environmental variables. Valid options are: `snowflake_account_authentication_policy_attachment_resource` | `snowflake_account_password_policy_attachment_resource` | `snowflake_alert_resource` | `snowflake_alerts_datasource` | `snowflake_api_integration_resource` | `snowflake_authentication_policy_resource` | `snowflake_compute_pool_resource` | `snowflake_compute_pools_datasource` | `snowflake_cortex_search_service_resource` | `snowflake_cortex_search_services_datasource` | `snowflake_current_account_resource` | `snowflake_current_account_datasource` | `snowflake_current_organization_account_resource` | `snowflake_database_datasource` | `snowflake_database_role_datasource` | `snowflake_dynamic_table_resource` | `snowflake_dynamic_tables_datasource` | `snowflake_external_function_resource` | `snowflake_external_functions_datasource` | `snowflake_external_table_resource` | `snowflake_external_tables_datasource` | `snowflake_external_volume_resource` | `snowflake_failover_group_resource` | `snowflake_failover_groups_datasource` | `snowflake_file_format_resource` | `snowflake_file_formats_datasource` | `snowflake_function_java_resource` | `snowflake_function_javascript_resource` | `snowflake_function_python_resource` | `snowflake_function_scala_resource` | `snowflake_function_sql_resource` | `snowflake_functions_datasource` | `snowflake_git_repository_resource` | `snowflake_git_repositories_datasource` | `snowflake_image_repository_resource` | `snowflake_image_repositories_datasource` | `snowflake_job_service_resource` | `snowflake_listing_resource` | `snowflake_managed_account_resource` | `snowflake_materialized_view_resource` | `snowflake_materialized_views_datasource` | `snowflake_network_policy_attachment_resource` | `snowflake_network_rule_resource` | `snowflake_email_notification_integration_resource` | `snowflake_notification_integration_resource` | `snowflake_object_parameter_resource` | `snowflake_password_policy_resource` | `snowflake_pipe_resource` | `snowflake_pipes_datasource` | `snowflake_current_role_datasource` | `snowflake_service_resource` | `snowflake_services_datasource` | `snowflake_sequence_resource` | `snowflake_sequences_datasource` | `snowflake_share_resource` | `snowflake_shares_datasource` | `snowflake_parameters_datasource` | `snowflake_procedure_java_resource` | `snowflake_procedure_javascript_resource` | `snowflake_procedure_python_resource` | `snowflake_procedure_scala_resource` | `snowflake_procedure_sql_resource` | `snowflake_procedures_datasource` | `snowflake_stage_resource` | `snowflake_stages_datasource` | `snowflake_storage_integration_resource` | `snowflake_storage_integrations_datasource` | `snowflake_system_generate_scim_access_token_datasource` | `snowflake_system_get_aws_sns_iam_policy_datasource` | `snowflake_system_get_privatelink_config_datasource` | `snowflake_system_get_snowflake_platform_info_datasource` | `snowflake_table_column_masking_policy_application_resource` | `snowflake_table_constraint_resource` | `snowflake_table_resource` | `snowflake_tables_datasource` | `snowflake_user_authentication_policy_attachment_resource` | `snowflake_user_public_keys_resource` | `snowflake_user_password_policy_attachment_resource` | `snowflake_user_programmatic_access_token_resource` | `snowflake_user_programmatic_access_tokens_datasource`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#preview_features_enabled SnowflakeProvider#preview_features_enabled}
	PreviewFeaturesEnabled *[]*string `field:"optional" json:"previewFeaturesEnabled" yaml:"previewFeaturesEnabled"`
	// Private Key for username+private-key auth. Cannot be used with `password`. Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#private_key SnowflakeProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Supports the encryption ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc.
	//
	// Can also be sourced from the `SNOWFLAKE_PRIVATE_KEY_PASSPHRASE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#private_key_passphrase SnowflakeProvider#private_key_passphrase}
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
	// Sets the profile to read from ~/.snowflake/config file. Can also be sourced from the `SNOWFLAKE_PROFILE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#profile SnowflakeProvider#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// A protocol used in the connection.
	//
	// Valid options are: `http` | `https`. Can also be sourced from the `SNOWFLAKE_PROTOCOL` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#protocol SnowflakeProvider#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// request retry timeout in seconds EXCLUDING network roundtrip and read out http response.
	//
	// Can also be sourced from the `SNOWFLAKE_REQUEST_TIMEOUT` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#request_timeout SnowflakeProvider#request_timeout}
	RequestTimeout *float64 `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// Specifies the role to use by default for accessing Snowflake objects in the client session.
	//
	// Can also be sourced from the `SNOWFLAKE_ROLE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#role SnowflakeProvider#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// False by default.
	//
	// Skips TOML configuration file permission verification. This flag has no effect on Windows systems, as the permissions are not checked on this platform. Instead of skipping the permissions verification, we recommend setting the proper privileges - see [the section below](#toml-file-limitations). Can also be sourced from the `SNOWFLAKE_SKIP_TOML_FILE_PERMISSION_VERIFICATION` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#skip_toml_file_permission_verification SnowflakeProvider#skip_toml_file_permission_verification}
	SkipTomlFilePermissionVerification interface{} `field:"optional" json:"skipTomlFilePermissionVerification" yaml:"skipTomlFilePermissionVerification"`
	// Sets temporary directory used by the driver for operations like encrypting, compressing etc.
	//
	// Can also be sourced from the `SNOWFLAKE_TMP_DIRECTORY_PATH` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#tmp_directory_path SnowflakeProvider#tmp_directory_path}
	TmpDirectoryPath *string `field:"optional" json:"tmpDirectoryPath" yaml:"tmpDirectoryPath"`
	// Token to use for OAuth and other forms of token based auth.
	//
	// When this field is set here, or in the TOML file, the provider sets the `authenticator` to `OAUTH`. Optionally, set the `authenticator` field to the authenticator you want to use. Can also be sourced from the `SNOWFLAKE_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#token SnowflakeProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// token_accessor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#token_accessor SnowflakeProvider#token_accessor}
	TokenAccessor *SnowflakeProviderTokenAccessor `field:"optional" json:"tokenAccessor" yaml:"tokenAccessor"`
	// False by default.
	//
	// When this is set to true, the provider expects the legacy TOML format. Otherwise, it expects the new format. See more in [the section below](#examples) Can also be sourced from the `SNOWFLAKE_USE_LEGACY_TOML_FILE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#use_legacy_toml_file SnowflakeProvider#use_legacy_toml_file}
	UseLegacyTomlFile interface{} `field:"optional" json:"useLegacyTomlFile" yaml:"useLegacyTomlFile"`
	// Username. Required unless using `profile`. Can also be sourced from the `SNOWFLAKE_USER` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#user SnowflakeProvider#user}
	User *string `field:"optional" json:"user" yaml:"user"`
	// True by default.
	//
	// If false, disables the validation checks for Database, Schema, Warehouse and Role at the time a connection is established. Can also be sourced from the `SNOWFLAKE_VALIDATE_DEFAULT_PARAMETERS` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#validate_default_parameters SnowflakeProvider#validate_default_parameters}
	ValidateDefaultParameters *string `field:"optional" json:"validateDefaultParameters" yaml:"validateDefaultParameters"`
	// Specifies the virtual warehouse to use by default for queries, loading, etc.
	//
	// in the client session. Can also be sourced from the `SNOWFLAKE_WAREHOUSE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.9.0/docs#warehouse SnowflakeProvider#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}

