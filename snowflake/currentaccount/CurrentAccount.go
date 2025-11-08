// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package currentaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/currentaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/current_account snowflake_current_account}.
type CurrentAccount interface {
	cdktf.TerraformResource
	AbortDetachedQuery() interface{}
	SetAbortDetachedQuery(val interface{})
	AbortDetachedQueryInput() interface{}
	ActivePythonProfiler() *string
	SetActivePythonProfiler(val *string)
	ActivePythonProfilerInput() *string
	AllowClientMfaCaching() interface{}
	SetAllowClientMfaCaching(val interface{})
	AllowClientMfaCachingInput() interface{}
	AllowIdToken() interface{}
	SetAllowIdToken(val interface{})
	AllowIdTokenInput() interface{}
	AuthenticationPolicy() *string
	SetAuthenticationPolicy(val *string)
	AuthenticationPolicyInput() *string
	Autocommit() interface{}
	SetAutocommit(val interface{})
	AutocommitInput() interface{}
	BaseLocationPrefix() *string
	SetBaseLocationPrefix(val *string)
	BaseLocationPrefixInput() *string
	BinaryInputFormat() *string
	SetBinaryInputFormat(val *string)
	BinaryInputFormatInput() *string
	BinaryOutputFormat() *string
	SetBinaryOutputFormat(val *string)
	BinaryOutputFormatInput() *string
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
	CatalogSync() *string
	SetCatalogSync(val *string)
	CatalogSyncInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientEnableLogInfoStatementParameters() interface{}
	SetClientEnableLogInfoStatementParameters(val interface{})
	ClientEnableLogInfoStatementParametersInput() interface{}
	ClientEncryptionKeySize() *float64
	SetClientEncryptionKeySize(val *float64)
	ClientEncryptionKeySizeInput() *float64
	ClientMemoryLimit() *float64
	SetClientMemoryLimit(val *float64)
	ClientMemoryLimitInput() *float64
	ClientMetadataRequestUseConnectionCtx() interface{}
	SetClientMetadataRequestUseConnectionCtx(val interface{})
	ClientMetadataRequestUseConnectionCtxInput() interface{}
	ClientMetadataUseSessionDatabase() interface{}
	SetClientMetadataUseSessionDatabase(val interface{})
	ClientMetadataUseSessionDatabaseInput() interface{}
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
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CortexEnabledCrossRegion() *string
	SetCortexEnabledCrossRegion(val *string)
	CortexEnabledCrossRegionInput() *string
	CortexModelsAllowlist() *string
	SetCortexModelsAllowlist(val *string)
	CortexModelsAllowlistInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CsvTimestampFormat() *string
	SetCsvTimestampFormat(val *string)
	CsvTimestampFormatInput() *string
	DataRetentionTimeInDays() *float64
	SetDataRetentionTimeInDays(val *float64)
	DataRetentionTimeInDaysInput() *float64
	DateInputFormat() *string
	SetDateInputFormat(val *string)
	DateInputFormatInput() *string
	DateOutputFormat() *string
	SetDateOutputFormat(val *string)
	DateOutputFormatInput() *string
	DefaultDdlCollation() *string
	SetDefaultDdlCollation(val *string)
	DefaultDdlCollationInput() *string
	DefaultNotebookComputePoolCpu() *string
	SetDefaultNotebookComputePoolCpu(val *string)
	DefaultNotebookComputePoolCpuInput() *string
	DefaultNotebookComputePoolGpu() *string
	SetDefaultNotebookComputePoolGpu(val *string)
	DefaultNotebookComputePoolGpuInput() *string
	DefaultNullOrdering() *string
	SetDefaultNullOrdering(val *string)
	DefaultNullOrderingInput() *string
	DefaultStreamlitNotebookWarehouse() *string
	SetDefaultStreamlitNotebookWarehouse(val *string)
	DefaultStreamlitNotebookWarehouseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableUiDownloadButton() interface{}
	SetDisableUiDownloadButton(val interface{})
	DisableUiDownloadButtonInput() interface{}
	DisableUserPrivilegeGrants() interface{}
	SetDisableUserPrivilegeGrants(val interface{})
	DisableUserPrivilegeGrantsInput() interface{}
	EnableAutomaticSensitiveDataClassificationLog() interface{}
	SetEnableAutomaticSensitiveDataClassificationLog(val interface{})
	EnableAutomaticSensitiveDataClassificationLogInput() interface{}
	EnableEgressCostOptimizer() interface{}
	SetEnableEgressCostOptimizer(val interface{})
	EnableEgressCostOptimizerInput() interface{}
	EnableIdentifierFirstLogin() interface{}
	SetEnableIdentifierFirstLogin(val interface{})
	EnableIdentifierFirstLoginInput() interface{}
	EnableInternalStagesPrivatelink() interface{}
	SetEnableInternalStagesPrivatelink(val interface{})
	EnableInternalStagesPrivatelinkInput() interface{}
	EnableTriSecretAndRekeyOptOutForImageRepository() interface{}
	SetEnableTriSecretAndRekeyOptOutForImageRepository(val interface{})
	EnableTriSecretAndRekeyOptOutForImageRepositoryInput() interface{}
	EnableTriSecretAndRekeyOptOutForSpcsBlockStorage() interface{}
	SetEnableTriSecretAndRekeyOptOutForSpcsBlockStorage(val interface{})
	EnableTriSecretAndRekeyOptOutForSpcsBlockStorageInput() interface{}
	EnableUnhandledExceptionsReporting() interface{}
	SetEnableUnhandledExceptionsReporting(val interface{})
	EnableUnhandledExceptionsReportingInput() interface{}
	EnableUnloadPhysicalTypeOptimization() interface{}
	SetEnableUnloadPhysicalTypeOptimization(val interface{})
	EnableUnloadPhysicalTypeOptimizationInput() interface{}
	EnableUnredactedQuerySyntaxError() interface{}
	SetEnableUnredactedQuerySyntaxError(val interface{})
	EnableUnredactedQuerySyntaxErrorInput() interface{}
	EnableUnredactedSecureObjectError() interface{}
	SetEnableUnredactedSecureObjectError(val interface{})
	EnableUnredactedSecureObjectErrorInput() interface{}
	EnforceNetworkRulesForInternalStages() interface{}
	SetEnforceNetworkRulesForInternalStages(val interface{})
	EnforceNetworkRulesForInternalStagesInput() interface{}
	ErrorOnNondeterministicMerge() interface{}
	SetErrorOnNondeterministicMerge(val interface{})
	ErrorOnNondeterministicMergeInput() interface{}
	ErrorOnNondeterministicUpdate() interface{}
	SetErrorOnNondeterministicUpdate(val interface{})
	ErrorOnNondeterministicUpdateInput() interface{}
	EventTable() *string
	SetEventTable(val *string)
	EventTableInput() *string
	ExternalOauthAddPrivilegedRolesToBlockedList() interface{}
	SetExternalOauthAddPrivilegedRolesToBlockedList(val interface{})
	ExternalOauthAddPrivilegedRolesToBlockedListInput() interface{}
	ExternalVolume() *string
	SetExternalVolume(val *string)
	ExternalVolumeInput() *string
	FeaturePolicy() *string
	SetFeaturePolicy(val *string)
	FeaturePolicyInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeographyOutputFormat() *string
	SetGeographyOutputFormat(val *string)
	GeographyOutputFormatInput() *string
	GeometryOutputFormat() *string
	SetGeometryOutputFormat(val *string)
	GeometryOutputFormatInput() *string
	HybridTableLockTimeout() *float64
	SetHybridTableLockTimeout(val *float64)
	HybridTableLockTimeoutInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitialReplicationSizeLimitInTb() *string
	SetInitialReplicationSizeLimitInTb(val *string)
	InitialReplicationSizeLimitInTbInput() *string
	JdbcTreatDecimalAsInt() interface{}
	SetJdbcTreatDecimalAsInt(val interface{})
	JdbcTreatDecimalAsIntInput() interface{}
	JdbcTreatTimestampNtzAsUtc() interface{}
	SetJdbcTreatTimestampNtzAsUtc(val interface{})
	JdbcTreatTimestampNtzAsUtcInput() interface{}
	JdbcUseSessionTimezone() interface{}
	SetJdbcUseSessionTimezone(val interface{})
	JdbcUseSessionTimezoneInput() interface{}
	JsonIndent() *float64
	SetJsonIndent(val *float64)
	JsonIndentInput() *float64
	JsTreatIntegerAsBigint() interface{}
	SetJsTreatIntegerAsBigint(val interface{})
	JsTreatIntegerAsBigintInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListingAutoFulfillmentReplicationRefreshSchedule() *string
	SetListingAutoFulfillmentReplicationRefreshSchedule(val *string)
	ListingAutoFulfillmentReplicationRefreshScheduleInput() *string
	LockTimeout() *float64
	SetLockTimeout(val *float64)
	LockTimeoutInput() *float64
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	MaxConcurrencyLevel() *float64
	SetMaxConcurrencyLevel(val *float64)
	MaxConcurrencyLevelInput() *float64
	MaxDataExtensionTimeInDays() *float64
	SetMaxDataExtensionTimeInDays(val *float64)
	MaxDataExtensionTimeInDaysInput() *float64
	MetricLevel() *string
	SetMetricLevel(val *string)
	MetricLevelInput() *string
	MinDataRetentionTimeInDays() *float64
	SetMinDataRetentionTimeInDays(val *float64)
	MinDataRetentionTimeInDaysInput() *float64
	MultiStatementCount() *float64
	SetMultiStatementCount(val *float64)
	MultiStatementCountInput() *float64
	NetworkPolicy() *string
	SetNetworkPolicy(val *string)
	NetworkPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	NoorderSequenceAsDefault() interface{}
	SetNoorderSequenceAsDefault(val interface{})
	NoorderSequenceAsDefaultInput() interface{}
	OauthAddPrivilegedRolesToBlockedList() interface{}
	SetOauthAddPrivilegedRolesToBlockedList(val interface{})
	OauthAddPrivilegedRolesToBlockedListInput() interface{}
	OdbcTreatDecimalAsInt() interface{}
	SetOdbcTreatDecimalAsInt(val interface{})
	OdbcTreatDecimalAsIntInput() interface{}
	PackagesPolicy() *string
	SetPackagesPolicy(val *string)
	PackagesPolicyInput() *string
	PasswordPolicy() *string
	SetPasswordPolicy(val *string)
	PasswordPolicyInput() *string
	PeriodicDataRekeying() interface{}
	SetPeriodicDataRekeying(val interface{})
	PeriodicDataRekeyingInput() interface{}
	PipeExecutionPaused() interface{}
	SetPipeExecutionPaused(val interface{})
	PipeExecutionPausedInput() interface{}
	PreventUnloadToInlineUrl() interface{}
	SetPreventUnloadToInlineUrl(val interface{})
	PreventUnloadToInlineUrlInput() interface{}
	PreventUnloadToInternalStages() interface{}
	SetPreventUnloadToInternalStages(val interface{})
	PreventUnloadToInternalStagesInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PythonProfilerModules() *string
	SetPythonProfilerModules(val *string)
	PythonProfilerModulesInput() *string
	PythonProfilerTargetStage() *string
	SetPythonProfilerTargetStage(val *string)
	PythonProfilerTargetStageInput() *string
	QueryTag() *string
	SetQueryTag(val *string)
	QueryTagInput() *string
	QuotedIdentifiersIgnoreCase() interface{}
	SetQuotedIdentifiersIgnoreCase(val interface{})
	QuotedIdentifiersIgnoreCaseInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReplaceInvalidCharacters() interface{}
	SetReplaceInvalidCharacters(val interface{})
	ReplaceInvalidCharactersInput() interface{}
	RequireStorageIntegrationForStageCreation() interface{}
	SetRequireStorageIntegrationForStageCreation(val interface{})
	RequireStorageIntegrationForStageCreationInput() interface{}
	RequireStorageIntegrationForStageOperation() interface{}
	SetRequireStorageIntegrationForStageOperation(val interface{})
	RequireStorageIntegrationForStageOperationInput() interface{}
	ResourceMonitor() *string
	SetResourceMonitor(val *string)
	ResourceMonitorInput() *string
	RowsPerResultset() *float64
	SetRowsPerResultset(val *float64)
	RowsPerResultsetInput() *float64
	S3StageVpceDnsName() *string
	SetS3StageVpceDnsName(val *string)
	S3StageVpceDnsNameInput() *string
	SamlIdentityProvider() *string
	SetSamlIdentityProvider(val *string)
	SamlIdentityProviderInput() *string
	SearchPath() *string
	SetSearchPath(val *string)
	SearchPathInput() *string
	ServerlessTaskMaxStatementSize() *string
	SetServerlessTaskMaxStatementSize(val *string)
	ServerlessTaskMaxStatementSizeInput() *string
	ServerlessTaskMinStatementSize() *string
	SetServerlessTaskMinStatementSize(val *string)
	ServerlessTaskMinStatementSizeInput() *string
	SessionPolicy() *string
	SetSessionPolicy(val *string)
	SessionPolicyInput() *string
	SimulatedDataSharingConsumer() *string
	SetSimulatedDataSharingConsumer(val *string)
	SimulatedDataSharingConsumerInput() *string
	SsoLoginPage() interface{}
	SetSsoLoginPage(val interface{})
	SsoLoginPageInput() interface{}
	StatementQueuedTimeoutInSeconds() *float64
	SetStatementQueuedTimeoutInSeconds(val *float64)
	StatementQueuedTimeoutInSecondsInput() *float64
	StatementTimeoutInSeconds() *float64
	SetStatementTimeoutInSeconds(val *float64)
	StatementTimeoutInSecondsInput() *float64
	StorageSerializationPolicy() *string
	SetStorageSerializationPolicy(val *string)
	StorageSerializationPolicyInput() *string
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
	Timeouts() CurrentAccountTimeoutsOutputReference
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
	WeekOfYearPolicy() *float64
	SetWeekOfYearPolicy(val *float64)
	WeekOfYearPolicyInput() *float64
	WeekStart() *float64
	SetWeekStart(val *float64)
	WeekStartInput() *float64
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
	PutTimeouts(value *CurrentAccountTimeouts)
	ResetAbortDetachedQuery()
	ResetActivePythonProfiler()
	ResetAllowClientMfaCaching()
	ResetAllowIdToken()
	ResetAuthenticationPolicy()
	ResetAutocommit()
	ResetBaseLocationPrefix()
	ResetBinaryInputFormat()
	ResetBinaryOutputFormat()
	ResetCatalog()
	ResetCatalogSync()
	ResetClientEnableLogInfoStatementParameters()
	ResetClientEncryptionKeySize()
	ResetClientMemoryLimit()
	ResetClientMetadataRequestUseConnectionCtx()
	ResetClientMetadataUseSessionDatabase()
	ResetClientPrefetchThreads()
	ResetClientResultChunkSize()
	ResetClientResultColumnCaseInsensitive()
	ResetClientSessionKeepAlive()
	ResetClientSessionKeepAliveHeartbeatFrequency()
	ResetClientTimestampTypeMapping()
	ResetCortexEnabledCrossRegion()
	ResetCortexModelsAllowlist()
	ResetCsvTimestampFormat()
	ResetDataRetentionTimeInDays()
	ResetDateInputFormat()
	ResetDateOutputFormat()
	ResetDefaultDdlCollation()
	ResetDefaultNotebookComputePoolCpu()
	ResetDefaultNotebookComputePoolGpu()
	ResetDefaultNullOrdering()
	ResetDefaultStreamlitNotebookWarehouse()
	ResetDisableUiDownloadButton()
	ResetDisableUserPrivilegeGrants()
	ResetEnableAutomaticSensitiveDataClassificationLog()
	ResetEnableEgressCostOptimizer()
	ResetEnableIdentifierFirstLogin()
	ResetEnableInternalStagesPrivatelink()
	ResetEnableTriSecretAndRekeyOptOutForImageRepository()
	ResetEnableTriSecretAndRekeyOptOutForSpcsBlockStorage()
	ResetEnableUnhandledExceptionsReporting()
	ResetEnableUnloadPhysicalTypeOptimization()
	ResetEnableUnredactedQuerySyntaxError()
	ResetEnableUnredactedSecureObjectError()
	ResetEnforceNetworkRulesForInternalStages()
	ResetErrorOnNondeterministicMerge()
	ResetErrorOnNondeterministicUpdate()
	ResetEventTable()
	ResetExternalOauthAddPrivilegedRolesToBlockedList()
	ResetExternalVolume()
	ResetFeaturePolicy()
	ResetGeographyOutputFormat()
	ResetGeometryOutputFormat()
	ResetHybridTableLockTimeout()
	ResetId()
	ResetInitialReplicationSizeLimitInTb()
	ResetJdbcTreatDecimalAsInt()
	ResetJdbcTreatTimestampNtzAsUtc()
	ResetJdbcUseSessionTimezone()
	ResetJsonIndent()
	ResetJsTreatIntegerAsBigint()
	ResetListingAutoFulfillmentReplicationRefreshSchedule()
	ResetLockTimeout()
	ResetLogLevel()
	ResetMaxConcurrencyLevel()
	ResetMaxDataExtensionTimeInDays()
	ResetMetricLevel()
	ResetMinDataRetentionTimeInDays()
	ResetMultiStatementCount()
	ResetNetworkPolicy()
	ResetNoorderSequenceAsDefault()
	ResetOauthAddPrivilegedRolesToBlockedList()
	ResetOdbcTreatDecimalAsInt()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackagesPolicy()
	ResetPasswordPolicy()
	ResetPeriodicDataRekeying()
	ResetPipeExecutionPaused()
	ResetPreventUnloadToInlineUrl()
	ResetPreventUnloadToInternalStages()
	ResetPythonProfilerModules()
	ResetPythonProfilerTargetStage()
	ResetQueryTag()
	ResetQuotedIdentifiersIgnoreCase()
	ResetReplaceInvalidCharacters()
	ResetRequireStorageIntegrationForStageCreation()
	ResetRequireStorageIntegrationForStageOperation()
	ResetResourceMonitor()
	ResetRowsPerResultset()
	ResetS3StageVpceDnsName()
	ResetSamlIdentityProvider()
	ResetSearchPath()
	ResetServerlessTaskMaxStatementSize()
	ResetServerlessTaskMinStatementSize()
	ResetSessionPolicy()
	ResetSimulatedDataSharingConsumer()
	ResetSsoLoginPage()
	ResetStatementQueuedTimeoutInSeconds()
	ResetStatementTimeoutInSeconds()
	ResetStorageSerializationPolicy()
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
	ResetWeekOfYearPolicy()
	ResetWeekStart()
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

// The jsii proxy struct for CurrentAccount
type jsiiProxy_CurrentAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CurrentAccount) AbortDetachedQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) AbortDetachedQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ActivePythonProfiler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activePythonProfiler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ActivePythonProfilerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activePythonProfilerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) AllowClientMfaCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClientMfaCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) AllowClientMfaCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClientMfaCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) AllowIdToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowIdToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) AllowIdTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowIdTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) AuthenticationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) AuthenticationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Autocommit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) AutocommitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) BaseLocationPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseLocationPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) BaseLocationPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseLocationPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) BinaryInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) BinaryInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) BinaryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) BinaryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CatalogSync() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CatalogSyncInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientEnableLogInfoStatementParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientEnableLogInfoStatementParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientEnableLogInfoStatementParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientEnableLogInfoStatementParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientEncryptionKeySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientEncryptionKeySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientEncryptionKeySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientEncryptionKeySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientMemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientMemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientMetadataRequestUseConnectionCtx() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientMetadataRequestUseConnectionCtxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientMetadataUseSessionDatabase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataUseSessionDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientMetadataUseSessionDatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataUseSessionDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientPrefetchThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientPrefetchThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientResultChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientResultChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientResultColumnCaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientResultColumnCaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientSessionKeepAlive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientSessionKeepAliveHeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientSessionKeepAliveHeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientSessionKeepAliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientTimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ClientTimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CortexEnabledCrossRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cortexEnabledCrossRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CortexEnabledCrossRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cortexEnabledCrossRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CortexModelsAllowlist() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cortexModelsAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CortexModelsAllowlistInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cortexModelsAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CsvTimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvTimestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) CsvTimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvTimestampFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DataRetentionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DataRetentionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DateInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DateInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DateOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DateOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultDdlCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDdlCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultDdlCollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDdlCollationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultNotebookComputePoolCpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultNotebookComputePoolCpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultNotebookComputePoolGpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultNotebookComputePoolGpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolGpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultNullOrdering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNullOrdering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultNullOrderingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNullOrderingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultStreamlitNotebookWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStreamlitNotebookWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DefaultStreamlitNotebookWarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStreamlitNotebookWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DisableUiDownloadButton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUiDownloadButton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DisableUiDownloadButtonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUiDownloadButtonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DisableUserPrivilegeGrants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUserPrivilegeGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) DisableUserPrivilegeGrantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUserPrivilegeGrantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableAutomaticSensitiveDataClassificationLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticSensitiveDataClassificationLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableAutomaticSensitiveDataClassificationLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticSensitiveDataClassificationLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableEgressCostOptimizer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEgressCostOptimizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableEgressCostOptimizerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEgressCostOptimizerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableIdentifierFirstLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIdentifierFirstLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableIdentifierFirstLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIdentifierFirstLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableInternalStagesPrivatelink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternalStagesPrivatelink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableInternalStagesPrivatelinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternalStagesPrivatelinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableTriSecretAndRekeyOptOutForImageRepository() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForImageRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableTriSecretAndRekeyOptOutForImageRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForImageRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableTriSecretAndRekeyOptOutForSpcsBlockStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForSpcsBlockStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableTriSecretAndRekeyOptOutForSpcsBlockStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForSpcsBlockStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableUnhandledExceptionsReporting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnhandledExceptionsReporting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableUnhandledExceptionsReportingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnhandledExceptionsReportingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableUnloadPhysicalTypeOptimization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableUnloadPhysicalTypeOptimizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableUnredactedQuerySyntaxError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableUnredactedQuerySyntaxErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableUnredactedSecureObjectError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedSecureObjectError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnableUnredactedSecureObjectErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedSecureObjectErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnforceNetworkRulesForInternalStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceNetworkRulesForInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EnforceNetworkRulesForInternalStagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceNetworkRulesForInternalStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ErrorOnNondeterministicMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ErrorOnNondeterministicMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ErrorOnNondeterministicUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ErrorOnNondeterministicUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EventTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) EventTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ExternalOauthAddPrivilegedRolesToBlockedList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalOauthAddPrivilegedRolesToBlockedList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ExternalOauthAddPrivilegedRolesToBlockedListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalOauthAddPrivilegedRolesToBlockedListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ExternalVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ExternalVolumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) FeaturePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) FeaturePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) GeographyOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) GeographyOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) GeometryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) GeometryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) HybridTableLockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hybridTableLockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) HybridTableLockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hybridTableLockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) InitialReplicationSizeLimitInTb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialReplicationSizeLimitInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) InitialReplicationSizeLimitInTbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialReplicationSizeLimitInTbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JdbcTreatTimestampNtzAsUtc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JdbcTreatTimestampNtzAsUtcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JdbcUseSessionTimezone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JdbcUseSessionTimezoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JsonIndent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JsonIndentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JsTreatIntegerAsBigint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsTreatIntegerAsBigint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) JsTreatIntegerAsBigintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsTreatIntegerAsBigintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ListingAutoFulfillmentReplicationRefreshSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingAutoFulfillmentReplicationRefreshSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ListingAutoFulfillmentReplicationRefreshScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingAutoFulfillmentReplicationRefreshScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) LockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) LockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MaxConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MaxConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MaxDataExtensionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MaxDataExtensionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MetricLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MetricLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MinDataRetentionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MinDataRetentionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDataRetentionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MultiStatementCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) MultiStatementCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) NetworkPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) NetworkPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) NoorderSequenceAsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) NoorderSequenceAsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) OauthAddPrivilegedRolesToBlockedList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthAddPrivilegedRolesToBlockedList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) OauthAddPrivilegedRolesToBlockedListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthAddPrivilegedRolesToBlockedListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) OdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) OdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PackagesPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PackagesPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagesPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PasswordPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PasswordPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PeriodicDataRekeying() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"periodicDataRekeying",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PeriodicDataRekeyingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"periodicDataRekeyingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PipeExecutionPaused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipeExecutionPaused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PipeExecutionPausedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipeExecutionPausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PreventUnloadToInlineUrl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInlineUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PreventUnloadToInlineUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInlineUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PreventUnloadToInternalStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PreventUnloadToInternalStagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PythonProfilerModules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProfilerModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PythonProfilerModulesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProfilerModulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PythonProfilerTargetStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProfilerTargetStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) PythonProfilerTargetStageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProfilerTargetStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) QueryTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) QueryTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) QuotedIdentifiersIgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) QuotedIdentifiersIgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ReplaceInvalidCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ReplaceInvalidCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) RequireStorageIntegrationForStageCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) RequireStorageIntegrationForStageCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) RequireStorageIntegrationForStageOperation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) RequireStorageIntegrationForStageOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ResourceMonitor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ResourceMonitorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) RowsPerResultset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) RowsPerResultsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) S3StageVpceDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) S3StageVpceDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SamlIdentityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlIdentityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SamlIdentityProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlIdentityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SearchPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SearchPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ServerlessTaskMaxStatementSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessTaskMaxStatementSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ServerlessTaskMaxStatementSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessTaskMaxStatementSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ServerlessTaskMinStatementSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessTaskMinStatementSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) ServerlessTaskMinStatementSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessTaskMinStatementSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SessionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SessionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SimulatedDataSharingConsumer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SimulatedDataSharingConsumerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SsoLoginPage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssoLoginPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SsoLoginPageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssoLoginPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) StatementQueuedTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) StatementQueuedTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) StatementTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) StatementTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) StorageSerializationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) StorageSerializationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) StrictJsonOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) StrictJsonOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SuspendTaskAfterNumFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) SuspendTaskAfterNumFailuresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TaskAutoRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TaskAutoRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimeInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimeInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimeOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimeOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Timeouts() CurrentAccountTimeoutsOutputReference {
	var returns CurrentAccountTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampDayIsAlways24H() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampDayIsAlways24HInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24HInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampLtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampLtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampNtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampNtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampTzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimestampTzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TransactionAbortOnError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TransactionAbortOnErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TransactionDefaultIsolationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TransactionDefaultIsolationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TwoDigitCenturyStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) TwoDigitCenturyStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UnsupportedDdlAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UnsupportedDdlActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UseCachedResult() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UseCachedResultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UserTaskManagedInitialWarehouseSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UserTaskManagedInitialWarehouseSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UserTaskMinimumTriggerIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UserTaskMinimumTriggerIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UserTaskTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) UserTaskTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) WeekOfYearPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) WeekOfYearPolicyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) WeekStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentAccount) WeekStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStartInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/current_account snowflake_current_account} Resource.
