// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeusers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/datasnowflakeusers/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeUsersUsersParametersOutputReference interface {
	cdktf.ComplexObject
	AbortDetachedQuery() DataSnowflakeUsersUsersParametersAbortDetachedQueryList
	Autocommit() DataSnowflakeUsersUsersParametersAutocommitList
	BinaryInputFormat() DataSnowflakeUsersUsersParametersBinaryInputFormatList
	BinaryOutputFormat() DataSnowflakeUsersUsersParametersBinaryOutputFormatList
	ClientMemoryLimit() DataSnowflakeUsersUsersParametersClientMemoryLimitList
	ClientMetadataRequestUseConnectionCtx() DataSnowflakeUsersUsersParametersClientMetadataRequestUseConnectionCtxList
	ClientPrefetchThreads() DataSnowflakeUsersUsersParametersClientPrefetchThreadsList
	ClientResultChunkSize() DataSnowflakeUsersUsersParametersClientResultChunkSizeList
	ClientResultColumnCaseInsensitive() DataSnowflakeUsersUsersParametersClientResultColumnCaseInsensitiveList
	ClientSessionKeepAlive() DataSnowflakeUsersUsersParametersClientSessionKeepAliveList
	ClientSessionKeepAliveHeartbeatFrequency() DataSnowflakeUsersUsersParametersClientSessionKeepAliveHeartbeatFrequencyList
	ClientTimestampTypeMapping() DataSnowflakeUsersUsersParametersClientTimestampTypeMappingList
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
	DateInputFormat() DataSnowflakeUsersUsersParametersDateInputFormatList
	DateOutputFormat() DataSnowflakeUsersUsersParametersDateOutputFormatList
	EnableUnloadPhysicalTypeOptimization() DataSnowflakeUsersUsersParametersEnableUnloadPhysicalTypeOptimizationList
	EnableUnredactedQuerySyntaxError() DataSnowflakeUsersUsersParametersEnableUnredactedQuerySyntaxErrorList
	ErrorOnNondeterministicMerge() DataSnowflakeUsersUsersParametersErrorOnNondeterministicMergeList
	ErrorOnNondeterministicUpdate() DataSnowflakeUsersUsersParametersErrorOnNondeterministicUpdateList
	// Experimental.
	Fqn() *string
	GeographyOutputFormat() DataSnowflakeUsersUsersParametersGeographyOutputFormatList
	GeometryOutputFormat() DataSnowflakeUsersUsersParametersGeometryOutputFormatList
	InternalValue() *DataSnowflakeUsersUsersParameters
	SetInternalValue(val *DataSnowflakeUsersUsersParameters)
	JdbcTreatDecimalAsInt() DataSnowflakeUsersUsersParametersJdbcTreatDecimalAsIntList
	JdbcTreatTimestampNtzAsUtc() DataSnowflakeUsersUsersParametersJdbcTreatTimestampNtzAsUtcList
	JdbcUseSessionTimezone() DataSnowflakeUsersUsersParametersJdbcUseSessionTimezoneList
	JsonIndent() DataSnowflakeUsersUsersParametersJsonIndentList
	LockTimeout() DataSnowflakeUsersUsersParametersLockTimeoutList
	LogLevel() DataSnowflakeUsersUsersParametersLogLevelList
	MultiStatementCount() DataSnowflakeUsersUsersParametersMultiStatementCountList
	NetworkPolicy() DataSnowflakeUsersUsersParametersNetworkPolicyList
	NoorderSequenceAsDefault() DataSnowflakeUsersUsersParametersNoorderSequenceAsDefaultList
	OdbcTreatDecimalAsInt() DataSnowflakeUsersUsersParametersOdbcTreatDecimalAsIntList
	PreventUnloadToInternalStages() DataSnowflakeUsersUsersParametersPreventUnloadToInternalStagesList
	QueryTag() DataSnowflakeUsersUsersParametersQueryTagList
	QuotedIdentifiersIgnoreCase() DataSnowflakeUsersUsersParametersQuotedIdentifiersIgnoreCaseList
	RowsPerResultset() DataSnowflakeUsersUsersParametersRowsPerResultsetList
	S3StageVpceDnsName() DataSnowflakeUsersUsersParametersS3StageVpceDnsNameList
	SearchPath() DataSnowflakeUsersUsersParametersSearchPathList
	SimulatedDataSharingConsumer() DataSnowflakeUsersUsersParametersSimulatedDataSharingConsumerList
	StatementQueuedTimeoutInSeconds() DataSnowflakeUsersUsersParametersStatementQueuedTimeoutInSecondsList
	StatementTimeoutInSeconds() DataSnowflakeUsersUsersParametersStatementTimeoutInSecondsList
	StrictJsonOutput() DataSnowflakeUsersUsersParametersStrictJsonOutputList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeInputFormat() DataSnowflakeUsersUsersParametersTimeInputFormatList
	TimeOutputFormat() DataSnowflakeUsersUsersParametersTimeOutputFormatList
	TimestampDayIsAlways24H() DataSnowflakeUsersUsersParametersTimestampDayIsAlways24HList
	TimestampInputFormat() DataSnowflakeUsersUsersParametersTimestampInputFormatList
	TimestampLtzOutputFormat() DataSnowflakeUsersUsersParametersTimestampLtzOutputFormatList
	TimestampNtzOutputFormat() DataSnowflakeUsersUsersParametersTimestampNtzOutputFormatList
	TimestampOutputFormat() DataSnowflakeUsersUsersParametersTimestampOutputFormatList
	TimestampTypeMapping() DataSnowflakeUsersUsersParametersTimestampTypeMappingList
	TimestampTzOutputFormat() DataSnowflakeUsersUsersParametersTimestampTzOutputFormatList
	Timezone() DataSnowflakeUsersUsersParametersTimezoneList
	TraceLevel() DataSnowflakeUsersUsersParametersTraceLevelList
	TransactionAbortOnError() DataSnowflakeUsersUsersParametersTransactionAbortOnErrorList
	TransactionDefaultIsolationLevel() DataSnowflakeUsersUsersParametersTransactionDefaultIsolationLevelList
	TwoDigitCenturyStart() DataSnowflakeUsersUsersParametersTwoDigitCenturyStartList
	UnsupportedDdlAction() DataSnowflakeUsersUsersParametersUnsupportedDdlActionList
	UseCachedResult() DataSnowflakeUsersUsersParametersUseCachedResultList
	WeekOfYearPolicy() DataSnowflakeUsersUsersParametersWeekOfYearPolicyList
	WeekStart() DataSnowflakeUsersUsersParametersWeekStartList
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

