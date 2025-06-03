// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionscala

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/functionscala/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/function_scala snowflake_function_scala}.
type FunctionScala interface {
	cdktf.TerraformResource
	Arguments() FunctionScalaArgumentsList
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
	EnableConsoleOutput() interface{}
	SetEnableConsoleOutput(val interface{})
	EnableConsoleOutputInput() interface{}
	ExternalAccessIntegrations() *[]*string
	SetExternalAccessIntegrations(val *[]*string)
	ExternalAccessIntegrationsInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullyQualifiedName() *string
	FunctionDefinition() *string
	SetFunctionDefinition(val *string)
	FunctionDefinitionInput() *string
	FunctionLanguage() *string
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Imports() FunctionScalaImportsList
	ImportsInput() interface{}
	IsSecure() *string
	SetIsSecure(val *string)
	IsSecureInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	MetricLevel() *string
	SetMetricLevel(val *string)
	MetricLevelInput() *string
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
	Parameters() FunctionScalaParametersList
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
	ReturnResultsBehavior() *string
	SetReturnResultsBehavior(val *string)
	ReturnResultsBehaviorInput() *string
	ReturnType() *string
	SetReturnType(val *string)
	ReturnTypeInput() *string
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Secrets() FunctionScalaSecretsList
	SecretsInput() interface{}
	ShowOutput() FunctionScalaShowOutputList
	TargetPath() FunctionScalaTargetPathOutputReference
	TargetPathInput() *FunctionScalaTargetPath
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FunctionScalaTimeoutsOutputReference
	TimeoutsInput() interface{}
	TraceLevel() *string
	SetTraceLevel(val *string)
	TraceLevelInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutArguments(value interface{})
	PutImports(value interface{})
	PutSecrets(value interface{})
	PutTargetPath(value *FunctionScalaTargetPath)
	PutTimeouts(value *FunctionScalaTimeouts)
	ResetArguments()
	ResetComment()
	ResetEnableConsoleOutput()
	ResetExternalAccessIntegrations()
	ResetFunctionDefinition()
	ResetId()
	ResetImports()
	ResetIsSecure()
	ResetLogLevel()
	ResetMetricLevel()
	ResetNullInputBehavior()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackages()
	ResetReturnResultsBehavior()
	ResetSecrets()
	ResetTargetPath()
	ResetTimeouts()
	ResetTraceLevel()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FunctionScala
type jsiiProxy_FunctionScala struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FunctionScala) Arguments() FunctionScalaArgumentsList {
	var returns FunctionScalaArgumentsList
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) EnableConsoleOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) EnableConsoleOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ExternalAccessIntegrations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalAccessIntegrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ExternalAccessIntegrationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalAccessIntegrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) FunctionDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) FunctionDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) FunctionLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Imports() FunctionScalaImportsList {
	var returns FunctionScalaImportsList
	_jsii_.Get(
		j,
		"imports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ImportsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) IsSecure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) IsSecureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) MetricLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) MetricLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) NullInputBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) NullInputBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Packages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) PackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Parameters() FunctionScalaParametersList {
	var returns FunctionScalaParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ReturnResultsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnResultsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ReturnResultsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnResultsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ReturnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ReturnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Secrets() FunctionScalaSecretsList {
	var returns FunctionScalaSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) ShowOutput() FunctionScalaShowOutputList {
	var returns FunctionScalaShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) TargetPath() FunctionScalaTargetPathOutputReference {
	var returns FunctionScalaTargetPathOutputReference
	_jsii_.Get(
		j,
		"targetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) TargetPathInput() *FunctionScalaTargetPath {
	var returns *FunctionScalaTargetPath
	_jsii_.Get(
		j,
		"targetPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) Timeouts() FunctionScalaTimeoutsOutputReference {
	var returns FunctionScalaTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionScala) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/function_scala snowflake_function_scala} Resource.
func NewFunctionScala(scope constructs.Construct, id *string, config *FunctionScalaConfig) FunctionScala {
	_init_.Initialize()

	if err := validateNewFunctionScalaParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionScala{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.functionScala.FunctionScala",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/function_scala snowflake_function_scala} Resource.
func NewFunctionScala_Override(f FunctionScala, scope constructs.Construct, id *string, config *FunctionScalaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.functionScala.FunctionScala",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FunctionScala)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetEnableConsoleOutput(val interface{}) {
	if err := j.validateSetEnableConsoleOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsoleOutput",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetExternalAccessIntegrations(val *[]*string) {
	if err := j.validateSetExternalAccessIntegrationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAccessIntegrations",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetFunctionDefinition(val *string) {
	if err := j.validateSetFunctionDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionDefinition",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetIsSecure(val *string) {
	if err := j.validateSetIsSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecure",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetMetricLevel(val *string) {
	if err := j.validateSetMetricLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricLevel",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetNullInputBehavior(val *string) {
	if err := j.validateSetNullInputBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullInputBehavior",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetPackages(val *[]*string) {
	if err := j.validateSetPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packages",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetReturnResultsBehavior(val *string) {
	if err := j.validateSetReturnResultsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnResultsBehavior",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetReturnType(val *string) {
	if err := j.validateSetReturnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnType",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_FunctionScala)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

// Generates CDKTF code for importing a FunctionScala resource upon running "cdktf plan <stack-name>".
func FunctionScala_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFunctionScala_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionScala.FunctionScala",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func FunctionScala_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionScala_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionScala.FunctionScala",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionScala_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionScala_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionScala.FunctionScala",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionScala_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionScala_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionScala.FunctionScala",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FunctionScala_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.functionScala.FunctionScala",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FunctionScala) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FunctionScala) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FunctionScala) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FunctionScala) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionScala) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FunctionScala) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FunctionScala) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FunctionScala) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FunctionScala) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FunctionScala) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FunctionScala) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FunctionScala) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionScala) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FunctionScala) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionScala) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FunctionScala) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FunctionScala) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FunctionScala) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FunctionScala) PutArguments(value interface{}) {
	if err := f.validatePutArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putArguments",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionScala) PutImports(value interface{}) {
	if err := f.validatePutImportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putImports",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionScala) PutSecrets(value interface{}) {
	if err := f.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSecrets",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionScala) PutTargetPath(value *FunctionScalaTargetPath) {
	if err := f.validatePutTargetPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTargetPath",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionScala) PutTimeouts(value *FunctionScalaTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionScala) ResetArguments() {
	_jsii_.InvokeVoid(
		f,
		"resetArguments",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetComment() {
	_jsii_.InvokeVoid(
		f,
		"resetComment",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetEnableConsoleOutput() {
	_jsii_.InvokeVoid(
		f,
		"resetEnableConsoleOutput",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetExternalAccessIntegrations() {
	_jsii_.InvokeVoid(
		f,
		"resetExternalAccessIntegrations",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetFunctionDefinition() {
	_jsii_.InvokeVoid(
		f,
		"resetFunctionDefinition",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetImports() {
	_jsii_.InvokeVoid(
		f,
		"resetImports",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetIsSecure() {
	_jsii_.InvokeVoid(
		f,
		"resetIsSecure",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetLogLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetMetricLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetMetricLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetNullInputBehavior() {
	_jsii_.InvokeVoid(
		f,
		"resetNullInputBehavior",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetPackages() {
	_jsii_.InvokeVoid(
		f,
		"resetPackages",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetReturnResultsBehavior() {
	_jsii_.InvokeVoid(
		f,
		"resetReturnResultsBehavior",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetSecrets() {
	_jsii_.InvokeVoid(
		f,
		"resetSecrets",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetTargetPath() {
	_jsii_.InvokeVoid(
		f,
		"resetTargetPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionScala) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionScala) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionScala) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionScala) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionScala) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionScala) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

