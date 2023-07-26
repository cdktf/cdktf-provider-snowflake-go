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
			_jsii_.MemberProperty{JsiiProperty: "browserAuth", GoGetter: "BrowserAuth"},
			_jsii_.MemberProperty{JsiiProperty: "browserAuthInput", GoGetter: "BrowserAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "hostInput", GoGetter: "HostInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecureMode", GoGetter: "InsecureMode"},
			_jsii_.MemberProperty{JsiiProperty: "insecureModeInput", GoGetter: "InsecureModeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAccount", GoMethod: "ResetAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetBrowserAuth", GoMethod: "ResetBrowserAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetHost", GoMethod: "ResetHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureMode", GoMethod: "ResetInsecureMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthAccessToken", GoMethod: "ResetOauthAccessToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthClientId", GoMethod: "ResetOauthClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthClientSecret", GoMethod: "ResetOauthClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthEndpoint", GoMethod: "ResetOauthEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthRedirectUrl", GoMethod: "ResetOauthRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthRefreshToken", GoMethod: "ResetOauthRefreshToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetPort", GoMethod: "ResetPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKey", GoMethod: "ResetPrivateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKeyPassphrase", GoMethod: "ResetPrivateKeyPassphrase"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKeyPath", GoMethod: "ResetPrivateKeyPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetProfile", GoMethod: "ResetProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocol", GoMethod: "ResetProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRole", GoMethod: "ResetRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionParams", GoMethod: "ResetSessionParams"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsername", GoMethod: "ResetUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarehouse", GoMethod: "ResetWarehouse"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "roleInput", GoGetter: "RoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionParams", GoGetter: "SessionParams"},
			_jsii_.MemberProperty{JsiiProperty: "sessionParamsInput", GoGetter: "SessionParamsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
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
}
