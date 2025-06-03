// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package task

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/task/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/task snowflake_task}.
type Task interface {
	cdktf.TerraformResource
	AbortDetachedQuery() interface{}
	SetAbortDetachedQuery(val interface{})
	AbortDetachedQueryInput() interface{}
	After() *[]*string
	SetAfter(val *[]*string)
	AfterInput() *[]*string
	AllowOverlappingExecution() *string
	SetAllowOverlappingExecution(val *string)
	AllowOverlappingExecutionInput() *string
	Autocommit() interface{}
	SetAutocommit(val interface{})
	AutocommitInput() interface{}
	BinaryInputFormat() *string
	SetBinaryInputFormat(val *string)
	BinaryInputFormatInput() *string
	BinaryOutputFormat() *string
	SetBinaryOutputFormat(val *string)
	BinaryOutputFormatInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientMemoryLimit() *float64
	SetClientMemoryLimit(val *float64)
	ClientMemoryLimitInput() *float64
	ClientMetadataRequestUseConnectionCtx() interface{}
	SetClientMetadataRequestUseConnectionCtx(val interface{})
	ClientMetadataRequestUseConnectionCtxInput() interface{}
	ClientPrefetchThreads() *float64
	SetClientPrefetchThreads(val *float64)
	ClientPrefetchThreadsInput() *float64
	ClientResultChunkSize() *float64
	SetClientResultChunkSize(val *float64)
	ClientResultChunkSizeInput() *float64
	ClientResultColumnCaseInsensitive() interface{}
	SetClientResultColumnCaseInsensitive(val interface{})
	ClientResultColumnCaseInsensitiveInput() interface{}
	ClientSessionKeepAlive() interface{}
	SetClientSessionKeepAlive(val interface{})
	ClientSessionKeepAliveHeartbeatFrequency() *float64
	SetClientSessionKeepAliveHeartbeatFrequency(val *float64)
	ClientSessionKeepAliveHeartbeatFrequencyInput() *float64
	ClientSessionKeepAliveInput() interface{}
	ClientTimestampTypeMapping() *string
	SetClientTimestampTypeMapping(val *string)
	ClientTimestampTypeMappingInput() *string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	Config() *string
	SetConfig(val *string)
	ConfigInput() *string
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
	DateInputFormat() *string
	SetDateInputFormat(val *string)
	DateInputFormatInput() *string
	DateOutputFormat() *string
	SetDateOutputFormat(val *string)
	DateOutputFormatInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableUnloadPhysicalTypeOptimization() interface{}
	SetEnableUnloadPhysicalTypeOptimization(val interface{})
	EnableUnloadPhysicalTypeOptimizationInput() interface{}
	ErrorIntegration() *string
	SetErrorIntegration(val *string)
	ErrorIntegrationInput() *string
	ErrorOnNondeterministicMerge() interface{}
	SetErrorOnNondeterministicMerge(val interface{})
	ErrorOnNondeterministicMergeInput() interface{}
	ErrorOnNondeterministicUpdate() interface{}
	SetErrorOnNondeterministicUpdate(val interface{})
	ErrorOnNondeterministicUpdateInput() interface{}
	Finalize() *string
	SetFinalize(val *string)
	FinalizeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullyQualifiedName() *string
	GeographyOutputFormat() *string
	SetGeographyOutputFormat(val *string)
	GeographyOutputFormatInput() *string
	GeometryOutputFormat() *string
	SetGeometryOutputFormat(val *string)
	GeometryOutputFormatInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	JdbcTreatTimestampNtzAsUtc() interface{}
	SetJdbcTreatTimestampNtzAsUtc(val interface{})
	JdbcTreatTimestampNtzAsUtcInput() interface{}
	JdbcUseSessionTimezone() interface{}
	SetJdbcUseSessionTimezone(val interface{})
	JdbcUseSessionTimezoneInput() interface{}
	JsonIndent() *float64
	SetJsonIndent(val *float64)
	JsonIndentInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LockTimeout() *float64
	SetLockTimeout(val *float64)
	LockTimeoutInput() *float64
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	MultiStatementCount() *float64
	SetMultiStatementCount(val *float64)
	MultiStatementCountInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NoorderSequenceAsDefault() interface{}
	SetNoorderSequenceAsDefault(val interface{})
	NoorderSequenceAsDefaultInput() interface{}
	OdbcTreatDecimalAsInt() interface{}
	SetOdbcTreatDecimalAsInt(val interface{})
	OdbcTreatDecimalAsIntInput() interface{}
	Parameters() TaskParametersList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueryTag() *string
	SetQueryTag(val *string)
	QueryTagInput() *string
	QuotedIdentifiersIgnoreCase() interface{}
	SetQuotedIdentifiersIgnoreCase(val interface{})
	QuotedIdentifiersIgnoreCaseInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RowsPerResultset() *float64
	SetRowsPerResultset(val *float64)
	RowsPerResultsetInput() *float64
	S3StageVpceDnsName() *string
	SetS3StageVpceDnsName(val *string)
	S3StageVpceDnsNameInput() *string
	Schedule() TaskScheduleOutputReference
	ScheduleInput() *TaskSchedule
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	SearchPath() *string
	SetSearchPath(val *string)
	SearchPathInput() *string
	ShowOutput() TaskShowOutputList
	SqlStatement() *string
	SetSqlStatement(val *string)
	SqlStatementInput() *string
	Started() interface{}
	SetStarted(val interface{})
	StartedInput() interface{}
	StatementQueuedTimeoutInSeconds() *float64
	SetStatementQueuedTimeoutInSeconds(val *float64)
	StatementQueuedTimeoutInSecondsInput() *float64
	StatementTimeoutInSeconds() *float64
	SetStatementTimeoutInSeconds(val *float64)
	StatementTimeoutInSecondsInput() *float64
	StrictJsonOutput() interface{}
	SetStrictJsonOutput(val interface{})
	StrictJsonOutputInput() interface{}
	SuspendTaskAfterNumFailures() *float64
	SetSuspendTaskAfterNumFailures(val *float64)
	SuspendTaskAfterNumFailuresInput() *float64
	TaskAutoRetryAttempts() *float64
	SetTaskAutoRetryAttempts(val *float64)
	TaskAutoRetryAttemptsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeInputFormat() *string
	SetTimeInputFormat(val *string)
	TimeInputFormatInput() *string
	TimeOutputFormat() *string
	SetTimeOutputFormat(val *string)
	TimeOutputFormatInput() *string
	Timeouts() TaskTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimestampDayIsAlways24H() interface{}
	SetTimestampDayIsAlways24H(val interface{})
	TimestampDayIsAlways24HInput() interface{}
	TimestampInputFormat() *string
	SetTimestampInputFormat(val *string)
	TimestampInputFormatInput() *string
	TimestampLtzOutputFormat() *string
	SetTimestampLtzOutputFormat(val *string)
	TimestampLtzOutputFormatInput() *string
	TimestampNtzOutputFormat() *string
	SetTimestampNtzOutputFormat(val *string)
	TimestampNtzOutputFormatInput() *string
	TimestampOutputFormat() *string
	SetTimestampOutputFormat(val *string)
	TimestampOutputFormatInput() *string
	TimestampTypeMapping() *string
	SetTimestampTypeMapping(val *string)
	TimestampTypeMappingInput() *string
	TimestampTzOutputFormat() *string
	SetTimestampTzOutputFormat(val *string)
	TimestampTzOutputFormatInput() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	TraceLevel() *string
	SetTraceLevel(val *string)
	TraceLevelInput() *string
	TransactionAbortOnError() interface{}
	SetTransactionAbortOnError(val interface{})
	TransactionAbortOnErrorInput() interface{}
	TransactionDefaultIsolationLevel() *string
	SetTransactionDefaultIsolationLevel(val *string)
	TransactionDefaultIsolationLevelInput() *string
	TwoDigitCenturyStart() *float64
	SetTwoDigitCenturyStart(val *float64)
	TwoDigitCenturyStartInput() *float64
	UnsupportedDdlAction() *string
	SetUnsupportedDdlAction(val *string)
	UnsupportedDdlActionInput() *string
	UseCachedResult() interface{}
	SetUseCachedResult(val interface{})
	UseCachedResultInput() interface{}
	UserTaskManagedInitialWarehouseSize() *string
	SetUserTaskManagedInitialWarehouseSize(val *string)
	UserTaskManagedInitialWarehouseSizeInput() *string
	UserTaskMinimumTriggerIntervalInSeconds() *float64
	SetUserTaskMinimumTriggerIntervalInSeconds(val *float64)
	UserTaskMinimumTriggerIntervalInSecondsInput() *float64
	UserTaskTimeoutMs() *float64
	SetUserTaskTimeoutMs(val *float64)
	UserTaskTimeoutMsInput() *float64
	Warehouse() *string
	SetWarehouse(val *string)
	WarehouseInput() *string
	WeekOfYearPolicy() *float64
	SetWeekOfYearPolicy(val *float64)
	WeekOfYearPolicyInput() *float64
	WeekStart() *float64
	SetWeekStart(val *float64)
	WeekStartInput() *float64
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
	PutSchedule(value *TaskSchedule)
	PutTimeouts(value *TaskTimeouts)
	ResetAbortDetachedQuery()
	ResetAfter()
	ResetAllowOverlappingExecution()
	ResetAutocommit()
	ResetBinaryInputFormat()
	ResetBinaryOutputFormat()
	ResetClientMemoryLimit()
	ResetClientMetadataRequestUseConnectionCtx()
	ResetClientPrefetchThreads()
	ResetClientResultChunkSize()
	ResetClientResultColumnCaseInsensitive()
	ResetClientSessionKeepAlive()
	ResetClientSessionKeepAliveHeartbeatFrequency()
	ResetClientTimestampTypeMapping()
	ResetComment()
	ResetConfig()
	ResetDateInputFormat()
	ResetDateOutputFormat()
	ResetEnableUnloadPhysicalTypeOptimization()
	ResetErrorIntegration()
	ResetErrorOnNondeterministicMerge()
	ResetErrorOnNondeterministicUpdate()
	ResetFinalize()
	ResetGeographyOutputFormat()
	ResetGeometryOutputFormat()
	ResetId()
	ResetJdbcTreatTimestampNtzAsUtc()
	ResetJdbcUseSessionTimezone()
	ResetJsonIndent()
	ResetLockTimeout()
	ResetLogLevel()
	ResetMultiStatementCount()
	ResetNoorderSequenceAsDefault()
	ResetOdbcTreatDecimalAsInt()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueryTag()
	ResetQuotedIdentifiersIgnoreCase()
	ResetRowsPerResultset()
	ResetS3StageVpceDnsName()
	ResetSchedule()
	ResetSearchPath()
	ResetStatementQueuedTimeoutInSeconds()
	ResetStatementTimeoutInSeconds()
	ResetStrictJsonOutput()
	ResetSuspendTaskAfterNumFailures()
	ResetTaskAutoRetryAttempts()
	ResetTimeInputFormat()
	ResetTimeOutputFormat()
	ResetTimeouts()
	ResetTimestampDayIsAlways24H()
	ResetTimestampInputFormat()
	ResetTimestampLtzOutputFormat()
	ResetTimestampNtzOutputFormat()
	ResetTimestampOutputFormat()
	ResetTimestampTypeMapping()
	ResetTimestampTzOutputFormat()
	ResetTimezone()
	ResetTraceLevel()
	ResetTransactionAbortOnError()
	ResetTransactionDefaultIsolationLevel()
	ResetTwoDigitCenturyStart()
	ResetUnsupportedDdlAction()
	ResetUseCachedResult()
	ResetUserTaskManagedInitialWarehouseSize()
	ResetUserTaskMinimumTriggerIntervalInSeconds()
	ResetUserTaskTimeoutMs()
	ResetWarehouse()
	ResetWeekOfYearPolicy()
	ResetWeekStart()
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

func (j *jsiiProxy_Task) AbortDetachedQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) AbortDetachedQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQueryInput",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_Task) AllowOverlappingExecution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowOverlappingExecution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) AllowOverlappingExecutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowOverlappingExecutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Autocommit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) AutocommitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) BinaryInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) BinaryInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) BinaryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) BinaryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormatInput",
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

