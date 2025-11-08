// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package legacyserviceuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/legacyserviceuser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LegacyServiceUserParametersOutputReference interface {
	cdktf.ComplexObject
	AbortDetachedQuery() LegacyServiceUserParametersAbortDetachedQueryList
	Autocommit() LegacyServiceUserParametersAutocommitList
	BinaryInputFormat() LegacyServiceUserParametersBinaryInputFormatList
	BinaryOutputFormat() LegacyServiceUserParametersBinaryOutputFormatList
	ClientMemoryLimit() LegacyServiceUserParametersClientMemoryLimitList
	ClientMetadataRequestUseConnectionCtx() LegacyServiceUserParametersClientMetadataRequestUseConnectionCtxList
	ClientPrefetchThreads() LegacyServiceUserParametersClientPrefetchThreadsList
	ClientResultChunkSize() LegacyServiceUserParametersClientResultChunkSizeList
	ClientResultColumnCaseInsensitive() LegacyServiceUserParametersClientResultColumnCaseInsensitiveList
	ClientSessionKeepAlive() LegacyServiceUserParametersClientSessionKeepAliveList
	ClientSessionKeepAliveHeartbeatFrequency() LegacyServiceUserParametersClientSessionKeepAliveHeartbeatFrequencyList
	ClientTimestampTypeMapping() LegacyServiceUserParametersClientTimestampTypeMappingList
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
	DateInputFormat() LegacyServiceUserParametersDateInputFormatList
	DateOutputFormat() LegacyServiceUserParametersDateOutputFormatList
	EnableUnloadPhysicalTypeOptimization() LegacyServiceUserParametersEnableUnloadPhysicalTypeOptimizationList
	EnableUnredactedQuerySyntaxError() LegacyServiceUserParametersEnableUnredactedQuerySyntaxErrorList
	ErrorOnNondeterministicMerge() LegacyServiceUserParametersErrorOnNondeterministicMergeList
	ErrorOnNondeterministicUpdate() LegacyServiceUserParametersErrorOnNondeterministicUpdateList
	// Experimental.
	Fqn() *string
	GeographyOutputFormat() LegacyServiceUserParametersGeographyOutputFormatList
	GeometryOutputFormat() LegacyServiceUserParametersGeometryOutputFormatList
	InternalValue() *LegacyServiceUserParameters
	SetInternalValue(val *LegacyServiceUserParameters)
	JdbcTreatDecimalAsInt() LegacyServiceUserParametersJdbcTreatDecimalAsIntList
	JdbcTreatTimestampNtzAsUtc() LegacyServiceUserParametersJdbcTreatTimestampNtzAsUtcList
	JdbcUseSessionTimezone() LegacyServiceUserParametersJdbcUseSessionTimezoneList
	JsonIndent() LegacyServiceUserParametersJsonIndentList
	LockTimeout() LegacyServiceUserParametersLockTimeoutList
	LogLevel() LegacyServiceUserParametersLogLevelList
	MultiStatementCount() LegacyServiceUserParametersMultiStatementCountList
	NetworkPolicy() LegacyServiceUserParametersNetworkPolicyList
	NoorderSequenceAsDefault() LegacyServiceUserParametersNoorderSequenceAsDefaultList
	OdbcTreatDecimalAsInt() LegacyServiceUserParametersOdbcTreatDecimalAsIntList
	PreventUnloadToInternalStages() LegacyServiceUserParametersPreventUnloadToInternalStagesList
	QueryTag() LegacyServiceUserParametersQueryTagList
	QuotedIdentifiersIgnoreCase() LegacyServiceUserParametersQuotedIdentifiersIgnoreCaseList
	RowsPerResultset() LegacyServiceUserParametersRowsPerResultsetList
	S3StageVpceDnsName() LegacyServiceUserParametersS3StageVpceDnsNameList
	SearchPath() LegacyServiceUserParametersSearchPathList
	SimulatedDataSharingConsumer() LegacyServiceUserParametersSimulatedDataSharingConsumerList
	StatementQueuedTimeoutInSeconds() LegacyServiceUserParametersStatementQueuedTimeoutInSecondsList
	StatementTimeoutInSeconds() LegacyServiceUserParametersStatementTimeoutInSecondsList
	StrictJsonOutput() LegacyServiceUserParametersStrictJsonOutputList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeInputFormat() LegacyServiceUserParametersTimeInputFormatList
	TimeOutputFormat() LegacyServiceUserParametersTimeOutputFormatList
	TimestampDayIsAlways24H() LegacyServiceUserParametersTimestampDayIsAlways24HList
	TimestampInputFormat() LegacyServiceUserParametersTimestampInputFormatList
	TimestampLtzOutputFormat() LegacyServiceUserParametersTimestampLtzOutputFormatList
	TimestampNtzOutputFormat() LegacyServiceUserParametersTimestampNtzOutputFormatList
	TimestampOutputFormat() LegacyServiceUserParametersTimestampOutputFormatList
	TimestampTypeMapping() LegacyServiceUserParametersTimestampTypeMappingList
	TimestampTzOutputFormat() LegacyServiceUserParametersTimestampTzOutputFormatList
	Timezone() LegacyServiceUserParametersTimezoneList
	TraceLevel() LegacyServiceUserParametersTraceLevelList
	TransactionAbortOnError() LegacyServiceUserParametersTransactionAbortOnErrorList
	TransactionDefaultIsolationLevel() LegacyServiceUserParametersTransactionDefaultIsolationLevelList
	TwoDigitCenturyStart() LegacyServiceUserParametersTwoDigitCenturyStartList
	UnsupportedDdlAction() LegacyServiceUserParametersUnsupportedDdlActionList
	UseCachedResult() LegacyServiceUserParametersUseCachedResultList
	WeekOfYearPolicy() LegacyServiceUserParametersWeekOfYearPolicyList
	WeekStart() LegacyServiceUserParametersWeekStartList
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LegacyServiceUserParametersOutputReference
type jsiiProxy_LegacyServiceUserParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) AbortDetachedQuery() LegacyServiceUserParametersAbortDetachedQueryList {
	var returns LegacyServiceUserParametersAbortDetachedQueryList
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) Autocommit() LegacyServiceUserParametersAutocommitList {
	var returns LegacyServiceUserParametersAutocommitList
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) BinaryInputFormat() LegacyServiceUserParametersBinaryInputFormatList {
	var returns LegacyServiceUserParametersBinaryInputFormatList
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) BinaryOutputFormat() LegacyServiceUserParametersBinaryOutputFormatList {
	var returns LegacyServiceUserParametersBinaryOutputFormatList
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ClientMemoryLimit() LegacyServiceUserParametersClientMemoryLimitList {
	var returns LegacyServiceUserParametersClientMemoryLimitList
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ClientMetadataRequestUseConnectionCtx() LegacyServiceUserParametersClientMetadataRequestUseConnectionCtxList {
	var returns LegacyServiceUserParametersClientMetadataRequestUseConnectionCtxList
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ClientPrefetchThreads() LegacyServiceUserParametersClientPrefetchThreadsList {
	var returns LegacyServiceUserParametersClientPrefetchThreadsList
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ClientResultChunkSize() LegacyServiceUserParametersClientResultChunkSizeList {
	var returns LegacyServiceUserParametersClientResultChunkSizeList
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ClientResultColumnCaseInsensitive() LegacyServiceUserParametersClientResultColumnCaseInsensitiveList {
	var returns LegacyServiceUserParametersClientResultColumnCaseInsensitiveList
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ClientSessionKeepAlive() LegacyServiceUserParametersClientSessionKeepAliveList {
	var returns LegacyServiceUserParametersClientSessionKeepAliveList
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ClientSessionKeepAliveHeartbeatFrequency() LegacyServiceUserParametersClientSessionKeepAliveHeartbeatFrequencyList {
	var returns LegacyServiceUserParametersClientSessionKeepAliveHeartbeatFrequencyList
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ClientTimestampTypeMapping() LegacyServiceUserParametersClientTimestampTypeMappingList {
	var returns LegacyServiceUserParametersClientTimestampTypeMappingList
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) DateInputFormat() LegacyServiceUserParametersDateInputFormatList {
	var returns LegacyServiceUserParametersDateInputFormatList
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) DateOutputFormat() LegacyServiceUserParametersDateOutputFormatList {
	var returns LegacyServiceUserParametersDateOutputFormatList
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) EnableUnloadPhysicalTypeOptimization() LegacyServiceUserParametersEnableUnloadPhysicalTypeOptimizationList {
	var returns LegacyServiceUserParametersEnableUnloadPhysicalTypeOptimizationList
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) EnableUnredactedQuerySyntaxError() LegacyServiceUserParametersEnableUnredactedQuerySyntaxErrorList {
	var returns LegacyServiceUserParametersEnableUnredactedQuerySyntaxErrorList
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ErrorOnNondeterministicMerge() LegacyServiceUserParametersErrorOnNondeterministicMergeList {
	var returns LegacyServiceUserParametersErrorOnNondeterministicMergeList
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) ErrorOnNondeterministicUpdate() LegacyServiceUserParametersErrorOnNondeterministicUpdateList {
	var returns LegacyServiceUserParametersErrorOnNondeterministicUpdateList
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) GeographyOutputFormat() LegacyServiceUserParametersGeographyOutputFormatList {
	var returns LegacyServiceUserParametersGeographyOutputFormatList
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) GeometryOutputFormat() LegacyServiceUserParametersGeometryOutputFormatList {
	var returns LegacyServiceUserParametersGeometryOutputFormatList
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) InternalValue() *LegacyServiceUserParameters {
	var returns *LegacyServiceUserParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) JdbcTreatDecimalAsInt() LegacyServiceUserParametersJdbcTreatDecimalAsIntList {
	var returns LegacyServiceUserParametersJdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) JdbcTreatTimestampNtzAsUtc() LegacyServiceUserParametersJdbcTreatTimestampNtzAsUtcList {
	var returns LegacyServiceUserParametersJdbcTreatTimestampNtzAsUtcList
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) JdbcUseSessionTimezone() LegacyServiceUserParametersJdbcUseSessionTimezoneList {
	var returns LegacyServiceUserParametersJdbcUseSessionTimezoneList
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) JsonIndent() LegacyServiceUserParametersJsonIndentList {
	var returns LegacyServiceUserParametersJsonIndentList
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) LockTimeout() LegacyServiceUserParametersLockTimeoutList {
	var returns LegacyServiceUserParametersLockTimeoutList
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) LogLevel() LegacyServiceUserParametersLogLevelList {
	var returns LegacyServiceUserParametersLogLevelList
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) MultiStatementCount() LegacyServiceUserParametersMultiStatementCountList {
	var returns LegacyServiceUserParametersMultiStatementCountList
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) NetworkPolicy() LegacyServiceUserParametersNetworkPolicyList {
	var returns LegacyServiceUserParametersNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) NoorderSequenceAsDefault() LegacyServiceUserParametersNoorderSequenceAsDefaultList {
	var returns LegacyServiceUserParametersNoorderSequenceAsDefaultList
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) OdbcTreatDecimalAsInt() LegacyServiceUserParametersOdbcTreatDecimalAsIntList {
	var returns LegacyServiceUserParametersOdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) PreventUnloadToInternalStages() LegacyServiceUserParametersPreventUnloadToInternalStagesList {
	var returns LegacyServiceUserParametersPreventUnloadToInternalStagesList
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) QueryTag() LegacyServiceUserParametersQueryTagList {
	var returns LegacyServiceUserParametersQueryTagList
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) QuotedIdentifiersIgnoreCase() LegacyServiceUserParametersQuotedIdentifiersIgnoreCaseList {
	var returns LegacyServiceUserParametersQuotedIdentifiersIgnoreCaseList
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) RowsPerResultset() LegacyServiceUserParametersRowsPerResultsetList {
	var returns LegacyServiceUserParametersRowsPerResultsetList
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) S3StageVpceDnsName() LegacyServiceUserParametersS3StageVpceDnsNameList {
	var returns LegacyServiceUserParametersS3StageVpceDnsNameList
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) SearchPath() LegacyServiceUserParametersSearchPathList {
	var returns LegacyServiceUserParametersSearchPathList
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) SimulatedDataSharingConsumer() LegacyServiceUserParametersSimulatedDataSharingConsumerList {
	var returns LegacyServiceUserParametersSimulatedDataSharingConsumerList
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) StatementQueuedTimeoutInSeconds() LegacyServiceUserParametersStatementQueuedTimeoutInSecondsList {
	var returns LegacyServiceUserParametersStatementQueuedTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) StatementTimeoutInSeconds() LegacyServiceUserParametersStatementTimeoutInSecondsList {
	var returns LegacyServiceUserParametersStatementTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) StrictJsonOutput() LegacyServiceUserParametersStrictJsonOutputList {
	var returns LegacyServiceUserParametersStrictJsonOutputList
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimeInputFormat() LegacyServiceUserParametersTimeInputFormatList {
	var returns LegacyServiceUserParametersTimeInputFormatList
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimeOutputFormat() LegacyServiceUserParametersTimeOutputFormatList {
	var returns LegacyServiceUserParametersTimeOutputFormatList
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimestampDayIsAlways24H() LegacyServiceUserParametersTimestampDayIsAlways24HList {
	var returns LegacyServiceUserParametersTimestampDayIsAlways24HList
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimestampInputFormat() LegacyServiceUserParametersTimestampInputFormatList {
	var returns LegacyServiceUserParametersTimestampInputFormatList
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimestampLtzOutputFormat() LegacyServiceUserParametersTimestampLtzOutputFormatList {
	var returns LegacyServiceUserParametersTimestampLtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimestampNtzOutputFormat() LegacyServiceUserParametersTimestampNtzOutputFormatList {
	var returns LegacyServiceUserParametersTimestampNtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimestampOutputFormat() LegacyServiceUserParametersTimestampOutputFormatList {
	var returns LegacyServiceUserParametersTimestampOutputFormatList
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimestampTypeMapping() LegacyServiceUserParametersTimestampTypeMappingList {
	var returns LegacyServiceUserParametersTimestampTypeMappingList
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TimestampTzOutputFormat() LegacyServiceUserParametersTimestampTzOutputFormatList {
	var returns LegacyServiceUserParametersTimestampTzOutputFormatList
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) Timezone() LegacyServiceUserParametersTimezoneList {
	var returns LegacyServiceUserParametersTimezoneList
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TraceLevel() LegacyServiceUserParametersTraceLevelList {
	var returns LegacyServiceUserParametersTraceLevelList
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TransactionAbortOnError() LegacyServiceUserParametersTransactionAbortOnErrorList {
	var returns LegacyServiceUserParametersTransactionAbortOnErrorList
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TransactionDefaultIsolationLevel() LegacyServiceUserParametersTransactionDefaultIsolationLevelList {
	var returns LegacyServiceUserParametersTransactionDefaultIsolationLevelList
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) TwoDigitCenturyStart() LegacyServiceUserParametersTwoDigitCenturyStartList {
	var returns LegacyServiceUserParametersTwoDigitCenturyStartList
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) UnsupportedDdlAction() LegacyServiceUserParametersUnsupportedDdlActionList {
	var returns LegacyServiceUserParametersUnsupportedDdlActionList
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) UseCachedResult() LegacyServiceUserParametersUseCachedResultList {
	var returns LegacyServiceUserParametersUseCachedResultList
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) WeekOfYearPolicy() LegacyServiceUserParametersWeekOfYearPolicyList {
	var returns LegacyServiceUserParametersWeekOfYearPolicyList
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference) WeekStart() LegacyServiceUserParametersWeekStartList {
	var returns LegacyServiceUserParametersWeekStartList
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}


func NewLegacyServiceUserParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LegacyServiceUserParametersOutputReference {
	_init_.Initialize()

	if err := validateNewLegacyServiceUserParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LegacyServiceUserParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUserParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLegacyServiceUserParametersOutputReference_Override(l LegacyServiceUserParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUserParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference)SetInternalValue(val *LegacyServiceUserParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

