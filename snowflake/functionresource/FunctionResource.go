// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v9/functionresource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/function snowflake_function}.
type FunctionResource interface {
	cdktf.TerraformResource
	Arguments() FunctionResourceArgumentsList
	ArgumentsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Imports() *[]*string
	SetImports(val *[]*string)
	ImportsInput() *[]*string
	IsSecure() interface{}
	SetIsSecure(val interface{})
	IsSecureInput() interface{}
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NullInputBehavior() *string
	SetNullInputBehavior(val *string)
	NullInputBehaviorInput() *string
	Packages() *[]*string
	SetPackages(val *[]*string)
	PackagesInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReturnBehavior() *string
	SetReturnBehavior(val *string)
	ReturnBehaviorInput() *string
	ReturnType() *string
	SetReturnType(val *string)
	ReturnTypeInput() *string
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Statement() *string
	SetStatement(val *string)
	StatementInput() *string
	TargetPath() *string
	SetTargetPath(val *string)
	TargetPathInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutArguments(value interface{})
	ResetArguments()
	ResetComment()
	ResetHandler()
	ResetId()
	ResetImports()
	ResetIsSecure()
	ResetLanguage()
	ResetNullInputBehavior()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackages()
	ResetReturnBehavior()
	ResetRuntimeVersion()
	ResetTargetPath()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FunctionResource
type jsiiProxy_FunctionResource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FunctionResource) Arguments() FunctionResourceArgumentsList {
	var returns FunctionResourceArgumentsList
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) ArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Imports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) ImportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) IsSecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) IsSecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) NullInputBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) NullInputBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Packages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) PackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) ReturnBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) ReturnBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) ReturnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) ReturnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) Statement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) StatementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) TargetPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) TargetPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/function snowflake_function} Resource.
func NewFunctionResource(scope constructs.Construct, id *string, config *FunctionResourceConfig) FunctionResource {
	_init_.Initialize()

	if err := validateNewFunctionResourceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionResource{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.functionResource.FunctionResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/function snowflake_function} Resource.
func NewFunctionResource_Override(f FunctionResource, scope constructs.Construct, id *string, config *FunctionResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.functionResource.FunctionResource",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FunctionResource)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetImports(val *[]*string) {
	if err := j.validateSetImportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imports",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetIsSecure(val interface{}) {
	if err := j.validateSetIsSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecure",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetLanguage(val *string) {
	if err := j.validateSetLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetNullInputBehavior(val *string) {
	if err := j.validateSetNullInputBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullInputBehavior",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetPackages(val *[]*string) {
	if err := j.validateSetPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packages",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetReturnBehavior(val *string) {
	if err := j.validateSetReturnBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnBehavior",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetReturnType(val *string) {
	if err := j.validateSetReturnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnType",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetStatement(val *string) {
	if err := j.validateSetStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statement",
		val,
	)
}

func (j *jsiiProxy_FunctionResource)SetTargetPath(val *string) {
	if err := j.validateSetTargetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPath",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func FunctionResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionResource.FunctionResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionResource_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionResource_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionResource.FunctionResource",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionResource_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionResource_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionResource.FunctionResource",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FunctionResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.functionResource.FunctionResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FunctionResource) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FunctionResource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FunctionResource) PutArguments(value interface{}) {
	if err := f.validatePutArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putArguments",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionResource) ResetArguments() {
	_jsii_.InvokeVoid(
		f,
		"resetArguments",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetComment() {
	_jsii_.InvokeVoid(
		f,
		"resetComment",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetHandler() {
	_jsii_.InvokeVoid(
		f,
		"resetHandler",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetImports() {
	_jsii_.InvokeVoid(
		f,
		"resetImports",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetIsSecure() {
	_jsii_.InvokeVoid(
		f,
		"resetIsSecure",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetLanguage() {
	_jsii_.InvokeVoid(
		f,
		"resetLanguage",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetNullInputBehavior() {
	_jsii_.InvokeVoid(
		f,
		"resetNullInputBehavior",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetPackages() {
	_jsii_.InvokeVoid(
		f,
		"resetPackages",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetReturnBehavior() {
	_jsii_.InvokeVoid(
		f,
		"resetReturnBehavior",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) ResetTargetPath() {
	_jsii_.InvokeVoid(
		f,
		"resetTargetPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

