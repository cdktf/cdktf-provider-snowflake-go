// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflaketasks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v13/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v13/datasnowflaketasks/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeTasksTasksParametersOutputReference interface {
	cdktf.ComplexObject
	AbortDetachedQuery() DataSnowflakeTasksTasksParametersAbortDetachedQueryList
	Autocommit() DataSnowflakeTasksTasksParametersAutocommitList
	BinaryInputFormat() DataSnowflakeTasksTasksParametersBinaryInputFormatList
	BinaryOutputFormat() DataSnowflakeTasksTasksParametersBinaryOutputFormatList
	ClientMemoryLimit() DataSnowflakeTasksTasksParametersClientMemoryLimitList
	ClientMetadataRequestUseConnectionCtx() DataSnowflakeTasksTasksParametersClientMetadataRequestUseConnectionCtxList
	ClientPrefetchThreads() DataSnowflakeTasksTasksParametersClientPrefetchThreadsList
	ClientResultChunkSize() DataSnowflakeTasksTasksParametersClientResultChunkSizeList
	ClientResultColumnCaseInsensitive() DataSnowflakeTasksTasksParametersClientResultColumnCaseInsensitiveList
	ClientSessionKeepAlive() DataSnowflakeTasksTasksParametersClientSessionKeepAliveList
	ClientSessionKeepAliveHeartbeatFrequency() DataSnowflakeTasksTasksParametersClientSessionKeepAliveHeartbeatFrequencyList
	ClientTimestampTypeMapping() DataSnowflakeTasksTasksParametersClientTimestampTypeMappingList
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
	DateInputFormat() DataSnowflakeTasksTasksParametersDateInputFormatList
	DateOutputFormat() DataSnowflakeTasksTasksParametersDateOutputFormatList
	EnableUnloadPhysicalTypeOptimization() DataSnowflakeTasksTasksParametersEnableUnloadPhysicalTypeOptimizationList
	ErrorOnNondeterministicMerge() DataSnowflakeTasksTasksParametersErrorOnNondeterministicMergeList
	ErrorOnNondeterministicUpdate() DataSnowflakeTasksTasksParametersErrorOnNondeterministicUpdateList
	// Experimental.
	Fqn() *string
	GeographyOutputFormat() DataSnowflakeTasksTasksParametersGeographyOutputFormatList
	GeometryOutputFormat() DataSnowflakeTasksTasksParametersGeometryOutputFormatList
	InternalValue() *DataSnowflakeTasksTasksParameters
	SetInternalValue(val *DataSnowflakeTasksTasksParameters)
	JdbcTreatTimestampNtzAsUtc() DataSnowflakeTasksTasksParametersJdbcTreatTimestampNtzAsUtcList
	JdbcUseSessionTimezone() DataSnowflakeTasksTasksParametersJdbcUseSessionTimezoneList
	JsonIndent() DataSnowflakeTasksTasksParametersJsonIndentList
	LockTimeout() DataSnowflakeTasksTasksParametersLockTimeoutList
	LogLevel() DataSnowflakeTasksTasksParametersLogLevelList
	MultiStatementCount() DataSnowflakeTasksTasksParametersMultiStatementCountList
	NoorderSequenceAsDefault() DataSnowflakeTasksTasksParametersNoorderSequenceAsDefaultList
	OdbcTreatDecimalAsInt() DataSnowflakeTasksTasksParametersOdbcTreatDecimalAsIntList
	QueryTag() DataSnowflakeTasksTasksParametersQueryTagList
	QuotedIdentifiersIgnoreCase() DataSnowflakeTasksTasksParametersQuotedIdentifiersIgnoreCaseList
	RowsPerResultset() DataSnowflakeTasksTasksParametersRowsPerResultsetList
	S3StageVpceDnsName() DataSnowflakeTasksTasksParametersS3StageVpceDnsNameList
	SearchPath() DataSnowflakeTasksTasksParametersSearchPathList
	StatementQueuedTimeoutInSeconds() DataSnowflakeTasksTasksParametersStatementQueuedTimeoutInSecondsList
	StatementTimeoutInSeconds() DataSnowflakeTasksTasksParametersStatementTimeoutInSecondsList
	StrictJsonOutput() DataSnowflakeTasksTasksParametersStrictJsonOutputList
	SuspendTaskAfterNumFailures() DataSnowflakeTasksTasksParametersSuspendTaskAfterNumFailuresList
	TaskAutoRetryAttempts() DataSnowflakeTasksTasksParametersTaskAutoRetryAttemptsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeInputFormat() DataSnowflakeTasksTasksParametersTimeInputFormatList
	TimeOutputFormat() DataSnowflakeTasksTasksParametersTimeOutputFormatList
	TimestampDayIsAlways24H() DataSnowflakeTasksTasksParametersTimestampDayIsAlways24HList
	TimestampInputFormat() DataSnowflakeTasksTasksParametersTimestampInputFormatList
	TimestampLtzOutputFormat() DataSnowflakeTasksTasksParametersTimestampLtzOutputFormatList
	TimestampNtzOutputFormat() DataSnowflakeTasksTasksParametersTimestampNtzOutputFormatList
	TimestampOutputFormat() DataSnowflakeTasksTasksParametersTimestampOutputFormatList
	TimestampTypeMapping() DataSnowflakeTasksTasksParametersTimestampTypeMappingList
	TimestampTzOutputFormat() DataSnowflakeTasksTasksParametersTimestampTzOutputFormatList
	Timezone() DataSnowflakeTasksTasksParametersTimezoneList
	TraceLevel() DataSnowflakeTasksTasksParametersTraceLevelList
	TransactionAbortOnError() DataSnowflakeTasksTasksParametersTransactionAbortOnErrorList
	TransactionDefaultIsolationLevel() DataSnowflakeTasksTasksParametersTransactionDefaultIsolationLevelList
	TwoDigitCenturyStart() DataSnowflakeTasksTasksParametersTwoDigitCenturyStartList
	UnsupportedDdlAction() DataSnowflakeTasksTasksParametersUnsupportedDdlActionList
	UseCachedResult() DataSnowflakeTasksTasksParametersUseCachedResultList
	UserTaskManagedInitialWarehouseSize() DataSnowflakeTasksTasksParametersUserTaskManagedInitialWarehouseSizeList
	UserTaskMinimumTriggerIntervalInSeconds() DataSnowflakeTasksTasksParametersUserTaskMinimumTriggerIntervalInSecondsList
	UserTaskTimeoutMs() DataSnowflakeTasksTasksParametersUserTaskTimeoutMsList
	WeekOfYearPolicy() DataSnowflakeTasksTasksParametersWeekOfYearPolicyList
	WeekStart() DataSnowflakeTasksTasksParametersWeekStartList
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

