// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiintegration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.apiIntegration.ApiIntegration",
		reflect.TypeOf((*ApiIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiAllowedPrefixes", GoGetter: "ApiAllowedPrefixes"},
			_jsii_.MemberProperty{JsiiProperty: "apiAllowedPrefixesInput", GoGetter: "ApiAllowedPrefixesInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiAwsExternalId", GoGetter: "ApiAwsExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "apiAwsIamUserArn", GoGetter: "ApiAwsIamUserArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiAwsRoleArn", GoGetter: "ApiAwsRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiAwsRoleArnInput", GoGetter: "ApiAwsRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiBlockedPrefixes", GoGetter: "ApiBlockedPrefixes"},
			_jsii_.MemberProperty{JsiiProperty: "apiBlockedPrefixesInput", GoGetter: "ApiBlockedPrefixesInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiGcpServiceAccount", GoGetter: "ApiGcpServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "apiGcpServiceAccountInput", GoGetter: "ApiGcpServiceAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyInput", GoGetter: "ApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiProvider", GoGetter: "ApiProvider"},
			_jsii_.MemberProperty{JsiiProperty: "apiProviderInput", GoGetter: "ApiProviderInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureAdApplicationId", GoGetter: "AzureAdApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "azureAdApplicationIdInput", GoGetter: "AzureAdApplicationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureConsentUrl", GoGetter: "AzureConsentUrl"},
			_jsii_.MemberProperty{JsiiProperty: "azureMultiTenantAppName", GoGetter: "AzureMultiTenantAppName"},
			_jsii_.MemberProperty{JsiiProperty: "azureTenantId", GoGetter: "AzureTenantId"},
			_jsii_.MemberProperty{JsiiProperty: "azureTenantIdInput", GoGetter: "AzureTenantIdInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "googleAudience", GoGetter: "GoogleAudience"},
			_jsii_.MemberProperty{JsiiProperty: "googleAudienceInput", GoGetter: "GoogleAudienceInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiAwsRoleArn", GoMethod: "ResetApiAwsRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiBlockedPrefixes", GoMethod: "ResetApiBlockedPrefixes"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiGcpServiceAccount", GoMethod: "ResetApiGcpServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiKey", GoMethod: "ResetApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureAdApplicationId", GoMethod: "ResetAzureAdApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureTenantId", GoMethod: "ResetAzureTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleAudience", GoMethod: "ResetGoogleAudience"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ApiIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.apiIntegration.ApiIntegrationConfig",
		reflect.TypeOf((*ApiIntegrationConfig)(nil)).Elem(),
	)
}
