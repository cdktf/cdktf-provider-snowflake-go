// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externaloauthintegration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegration",
		reflect.TypeOf((*ExternalOauthIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedRoles", GoGetter: "AllowedRoles"},
			_jsii_.MemberProperty{JsiiProperty: "allowedRolesInput", GoGetter: "AllowedRolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "anyRoleMode", GoGetter: "AnyRoleMode"},
			_jsii_.MemberProperty{JsiiProperty: "anyRoleModeInput", GoGetter: "AnyRoleModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "audienceUrls", GoGetter: "AudienceUrls"},
			_jsii_.MemberProperty{JsiiProperty: "audienceUrlsInput", GoGetter: "AudienceUrlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "blockedRoles", GoGetter: "BlockedRoles"},
			_jsii_.MemberProperty{JsiiProperty: "blockedRolesInput", GoGetter: "BlockedRolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdOn", GoGetter: "CreatedOn"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuer", GoGetter: "Issuer"},
			_jsii_.MemberProperty{JsiiProperty: "issuerInput", GoGetter: "IssuerInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwsKeysUrls", GoGetter: "JwsKeysUrls"},
			_jsii_.MemberProperty{JsiiProperty: "jwsKeysUrlsInput", GoGetter: "JwsKeysUrlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedRoles", GoMethod: "ResetAllowedRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnyRoleMode", GoMethod: "ResetAnyRoleMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudienceUrls", GoMethod: "ResetAudienceUrls"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockedRoles", GoMethod: "ResetBlockedRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetJwsKeysUrls", GoMethod: "ResetJwsKeysUrls"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRsaPublicKey", GoMethod: "ResetRsaPublicKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetRsaPublicKey2", GoMethod: "ResetRsaPublicKey2"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopeDelimiter", GoMethod: "ResetScopeDelimiter"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopeMappingAttribute", GoMethod: "ResetScopeMappingAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "rsaPublicKey", GoGetter: "RsaPublicKey"},
			_jsii_.MemberProperty{JsiiProperty: "rsaPublicKey2", GoGetter: "RsaPublicKey2"},
			_jsii_.MemberProperty{JsiiProperty: "rsaPublicKey2Input", GoGetter: "RsaPublicKey2Input"},
			_jsii_.MemberProperty{JsiiProperty: "rsaPublicKeyInput", GoGetter: "RsaPublicKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopeDelimiter", GoGetter: "ScopeDelimiter"},
			_jsii_.MemberProperty{JsiiProperty: "scopeDelimiterInput", GoGetter: "ScopeDelimiterInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopeMappingAttribute", GoGetter: "ScopeMappingAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "scopeMappingAttributeInput", GoGetter: "ScopeMappingAttributeInput"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeUserMappingAttribute", GoGetter: "SnowflakeUserMappingAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeUserMappingAttributeInput", GoGetter: "SnowflakeUserMappingAttributeInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tokenUserMappingClaims", GoGetter: "TokenUserMappingClaims"},
			_jsii_.MemberProperty{JsiiProperty: "tokenUserMappingClaimsInput", GoGetter: "TokenUserMappingClaimsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_ExternalOauthIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegrationConfig",
		reflect.TypeOf((*ExternalOauthIntegrationConfig)(nil)).Elem(),
	)
}
