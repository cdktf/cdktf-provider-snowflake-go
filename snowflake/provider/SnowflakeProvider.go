// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs snowflake}.
type SnowflakeProvider interface {
	cdktf.TerraformProvider
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	Authenticator() *string
	SetAuthenticator(val *string)
	AuthenticatorInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientIp() *string
	SetClientIp(val *string)
	ClientIpInput() *string
	ClientRequestMfaToken() *string
	SetClientRequestMfaToken(val *string)
	ClientRequestMfaTokenInput() *string
	ClientStoreTemporaryCredential() *string
	SetClientStoreTemporaryCredential(val *string)
	ClientStoreTemporaryCredentialInput() *string
	ClientTimeout() *float64
	SetClientTimeout(val *float64)
	ClientTimeoutInput() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DisableConsoleLogin() *string
	SetDisableConsoleLogin(val *string)
	DisableConsoleLoginInput() *string
	DisableQueryContextCache() interface{}
	SetDisableQueryContextCache(val interface{})
	DisableQueryContextCacheInput() interface{}
	DisableTelemetry() interface{}
	SetDisableTelemetry(val interface{})
	DisableTelemetryInput() interface{}
	DriverTracing() *string
	SetDriverTracing(val *string)
	DriverTracingInput() *string
	ExternalBrowserTimeout() *float64
	SetExternalBrowserTimeout(val *float64)
	ExternalBrowserTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IncludeRetryReason() *string
	SetIncludeRetryReason(val *string)
	IncludeRetryReasonInput() *string
	InsecureMode() interface{}
	SetInsecureMode(val interface{})
	InsecureModeInput() interface{}
	JwtClientTimeout() *float64
	SetJwtClientTimeout(val *float64)
	JwtClientTimeoutInput() *float64
	JwtExpireTimeout() *float64
	SetJwtExpireTimeout(val *float64)
	JwtExpireTimeoutInput() *float64
	KeepSessionAlive() interface{}
	SetKeepSessionAlive(val interface{})
	KeepSessionAliveInput() interface{}
	LoginTimeout() *float64
	SetLoginTimeout(val *float64)
	LoginTimeoutInput() *float64
	MaxRetryCount() *float64
	SetMaxRetryCount(val *float64)
	MaxRetryCountInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	OcspFailOpen() *string
	SetOcspFailOpen(val *string)
	OcspFailOpenInput() *string
	OktaUrl() *string
	SetOktaUrl(val *string)
	OktaUrlInput() *string
	OrganizationName() *string
	SetOrganizationName(val *string)
	OrganizationNameInput() *string
	Params() *map[string]*string
	SetParams(val *map[string]*string)
	ParamsInput() *map[string]*string
	Passcode() *string
	SetPasscode(val *string)
	PasscodeInPassword() interface{}
	SetPasscodeInPassword(val interface{})
	PasscodeInPasswordInput() interface{}
	PasscodeInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PreviewFeaturesEnabled() *[]*string
	SetPreviewFeaturesEnabled(val *[]*string)
	PreviewFeaturesEnabledInput() *[]*string
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	PrivateKeyPassphrase() *string
	SetPrivateKeyPassphrase(val *string)
	PrivateKeyPassphraseInput() *string
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	RawOverrides() interface{}
	RequestTimeout() *float64
	SetRequestTimeout(val *float64)
	RequestTimeoutInput() *float64
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SkipTomlFilePermissionVerification() interface{}
	SetSkipTomlFilePermissionVerification(val interface{})
	SkipTomlFilePermissionVerificationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TmpDirectoryPath() *string
	SetTmpDirectoryPath(val *string)
	TmpDirectoryPathInput() *string
	Token() *string
	SetToken(val *string)
	TokenAccessor() *SnowflakeProviderTokenAccessor
	SetTokenAccessor(val *SnowflakeProviderTokenAccessor)
	TokenAccessorInput() *SnowflakeProviderTokenAccessor
	TokenInput() *string
	UseLegacyTomlFile() interface{}
	SetUseLegacyTomlFile(val interface{})
	UseLegacyTomlFileInput() interface{}
	User() *string
	SetUser(val *string)
	UserInput() *string
	ValidateDefaultParameters() *string
	SetValidateDefaultParameters(val *string)
	ValidateDefaultParametersInput() *string
	Warehouse() *string
	SetWarehouse(val *string)
	WarehouseInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccountName()
	ResetAlias()
	ResetAuthenticator()
	ResetClientIp()
	ResetClientRequestMfaToken()
	ResetClientStoreTemporaryCredential()
	ResetClientTimeout()
	ResetDisableConsoleLogin()
	ResetDisableQueryContextCache()
	ResetDisableTelemetry()
	ResetDriverTracing()
	ResetExternalBrowserTimeout()
	ResetHost()
	ResetIncludeRetryReason()
	ResetInsecureMode()
	ResetJwtClientTimeout()
	ResetJwtExpireTimeout()
	ResetKeepSessionAlive()
	ResetLoginTimeout()
	ResetMaxRetryCount()
	ResetOcspFailOpen()
	ResetOktaUrl()
	ResetOrganizationName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParams()
	ResetPasscode()
	ResetPasscodeInPassword()
	ResetPassword()
	ResetPort()
	ResetPreviewFeaturesEnabled()
	ResetPrivateKey()
	ResetPrivateKeyPassphrase()
	ResetProfile()
	ResetProtocol()
	ResetRequestTimeout()
	ResetRole()
	ResetSkipTomlFilePermissionVerification()
	ResetTmpDirectoryPath()
	ResetToken()
	ResetTokenAccessor()
	ResetUseLegacyTomlFile()
	ResetUser()
	ResetValidateDefaultParameters()
	ResetWarehouse()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SnowflakeProvider
type jsiiProxy_SnowflakeProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_SnowflakeProvider) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Authenticator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) AuthenticatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ClientIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ClientIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ClientRequestMfaToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientRequestMfaToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ClientRequestMfaTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientRequestMfaTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ClientStoreTemporaryCredential() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientStoreTemporaryCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ClientStoreTemporaryCredentialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientStoreTemporaryCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ClientTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ClientTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) DisableConsoleLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disableConsoleLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) DisableConsoleLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disableConsoleLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) DisableQueryContextCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableQueryContextCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) DisableQueryContextCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableQueryContextCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) DisableTelemetry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) DisableTelemetryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTelemetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) DriverTracing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) DriverTracingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverTracingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ExternalBrowserTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalBrowserTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ExternalBrowserTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalBrowserTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) IncludeRetryReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeRetryReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) IncludeRetryReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeRetryReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) InsecureMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) InsecureModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) JwtClientTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jwtClientTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) JwtClientTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jwtClientTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) JwtExpireTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jwtExpireTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) JwtExpireTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jwtExpireTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) KeepSessionAlive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepSessionAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) KeepSessionAliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepSessionAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) LoginTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) LoginTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loginTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) MaxRetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) MaxRetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OcspFailOpen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspFailOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OcspFailOpenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspFailOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OktaUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OktaUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Params() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Passcode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passcode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PasscodeInPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passcodeInPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PasscodeInPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passcodeInPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PasscodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passcodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PreviewFeaturesEnabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewFeaturesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PreviewFeaturesEnabledInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewFeaturesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKeyPassphrase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKeyPassphraseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphraseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) RequestTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) RequestTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) SkipTomlFilePermissionVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipTomlFilePermissionVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) SkipTomlFilePermissionVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipTomlFilePermissionVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TmpDirectoryPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tmpDirectoryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TmpDirectoryPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tmpDirectoryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TokenAccessor() *SnowflakeProviderTokenAccessor {
	var returns *SnowflakeProviderTokenAccessor
	_jsii_.Get(
		j,
		"tokenAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TokenAccessorInput() *SnowflakeProviderTokenAccessor {
	var returns *SnowflakeProviderTokenAccessor
	_jsii_.Get(
		j,
		"tokenAccessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) UseLegacyTomlFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLegacyTomlFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) UseLegacyTomlFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLegacyTomlFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ValidateDefaultParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateDefaultParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ValidateDefaultParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateDefaultParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Warehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) WarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs snowflake} Resource.
func NewSnowflakeProvider(scope constructs.Construct, id *string, config *SnowflakeProviderConfig) SnowflakeProvider {
	_init_.Initialize()

	if err := validateNewSnowflakeProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnowflakeProvider{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.provider.SnowflakeProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs snowflake} Resource.
func NewSnowflakeProvider_Override(s SnowflakeProvider, scope constructs.Construct, id *string, config *SnowflakeProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.provider.SnowflakeProvider",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetAuthenticator(val *string) {
	_jsii_.Set(
		j,
		"authenticator",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetClientIp(val *string) {
	_jsii_.Set(
		j,
		"clientIp",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetClientRequestMfaToken(val *string) {
	_jsii_.Set(
		j,
		"clientRequestMfaToken",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetClientStoreTemporaryCredential(val *string) {
	_jsii_.Set(
		j,
		"clientStoreTemporaryCredential",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetClientTimeout(val *float64) {
	_jsii_.Set(
		j,
		"clientTimeout",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetDisableConsoleLogin(val *string) {
	_jsii_.Set(
		j,
		"disableConsoleLogin",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetDisableQueryContextCache(val interface{}) {
	if err := j.validateSetDisableQueryContextCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableQueryContextCache",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetDisableTelemetry(val interface{}) {
	if err := j.validateSetDisableTelemetryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableTelemetry",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetDriverTracing(val *string) {
	_jsii_.Set(
		j,
		"driverTracing",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetExternalBrowserTimeout(val *float64) {
	_jsii_.Set(
		j,
		"externalBrowserTimeout",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetIncludeRetryReason(val *string) {
	_jsii_.Set(
		j,
		"includeRetryReason",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetInsecureMode(val interface{}) {
	if err := j.validateSetInsecureModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureMode",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetJwtClientTimeout(val *float64) {
	_jsii_.Set(
		j,
		"jwtClientTimeout",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetJwtExpireTimeout(val *float64) {
	_jsii_.Set(
		j,
		"jwtExpireTimeout",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetKeepSessionAlive(val interface{}) {
	if err := j.validateSetKeepSessionAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepSessionAlive",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetLoginTimeout(val *float64) {
	_jsii_.Set(
		j,
		"loginTimeout",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetMaxRetryCount(val *float64) {
	_jsii_.Set(
		j,
		"maxRetryCount",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOcspFailOpen(val *string) {
	_jsii_.Set(
		j,
		"ocspFailOpen",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOktaUrl(val *string) {
	_jsii_.Set(
		j,
		"oktaUrl",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOrganizationName(val *string) {
	_jsii_.Set(
		j,
		"organizationName",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetParams(val *map[string]*string) {
	_jsii_.Set(
		j,
		"params",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPasscode(val *string) {
	_jsii_.Set(
		j,
		"passcode",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPasscodeInPassword(val interface{}) {
	if err := j.validateSetPasscodeInPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passcodeInPassword",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPreviewFeaturesEnabled(val *[]*string) {
	_jsii_.Set(
		j,
		"previewFeaturesEnabled",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPrivateKeyPassphrase(val *string) {
	_jsii_.Set(
		j,
		"privateKeyPassphrase",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetProfile(val *string) {
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetRequestTimeout(val *float64) {
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetSkipTomlFilePermissionVerification(val interface{}) {
	if err := j.validateSetSkipTomlFilePermissionVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipTomlFilePermissionVerification",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetTmpDirectoryPath(val *string) {
	_jsii_.Set(
		j,
		"tmpDirectoryPath",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetTokenAccessor(val *SnowflakeProviderTokenAccessor) {
	if err := j.validateSetTokenAccessorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenAccessor",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetUseLegacyTomlFile(val interface{}) {
	if err := j.validateSetUseLegacyTomlFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLegacyTomlFile",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetValidateDefaultParameters(val *string) {
	_jsii_.Set(
		j,
		"validateDefaultParameters",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetWarehouse(val *string) {
	_jsii_.Set(
		j,
		"warehouse",
		val,
	)
}

// Generates CDKTF code for importing a SnowflakeProvider resource upon running "cdktf plan <stack-name>".
func SnowflakeProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSnowflakeProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.provider.SnowflakeProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SnowflakeProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnowflakeProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.provider.SnowflakeProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnowflakeProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnowflakeProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.provider.SnowflakeProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnowflakeProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnowflakeProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.provider.SnowflakeProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnowflakeProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.provider.SnowflakeProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SnowflakeProvider) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SnowflakeProvider) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetAccountName() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		s,
		"resetAlias",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetAuthenticator() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthenticator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetClientIp() {
	_jsii_.InvokeVoid(
		s,
		"resetClientIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetClientRequestMfaToken() {
	_jsii_.InvokeVoid(
		s,
		"resetClientRequestMfaToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetClientStoreTemporaryCredential() {
	_jsii_.InvokeVoid(
		s,
		"resetClientStoreTemporaryCredential",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetClientTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetClientTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetDisableConsoleLogin() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableConsoleLogin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetDisableQueryContextCache() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableQueryContextCache",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetDisableTelemetry() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableTelemetry",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetDriverTracing() {
	_jsii_.InvokeVoid(
		s,
		"resetDriverTracing",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetExternalBrowserTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalBrowserTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetHost() {
	_jsii_.InvokeVoid(
		s,
		"resetHost",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetIncludeRetryReason() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeRetryReason",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetInsecureMode() {
	_jsii_.InvokeVoid(
		s,
		"resetInsecureMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetJwtClientTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetJwtClientTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetJwtExpireTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetJwtExpireTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetKeepSessionAlive() {
	_jsii_.InvokeVoid(
		s,
		"resetKeepSessionAlive",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetLoginTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetLoginTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetMaxRetryCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxRetryCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOcspFailOpen() {
	_jsii_.InvokeVoid(
		s,
		"resetOcspFailOpen",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOktaUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetOktaUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOrganizationName() {
	_jsii_.InvokeVoid(
		s,
		"resetOrganizationName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetParams() {
	_jsii_.InvokeVoid(
		s,
		"resetParams",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPasscode() {
	_jsii_.InvokeVoid(
		s,
		"resetPasscode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPasscodeInPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetPasscodeInPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPort() {
	_jsii_.InvokeVoid(
		s,
		"resetPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPreviewFeaturesEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetPreviewFeaturesEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPrivateKeyPassphrase() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKeyPassphrase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetProfile() {
	_jsii_.InvokeVoid(
		s,
		"resetProfile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetRole() {
	_jsii_.InvokeVoid(
		s,
		"resetRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetSkipTomlFilePermissionVerification() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipTomlFilePermissionVerification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetTmpDirectoryPath() {
	_jsii_.InvokeVoid(
		s,
		"resetTmpDirectoryPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetToken() {
	_jsii_.InvokeVoid(
		s,
		"resetToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetTokenAccessor() {
	_jsii_.InvokeVoid(
		s,
		"resetTokenAccessor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetUseLegacyTomlFile() {
	_jsii_.InvokeVoid(
		s,
		"resetUseLegacyTomlFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetUser() {
	_jsii_.InvokeVoid(
		s,
		"resetUser",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetValidateDefaultParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetValidateDefaultParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetWarehouse() {
	_jsii_.InvokeVoid(
		s,
		"resetWarehouse",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnowflakeProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnowflakeProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnowflakeProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnowflakeProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnowflakeProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