func (j *jsiiProxy_Task) ClientMemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientMemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientMetadataRequestUseConnectionCtx() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientMetadataRequestUseConnectionCtxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientPrefetchThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientPrefetchThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientResultChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientResultChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientResultColumnCaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientResultColumnCaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientSessionKeepAlive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientSessionKeepAliveHeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientSessionKeepAliveHeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientSessionKeepAliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientTimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ClientTimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMappingInput",
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

func (j *jsiiProxy_Task) Config() *string {
	var returns *string
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configInput",
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

func (j *jsiiProxy_Task) DateInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) DateInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) DateOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) DateOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormatInput",
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

func (j *jsiiProxy_Task) EnableUnloadPhysicalTypeOptimization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) EnableUnloadPhysicalTypeOptimizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimizationInput",
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

func (j *jsiiProxy_Task) ErrorOnNondeterministicMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ErrorOnNondeterministicMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ErrorOnNondeterministicUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ErrorOnNondeterministicUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Finalize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) FinalizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalizeInput",
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

func (j *jsiiProxy_Task) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) GeographyOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) GeographyOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) GeometryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) GeometryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormatInput",
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

func (j *jsiiProxy_Task) JdbcTreatTimestampNtzAsUtc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) JdbcTreatTimestampNtzAsUtcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) JdbcUseSessionTimezone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) JdbcUseSessionTimezoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) JsonIndent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) JsonIndentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndentInput",
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

