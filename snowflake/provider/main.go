// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.provider.SnowflakeProvider",
		reflect.TypeOf((*SnowflakeProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberProperty{JsiiProperty: "accountInput", GoGetter: "AccountInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "authenticator", GoGetter: "Authenticator"},
			_jsii_.MemberProperty{JsiiProperty: "authenticatorInput", GoGetter: "AuthenticatorInput"},
			_jsii_.MemberProperty{JsiiProperty: "browserAuth", GoGetter: "BrowserAuth"},
			_jsii_.MemberProperty{JsiiProperty: "browserAuthInput", GoGetter: "BrowserAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientIp", GoGetter: "ClientIp"},
			_jsii_.MemberProperty{JsiiProperty: "clientIpInput", GoGetter: "ClientIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientRequestMfaToken", GoGetter: "ClientRequestMfaToken"},
			_jsii_.MemberProperty{JsiiProperty: "clientRequestMfaTokenInput", GoGetter: "ClientRequestMfaTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientStoreTemporaryCredential", GoGetter: "ClientStoreTemporaryCredential"},
			_jsii_.MemberProperty{JsiiProperty: "clientStoreTemporaryCredentialInput", GoGetter: "ClientStoreTemporaryCredentialInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientTimeout", GoGetter: "ClientTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "clientTimeoutInput", GoGetter: "ClientTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "disableQueryContextCache", GoGetter: "DisableQueryContextCache"},
			_jsii_.MemberProperty{JsiiProperty: "disableQueryContextCacheInput", GoGetter: "DisableQueryContextCacheInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableTelemetry", GoGetter: "DisableTelemetry"},
			_jsii_.MemberProperty{JsiiProperty: "disableTelemetryInput", GoGetter: "DisableTelemetryInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalBrowserTimeout", GoGetter: "ExternalBrowserTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "externalBrowserTimeoutInput", GoGetter: "ExternalBrowserTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "hostInput", GoGetter: "HostInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecureMode", GoGetter: "InsecureMode"},
			_jsii_.MemberProperty{JsiiProperty: "insecureModeInput", GoGetter: "InsecureModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwtClientTimeout", GoGetter: "JwtClientTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "jwtClientTimeoutInput", GoGetter: "JwtClientTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwtExpireTimeout", GoGetter: "JwtExpireTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "jwtExpireTimeoutInput", GoGetter: "JwtExpireTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "keepSessionAlive", GoGetter: "KeepSessionAlive"},
			_jsii_.MemberProperty{JsiiProperty: "keepSessionAliveInput", GoGetter: "KeepSessionAliveInput"},
			_jsii_.MemberProperty{JsiiProperty: "loginTimeout", GoGetter: "LoginTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "loginTimeoutInput", GoGetter: "LoginTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oauthAccessToken", GoGetter: "OauthAccessToken"},
			_jsii_.MemberProperty{JsiiProperty: "oauthAccessTokenInput", GoGetter: "OauthAccessTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "oauthClientId", GoGetter: "OauthClientId"},
			_jsii_.MemberProperty{JsiiProperty: "oauthClientIdInput", GoGetter: "OauthClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "oauthClientSecret", GoGetter: "OauthClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "oauthClientSecretInput", GoGetter: "OauthClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "oauthEndpoint", GoGetter: "OauthEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "oauthEndpointInput", GoGetter: "OauthEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "oauthRedirectUrl", GoGetter: "OauthRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "oauthRedirectUrlInput", GoGetter: "OauthRedirectUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "oauthRefreshToken", GoGetter: "OauthRefreshToken"},
			_jsii_.MemberProperty{JsiiProperty: "oauthRefreshTokenInput", GoGetter: "OauthRefreshTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "ocspFailOpen", GoGetter: "OcspFailOpen"},
			_jsii_.MemberProperty{JsiiProperty: "ocspFailOpenInput", GoGetter: "OcspFailOpenInput"},
			_jsii_.MemberProperty{JsiiProperty: "oktaUrl", GoGetter: "OktaUrl"},
			_jsii_.MemberProperty{JsiiProperty: "oktaUrlInput", GoGetter: "OktaUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "params", GoGetter: "Params"},
			_jsii_.MemberProperty{JsiiProperty: "paramsInput", GoGetter: "ParamsInput"},
			_jsii_.MemberProperty{JsiiProperty: "passcode", GoGetter: "Passcode"},
			_jsii_.MemberProperty{JsiiProperty: "passcodeInPassword", GoGetter: "PasscodeInPassword"},
			_jsii_.MemberProperty{JsiiProperty: "passcodeInPasswordInput", GoGetter: "PasscodeInPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "passcodeInput", GoGetter: "PasscodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPassphrase", GoGetter: "PrivateKeyPassphrase"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPassphraseInput", GoGetter: "PrivateKeyPassphraseInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPath", GoGetter: "PrivateKeyPath"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPathInput", GoGetter: "PrivateKeyPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "profile", GoGetter: "Profile"},
			_jsii_.MemberProperty{JsiiProperty: "profileInput", GoGetter: "ProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeout", GoGetter: "RequestTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeoutInput", GoGetter: "RequestTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccount", GoMethod: "ResetAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthenticator", GoMethod: "ResetAuthenticator"},
			_jsii_.MemberMethod{JsiiMethod: "resetBrowserAuth", GoMethod: "ResetBrowserAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientIp", GoMethod: "ResetClientIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientRequestMfaToken", GoMethod: "ResetClientRequestMfaToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientStoreTemporaryCredential", GoMethod: "ResetClientStoreTemporaryCredential"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTimeout", GoMethod: "ResetClientTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableQueryContextCache", GoMethod: "ResetDisableQueryContextCache"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableTelemetry", GoMethod: "ResetDisableTelemetry"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalBrowserTimeout", GoMethod: "ResetExternalBrowserTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetHost", GoMethod: "ResetHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureMode", GoMethod: "ResetInsecureMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetJwtClientTimeout", GoMethod: "ResetJwtClientTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetJwtExpireTimeout", GoMethod: "ResetJwtExpireTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeepSessionAlive", GoMethod: "ResetKeepSessionAlive"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginTimeout", GoMethod: "ResetLoginTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthAccessToken", GoMethod: "ResetOauthAccessToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthClientId", GoMethod: "ResetOauthClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthClientSecret", GoMethod: "ResetOauthClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthEndpoint", GoMethod: "ResetOauthEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthRedirectUrl", GoMethod: "ResetOauthRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthRefreshToken", GoMethod: "ResetOauthRefreshToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetOcspFailOpen", GoMethod: "ResetOcspFailOpen"},
			_jsii_.MemberMethod{JsiiMethod: "resetOktaUrl", GoMethod: "ResetOktaUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParams", GoMethod: "ResetParams"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasscode", GoMethod: "ResetPasscode"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasscodeInPassword", GoMethod: "ResetPasscodeInPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetPort", GoMethod: "ResetPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKey", GoMethod: "ResetPrivateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKeyPassphrase", GoMethod: "ResetPrivateKeyPassphrase"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKeyPath", GoMethod: "ResetPrivateKeyPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetProfile", GoMethod: "ResetProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocol", GoMethod: "ResetProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestTimeout", GoMethod: "ResetRequestTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetRole", GoMethod: "ResetRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionParams", GoMethod: "ResetSessionParams"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenAccessor", GoMethod: "ResetTokenAccessor"},
			_jsii_.MemberMethod{JsiiMethod: "resetUser", GoMethod: "ResetUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsername", GoMethod: "ResetUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidateDefaultParameters", GoMethod: "ResetValidateDefaultParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarehouse", GoMethod: "ResetWarehouse"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "roleInput", GoGetter: "RoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionParams", GoGetter: "SessionParams"},
			_jsii_.MemberProperty{JsiiProperty: "sessionParamsInput", GoGetter: "SessionParamsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenAccessor", GoGetter: "TokenAccessor"},
			_jsii_.MemberProperty{JsiiProperty: "tokenAccessorInput", GoGetter: "TokenAccessorInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userInput", GoGetter: "UserInput"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "validateDefaultParameters", GoGetter: "ValidateDefaultParameters"},
			_jsii_.MemberProperty{JsiiProperty: "validateDefaultParametersInput", GoGetter: "ValidateDefaultParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "warehouse", GoGetter: "Warehouse"},
			_jsii_.MemberProperty{JsiiProperty: "warehouseInput", GoGetter: "WarehouseInput"},
		},
		func() interface{} {
			j := jsiiProxy_SnowflakeProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.provider.SnowflakeProviderConfig",
		reflect.TypeOf((*SnowflakeProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.provider.SnowflakeProviderTokenAccessor",
		reflect.TypeOf((*SnowflakeProviderTokenAccessor)(nil)).Elem(),
	)
}
