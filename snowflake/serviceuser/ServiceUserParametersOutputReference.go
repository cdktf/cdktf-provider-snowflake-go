// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/serviceuser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceUserParametersOutputReference interface {
	cdktf.ComplexObject
	AbortDetachedQuery() ServiceUserParametersAbortDetachedQueryList
	Autocommit() ServiceUserParametersAutocommitList
	BinaryInputFormat() ServiceUserParametersBinaryInputFormatList
	BinaryOutputFormat() ServiceUserParametersBinaryOutputFormatList
	ClientMemoryLimit() ServiceUserParametersClientMemoryLimitList
	ClientMetadataRequestUseConnectionCtx() ServiceUserParametersClientMetadataRequestUseConnectionCtxList
	ClientPrefetchThreads() ServiceUserParametersClientPrefetchThreadsList
	ClientResultChunkSize() ServiceUserParametersClientResultChunkSizeList
	ClientResultColumnCaseInsensitive() ServiceUserParametersClientResultColumnCaseInsensitiveList
	ClientSessionKeepAlive() ServiceUserParametersClientSessionKeepAliveList
	ClientSessionKeepAliveHeartbeatFrequency() ServiceUserParametersClientSessionKeepAliveHeartbeatFrequencyList
	ClientTimestampTypeMapping() ServiceUserParametersClientTimestampTypeMappingList
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
	DateInputFormat() ServiceUserParametersDateInputFormatList
	DateOutputFormat() ServiceUserParametersDateOutputFormatList
	EnableUnloadPhysicalTypeOptimization() ServiceUserParametersEnableUnloadPhysicalTypeOptimizationList
	EnableUnredactedQuerySyntaxError() ServiceUserParametersEnableUnredactedQuerySyntaxErrorList
	ErrorOnNondeterministicMerge() ServiceUserParametersErrorOnNondeterministicMergeList
	ErrorOnNondeterministicUpdate() ServiceUserParametersErrorOnNondeterministicUpdateList
	// Experimental.
	Fqn() *string
	GeographyOutputFormat() ServiceUserParametersGeographyOutputFormatList
	GeometryOutputFormat() ServiceUserParametersGeometryOutputFormatList
	InternalValue() *ServiceUserParameters
	SetInternalValue(val *ServiceUserParameters)
	JdbcTreatDecimalAsInt() ServiceUserParametersJdbcTreatDecimalAsIntList
	JdbcTreatTimestampNtzAsUtc() ServiceUserParametersJdbcTreatTimestampNtzAsUtcList
	JdbcUseSessionTimezone() ServiceUserParametersJdbcUseSessionTimezoneList
	JsonIndent() ServiceUserParametersJsonIndentList
	LockTimeout() ServiceUserParametersLockTimeoutList
	LogLevel() ServiceUserParametersLogLevelList
	MultiStatementCount() ServiceUserParametersMultiStatementCountList
	NetworkPolicy() ServiceUserParametersNetworkPolicyList
	NoorderSequenceAsDefault() ServiceUserParametersNoorderSequenceAsDefaultList
	OdbcTreatDecimalAsInt() ServiceUserParametersOdbcTreatDecimalAsIntList
	PreventUnloadToInternalStages() ServiceUserParametersPreventUnloadToInternalStagesList
	QueryTag() ServiceUserParametersQueryTagList
	QuotedIdentifiersIgnoreCase() ServiceUserParametersQuotedIdentifiersIgnoreCaseList
	RowsPerResultset() ServiceUserParametersRowsPerResultsetList
	S3StageVpceDnsName() ServiceUserParametersS3StageVpceDnsNameList
	SearchPath() ServiceUserParametersSearchPathList
	SimulatedDataSharingConsumer() ServiceUserParametersSimulatedDataSharingConsumerList
	StatementQueuedTimeoutInSeconds() ServiceUserParametersStatementQueuedTimeoutInSecondsList
	StatementTimeoutInSeconds() ServiceUserParametersStatementTimeoutInSecondsList
	StrictJsonOutput() ServiceUserParametersStrictJsonOutputList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeInputFormat() ServiceUserParametersTimeInputFormatList
	TimeOutputFormat() ServiceUserParametersTimeOutputFormatList
	TimestampDayIsAlways24H() ServiceUserParametersTimestampDayIsAlways24HList
	TimestampInputFormat() ServiceUserParametersTimestampInputFormatList
	TimestampLtzOutputFormat() ServiceUserParametersTimestampLtzOutputFormatList
	TimestampNtzOutputFormat() ServiceUserParametersTimestampNtzOutputFormatList
	TimestampOutputFormat() ServiceUserParametersTimestampOutputFormatList
	TimestampTypeMapping() ServiceUserParametersTimestampTypeMappingList
	TimestampTzOutputFormat() ServiceUserParametersTimestampTzOutputFormatList
	Timezone() ServiceUserParametersTimezoneList
	TraceLevel() ServiceUserParametersTraceLevelList
	TransactionAbortOnError() ServiceUserParametersTransactionAbortOnErrorList
	TransactionDefaultIsolationLevel() ServiceUserParametersTransactionDefaultIsolationLevelList
	TwoDigitCenturyStart() ServiceUserParametersTwoDigitCenturyStartList
	UnsupportedDdlAction() ServiceUserParametersUnsupportedDdlActionList
	UseCachedResult() ServiceUserParametersUseCachedResultList
	WeekOfYearPolicy() ServiceUserParametersWeekOfYearPolicyList
	WeekStart() ServiceUserParametersWeekStartList
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