func (j *jsiiProxy_Task) LockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) LockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) MultiStatementCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) MultiStatementCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCountInput",
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

func (j *jsiiProxy_Task) NoorderSequenceAsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) NoorderSequenceAsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) OdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) OdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Parameters() TaskParametersList {
	var returns TaskParametersList
	_jsii_.Get(
		j,
		"parameters",
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

func (j *jsiiProxy_Task) QueryTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) QueryTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) QuotedIdentifiersIgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) QuotedIdentifiersIgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCaseInput",
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

func (j *jsiiProxy_Task) RowsPerResultset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) RowsPerResultsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) S3StageVpceDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) S3StageVpceDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Schedule() TaskScheduleOutputReference {
	var returns TaskScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ScheduleInput() *TaskSchedule {
	var returns *TaskSchedule
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

func (j *jsiiProxy_Task) SearchPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) SearchPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ShowOutput() TaskShowOutputList {
	var returns TaskShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
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

func (j *jsiiProxy_Task) Started() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"started",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StartedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StatementQueuedTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StatementQueuedTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StatementTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StatementTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StrictJsonOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StrictJsonOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutputInput",
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

func (j *jsiiProxy_Task) TaskAutoRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TaskAutoRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttemptsInput",
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

func (j *jsiiProxy_Task) TimeInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimeInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimeOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimeOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Timeouts() TaskTimeoutsOutputReference {
	var returns TaskTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampDayIsAlways24H() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampDayIsAlways24HInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24HInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampLtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampLtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampNtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampNtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampTzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimestampTzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TransactionAbortOnError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TransactionAbortOnErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TransactionDefaultIsolationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TransactionDefaultIsolationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TwoDigitCenturyStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) TwoDigitCenturyStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UnsupportedDdlAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UnsupportedDdlActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UseCachedResult() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UseCachedResultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResultInput",
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

func (j *jsiiProxy_Task) UserTaskMinimumTriggerIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) UserTaskMinimumTriggerIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSecondsInput",
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

func (j *jsiiProxy_Task) WeekOfYearPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) WeekOfYearPolicyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) WeekStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) WeekStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStartInput",
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


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/task snowflake_task} Resource.
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

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/task snowflake_task} Resource.
func NewTask_Override(t Task, scope constructs.Construct, id *string, config *TaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.task.Task",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_Task)SetAbortDetachedQuery(val interface{}) {
	if err := j.validateSetAbortDetachedQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortDetachedQuery",
		val,
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

func (j *jsiiProxy_Task)SetAllowOverlappingExecution(val *string) {
	if err := j.validateSetAllowOverlappingExecutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOverlappingExecution",
		val,
	)
}

