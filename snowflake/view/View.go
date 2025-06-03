// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/view/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/view snowflake_view}.
type View interface {
	cdktf.TerraformResource
	AggregationPolicy() ViewAggregationPolicyOutputReference
	AggregationPolicyInput() *ViewAggregationPolicy
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChangeTracking() *string
	SetChangeTracking(val *string)
	ChangeTrackingInput() *string
	Column() ViewColumnList
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
	DataMetricFunction() ViewDataMetricFunctionList
	DataMetricFunctionInput() interface{}
	DataMetricSchedule() ViewDataMetricScheduleOutputReference
	DataMetricScheduleInput() *ViewDataMetricSchedule
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DescribeOutput() ViewDescribeOutputList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullyQualifiedName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsRecursive() *string
	SetIsRecursive(val *string)
	IsRecursiveInput() *string
	IsSecure() *string
	SetIsSecure(val *string)
	IsSecureInput() *string
	IsTemporary() *string
	SetIsTemporary(val *string)
	IsTemporaryInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	RowAccessPolicy() ViewRowAccessPolicyOutputReference
	RowAccessPolicyInput() *ViewRowAccessPolicy
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	ShowOutput() ViewShowOutputList
	Statement() *string
	SetStatement(val *string)
	StatementInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ViewTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAggregationPolicy(value *ViewAggregationPolicy)
	PutColumn(value interface{})
	PutDataMetricFunction(value interface{})
	PutDataMetricSchedule(value *ViewDataMetricSchedule)
	PutRowAccessPolicy(value *ViewRowAccessPolicy)
	PutTimeouts(value *ViewTimeouts)
	ResetAggregationPolicy()
	ResetChangeTracking()
	ResetColumn()
	ResetComment()
	ResetCopyGrants()
	ResetDataMetricFunction()
	ResetDataMetricSchedule()
	ResetId()
	ResetIsRecursive()
	ResetIsSecure()
	ResetIsTemporary()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRowAccessPolicy()
	ResetTimeouts()
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

// The jsii proxy struct for View
type jsiiProxy_View struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_View) AggregationPolicy() ViewAggregationPolicyOutputReference {
	var returns ViewAggregationPolicyOutputReference
	_jsii_.Get(
		j,
		"aggregationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) AggregationPolicyInput() *ViewAggregationPolicy {
	var returns *ViewAggregationPolicy
	_jsii_.Get(
		j,
		"aggregationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) ChangeTracking() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeTracking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) ChangeTrackingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeTrackingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Column() ViewColumnList {
	var returns ViewColumnList
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) ColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) CopyGrants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) CopyGrantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyGrantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) DataMetricFunction() ViewDataMetricFunctionList {
	var returns ViewDataMetricFunctionList
	_jsii_.Get(
		j,
		"dataMetricFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) DataMetricFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataMetricFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) DataMetricSchedule() ViewDataMetricScheduleOutputReference {
	var returns ViewDataMetricScheduleOutputReference
	_jsii_.Get(
		j,
		"dataMetricSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) DataMetricScheduleInput() *ViewDataMetricSchedule {
	var returns *ViewDataMetricSchedule
	_jsii_.Get(
		j,
		"dataMetricScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) DescribeOutput() ViewDescribeOutputList {
	var returns ViewDescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) IsRecursive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isRecursive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) IsRecursiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isRecursiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) IsSecure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) IsSecureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isSecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) IsTemporary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isTemporary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) IsTemporaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isTemporaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) RowAccessPolicy() ViewRowAccessPolicyOutputReference {
	var returns ViewRowAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"rowAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) RowAccessPolicyInput() *ViewRowAccessPolicy {
	var returns *ViewRowAccessPolicy
	_jsii_.Get(
		j,
		"rowAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) ShowOutput() ViewShowOutputList {
	var returns ViewShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Statement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) StatementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) Timeouts() ViewTimeoutsOutputReference {
	var returns ViewTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_View) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/view snowflake_view} Resource.
func NewView(scope constructs.Construct, id *string, config *ViewConfig) View {
	_init_.Initialize()

	if err := validateNewViewParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_View{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.view.View",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/view snowflake_view} Resource.
func NewView_Override(v View, scope constructs.Construct, id *string, config *ViewConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.view.View",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_View)SetChangeTracking(val *string) {
	if err := j.validateSetChangeTrackingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeTracking",
		val,
	)
}

func (j *jsiiProxy_View)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_View)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_View)SetCopyGrants(val interface{}) {
	if err := j.validateSetCopyGrantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyGrants",
		val,
	)
}

func (j *jsiiProxy_View)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_View)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_View)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_View)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_View)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_View)SetIsRecursive(val *string) {
	if err := j.validateSetIsRecursiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRecursive",
		val,
	)
}

func (j *jsiiProxy_View)SetIsSecure(val *string) {
	if err := j.validateSetIsSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecure",
		val,
	)
}

func (j *jsiiProxy_View)SetIsTemporary(val *string) {
	if err := j.validateSetIsTemporaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isTemporary",
		val,
	)
}

func (j *jsiiProxy_View)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_View)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_View)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_View)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_View)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_View)SetStatement(val *string) {
	if err := j.validateSetStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statement",
		val,
	)
}

// Generates CDKTF code for importing a View resource upon running "cdktf plan <stack-name>".
func View_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateView_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.view.View",
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
func View_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateView_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.view.View",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func View_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateView_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.view.View",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func View_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateView_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.view.View",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func View_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.view.View",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_View) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_View) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_View) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_View) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_View) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_View) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_View) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_View) PutAggregationPolicy(value *ViewAggregationPolicy) {
	if err := v.validatePutAggregationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAggregationPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_View) PutColumn(value interface{}) {
	if err := v.validatePutColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putColumn",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_View) PutDataMetricFunction(value interface{}) {
	if err := v.validatePutDataMetricFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDataMetricFunction",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_View) PutDataMetricSchedule(value *ViewDataMetricSchedule) {
	if err := v.validatePutDataMetricScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDataMetricSchedule",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_View) PutRowAccessPolicy(value *ViewRowAccessPolicy) {
	if err := v.validatePutRowAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRowAccessPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_View) PutTimeouts(value *ViewTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_View) ResetAggregationPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetAggregationPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetChangeTracking() {
	_jsii_.InvokeVoid(
		v,
		"resetChangeTracking",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetColumn() {
	_jsii_.InvokeVoid(
		v,
		"resetColumn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetComment() {
	_jsii_.InvokeVoid(
		v,
		"resetComment",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetCopyGrants() {
	_jsii_.InvokeVoid(
		v,
		"resetCopyGrants",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetDataMetricFunction() {
	_jsii_.InvokeVoid(
		v,
		"resetDataMetricFunction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetDataMetricSchedule() {
	_jsii_.InvokeVoid(
		v,
		"resetDataMetricSchedule",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetIsRecursive() {
	_jsii_.InvokeVoid(
		v,
		"resetIsRecursive",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetIsSecure() {
	_jsii_.InvokeVoid(
		v,
		"resetIsSecure",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetIsTemporary() {
	_jsii_.InvokeVoid(
		v,
		"resetIsTemporary",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetRowAccessPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetRowAccessPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_View) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_View) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

