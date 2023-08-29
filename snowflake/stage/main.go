// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stage

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.stage.Stage",
		reflect.TypeOf((*Stage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "awsExternalId", GoGetter: "AwsExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "awsExternalIdInput", GoGetter: "AwsExternalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "copyOptions", GoGetter: "CopyOptions"},
			_jsii_.MemberProperty{JsiiProperty: "copyOptionsInput", GoGetter: "CopyOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsInput", GoGetter: "CredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directory", GoGetter: "Directory"},
			_jsii_.MemberProperty{JsiiProperty: "directoryInput", GoGetter: "DirectoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionInput", GoGetter: "EncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileFormat", GoGetter: "FileFormat"},
			_jsii_.MemberProperty{JsiiProperty: "fileFormatInput", GoGetter: "FileFormatInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTag", GoMethod: "PutTag"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsExternalId", GoMethod: "ResetAwsExternalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyOptions", GoMethod: "ResetCopyOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCredentials", GoMethod: "ResetCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirectory", GoMethod: "ResetDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryption", GoMethod: "ResetEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileFormat", GoMethod: "ResetFileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnowflakeIamUser", GoMethod: "ResetSnowflakeIamUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageIntegration", GoMethod: "ResetStorageIntegration"},
			_jsii_.MemberMethod{JsiiMethod: "resetTag", GoMethod: "ResetTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrl", GoMethod: "ResetUrl"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeIamUser", GoGetter: "SnowflakeIamUser"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeIamUserInput", GoGetter: "SnowflakeIamUserInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageIntegration", GoGetter: "StorageIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "storageIntegrationInput", GoGetter: "StorageIntegrationInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberProperty{JsiiProperty: "tagInput", GoGetter: "TagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
		},
		func() interface{} {
			j := jsiiProxy_Stage{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.stage.StageConfig",
		reflect.TypeOf((*StageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-snowflake.stage.StageTag",
		reflect.TypeOf((*StageTag)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.stage.StageTagList",
		reflect.TypeOf((*StageTagList)(nil)).Elem(),
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
			j := jsiiProxy_StageTagList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-snowflake.stage.StageTagOutputReference",
		reflect.TypeOf((*StageTagOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_StageTagOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
