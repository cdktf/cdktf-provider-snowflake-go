// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/serviceuser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/service_user snowflake_service_user}.
type ServiceUser interface {
	cdktf.TerraformResource
	AbortDetachedQuery() interface{}
	SetAbortDetachedQuery(val interface{})
	AbortDetachedQueryInput() interface{}
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
	DateInputFormat() *string
	SetDateInputFormat(val *string)
	DateInputFormatInput() *string
	DateOutputFormat() *string
	SetDateOutputFormat(val *string)
	DateOutputFormatInput() *string
	DaysToExpiry() *float64
	SetDaysToExpiry(val *float64)
	DaysToExpiryInput() *float64
	DefaultNamespace() *string
	SetDefaultNamespace(val *string)
	DefaultNamespaceInput() *string
	DefaultRole() *string
	SetDefaultRole(val *string)
	DefaultRoleInput() *string
	DefaultSecondaryRolesOption() *string
	SetDefaultSecondaryRolesOption(val *string)
	DefaultSecondaryRolesOptionInput() *string
	DefaultWarehouse() *string
	SetDefaultWarehouse(val *string)
	DefaultWarehouseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() *string
	SetDisabled(val *string)
	DisabledInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	EnableUnloadPhysicalTypeOptimization() interface{}
	SetEnableUnloadPhysicalTypeOptimization(val interface{})
	EnableUnloadPhysicalTypeOptimizationInput() interface{}
	EnableUnredactedQuerySyntaxError() interface{}
	SetEnableUnredactedQuerySyntaxError(val interface{})
	EnableUnredactedQuerySyntaxErrorInput() interface{}
	ErrorOnNondeterministicMerge() interface{}
	SetErrorOnNondeterministicMerge(val interface{})
	ErrorOnNondeterministicMergeInput() interface{}
	ErrorOnNondeterministicUpdate() interface{}
	SetErrorOnNondeterministicUpdate(val interface{})
	ErrorOnNondeterministicUpdateInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LockTimeout() *float64
	SetLockTimeout(val *float64)
	LockTimeoutInput() *float64
	LoginName() *string
	SetLoginName(val *string)
	LoginNameInput() *string
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	MinsToUnlock() *float64
	SetMinsToUnlock(val *float64)
	MinsToUnlockInput() *float64
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
	OdbcTreatDecimalAsInt() interface{}
	SetOdbcTreatDecimalAsInt(val interface{})
	OdbcTreatDecimalAsIntInput() interface{}
	Parameters() ServiceUserParametersList
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
	RsaPublicKey() *string
	SetRsaPublicKey(val *string)
	RsaPublicKey2() *string
	SetRsaPublicKey2(val *string)
	RsaPublicKey2Input() *string
	RsaPublicKeyInput() *string
	S3StageVpceDnsName() *string
	SetS3StageVpceDnsName(val *string)
	S3StageVpceDnsNameInput() *string
	SearchPath() *string
	SetSearchPath(val *string)
	SearchPathInput() *string
	ShowOutput() ServiceUserShowOutputList
	SimulatedDataSharingConsumer() *string
	SetSimulatedDataSharingConsumer(val *string)
	SimulatedDataSharingConsumerInput() *string
	StatementQueuedTimeoutInSeconds() *float64
	SetStatementQueuedTimeoutInSeconds(val *float64)
	StatementQueuedTimeoutInSecondsInput() *float64
	StatementTimeoutInSeconds() *float64
	SetStatementTimeoutInSeconds(val *float64)
	StatementTimeoutInSecondsInput() *float64
	StrictJsonOutput() interface{}
	SetStrictJsonOutput(val interface{})
	StrictJsonOutputInput() interface{}
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
	UserType() *string
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
	ResetAbortDetachedQuery()
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
	ResetDateInputFormat()
	ResetDateOutputFormat()
	ResetDaysToExpiry()
	ResetDefaultNamespace()
	ResetDefaultRole()
	ResetDefaultSecondaryRolesOption()
	ResetDefaultWarehouse()
	ResetDisabled()
	ResetDisplayName()
	ResetEmail()
	ResetEnableUnloadPhysicalTypeOptimization()
	ResetEnableUnredactedQuerySyntaxError()
	ResetErrorOnNondeterministicMerge()
	ResetErrorOnNondeterministicUpdate()
	ResetGeographyOutputFormat()
	ResetGeometryOutputFormat()
	ResetId()
	ResetJdbcTreatDecimalAsInt()
	ResetJdbcTreatTimestampNtzAsUtc()
	ResetJdbcUseSessionTimezone()
	ResetJsonIndent()
	ResetLockTimeout()
	ResetLoginName()
	ResetLogLevel()
	ResetMinsToUnlock()
	ResetMultiStatementCount()
	ResetNetworkPolicy()
	ResetNoorderSequenceAsDefault()
	ResetOdbcTreatDecimalAsInt()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreventUnloadToInternalStages()
	ResetQueryTag()
	ResetQuotedIdentifiersIgnoreCase()
	ResetRowsPerResultset()
	ResetRsaPublicKey()
	ResetRsaPublicKey2()
	ResetS3StageVpceDnsName()
	ResetSearchPath()
	ResetSimulatedDataSharingConsumer()
	ResetStatementQueuedTimeoutInSeconds()
	ResetStatementTimeoutInSeconds()
	ResetStrictJsonOutput()
	ResetTimeInputFormat()
	ResetTimeOutputFormat()
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

// The jsii proxy struct for ServiceUser
type jsiiProxy_ServiceUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServiceUser) AbortDetachedQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) AbortDetachedQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Autocommit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) AutocommitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) BinaryInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) BinaryInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) BinaryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) BinaryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientMemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientMemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientMetadataRequestUseConnectionCtx() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientMetadataRequestUseConnectionCtxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientPrefetchThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientPrefetchThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientResultChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientResultChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientResultColumnCaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientResultColumnCaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientSessionKeepAlive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientSessionKeepAliveHeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientSessionKeepAliveHeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientSessionKeepAliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientTimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ClientTimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DateInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DateInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DateOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DateOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DaysToExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DaysToExpiryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DefaultNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DefaultNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DefaultRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DefaultRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DefaultSecondaryRolesOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryRolesOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DefaultSecondaryRolesOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryRolesOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DefaultWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DefaultWarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Disabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DisabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) EnableUnloadPhysicalTypeOptimization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) EnableUnloadPhysicalTypeOptimizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) EnableUnredactedQuerySyntaxError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) EnableUnredactedQuerySyntaxErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ErrorOnNondeterministicMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ErrorOnNondeterministicMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ErrorOnNondeterministicUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ErrorOnNondeterministicUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) GeographyOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) GeographyOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) GeometryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) GeometryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) JdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) JdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) JdbcTreatTimestampNtzAsUtc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) JdbcTreatTimestampNtzAsUtcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) JdbcUseSessionTimezone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) JdbcUseSessionTimezoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) JsonIndent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) JsonIndentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) LockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) LockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) LoginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) LoginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) MinsToUnlock() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToUnlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) MinsToUnlockInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToUnlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) MultiStatementCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) MultiStatementCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) NetworkPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) NetworkPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) NoorderSequenceAsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) NoorderSequenceAsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) OdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) OdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Parameters() ServiceUserParametersList {
	var returns ServiceUserParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) PreventUnloadToInternalStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) PreventUnloadToInternalStagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) QueryTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) QueryTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) QuotedIdentifiersIgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) QuotedIdentifiersIgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) RowsPerResultset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) RowsPerResultsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) RsaPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) RsaPublicKey2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) RsaPublicKey2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) RsaPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) S3StageVpceDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) S3StageVpceDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) SearchPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) SearchPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) ShowOutput() ServiceUserShowOutputList {
	var returns ServiceUserShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) SimulatedDataSharingConsumer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) SimulatedDataSharingConsumerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) StatementQueuedTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) StatementQueuedTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) StatementTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) StatementTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) StrictJsonOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) StrictJsonOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimeInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimeInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimeOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimeOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampDayIsAlways24H() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampDayIsAlways24HInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24HInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampLtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampLtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampNtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampNtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampTzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimestampTzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TransactionAbortOnError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TransactionAbortOnErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TransactionDefaultIsolationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TransactionDefaultIsolationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TwoDigitCenturyStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) TwoDigitCenturyStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) UnsupportedDdlAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) UnsupportedDdlActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) UseCachedResult() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) UseCachedResultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) WeekOfYearPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) WeekOfYearPolicyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) WeekStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUser) WeekStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStartInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/service_user snowflake_service_user} Resource.
