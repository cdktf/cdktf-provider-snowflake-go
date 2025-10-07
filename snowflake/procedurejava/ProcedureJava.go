// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedurejava

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/procedurejava/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/procedure_java snowflake_procedure_java}.
type ProcedureJava interface {
	cdktf.TerraformResource
	Arguments() ProcedureJavaArgumentsList
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
	ExecuteAs() *string
	SetExecuteAs(val *string)
	ExecuteAsInput() *string
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
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Imports() ProcedureJavaImportsList
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
	Parameters() ProcedureJavaParametersList
	ProcedureDefinition() *string
	SetProcedureDefinition(val *string)
	ProcedureDefinitionInput() *string
	ProcedureLanguage() *string
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
	ReturnType() *string
	SetReturnType(val *string)
	ReturnTypeInput() *string
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Secrets() ProcedureJavaSecretsList
	SecretsInput() interface{}
	ShowOutput() ProcedureJavaShowOutputList
	SnowparkPackage() *string
	SetSnowparkPackage(val *string)
	SnowparkPackageInput() *string
	TargetPath() ProcedureJavaTargetPathOutputReference
	TargetPathInput() *ProcedureJavaTargetPath
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ProcedureJavaTimeoutsOutputReference
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
	PutTargetPath(value *ProcedureJavaTargetPath)
	PutTimeouts(value *ProcedureJavaTimeouts)
	ResetArguments()
	ResetComment()
	ResetEnableConsoleOutput()
	ResetExecuteAs()
	ResetExternalAccessIntegrations()
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
	ResetProcedureDefinition()
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

// The jsii proxy struct for ProcedureJava
type jsiiProxy_ProcedureJava struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProcedureJava) Arguments() ProcedureJavaArgumentsList {
	var returns ProcedureJavaArgumentsList
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) EnableConsoleOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) EnableConsoleOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ExecuteAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executeAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ExecuteAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executeAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ExternalAccessIntegrations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalAccessIntegrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ExternalAccessIntegrationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalAccessIntegrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Imports() ProcedureJavaImportsList {
	var returns ProcedureJavaImportsList
	_jsii_.Get(
		j,
		"imports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ImportsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) IsSecure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) IsSecureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) MetricLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) MetricLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) NullInputBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) NullInputBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Packages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) PackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Parameters() ProcedureJavaParametersList {
	var returns ProcedureJavaParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ProcedureDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procedureDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ProcedureDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procedureDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ProcedureLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procedureLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ReturnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ReturnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Secrets() ProcedureJavaSecretsList {
	var returns ProcedureJavaSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) ShowOutput() ProcedureJavaShowOutputList {
	var returns ProcedureJavaShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) SnowparkPackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowparkPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) SnowparkPackageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowparkPackageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) TargetPath() ProcedureJavaTargetPathOutputReference {
	var returns ProcedureJavaTargetPathOutputReference
	_jsii_.Get(
		j,
		"targetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) TargetPathInput() *ProcedureJavaTargetPath {
	var returns *ProcedureJavaTargetPath
	_jsii_.Get(
		j,
		"targetPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) Timeouts() ProcedureJavaTimeoutsOutputReference {
	var returns ProcedureJavaTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureJava) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/procedure_java snowflake_procedure_java} Resource.
func NewProcedureJava(scope constructs.Construct, id *string, config *ProcedureJavaConfig) ProcedureJava {
	_init_.Initialize()

	if err := validateNewProcedureJavaParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProcedureJava{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.procedureJava.ProcedureJava",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/procedure_java snowflake_procedure_java} Resource.
func NewProcedureJava_Override(p ProcedureJava, scope constructs.Construct, id *string, config *ProcedureJavaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.procedureJava.ProcedureJava",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProcedureJava)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetEnableConsoleOutput(val interface{}) {
	if err := j.validateSetEnableConsoleOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsoleOutput",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetExecuteAs(val *string) {
	if err := j.validateSetExecuteAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeAs",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetExternalAccessIntegrations(val *[]*string) {
	if err := j.validateSetExternalAccessIntegrationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAccessIntegrations",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetIsSecure(val *string) {
	if err := j.validateSetIsSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecure",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetMetricLevel(val *string) {
	if err := j.validateSetMetricLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricLevel",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetNullInputBehavior(val *string) {
	if err := j.validateSetNullInputBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullInputBehavior",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetPackages(val *[]*string) {
	if err := j.validateSetPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packages",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetProcedureDefinition(val *string) {
	if err := j.validateSetProcedureDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"procedureDefinition",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetReturnType(val *string) {
	if err := j.validateSetReturnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnType",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetSnowparkPackage(val *string) {
	if err := j.validateSetSnowparkPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowparkPackage",
		val,
	)
}

func (j *jsiiProxy_ProcedureJava)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

// Generates CDKTF code for importing a ProcedureJava resource upon running "cdktf plan <stack-name>".
func ProcedureJava_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProcedureJava_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.procedureJava.ProcedureJava",
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
func ProcedureJava_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProcedureJava_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.procedureJava.ProcedureJava",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProcedureJava_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProcedureJava_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.procedureJava.ProcedureJava",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProcedureJava_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProcedureJava_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.procedureJava.ProcedureJava",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProcedureJava_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.procedureJava.ProcedureJava",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProcedureJava) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_ProcedureJava) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProcedureJava) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_ProcedureJava) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProcedureJava) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_ProcedureJava) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProcedureJava) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProcedureJava) PutArguments(value interface{}) {
	if err := p.validatePutArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putArguments",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureJava) PutImports(value interface{}) {
	if err := p.validatePutImportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putImports",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureJava) PutSecrets(value interface{}) {
	if err := p.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecrets",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureJava) PutTargetPath(value *ProcedureJavaTargetPath) {
	if err := p.validatePutTargetPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTargetPath",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureJava) PutTimeouts(value *ProcedureJavaTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureJava) ResetArguments() {
	_jsii_.InvokeVoid(
		p,
		"resetArguments",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetComment() {
	_jsii_.InvokeVoid(
		p,
		"resetComment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetEnableConsoleOutput() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableConsoleOutput",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetExecuteAs() {
	_jsii_.InvokeVoid(
		p,
		"resetExecuteAs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetExternalAccessIntegrations() {
	_jsii_.InvokeVoid(
		p,
		"resetExternalAccessIntegrations",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetImports() {
	_jsii_.InvokeVoid(
		p,
		"resetImports",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetIsSecure() {
	_jsii_.InvokeVoid(
		p,
		"resetIsSecure",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetLogLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetMetricLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetNullInputBehavior() {
	_jsii_.InvokeVoid(
		p,
		"resetNullInputBehavior",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetPackages() {
	_jsii_.InvokeVoid(
		p,
		"resetPackages",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetProcedureDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetProcedureDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetSecrets() {
	_jsii_.InvokeVoid(
		p,
		"resetSecrets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetTargetPath() {
	_jsii_.InvokeVoid(
		p,
		"resetTargetPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureJava) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureJava) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

