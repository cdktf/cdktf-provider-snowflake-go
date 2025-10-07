// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package currentorganizationaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/currentorganizationaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/current_organization_account snowflake_current_organization_account}.
type CurrentOrganizationAccount interface {
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
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ShowOutput() CurrentOrganizationAccountShowOutputList
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
	Timeouts() CurrentOrganizationAccountTimeoutsOutputReference
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
	PutTimeouts(value *CurrentOrganizationAccountTimeouts)
	ResetAbortDetachedQuery()
	ResetActivePythonProfiler()
	ResetAllowClientMfaCaching()
	ResetAllowIdToken()
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
	ResetComment()
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

// The jsii proxy struct for CurrentOrganizationAccount
type jsiiProxy_CurrentOrganizationAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CurrentOrganizationAccount) AbortDetachedQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) AbortDetachedQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ActivePythonProfiler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activePythonProfiler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ActivePythonProfilerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activePythonProfilerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) AllowClientMfaCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClientMfaCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) AllowClientMfaCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClientMfaCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) AllowIdToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowIdToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) AllowIdTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowIdTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Autocommit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) AutocommitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) BaseLocationPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseLocationPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) BaseLocationPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseLocationPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) BinaryInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) BinaryInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) BinaryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) BinaryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CatalogSync() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CatalogSyncInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientEnableLogInfoStatementParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientEnableLogInfoStatementParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientEnableLogInfoStatementParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientEnableLogInfoStatementParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientEncryptionKeySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientEncryptionKeySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientEncryptionKeySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientEncryptionKeySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientMemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientMemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientMetadataRequestUseConnectionCtx() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientMetadataRequestUseConnectionCtxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientMetadataUseSessionDatabase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataUseSessionDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientMetadataUseSessionDatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataUseSessionDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientPrefetchThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientPrefetchThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientResultChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientResultChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientResultColumnCaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientResultColumnCaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientSessionKeepAlive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientSessionKeepAliveHeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientSessionKeepAliveHeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientSessionKeepAliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientTimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ClientTimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CortexEnabledCrossRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cortexEnabledCrossRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CortexEnabledCrossRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cortexEnabledCrossRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CortexModelsAllowlist() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cortexModelsAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CortexModelsAllowlistInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cortexModelsAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CsvTimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvTimestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) CsvTimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvTimestampFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DataRetentionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DataRetentionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DateInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DateInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DateOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DateOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultDdlCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDdlCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultDdlCollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDdlCollationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultNotebookComputePoolCpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultNotebookComputePoolCpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultNotebookComputePoolGpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultNotebookComputePoolGpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolGpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultNullOrdering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNullOrdering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultNullOrderingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNullOrderingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultStreamlitNotebookWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStreamlitNotebookWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DefaultStreamlitNotebookWarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStreamlitNotebookWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DisableUiDownloadButton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUiDownloadButton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DisableUiDownloadButtonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUiDownloadButtonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DisableUserPrivilegeGrants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUserPrivilegeGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) DisableUserPrivilegeGrantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUserPrivilegeGrantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableAutomaticSensitiveDataClassificationLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticSensitiveDataClassificationLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableAutomaticSensitiveDataClassificationLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticSensitiveDataClassificationLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableEgressCostOptimizer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEgressCostOptimizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableEgressCostOptimizerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEgressCostOptimizerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableIdentifierFirstLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIdentifierFirstLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableIdentifierFirstLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIdentifierFirstLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableInternalStagesPrivatelink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternalStagesPrivatelink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableInternalStagesPrivatelinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternalStagesPrivatelinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableTriSecretAndRekeyOptOutForImageRepository() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForImageRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableTriSecretAndRekeyOptOutForImageRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForImageRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableTriSecretAndRekeyOptOutForSpcsBlockStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForSpcsBlockStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableTriSecretAndRekeyOptOutForSpcsBlockStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForSpcsBlockStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableUnhandledExceptionsReporting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnhandledExceptionsReporting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableUnhandledExceptionsReportingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnhandledExceptionsReportingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableUnloadPhysicalTypeOptimization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableUnloadPhysicalTypeOptimizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableUnredactedQuerySyntaxError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableUnredactedQuerySyntaxErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableUnredactedSecureObjectError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedSecureObjectError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnableUnredactedSecureObjectErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedSecureObjectErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnforceNetworkRulesForInternalStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceNetworkRulesForInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EnforceNetworkRulesForInternalStagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceNetworkRulesForInternalStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ErrorOnNondeterministicMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ErrorOnNondeterministicMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ErrorOnNondeterministicUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ErrorOnNondeterministicUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EventTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) EventTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ExternalOauthAddPrivilegedRolesToBlockedList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalOauthAddPrivilegedRolesToBlockedList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ExternalOauthAddPrivilegedRolesToBlockedListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalOauthAddPrivilegedRolesToBlockedListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ExternalVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ExternalVolumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) GeographyOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) GeographyOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) GeometryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) GeometryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) HybridTableLockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hybridTableLockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) HybridTableLockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hybridTableLockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) InitialReplicationSizeLimitInTb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialReplicationSizeLimitInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) InitialReplicationSizeLimitInTbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialReplicationSizeLimitInTbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JdbcTreatTimestampNtzAsUtc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JdbcTreatTimestampNtzAsUtcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JdbcUseSessionTimezone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JdbcUseSessionTimezoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JsonIndent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JsonIndentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JsTreatIntegerAsBigint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsTreatIntegerAsBigint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) JsTreatIntegerAsBigintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsTreatIntegerAsBigintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ListingAutoFulfillmentReplicationRefreshSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingAutoFulfillmentReplicationRefreshSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ListingAutoFulfillmentReplicationRefreshScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingAutoFulfillmentReplicationRefreshScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) LockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) LockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MaxConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MaxConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MaxDataExtensionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MaxDataExtensionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MetricLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MetricLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MinDataRetentionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MinDataRetentionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDataRetentionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MultiStatementCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) MultiStatementCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) NetworkPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) NetworkPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) NoorderSequenceAsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) NoorderSequenceAsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) OauthAddPrivilegedRolesToBlockedList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthAddPrivilegedRolesToBlockedList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) OauthAddPrivilegedRolesToBlockedListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthAddPrivilegedRolesToBlockedListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) OdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) OdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PasswordPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PasswordPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PeriodicDataRekeying() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"periodicDataRekeying",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PeriodicDataRekeyingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"periodicDataRekeyingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PipeExecutionPaused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipeExecutionPaused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PipeExecutionPausedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipeExecutionPausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PreventUnloadToInlineUrl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInlineUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PreventUnloadToInlineUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInlineUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PreventUnloadToInternalStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PreventUnloadToInternalStagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PythonProfilerModules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProfilerModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PythonProfilerModulesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProfilerModulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PythonProfilerTargetStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProfilerTargetStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) PythonProfilerTargetStageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProfilerTargetStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) QueryTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) QueryTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) QuotedIdentifiersIgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) QuotedIdentifiersIgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ReplaceInvalidCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ReplaceInvalidCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) RequireStorageIntegrationForStageCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) RequireStorageIntegrationForStageCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) RequireStorageIntegrationForStageOperation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) RequireStorageIntegrationForStageOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ResourceMonitor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ResourceMonitorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) RowsPerResultset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) RowsPerResultsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) S3StageVpceDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) S3StageVpceDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SamlIdentityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlIdentityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SamlIdentityProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlIdentityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SearchPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SearchPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ServerlessTaskMaxStatementSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessTaskMaxStatementSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ServerlessTaskMaxStatementSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessTaskMaxStatementSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ServerlessTaskMinStatementSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessTaskMinStatementSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ServerlessTaskMinStatementSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessTaskMinStatementSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SessionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SessionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) ShowOutput() CurrentOrganizationAccountShowOutputList {
	var returns CurrentOrganizationAccountShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SimulatedDataSharingConsumer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SimulatedDataSharingConsumerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SsoLoginPage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssoLoginPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SsoLoginPageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssoLoginPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) StatementQueuedTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) StatementQueuedTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) StatementTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) StatementTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) StorageSerializationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) StorageSerializationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) StrictJsonOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) StrictJsonOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SuspendTaskAfterNumFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) SuspendTaskAfterNumFailuresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TaskAutoRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TaskAutoRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimeInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimeInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimeOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimeOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Timeouts() CurrentOrganizationAccountTimeoutsOutputReference {
	var returns CurrentOrganizationAccountTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampDayIsAlways24H() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampDayIsAlways24HInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24HInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampLtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampLtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampNtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampNtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampTzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimestampTzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TransactionAbortOnError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TransactionAbortOnErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TransactionDefaultIsolationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TransactionDefaultIsolationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TwoDigitCenturyStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) TwoDigitCenturyStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UnsupportedDdlAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UnsupportedDdlActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UseCachedResult() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UseCachedResultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UserTaskManagedInitialWarehouseSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UserTaskManagedInitialWarehouseSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UserTaskMinimumTriggerIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UserTaskMinimumTriggerIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UserTaskTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) UserTaskTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) WeekOfYearPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) WeekOfYearPolicyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) WeekStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccount) WeekStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStartInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/current_organization_account snowflake_current_organization_account} Resource.