// The jsii proxy struct for DataSnowflakeTasksTasksParametersOutputReference
type jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) AbortDetachedQuery() DataSnowflakeTasksTasksParametersAbortDetachedQueryList {
	var returns DataSnowflakeTasksTasksParametersAbortDetachedQueryList
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) Autocommit() DataSnowflakeTasksTasksParametersAutocommitList {
	var returns DataSnowflakeTasksTasksParametersAutocommitList
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) BinaryInputFormat() DataSnowflakeTasksTasksParametersBinaryInputFormatList {
	var returns DataSnowflakeTasksTasksParametersBinaryInputFormatList
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) BinaryOutputFormat() DataSnowflakeTasksTasksParametersBinaryOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersBinaryOutputFormatList
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ClientMemoryLimit() DataSnowflakeTasksTasksParametersClientMemoryLimitList {
	var returns DataSnowflakeTasksTasksParametersClientMemoryLimitList
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ClientMetadataRequestUseConnectionCtx() DataSnowflakeTasksTasksParametersClientMetadataRequestUseConnectionCtxList {
	var returns DataSnowflakeTasksTasksParametersClientMetadataRequestUseConnectionCtxList
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ClientPrefetchThreads() DataSnowflakeTasksTasksParametersClientPrefetchThreadsList {
	var returns DataSnowflakeTasksTasksParametersClientPrefetchThreadsList
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ClientResultChunkSize() DataSnowflakeTasksTasksParametersClientResultChunkSizeList {
	var returns DataSnowflakeTasksTasksParametersClientResultChunkSizeList
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ClientResultColumnCaseInsensitive() DataSnowflakeTasksTasksParametersClientResultColumnCaseInsensitiveList {
	var returns DataSnowflakeTasksTasksParametersClientResultColumnCaseInsensitiveList
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ClientSessionKeepAlive() DataSnowflakeTasksTasksParametersClientSessionKeepAliveList {
	var returns DataSnowflakeTasksTasksParametersClientSessionKeepAliveList
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ClientSessionKeepAliveHeartbeatFrequency() DataSnowflakeTasksTasksParametersClientSessionKeepAliveHeartbeatFrequencyList {
	var returns DataSnowflakeTasksTasksParametersClientSessionKeepAliveHeartbeatFrequencyList
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ClientTimestampTypeMapping() DataSnowflakeTasksTasksParametersClientTimestampTypeMappingList {
	var returns DataSnowflakeTasksTasksParametersClientTimestampTypeMappingList
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) DateInputFormat() DataSnowflakeTasksTasksParametersDateInputFormatList {
	var returns DataSnowflakeTasksTasksParametersDateInputFormatList
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) DateOutputFormat() DataSnowflakeTasksTasksParametersDateOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersDateOutputFormatList
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) EnableUnloadPhysicalTypeOptimization() DataSnowflakeTasksTasksParametersEnableUnloadPhysicalTypeOptimizationList {
	var returns DataSnowflakeTasksTasksParametersEnableUnloadPhysicalTypeOptimizationList
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ErrorOnNondeterministicMerge() DataSnowflakeTasksTasksParametersErrorOnNondeterministicMergeList {
	var returns DataSnowflakeTasksTasksParametersErrorOnNondeterministicMergeList
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ErrorOnNondeterministicUpdate() DataSnowflakeTasksTasksParametersErrorOnNondeterministicUpdateList {
	var returns DataSnowflakeTasksTasksParametersErrorOnNondeterministicUpdateList
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GeographyOutputFormat() DataSnowflakeTasksTasksParametersGeographyOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersGeographyOutputFormatList
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GeometryOutputFormat() DataSnowflakeTasksTasksParametersGeometryOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersGeometryOutputFormatList
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) InternalValue() *DataSnowflakeTasksTasksParameters {
	var returns *DataSnowflakeTasksTasksParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) JdbcTreatTimestampNtzAsUtc() DataSnowflakeTasksTasksParametersJdbcTreatTimestampNtzAsUtcList {
	var returns DataSnowflakeTasksTasksParametersJdbcTreatTimestampNtzAsUtcList
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) JdbcUseSessionTimezone() DataSnowflakeTasksTasksParametersJdbcUseSessionTimezoneList {
	var returns DataSnowflakeTasksTasksParametersJdbcUseSessionTimezoneList
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) JsonIndent() DataSnowflakeTasksTasksParametersJsonIndentList {
	var returns DataSnowflakeTasksTasksParametersJsonIndentList
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) LockTimeout() DataSnowflakeTasksTasksParametersLockTimeoutList {
	var returns DataSnowflakeTasksTasksParametersLockTimeoutList
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) LogLevel() DataSnowflakeTasksTasksParametersLogLevelList {
	var returns DataSnowflakeTasksTasksParametersLogLevelList
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) MultiStatementCount() DataSnowflakeTasksTasksParametersMultiStatementCountList {
	var returns DataSnowflakeTasksTasksParametersMultiStatementCountList
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) NoorderSequenceAsDefault() DataSnowflakeTasksTasksParametersNoorderSequenceAsDefaultList {
	var returns DataSnowflakeTasksTasksParametersNoorderSequenceAsDefaultList
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) OdbcTreatDecimalAsInt() DataSnowflakeTasksTasksParametersOdbcTreatDecimalAsIntList {
	var returns DataSnowflakeTasksTasksParametersOdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) QueryTag() DataSnowflakeTasksTasksParametersQueryTagList {
	var returns DataSnowflakeTasksTasksParametersQueryTagList
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) QuotedIdentifiersIgnoreCase() DataSnowflakeTasksTasksParametersQuotedIdentifiersIgnoreCaseList {
	var returns DataSnowflakeTasksTasksParametersQuotedIdentifiersIgnoreCaseList
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) RowsPerResultset() DataSnowflakeTasksTasksParametersRowsPerResultsetList {
	var returns DataSnowflakeTasksTasksParametersRowsPerResultsetList
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) S3StageVpceDnsName() DataSnowflakeTasksTasksParametersS3StageVpceDnsNameList {
	var returns DataSnowflakeTasksTasksParametersS3StageVpceDnsNameList
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) SearchPath() DataSnowflakeTasksTasksParametersSearchPathList {
	var returns DataSnowflakeTasksTasksParametersSearchPathList
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) StatementQueuedTimeoutInSeconds() DataSnowflakeTasksTasksParametersStatementQueuedTimeoutInSecondsList {
	var returns DataSnowflakeTasksTasksParametersStatementQueuedTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) StatementTimeoutInSeconds() DataSnowflakeTasksTasksParametersStatementTimeoutInSecondsList {
	var returns DataSnowflakeTasksTasksParametersStatementTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) StrictJsonOutput() DataSnowflakeTasksTasksParametersStrictJsonOutputList {
	var returns DataSnowflakeTasksTasksParametersStrictJsonOutputList
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) SuspendTaskAfterNumFailures() DataSnowflakeTasksTasksParametersSuspendTaskAfterNumFailuresList {
	var returns DataSnowflakeTasksTasksParametersSuspendTaskAfterNumFailuresList
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TaskAutoRetryAttempts() DataSnowflakeTasksTasksParametersTaskAutoRetryAttemptsList {
	var returns DataSnowflakeTasksTasksParametersTaskAutoRetryAttemptsList
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimeInputFormat() DataSnowflakeTasksTasksParametersTimeInputFormatList {
	var returns DataSnowflakeTasksTasksParametersTimeInputFormatList
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimeOutputFormat() DataSnowflakeTasksTasksParametersTimeOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersTimeOutputFormatList
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimestampDayIsAlways24H() DataSnowflakeTasksTasksParametersTimestampDayIsAlways24HList {
	var returns DataSnowflakeTasksTasksParametersTimestampDayIsAlways24HList
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimestampInputFormat() DataSnowflakeTasksTasksParametersTimestampInputFormatList {
	var returns DataSnowflakeTasksTasksParametersTimestampInputFormatList
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimestampLtzOutputFormat() DataSnowflakeTasksTasksParametersTimestampLtzOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersTimestampLtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimestampNtzOutputFormat() DataSnowflakeTasksTasksParametersTimestampNtzOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersTimestampNtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimestampOutputFormat() DataSnowflakeTasksTasksParametersTimestampOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersTimestampOutputFormatList
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimestampTypeMapping() DataSnowflakeTasksTasksParametersTimestampTypeMappingList {
	var returns DataSnowflakeTasksTasksParametersTimestampTypeMappingList
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TimestampTzOutputFormat() DataSnowflakeTasksTasksParametersTimestampTzOutputFormatList {
	var returns DataSnowflakeTasksTasksParametersTimestampTzOutputFormatList
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) Timezone() DataSnowflakeTasksTasksParametersTimezoneList {
	var returns DataSnowflakeTasksTasksParametersTimezoneList
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TraceLevel() DataSnowflakeTasksTasksParametersTraceLevelList {
	var returns DataSnowflakeTasksTasksParametersTraceLevelList
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TransactionAbortOnError() DataSnowflakeTasksTasksParametersTransactionAbortOnErrorList {
	var returns DataSnowflakeTasksTasksParametersTransactionAbortOnErrorList
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TransactionDefaultIsolationLevel() DataSnowflakeTasksTasksParametersTransactionDefaultIsolationLevelList {
	var returns DataSnowflakeTasksTasksParametersTransactionDefaultIsolationLevelList
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) TwoDigitCenturyStart() DataSnowflakeTasksTasksParametersTwoDigitCenturyStartList {
	var returns DataSnowflakeTasksTasksParametersTwoDigitCenturyStartList
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) UnsupportedDdlAction() DataSnowflakeTasksTasksParametersUnsupportedDdlActionList {
	var returns DataSnowflakeTasksTasksParametersUnsupportedDdlActionList
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) UseCachedResult() DataSnowflakeTasksTasksParametersUseCachedResultList {
	var returns DataSnowflakeTasksTasksParametersUseCachedResultList
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) UserTaskManagedInitialWarehouseSize() DataSnowflakeTasksTasksParametersUserTaskManagedInitialWarehouseSizeList {
	var returns DataSnowflakeTasksTasksParametersUserTaskManagedInitialWarehouseSizeList
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) UserTaskMinimumTriggerIntervalInSeconds() DataSnowflakeTasksTasksParametersUserTaskMinimumTriggerIntervalInSecondsList {
	var returns DataSnowflakeTasksTasksParametersUserTaskMinimumTriggerIntervalInSecondsList
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) UserTaskTimeoutMs() DataSnowflakeTasksTasksParametersUserTaskTimeoutMsList {
	var returns DataSnowflakeTasksTasksParametersUserTaskTimeoutMsList
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) WeekOfYearPolicy() DataSnowflakeTasksTasksParametersWeekOfYearPolicyList {
	var returns DataSnowflakeTasksTasksParametersWeekOfYearPolicyList
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) WeekStart() DataSnowflakeTasksTasksParametersWeekStartList {
	var returns DataSnowflakeTasksTasksParametersWeekStartList
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}


func NewDataSnowflakeTasksTasksParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeTasksTasksParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeTasksTasksParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeTasks.DataSnowflakeTasksTasksParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeTasksTasksParametersOutputReference_Override(d DataSnowflakeTasksTasksParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeTasks.DataSnowflakeTasksTasksParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference)SetInternalValue(val *DataSnowflakeTasksTasksParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeTasksTasksParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

