// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserParametersOutputReference interface {
	cdktf.ComplexObject
	AbortDetachedQuery() UserParametersAbortDetachedQueryList
	Autocommit() UserParametersAutocommitList
	BinaryInputFormat() UserParametersBinaryInputFormatList
	BinaryOutputFormat() UserParametersBinaryOutputFormatList
	ClientMemoryLimit() UserParametersClientMemoryLimitList
	ClientMetadataRequestUseConnectionCtx() UserParametersClientMetadataRequestUseConnectionCtxList
	ClientPrefetchThreads() UserParametersClientPrefetchThreadsList
	ClientResultChunkSize() UserParametersClientResultChunkSizeList
	ClientResultColumnCaseInsensitive() UserParametersClientResultColumnCaseInsensitiveList
	ClientSessionKeepAlive() UserParametersClientSessionKeepAliveList
	ClientSessionKeepAliveHeartbeatFrequency() UserParametersClientSessionKeepAliveHeartbeatFrequencyList
	ClientTimestampTypeMapping() UserParametersClientTimestampTypeMappingList
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
	DateInputFormat() UserParametersDateInputFormatList
	DateOutputFormat() UserParametersDateOutputFormatList
	EnableUnloadPhysicalTypeOptimization() UserParametersEnableUnloadPhysicalTypeOptimizationList
	EnableUnredactedQuerySyntaxError() UserParametersEnableUnredactedQuerySyntaxErrorList
	ErrorOnNondeterministicMerge() UserParametersErrorOnNondeterministicMergeList
	ErrorOnNondeterministicUpdate() UserParametersErrorOnNondeterministicUpdateList
	// Experimental.
	Fqn() *string
	GeographyOutputFormat() UserParametersGeographyOutputFormatList
	GeometryOutputFormat() UserParametersGeometryOutputFormatList
	InternalValue() *UserParameters
	SetInternalValue(val *UserParameters)
	JdbcTreatDecimalAsInt() UserParametersJdbcTreatDecimalAsIntList
	JdbcTreatTimestampNtzAsUtc() UserParametersJdbcTreatTimestampNtzAsUtcList
	JdbcUseSessionTimezone() UserParametersJdbcUseSessionTimezoneList
	JsonIndent() UserParametersJsonIndentList
	LockTimeout() UserParametersLockTimeoutList
	LogLevel() UserParametersLogLevelList
	MultiStatementCount() UserParametersMultiStatementCountList
	NetworkPolicy() UserParametersNetworkPolicyList
	NoorderSequenceAsDefault() UserParametersNoorderSequenceAsDefaultList
	OdbcTreatDecimalAsInt() UserParametersOdbcTreatDecimalAsIntList
	PreventUnloadToInternalStages() UserParametersPreventUnloadToInternalStagesList
	QueryTag() UserParametersQueryTagList
	QuotedIdentifiersIgnoreCase() UserParametersQuotedIdentifiersIgnoreCaseList
	RowsPerResultset() UserParametersRowsPerResultsetList
	S3StageVpceDnsName() UserParametersS3StageVpceDnsNameList
	SearchPath() UserParametersSearchPathList
	SimulatedDataSharingConsumer() UserParametersSimulatedDataSharingConsumerList
	StatementQueuedTimeoutInSeconds() UserParametersStatementQueuedTimeoutInSecondsList
	StatementTimeoutInSeconds() UserParametersStatementTimeoutInSecondsList
	StrictJsonOutput() UserParametersStrictJsonOutputList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeInputFormat() UserParametersTimeInputFormatList
	TimeOutputFormat() UserParametersTimeOutputFormatList
	TimestampDayIsAlways24H() UserParametersTimestampDayIsAlways24HList
	TimestampInputFormat() UserParametersTimestampInputFormatList
	TimestampLtzOutputFormat() UserParametersTimestampLtzOutputFormatList
	TimestampNtzOutputFormat() UserParametersTimestampNtzOutputFormatList
	TimestampOutputFormat() UserParametersTimestampOutputFormatList
	TimestampTypeMapping() UserParametersTimestampTypeMappingList
	TimestampTzOutputFormat() UserParametersTimestampTzOutputFormatList
	Timezone() UserParametersTimezoneList
	TraceLevel() UserParametersTraceLevelList
	TransactionAbortOnError() UserParametersTransactionAbortOnErrorList
	TransactionDefaultIsolationLevel() UserParametersTransactionDefaultIsolationLevelList
	TwoDigitCenturyStart() UserParametersTwoDigitCenturyStartList
	UnsupportedDdlAction() UserParametersUnsupportedDdlActionList
	UseCachedResult() UserParametersUseCachedResultList
	WeekOfYearPolicy() UserParametersWeekOfYearPolicyList
	WeekStart() UserParametersWeekStartList
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