func NewServiceUser(scope constructs.Construct, id *string, config *ServiceUserConfig) ServiceUser {
	_init_.Initialize()

	if err := validateNewServiceUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceUser{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.serviceUser.ServiceUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/service_user snowflake_service_user} Resource.
func NewServiceUser_Override(s ServiceUser, scope constructs.Construct, id *string, config *ServiceUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.serviceUser.ServiceUser",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceUser)SetAbortDetachedQuery(val interface{}) {
	if err := j.validateSetAbortDetachedQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortDetachedQuery",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetAutocommit(val interface{}) {
	if err := j.validateSetAutocommitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autocommit",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetBinaryInputFormat(val *string) {
	if err := j.validateSetBinaryInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryInputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetBinaryOutputFormat(val *string) {
	if err := j.validateSetBinaryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetClientMemoryLimit(val *float64) {
	if err := j.validateSetClientMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMemoryLimit",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetClientMetadataRequestUseConnectionCtx(val interface{}) {
	if err := j.validateSetClientMetadataRequestUseConnectionCtxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadataRequestUseConnectionCtx",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetClientPrefetchThreads(val *float64) {
	if err := j.validateSetClientPrefetchThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientPrefetchThreads",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetClientResultChunkSize(val *float64) {
	if err := j.validateSetClientResultChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultChunkSize",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetClientResultColumnCaseInsensitive(val interface{}) {
	if err := j.validateSetClientResultColumnCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultColumnCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetClientSessionKeepAlive(val interface{}) {
	if err := j.validateSetClientSessionKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAlive",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetClientSessionKeepAliveHeartbeatFrequency(val *float64) {
	if err := j.validateSetClientSessionKeepAliveHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetClientTimestampTypeMapping(val *string) {
	if err := j.validateSetClientTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTimestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDateInputFormat(val *string) {
	if err := j.validateSetDateInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateInputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDateOutputFormat(val *string) {
	if err := j.validateSetDateOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDaysToExpiry(val *float64) {
	if err := j.validateSetDaysToExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysToExpiry",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDefaultNamespace(val *string) {
	if err := j.validateSetDefaultNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNamespace",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDefaultRole(val *string) {
	if err := j.validateSetDefaultRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRole",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDefaultSecondaryRolesOption(val *string) {
	if err := j.validateSetDefaultSecondaryRolesOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSecondaryRolesOption",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDefaultWarehouse(val *string) {
	if err := j.validateSetDefaultWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultWarehouse",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDisabled(val *string) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetEnableUnloadPhysicalTypeOptimization(val interface{}) {
	if err := j.validateSetEnableUnloadPhysicalTypeOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnloadPhysicalTypeOptimization",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetEnableUnredactedQuerySyntaxError(val interface{}) {
	if err := j.validateSetEnableUnredactedQuerySyntaxErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnredactedQuerySyntaxError",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetErrorOnNondeterministicMerge(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicMerge",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetErrorOnNondeterministicUpdate(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicUpdate",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetGeographyOutputFormat(val *string) {
	if err := j.validateSetGeographyOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geographyOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetGeometryOutputFormat(val *string) {
	if err := j.validateSetGeometryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geometryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetJdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetJdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetJdbcTreatTimestampNtzAsUtc(val interface{}) {
	if err := j.validateSetJdbcTreatTimestampNtzAsUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetJdbcUseSessionTimezone(val interface{}) {
	if err := j.validateSetJdbcUseSessionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcUseSessionTimezone",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetJsonIndent(val *float64) {
	if err := j.validateSetJsonIndentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonIndent",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetLockTimeout(val *float64) {
	if err := j.validateSetLockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockTimeout",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetLoginName(val *string) {
	if err := j.validateSetLoginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginName",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetMinsToUnlock(val *float64) {
	if err := j.validateSetMinsToUnlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minsToUnlock",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetMultiStatementCount(val *float64) {
	if err := j.validateSetMultiStatementCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiStatementCount",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetNetworkPolicy(val *string) {
	if err := j.validateSetNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetNoorderSequenceAsDefault(val interface{}) {
	if err := j.validateSetNoorderSequenceAsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noorderSequenceAsDefault",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetOdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetOdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetPreventUnloadToInternalStages(val interface{}) {
	if err := j.validateSetPreventUnloadToInternalStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUnloadToInternalStages",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetQueryTag(val *string) {
	if err := j.validateSetQueryTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryTag",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetQuotedIdentifiersIgnoreCase(val interface{}) {
	if err := j.validateSetQuotedIdentifiersIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotedIdentifiersIgnoreCase",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetRowsPerResultset(val *float64) {
	if err := j.validateSetRowsPerResultsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsPerResultset",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetRsaPublicKey(val *string) {
	if err := j.validateSetRsaPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetRsaPublicKey2(val *string) {
	if err := j.validateSetRsaPublicKey2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey2",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetS3StageVpceDnsName(val *string) {
	if err := j.validateSetS3StageVpceDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3StageVpceDnsName",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetSearchPath(val *string) {
	if err := j.validateSetSearchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchPath",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetSimulatedDataSharingConsumer(val *string) {
	if err := j.validateSetSimulatedDataSharingConsumerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"simulatedDataSharingConsumer",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetStatementQueuedTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementQueuedTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementQueuedTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetStatementTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetStrictJsonOutput(val interface{}) {
	if err := j.validateSetStrictJsonOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictJsonOutput",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimeInputFormat(val *string) {
	if err := j.validateSetTimeInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeInputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimeOutputFormat(val *string) {
	if err := j.validateSetTimeOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimestampDayIsAlways24H(val interface{}) {
	if err := j.validateSetTimestampDayIsAlways24HParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDayIsAlways24H",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimestampInputFormat(val *string) {
	if err := j.validateSetTimestampInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampInputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimestampLtzOutputFormat(val *string) {
	if err := j.validateSetTimestampLtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampLtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimestampNtzOutputFormat(val *string) {
	if err := j.validateSetTimestampNtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampNtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimestampOutputFormat(val *string) {
	if err := j.validateSetTimestampOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimestampTypeMapping(val *string) {
	if err := j.validateSetTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimestampTzOutputFormat(val *string) {
	if err := j.validateSetTimestampTzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTransactionAbortOnError(val interface{}) {
	if err := j.validateSetTransactionAbortOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionAbortOnError",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTransactionDefaultIsolationLevel(val *string) {
	if err := j.validateSetTransactionDefaultIsolationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionDefaultIsolationLevel",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetTwoDigitCenturyStart(val *float64) {
	if err := j.validateSetTwoDigitCenturyStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoDigitCenturyStart",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetUnsupportedDdlAction(val *string) {
	if err := j.validateSetUnsupportedDdlActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unsupportedDdlAction",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetUseCachedResult(val interface{}) {
	if err := j.validateSetUseCachedResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCachedResult",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetWeekOfYearPolicy(val *float64) {
	if err := j.validateSetWeekOfYearPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekOfYearPolicy",
		val,
	)
}

func (j *jsiiProxy_ServiceUser)SetWeekStart(val *float64) {
	if err := j.validateSetWeekStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekStart",
		val,
	)
}

// Generates CDKTF code for importing a ServiceUser resource upon running "cdktf plan <stack-name>".
func ServiceUser_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServiceUser_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.serviceUser.ServiceUser",
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
func ServiceUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.serviceUser.ServiceUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.serviceUser.ServiceUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceUser_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.serviceUser.ServiceUser",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServiceUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.serviceUser.ServiceUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServiceUser) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServiceUser) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceUser) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceUser) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceUser) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUser) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServiceUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUser) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceUser) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServiceUser) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceUser) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceUser) ResetAbortDetachedQuery() {
	_jsii_.InvokeVoid(
		s,
		"resetAbortDetachedQuery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetAutocommit() {
	_jsii_.InvokeVoid(
		s,
		"resetAutocommit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetBinaryInputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetBinaryInputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetBinaryOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetBinaryOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetClientMemoryLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetClientMemoryLimit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetClientMetadataRequestUseConnectionCtx() {
	_jsii_.InvokeVoid(
		s,
		"resetClientMetadataRequestUseConnectionCtx",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetClientPrefetchThreads() {
	_jsii_.InvokeVoid(
		s,
		"resetClientPrefetchThreads",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetClientResultChunkSize() {
	_jsii_.InvokeVoid(
		s,
		"resetClientResultChunkSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetClientResultColumnCaseInsensitive() {
	_jsii_.InvokeVoid(
		s,
		"resetClientResultColumnCaseInsensitive",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetClientSessionKeepAlive() {
	_jsii_.InvokeVoid(
		s,
		"resetClientSessionKeepAlive",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetClientSessionKeepAliveHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		s,
		"resetClientSessionKeepAliveHeartbeatFrequency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetClientTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		s,
		"resetClientTimestampTypeMapping",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDateInputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDateInputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDateOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDateOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDaysToExpiry() {
	_jsii_.InvokeVoid(
		s,
		"resetDaysToExpiry",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDefaultNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDefaultRole() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDefaultSecondaryRolesOption() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultSecondaryRolesOption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDefaultWarehouse() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultWarehouse",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDisabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDisabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetDisplayName() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetEmail() {
	_jsii_.InvokeVoid(
		s,
		"resetEmail",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetEnableUnloadPhysicalTypeOptimization() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableUnloadPhysicalTypeOptimization",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetEnableUnredactedQuerySyntaxError() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableUnredactedQuerySyntaxError",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetErrorOnNondeterministicMerge() {
	_jsii_.InvokeVoid(
		s,
		"resetErrorOnNondeterministicMerge",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetErrorOnNondeterministicUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetErrorOnNondeterministicUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetGeographyOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetGeographyOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetGeometryOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetGeometryOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetJdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		s,
		"resetJdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetJdbcTreatTimestampNtzAsUtc() {
	_jsii_.InvokeVoid(
		s,
		"resetJdbcTreatTimestampNtzAsUtc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetJdbcUseSessionTimezone() {
	_jsii_.InvokeVoid(
		s,
		"resetJdbcUseSessionTimezone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetJsonIndent() {
	_jsii_.InvokeVoid(
		s,
		"resetJsonIndent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetLockTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetLockTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetLoginName() {
	_jsii_.InvokeVoid(
		s,
		"resetLoginName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetLogLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetMinsToUnlock() {
	_jsii_.InvokeVoid(
		s,
		"resetMinsToUnlock",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetMultiStatementCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiStatementCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetNoorderSequenceAsDefault() {
	_jsii_.InvokeVoid(
		s,
		"resetNoorderSequenceAsDefault",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetOdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		s,
		"resetOdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetPreventUnloadToInternalStages() {
	_jsii_.InvokeVoid(
		s,
		"resetPreventUnloadToInternalStages",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetQueryTag() {
	_jsii_.InvokeVoid(
		s,
		"resetQueryTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetQuotedIdentifiersIgnoreCase() {
	_jsii_.InvokeVoid(
		s,
		"resetQuotedIdentifiersIgnoreCase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetRowsPerResultset() {
	_jsii_.InvokeVoid(
		s,
		"resetRowsPerResultset",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetRsaPublicKey() {
	_jsii_.InvokeVoid(
		s,
		"resetRsaPublicKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetRsaPublicKey2() {
	_jsii_.InvokeVoid(
		s,
		"resetRsaPublicKey2",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetS3StageVpceDnsName() {
	_jsii_.InvokeVoid(
		s,
		"resetS3StageVpceDnsName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetSearchPath() {
	_jsii_.InvokeVoid(
		s,
		"resetSearchPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetSimulatedDataSharingConsumer() {
	_jsii_.InvokeVoid(
		s,
		"resetSimulatedDataSharingConsumer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetStatementQueuedTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetStatementQueuedTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetStatementTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetStatementTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetStrictJsonOutput() {
	_jsii_.InvokeVoid(
		s,
		"resetStrictJsonOutput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimeInputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeInputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimeOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimestampDayIsAlways24H() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampDayIsAlways24H",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimestampInputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampInputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimestampLtzOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampLtzOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimestampNtzOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampNtzOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimestampOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampTypeMapping",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimestampTzOutputFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampTzOutputFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTimezone() {
	_jsii_.InvokeVoid(
		s,
		"resetTimezone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTransactionAbortOnError() {
	_jsii_.InvokeVoid(
		s,
		"resetTransactionAbortOnError",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTransactionDefaultIsolationLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetTransactionDefaultIsolationLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetTwoDigitCenturyStart() {
	_jsii_.InvokeVoid(
		s,
		"resetTwoDigitCenturyStart",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetUnsupportedDdlAction() {
	_jsii_.InvokeVoid(
		s,
		"resetUnsupportedDdlAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetUseCachedResult() {
	_jsii_.InvokeVoid(
		s,
		"resetUseCachedResult",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetWeekOfYearPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetWeekOfYearPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) ResetWeekStart() {
	_jsii_.InvokeVoid(
		s,
		"resetWeekStart",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUser) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUser) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

