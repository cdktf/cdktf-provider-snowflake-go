// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package passwordpolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicy",
		reflect.TypeOf((*PasswordPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ifNotExists", GoGetter: "IfNotExists"},
			_jsii_.MemberProperty{JsiiProperty: "ifNotExistsInput", GoGetter: "IfNotExistsInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "lockoutTimeMins", GoGetter: "LockoutTimeMins"},
			_jsii_.MemberProperty{JsiiProperty: "lockoutTimeMinsInput", GoGetter: "LockoutTimeMinsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxAgeDays", GoGetter: "MaxAgeDays"},
			_jsii_.MemberProperty{JsiiProperty: "maxAgeDaysInput", GoGetter: "MaxAgeDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxLength", GoGetter: "MaxLength"},
			_jsii_.MemberProperty{JsiiProperty: "maxLengthInput", GoGetter: "MaxLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "minLength", GoGetter: "MinLength"},
			_jsii_.MemberProperty{JsiiProperty: "minLengthInput", GoGetter: "MinLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "minLowerCaseChars", GoGetter: "MinLowerCaseChars"},
			_jsii_.MemberProperty{JsiiProperty: "minLowerCaseCharsInput", GoGetter: "MinLowerCaseCharsInput"},
			_jsii_.MemberProperty{JsiiProperty: "minNumericChars", GoGetter: "MinNumericChars"},
			_jsii_.MemberProperty{JsiiProperty: "minNumericCharsInput", GoGetter: "MinNumericCharsInput"},
			_jsii_.MemberProperty{JsiiProperty: "minSpecialChars", GoGetter: "MinSpecialChars"},
			_jsii_.MemberProperty{JsiiProperty: "minSpecialCharsInput", GoGetter: "MinSpecialCharsInput"},
			_jsii_.MemberProperty{JsiiProperty: "minUpperCaseChars", GoGetter: "MinUpperCaseChars"},
			_jsii_.MemberProperty{JsiiProperty: "minUpperCaseCharsInput", GoGetter: "MinUpperCaseCharsInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orReplace", GoGetter: "OrReplace"},
			_jsii_.MemberProperty{JsiiProperty: "orReplaceInput", GoGetter: "OrReplaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "qualifiedName", GoGetter: "QualifiedName"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIfNotExists", GoMethod: "ResetIfNotExists"},
			_jsii_.MemberMethod{JsiiMethod: "resetLockoutTimeMins", GoMethod: "ResetLockoutTimeMins"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxAgeDays", GoMethod: "ResetMaxAgeDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLength", GoMethod: "ResetMaxLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinLength", GoMethod: "ResetMinLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinLowerCaseChars", GoMethod: "ResetMinLowerCaseChars"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinNumericChars", GoMethod: "ResetMinNumericChars"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinSpecialChars", GoMethod: "ResetMinSpecialChars"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinUpperCaseChars", GoMethod: "ResetMinUpperCaseChars"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrReplace", GoMethod: "ResetOrReplace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_PasswordPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicyConfig",
		reflect.TypeOf((*PasswordPolicyConfig)(nil)).Elem(),
	)
}
