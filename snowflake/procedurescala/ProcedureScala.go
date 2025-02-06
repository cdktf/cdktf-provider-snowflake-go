// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedurescala

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/procedurescala/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/procedure_scala snowflake_procedure_scala}.
type ProcedureScala interface {
	cdktf.TerraformResource
	Arguments() ProcedureScalaArgumentsList
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
	Imports() ProcedureScalaImportsList
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
	Parameters() ProcedureScalaParametersList
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
	Secrets() ProcedureScalaSecretsList
	SecretsInput() interface{}
	ShowOutput() ProcedureScalaShowOutputList
	SnowparkPackage() *string
	SetSnowparkPackage(val *string)
	SnowparkPackageInput() *string
	TargetPath() ProcedureScalaTargetPathOutputReference
	TargetPathInput() *ProcedureScalaTargetPath
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutTargetPath(value *ProcedureScalaTargetPath)
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

// The jsii proxy struct for ProcedureScala
type jsiiProxy_ProcedureScala struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ProcedureScala) Arguments() ProcedureScalaArgumentsList {
	var returns ProcedureScalaArgumentsList
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) EnableConsoleOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) EnableConsoleOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ExecuteAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executeAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ExecuteAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executeAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ExternalAccessIntegrations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalAccessIntegrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ExternalAccessIntegrationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalAccessIntegrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Imports() ProcedureScalaImportsList {
	var returns ProcedureScalaImportsList
	_jsii_.Get(
		j,
		"imports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ImportsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) IsSecure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) IsSecureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) MetricLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) MetricLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) NullInputBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) NullInputBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Packages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) PackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Parameters() ProcedureScalaParametersList {
	var returns ProcedureScalaParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ProcedureDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procedureDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ProcedureDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procedureDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ProcedureLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procedureLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ReturnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ReturnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) Secrets() ProcedureScalaSecretsList {
	var returns ProcedureScalaSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) ShowOutput() ProcedureScalaShowOutputList {
	var returns ProcedureScalaShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) SnowparkPackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowparkPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) SnowparkPackageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowparkPackageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) TargetPath() ProcedureScalaTargetPathOutputReference {
	var returns ProcedureScalaTargetPathOutputReference
	_jsii_.Get(
		j,
		"targetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) TargetPathInput() *ProcedureScalaTargetPath {
	var returns *ProcedureScalaTargetPath
	_jsii_.Get(
		j,
		"targetPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProcedureScala) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/procedure_scala snowflake_procedure_scala} Resource.
func NewProcedureScala(scope constructs.Construct, id *string, config *ProcedureScalaConfig) ProcedureScala {
	_init_.Initialize()

	if err := validateNewProcedureScalaParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProcedureScala{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.procedureScala.ProcedureScala",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/procedure_scala snowflake_procedure_scala} Resource.
func NewProcedureScala_Override(p ProcedureScala, scope constructs.Construct, id *string, config *ProcedureScalaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.procedureScala.ProcedureScala",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProcedureScala)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetEnableConsoleOutput(val interface{}) {
	if err := j.validateSetEnableConsoleOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsoleOutput",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetExecuteAs(val *string) {
	if err := j.validateSetExecuteAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeAs",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetExternalAccessIntegrations(val *[]*string) {
	if err := j.validateSetExternalAccessIntegrationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAccessIntegrations",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetIsSecure(val *string) {
	if err := j.validateSetIsSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecure",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetMetricLevel(val *string) {
	if err := j.validateSetMetricLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricLevel",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetNullInputBehavior(val *string) {
	if err := j.validateSetNullInputBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullInputBehavior",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetPackages(val *[]*string) {
	if err := j.validateSetPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packages",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetProcedureDefinition(val *string) {
	if err := j.validateSetProcedureDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"procedureDefinition",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetReturnType(val *string) {
	if err := j.validateSetReturnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnType",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetSnowparkPackage(val *string) {
	if err := j.validateSetSnowparkPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowparkPackage",
		val,
	)
}

func (j *jsiiProxy_ProcedureScala)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

// Generates CDKTF code for importing a ProcedureScala resource upon running "cdktf plan <stack-name>".
func ProcedureScala_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProcedureScala_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.procedureScala.ProcedureScala",
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
func ProcedureScala_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProcedureScala_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.procedureScala.ProcedureScala",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProcedureScala_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProcedureScala_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.procedureScala.ProcedureScala",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProcedureScala_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProcedureScala_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.procedureScala.ProcedureScala",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ProcedureScala_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.procedureScala.ProcedureScala",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ProcedureScala) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_ProcedureScala) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProcedureScala) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_ProcedureScala) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProcedureScala) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_ProcedureScala) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_ProcedureScala) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_ProcedureScala) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_ProcedureScala) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_ProcedureScala) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_ProcedureScala) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_ProcedureScala) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureScala) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_ProcedureScala) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_ProcedureScala) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProcedureScala) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_ProcedureScala) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_ProcedureScala) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProcedureScala) PutArguments(value interface{}) {
	if err := p.validatePutArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putArguments",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureScala) PutImports(value interface{}) {
	if err := p.validatePutImportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putImports",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureScala) PutSecrets(value interface{}) {
	if err := p.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecrets",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureScala) PutTargetPath(value *ProcedureScalaTargetPath) {
	if err := p.validatePutTargetPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTargetPath",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_ProcedureScala) ResetArguments() {
	_jsii_.InvokeVoid(
		p,
		"resetArguments",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetComment() {
	_jsii_.InvokeVoid(
		p,
		"resetComment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetEnableConsoleOutput() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableConsoleOutput",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetExecuteAs() {
	_jsii_.InvokeVoid(
		p,
		"resetExecuteAs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetExternalAccessIntegrations() {
	_jsii_.InvokeVoid(
		p,
		"resetExternalAccessIntegrations",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetImports() {
	_jsii_.InvokeVoid(
		p,
		"resetImports",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetIsSecure() {
	_jsii_.InvokeVoid(
		p,
		"resetIsSecure",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetLogLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetMetricLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetNullInputBehavior() {
	_jsii_.InvokeVoid(
		p,
		"resetNullInputBehavior",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetPackages() {
	_jsii_.InvokeVoid(
		p,
		"resetPackages",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetProcedureDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetProcedureDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetSecrets() {
	_jsii_.InvokeVoid(
		p,
		"resetSecrets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetTargetPath() {
	_jsii_.InvokeVoid(
		p,
		"resetTargetPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProcedureScala) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureScala) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureScala) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureScala) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureScala) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProcedureScala) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

