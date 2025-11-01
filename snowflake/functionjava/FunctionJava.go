// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionjava

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/functionjava/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/function_java snowflake_function_java}.
type FunctionJava interface {
	cdktf.TerraformResource
	Arguments() FunctionJavaArgumentsList
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
	Imports() FunctionJavaImportsList
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
	Parameters() FunctionJavaParametersList
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
	Secrets() FunctionJavaSecretsList
	SecretsInput() interface{}
	ShowOutput() FunctionJavaShowOutputList
	TargetPath() FunctionJavaTargetPathOutputReference
	TargetPathInput() *FunctionJavaTargetPath
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FunctionJavaTimeoutsOutputReference
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
	PutTargetPath(value *FunctionJavaTargetPath)
	PutTimeouts(value *FunctionJavaTimeouts)
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
	ResetRuntimeVersion()
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

// The jsii proxy struct for FunctionJava
type jsiiProxy_FunctionJava struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FunctionJava) Arguments() FunctionJavaArgumentsList {
	var returns FunctionJavaArgumentsList
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) EnableConsoleOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) EnableConsoleOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ExternalAccessIntegrations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalAccessIntegrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ExternalAccessIntegrationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalAccessIntegrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) FunctionDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) FunctionDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) FunctionLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Imports() FunctionJavaImportsList {
	var returns FunctionJavaImportsList
	_jsii_.Get(
		j,
		"imports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ImportsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) IsSecure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) IsSecureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) MetricLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) MetricLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) NullInputBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) NullInputBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Packages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) PackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Parameters() FunctionJavaParametersList {
	var returns FunctionJavaParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ReturnResultsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnResultsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ReturnResultsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnResultsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ReturnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ReturnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Secrets() FunctionJavaSecretsList {
	var returns FunctionJavaSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) ShowOutput() FunctionJavaShowOutputList {
	var returns FunctionJavaShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) TargetPath() FunctionJavaTargetPathOutputReference {
	var returns FunctionJavaTargetPathOutputReference
	_jsii_.Get(
		j,
		"targetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) TargetPathInput() *FunctionJavaTargetPath {
	var returns *FunctionJavaTargetPath
	_jsii_.Get(
		j,
		"targetPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) Timeouts() FunctionJavaTimeoutsOutputReference {
	var returns FunctionJavaTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionJava) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/function_java snowflake_function_java} Resource.
func NewFunctionJava(scope constructs.Construct, id *string, config *FunctionJavaConfig) FunctionJava {
	_init_.Initialize()

	if err := validateNewFunctionJavaParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionJava{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.functionJava.FunctionJava",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/function_java snowflake_function_java} Resource.
func NewFunctionJava_Override(f FunctionJava, scope constructs.Construct, id *string, config *FunctionJavaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.functionJava.FunctionJava",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FunctionJava)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetEnableConsoleOutput(val interface{}) {
	if err := j.validateSetEnableConsoleOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsoleOutput",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetExternalAccessIntegrations(val *[]*string) {
	if err := j.validateSetExternalAccessIntegrationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAccessIntegrations",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetFunctionDefinition(val *string) {
	if err := j.validateSetFunctionDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionDefinition",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetIsSecure(val *string) {
	if err := j.validateSetIsSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecure",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetMetricLevel(val *string) {
	if err := j.validateSetMetricLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricLevel",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetNullInputBehavior(val *string) {
	if err := j.validateSetNullInputBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullInputBehavior",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetPackages(val *[]*string) {
	if err := j.validateSetPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packages",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetReturnResultsBehavior(val *string) {
	if err := j.validateSetReturnResultsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnResultsBehavior",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetReturnType(val *string) {
	if err := j.validateSetReturnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnType",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_FunctionJava)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

// Generates CDKTF code for importing a FunctionJava resource upon running "cdktf plan <stack-name>".
func FunctionJava_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFunctionJava_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionJava.FunctionJava",
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
func FunctionJava_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionJava_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionJava.FunctionJava",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionJava_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionJava_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionJava.FunctionJava",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FunctionJava_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionJava_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.functionJava.FunctionJava",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FunctionJava_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.functionJava.FunctionJava",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FunctionJava) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FunctionJava) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FunctionJava) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FunctionJava) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionJava) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FunctionJava) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FunctionJava) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FunctionJava) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FunctionJava) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FunctionJava) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FunctionJava) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FunctionJava) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionJava) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FunctionJava) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionJava) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FunctionJava) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FunctionJava) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FunctionJava) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FunctionJava) PutArguments(value interface{}) {
	if err := f.validatePutArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putArguments",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionJava) PutImports(value interface{}) {
	if err := f.validatePutImportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putImports",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionJava) PutSecrets(value interface{}) {
	if err := f.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSecrets",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionJava) PutTargetPath(value *FunctionJavaTargetPath) {
	if err := f.validatePutTargetPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTargetPath",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionJava) PutTimeouts(value *FunctionJavaTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionJava) ResetArguments() {
	_jsii_.InvokeVoid(
		f,
		"resetArguments",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetComment() {
	_jsii_.InvokeVoid(
		f,
		"resetComment",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetEnableConsoleOutput() {
	_jsii_.InvokeVoid(
		f,
		"resetEnableConsoleOutput",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetExternalAccessIntegrations() {
	_jsii_.InvokeVoid(
		f,
		"resetExternalAccessIntegrations",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetFunctionDefinition() {
	_jsii_.InvokeVoid(
		f,
		"resetFunctionDefinition",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetImports() {
	_jsii_.InvokeVoid(
		f,
		"resetImports",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetIsSecure() {
	_jsii_.InvokeVoid(
		f,
		"resetIsSecure",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetLogLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetMetricLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetMetricLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetNullInputBehavior() {
	_jsii_.InvokeVoid(
		f,
		"resetNullInputBehavior",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetPackages() {
	_jsii_.InvokeVoid(
		f,
		"resetPackages",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetReturnResultsBehavior() {
	_jsii_.InvokeVoid(
		f,
		"resetReturnResultsBehavior",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetSecrets() {
	_jsii_.InvokeVoid(
		f,
		"resetSecrets",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetTargetPath() {
	_jsii_.InvokeVoid(
		f,
		"resetTargetPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionJava) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionJava) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionJava) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionJava) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionJava) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionJava) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

