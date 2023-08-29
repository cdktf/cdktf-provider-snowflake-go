// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.user.User",
		reflect.TypeOf((*User)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNamespace", GoGetter: "DefaultNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNamespaceInput", GoGetter: "DefaultNamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRole", GoGetter: "DefaultRole"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRoleInput", GoGetter: "DefaultRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSecondaryRoles", GoGetter: "DefaultSecondaryRoles"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSecondaryRolesInput", GoGetter: "DefaultSecondaryRolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultWarehouse", GoGetter: "DefaultWarehouse"},
			_jsii_.MemberProperty{JsiiProperty: "defaultWarehouseInput", GoGetter: "DefaultWarehouseInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disabled", GoGetter: "Disabled"},
			_jsii_.MemberProperty{JsiiProperty: "disabledInput", GoGetter: "DisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "firstName", GoGetter: "FirstName"},
			_jsii_.MemberProperty{JsiiProperty: "firstNameInput", GoGetter: "FirstNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hasRsaPublicKey", GoGetter: "HasRsaPublicKey"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastName", GoGetter: "LastName"},
			_jsii_.MemberProperty{JsiiProperty: "lastNameInput", GoGetter: "LastNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loginName", GoGetter: "LoginName"},
			_jsii_.MemberProperty{JsiiProperty: "loginNameInput", GoGetter: "LoginNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "mustChangePassword", GoGetter: "MustChangePassword"},
			_jsii_.MemberProperty{JsiiProperty: "mustChangePasswordInput", GoGetter: "MustChangePasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultNamespace", GoMethod: "ResetDefaultNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRole", GoMethod: "ResetDefaultRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultSecondaryRoles", GoMethod: "ResetDefaultSecondaryRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultWarehouse", GoMethod: "ResetDefaultWarehouse"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisabled", GoMethod: "ResetDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmail", GoMethod: "ResetEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetFirstName", GoMethod: "ResetFirstName"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLastName", GoMethod: "ResetLastName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginName", GoMethod: "ResetLoginName"},
			_jsii_.MemberMethod{JsiiMethod: "resetMustChangePassword", GoMethod: "ResetMustChangePassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetRsaPublicKey", GoMethod: "ResetRsaPublicKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetRsaPublicKey2", GoMethod: "ResetRsaPublicKey2"},
			_jsii_.MemberProperty{JsiiProperty: "rsaPublicKey", GoGetter: "RsaPublicKey"},
			_jsii_.MemberProperty{JsiiProperty: "rsaPublicKey2", GoGetter: "RsaPublicKey2"},
			_jsii_.MemberProperty{JsiiProperty: "rsaPublicKey2Input", GoGetter: "RsaPublicKey2Input"},
			_jsii_.MemberProperty{JsiiProperty: "rsaPublicKeyInput", GoGetter: "RsaPublicKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_User{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.user.UserConfig",
		reflect.TypeOf((*UserConfig)(nil)).Elem(),
	)
}