// The jsii proxy struct for ServiceUserParametersOutputReference
type jsiiProxy_ServiceUserParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) AbortDetachedQuery() ServiceUserParametersAbortDetachedQueryList {
	var returns ServiceUserParametersAbortDetachedQueryList
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) Autocommit() ServiceUserParametersAutocommitList {
	var returns ServiceUserParametersAutocommitList
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) BinaryInputFormat() ServiceUserParametersBinaryInputFormatList {
	var returns ServiceUserParametersBinaryInputFormatList
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) BinaryOutputFormat() ServiceUserParametersBinaryOutputFormatList {
	var returns ServiceUserParametersBinaryOutputFormatList
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ClientMemoryLimit() ServiceUserParametersClientMemoryLimitList {
	var returns ServiceUserParametersClientMemoryLimitList
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ClientMetadataRequestUseConnectionCtx() ServiceUserParametersClientMetadataRequestUseConnectionCtxList {
	var returns ServiceUserParametersClientMetadataRequestUseConnectionCtxList
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ClientPrefetchThreads() ServiceUserParametersClientPrefetchThreadsList {
	var returns ServiceUserParametersClientPrefetchThreadsList
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ClientResultChunkSize() ServiceUserParametersClientResultChunkSizeList {
	var returns ServiceUserParametersClientResultChunkSizeList
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ClientResultColumnCaseInsensitive() ServiceUserParametersClientResultColumnCaseInsensitiveList {
	var returns ServiceUserParametersClientResultColumnCaseInsensitiveList
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ClientSessionKeepAlive() ServiceUserParametersClientSessionKeepAliveList {
	var returns ServiceUserParametersClientSessionKeepAliveList
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ClientSessionKeepAliveHeartbeatFrequency() ServiceUserParametersClientSessionKeepAliveHeartbeatFrequencyList {
	var returns ServiceUserParametersClientSessionKeepAliveHeartbeatFrequencyList
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ClientTimestampTypeMapping() ServiceUserParametersClientTimestampTypeMappingList {
	var returns ServiceUserParametersClientTimestampTypeMappingList
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) DateInputFormat() ServiceUserParametersDateInputFormatList {
	var returns ServiceUserParametersDateInputFormatList
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) DateOutputFormat() ServiceUserParametersDateOutputFormatList {
	var returns ServiceUserParametersDateOutputFormatList
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) EnableUnloadPhysicalTypeOptimization() ServiceUserParametersEnableUnloadPhysicalTypeOptimizationList {
	var returns ServiceUserParametersEnableUnloadPhysicalTypeOptimizationList
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) EnableUnredactedQuerySyntaxError() ServiceUserParametersEnableUnredactedQuerySyntaxErrorList {
	var returns ServiceUserParametersEnableUnredactedQuerySyntaxErrorList
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ErrorOnNondeterministicMerge() ServiceUserParametersErrorOnNondeterministicMergeList {
	var returns ServiceUserParametersErrorOnNondeterministicMergeList
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) ErrorOnNondeterministicUpdate() ServiceUserParametersErrorOnNondeterministicUpdateList {
	var returns ServiceUserParametersErrorOnNondeterministicUpdateList
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) GeographyOutputFormat() ServiceUserParametersGeographyOutputFormatList {
	var returns ServiceUserParametersGeographyOutputFormatList
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) GeometryOutputFormat() ServiceUserParametersGeometryOutputFormatList {
	var returns ServiceUserParametersGeometryOutputFormatList
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) InternalValue() *ServiceUserParameters {
	var returns *ServiceUserParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) JdbcTreatDecimalAsInt() ServiceUserParametersJdbcTreatDecimalAsIntList {
	var returns ServiceUserParametersJdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) JdbcTreatTimestampNtzAsUtc() ServiceUserParametersJdbcTreatTimestampNtzAsUtcList {
	var returns ServiceUserParametersJdbcTreatTimestampNtzAsUtcList
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) JdbcUseSessionTimezone() ServiceUserParametersJdbcUseSessionTimezoneList {
	var returns ServiceUserParametersJdbcUseSessionTimezoneList
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) JsonIndent() ServiceUserParametersJsonIndentList {
	var returns ServiceUserParametersJsonIndentList
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) LockTimeout() ServiceUserParametersLockTimeoutList {
	var returns ServiceUserParametersLockTimeoutList
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) LogLevel() ServiceUserParametersLogLevelList {
	var returns ServiceUserParametersLogLevelList
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) MultiStatementCount() ServiceUserParametersMultiStatementCountList {
	var returns ServiceUserParametersMultiStatementCountList
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) NetworkPolicy() ServiceUserParametersNetworkPolicyList {
	var returns ServiceUserParametersNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) NoorderSequenceAsDefault() ServiceUserParametersNoorderSequenceAsDefaultList {
	var returns ServiceUserParametersNoorderSequenceAsDefaultList
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) OdbcTreatDecimalAsInt() ServiceUserParametersOdbcTreatDecimalAsIntList {
	var returns ServiceUserParametersOdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) PreventUnloadToInternalStages() ServiceUserParametersPreventUnloadToInternalStagesList {
	var returns ServiceUserParametersPreventUnloadToInternalStagesList
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) QueryTag() ServiceUserParametersQueryTagList {
	var returns ServiceUserParametersQueryTagList
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) QuotedIdentifiersIgnoreCase() ServiceUserParametersQuotedIdentifiersIgnoreCaseList {
	var returns ServiceUserParametersQuotedIdentifiersIgnoreCaseList
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) RowsPerResultset() ServiceUserParametersRowsPerResultsetList {
	var returns ServiceUserParametersRowsPerResultsetList
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) S3StageVpceDnsName() ServiceUserParametersS3StageVpceDnsNameList {
	var returns ServiceUserParametersS3StageVpceDnsNameList
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) SearchPath() ServiceUserParametersSearchPathList {
	var returns ServiceUserParametersSearchPathList
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) SimulatedDataSharingConsumer() ServiceUserParametersSimulatedDataSharingConsumerList {
	var returns ServiceUserParametersSimulatedDataSharingConsumerList
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) StatementQueuedTimeoutInSeconds() ServiceUserParametersStatementQueuedTimeoutInSecondsList {
	var returns ServiceUserParametersStatementQueuedTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) StatementTimeoutInSeconds() ServiceUserParametersStatementTimeoutInSecondsList {
	var returns ServiceUserParametersStatementTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) StrictJsonOutput() ServiceUserParametersStrictJsonOutputList {
	var returns ServiceUserParametersStrictJsonOutputList
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimeInputFormat() ServiceUserParametersTimeInputFormatList {
	var returns ServiceUserParametersTimeInputFormatList
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimeOutputFormat() ServiceUserParametersTimeOutputFormatList {
	var returns ServiceUserParametersTimeOutputFormatList
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimestampDayIsAlways24H() ServiceUserParametersTimestampDayIsAlways24HList {
	var returns ServiceUserParametersTimestampDayIsAlways24HList
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimestampInputFormat() ServiceUserParametersTimestampInputFormatList {
	var returns ServiceUserParametersTimestampInputFormatList
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimestampLtzOutputFormat() ServiceUserParametersTimestampLtzOutputFormatList {
	var returns ServiceUserParametersTimestampLtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimestampNtzOutputFormat() ServiceUserParametersTimestampNtzOutputFormatList {
	var returns ServiceUserParametersTimestampNtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimestampOutputFormat() ServiceUserParametersTimestampOutputFormatList {
	var returns ServiceUserParametersTimestampOutputFormatList
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimestampTypeMapping() ServiceUserParametersTimestampTypeMappingList {
	var returns ServiceUserParametersTimestampTypeMappingList
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TimestampTzOutputFormat() ServiceUserParametersTimestampTzOutputFormatList {
	var returns ServiceUserParametersTimestampTzOutputFormatList
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) Timezone() ServiceUserParametersTimezoneList {
	var returns ServiceUserParametersTimezoneList
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TraceLevel() ServiceUserParametersTraceLevelList {
	var returns ServiceUserParametersTraceLevelList
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TransactionAbortOnError() ServiceUserParametersTransactionAbortOnErrorList {
	var returns ServiceUserParametersTransactionAbortOnErrorList
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TransactionDefaultIsolationLevel() ServiceUserParametersTransactionDefaultIsolationLevelList {
	var returns ServiceUserParametersTransactionDefaultIsolationLevelList
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) TwoDigitCenturyStart() ServiceUserParametersTwoDigitCenturyStartList {
	var returns ServiceUserParametersTwoDigitCenturyStartList
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) UnsupportedDdlAction() ServiceUserParametersUnsupportedDdlActionList {
	var returns ServiceUserParametersUnsupportedDdlActionList
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) UseCachedResult() ServiceUserParametersUseCachedResultList {
	var returns ServiceUserParametersUseCachedResultList
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) WeekOfYearPolicy() ServiceUserParametersWeekOfYearPolicyList {
	var returns ServiceUserParametersWeekOfYearPolicyList
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserParametersOutputReference) WeekStart() ServiceUserParametersWeekStartList {
	var returns ServiceUserParametersWeekStartList
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}


func NewServiceUserParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceUserParametersOutputReference {
	_init_.Initialize()

	if err := validateNewServiceUserParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceUserParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.serviceUser.ServiceUserParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceUserParametersOutputReference_Override(s ServiceUserParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.serviceUser.ServiceUserParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceUserParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceUserParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceUserParametersOutputReference)SetInternalValue(val *ServiceUserParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceUserParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceUserParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

