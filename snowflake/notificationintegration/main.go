// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationintegration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegration",
		reflect.TypeOf((*NotificationIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "awsSnsExternalId", GoGetter: "AwsSnsExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "awsSnsIamUserArn", GoGetter: "AwsSnsIamUserArn"},
			_jsii_.MemberProperty{JsiiProperty: "awsSnsRoleArn", GoGetter: "AwsSnsRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "awsSnsRoleArnInput", GoGetter: "AwsSnsRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSnsTopicArn", GoGetter: "AwsSnsTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "awsSnsTopicArnInput", GoGetter: "AwsSnsTopicArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSqsArn", GoGetter: "AwsSqsArn"},
			_jsii_.MemberProperty{JsiiProperty: "awsSqsArnInput", GoGetter: "AwsSqsArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSqsExternalId", GoGetter: "AwsSqsExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "awsSqsIamUserArn", GoGetter: "AwsSqsIamUserArn"},
			_jsii_.MemberProperty{JsiiProperty: "awsSqsRoleArn", GoGetter: "AwsSqsRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "awsSqsRoleArnInput", GoGetter: "AwsSqsRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureStorageQueuePrimaryUri", GoGetter: "AzureStorageQueuePrimaryUri"},
			_jsii_.MemberProperty{JsiiProperty: "azureStorageQueuePrimaryUriInput", GoGetter: "AzureStorageQueuePrimaryUriInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "direction", GoGetter: "Direction"},
			_jsii_.MemberProperty{JsiiProperty: "directionInput", GoGetter: "DirectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "gcpPubsubServiceAccount", GoGetter: "GcpPubsubServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "gcpPubsubSubscriptionName", GoGetter: "GcpPubsubSubscriptionName"},
			_jsii_.MemberProperty{JsiiProperty: "gcpPubsubSubscriptionNameInput", GoGetter: "GcpPubsubSubscriptionNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "gcpPubsubTopicName", GoGetter: "GcpPubsubTopicName"},
			_jsii_.MemberProperty{JsiiProperty: "gcpPubsubTopicNameInput", GoGetter: "GcpPubsubTopicNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationProvider", GoGetter: "NotificationProvider"},
			_jsii_.MemberProperty{JsiiProperty: "notificationProviderInput", GoGetter: "NotificationProviderInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSnsRoleArn", GoMethod: "ResetAwsSnsRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSnsTopicArn", GoMethod: "ResetAwsSnsTopicArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSqsArn", GoMethod: "ResetAwsSqsArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSqsRoleArn", GoMethod: "ResetAwsSqsRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureStorageQueuePrimaryUri", GoMethod: "ResetAzureStorageQueuePrimaryUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureTenantId", GoMethod: "ResetAzureTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirection", GoMethod: "ResetDirection"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetGcpPubsubSubscriptionName", GoMethod: "ResetGcpPubsubSubscriptionName"},
			_jsii_.MemberMethod{JsiiMethod: "resetGcpPubsubTopicName", GoMethod: "ResetGcpPubsubTopicName"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationProvider", GoMethod: "ResetNotificationProvider"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_NotificationIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegrationConfig",
		reflect.TypeOf((*NotificationIntegrationConfig)(nil)).Elem(),
	)
}