func (j *jsiiProxy_Task)SetAutocommit(val interface{}) {
	if err := j.validateSetAutocommitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autocommit",
		val,
	)
}

func (j *jsiiProxy_Task)SetBinaryInputFormat(val *string) {
	if err := j.validateSetBinaryInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryInputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetBinaryOutputFormat(val *string) {
	if err := j.validateSetBinaryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetClientMemoryLimit(val *float64) {
	if err := j.validateSetClientMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMemoryLimit",
		val,
	)
}

func (j *jsiiProxy_Task)SetClientMetadataRequestUseConnectionCtx(val interface{}) {
	if err := j.validateSetClientMetadataRequestUseConnectionCtxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadataRequestUseConnectionCtx",
		val,
	)
}

func (j *jsiiProxy_Task)SetClientPrefetchThreads(val *float64) {
	if err := j.validateSetClientPrefetchThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientPrefetchThreads",
		val,
	)
}

func (j *jsiiProxy_Task)SetClientResultChunkSize(val *float64) {
	if err := j.validateSetClientResultChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultChunkSize",
		val,
	)
}

func (j *jsiiProxy_Task)SetClientResultColumnCaseInsensitive(val interface{}) {
	if err := j.validateSetClientResultColumnCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultColumnCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_Task)SetClientSessionKeepAlive(val interface{}) {
	if err := j.validateSetClientSessionKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAlive",
		val,
	)
}

