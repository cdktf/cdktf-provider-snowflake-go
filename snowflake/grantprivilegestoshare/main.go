// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoshare

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShare",
		reflect.TypeOf((*GrantPrivilegesToShare)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onAllTablesInSchema", GoGetter: "OnAllTablesInSchema"},
			_jsii_.MemberProperty{JsiiProperty: "onAllTablesInSchemaInput", GoGetter: "OnAllTablesInSchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "onDatabase", GoGetter: "OnDatabase"},
			_jsii_.MemberProperty{JsiiProperty: "onDatabaseInput", GoGetter: "OnDatabaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "onSchema", GoGetter: "OnSchema"},
			_jsii_.MemberProperty{JsiiProperty: "onSchemaInput", GoGetter: "OnSchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "onTable", GoGetter: "OnTable"},
			_jsii_.MemberProperty{JsiiProperty: "onTableInput", GoGetter: "OnTableInput"},
			_jsii_.MemberProperty{JsiiProperty: "onTag", GoGetter: "OnTag"},
			_jsii_.MemberProperty{JsiiProperty: "onTagInput", GoGetter: "OnTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "onView", GoGetter: "OnView"},
			_jsii_.MemberProperty{JsiiProperty: "onViewInput", GoGetter: "OnViewInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privileges", GoGetter: "Privileges"},
			_jsii_.MemberProperty{JsiiProperty: "privilegesInput", GoGetter: "PrivilegesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnAllTablesInSchema", GoMethod: "ResetOnAllTablesInSchema"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnDatabase", GoMethod: "ResetOnDatabase"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnSchema", GoMethod: "ResetOnSchema"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnTable", GoMethod: "ResetOnTable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnTag", GoMethod: "ResetOnTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnView", GoMethod: "ResetOnView"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "toShare", GoGetter: "ToShare"},
			_jsii_.MemberProperty{JsiiProperty: "toShareInput", GoGetter: "ToShareInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_GrantPrivilegesToShare{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShareConfig",
		reflect.TypeOf((*GrantPrivilegesToShareConfig)(nil)).Elem(),
	)
}