// The jsii proxy struct for UserParametersOutputReference
type jsiiProxy_UserParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_UserParametersOutputReference) AbortDetachedQuery() UserParametersAbortDetachedQueryList {
	var returns UserParametersAbortDetachedQueryList
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) Autocommit() UserParametersAutocommitList {
	var returns UserParametersAutocommitList
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) BinaryInputFormat() UserParametersBinaryInputFormatList {
	var returns UserParametersBinaryInputFormatList
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) BinaryOutputFormat() UserParametersBinaryOutputFormatList {
	var returns UserParametersBinaryOutputFormatList
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ClientMemoryLimit() UserParametersClientMemoryLimitList {
	var returns UserParametersClientMemoryLimitList
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ClientMetadataRequestUseConnectionCtx() UserParametersClientMetadataRequestUseConnectionCtxList {
	var returns UserParametersClientMetadataRequestUseConnectionCtxList
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ClientPrefetchThreads() UserParametersClientPrefetchThreadsList {
	var returns UserParametersClientPrefetchThreadsList
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ClientResultChunkSize() UserParametersClientResultChunkSizeList {
	var returns UserParametersClientResultChunkSizeList
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ClientResultColumnCaseInsensitive() UserParametersClientResultColumnCaseInsensitiveList {
	var returns UserParametersClientResultColumnCaseInsensitiveList
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ClientSessionKeepAlive() UserParametersClientSessionKeepAliveList {
	var returns UserParametersClientSessionKeepAliveList
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ClientSessionKeepAliveHeartbeatFrequency() UserParametersClientSessionKeepAliveHeartbeatFrequencyList {
	var returns UserParametersClientSessionKeepAliveHeartbeatFrequencyList
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ClientTimestampTypeMapping() UserParametersClientTimestampTypeMappingList {
	var returns UserParametersClientTimestampTypeMappingList
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) DateInputFormat() UserParametersDateInputFormatList {
	var returns UserParametersDateInputFormatList
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) DateOutputFormat() UserParametersDateOutputFormatList {
	var returns UserParametersDateOutputFormatList
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) EnableUnloadPhysicalTypeOptimization() UserParametersEnableUnloadPhysicalTypeOptimizationList {
	var returns UserParametersEnableUnloadPhysicalTypeOptimizationList
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) EnableUnredactedQuerySyntaxError() UserParametersEnableUnredactedQuerySyntaxErrorList {
	var returns UserParametersEnableUnredactedQuerySyntaxErrorList
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ErrorOnNondeterministicMerge() UserParametersErrorOnNondeterministicMergeList {
	var returns UserParametersErrorOnNondeterministicMergeList
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) ErrorOnNondeterministicUpdate() UserParametersErrorOnNondeterministicUpdateList {
	var returns UserParametersErrorOnNondeterministicUpdateList
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) GeographyOutputFormat() UserParametersGeographyOutputFormatList {
	var returns UserParametersGeographyOutputFormatList
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) GeometryOutputFormat() UserParametersGeometryOutputFormatList {
	var returns UserParametersGeometryOutputFormatList
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) InternalValue() *UserParameters {
	var returns *UserParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) JdbcTreatDecimalAsInt() UserParametersJdbcTreatDecimalAsIntList {
	var returns UserParametersJdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) JdbcTreatTimestampNtzAsUtc() UserParametersJdbcTreatTimestampNtzAsUtcList {
	var returns UserParametersJdbcTreatTimestampNtzAsUtcList
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) JdbcUseSessionTimezone() UserParametersJdbcUseSessionTimezoneList {
	var returns UserParametersJdbcUseSessionTimezoneList
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) JsonIndent() UserParametersJsonIndentList {
	var returns UserParametersJsonIndentList
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) LockTimeout() UserParametersLockTimeoutList {
	var returns UserParametersLockTimeoutList
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) LogLevel() UserParametersLogLevelList {
	var returns UserParametersLogLevelList
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) MultiStatementCount() UserParametersMultiStatementCountList {
	var returns UserParametersMultiStatementCountList
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) NetworkPolicy() UserParametersNetworkPolicyList {
	var returns UserParametersNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) NoorderSequenceAsDefault() UserParametersNoorderSequenceAsDefaultList {
	var returns UserParametersNoorderSequenceAsDefaultList
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) OdbcTreatDecimalAsInt() UserParametersOdbcTreatDecimalAsIntList {
	var returns UserParametersOdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) PreventUnloadToInternalStages() UserParametersPreventUnloadToInternalStagesList {
	var returns UserParametersPreventUnloadToInternalStagesList
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) QueryTag() UserParametersQueryTagList {
	var returns UserParametersQueryTagList
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) QuotedIdentifiersIgnoreCase() UserParametersQuotedIdentifiersIgnoreCaseList {
	var returns UserParametersQuotedIdentifiersIgnoreCaseList
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) RowsPerResultset() UserParametersRowsPerResultsetList {
	var returns UserParametersRowsPerResultsetList
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) S3StageVpceDnsName() UserParametersS3StageVpceDnsNameList {
	var returns UserParametersS3StageVpceDnsNameList
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) SearchPath() UserParametersSearchPathList {
	var returns UserParametersSearchPathList
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) SimulatedDataSharingConsumer() UserParametersSimulatedDataSharingConsumerList {
	var returns UserParametersSimulatedDataSharingConsumerList
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) StatementQueuedTimeoutInSeconds() UserParametersStatementQueuedTimeoutInSecondsList {
	var returns UserParametersStatementQueuedTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) StatementTimeoutInSeconds() UserParametersStatementTimeoutInSecondsList {
	var returns UserParametersStatementTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) StrictJsonOutput() UserParametersStrictJsonOutputList {
	var returns UserParametersStrictJsonOutputList
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimeInputFormat() UserParametersTimeInputFormatList {
	var returns UserParametersTimeInputFormatList
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimeOutputFormat() UserParametersTimeOutputFormatList {
	var returns UserParametersTimeOutputFormatList
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimestampDayIsAlways24H() UserParametersTimestampDayIsAlways24HList {
	var returns UserParametersTimestampDayIsAlways24HList
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimestampInputFormat() UserParametersTimestampInputFormatList {
	var returns UserParametersTimestampInputFormatList
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimestampLtzOutputFormat() UserParametersTimestampLtzOutputFormatList {
	var returns UserParametersTimestampLtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimestampNtzOutputFormat() UserParametersTimestampNtzOutputFormatList {
	var returns UserParametersTimestampNtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimestampOutputFormat() UserParametersTimestampOutputFormatList {
	var returns UserParametersTimestampOutputFormatList
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimestampTypeMapping() UserParametersTimestampTypeMappingList {
	var returns UserParametersTimestampTypeMappingList
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TimestampTzOutputFormat() UserParametersTimestampTzOutputFormatList {
	var returns UserParametersTimestampTzOutputFormatList
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) Timezone() UserParametersTimezoneList {
	var returns UserParametersTimezoneList
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TraceLevel() UserParametersTraceLevelList {
	var returns UserParametersTraceLevelList
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TransactionAbortOnError() UserParametersTransactionAbortOnErrorList {
	var returns UserParametersTransactionAbortOnErrorList
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TransactionDefaultIsolationLevel() UserParametersTransactionDefaultIsolationLevelList {
	var returns UserParametersTransactionDefaultIsolationLevelList
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) TwoDigitCenturyStart() UserParametersTwoDigitCenturyStartList {
	var returns UserParametersTwoDigitCenturyStartList
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) UnsupportedDdlAction() UserParametersUnsupportedDdlActionList {
	var returns UserParametersUnsupportedDdlActionList
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) UseCachedResult() UserParametersUseCachedResultList {
	var returns UserParametersUseCachedResultList
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) WeekOfYearPolicy() UserParametersWeekOfYearPolicyList {
	var returns UserParametersWeekOfYearPolicyList
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserParametersOutputReference) WeekStart() UserParametersWeekStartList {
	var returns UserParametersWeekStartList
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}


func NewUserParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) UserParametersOutputReference {
	_init_.Initialize()

	if err := validateNewUserParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.user.UserParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewUserParametersOutputReference_Override(u UserParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.user.UserParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		u,
	)
}

func (j *jsiiProxy_UserParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_UserParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_UserParametersOutputReference)SetInternalValue(val *UserParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_UserParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_UserParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (u *jsiiProxy_UserParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := u.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		u,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

