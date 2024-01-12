// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package task

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/task/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task snowflake_task}.
type Task interface {
	cdktf.TerraformResource
	After() *[]*string
	SetAfter(val *[]*string)
	AfterInput() *[]*string
	AllowOverlappingExecution() interface{}
	SetAllowOverlappingExecution(val interface{})
	AllowOverlappingExecutionInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ErrorIntegration() *string
	SetErrorIntegration(val *string)
	ErrorIntegrationInput() *string
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
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	SessionParameters() *map[string]*string
	SetSessionParameters(val *map[string]*string)
	SessionParametersInput() *map[string]*string
	SqlStatement() *string
	SetSqlStatement(val *string)
	SqlStatementInput() *string
	SuspendTaskAfterNumFailures() *float64
	SetSuspendTaskAfterNumFailures(val *float64)
	SuspendTaskAfterNumFailuresInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserTaskManagedInitialWarehouseSize() *string
	SetUserTaskManagedInitialWarehouseSize(val *string)
	UserTaskManagedInitialWarehouseSizeInput() *string
	UserTaskTimeoutMs() *float64
	SetUserTaskTimeoutMs(val *float64)
	UserTaskTimeoutMsInput() *float64
	Warehouse() *string
	SetWarehouse(val *string)
	WarehouseInput() *string
	When() *string
	SetWhen(val *string)
	WhenInput() *string
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
	ResetAfter()
	ResetAllowOverlappingExecution()
	ResetComment()
	ResetEnabled()
	ResetErrorIntegration()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchedule()
	ResetSessionParameters()
	ResetSuspendTaskAfterNumFailures()
	ResetUserTaskManagedInitialWarehouseSize()
	ResetUserTaskTimeoutMs()
	ResetWarehouse()
	ResetWhen()
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

// The jsii proxy struct for Task
type jsiiProxy_Task struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Task) After() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"after",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) AfterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"afterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) AllowOverlappingExecution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverlappingExecution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) AllowOverlappingExecutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverlappingExecutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ErrorIntegration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ErrorIntegrationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SessionParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sessionParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SessionParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sessionParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SqlStatement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SqlStatementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SuspendTaskAfterNumFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SuspendTaskAfterNumFailuresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UserTaskManagedInitialWarehouseSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UserTaskManagedInitialWarehouseSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UserTaskTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UserTaskTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Warehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) WarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) When() *string {
	var returns *string
	_jsii_.Get(
		j,
		"when",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) WhenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task snowflake_task} Resource.
func NewTask(scope constructs.Construct, id *string, config *TaskConfig) Task {
	_init_.Initialize()

	if err := validateNewTaskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Task{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.task.Task",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task snowflake_task} Resource.
func NewTask_Override(t Task, scope constructs.Construct, id *string, config *TaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.task.Task",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_Task)SetAfter(val *[]*string) {
	if err := j.validateSetAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"after",
		val,
	)
}

func (j *jsiiProxy_Task)SetAllowOverlappingExecution(val interface{}) {
	if err := j.validateSetAllowOverlappingExecutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOverlappingExecution",
		val,
	)
}

func (j *jsiiProxy_Task)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_Task)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Task)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Task)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_Task)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Task)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_Task)SetErrorIntegration(val *string) {
	if err := j.validateSetErrorIntegrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorIntegration",
		val,
	)
}

func (j *jsiiProxy_Task)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Task)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Task)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Task)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Task)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Task)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Task)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_Task)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_Task)SetSessionParameters(val *map[string]*string) {
	if err := j.validateSetSessionParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionParameters",
		val,
	)
}

func (j *jsiiProxy_Task)SetSqlStatement(val *string) {
	if err := j.validateSetSqlStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlStatement",
		val,
	)
}

func (j *jsiiProxy_Task)SetSuspendTaskAfterNumFailures(val *float64) {
	if err := j.validateSetSuspendTaskAfterNumFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTaskAfterNumFailures",
		val,
	)
}

func (j *jsiiProxy_Task)SetUserTaskManagedInitialWarehouseSize(val *string) {
	if err := j.validateSetUserTaskManagedInitialWarehouseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskManagedInitialWarehouseSize",
		val,
	)
}

func (j *jsiiProxy_Task)SetUserTaskTimeoutMs(val *float64) {
	if err := j.validateSetUserTaskTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_Task)SetWarehouse(val *string) {
	if err := j.validateSetWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouse",
		val,
	)
}

func (j *jsiiProxy_Task)SetWhen(val *string) {
	if err := j.validateSetWhenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"when",
		val,
	)
}

// Generates CDKTF code for importing a Task resource upon running "cdktf plan <stack-name>".
func Task_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateTask_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.task.Task",
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
func Task_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.task.Task",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Task_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTask_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.task.Task",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Task_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTask_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.task.Task",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Task_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.task.Task",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_Task) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_Task) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_Task) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_Task) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_Task) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_Task) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_Task) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_Task) ResetAfter() {
	_jsii_.InvokeVoid(
		t,
		"resetAfter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetAllowOverlappingExecution() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowOverlappingExecution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetComment() {
	_jsii_.InvokeVoid(
		t,
		"resetComment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetErrorIntegration() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorIntegration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetSchedule() {
	_jsii_.InvokeVoid(
		t,
		"resetSchedule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetSessionParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSessionParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetSuspendTaskAfterNumFailures() {
	_jsii_.InvokeVoid(
		t,
		"resetSuspendTaskAfterNumFailures",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetUserTaskManagedInitialWarehouseSize() {
	_jsii_.InvokeVoid(
		t,
		"resetUserTaskManagedInitialWarehouseSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetUserTaskTimeoutMs() {
	_jsii_.InvokeVoid(
		t,
		"resetUserTaskTimeoutMs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetWarehouse() {
	_jsii_.InvokeVoid(
		t,
		"resetWarehouse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetWhen() {
	_jsii_.InvokeVoid(
		t,
		"resetWhen",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

