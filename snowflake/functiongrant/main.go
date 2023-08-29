// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functiongrant

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.functionGrant.FunctionGrant",
		reflect.TypeOf((*FunctionGrant)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "argumentDataTypes", GoGetter: "ArgumentDataTypes"},
			_jsii_.MemberProperty{JsiiProperty: "argumentDataTypesInput", GoGetter: "ArgumentDataTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "databaseNameInput", GoGetter: "DatabaseNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableMultipleGrants", GoGetter: "EnableMultipleGrants"},
			_jsii_.MemberProperty{JsiiProperty: "enableMultipleGrantsInput", GoGetter: "EnableMultipleGrantsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberProperty{JsiiProperty: "functionNameInput", GoGetter: "FunctionNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onAll", GoGetter: "OnAll"},
			_jsii_.MemberProperty{JsiiProperty: "onAllInput", GoGetter: "OnAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "onFuture", GoGetter: "OnFuture"},
			_jsii_.MemberProperty{JsiiProperty: "onFutureInput", GoGetter: "OnFutureInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privilege", GoGetter: "Privilege"},
			_jsii_.MemberProperty{JsiiProperty: "privilegeInput", GoGetter: "PrivilegeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetArgumentDataTypes", GoMethod: "ResetArgumentDataTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableMultipleGrants", GoMethod: "ResetEnableMultipleGrants"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionName", GoMethod: "ResetFunctionName"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnAll", GoMethod: "ResetOnAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnFuture", GoMethod: "ResetOnFuture"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivilege", GoMethod: "ResetPrivilege"},
			_jsii_.MemberMethod{JsiiMethod: "resetRevertOwnershipToRoleName", GoMethod: "ResetRevertOwnershipToRoleName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchemaName", GoMethod: "ResetSchemaName"},
			_jsii_.MemberMethod{JsiiMethod: "resetShares", GoMethod: "ResetShares"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithGrantOption", GoMethod: "ResetWithGrantOption"},
			_jsii_.MemberProperty{JsiiProperty: "revertOwnershipToRoleName", GoGetter: "RevertOwnershipToRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "revertOwnershipToRoleNameInput", GoGetter: "RevertOwnershipToRoleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "roles", GoGetter: "Roles"},
			_jsii_.MemberProperty{JsiiProperty: "rolesInput", GoGetter: "RolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "schemaName", GoGetter: "SchemaName"},
			_jsii_.MemberProperty{JsiiProperty: "schemaNameInput", GoGetter: "SchemaNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "shares", GoGetter: "Shares"},
			_jsii_.MemberProperty{JsiiProperty: "sharesInput", GoGetter: "SharesInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "withGrantOption", GoGetter: "WithGrantOption"},
			_jsii_.MemberProperty{JsiiProperty: "withGrantOptionInput", GoGetter: "WithGrantOptionInput"},
		},
		func() interface{} {
			j := jsiiProxy_FunctionGrant{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.functionGrant.FunctionGrantConfig",
		reflect.TypeOf((*FunctionGrantConfig)(nil)).Elem(),
	)
}