func (j *jsiiProxy_Task)SetClientSessionKeepAliveHeartbeatFrequency(val *float64) {
	if err := j.validateSetClientSessionKeepAliveHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_Task)SetClientTimestampTypeMapping(val *string) {
	if err := j.validateSetClientTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTimestampTypeMapping",
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

func (j *jsiiProxy_Task)SetConfig(val *string) {
	if err := j.validateSetConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"config",
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

func (j *jsiiProxy_Task)SetDateInputFormat(val *string) {
	if err := j.validateSetDateInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateInputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetDateOutputFormat(val *string) {
	if err := j.validateSetDateOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateOutputFormat",
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

func (j *jsiiProxy_Task)SetEnableUnloadPhysicalTypeOptimization(val interface{}) {
	if err := j.validateSetEnableUnloadPhysicalTypeOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnloadPhysicalTypeOptimization",
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

func (j *jsiiProxy_Task)SetErrorOnNondeterministicMerge(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicMerge",
		val,
	)
}

func (j *jsiiProxy_Task)SetErrorOnNondeterministicUpdate(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicUpdate",
		val,
	)
}

func (j *jsiiProxy_Task)SetFinalize(val *string) {
	if err := j.validateSetFinalizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalize",
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

func (j *jsiiProxy_Task)SetGeographyOutputFormat(val *string) {
	if err := j.validateSetGeographyOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geographyOutputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetGeometryOutputFormat(val *string) {
	if err := j.validateSetGeometryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geometryOutputFormat",
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

func (j *jsiiProxy_Task)SetJdbcTreatTimestampNtzAsUtc(val interface{}) {
	if err := j.validateSetJdbcTreatTimestampNtzAsUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		val,
	)
}

func (j *jsiiProxy_Task)SetJdbcUseSessionTimezone(val interface{}) {
	if err := j.validateSetJdbcUseSessionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcUseSessionTimezone",
		val,
	)
}

func (j *jsiiProxy_Task)SetJsonIndent(val *float64) {
	if err := j.validateSetJsonIndentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonIndent",
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

func (j *jsiiProxy_Task)SetLockTimeout(val *float64) {
	if err := j.validateSetLockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockTimeout",
		val,
	)
}

func (j *jsiiProxy_Task)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_Task)SetMultiStatementCount(val *float64) {
	if err := j.validateSetMultiStatementCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiStatementCount",
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

func (j *jsiiProxy_Task)SetNoorderSequenceAsDefault(val interface{}) {
	if err := j.validateSetNoorderSequenceAsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noorderSequenceAsDefault",
		val,
	)
}

func (j *jsiiProxy_Task)SetOdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetOdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbcTreatDecimalAsInt",
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

func (j *jsiiProxy_Task)SetQueryTag(val *string) {
	if err := j.validateSetQueryTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryTag",
		val,
	)
}

func (j *jsiiProxy_Task)SetQuotedIdentifiersIgnoreCase(val interface{}) {
	if err := j.validateSetQuotedIdentifiersIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotedIdentifiersIgnoreCase",
		val,
	)
}

func (j *jsiiProxy_Task)SetRowsPerResultset(val *float64) {
	if err := j.validateSetRowsPerResultsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsPerResultset",
		val,
	)
}

