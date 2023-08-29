// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemonitor

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitor",
		reflect.TypeOf((*ResourceMonitor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creditQuota", GoGetter: "CreditQuota"},
			_jsii_.MemberProperty{JsiiProperty: "creditQuotaInput", GoGetter: "CreditQuotaInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "endTimestamp", GoGetter: "EndTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "endTimestampInput", GoGetter: "EndTimestampInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequency", GoGetter: "Frequency"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyInput", GoGetter: "FrequencyInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notifyTriggers", GoGetter: "NotifyTriggers"},
			_jsii_.MemberProperty{JsiiProperty: "notifyTriggersInput", GoGetter: "NotifyTriggersInput"},
			_jsii_.MemberProperty{JsiiProperty: "notifyUsers", GoGetter: "NotifyUsers"},
			_jsii_.MemberProperty{JsiiProperty: "notifyUsersInput", GoGetter: "NotifyUsersInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreditQuota", GoMethod: "ResetCreditQuota"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndTimestamp", GoMethod: "ResetEndTimestamp"},
			_jsii_.MemberMethod{JsiiMethod: "resetFrequency", GoMethod: "ResetFrequency"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotifyTriggers", GoMethod: "ResetNotifyTriggers"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotifyUsers", GoMethod: "ResetNotifyUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetForAccount", GoMethod: "ResetSetForAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartTimestamp", GoMethod: "ResetStartTimestamp"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuspendImmediateTrigger", GoMethod: "ResetSuspendImmediateTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuspendImmediateTriggers", GoMethod: "ResetSuspendImmediateTriggers"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuspendTrigger", GoMethod: "ResetSuspendTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuspendTriggers", GoMethod: "ResetSuspendTriggers"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarehouses", GoMethod: "ResetWarehouses"},
			_jsii_.MemberProperty{JsiiProperty: "setForAccount", GoGetter: "SetForAccount"},
			_jsii_.MemberProperty{JsiiProperty: "setForAccountInput", GoGetter: "SetForAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "startTimestamp", GoGetter: "StartTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "startTimestampInput", GoGetter: "StartTimestampInput"},
			_jsii_.MemberProperty{JsiiProperty: "suspendImmediateTrigger", GoGetter: "SuspendImmediateTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "suspendImmediateTriggerInput", GoGetter: "SuspendImmediateTriggerInput"},
			_jsii_.MemberProperty{JsiiProperty: "suspendImmediateTriggers", GoGetter: "SuspendImmediateTriggers"},
			_jsii_.MemberProperty{JsiiProperty: "suspendImmediateTriggersInput", GoGetter: "SuspendImmediateTriggersInput"},
			_jsii_.MemberProperty{JsiiProperty: "suspendTrigger", GoGetter: "SuspendTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "suspendTriggerInput", GoGetter: "SuspendTriggerInput"},
			_jsii_.MemberProperty{JsiiProperty: "suspendTriggers", GoGetter: "SuspendTriggers"},
			_jsii_.MemberProperty{JsiiProperty: "suspendTriggersInput", GoGetter: "SuspendTriggersInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "warehouses", GoGetter: "Warehouses"},
			_jsii_.MemberProperty{JsiiProperty: "warehousesInput", GoGetter: "WarehousesInput"},
		},
		func() interface{} {
			j := jsiiProxy_ResourceMonitor{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitorConfig",
		reflect.TypeOf((*ResourceMonitorConfig)(nil)).Elem(),
	)
}
