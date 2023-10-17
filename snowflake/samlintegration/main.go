// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package samlintegration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.samlIntegration.SamlIntegration",
		reflect.TypeOf((*SamlIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2EnableSpInitiated", GoMethod: "ResetSaml2EnableSpInitiated"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2ForceAuthn", GoMethod: "ResetSaml2ForceAuthn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2PostLogoutRedirectUrl", GoMethod: "ResetSaml2PostLogoutRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2RequestedNameidFormat", GoMethod: "ResetSaml2RequestedNameidFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2SignRequest", GoMethod: "ResetSaml2SignRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2SnowflakeAcsUrl", GoMethod: "ResetSaml2SnowflakeAcsUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2SnowflakeIssuerUrl", GoMethod: "ResetSaml2SnowflakeIssuerUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2SnowflakeX509Cert", GoMethod: "ResetSaml2SnowflakeX509Cert"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml2SpInitiatedLoginPageLabel", GoMethod: "ResetSaml2SpInitiatedLoginPageLabel"},
			_jsii_.MemberProperty{JsiiProperty: "saml2DigestMethodsUsed", GoGetter: "Saml2DigestMethodsUsed"},
			_jsii_.MemberProperty{JsiiProperty: "saml2EnableSpInitiated", GoGetter: "Saml2EnableSpInitiated"},
			_jsii_.MemberProperty{JsiiProperty: "saml2EnableSpInitiatedInput", GoGetter: "Saml2EnableSpInitiatedInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2ForceAuthn", GoGetter: "Saml2ForceAuthn"},
			_jsii_.MemberProperty{JsiiProperty: "saml2ForceAuthnInput", GoGetter: "Saml2ForceAuthnInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2Issuer", GoGetter: "Saml2Issuer"},
			_jsii_.MemberProperty{JsiiProperty: "saml2IssuerInput", GoGetter: "Saml2IssuerInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2PostLogoutRedirectUrl", GoGetter: "Saml2PostLogoutRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "saml2PostLogoutRedirectUrlInput", GoGetter: "Saml2PostLogoutRedirectUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2Provider", GoGetter: "Saml2Provider"},
			_jsii_.MemberProperty{JsiiProperty: "saml2ProviderInput", GoGetter: "Saml2ProviderInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2RequestedNameidFormat", GoGetter: "Saml2RequestedNameidFormat"},
			_jsii_.MemberProperty{JsiiProperty: "saml2RequestedNameidFormatInput", GoGetter: "Saml2RequestedNameidFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SignatureMethodsUsed", GoGetter: "Saml2SignatureMethodsUsed"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SignRequest", GoGetter: "Saml2SignRequest"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SignRequestInput", GoGetter: "Saml2SignRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SnowflakeAcsUrl", GoGetter: "Saml2SnowflakeAcsUrl"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SnowflakeAcsUrlInput", GoGetter: "Saml2SnowflakeAcsUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SnowflakeIssuerUrl", GoGetter: "Saml2SnowflakeIssuerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SnowflakeIssuerUrlInput", GoGetter: "Saml2SnowflakeIssuerUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SnowflakeMetadata", GoGetter: "Saml2SnowflakeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SnowflakeX509Cert", GoGetter: "Saml2SnowflakeX509Cert"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SnowflakeX509CertInput", GoGetter: "Saml2SnowflakeX509CertInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SpInitiatedLoginPageLabel", GoGetter: "Saml2SpInitiatedLoginPageLabel"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SpInitiatedLoginPageLabelInput", GoGetter: "Saml2SpInitiatedLoginPageLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SsoUrl", GoGetter: "Saml2SsoUrl"},
			_jsii_.MemberProperty{JsiiProperty: "saml2SsoUrlInput", GoGetter: "Saml2SsoUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "saml2X509Cert", GoGetter: "Saml2X509Cert"},
			_jsii_.MemberProperty{JsiiProperty: "saml2X509CertInput", GoGetter: "Saml2X509CertInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SamlIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.samlIntegration.SamlIntegrationConfig",
		reflect.TypeOf((*SamlIntegrationConfig)(nil)).Elem(),
	)
}