func (j *jsiiProxy_Task)SetS3StageVpceDnsName(val *string) {
	if err := j.validateSetS3StageVpceDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3StageVpceDnsName",
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

func (j *jsiiProxy_Task)SetSearchPath(val *string) {
	if err := j.validateSetSearchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchPath",
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

func (j *jsiiProxy_Task)SetStarted(val interface{}) {
	if err := j.validateSetStartedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"started",
		val,
	)
}

func (j *jsiiProxy_Task)SetStatementQueuedTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementQueuedTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementQueuedTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_Task)SetStatementTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_Task)SetStrictJsonOutput(val interface{}) {
	if err := j.validateSetStrictJsonOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictJsonOutput",
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

func (j *jsiiProxy_Task)SetTaskAutoRetryAttempts(val *float64) {
	if err := j.validateSetTaskAutoRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAutoRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimeInputFormat(val *string) {
	if err := j.validateSetTimeInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeInputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimeOutputFormat(val *string) {
	if err := j.validateSetTimeOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOutputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimestampDayIsAlways24H(val interface{}) {
	if err := j.validateSetTimestampDayIsAlways24HParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDayIsAlways24H",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimestampInputFormat(val *string) {
	if err := j.validateSetTimestampInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampInputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimestampLtzOutputFormat(val *string) {
	if err := j.validateSetTimestampLtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampLtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimestampNtzOutputFormat(val *string) {
	if err := j.validateSetTimestampNtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampNtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimestampOutputFormat(val *string) {
	if err := j.validateSetTimestampOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOutputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimestampTypeMapping(val *string) {
	if err := j.validateSetTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimestampTzOutputFormat(val *string) {
	if err := j.validateSetTimestampTzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_Task)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_Task)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

func (j *jsiiProxy_Task)SetTransactionAbortOnError(val interface{}) {
	if err := j.validateSetTransactionAbortOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionAbortOnError",
		val,
	)
}

func (j *jsiiProxy_Task)SetTransactionDefaultIsolationLevel(val *string) {
	if err := j.validateSetTransactionDefaultIsolationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionDefaultIsolationLevel",
		val,
	)
}

func (j *jsiiProxy_Task)SetTwoDigitCenturyStart(val *float64) {
	if err := j.validateSetTwoDigitCenturyStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoDigitCenturyStart",
		val,
	)
}

func (j *jsiiProxy_Task)SetUnsupportedDdlAction(val *string) {
	if err := j.validateSetUnsupportedDdlActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unsupportedDdlAction",
		val,
	)
}

func (j *jsiiProxy_Task)SetUseCachedResult(val interface{}) {
	if err := j.validateSetUseCachedResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCachedResult",
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

func (j *jsiiProxy_Task)SetUserTaskMinimumTriggerIntervalInSeconds(val *float64) {
	if err := j.validateSetUserTaskMinimumTriggerIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
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

func (j *jsiiProxy_Task)SetWeekOfYearPolicy(val *float64) {
	if err := j.validateSetWeekOfYearPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekOfYearPolicy",
		val,
	)
}

func (j *jsiiProxy_Task)SetWeekStart(val *float64) {
	if err := j.validateSetWeekStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekStart",
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

func (t *jsiiProxy_Task) PutSchedule(value *TaskSchedule) {
	if err := t.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchedule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_Task) PutTimeouts(value *TaskTimeouts) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_Task) ResetAbortDetachedQuery() {
	_jsii_.InvokeVoid(
		t,
		"resetAbortDetachedQuery",
		nil, // no parameters
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

func (t *jsiiProxy_Task) ResetAutocommit() {
	_jsii_.InvokeVoid(
		t,
		"resetAutocommit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetBinaryInputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetBinaryInputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetBinaryOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetBinaryOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetClientMemoryLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetClientMemoryLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetClientMetadataRequestUseConnectionCtx() {
	_jsii_.InvokeVoid(
		t,
		"resetClientMetadataRequestUseConnectionCtx",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetClientPrefetchThreads() {
	_jsii_.InvokeVoid(
		t,
		"resetClientPrefetchThreads",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetClientResultChunkSize() {
	_jsii_.InvokeVoid(
		t,
		"resetClientResultChunkSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetClientResultColumnCaseInsensitive() {
	_jsii_.InvokeVoid(
		t,
		"resetClientResultColumnCaseInsensitive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetClientSessionKeepAlive() {
	_jsii_.InvokeVoid(
		t,
		"resetClientSessionKeepAlive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetClientSessionKeepAliveHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		t,
		"resetClientSessionKeepAliveHeartbeatFrequency",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetClientTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		t,
		"resetClientTimestampTypeMapping",
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

func (t *jsiiProxy_Task) ResetConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetDateInputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetDateInputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetDateOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetDateOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetEnableUnloadPhysicalTypeOptimization() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableUnloadPhysicalTypeOptimization",
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

func (t *jsiiProxy_Task) ResetErrorOnNondeterministicMerge() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorOnNondeterministicMerge",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetErrorOnNondeterministicUpdate() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorOnNondeterministicUpdate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetFinalize() {
	_jsii_.InvokeVoid(
		t,
		"resetFinalize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetGeographyOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetGeographyOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetGeometryOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetGeometryOutputFormat",
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

func (t *jsiiProxy_Task) ResetJdbcTreatTimestampNtzAsUtc() {
	_jsii_.InvokeVoid(
		t,
		"resetJdbcTreatTimestampNtzAsUtc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetJdbcUseSessionTimezone() {
	_jsii_.InvokeVoid(
		t,
		"resetJdbcUseSessionTimezone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetJsonIndent() {
	_jsii_.InvokeVoid(
		t,
		"resetJsonIndent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetLockTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetLockTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetLogLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetMultiStatementCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiStatementCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetNoorderSequenceAsDefault() {
	_jsii_.InvokeVoid(
		t,
		"resetNoorderSequenceAsDefault",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetOdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		t,
		"resetOdbcTreatDecimalAsInt",
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

func (t *jsiiProxy_Task) ResetQueryTag() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryTag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetQuotedIdentifiersIgnoreCase() {
	_jsii_.InvokeVoid(
		t,
		"resetQuotedIdentifiersIgnoreCase",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetRowsPerResultset() {
	_jsii_.InvokeVoid(
		t,
		"resetRowsPerResultset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetS3StageVpceDnsName() {
	_jsii_.InvokeVoid(
		t,
		"resetS3StageVpceDnsName",
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

func (t *jsiiProxy_Task) ResetSearchPath() {
	_jsii_.InvokeVoid(
		t,
		"resetSearchPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetStatementQueuedTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetStatementQueuedTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetStatementTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetStatementTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetStrictJsonOutput() {
	_jsii_.InvokeVoid(
		t,
		"resetStrictJsonOutput",
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

func (t *jsiiProxy_Task) ResetTaskAutoRetryAttempts() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskAutoRetryAttempts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimeInputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeInputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimeOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimestampDayIsAlways24H() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampDayIsAlways24H",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimestampInputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampInputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimestampLtzOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampLtzOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimestampNtzOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampNtzOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimestampOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampTypeMapping",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimestampTzOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampTzOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTimezone() {
	_jsii_.InvokeVoid(
		t,
		"resetTimezone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTransactionAbortOnError() {
	_jsii_.InvokeVoid(
		t,
		"resetTransactionAbortOnError",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTransactionDefaultIsolationLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetTransactionDefaultIsolationLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetTwoDigitCenturyStart() {
	_jsii_.InvokeVoid(
		t,
		"resetTwoDigitCenturyStart",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetUnsupportedDdlAction() {
	_jsii_.InvokeVoid(
		t,
		"resetUnsupportedDdlAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetUseCachedResult() {
	_jsii_.InvokeVoid(
		t,
		"resetUseCachedResult",
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

func (t *jsiiProxy_Task) ResetUserTaskMinimumTriggerIntervalInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetUserTaskMinimumTriggerIntervalInSeconds",
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

func (t *jsiiProxy_Task) ResetWeekOfYearPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetWeekOfYearPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) ResetWeekStart() {
	_jsii_.InvokeVoid(
		t,
		"resetWeekStart",
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