func NewCurrentAccount(scope constructs.Construct, id *string, config *CurrentAccountConfig) CurrentAccount {
	_init_.Initialize()

	if err := validateNewCurrentAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CurrentAccount{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.currentAccount.CurrentAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/current_account snowflake_current_account} Resource.
func NewCurrentAccount_Override(c CurrentAccount, scope constructs.Construct, id *string, config *CurrentAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.currentAccount.CurrentAccount",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CurrentAccount)SetAbortDetachedQuery(val interface{}) {
	if err := j.validateSetAbortDetachedQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortDetachedQuery",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetActivePythonProfiler(val *string) {
	if err := j.validateSetActivePythonProfilerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activePythonProfiler",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetAllowClientMfaCaching(val interface{}) {
	if err := j.validateSetAllowClientMfaCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClientMfaCaching",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetAllowIdToken(val interface{}) {
	if err := j.validateSetAllowIdTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowIdToken",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetAuthenticationPolicy(val *string) {
	if err := j.validateSetAuthenticationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetAutocommit(val interface{}) {
	if err := j.validateSetAutocommitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autocommit",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetBaseLocationPrefix(val *string) {
	if err := j.validateSetBaseLocationPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseLocationPrefix",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetBinaryInputFormat(val *string) {
	if err := j.validateSetBinaryInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryInputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetBinaryOutputFormat(val *string) {
	if err := j.validateSetBinaryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetCatalogSync(val *string) {
	if err := j.validateSetCatalogSyncParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogSync",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientEnableLogInfoStatementParameters(val interface{}) {
	if err := j.validateSetClientEnableLogInfoStatementParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientEnableLogInfoStatementParameters",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientEncryptionKeySize(val *float64) {
	if err := j.validateSetClientEncryptionKeySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientEncryptionKeySize",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientMemoryLimit(val *float64) {
	if err := j.validateSetClientMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMemoryLimit",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientMetadataRequestUseConnectionCtx(val interface{}) {
	if err := j.validateSetClientMetadataRequestUseConnectionCtxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadataRequestUseConnectionCtx",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientMetadataUseSessionDatabase(val interface{}) {
	if err := j.validateSetClientMetadataUseSessionDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadataUseSessionDatabase",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientPrefetchThreads(val *float64) {
	if err := j.validateSetClientPrefetchThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientPrefetchThreads",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientResultChunkSize(val *float64) {
	if err := j.validateSetClientResultChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultChunkSize",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientResultColumnCaseInsensitive(val interface{}) {
	if err := j.validateSetClientResultColumnCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultColumnCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientSessionKeepAlive(val interface{}) {
	if err := j.validateSetClientSessionKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAlive",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientSessionKeepAliveHeartbeatFrequency(val *float64) {
	if err := j.validateSetClientSessionKeepAliveHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetClientTimestampTypeMapping(val *string) {
	if err := j.validateSetClientTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTimestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetCortexEnabledCrossRegion(val *string) {
	if err := j.validateSetCortexEnabledCrossRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cortexEnabledCrossRegion",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetCortexModelsAllowlist(val *string) {
	if err := j.validateSetCortexModelsAllowlistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cortexModelsAllowlist",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetCsvTimestampFormat(val *string) {
	if err := j.validateSetCsvTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvTimestampFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDataRetentionTimeInDays(val *float64) {
	if err := j.validateSetDataRetentionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRetentionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDateInputFormat(val *string) {
	if err := j.validateSetDateInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateInputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDateOutputFormat(val *string) {
	if err := j.validateSetDateOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDefaultDdlCollation(val *string) {
	if err := j.validateSetDefaultDdlCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDdlCollation",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDefaultNotebookComputePoolCpu(val *string) {
	if err := j.validateSetDefaultNotebookComputePoolCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNotebookComputePoolCpu",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDefaultNotebookComputePoolGpu(val *string) {
	if err := j.validateSetDefaultNotebookComputePoolGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNotebookComputePoolGpu",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDefaultNullOrdering(val *string) {
	if err := j.validateSetDefaultNullOrderingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNullOrdering",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDefaultStreamlitNotebookWarehouse(val *string) {
	if err := j.validateSetDefaultStreamlitNotebookWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStreamlitNotebookWarehouse",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDisableUiDownloadButton(val interface{}) {
	if err := j.validateSetDisableUiDownloadButtonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUiDownloadButton",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetDisableUserPrivilegeGrants(val interface{}) {
	if err := j.validateSetDisableUserPrivilegeGrantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUserPrivilegeGrants",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableAutomaticSensitiveDataClassificationLog(val interface{}) {
	if err := j.validateSetEnableAutomaticSensitiveDataClassificationLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticSensitiveDataClassificationLog",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableEgressCostOptimizer(val interface{}) {
	if err := j.validateSetEnableEgressCostOptimizerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEgressCostOptimizer",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableIdentifierFirstLogin(val interface{}) {
	if err := j.validateSetEnableIdentifierFirstLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIdentifierFirstLogin",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableInternalStagesPrivatelink(val interface{}) {
	if err := j.validateSetEnableInternalStagesPrivatelinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInternalStagesPrivatelink",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableTriSecretAndRekeyOptOutForImageRepository(val interface{}) {
	if err := j.validateSetEnableTriSecretAndRekeyOptOutForImageRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTriSecretAndRekeyOptOutForImageRepository",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableTriSecretAndRekeyOptOutForSpcsBlockStorage(val interface{}) {
	if err := j.validateSetEnableTriSecretAndRekeyOptOutForSpcsBlockStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTriSecretAndRekeyOptOutForSpcsBlockStorage",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableUnhandledExceptionsReporting(val interface{}) {
	if err := j.validateSetEnableUnhandledExceptionsReportingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnhandledExceptionsReporting",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableUnloadPhysicalTypeOptimization(val interface{}) {
	if err := j.validateSetEnableUnloadPhysicalTypeOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnloadPhysicalTypeOptimization",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableUnredactedQuerySyntaxError(val interface{}) {
	if err := j.validateSetEnableUnredactedQuerySyntaxErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnredactedQuerySyntaxError",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnableUnredactedSecureObjectError(val interface{}) {
	if err := j.validateSetEnableUnredactedSecureObjectErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnredactedSecureObjectError",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEnforceNetworkRulesForInternalStages(val interface{}) {
	if err := j.validateSetEnforceNetworkRulesForInternalStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceNetworkRulesForInternalStages",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetErrorOnNondeterministicMerge(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicMerge",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetErrorOnNondeterministicUpdate(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicUpdate",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetEventTable(val *string) {
	if err := j.validateSetEventTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventTable",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetExternalOauthAddPrivilegedRolesToBlockedList(val interface{}) {
	if err := j.validateSetExternalOauthAddPrivilegedRolesToBlockedListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalOauthAddPrivilegedRolesToBlockedList",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetExternalVolume(val *string) {
	if err := j.validateSetExternalVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalVolume",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetFeaturePolicy(val *string) {
	if err := j.validateSetFeaturePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featurePolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetGeographyOutputFormat(val *string) {
	if err := j.validateSetGeographyOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geographyOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetGeometryOutputFormat(val *string) {
	if err := j.validateSetGeometryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geometryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetHybridTableLockTimeout(val *float64) {
	if err := j.validateSetHybridTableLockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridTableLockTimeout",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetInitialReplicationSizeLimitInTb(val *string) {
	if err := j.validateSetInitialReplicationSizeLimitInTbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialReplicationSizeLimitInTb",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetJdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetJdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetJdbcTreatTimestampNtzAsUtc(val interface{}) {
	if err := j.validateSetJdbcTreatTimestampNtzAsUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetJdbcUseSessionTimezone(val interface{}) {
	if err := j.validateSetJdbcUseSessionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcUseSessionTimezone",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetJsonIndent(val *float64) {
	if err := j.validateSetJsonIndentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonIndent",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetJsTreatIntegerAsBigint(val interface{}) {
	if err := j.validateSetJsTreatIntegerAsBigintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsTreatIntegerAsBigint",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetListingAutoFulfillmentReplicationRefreshSchedule(val *string) {
	if err := j.validateSetListingAutoFulfillmentReplicationRefreshScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingAutoFulfillmentReplicationRefreshSchedule",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetLockTimeout(val *float64) {
	if err := j.validateSetLockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockTimeout",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetMaxConcurrencyLevel(val *float64) {
	if err := j.validateSetMaxConcurrencyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetMaxDataExtensionTimeInDays(val *float64) {
	if err := j.validateSetMaxDataExtensionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDataExtensionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetMetricLevel(val *string) {
	if err := j.validateSetMetricLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetMinDataRetentionTimeInDays(val *float64) {
	if err := j.validateSetMinDataRetentionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDataRetentionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetMultiStatementCount(val *float64) {
	if err := j.validateSetMultiStatementCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiStatementCount",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetNetworkPolicy(val *string) {
	if err := j.validateSetNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetNoorderSequenceAsDefault(val interface{}) {
	if err := j.validateSetNoorderSequenceAsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noorderSequenceAsDefault",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetOauthAddPrivilegedRolesToBlockedList(val interface{}) {
	if err := j.validateSetOauthAddPrivilegedRolesToBlockedListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAddPrivilegedRolesToBlockedList",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetOdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetOdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetPackagesPolicy(val *string) {
	if err := j.validateSetPackagesPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packagesPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetPasswordPolicy(val *string) {
	if err := j.validateSetPasswordPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetPeriodicDataRekeying(val interface{}) {
	if err := j.validateSetPeriodicDataRekeyingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodicDataRekeying",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetPipeExecutionPaused(val interface{}) {
	if err := j.validateSetPipeExecutionPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipeExecutionPaused",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetPreventUnloadToInlineUrl(val interface{}) {
	if err := j.validateSetPreventUnloadToInlineUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUnloadToInlineUrl",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetPreventUnloadToInternalStages(val interface{}) {
	if err := j.validateSetPreventUnloadToInternalStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUnloadToInternalStages",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetPythonProfilerModules(val *string) {
	if err := j.validateSetPythonProfilerModulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonProfilerModules",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetPythonProfilerTargetStage(val *string) {
	if err := j.validateSetPythonProfilerTargetStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonProfilerTargetStage",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetQueryTag(val *string) {
	if err := j.validateSetQueryTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryTag",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetQuotedIdentifiersIgnoreCase(val interface{}) {
	if err := j.validateSetQuotedIdentifiersIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotedIdentifiersIgnoreCase",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetReplaceInvalidCharacters(val interface{}) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetRequireStorageIntegrationForStageCreation(val interface{}) {
	if err := j.validateSetRequireStorageIntegrationForStageCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireStorageIntegrationForStageCreation",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetRequireStorageIntegrationForStageOperation(val interface{}) {
	if err := j.validateSetRequireStorageIntegrationForStageOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireStorageIntegrationForStageOperation",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetResourceMonitor(val *string) {
	if err := j.validateSetResourceMonitorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceMonitor",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetRowsPerResultset(val *float64) {
	if err := j.validateSetRowsPerResultsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsPerResultset",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetS3StageVpceDnsName(val *string) {
	if err := j.validateSetS3StageVpceDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3StageVpceDnsName",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetSamlIdentityProvider(val *string) {
	if err := j.validateSetSamlIdentityProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlIdentityProvider",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetSearchPath(val *string) {
	if err := j.validateSetSearchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchPath",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetServerlessTaskMaxStatementSize(val *string) {
	if err := j.validateSetServerlessTaskMaxStatementSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessTaskMaxStatementSize",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetServerlessTaskMinStatementSize(val *string) {
	if err := j.validateSetServerlessTaskMinStatementSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessTaskMinStatementSize",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetSessionPolicy(val *string) {
	if err := j.validateSetSessionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetSimulatedDataSharingConsumer(val *string) {
	if err := j.validateSetSimulatedDataSharingConsumerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"simulatedDataSharingConsumer",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetSsoLoginPage(val interface{}) {
	if err := j.validateSetSsoLoginPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoLoginPage",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetStatementQueuedTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementQueuedTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementQueuedTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetStatementTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetStorageSerializationPolicy(val *string) {
	if err := j.validateSetStorageSerializationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSerializationPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetStrictJsonOutput(val interface{}) {
	if err := j.validateSetStrictJsonOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictJsonOutput",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetSuspendTaskAfterNumFailures(val *float64) {
	if err := j.validateSetSuspendTaskAfterNumFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTaskAfterNumFailures",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTaskAutoRetryAttempts(val *float64) {
	if err := j.validateSetTaskAutoRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAutoRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimeInputFormat(val *string) {
	if err := j.validateSetTimeInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeInputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimeOutputFormat(val *string) {
	if err := j.validateSetTimeOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimestampDayIsAlways24H(val interface{}) {
	if err := j.validateSetTimestampDayIsAlways24HParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDayIsAlways24H",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimestampInputFormat(val *string) {
	if err := j.validateSetTimestampInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampInputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimestampLtzOutputFormat(val *string) {
	if err := j.validateSetTimestampLtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampLtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimestampNtzOutputFormat(val *string) {
	if err := j.validateSetTimestampNtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampNtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimestampOutputFormat(val *string) {
	if err := j.validateSetTimestampOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimestampTypeMapping(val *string) {
	if err := j.validateSetTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimestampTzOutputFormat(val *string) {
	if err := j.validateSetTimestampTzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTransactionAbortOnError(val interface{}) {
	if err := j.validateSetTransactionAbortOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionAbortOnError",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTransactionDefaultIsolationLevel(val *string) {
	if err := j.validateSetTransactionDefaultIsolationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionDefaultIsolationLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetTwoDigitCenturyStart(val *float64) {
	if err := j.validateSetTwoDigitCenturyStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoDigitCenturyStart",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetUnsupportedDdlAction(val *string) {
	if err := j.validateSetUnsupportedDdlActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unsupportedDdlAction",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetUseCachedResult(val interface{}) {
	if err := j.validateSetUseCachedResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCachedResult",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetUserTaskManagedInitialWarehouseSize(val *string) {
	if err := j.validateSetUserTaskManagedInitialWarehouseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskManagedInitialWarehouseSize",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetUserTaskMinimumTriggerIntervalInSeconds(val *float64) {
	if err := j.validateSetUserTaskMinimumTriggerIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetUserTaskTimeoutMs(val *float64) {
	if err := j.validateSetUserTaskTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetWeekOfYearPolicy(val *float64) {
	if err := j.validateSetWeekOfYearPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekOfYearPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentAccount)SetWeekStart(val *float64) {
	if err := j.validateSetWeekStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekStart",
		val,
	)
}

// Generates CDKTF code for importing a CurrentAccount resource upon running "cdktf plan <stack-name>".
func CurrentAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCurrentAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.currentAccount.CurrentAccount",
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
func CurrentAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurrentAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.currentAccount.CurrentAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CurrentAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurrentAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.currentAccount.CurrentAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CurrentAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurrentAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.currentAccount.CurrentAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CurrentAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.currentAccount.CurrentAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CurrentAccount) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CurrentAccount) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CurrentAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CurrentAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CurrentAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CurrentAccount) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CurrentAccount) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CurrentAccount) PutTimeouts(value *CurrentAccountTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CurrentAccount) ResetAbortDetachedQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetAbortDetachedQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetActivePythonProfiler() {
	_jsii_.InvokeVoid(
		c,
		"resetActivePythonProfiler",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetAllowClientMfaCaching() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowClientMfaCaching",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetAllowIdToken() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowIdToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetAuthenticationPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticationPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetAutocommit() {
	_jsii_.InvokeVoid(
		c,
		"resetAutocommit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetBaseLocationPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetBaseLocationPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetBinaryInputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetBinaryInputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetBinaryOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetBinaryOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetCatalog() {
	_jsii_.InvokeVoid(
		c,
		"resetCatalog",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetCatalogSync() {
	_jsii_.InvokeVoid(
		c,
		"resetCatalogSync",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientEnableLogInfoStatementParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetClientEnableLogInfoStatementParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientEncryptionKeySize() {
	_jsii_.InvokeVoid(
		c,
		"resetClientEncryptionKeySize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientMemoryLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetClientMemoryLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientMetadataRequestUseConnectionCtx() {
	_jsii_.InvokeVoid(
		c,
		"resetClientMetadataRequestUseConnectionCtx",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientMetadataUseSessionDatabase() {
	_jsii_.InvokeVoid(
		c,
		"resetClientMetadataUseSessionDatabase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientPrefetchThreads() {
	_jsii_.InvokeVoid(
		c,
		"resetClientPrefetchThreads",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientResultChunkSize() {
	_jsii_.InvokeVoid(
		c,
		"resetClientResultChunkSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientResultColumnCaseInsensitive() {
	_jsii_.InvokeVoid(
		c,
		"resetClientResultColumnCaseInsensitive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientSessionKeepAlive() {
	_jsii_.InvokeVoid(
		c,
		"resetClientSessionKeepAlive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientSessionKeepAliveHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetClientSessionKeepAliveHeartbeatFrequency",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetClientTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTimestampTypeMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetCortexEnabledCrossRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetCortexEnabledCrossRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetCortexModelsAllowlist() {
	_jsii_.InvokeVoid(
		c,
		"resetCortexModelsAllowlist",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetCsvTimestampFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetCsvTimestampFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDataRetentionTimeInDays() {
	_jsii_.InvokeVoid(
		c,
		"resetDataRetentionTimeInDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDateInputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetDateInputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDateOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetDateOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDefaultDdlCollation() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultDdlCollation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDefaultNotebookComputePoolCpu() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultNotebookComputePoolCpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDefaultNotebookComputePoolGpu() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultNotebookComputePoolGpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDefaultNullOrdering() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultNullOrdering",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDefaultStreamlitNotebookWarehouse() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultStreamlitNotebookWarehouse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDisableUiDownloadButton() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableUiDownloadButton",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetDisableUserPrivilegeGrants() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableUserPrivilegeGrants",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableAutomaticSensitiveDataClassificationLog() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableAutomaticSensitiveDataClassificationLog",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableEgressCostOptimizer() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableEgressCostOptimizer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableIdentifierFirstLogin() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableIdentifierFirstLogin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableInternalStagesPrivatelink() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableInternalStagesPrivatelink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableTriSecretAndRekeyOptOutForImageRepository() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableTriSecretAndRekeyOptOutForImageRepository",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableTriSecretAndRekeyOptOutForSpcsBlockStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableTriSecretAndRekeyOptOutForSpcsBlockStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableUnhandledExceptionsReporting() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableUnhandledExceptionsReporting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableUnloadPhysicalTypeOptimization() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableUnloadPhysicalTypeOptimization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableUnredactedQuerySyntaxError() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableUnredactedQuerySyntaxError",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnableUnredactedSecureObjectError() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableUnredactedSecureObjectError",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEnforceNetworkRulesForInternalStages() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforceNetworkRulesForInternalStages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetErrorOnNondeterministicMerge() {
	_jsii_.InvokeVoid(
		c,
		"resetErrorOnNondeterministicMerge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetErrorOnNondeterministicUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetErrorOnNondeterministicUpdate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetEventTable() {
	_jsii_.InvokeVoid(
		c,
		"resetEventTable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetExternalOauthAddPrivilegedRolesToBlockedList() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalOauthAddPrivilegedRolesToBlockedList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetExternalVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetFeaturePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetFeaturePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetGeographyOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetGeographyOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetGeometryOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetGeometryOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetHybridTableLockTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetHybridTableLockTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetInitialReplicationSizeLimitInTb() {
	_jsii_.InvokeVoid(
		c,
		"resetInitialReplicationSizeLimitInTb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetJdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		c,
		"resetJdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetJdbcTreatTimestampNtzAsUtc() {
	_jsii_.InvokeVoid(
		c,
		"resetJdbcTreatTimestampNtzAsUtc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetJdbcUseSessionTimezone() {
	_jsii_.InvokeVoid(
		c,
		"resetJdbcUseSessionTimezone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetJsonIndent() {
	_jsii_.InvokeVoid(
		c,
		"resetJsonIndent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetJsTreatIntegerAsBigint() {
	_jsii_.InvokeVoid(
		c,
		"resetJsTreatIntegerAsBigint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetListingAutoFulfillmentReplicationRefreshSchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetListingAutoFulfillmentReplicationRefreshSchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetLockTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetLockTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetLogLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetMaxConcurrencyLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConcurrencyLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetMaxDataExtensionTimeInDays() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxDataExtensionTimeInDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetMetricLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetMinDataRetentionTimeInDays() {
	_jsii_.InvokeVoid(
		c,
		"resetMinDataRetentionTimeInDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetMultiStatementCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMultiStatementCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetNoorderSequenceAsDefault() {
	_jsii_.InvokeVoid(
		c,
		"resetNoorderSequenceAsDefault",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetOauthAddPrivilegedRolesToBlockedList() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthAddPrivilegedRolesToBlockedList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetOdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		c,
		"resetOdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetPackagesPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetPackagesPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetPeriodicDataRekeying() {
	_jsii_.InvokeVoid(
		c,
		"resetPeriodicDataRekeying",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetPipeExecutionPaused() {
	_jsii_.InvokeVoid(
		c,
		"resetPipeExecutionPaused",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetPreventUnloadToInlineUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetPreventUnloadToInlineUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetPreventUnloadToInternalStages() {
	_jsii_.InvokeVoid(
		c,
		"resetPreventUnloadToInternalStages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetPythonProfilerModules() {
	_jsii_.InvokeVoid(
		c,
		"resetPythonProfilerModules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetPythonProfilerTargetStage() {
	_jsii_.InvokeVoid(
		c,
		"resetPythonProfilerTargetStage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetQueryTag() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryTag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetQuotedIdentifiersIgnoreCase() {
	_jsii_.InvokeVoid(
		c,
		"resetQuotedIdentifiersIgnoreCase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		c,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetRequireStorageIntegrationForStageCreation() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireStorageIntegrationForStageCreation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetRequireStorageIntegrationForStageOperation() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireStorageIntegrationForStageOperation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetResourceMonitor() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceMonitor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetRowsPerResultset() {
	_jsii_.InvokeVoid(
		c,
		"resetRowsPerResultset",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetS3StageVpceDnsName() {
	_jsii_.InvokeVoid(
		c,
		"resetS3StageVpceDnsName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetSamlIdentityProvider() {
	_jsii_.InvokeVoid(
		c,
		"resetSamlIdentityProvider",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetSearchPath() {
	_jsii_.InvokeVoid(
		c,
		"resetSearchPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetServerlessTaskMaxStatementSize() {
	_jsii_.InvokeVoid(
		c,
		"resetServerlessTaskMaxStatementSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetServerlessTaskMinStatementSize() {
	_jsii_.InvokeVoid(
		c,
		"resetServerlessTaskMinStatementSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetSessionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetSessionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetSimulatedDataSharingConsumer() {
	_jsii_.InvokeVoid(
		c,
		"resetSimulatedDataSharingConsumer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetSsoLoginPage() {
	_jsii_.InvokeVoid(
		c,
		"resetSsoLoginPage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetStatementQueuedTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetStatementQueuedTimeoutInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetStatementTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetStatementTimeoutInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetStorageSerializationPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageSerializationPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetStrictJsonOutput() {
	_jsii_.InvokeVoid(
		c,
		"resetStrictJsonOutput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetSuspendTaskAfterNumFailures() {
	_jsii_.InvokeVoid(
		c,
		"resetSuspendTaskAfterNumFailures",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTaskAutoRetryAttempts() {
	_jsii_.InvokeVoid(
		c,
		"resetTaskAutoRetryAttempts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimeInputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeInputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimeOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimestampDayIsAlways24H() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampDayIsAlways24H",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimestampInputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampInputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimestampLtzOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampLtzOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimestampNtzOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampNtzOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimestampOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampTypeMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimestampTzOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampTzOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTimezone() {
	_jsii_.InvokeVoid(
		c,
		"resetTimezone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTransactionAbortOnError() {
	_jsii_.InvokeVoid(
		c,
		"resetTransactionAbortOnError",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTransactionDefaultIsolationLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetTransactionDefaultIsolationLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetTwoDigitCenturyStart() {
	_jsii_.InvokeVoid(
		c,
		"resetTwoDigitCenturyStart",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetUnsupportedDdlAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUnsupportedDdlAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetUseCachedResult() {
	_jsii_.InvokeVoid(
		c,
		"resetUseCachedResult",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetUserTaskManagedInitialWarehouseSize() {
	_jsii_.InvokeVoid(
		c,
		"resetUserTaskManagedInitialWarehouseSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetUserTaskMinimumTriggerIntervalInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetUserTaskMinimumTriggerIntervalInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetUserTaskTimeoutMs() {
	_jsii_.InvokeVoid(
		c,
		"resetUserTaskTimeoutMs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetWeekOfYearPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetWeekOfYearPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) ResetWeekStart() {
	_jsii_.InvokeVoid(
		c,
		"resetWeekStart",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

