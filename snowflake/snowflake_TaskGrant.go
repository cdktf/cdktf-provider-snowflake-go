// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-snowflake-go/snowflake/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-snowflake-go/snowflake/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/snowflake/r/task_grant snowflake_task_grant}.
type TaskGrant interface {
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableMultipleGrants() interface{}
	SetEnableMultipleGrants(val interface{})
	EnableMultipleGrantsInput() interface{}
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
	// The tree node.
	Node() constructs.Node
	OnFuture() interface{}
	SetOnFuture(val interface{})
	OnFutureInput() interface{}
	Privilege() *string
	SetPrivilege(val *string)
	PrivilegeInput() *string
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
	Roles() *[]*string
	SetRoles(val *[]*string)
	RolesInput() *[]*string
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
	TaskName() *string
	SetTaskName(val *string)
	TaskNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithGrantOption() interface{}
	SetWithGrantOption(val interface{})
	WithGrantOptionInput() interface{}
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
	ResetEnableMultipleGrants()
	ResetId()
	ResetOnFuture()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivilege()
	ResetRoles()
	ResetTaskName()
	ResetWithGrantOption()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TaskGrant
type jsiiProxy_TaskGrant struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TaskGrant) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) EnableMultipleGrants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) EnableMultipleGrantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleGrantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) OnFuture() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFuture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) OnFutureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFutureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Privilege() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilege",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) PrivilegeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) RolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) TaskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) TaskNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) WithGrantOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskGrant) WithGrantOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/task_grant snowflake_task_grant} Resource.
func NewTaskGrant(scope constructs.Construct, id *string, config *TaskGrantConfig) TaskGrant {
	_init_.Initialize()

	j := jsiiProxy_TaskGrant{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.TaskGrant",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/task_grant snowflake_task_grant} Resource.
func NewTaskGrant_Override(t TaskGrant, scope constructs.Construct, id *string, config *TaskGrantConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.TaskGrant",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TaskGrant) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetEnableMultipleGrants(val interface{}) {
	_jsii_.Set(
		j,
		"enableMultipleGrants",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetOnFuture(val interface{}) {
	_jsii_.Set(
		j,
		"onFuture",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetPrivilege(val *string) {
	_jsii_.Set(
		j,
		"privilege",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetSchemaName(val *string) {
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetTaskName(val *string) {
	_jsii_.Set(
		j,
		"taskName",
		val,
	)
}

func (j *jsiiProxy_TaskGrant) SetWithGrantOption(val interface{}) {
	_jsii_.Set(
		j,
		"withGrantOption",
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
func TaskGrant_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.TaskGrant",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TaskGrant_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.TaskGrant",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TaskGrant) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TaskGrant) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TaskGrant) ResetEnableMultipleGrants() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableMultipleGrants",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskGrant) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskGrant) ResetOnFuture() {
	_jsii_.InvokeVoid(
		t,
		"resetOnFuture",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskGrant) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskGrant) ResetPrivilege() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivilege",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskGrant) ResetRoles() {
	_jsii_.InvokeVoid(
		t,
		"resetRoles",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskGrant) ResetTaskName() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskGrant) ResetWithGrantOption() {
	_jsii_.InvokeVoid(
		t,
		"resetWithGrantOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskGrant) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskGrant) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
