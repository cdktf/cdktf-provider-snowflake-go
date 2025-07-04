// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package failovergroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/failovergroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/failover_group snowflake_failover_group}.
type FailoverGroup interface {
	cdktf.TerraformResource
	AllowedAccounts() *[]*string
	SetAllowedAccounts(val *[]*string)
	AllowedAccountsInput() *[]*string
	AllowedDatabases() *[]*string
	SetAllowedDatabases(val *[]*string)
	AllowedDatabasesInput() *[]*string
	AllowedIntegrationTypes() *[]*string
	SetAllowedIntegrationTypes(val *[]*string)
	AllowedIntegrationTypesInput() *[]*string
	AllowedShares() *[]*string
	SetAllowedShares(val *[]*string)
	AllowedSharesInput() *[]*string
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
	FromReplica() FailoverGroupFromReplicaOutputReference
	FromReplicaInput() *FailoverGroupFromReplica
	FullyQualifiedName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreEditionCheck() interface{}
	SetIgnoreEditionCheck(val interface{})
	IgnoreEditionCheckInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ObjectTypes() *[]*string
	SetObjectTypes(val *[]*string)
	ObjectTypesInput() *[]*string
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
	ReplicationSchedule() FailoverGroupReplicationScheduleOutputReference
	ReplicationScheduleInput() *FailoverGroupReplicationSchedule
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FailoverGroupTimeoutsOutputReference
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
	PutFromReplica(value *FailoverGroupFromReplica)
	PutReplicationSchedule(value *FailoverGroupReplicationSchedule)
	PutTimeouts(value *FailoverGroupTimeouts)
	ResetAllowedAccounts()
	ResetAllowedDatabases()
	ResetAllowedIntegrationTypes()
	ResetAllowedShares()
	ResetFromReplica()
	ResetId()
	ResetIgnoreEditionCheck()
	ResetObjectTypes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplicationSchedule()
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

// The jsii proxy struct for FailoverGroup
type jsiiProxy_FailoverGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FailoverGroup) AllowedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) AllowedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) AllowedDatabases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) AllowedDatabasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) AllowedIntegrationTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIntegrationTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) AllowedIntegrationTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIntegrationTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) AllowedShares() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) AllowedSharesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) FromReplica() FailoverGroupFromReplicaOutputReference {
	var returns FailoverGroupFromReplicaOutputReference
	_jsii_.Get(
		j,
		"fromReplica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) FromReplicaInput() *FailoverGroupFromReplica {
	var returns *FailoverGroupFromReplica
	_jsii_.Get(
		j,
		"fromReplicaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) IgnoreEditionCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreEditionCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) IgnoreEditionCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreEditionCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) ObjectTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) ObjectTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) ReplicationSchedule() FailoverGroupReplicationScheduleOutputReference {
	var returns FailoverGroupReplicationScheduleOutputReference
	_jsii_.Get(
		j,
		"replicationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) ReplicationScheduleInput() *FailoverGroupReplicationSchedule {
	var returns *FailoverGroupReplicationSchedule
	_jsii_.Get(
		j,
		"replicationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) Timeouts() FailoverGroupTimeoutsOutputReference {
	var returns FailoverGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FailoverGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/failover_group snowflake_failover_group} Resource.
func NewFailoverGroup(scope constructs.Construct, id *string, config *FailoverGroupConfig) FailoverGroup {
	_init_.Initialize()

	if err := validateNewFailoverGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FailoverGroup{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.failoverGroup.FailoverGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/failover_group snowflake_failover_group} Resource.
func NewFailoverGroup_Override(f FailoverGroup, scope constructs.Construct, id *string, config *FailoverGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.failoverGroup.FailoverGroup",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FailoverGroup)SetAllowedAccounts(val *[]*string) {
	if err := j.validateSetAllowedAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAccounts",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetAllowedDatabases(val *[]*string) {
	if err := j.validateSetAllowedDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDatabases",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetAllowedIntegrationTypes(val *[]*string) {
	if err := j.validateSetAllowedIntegrationTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIntegrationTypes",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetAllowedShares(val *[]*string) {
	if err := j.validateSetAllowedSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedShares",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetIgnoreEditionCheck(val interface{}) {
	if err := j.validateSetIgnoreEditionCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreEditionCheck",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetObjectTypes(val *[]*string) {
	if err := j.validateSetObjectTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectTypes",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FailoverGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a FailoverGroup resource upon running "cdktf plan <stack-name>".
func FailoverGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFailoverGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.failoverGroup.FailoverGroup",
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
func FailoverGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFailoverGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.failoverGroup.FailoverGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FailoverGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFailoverGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.failoverGroup.FailoverGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FailoverGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFailoverGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.failoverGroup.FailoverGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FailoverGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.failoverGroup.FailoverGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FailoverGroup) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FailoverGroup) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FailoverGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FailoverGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FailoverGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FailoverGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FailoverGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FailoverGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FailoverGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FailoverGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FailoverGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FailoverGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FailoverGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FailoverGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FailoverGroup) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FailoverGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FailoverGroup) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FailoverGroup) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FailoverGroup) PutFromReplica(value *FailoverGroupFromReplica) {
	if err := f.validatePutFromReplicaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFromReplica",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FailoverGroup) PutReplicationSchedule(value *FailoverGroupReplicationSchedule) {
	if err := f.validatePutReplicationScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putReplicationSchedule",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FailoverGroup) PutTimeouts(value *FailoverGroupTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FailoverGroup) ResetAllowedAccounts() {
	_jsii_.InvokeVoid(
		f,
		"resetAllowedAccounts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetAllowedDatabases() {
	_jsii_.InvokeVoid(
		f,
		"resetAllowedDatabases",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetAllowedIntegrationTypes() {
	_jsii_.InvokeVoid(
		f,
		"resetAllowedIntegrationTypes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetAllowedShares() {
	_jsii_.InvokeVoid(
		f,
		"resetAllowedShares",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetFromReplica() {
	_jsii_.InvokeVoid(
		f,
		"resetFromReplica",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetIgnoreEditionCheck() {
	_jsii_.InvokeVoid(
		f,
		"resetIgnoreEditionCheck",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetObjectTypes() {
	_jsii_.InvokeVoid(
		f,
		"resetObjectTypes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetReplicationSchedule() {
	_jsii_.InvokeVoid(
		f,
		"resetReplicationSchedule",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FailoverGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FailoverGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FailoverGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FailoverGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FailoverGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FailoverGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

