package externalfunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v6/externalfunction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/external_function snowflake_external_function}.
type ExternalFunction interface {
	cdktf.TerraformResource
	ApiIntegration() *string
	SetApiIntegration(val *string)
	ApiIntegrationInput() *string
	Arg() ExternalFunctionArgList
	ArgInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContextHeaders() *[]*string
	SetContextHeaders(val *[]*string)
	ContextHeadersInput() *[]*string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedOn() *string
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
	Header() ExternalFunctionHeaderList
	HeaderInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxBatchRows() *float64
	SetMaxBatchRows(val *float64)
	MaxBatchRowsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NullInputBehavior() *string
	SetNullInputBehavior(val *string)
	NullInputBehaviorInput() *string
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
	RequestTranslator() *string
	SetRequestTranslator(val *string)
	RequestTranslatorInput() *string
	ResponseTranslator() *string
	SetResponseTranslator(val *string)
	ResponseTranslatorInput() *string
	ReturnBehavior() *string
	SetReturnBehavior(val *string)
	ReturnBehaviorInput() *string
	ReturnNullAllowed() interface{}
	SetReturnNullAllowed(val interface{})
	ReturnNullAllowedInput() interface{}
	ReturnType() *string
	SetReturnType(val *string)
	ReturnTypeInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UrlOfProxyAndResource() *string
	SetUrlOfProxyAndResource(val *string)
	UrlOfProxyAndResourceInput() *string
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
	PutArg(value interface{})
	PutHeader(value interface{})
	ResetArg()
	ResetComment()
	ResetCompression()
	ResetContextHeaders()
	ResetHeader()
	ResetId()
	ResetMaxBatchRows()
	ResetNullInputBehavior()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestTranslator()
	ResetResponseTranslator()
	ResetReturnNullAllowed()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ExternalFunction
type jsiiProxy_ExternalFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ExternalFunction) ApiIntegration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ApiIntegrationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Arg() ExternalFunctionArgList {
	var returns ExternalFunctionArgList
	_jsii_.Get(
		j,
		"arg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ArgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ContextHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contextHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ContextHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contextHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Header() ExternalFunctionHeaderList {
	var returns ExternalFunctionHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) MaxBatchRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) MaxBatchRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) NullInputBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) NullInputBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullInputBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) RequestTranslator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTranslator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) RequestTranslatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTranslatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ResponseTranslator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTranslator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ResponseTranslatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTranslatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ReturnBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ReturnBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ReturnNullAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnNullAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ReturnNullAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnNullAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ReturnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) ReturnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) UrlOfProxyAndResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlOfProxyAndResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalFunction) UrlOfProxyAndResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlOfProxyAndResourceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/external_function snowflake_external_function} Resource.
func NewExternalFunction(scope constructs.Construct, id *string, config *ExternalFunctionConfig) ExternalFunction {
	_init_.Initialize()

	if err := validateNewExternalFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalFunction{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalFunction.ExternalFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/external_function snowflake_external_function} Resource.
func NewExternalFunction_Override(e ExternalFunction, scope constructs.Construct, id *string, config *ExternalFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalFunction.ExternalFunction",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExternalFunction)SetApiIntegration(val *string) {
	if err := j.validateSetApiIntegrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiIntegration",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetContextHeaders(val *[]*string) {
	if err := j.validateSetContextHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextHeaders",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetMaxBatchRows(val *float64) {
	if err := j.validateSetMaxBatchRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchRows",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetNullInputBehavior(val *string) {
	if err := j.validateSetNullInputBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullInputBehavior",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetRequestTranslator(val *string) {
	if err := j.validateSetRequestTranslatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTranslator",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetResponseTranslator(val *string) {
	if err := j.validateSetResponseTranslatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTranslator",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetReturnBehavior(val *string) {
	if err := j.validateSetReturnBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnBehavior",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetReturnNullAllowed(val interface{}) {
	if err := j.validateSetReturnNullAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnNullAllowed",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetReturnType(val *string) {
	if err := j.validateSetReturnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnType",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_ExternalFunction)SetUrlOfProxyAndResource(val *string) {
	if err := j.validateSetUrlOfProxyAndResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlOfProxyAndResource",
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
func ExternalFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalFunction.ExternalFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalFunction.ExternalFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalFunction.ExternalFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExternalFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.externalFunction.ExternalFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExternalFunction) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExternalFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExternalFunction) PutArg(value interface{}) {
	if err := e.validatePutArgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putArg",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalFunction) PutHeader(value interface{}) {
	if err := e.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHeader",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalFunction) ResetArg() {
	_jsii_.InvokeVoid(
		e,
		"resetArg",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetComment() {
	_jsii_.InvokeVoid(
		e,
		"resetComment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetCompression() {
	_jsii_.InvokeVoid(
		e,
		"resetCompression",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetContextHeaders() {
	_jsii_.InvokeVoid(
		e,
		"resetContextHeaders",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetHeader() {
	_jsii_.InvokeVoid(
		e,
		"resetHeader",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetMaxBatchRows() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxBatchRows",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetNullInputBehavior() {
	_jsii_.InvokeVoid(
		e,
		"resetNullInputBehavior",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetRequestTranslator() {
	_jsii_.InvokeVoid(
		e,
		"resetRequestTranslator",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetResponseTranslator() {
	_jsii_.InvokeVoid(
		e,
		"resetResponseTranslator",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) ResetReturnNullAllowed() {
	_jsii_.InvokeVoid(
		e,
		"resetReturnNullAllowed",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

