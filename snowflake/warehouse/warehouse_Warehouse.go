package warehouse

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v3/warehouse/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/snowflake/r/warehouse snowflake_warehouse}.
type Warehouse interface {
	cdktf.TerraformResource
	AutoResume() interface{}
	SetAutoResume(val interface{})
	AutoResumeInput() interface{}
	AutoSuspend() *float64
	SetAutoSuspend(val *float64)
	AutoSuspendInput() *float64
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitiallySuspended() interface{}
	SetInitiallySuspended(val interface{})
	InitiallySuspendedInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxClusterCount() *float64
	SetMaxClusterCount(val *float64)
	MaxClusterCountInput() *float64
	MaxConcurrencyLevel() *float64
	SetMaxConcurrencyLevel(val *float64)
	MaxConcurrencyLevelInput() *float64
	MinClusterCount() *float64
	SetMinClusterCount(val *float64)
	MinClusterCountInput() *float64
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
	ResourceMonitor() *string
	SetResourceMonitor(val *string)
	ResourceMonitorInput() *string
	ScalingPolicy() *string
	SetScalingPolicy(val *string)
	ScalingPolicyInput() *string
	StatementQueuedTimeoutInSeconds() *float64
	SetStatementQueuedTimeoutInSeconds(val *float64)
	StatementQueuedTimeoutInSecondsInput() *float64
	StatementTimeoutInSeconds() *float64
	SetStatementTimeoutInSeconds(val *float64)
	StatementTimeoutInSecondsInput() *float64
	Tag() WarehouseTagList
	TagInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WaitForProvisioning() interface{}
	SetWaitForProvisioning(val interface{})
	WaitForProvisioningInput() interface{}
	WarehouseSize() *string
	SetWarehouseSize(val *string)
	WarehouseSizeInput() *string
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
	PutTag(value interface{})
	ResetAutoResume()
	ResetAutoSuspend()
	ResetComment()
	ResetId()
	ResetInitiallySuspended()
	ResetMaxClusterCount()
	ResetMaxConcurrencyLevel()
	ResetMinClusterCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceMonitor()
	ResetScalingPolicy()
	ResetStatementQueuedTimeoutInSeconds()
	ResetStatementTimeoutInSeconds()
	ResetTag()
	ResetWaitForProvisioning()
	ResetWarehouseSize()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Warehouse
type jsiiProxy_Warehouse struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Warehouse) AutoResume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoResume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) AutoResumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoResumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) AutoSuspend() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoSuspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) AutoSuspendInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoSuspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) InitiallySuspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initiallySuspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) InitiallySuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initiallySuspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) MaxClusterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) MaxClusterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) MaxConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) MaxConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) MinClusterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) MinClusterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) ResourceMonitor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) ResourceMonitorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) ScalingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) ScalingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) StatementQueuedTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) StatementQueuedTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) StatementTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) StatementTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) Tag() WarehouseTagList {
	var returns WarehouseTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) WaitForProvisioning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForProvisioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) WaitForProvisioningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForProvisioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) WarehouseSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Warehouse) WarehouseSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseSizeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/warehouse snowflake_warehouse} Resource.
func NewWarehouse(scope constructs.Construct, id *string, config *WarehouseConfig) Warehouse {
	_init_.Initialize()

	if err := validateNewWarehouseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Warehouse{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.warehouse.Warehouse",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/warehouse snowflake_warehouse} Resource.
func NewWarehouse_Override(w Warehouse, scope constructs.Construct, id *string, config *WarehouseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.warehouse.Warehouse",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_Warehouse)SetAutoResume(val interface{}) {
	if err := j.validateSetAutoResumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoResume",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetAutoSuspend(val *float64) {
	if err := j.validateSetAutoSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSuspend",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetInitiallySuspended(val interface{}) {
	if err := j.validateSetInitiallySuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initiallySuspended",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetMaxClusterCount(val *float64) {
	if err := j.validateSetMaxClusterCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClusterCount",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetMaxConcurrencyLevel(val *float64) {
	if err := j.validateSetMaxConcurrencyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetMinClusterCount(val *float64) {
	if err := j.validateSetMinClusterCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minClusterCount",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetResourceMonitor(val *string) {
	if err := j.validateSetResourceMonitorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceMonitor",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetScalingPolicy(val *string) {
	if err := j.validateSetScalingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingPolicy",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetStatementQueuedTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementQueuedTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementQueuedTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetStatementTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetWaitForProvisioning(val interface{}) {
	if err := j.validateSetWaitForProvisioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForProvisioning",
		val,
	)
}

func (j *jsiiProxy_Warehouse)SetWarehouseSize(val *string) {
	if err := j.validateSetWarehouseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseSize",
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
func Warehouse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWarehouse_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.warehouse.Warehouse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Warehouse_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.warehouse.Warehouse",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_Warehouse) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_Warehouse) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_Warehouse) PutTag(value interface{}) {
	if err := w.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTag",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Warehouse) ResetAutoResume() {
	_jsii_.InvokeVoid(
		w,
		"resetAutoResume",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetAutoSuspend() {
	_jsii_.InvokeVoid(
		w,
		"resetAutoSuspend",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetComment() {
	_jsii_.InvokeVoid(
		w,
		"resetComment",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetInitiallySuspended() {
	_jsii_.InvokeVoid(
		w,
		"resetInitiallySuspended",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetMaxClusterCount() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxClusterCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetMaxConcurrencyLevel() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxConcurrencyLevel",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetMinClusterCount() {
	_jsii_.InvokeVoid(
		w,
		"resetMinClusterCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetResourceMonitor() {
	_jsii_.InvokeVoid(
		w,
		"resetResourceMonitor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetScalingPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetScalingPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetStatementQueuedTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		w,
		"resetStatementQueuedTimeoutInSeconds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetStatementTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		w,
		"resetStatementTimeoutInSeconds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetTag() {
	_jsii_.InvokeVoid(
		w,
		"resetTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetWaitForProvisioning() {
	_jsii_.InvokeVoid(
		w,
		"resetWaitForProvisioning",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) ResetWarehouseSize() {
	_jsii_.InvokeVoid(
		w,
		"resetWarehouseSize",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Warehouse) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Warehouse) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

