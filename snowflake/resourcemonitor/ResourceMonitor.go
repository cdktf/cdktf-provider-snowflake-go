// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/resourcemonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/resource_monitor snowflake_resource_monitor}.
type ResourceMonitor interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreditQuota() *float64
	SetCreditQuota(val *float64)
	CreditQuotaInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndTimestamp() *string
	SetEndTimestamp(val *string)
	EndTimestampInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullyQualifiedName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotifyTriggers() *[]*float64
	SetNotifyTriggers(val *[]*float64)
	NotifyTriggersInput() *[]*float64
	NotifyUsers() *[]*string
	SetNotifyUsers(val *[]*string)
	NotifyUsersInput() *[]*string
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
	ShowOutput() ResourceMonitorShowOutputList
	StartTimestamp() *string
	SetStartTimestamp(val *string)
	StartTimestampInput() *string
	SuspendImmediateTrigger() *float64
	SetSuspendImmediateTrigger(val *float64)
	SuspendImmediateTriggerInput() *float64
	SuspendTrigger() *float64
	SetSuspendTrigger(val *float64)
	SuspendTriggerInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetCreditQuota()
	ResetEndTimestamp()
	ResetFrequency()
	ResetId()
	ResetNotifyTriggers()
	ResetNotifyUsers()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartTimestamp()
	ResetSuspendImmediateTrigger()
	ResetSuspendTrigger()
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

// The jsii proxy struct for ResourceMonitor
type jsiiProxy_ResourceMonitor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResourceMonitor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) CreditQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creditQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) CreditQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creditQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) EndTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) EndTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) NotifyTriggers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"notifyTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) NotifyTriggersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"notifyTriggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) NotifyUsers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) NotifyUsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) ShowOutput() ResourceMonitorShowOutputList {
	var returns ResourceMonitorShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) StartTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) StartTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) SuspendImmediateTrigger() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendImmediateTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) SuspendImmediateTriggerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendImmediateTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) SuspendTrigger() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) SuspendTriggerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceMonitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/resource_monitor snowflake_resource_monitor} Resource.
func NewResourceMonitor(scope constructs.Construct, id *string, config *ResourceMonitorConfig) ResourceMonitor {
	_init_.Initialize()

	if err := validateNewResourceMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceMonitor{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/resource_monitor snowflake_resource_monitor} Resource.
func NewResourceMonitor_Override(r ResourceMonitor, scope constructs.Construct, id *string, config *ResourceMonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitor",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetCreditQuota(val *float64) {
	if err := j.validateSetCreditQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creditQuota",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetEndTimestamp(val *string) {
	if err := j.validateSetEndTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTimestamp",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetNotifyTriggers(val *[]*float64) {
	if err := j.validateSetNotifyTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyTriggers",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetNotifyUsers(val *[]*string) {
	if err := j.validateSetNotifyUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyUsers",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetStartTimestamp(val *string) {
	if err := j.validateSetStartTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimestamp",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetSuspendImmediateTrigger(val *float64) {
	if err := j.validateSetSuspendImmediateTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendImmediateTrigger",
		val,
	)
}

func (j *jsiiProxy_ResourceMonitor)SetSuspendTrigger(val *float64) {
	if err := j.validateSetSuspendTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTrigger",
		val,
	)
}

// Generates CDKTF code for importing a ResourceMonitor resource upon running "cdktf plan <stack-name>".
func ResourceMonitor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateResourceMonitor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitor",
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
func ResourceMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceMonitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceMonitor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceMonitor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResourceMonitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.resourceMonitor.ResourceMonitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ResourceMonitor) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_ResourceMonitor) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ResourceMonitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_ResourceMonitor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceMonitor) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_ResourceMonitor) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetCreditQuota() {
	_jsii_.InvokeVoid(
		r,
		"resetCreditQuota",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetEndTimestamp() {
	_jsii_.InvokeVoid(
		r,
		"resetEndTimestamp",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetFrequency() {
	_jsii_.InvokeVoid(
		r,
		"resetFrequency",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetNotifyTriggers() {
	_jsii_.InvokeVoid(
		r,
		"resetNotifyTriggers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetNotifyUsers() {
	_jsii_.InvokeVoid(
		r,
		"resetNotifyUsers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetStartTimestamp() {
	_jsii_.InvokeVoid(
		r,
		"resetStartTimestamp",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetSuspendImmediateTrigger() {
	_jsii_.InvokeVoid(
		r,
		"resetSuspendImmediateTrigger",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) ResetSuspendTrigger() {
	_jsii_.InvokeVoid(
		r,
		"resetSuspendTrigger",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceMonitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceMonitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

