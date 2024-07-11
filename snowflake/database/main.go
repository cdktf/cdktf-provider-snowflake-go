// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package database

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.database.Database",
		reflect.TypeOf((*Database)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "catalog", GoGetter: "Catalog"},
			_jsii_.MemberProperty{JsiiProperty: "catalogInput", GoGetter: "CatalogInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dataRetentionTimeInDays", GoGetter: "DataRetentionTimeInDays"},
			_jsii_.MemberProperty{JsiiProperty: "dataRetentionTimeInDaysInput", GoGetter: "DataRetentionTimeInDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDdlCollation", GoGetter: "DefaultDdlCollation"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDdlCollationInput", GoGetter: "DefaultDdlCollationInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableConsoleOutput", GoGetter: "EnableConsoleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "enableConsoleOutputInput", GoGetter: "EnableConsoleOutputInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalVolume", GoGetter: "ExternalVolume"},
			_jsii_.MemberProperty{JsiiProperty: "externalVolumeInput", GoGetter: "ExternalVolumeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "isTransient", GoGetter: "IsTransient"},
			_jsii_.MemberProperty{JsiiProperty: "isTransientInput", GoGetter: "IsTransientInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "logLevelInput", GoGetter: "LogLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxDataExtensionTimeInDays", GoGetter: "MaxDataExtensionTimeInDays"},
			_jsii_.MemberProperty{JsiiProperty: "maxDataExtensionTimeInDaysInput", GoGetter: "MaxDataExtensionTimeInDaysInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putReplication", GoMethod: "PutReplication"},
			_jsii_.MemberProperty{JsiiProperty: "quotedIdentifiersIgnoreCase", GoGetter: "QuotedIdentifiersIgnoreCase"},
			_jsii_.MemberProperty{JsiiProperty: "quotedIdentifiersIgnoreCaseInput", GoGetter: "QuotedIdentifiersIgnoreCaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharactersInput", GoGetter: "ReplaceInvalidCharactersInput"},
			_jsii_.MemberProperty{JsiiProperty: "replication", GoGetter: "Replication"},
			_jsii_.MemberProperty{JsiiProperty: "replicationInput", GoGetter: "ReplicationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCatalog", GoMethod: "ResetCatalog"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataRetentionTimeInDays", GoMethod: "ResetDataRetentionTimeInDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultDdlCollation", GoMethod: "ResetDefaultDdlCollation"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableConsoleOutput", GoMethod: "ResetEnableConsoleOutput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalVolume", GoMethod: "ResetExternalVolume"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsTransient", GoMethod: "ResetIsTransient"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogLevel", GoMethod: "ResetLogLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxDataExtensionTimeInDays", GoMethod: "ResetMaxDataExtensionTimeInDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetQuotedIdentifiersIgnoreCase", GoMethod: "ResetQuotedIdentifiersIgnoreCase"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplaceInvalidCharacters", GoMethod: "ResetReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplication", GoMethod: "ResetReplication"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageSerializationPolicy", GoMethod: "ResetStorageSerializationPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuspendTaskAfterNumFailures", GoMethod: "ResetSuspendTaskAfterNumFailures"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaskAutoRetryAttempts", GoMethod: "ResetTaskAutoRetryAttempts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTraceLevel", GoMethod: "ResetTraceLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserTaskManagedInitialWarehouseSize", GoMethod: "ResetUserTaskManagedInitialWarehouseSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserTaskMinimumTriggerIntervalInSeconds", GoMethod: "ResetUserTaskMinimumTriggerIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserTaskTimeoutMs", GoMethod: "ResetUserTaskTimeoutMs"},
			_jsii_.MemberProperty{JsiiProperty: "storageSerializationPolicy", GoGetter: "StorageSerializationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "storageSerializationPolicyInput", GoGetter: "StorageSerializationPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "suspendTaskAfterNumFailures", GoGetter: "SuspendTaskAfterNumFailures"},
			_jsii_.MemberProperty{JsiiProperty: "suspendTaskAfterNumFailuresInput", GoGetter: "SuspendTaskAfterNumFailuresInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "taskAutoRetryAttempts", GoGetter: "TaskAutoRetryAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "taskAutoRetryAttemptsInput", GoGetter: "TaskAutoRetryAttemptsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "traceLevel", GoGetter: "TraceLevel"},
			_jsii_.MemberProperty{JsiiProperty: "traceLevelInput", GoGetter: "TraceLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "userTaskManagedInitialWarehouseSize", GoGetter: "UserTaskManagedInitialWarehouseSize"},
			_jsii_.MemberProperty{JsiiProperty: "userTaskManagedInitialWarehouseSizeInput", GoGetter: "UserTaskManagedInitialWarehouseSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "userTaskMinimumTriggerIntervalInSeconds", GoGetter: "UserTaskMinimumTriggerIntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "userTaskMinimumTriggerIntervalInSecondsInput", GoGetter: "UserTaskMinimumTriggerIntervalInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "userTaskTimeoutMs", GoGetter: "UserTaskTimeoutMs"},
			_jsii_.MemberProperty{JsiiProperty: "userTaskTimeoutMsInput", GoGetter: "UserTaskTimeoutMsInput"},
		},
		func() interface{} {
			j := jsiiProxy_Database{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.database.DatabaseConfig",
		reflect.TypeOf((*DatabaseConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.database.DatabaseReplication",
		reflect.TypeOf((*DatabaseReplication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.database.DatabaseReplicationEnableToAccount",
		reflect.TypeOf((*DatabaseReplicationEnableToAccount)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.database.DatabaseReplicationEnableToAccountList",
		reflect.TypeOf((*DatabaseReplicationEnableToAccountList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseReplicationEnableToAccountList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.database.DatabaseReplicationEnableToAccountOutputReference",
		reflect.TypeOf((*DatabaseReplicationEnableToAccountOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountIdentifier", GoGetter: "AccountIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdentifierInput", GoGetter: "AccountIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithFailover", GoMethod: "ResetWithFailover"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "withFailover", GoGetter: "WithFailover"},
			_jsii_.MemberProperty{JsiiProperty: "withFailoverInput", GoGetter: "WithFailoverInput"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseReplicationEnableToAccountOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.database.DatabaseReplicationOutputReference",
		reflect.TypeOf((*DatabaseReplicationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableToAccount", GoGetter: "EnableToAccount"},
			_jsii_.MemberProperty{JsiiProperty: "enableToAccountInput", GoGetter: "EnableToAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreEditionCheck", GoGetter: "IgnoreEditionCheck"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreEditionCheckInput", GoGetter: "IgnoreEditionCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putEnableToAccount", GoMethod: "PutEnableToAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreEditionCheck", GoMethod: "ResetIgnoreEditionCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseReplicationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
