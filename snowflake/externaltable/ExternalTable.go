package externaltable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v6/externaltable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/external_table snowflake_external_table}.
type ExternalTable interface {
	cdktf.TerraformResource
	AutoRefresh() interface{}
	SetAutoRefresh(val interface{})
	AutoRefreshInput() interface{}
	AwsSnsTopic() *string
	SetAwsSnsTopic(val *string)
	AwsSnsTopicInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Column() ExternalTableColumnList
	ColumnInput() interface{}
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyGrants() interface{}
	SetCopyGrants(val interface{})
	CopyGrantsInput() interface{}
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
	FileFormat() *string
	SetFileFormat(val *string)
	FileFormatInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	PartitionBy() *[]*string
	SetPartitionBy(val *[]*string)
	PartitionByInput() *[]*string
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
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
	RefreshOnCreate() interface{}
	SetRefreshOnCreate(val interface{})
	RefreshOnCreateInput() interface{}
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Tag() ExternalTableTagList
	TagInput() interface{}
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
	PutColumn(value interface{})
	PutTag(value interface{})
	ResetAutoRefresh()
	ResetAwsSnsTopic()
	ResetComment()
	ResetCopyGrants()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionBy()
	ResetPattern()
	ResetRefreshOnCreate()
	ResetTag()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ExternalTable
type jsiiProxy_ExternalTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ExternalTable) AutoRefresh() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) AutoRefreshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRefreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) AwsSnsTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSnsTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) AwsSnsTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSnsTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Column() ExternalTableColumnList {
	var returns ExternalTableColumnList
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) ColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) CopyGrants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) CopyGrantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyGrantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) FileFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) FileFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) PartitionBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"partitionBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) PartitionByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"partitionByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) RefreshOnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshOnCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) RefreshOnCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshOnCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) Tag() ExternalTableTagList {
	var returns ExternalTableTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/external_table snowflake_external_table} Resource.
func NewExternalTable(scope constructs.Construct, id *string, config *ExternalTableConfig) ExternalTable {
	_init_.Initialize()

	if err := validateNewExternalTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalTable{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalTable.ExternalTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/external_table snowflake_external_table} Resource.
func NewExternalTable_Override(e ExternalTable, scope constructs.Construct, id *string, config *ExternalTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalTable.ExternalTable",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExternalTable)SetAutoRefresh(val interface{}) {
	if err := j.validateSetAutoRefreshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRefresh",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetAwsSnsTopic(val *string) {
	if err := j.validateSetAwsSnsTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSnsTopic",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetCopyGrants(val interface{}) {
	if err := j.validateSetCopyGrantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyGrants",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetFileFormat(val *string) {
	if err := j.validateSetFileFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileFormat",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetPartitionBy(val *[]*string) {
	if err := j.validateSetPartitionByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionBy",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetRefreshOnCreate(val interface{}) {
	if err := j.validateSetRefreshOnCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshOnCreate",
		val,
	)
}

func (j *jsiiProxy_ExternalTable)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
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
func ExternalTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalTable.ExternalTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalTable.ExternalTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalTable.ExternalTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExternalTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.externalTable.ExternalTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExternalTable) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExternalTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ExternalTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ExternalTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ExternalTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ExternalTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ExternalTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ExternalTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ExternalTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ExternalTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ExternalTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ExternalTable) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExternalTable) PutColumn(value interface{}) {
	if err := e.validatePutColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putColumn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalTable) PutTag(value interface{}) {
	if err := e.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTag",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalTable) ResetAutoRefresh() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoRefresh",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetAwsSnsTopic() {
	_jsii_.InvokeVoid(
		e,
		"resetAwsSnsTopic",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetComment() {
	_jsii_.InvokeVoid(
		e,
		"resetComment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetCopyGrants() {
	_jsii_.InvokeVoid(
		e,
		"resetCopyGrants",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetPartitionBy() {
	_jsii_.InvokeVoid(
		e,
		"resetPartitionBy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetPattern() {
	_jsii_.InvokeVoid(
		e,
		"resetPattern",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetRefreshOnCreate() {
	_jsii_.InvokeVoid(
		e,
		"resetRefreshOnCreate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) ResetTag() {
	_jsii_.InvokeVoid(
		e,
		"resetTag",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

