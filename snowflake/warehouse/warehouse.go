package warehouse

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.warehouse.Warehouse",
		reflect.TypeOf((*Warehouse)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoResume", GoGetter: "AutoResume"},
			_jsii_.MemberProperty{JsiiProperty: "autoResumeInput", GoGetter: "AutoResumeInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoSuspend", GoGetter: "AutoSuspend"},
			_jsii_.MemberProperty{JsiiProperty: "autoSuspendInput", GoGetter: "AutoSuspendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "initiallySuspended", GoGetter: "InitiallySuspended"},
			_jsii_.MemberProperty{JsiiProperty: "initiallySuspendedInput", GoGetter: "InitiallySuspendedInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxClusterCount", GoGetter: "MaxClusterCount"},
			_jsii_.MemberProperty{JsiiProperty: "maxClusterCountInput", GoGetter: "MaxClusterCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrencyLevel", GoGetter: "MaxConcurrencyLevel"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrencyLevelInput", GoGetter: "MaxConcurrencyLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "minClusterCount", GoGetter: "MinClusterCount"},
			_jsii_.MemberProperty{JsiiProperty: "minClusterCountInput", GoGetter: "MinClusterCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTag", GoMethod: "PutTag"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoResume", GoMethod: "ResetAutoResume"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoSuspend", GoMethod: "ResetAutoSuspend"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInitiallySuspended", GoMethod: "ResetInitiallySuspended"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxClusterCount", GoMethod: "ResetMaxClusterCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrencyLevel", GoMethod: "ResetMaxConcurrencyLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinClusterCount", GoMethod: "ResetMinClusterCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceMonitor", GoMethod: "ResetResourceMonitor"},
			_jsii_.MemberMethod{JsiiMethod: "resetScalingPolicy", GoMethod: "ResetScalingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatementQueuedTimeoutInSeconds", GoMethod: "ResetStatementQueuedTimeoutInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatementTimeoutInSeconds", GoMethod: "ResetStatementTimeoutInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTag", GoMethod: "ResetTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetWaitForProvisioning", GoMethod: "ResetWaitForProvisioning"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarehouseSize", GoMethod: "ResetWarehouseSize"},
			_jsii_.MemberProperty{JsiiProperty: "resourceMonitor", GoGetter: "ResourceMonitor"},
			_jsii_.MemberProperty{JsiiProperty: "resourceMonitorInput", GoGetter: "ResourceMonitorInput"},
			_jsii_.MemberProperty{JsiiProperty: "scalingPolicy", GoGetter: "ScalingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "scalingPolicyInput", GoGetter: "ScalingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "statementQueuedTimeoutInSeconds", GoGetter: "StatementQueuedTimeoutInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "statementQueuedTimeoutInSecondsInput", GoGetter: "StatementQueuedTimeoutInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "statementTimeoutInSeconds", GoGetter: "StatementTimeoutInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "statementTimeoutInSecondsInput", GoGetter: "StatementTimeoutInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberProperty{JsiiProperty: "tagInput", GoGetter: "TagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "waitForProvisioning", GoGetter: "WaitForProvisioning"},
			_jsii_.MemberProperty{JsiiProperty: "waitForProvisioningInput", GoGetter: "WaitForProvisioningInput"},
			_jsii_.MemberProperty{JsiiProperty: "warehouseSize", GoGetter: "WarehouseSize"},
			_jsii_.MemberProperty{JsiiProperty: "warehouseSizeInput", GoGetter: "WarehouseSizeInput"},
		},
		func() interface{} {
			j := jsiiProxy_Warehouse{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.warehouse.WarehouseConfig",
		reflect.TypeOf((*WarehouseConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.warehouse.WarehouseTag",
		reflect.TypeOf((*WarehouseTag)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.warehouse.WarehouseTagList",
		reflect.TypeOf((*WarehouseTagList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_WarehouseTagList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.warehouse.WarehouseTagOutputReference",
		reflect.TypeOf((*WarehouseTagOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabase", GoMethod: "ResetDatabase"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchema", GoMethod: "ResetSchema"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_WarehouseTagOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