func NewCurrentOrganizationAccount(scope constructs.Construct, id *string, config *CurrentOrganizationAccountConfig) CurrentOrganizationAccount {
	_init_.Initialize()

	if err := validateNewCurrentOrganizationAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CurrentOrganizationAccount{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/current_organization_account snowflake_current_organization_account} Resource.
func NewCurrentOrganizationAccount_Override(c CurrentOrganizationAccount, scope constructs.Construct, id *string, config *CurrentOrganizationAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccount",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetAbortDetachedQuery(val interface{}) {
	if err := j.validateSetAbortDetachedQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortDetachedQuery",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetActivePythonProfiler(val *string) {
	if err := j.validateSetActivePythonProfilerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activePythonProfiler",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetAllowClientMfaCaching(val interface{}) {
	if err := j.validateSetAllowClientMfaCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClientMfaCaching",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetAllowIdToken(val interface{}) {
	if err := j.validateSetAllowIdTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowIdToken",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetAutocommit(val interface{}) {
	if err := j.validateSetAutocommitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autocommit",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetBaseLocationPrefix(val *string) {
	if err := j.validateSetBaseLocationPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseLocationPrefix",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetBinaryInputFormat(val *string) {
	if err := j.validateSetBinaryInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryInputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetBinaryOutputFormat(val *string) {
	if err := j.validateSetBinaryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetCatalogSync(val *string) {
	if err := j.validateSetCatalogSyncParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogSync",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientEnableLogInfoStatementParameters(val interface{}) {
	if err := j.validateSetClientEnableLogInfoStatementParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientEnableLogInfoStatementParameters",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientEncryptionKeySize(val *float64) {
	if err := j.validateSetClientEncryptionKeySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientEncryptionKeySize",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientMemoryLimit(val *float64) {
	if err := j.validateSetClientMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMemoryLimit",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientMetadataRequestUseConnectionCtx(val interface{}) {
	if err := j.validateSetClientMetadataRequestUseConnectionCtxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadataRequestUseConnectionCtx",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientMetadataUseSessionDatabase(val interface{}) {
	if err := j.validateSetClientMetadataUseSessionDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadataUseSessionDatabase",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientPrefetchThreads(val *float64) {
	if err := j.validateSetClientPrefetchThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientPrefetchThreads",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientResultChunkSize(val *float64) {
	if err := j.validateSetClientResultChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultChunkSize",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientResultColumnCaseInsensitive(val interface{}) {
	if err := j.validateSetClientResultColumnCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultColumnCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientSessionKeepAlive(val interface{}) {
	if err := j.validateSetClientSessionKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAlive",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientSessionKeepAliveHeartbeatFrequency(val *float64) {
	if err := j.validateSetClientSessionKeepAliveHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetClientTimestampTypeMapping(val *string) {
	if err := j.validateSetClientTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTimestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetCortexEnabledCrossRegion(val *string) {
	if err := j.validateSetCortexEnabledCrossRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cortexEnabledCrossRegion",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetCortexModelsAllowlist(val *string) {
	if err := j.validateSetCortexModelsAllowlistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cortexModelsAllowlist",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetCsvTimestampFormat(val *string) {
	if err := j.validateSetCsvTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvTimestampFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDataRetentionTimeInDays(val *float64) {
	if err := j.validateSetDataRetentionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRetentionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDateInputFormat(val *string) {
	if err := j.validateSetDateInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateInputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDateOutputFormat(val *string) {
	if err := j.validateSetDateOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDefaultDdlCollation(val *string) {
	if err := j.validateSetDefaultDdlCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDdlCollation",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDefaultNotebookComputePoolCpu(val *string) {
	if err := j.validateSetDefaultNotebookComputePoolCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNotebookComputePoolCpu",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDefaultNotebookComputePoolGpu(val *string) {
	if err := j.validateSetDefaultNotebookComputePoolGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNotebookComputePoolGpu",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDefaultNullOrdering(val *string) {
	if err := j.validateSetDefaultNullOrderingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNullOrdering",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDefaultStreamlitNotebookWarehouse(val *string) {
	if err := j.validateSetDefaultStreamlitNotebookWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStreamlitNotebookWarehouse",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDisableUiDownloadButton(val interface{}) {
	if err := j.validateSetDisableUiDownloadButtonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUiDownloadButton",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetDisableUserPrivilegeGrants(val interface{}) {
	if err := j.validateSetDisableUserPrivilegeGrantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUserPrivilegeGrants",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableAutomaticSensitiveDataClassificationLog(val interface{}) {
	if err := j.validateSetEnableAutomaticSensitiveDataClassificationLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticSensitiveDataClassificationLog",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableEgressCostOptimizer(val interface{}) {
	if err := j.validateSetEnableEgressCostOptimizerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEgressCostOptimizer",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableIdentifierFirstLogin(val interface{}) {
	if err := j.validateSetEnableIdentifierFirstLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIdentifierFirstLogin",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableInternalStagesPrivatelink(val interface{}) {
	if err := j.validateSetEnableInternalStagesPrivatelinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInternalStagesPrivatelink",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableTriSecretAndRekeyOptOutForImageRepository(val interface{}) {
	if err := j.validateSetEnableTriSecretAndRekeyOptOutForImageRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTriSecretAndRekeyOptOutForImageRepository",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableTriSecretAndRekeyOptOutForSpcsBlockStorage(val interface{}) {
	if err := j.validateSetEnableTriSecretAndRekeyOptOutForSpcsBlockStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTriSecretAndRekeyOptOutForSpcsBlockStorage",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableUnhandledExceptionsReporting(val interface{}) {
	if err := j.validateSetEnableUnhandledExceptionsReportingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnhandledExceptionsReporting",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableUnloadPhysicalTypeOptimization(val interface{}) {
	if err := j.validateSetEnableUnloadPhysicalTypeOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnloadPhysicalTypeOptimization",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableUnredactedQuerySyntaxError(val interface{}) {
	if err := j.validateSetEnableUnredactedQuerySyntaxErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnredactedQuerySyntaxError",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnableUnredactedSecureObjectError(val interface{}) {
	if err := j.validateSetEnableUnredactedSecureObjectErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnredactedSecureObjectError",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEnforceNetworkRulesForInternalStages(val interface{}) {
	if err := j.validateSetEnforceNetworkRulesForInternalStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceNetworkRulesForInternalStages",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetErrorOnNondeterministicMerge(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicMerge",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetErrorOnNondeterministicUpdate(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicUpdate",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetEventTable(val *string) {
	if err := j.validateSetEventTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventTable",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetExternalOauthAddPrivilegedRolesToBlockedList(val interface{}) {
	if err := j.validateSetExternalOauthAddPrivilegedRolesToBlockedListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalOauthAddPrivilegedRolesToBlockedList",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetExternalVolume(val *string) {
	if err := j.validateSetExternalVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalVolume",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetGeographyOutputFormat(val *string) {
	if err := j.validateSetGeographyOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geographyOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetGeometryOutputFormat(val *string) {
	if err := j.validateSetGeometryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geometryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetHybridTableLockTimeout(val *float64) {
	if err := j.validateSetHybridTableLockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridTableLockTimeout",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetInitialReplicationSizeLimitInTb(val *string) {
	if err := j.validateSetInitialReplicationSizeLimitInTbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialReplicationSizeLimitInTb",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetJdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetJdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetJdbcTreatTimestampNtzAsUtc(val interface{}) {
	if err := j.validateSetJdbcTreatTimestampNtzAsUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetJdbcUseSessionTimezone(val interface{}) {
	if err := j.validateSetJdbcUseSessionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcUseSessionTimezone",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetJsonIndent(val *float64) {
	if err := j.validateSetJsonIndentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonIndent",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetJsTreatIntegerAsBigint(val interface{}) {
	if err := j.validateSetJsTreatIntegerAsBigintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsTreatIntegerAsBigint",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetListingAutoFulfillmentReplicationRefreshSchedule(val *string) {
	if err := j.validateSetListingAutoFulfillmentReplicationRefreshScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingAutoFulfillmentReplicationRefreshSchedule",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetLockTimeout(val *float64) {
	if err := j.validateSetLockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockTimeout",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetMaxConcurrencyLevel(val *float64) {
	if err := j.validateSetMaxConcurrencyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetMaxDataExtensionTimeInDays(val *float64) {
	if err := j.validateSetMaxDataExtensionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDataExtensionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetMetricLevel(val *string) {
	if err := j.validateSetMetricLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetMinDataRetentionTimeInDays(val *float64) {
	if err := j.validateSetMinDataRetentionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDataRetentionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetMultiStatementCount(val *float64) {
	if err := j.validateSetMultiStatementCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiStatementCount",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetNetworkPolicy(val *string) {
	if err := j.validateSetNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetNoorderSequenceAsDefault(val interface{}) {
	if err := j.validateSetNoorderSequenceAsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noorderSequenceAsDefault",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetOauthAddPrivilegedRolesToBlockedList(val interface{}) {
	if err := j.validateSetOauthAddPrivilegedRolesToBlockedListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAddPrivilegedRolesToBlockedList",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetOdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetOdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetPasswordPolicy(val *string) {
	if err := j.validateSetPasswordPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetPeriodicDataRekeying(val interface{}) {
	if err := j.validateSetPeriodicDataRekeyingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodicDataRekeying",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetPipeExecutionPaused(val interface{}) {
	if err := j.validateSetPipeExecutionPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipeExecutionPaused",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetPreventUnloadToInlineUrl(val interface{}) {
	if err := j.validateSetPreventUnloadToInlineUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUnloadToInlineUrl",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetPreventUnloadToInternalStages(val interface{}) {
	if err := j.validateSetPreventUnloadToInternalStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUnloadToInternalStages",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetPythonProfilerModules(val *string) {
	if err := j.validateSetPythonProfilerModulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonProfilerModules",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetPythonProfilerTargetStage(val *string) {
	if err := j.validateSetPythonProfilerTargetStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonProfilerTargetStage",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetQueryTag(val *string) {
	if err := j.validateSetQueryTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryTag",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetQuotedIdentifiersIgnoreCase(val interface{}) {
	if err := j.validateSetQuotedIdentifiersIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotedIdentifiersIgnoreCase",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetReplaceInvalidCharacters(val interface{}) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetRequireStorageIntegrationForStageCreation(val interface{}) {
	if err := j.validateSetRequireStorageIntegrationForStageCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireStorageIntegrationForStageCreation",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetRequireStorageIntegrationForStageOperation(val interface{}) {
	if err := j.validateSetRequireStorageIntegrationForStageOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireStorageIntegrationForStageOperation",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetResourceMonitor(val *string) {
	if err := j.validateSetResourceMonitorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceMonitor",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetRowsPerResultset(val *float64) {
	if err := j.validateSetRowsPerResultsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsPerResultset",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetS3StageVpceDnsName(val *string) {
	if err := j.validateSetS3StageVpceDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3StageVpceDnsName",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetSamlIdentityProvider(val *string) {
	if err := j.validateSetSamlIdentityProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlIdentityProvider",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetSearchPath(val *string) {
	if err := j.validateSetSearchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchPath",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetServerlessTaskMaxStatementSize(val *string) {
	if err := j.validateSetServerlessTaskMaxStatementSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessTaskMaxStatementSize",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetServerlessTaskMinStatementSize(val *string) {
	if err := j.validateSetServerlessTaskMinStatementSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessTaskMinStatementSize",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetSessionPolicy(val *string) {
	if err := j.validateSetSessionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetSimulatedDataSharingConsumer(val *string) {
	if err := j.validateSetSimulatedDataSharingConsumerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"simulatedDataSharingConsumer",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetSsoLoginPage(val interface{}) {
	if err := j.validateSetSsoLoginPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoLoginPage",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetStatementQueuedTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementQueuedTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementQueuedTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetStatementTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetStorageSerializationPolicy(val *string) {
	if err := j.validateSetStorageSerializationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSerializationPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetStrictJsonOutput(val interface{}) {
	if err := j.validateSetStrictJsonOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictJsonOutput",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetSuspendTaskAfterNumFailures(val *float64) {
	if err := j.validateSetSuspendTaskAfterNumFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTaskAfterNumFailures",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTaskAutoRetryAttempts(val *float64) {
	if err := j.validateSetTaskAutoRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAutoRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimeInputFormat(val *string) {
	if err := j.validateSetTimeInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeInputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimeOutputFormat(val *string) {
	if err := j.validateSetTimeOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimestampDayIsAlways24H(val interface{}) {
	if err := j.validateSetTimestampDayIsAlways24HParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDayIsAlways24H",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimestampInputFormat(val *string) {
	if err := j.validateSetTimestampInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampInputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimestampLtzOutputFormat(val *string) {
	if err := j.validateSetTimestampLtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampLtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimestampNtzOutputFormat(val *string) {
	if err := j.validateSetTimestampNtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampNtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimestampOutputFormat(val *string) {
	if err := j.validateSetTimestampOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimestampTypeMapping(val *string) {
	if err := j.validateSetTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimestampTzOutputFormat(val *string) {
	if err := j.validateSetTimestampTzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTransactionAbortOnError(val interface{}) {
	if err := j.validateSetTransactionAbortOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionAbortOnError",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTransactionDefaultIsolationLevel(val *string) {
	if err := j.validateSetTransactionDefaultIsolationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionDefaultIsolationLevel",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetTwoDigitCenturyStart(val *float64) {
	if err := j.validateSetTwoDigitCenturyStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoDigitCenturyStart",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetUnsupportedDdlAction(val *string) {
	if err := j.validateSetUnsupportedDdlActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unsupportedDdlAction",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetUseCachedResult(val interface{}) {
	if err := j.validateSetUseCachedResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCachedResult",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetUserTaskManagedInitialWarehouseSize(val *string) {
	if err := j.validateSetUserTaskManagedInitialWarehouseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskManagedInitialWarehouseSize",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetUserTaskMinimumTriggerIntervalInSeconds(val *float64) {
	if err := j.validateSetUserTaskMinimumTriggerIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetUserTaskTimeoutMs(val *float64) {
	if err := j.validateSetUserTaskTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetWeekOfYearPolicy(val *float64) {
	if err := j.validateSetWeekOfYearPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekOfYearPolicy",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccount)SetWeekStart(val *float64) {
	if err := j.validateSetWeekStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekStart",
		val,
	)
}

// Generates CDKTF code for importing a CurrentOrganizationAccount resource upon running "cdktf plan <stack-name>".
func CurrentOrganizationAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCurrentOrganizationAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccount",
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
func CurrentOrganizationAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurrentOrganizationAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CurrentOrganizationAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurrentOrganizationAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CurrentOrganizationAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCurrentOrganizationAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CurrentOrganizationAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccount) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CurrentOrganizationAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CurrentOrganizationAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CurrentOrganizationAccount) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CurrentOrganizationAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CurrentOrganizationAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CurrentOrganizationAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CurrentOrganizationAccount) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CurrentOrganizationAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CurrentOrganizationAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CurrentOrganizationAccount) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) PutTimeouts(value *CurrentOrganizationAccountTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetAbortDetachedQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetAbortDetachedQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetActivePythonProfiler() {
	_jsii_.InvokeVoid(
		c,
		"resetActivePythonProfiler",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetAllowClientMfaCaching() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowClientMfaCaching",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetAllowIdToken() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowIdToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetAutocommit() {
	_jsii_.InvokeVoid(
		c,
		"resetAutocommit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetBaseLocationPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetBaseLocationPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetBinaryInputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetBinaryInputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetBinaryOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetBinaryOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetCatalog() {
	_jsii_.InvokeVoid(
		c,
		"resetCatalog",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetCatalogSync() {
	_jsii_.InvokeVoid(
		c,
		"resetCatalogSync",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientEnableLogInfoStatementParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetClientEnableLogInfoStatementParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientEncryptionKeySize() {
	_jsii_.InvokeVoid(
		c,
		"resetClientEncryptionKeySize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientMemoryLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetClientMemoryLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientMetadataRequestUseConnectionCtx() {
	_jsii_.InvokeVoid(
		c,
		"resetClientMetadataRequestUseConnectionCtx",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientMetadataUseSessionDatabase() {
	_jsii_.InvokeVoid(
		c,
		"resetClientMetadataUseSessionDatabase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientPrefetchThreads() {
	_jsii_.InvokeVoid(
		c,
		"resetClientPrefetchThreads",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientResultChunkSize() {
	_jsii_.InvokeVoid(
		c,
		"resetClientResultChunkSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientResultColumnCaseInsensitive() {
	_jsii_.InvokeVoid(
		c,
		"resetClientResultColumnCaseInsensitive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientSessionKeepAlive() {
	_jsii_.InvokeVoid(
		c,
		"resetClientSessionKeepAlive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientSessionKeepAliveHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetClientSessionKeepAliveHeartbeatFrequency",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetClientTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTimestampTypeMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetCortexEnabledCrossRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetCortexEnabledCrossRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetCortexModelsAllowlist() {
	_jsii_.InvokeVoid(
		c,
		"resetCortexModelsAllowlist",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetCsvTimestampFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetCsvTimestampFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDataRetentionTimeInDays() {
	_jsii_.InvokeVoid(
		c,
		"resetDataRetentionTimeInDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDateInputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetDateInputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDateOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetDateOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDefaultDdlCollation() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultDdlCollation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDefaultNotebookComputePoolCpu() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultNotebookComputePoolCpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDefaultNotebookComputePoolGpu() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultNotebookComputePoolGpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDefaultNullOrdering() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultNullOrdering",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDefaultStreamlitNotebookWarehouse() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultStreamlitNotebookWarehouse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDisableUiDownloadButton() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableUiDownloadButton",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetDisableUserPrivilegeGrants() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableUserPrivilegeGrants",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableAutomaticSensitiveDataClassificationLog() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableAutomaticSensitiveDataClassificationLog",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableEgressCostOptimizer() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableEgressCostOptimizer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableIdentifierFirstLogin() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableIdentifierFirstLogin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableInternalStagesPrivatelink() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableInternalStagesPrivatelink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableTriSecretAndRekeyOptOutForImageRepository() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableTriSecretAndRekeyOptOutForImageRepository",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableTriSecretAndRekeyOptOutForSpcsBlockStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableTriSecretAndRekeyOptOutForSpcsBlockStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableUnhandledExceptionsReporting() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableUnhandledExceptionsReporting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableUnloadPhysicalTypeOptimization() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableUnloadPhysicalTypeOptimization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableUnredactedQuerySyntaxError() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableUnredactedQuerySyntaxError",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnableUnredactedSecureObjectError() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableUnredactedSecureObjectError",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEnforceNetworkRulesForInternalStages() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforceNetworkRulesForInternalStages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetErrorOnNondeterministicMerge() {
	_jsii_.InvokeVoid(
		c,
		"resetErrorOnNondeterministicMerge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetErrorOnNondeterministicUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetErrorOnNondeterministicUpdate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetEventTable() {
	_jsii_.InvokeVoid(
		c,
		"resetEventTable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetExternalOauthAddPrivilegedRolesToBlockedList() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalOauthAddPrivilegedRolesToBlockedList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetExternalVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetGeographyOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetGeographyOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetGeometryOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetGeometryOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetHybridTableLockTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetHybridTableLockTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetInitialReplicationSizeLimitInTb() {
	_jsii_.InvokeVoid(
		c,
		"resetInitialReplicationSizeLimitInTb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetJdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		c,
		"resetJdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetJdbcTreatTimestampNtzAsUtc() {
	_jsii_.InvokeVoid(
		c,
		"resetJdbcTreatTimestampNtzAsUtc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetJdbcUseSessionTimezone() {
	_jsii_.InvokeVoid(
		c,
		"resetJdbcUseSessionTimezone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetJsonIndent() {
	_jsii_.InvokeVoid(
		c,
		"resetJsonIndent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetJsTreatIntegerAsBigint() {
	_jsii_.InvokeVoid(
		c,
		"resetJsTreatIntegerAsBigint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetListingAutoFulfillmentReplicationRefreshSchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetListingAutoFulfillmentReplicationRefreshSchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetLockTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetLockTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetLogLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetMaxConcurrencyLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConcurrencyLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetMaxDataExtensionTimeInDays() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxDataExtensionTimeInDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetMetricLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetMinDataRetentionTimeInDays() {
	_jsii_.InvokeVoid(
		c,
		"resetMinDataRetentionTimeInDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetMultiStatementCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMultiStatementCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetNoorderSequenceAsDefault() {
	_jsii_.InvokeVoid(
		c,
		"resetNoorderSequenceAsDefault",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetOauthAddPrivilegedRolesToBlockedList() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthAddPrivilegedRolesToBlockedList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetOdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		c,
		"resetOdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetPeriodicDataRekeying() {
	_jsii_.InvokeVoid(
		c,
		"resetPeriodicDataRekeying",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetPipeExecutionPaused() {
	_jsii_.InvokeVoid(
		c,
		"resetPipeExecutionPaused",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetPreventUnloadToInlineUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetPreventUnloadToInlineUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetPreventUnloadToInternalStages() {
	_jsii_.InvokeVoid(
		c,
		"resetPreventUnloadToInternalStages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetPythonProfilerModules() {
	_jsii_.InvokeVoid(
		c,
		"resetPythonProfilerModules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetPythonProfilerTargetStage() {
	_jsii_.InvokeVoid(
		c,
		"resetPythonProfilerTargetStage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetQueryTag() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryTag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetQuotedIdentifiersIgnoreCase() {
	_jsii_.InvokeVoid(
		c,
		"resetQuotedIdentifiersIgnoreCase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		c,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetRequireStorageIntegrationForStageCreation() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireStorageIntegrationForStageCreation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetRequireStorageIntegrationForStageOperation() {
	_jsii_.InvokeVoid(
		c,
		"resetRequireStorageIntegrationForStageOperation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetResourceMonitor() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceMonitor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetRowsPerResultset() {
	_jsii_.InvokeVoid(
		c,
		"resetRowsPerResultset",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetS3StageVpceDnsName() {
	_jsii_.InvokeVoid(
		c,
		"resetS3StageVpceDnsName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetSamlIdentityProvider() {
	_jsii_.InvokeVoid(
		c,
		"resetSamlIdentityProvider",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetSearchPath() {
	_jsii_.InvokeVoid(
		c,
		"resetSearchPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetServerlessTaskMaxStatementSize() {
	_jsii_.InvokeVoid(
		c,
		"resetServerlessTaskMaxStatementSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetServerlessTaskMinStatementSize() {
	_jsii_.InvokeVoid(
		c,
		"resetServerlessTaskMinStatementSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetSessionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetSessionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetSimulatedDataSharingConsumer() {
	_jsii_.InvokeVoid(
		c,
		"resetSimulatedDataSharingConsumer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetSsoLoginPage() {
	_jsii_.InvokeVoid(
		c,
		"resetSsoLoginPage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetStatementQueuedTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetStatementQueuedTimeoutInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetStatementTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetStatementTimeoutInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetStorageSerializationPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageSerializationPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetStrictJsonOutput() {
	_jsii_.InvokeVoid(
		c,
		"resetStrictJsonOutput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetSuspendTaskAfterNumFailures() {
	_jsii_.InvokeVoid(
		c,
		"resetSuspendTaskAfterNumFailures",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTaskAutoRetryAttempts() {
	_jsii_.InvokeVoid(
		c,
		"resetTaskAutoRetryAttempts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimeInputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeInputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimeOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimestampDayIsAlways24H() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampDayIsAlways24H",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimestampInputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampInputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimestampLtzOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampLtzOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimestampNtzOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampNtzOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimestampOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampTypeMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimestampTzOutputFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTimestampTzOutputFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTimezone() {
	_jsii_.InvokeVoid(
		c,
		"resetTimezone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTransactionAbortOnError() {
	_jsii_.InvokeVoid(
		c,
		"resetTransactionAbortOnError",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTransactionDefaultIsolationLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetTransactionDefaultIsolationLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetTwoDigitCenturyStart() {
	_jsii_.InvokeVoid(
		c,
		"resetTwoDigitCenturyStart",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetUnsupportedDdlAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUnsupportedDdlAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetUseCachedResult() {
	_jsii_.InvokeVoid(
		c,
		"resetUseCachedResult",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetUserTaskManagedInitialWarehouseSize() {
	_jsii_.InvokeVoid(
		c,
		"resetUserTaskManagedInitialWarehouseSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetUserTaskMinimumTriggerIntervalInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetUserTaskMinimumTriggerIntervalInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetUserTaskTimeoutMs() {
	_jsii_.InvokeVoid(
		c,
		"resetUserTaskTimeoutMs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetWeekOfYearPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetWeekOfYearPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) ResetWeekStart() {
	_jsii_.InvokeVoid(
		c,
		"resetWeekStart",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CurrentOrganizationAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