// The jsii proxy struct for DataSnowflakeUsersUsersParametersOutputReference
type jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) AbortDetachedQuery() DataSnowflakeUsersUsersParametersAbortDetachedQueryList {
	var returns DataSnowflakeUsersUsersParametersAbortDetachedQueryList
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) Autocommit() DataSnowflakeUsersUsersParametersAutocommitList {
	var returns DataSnowflakeUsersUsersParametersAutocommitList
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) BinaryInputFormat() DataSnowflakeUsersUsersParametersBinaryInputFormatList {
	var returns DataSnowflakeUsersUsersParametersBinaryInputFormatList
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) BinaryOutputFormat() DataSnowflakeUsersUsersParametersBinaryOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersBinaryOutputFormatList
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ClientMemoryLimit() DataSnowflakeUsersUsersParametersClientMemoryLimitList {
	var returns DataSnowflakeUsersUsersParametersClientMemoryLimitList
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ClientMetadataRequestUseConnectionCtx() DataSnowflakeUsersUsersParametersClientMetadataRequestUseConnectionCtxList {
	var returns DataSnowflakeUsersUsersParametersClientMetadataRequestUseConnectionCtxList
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ClientPrefetchThreads() DataSnowflakeUsersUsersParametersClientPrefetchThreadsList {
	var returns DataSnowflakeUsersUsersParametersClientPrefetchThreadsList
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ClientResultChunkSize() DataSnowflakeUsersUsersParametersClientResultChunkSizeList {
	var returns DataSnowflakeUsersUsersParametersClientResultChunkSizeList
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ClientResultColumnCaseInsensitive() DataSnowflakeUsersUsersParametersClientResultColumnCaseInsensitiveList {
	var returns DataSnowflakeUsersUsersParametersClientResultColumnCaseInsensitiveList
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ClientSessionKeepAlive() DataSnowflakeUsersUsersParametersClientSessionKeepAliveList {
	var returns DataSnowflakeUsersUsersParametersClientSessionKeepAliveList
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ClientSessionKeepAliveHeartbeatFrequency() DataSnowflakeUsersUsersParametersClientSessionKeepAliveHeartbeatFrequencyList {
	var returns DataSnowflakeUsersUsersParametersClientSessionKeepAliveHeartbeatFrequencyList
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ClientTimestampTypeMapping() DataSnowflakeUsersUsersParametersClientTimestampTypeMappingList {
	var returns DataSnowflakeUsersUsersParametersClientTimestampTypeMappingList
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) DateInputFormat() DataSnowflakeUsersUsersParametersDateInputFormatList {
	var returns DataSnowflakeUsersUsersParametersDateInputFormatList
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) DateOutputFormat() DataSnowflakeUsersUsersParametersDateOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersDateOutputFormatList
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) EnableUnloadPhysicalTypeOptimization() DataSnowflakeUsersUsersParametersEnableUnloadPhysicalTypeOptimizationList {
	var returns DataSnowflakeUsersUsersParametersEnableUnloadPhysicalTypeOptimizationList
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) EnableUnredactedQuerySyntaxError() DataSnowflakeUsersUsersParametersEnableUnredactedQuerySyntaxErrorList {
	var returns DataSnowflakeUsersUsersParametersEnableUnredactedQuerySyntaxErrorList
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ErrorOnNondeterministicMerge() DataSnowflakeUsersUsersParametersErrorOnNondeterministicMergeList {
	var returns DataSnowflakeUsersUsersParametersErrorOnNondeterministicMergeList
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ErrorOnNondeterministicUpdate() DataSnowflakeUsersUsersParametersErrorOnNondeterministicUpdateList {
	var returns DataSnowflakeUsersUsersParametersErrorOnNondeterministicUpdateList
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GeographyOutputFormat() DataSnowflakeUsersUsersParametersGeographyOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersGeographyOutputFormatList
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GeometryOutputFormat() DataSnowflakeUsersUsersParametersGeometryOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersGeometryOutputFormatList
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) InternalValue() *DataSnowflakeUsersUsersParameters {
	var returns *DataSnowflakeUsersUsersParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) JdbcTreatDecimalAsInt() DataSnowflakeUsersUsersParametersJdbcTreatDecimalAsIntList {
	var returns DataSnowflakeUsersUsersParametersJdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) JdbcTreatTimestampNtzAsUtc() DataSnowflakeUsersUsersParametersJdbcTreatTimestampNtzAsUtcList {
	var returns DataSnowflakeUsersUsersParametersJdbcTreatTimestampNtzAsUtcList
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) JdbcUseSessionTimezone() DataSnowflakeUsersUsersParametersJdbcUseSessionTimezoneList {
	var returns DataSnowflakeUsersUsersParametersJdbcUseSessionTimezoneList
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) JsonIndent() DataSnowflakeUsersUsersParametersJsonIndentList {
	var returns DataSnowflakeUsersUsersParametersJsonIndentList
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) LockTimeout() DataSnowflakeUsersUsersParametersLockTimeoutList {
	var returns DataSnowflakeUsersUsersParametersLockTimeoutList
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) LogLevel() DataSnowflakeUsersUsersParametersLogLevelList {
	var returns DataSnowflakeUsersUsersParametersLogLevelList
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) MultiStatementCount() DataSnowflakeUsersUsersParametersMultiStatementCountList {
	var returns DataSnowflakeUsersUsersParametersMultiStatementCountList
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) NetworkPolicy() DataSnowflakeUsersUsersParametersNetworkPolicyList {
	var returns DataSnowflakeUsersUsersParametersNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) NoorderSequenceAsDefault() DataSnowflakeUsersUsersParametersNoorderSequenceAsDefaultList {
	var returns DataSnowflakeUsersUsersParametersNoorderSequenceAsDefaultList
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) OdbcTreatDecimalAsInt() DataSnowflakeUsersUsersParametersOdbcTreatDecimalAsIntList {
	var returns DataSnowflakeUsersUsersParametersOdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) PreventUnloadToInternalStages() DataSnowflakeUsersUsersParametersPreventUnloadToInternalStagesList {
	var returns DataSnowflakeUsersUsersParametersPreventUnloadToInternalStagesList
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) QueryTag() DataSnowflakeUsersUsersParametersQueryTagList {
	var returns DataSnowflakeUsersUsersParametersQueryTagList
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) QuotedIdentifiersIgnoreCase() DataSnowflakeUsersUsersParametersQuotedIdentifiersIgnoreCaseList {
	var returns DataSnowflakeUsersUsersParametersQuotedIdentifiersIgnoreCaseList
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) RowsPerResultset() DataSnowflakeUsersUsersParametersRowsPerResultsetList {
	var returns DataSnowflakeUsersUsersParametersRowsPerResultsetList
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) S3StageVpceDnsName() DataSnowflakeUsersUsersParametersS3StageVpceDnsNameList {
	var returns DataSnowflakeUsersUsersParametersS3StageVpceDnsNameList
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) SearchPath() DataSnowflakeUsersUsersParametersSearchPathList {
	var returns DataSnowflakeUsersUsersParametersSearchPathList
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) SimulatedDataSharingConsumer() DataSnowflakeUsersUsersParametersSimulatedDataSharingConsumerList {
	var returns DataSnowflakeUsersUsersParametersSimulatedDataSharingConsumerList
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) StatementQueuedTimeoutInSeconds() DataSnowflakeUsersUsersParametersStatementQueuedTimeoutInSecondsList {
	var returns DataSnowflakeUsersUsersParametersStatementQueuedTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) StatementTimeoutInSeconds() DataSnowflakeUsersUsersParametersStatementTimeoutInSecondsList {
	var returns DataSnowflakeUsersUsersParametersStatementTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) StrictJsonOutput() DataSnowflakeUsersUsersParametersStrictJsonOutputList {
	var returns DataSnowflakeUsersUsersParametersStrictJsonOutputList
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimeInputFormat() DataSnowflakeUsersUsersParametersTimeInputFormatList {
	var returns DataSnowflakeUsersUsersParametersTimeInputFormatList
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimeOutputFormat() DataSnowflakeUsersUsersParametersTimeOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersTimeOutputFormatList
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimestampDayIsAlways24H() DataSnowflakeUsersUsersParametersTimestampDayIsAlways24HList {
	var returns DataSnowflakeUsersUsersParametersTimestampDayIsAlways24HList
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimestampInputFormat() DataSnowflakeUsersUsersParametersTimestampInputFormatList {
	var returns DataSnowflakeUsersUsersParametersTimestampInputFormatList
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimestampLtzOutputFormat() DataSnowflakeUsersUsersParametersTimestampLtzOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersTimestampLtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimestampNtzOutputFormat() DataSnowflakeUsersUsersParametersTimestampNtzOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersTimestampNtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimestampOutputFormat() DataSnowflakeUsersUsersParametersTimestampOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersTimestampOutputFormatList
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimestampTypeMapping() DataSnowflakeUsersUsersParametersTimestampTypeMappingList {
	var returns DataSnowflakeUsersUsersParametersTimestampTypeMappingList
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TimestampTzOutputFormat() DataSnowflakeUsersUsersParametersTimestampTzOutputFormatList {
	var returns DataSnowflakeUsersUsersParametersTimestampTzOutputFormatList
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) Timezone() DataSnowflakeUsersUsersParametersTimezoneList {
	var returns DataSnowflakeUsersUsersParametersTimezoneList
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TraceLevel() DataSnowflakeUsersUsersParametersTraceLevelList {
	var returns DataSnowflakeUsersUsersParametersTraceLevelList
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TransactionAbortOnError() DataSnowflakeUsersUsersParametersTransactionAbortOnErrorList {
	var returns DataSnowflakeUsersUsersParametersTransactionAbortOnErrorList
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TransactionDefaultIsolationLevel() DataSnowflakeUsersUsersParametersTransactionDefaultIsolationLevelList {
	var returns DataSnowflakeUsersUsersParametersTransactionDefaultIsolationLevelList
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) TwoDigitCenturyStart() DataSnowflakeUsersUsersParametersTwoDigitCenturyStartList {
	var returns DataSnowflakeUsersUsersParametersTwoDigitCenturyStartList
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) UnsupportedDdlAction() DataSnowflakeUsersUsersParametersUnsupportedDdlActionList {
	var returns DataSnowflakeUsersUsersParametersUnsupportedDdlActionList
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) UseCachedResult() DataSnowflakeUsersUsersParametersUseCachedResultList {
	var returns DataSnowflakeUsersUsersParametersUseCachedResultList
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) WeekOfYearPolicy() DataSnowflakeUsersUsersParametersWeekOfYearPolicyList {
	var returns DataSnowflakeUsersUsersParametersWeekOfYearPolicyList
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) WeekStart() DataSnowflakeUsersUsersParametersWeekStartList {
	var returns DataSnowflakeUsersUsersParametersWeekStartList
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}


func NewDataSnowflakeUsersUsersParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeUsersUsersParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeUsersUsersParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeUsers.DataSnowflakeUsersUsersParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeUsersUsersParametersOutputReference_Override(d DataSnowflakeUsersUsersParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeUsers.DataSnowflakeUsersUsersParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference)SetInternalValue(val *DataSnowflakeUsersUsersParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

