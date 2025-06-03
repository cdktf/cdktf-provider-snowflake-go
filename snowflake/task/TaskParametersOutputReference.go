// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package task

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/task/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TaskParametersOutputReference interface {
	cdktf.ComplexObject
	AbortDetachedQuery() TaskParametersAbortDetachedQueryList
	Autocommit() TaskParametersAutocommitList
	BinaryInputFormat() TaskParametersBinaryInputFormatList
	BinaryOutputFormat() TaskParametersBinaryOutputFormatList
	ClientMemoryLimit() TaskParametersClientMemoryLimitList
	ClientMetadataRequestUseConnectionCtx() TaskParametersClientMetadataRequestUseConnectionCtxList
	ClientPrefetchThreads() TaskParametersClientPrefetchThreadsList
	ClientResultChunkSize() TaskParametersClientResultChunkSizeList
	ClientResultColumnCaseInsensitive() TaskParametersClientResultColumnCaseInsensitiveList
	ClientSessionKeepAlive() TaskParametersClientSessionKeepAliveList
	ClientSessionKeepAliveHeartbeatFrequency() TaskParametersClientSessionKeepAliveHeartbeatFrequencyList
	ClientTimestampTypeMapping() TaskParametersClientTimestampTypeMappingList
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DateInputFormat() TaskParametersDateInputFormatList
	DateOutputFormat() TaskParametersDateOutputFormatList
	EnableUnloadPhysicalTypeOptimization() TaskParametersEnableUnloadPhysicalTypeOptimizationList
	ErrorOnNondeterministicMerge() TaskParametersErrorOnNondeterministicMergeList
	ErrorOnNondeterministicUpdate() TaskParametersErrorOnNondeterministicUpdateList
	// Experimental.
	Fqn() *string
	GeographyOutputFormat() TaskParametersGeographyOutputFormatList
	GeometryOutputFormat() TaskParametersGeometryOutputFormatList
	InternalValue() *TaskParameters
	SetInternalValue(val *TaskParameters)
	JdbcTreatTimestampNtzAsUtc() TaskParametersJdbcTreatTimestampNtzAsUtcList
	JdbcUseSessionTimezone() TaskParametersJdbcUseSessionTimezoneList
	JsonIndent() TaskParametersJsonIndentList
	LockTimeout() TaskParametersLockTimeoutList
	LogLevel() TaskParametersLogLevelList
	MultiStatementCount() TaskParametersMultiStatementCountList
	NoorderSequenceAsDefault() TaskParametersNoorderSequenceAsDefaultList
	OdbcTreatDecimalAsInt() TaskParametersOdbcTreatDecimalAsIntList
	QueryTag() TaskParametersQueryTagList
	QuotedIdentifiersIgnoreCase() TaskParametersQuotedIdentifiersIgnoreCaseList
	RowsPerResultset() TaskParametersRowsPerResultsetList
	S3StageVpceDnsName() TaskParametersS3StageVpceDnsNameList
	SearchPath() TaskParametersSearchPathList
	StatementQueuedTimeoutInSeconds() TaskParametersStatementQueuedTimeoutInSecondsList
	StatementTimeoutInSeconds() TaskParametersStatementTimeoutInSecondsList
	StrictJsonOutput() TaskParametersStrictJsonOutputList
	SuspendTaskAfterNumFailures() TaskParametersSuspendTaskAfterNumFailuresList
	TaskAutoRetryAttempts() TaskParametersTaskAutoRetryAttemptsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeInputFormat() TaskParametersTimeInputFormatList
	TimeOutputFormat() TaskParametersTimeOutputFormatList
	TimestampDayIsAlways24H() TaskParametersTimestampDayIsAlways24HList
	TimestampInputFormat() TaskParametersTimestampInputFormatList
	TimestampLtzOutputFormat() TaskParametersTimestampLtzOutputFormatList
	TimestampNtzOutputFormat() TaskParametersTimestampNtzOutputFormatList
	TimestampOutputFormat() TaskParametersTimestampOutputFormatList
	TimestampTypeMapping() TaskParametersTimestampTypeMappingList
	TimestampTzOutputFormat() TaskParametersTimestampTzOutputFormatList
	Timezone() TaskParametersTimezoneList
	TraceLevel() TaskParametersTraceLevelList
	TransactionAbortOnError() TaskParametersTransactionAbortOnErrorList
	TransactionDefaultIsolationLevel() TaskParametersTransactionDefaultIsolationLevelList
	TwoDigitCenturyStart() TaskParametersTwoDigitCenturyStartList
	UnsupportedDdlAction() TaskParametersUnsupportedDdlActionList
	UseCachedResult() TaskParametersUseCachedResultList
	UserTaskManagedInitialWarehouseSize() TaskParametersUserTaskManagedInitialWarehouseSizeList
	UserTaskMinimumTriggerIntervalInSeconds() TaskParametersUserTaskMinimumTriggerIntervalInSecondsList
	UserTaskTimeoutMs() TaskParametersUserTaskTimeoutMsList
	WeekOfYearPolicy() TaskParametersWeekOfYearPolicyList
	WeekStart() TaskParametersWeekStartList
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TaskParametersOutputReference
type jsiiProxy_TaskParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TaskParametersOutputReference) AbortDetachedQuery() TaskParametersAbortDetachedQueryList {
	var returns TaskParametersAbortDetachedQueryList
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) Autocommit() TaskParametersAutocommitList {
	var returns TaskParametersAutocommitList
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) BinaryInputFormat() TaskParametersBinaryInputFormatList {
	var returns TaskParametersBinaryInputFormatList
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) BinaryOutputFormat() TaskParametersBinaryOutputFormatList {
	var returns TaskParametersBinaryOutputFormatList
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ClientMemoryLimit() TaskParametersClientMemoryLimitList {
	var returns TaskParametersClientMemoryLimitList
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ClientMetadataRequestUseConnectionCtx() TaskParametersClientMetadataRequestUseConnectionCtxList {
	var returns TaskParametersClientMetadataRequestUseConnectionCtxList
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ClientPrefetchThreads() TaskParametersClientPrefetchThreadsList {
	var returns TaskParametersClientPrefetchThreadsList
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ClientResultChunkSize() TaskParametersClientResultChunkSizeList {
	var returns TaskParametersClientResultChunkSizeList
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ClientResultColumnCaseInsensitive() TaskParametersClientResultColumnCaseInsensitiveList {
	var returns TaskParametersClientResultColumnCaseInsensitiveList
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ClientSessionKeepAlive() TaskParametersClientSessionKeepAliveList {
	var returns TaskParametersClientSessionKeepAliveList
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ClientSessionKeepAliveHeartbeatFrequency() TaskParametersClientSessionKeepAliveHeartbeatFrequencyList {
	var returns TaskParametersClientSessionKeepAliveHeartbeatFrequencyList
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ClientTimestampTypeMapping() TaskParametersClientTimestampTypeMappingList {
	var returns TaskParametersClientTimestampTypeMappingList
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) DateInputFormat() TaskParametersDateInputFormatList {
	var returns TaskParametersDateInputFormatList
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) DateOutputFormat() TaskParametersDateOutputFormatList {
	var returns TaskParametersDateOutputFormatList
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) EnableUnloadPhysicalTypeOptimization() TaskParametersEnableUnloadPhysicalTypeOptimizationList {
	var returns TaskParametersEnableUnloadPhysicalTypeOptimizationList
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ErrorOnNondeterministicMerge() TaskParametersErrorOnNondeterministicMergeList {
	var returns TaskParametersErrorOnNondeterministicMergeList
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) ErrorOnNondeterministicUpdate() TaskParametersErrorOnNondeterministicUpdateList {
	var returns TaskParametersErrorOnNondeterministicUpdateList
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) GeographyOutputFormat() TaskParametersGeographyOutputFormatList {
	var returns TaskParametersGeographyOutputFormatList
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) GeometryOutputFormat() TaskParametersGeometryOutputFormatList {
	var returns TaskParametersGeometryOutputFormatList
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) InternalValue() *TaskParameters {
	var returns *TaskParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) JdbcTreatTimestampNtzAsUtc() TaskParametersJdbcTreatTimestampNtzAsUtcList {
	var returns TaskParametersJdbcTreatTimestampNtzAsUtcList
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) JdbcUseSessionTimezone() TaskParametersJdbcUseSessionTimezoneList {
	var returns TaskParametersJdbcUseSessionTimezoneList
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) JsonIndent() TaskParametersJsonIndentList {
	var returns TaskParametersJsonIndentList
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) LockTimeout() TaskParametersLockTimeoutList {
	var returns TaskParametersLockTimeoutList
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) LogLevel() TaskParametersLogLevelList {
	var returns TaskParametersLogLevelList
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) MultiStatementCount() TaskParametersMultiStatementCountList {
	var returns TaskParametersMultiStatementCountList
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) NoorderSequenceAsDefault() TaskParametersNoorderSequenceAsDefaultList {
	var returns TaskParametersNoorderSequenceAsDefaultList
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) OdbcTreatDecimalAsInt() TaskParametersOdbcTreatDecimalAsIntList {
	var returns TaskParametersOdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) QueryTag() TaskParametersQueryTagList {
	var returns TaskParametersQueryTagList
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) QuotedIdentifiersIgnoreCase() TaskParametersQuotedIdentifiersIgnoreCaseList {
	var returns TaskParametersQuotedIdentifiersIgnoreCaseList
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) RowsPerResultset() TaskParametersRowsPerResultsetList {
	var returns TaskParametersRowsPerResultsetList
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) S3StageVpceDnsName() TaskParametersS3StageVpceDnsNameList {
	var returns TaskParametersS3StageVpceDnsNameList
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) SearchPath() TaskParametersSearchPathList {
	var returns TaskParametersSearchPathList
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) StatementQueuedTimeoutInSeconds() TaskParametersStatementQueuedTimeoutInSecondsList {
	var returns TaskParametersStatementQueuedTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) StatementTimeoutInSeconds() TaskParametersStatementTimeoutInSecondsList {
	var returns TaskParametersStatementTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) StrictJsonOutput() TaskParametersStrictJsonOutputList {
	var returns TaskParametersStrictJsonOutputList
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) SuspendTaskAfterNumFailures() TaskParametersSuspendTaskAfterNumFailuresList {
	var returns TaskParametersSuspendTaskAfterNumFailuresList
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TaskAutoRetryAttempts() TaskParametersTaskAutoRetryAttemptsList {
	var returns TaskParametersTaskAutoRetryAttemptsList
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimeInputFormat() TaskParametersTimeInputFormatList {
	var returns TaskParametersTimeInputFormatList
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimeOutputFormat() TaskParametersTimeOutputFormatList {
	var returns TaskParametersTimeOutputFormatList
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimestampDayIsAlways24H() TaskParametersTimestampDayIsAlways24HList {
	var returns TaskParametersTimestampDayIsAlways24HList
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimestampInputFormat() TaskParametersTimestampInputFormatList {
	var returns TaskParametersTimestampInputFormatList
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimestampLtzOutputFormat() TaskParametersTimestampLtzOutputFormatList {
	var returns TaskParametersTimestampLtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimestampNtzOutputFormat() TaskParametersTimestampNtzOutputFormatList {
	var returns TaskParametersTimestampNtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimestampOutputFormat() TaskParametersTimestampOutputFormatList {
	var returns TaskParametersTimestampOutputFormatList
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimestampTypeMapping() TaskParametersTimestampTypeMappingList {
	var returns TaskParametersTimestampTypeMappingList
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TimestampTzOutputFormat() TaskParametersTimestampTzOutputFormatList {
	var returns TaskParametersTimestampTzOutputFormatList
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) Timezone() TaskParametersTimezoneList {
	var returns TaskParametersTimezoneList
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TraceLevel() TaskParametersTraceLevelList {
	var returns TaskParametersTraceLevelList
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TransactionAbortOnError() TaskParametersTransactionAbortOnErrorList {
	var returns TaskParametersTransactionAbortOnErrorList
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TransactionDefaultIsolationLevel() TaskParametersTransactionDefaultIsolationLevelList {
	var returns TaskParametersTransactionDefaultIsolationLevelList
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) TwoDigitCenturyStart() TaskParametersTwoDigitCenturyStartList {
	var returns TaskParametersTwoDigitCenturyStartList
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) UnsupportedDdlAction() TaskParametersUnsupportedDdlActionList {
	var returns TaskParametersUnsupportedDdlActionList
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) UseCachedResult() TaskParametersUseCachedResultList {
	var returns TaskParametersUseCachedResultList
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) UserTaskManagedInitialWarehouseSize() TaskParametersUserTaskManagedInitialWarehouseSizeList {
	var returns TaskParametersUserTaskManagedInitialWarehouseSizeList
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) UserTaskMinimumTriggerIntervalInSeconds() TaskParametersUserTaskMinimumTriggerIntervalInSecondsList {
	var returns TaskParametersUserTaskMinimumTriggerIntervalInSecondsList
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) UserTaskTimeoutMs() TaskParametersUserTaskTimeoutMsList {
	var returns TaskParametersUserTaskTimeoutMsList
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) WeekOfYearPolicy() TaskParametersWeekOfYearPolicyList {
	var returns TaskParametersWeekOfYearPolicyList
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskParametersOutputReference) WeekStart() TaskParametersWeekStartList {
	var returns TaskParametersWeekStartList
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}


func NewTaskParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TaskParametersOutputReference {
	_init_.Initialize()

	if err := validateNewTaskParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaskParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.task.TaskParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTaskParametersOutputReference_Override(t TaskParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.task.TaskParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TaskParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TaskParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TaskParametersOutputReference)SetInternalValue(val *TaskParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TaskParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TaskParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TaskParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TaskParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TaskParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TaskParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TaskParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TaskParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TaskParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TaskParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TaskParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TaskParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

