// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package legacyserviceuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/legacyserviceuser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/legacy_service_user snowflake_legacy_service_user}.
type LegacyServiceUser interface {
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
	MustChangePassword() *string
	SetMustChangePassword(val *string)
	MustChangePasswordInput() *string
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
	Parameters() LegacyServiceUserParametersList
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	ShowOutput() LegacyServiceUserShowOutputList
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
	ResetMustChangePassword()
	ResetNetworkPolicy()
	ResetNoorderSequenceAsDefault()
	ResetOdbcTreatDecimalAsInt()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
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

// The jsii proxy struct for LegacyServiceUser
type jsiiProxy_LegacyServiceUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LegacyServiceUser) AbortDetachedQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) AbortDetachedQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Autocommit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) AutocommitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) BinaryInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) BinaryInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) BinaryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) BinaryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientMemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientMemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientMetadataRequestUseConnectionCtx() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientMetadataRequestUseConnectionCtxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientPrefetchThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientPrefetchThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientResultChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientResultChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientResultColumnCaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientResultColumnCaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientSessionKeepAlive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientSessionKeepAliveHeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientSessionKeepAliveHeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientSessionKeepAliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientTimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ClientTimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DateInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DateInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DateOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DateOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DaysToExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DaysToExpiryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DefaultNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DefaultNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DefaultRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DefaultRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DefaultSecondaryRolesOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryRolesOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DefaultSecondaryRolesOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryRolesOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DefaultWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DefaultWarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Disabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DisabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) EnableUnloadPhysicalTypeOptimization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) EnableUnloadPhysicalTypeOptimizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) EnableUnredactedQuerySyntaxError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) EnableUnredactedQuerySyntaxErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ErrorOnNondeterministicMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ErrorOnNondeterministicMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ErrorOnNondeterministicUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ErrorOnNondeterministicUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) GeographyOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) GeographyOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) GeometryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) GeometryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) JdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) JdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) JdbcTreatTimestampNtzAsUtc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) JdbcTreatTimestampNtzAsUtcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) JdbcUseSessionTimezone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) JdbcUseSessionTimezoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) JsonIndent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) JsonIndentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) LockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) LockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) LoginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) LoginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) MinsToUnlock() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToUnlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) MinsToUnlockInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToUnlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) MultiStatementCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) MultiStatementCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) MustChangePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mustChangePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) MustChangePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mustChangePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) NetworkPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) NetworkPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) NoorderSequenceAsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) NoorderSequenceAsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) OdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) OdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Parameters() LegacyServiceUserParametersList {
	var returns LegacyServiceUserParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) PreventUnloadToInternalStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) PreventUnloadToInternalStagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) QueryTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) QueryTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) QuotedIdentifiersIgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) QuotedIdentifiersIgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) RowsPerResultset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) RowsPerResultsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) RsaPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) RsaPublicKey2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) RsaPublicKey2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) RsaPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) S3StageVpceDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) S3StageVpceDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) SearchPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) SearchPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) ShowOutput() LegacyServiceUserShowOutputList {
	var returns LegacyServiceUserShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) SimulatedDataSharingConsumer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) SimulatedDataSharingConsumerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) StatementQueuedTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) StatementQueuedTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) StatementTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) StatementTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) StrictJsonOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) StrictJsonOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimeInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimeInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimeOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimeOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampDayIsAlways24H() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampDayIsAlways24HInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24HInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampLtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampLtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampNtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampNtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampTzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimestampTzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TransactionAbortOnError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TransactionAbortOnErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TransactionDefaultIsolationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TransactionDefaultIsolationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TwoDigitCenturyStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) TwoDigitCenturyStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) UnsupportedDdlAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) UnsupportedDdlActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) UseCachedResult() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) UseCachedResultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) WeekOfYearPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) WeekOfYearPolicyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) WeekStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUser) WeekStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStartInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/legacy_service_user snowflake_legacy_service_user} Resource.
func NewLegacyServiceUser(scope constructs.Construct, id *string, config *LegacyServiceUserConfig) LegacyServiceUser {
	_init_.Initialize()

	if err := validateNewLegacyServiceUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LegacyServiceUser{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/legacy_service_user snowflake_legacy_service_user} Resource.
func NewLegacyServiceUser_Override(l LegacyServiceUser, scope constructs.Construct, id *string, config *LegacyServiceUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUser",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetAbortDetachedQuery(val interface{}) {
	if err := j.validateSetAbortDetachedQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortDetachedQuery",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetAutocommit(val interface{}) {
	if err := j.validateSetAutocommitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autocommit",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetBinaryInputFormat(val *string) {
	if err := j.validateSetBinaryInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryInputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetBinaryOutputFormat(val *string) {
	if err := j.validateSetBinaryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetClientMemoryLimit(val *float64) {
	if err := j.validateSetClientMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMemoryLimit",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetClientMetadataRequestUseConnectionCtx(val interface{}) {
	if err := j.validateSetClientMetadataRequestUseConnectionCtxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadataRequestUseConnectionCtx",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetClientPrefetchThreads(val *float64) {
	if err := j.validateSetClientPrefetchThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientPrefetchThreads",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetClientResultChunkSize(val *float64) {
	if err := j.validateSetClientResultChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultChunkSize",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetClientResultColumnCaseInsensitive(val interface{}) {
	if err := j.validateSetClientResultColumnCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultColumnCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetClientSessionKeepAlive(val interface{}) {
	if err := j.validateSetClientSessionKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAlive",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetClientSessionKeepAliveHeartbeatFrequency(val *float64) {
	if err := j.validateSetClientSessionKeepAliveHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetClientTimestampTypeMapping(val *string) {
	if err := j.validateSetClientTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTimestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDateInputFormat(val *string) {
	if err := j.validateSetDateInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateInputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDateOutputFormat(val *string) {
	if err := j.validateSetDateOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDaysToExpiry(val *float64) {
	if err := j.validateSetDaysToExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysToExpiry",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDefaultNamespace(val *string) {
	if err := j.validateSetDefaultNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNamespace",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDefaultRole(val *string) {
	if err := j.validateSetDefaultRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRole",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDefaultSecondaryRolesOption(val *string) {
	if err := j.validateSetDefaultSecondaryRolesOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSecondaryRolesOption",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDefaultWarehouse(val *string) {
	if err := j.validateSetDefaultWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultWarehouse",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDisabled(val *string) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetEnableUnloadPhysicalTypeOptimization(val interface{}) {
	if err := j.validateSetEnableUnloadPhysicalTypeOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnloadPhysicalTypeOptimization",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetEnableUnredactedQuerySyntaxError(val interface{}) {
	if err := j.validateSetEnableUnredactedQuerySyntaxErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnredactedQuerySyntaxError",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetErrorOnNondeterministicMerge(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicMerge",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetErrorOnNondeterministicUpdate(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicUpdate",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetGeographyOutputFormat(val *string) {
	if err := j.validateSetGeographyOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geographyOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetGeometryOutputFormat(val *string) {
	if err := j.validateSetGeometryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geometryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetJdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetJdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetJdbcTreatTimestampNtzAsUtc(val interface{}) {
	if err := j.validateSetJdbcTreatTimestampNtzAsUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetJdbcUseSessionTimezone(val interface{}) {
	if err := j.validateSetJdbcUseSessionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcUseSessionTimezone",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetJsonIndent(val *float64) {
	if err := j.validateSetJsonIndentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonIndent",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetLockTimeout(val *float64) {
	if err := j.validateSetLockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockTimeout",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetLoginName(val *string) {
	if err := j.validateSetLoginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginName",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetMinsToUnlock(val *float64) {
	if err := j.validateSetMinsToUnlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minsToUnlock",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetMultiStatementCount(val *float64) {
	if err := j.validateSetMultiStatementCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiStatementCount",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetMustChangePassword(val *string) {
	if err := j.validateSetMustChangePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mustChangePassword",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetNetworkPolicy(val *string) {
	if err := j.validateSetNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetNoorderSequenceAsDefault(val interface{}) {
	if err := j.validateSetNoorderSequenceAsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noorderSequenceAsDefault",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetOdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetOdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetPreventUnloadToInternalStages(val interface{}) {
	if err := j.validateSetPreventUnloadToInternalStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUnloadToInternalStages",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetQueryTag(val *string) {
	if err := j.validateSetQueryTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryTag",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetQuotedIdentifiersIgnoreCase(val interface{}) {
	if err := j.validateSetQuotedIdentifiersIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotedIdentifiersIgnoreCase",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetRowsPerResultset(val *float64) {
	if err := j.validateSetRowsPerResultsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsPerResultset",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetRsaPublicKey(val *string) {
	if err := j.validateSetRsaPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetRsaPublicKey2(val *string) {
	if err := j.validateSetRsaPublicKey2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey2",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetS3StageVpceDnsName(val *string) {
	if err := j.validateSetS3StageVpceDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3StageVpceDnsName",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetSearchPath(val *string) {
	if err := j.validateSetSearchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchPath",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetSimulatedDataSharingConsumer(val *string) {
	if err := j.validateSetSimulatedDataSharingConsumerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"simulatedDataSharingConsumer",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetStatementQueuedTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementQueuedTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementQueuedTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetStatementTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetStrictJsonOutput(val interface{}) {
	if err := j.validateSetStrictJsonOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictJsonOutput",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimeInputFormat(val *string) {
	if err := j.validateSetTimeInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeInputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimeOutputFormat(val *string) {
	if err := j.validateSetTimeOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimestampDayIsAlways24H(val interface{}) {
	if err := j.validateSetTimestampDayIsAlways24HParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDayIsAlways24H",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimestampInputFormat(val *string) {
	if err := j.validateSetTimestampInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampInputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimestampLtzOutputFormat(val *string) {
	if err := j.validateSetTimestampLtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampLtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimestampNtzOutputFormat(val *string) {
	if err := j.validateSetTimestampNtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampNtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimestampOutputFormat(val *string) {
	if err := j.validateSetTimestampOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimestampTypeMapping(val *string) {
	if err := j.validateSetTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimestampTzOutputFormat(val *string) {
	if err := j.validateSetTimestampTzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTransactionAbortOnError(val interface{}) {
	if err := j.validateSetTransactionAbortOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionAbortOnError",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTransactionDefaultIsolationLevel(val *string) {
	if err := j.validateSetTransactionDefaultIsolationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionDefaultIsolationLevel",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetTwoDigitCenturyStart(val *float64) {
	if err := j.validateSetTwoDigitCenturyStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoDigitCenturyStart",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetUnsupportedDdlAction(val *string) {
	if err := j.validateSetUnsupportedDdlActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unsupportedDdlAction",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetUseCachedResult(val interface{}) {
	if err := j.validateSetUseCachedResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCachedResult",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetWeekOfYearPolicy(val *float64) {
	if err := j.validateSetWeekOfYearPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekOfYearPolicy",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUser)SetWeekStart(val *float64) {
	if err := j.validateSetWeekStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekStart",
		val,
	)
}

// Generates CDKTF code for importing a LegacyServiceUser resource upon running "cdktf plan <stack-name>".
func LegacyServiceUser_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLegacyServiceUser_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUser",
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
func LegacyServiceUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLegacyServiceUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LegacyServiceUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLegacyServiceUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LegacyServiceUser_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLegacyServiceUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUser",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LegacyServiceUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LegacyServiceUser) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LegacyServiceUser) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LegacyServiceUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LegacyServiceUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LegacyServiceUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LegacyServiceUser) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LegacyServiceUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LegacyServiceUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LegacyServiceUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LegacyServiceUser) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LegacyServiceUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LegacyServiceUser) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUser) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LegacyServiceUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LegacyServiceUser) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LegacyServiceUser) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LegacyServiceUser) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LegacyServiceUser) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetAbortDetachedQuery() {
	_jsii_.InvokeVoid(
		l,
		"resetAbortDetachedQuery",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetAutocommit() {
	_jsii_.InvokeVoid(
		l,
		"resetAutocommit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetBinaryInputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetBinaryInputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetBinaryOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetBinaryOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetClientMemoryLimit() {
	_jsii_.InvokeVoid(
		l,
		"resetClientMemoryLimit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetClientMetadataRequestUseConnectionCtx() {
	_jsii_.InvokeVoid(
		l,
		"resetClientMetadataRequestUseConnectionCtx",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetClientPrefetchThreads() {
	_jsii_.InvokeVoid(
		l,
		"resetClientPrefetchThreads",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetClientResultChunkSize() {
	_jsii_.InvokeVoid(
		l,
		"resetClientResultChunkSize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetClientResultColumnCaseInsensitive() {
	_jsii_.InvokeVoid(
		l,
		"resetClientResultColumnCaseInsensitive",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetClientSessionKeepAlive() {
	_jsii_.InvokeVoid(
		l,
		"resetClientSessionKeepAlive",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetClientSessionKeepAliveHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		l,
		"resetClientSessionKeepAliveHeartbeatFrequency",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetClientTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		l,
		"resetClientTimestampTypeMapping",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetComment() {
	_jsii_.InvokeVoid(
		l,
		"resetComment",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDateInputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetDateInputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDateOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetDateOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDaysToExpiry() {
	_jsii_.InvokeVoid(
		l,
		"resetDaysToExpiry",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDefaultNamespace() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultNamespace",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDefaultRole() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultRole",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDefaultSecondaryRolesOption() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultSecondaryRolesOption",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDefaultWarehouse() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultWarehouse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDisabled() {
	_jsii_.InvokeVoid(
		l,
		"resetDisabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetDisplayName() {
	_jsii_.InvokeVoid(
		l,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetEmail() {
	_jsii_.InvokeVoid(
		l,
		"resetEmail",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetEnableUnloadPhysicalTypeOptimization() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableUnloadPhysicalTypeOptimization",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetEnableUnredactedQuerySyntaxError() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableUnredactedQuerySyntaxError",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetErrorOnNondeterministicMerge() {
	_jsii_.InvokeVoid(
		l,
		"resetErrorOnNondeterministicMerge",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetErrorOnNondeterministicUpdate() {
	_jsii_.InvokeVoid(
		l,
		"resetErrorOnNondeterministicUpdate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetGeographyOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetGeographyOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetGeometryOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetGeometryOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetJdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		l,
		"resetJdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetJdbcTreatTimestampNtzAsUtc() {
	_jsii_.InvokeVoid(
		l,
		"resetJdbcTreatTimestampNtzAsUtc",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetJdbcUseSessionTimezone() {
	_jsii_.InvokeVoid(
		l,
		"resetJdbcUseSessionTimezone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetJsonIndent() {
	_jsii_.InvokeVoid(
		l,
		"resetJsonIndent",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetLockTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetLockTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetLoginName() {
	_jsii_.InvokeVoid(
		l,
		"resetLoginName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetLogLevel() {
	_jsii_.InvokeVoid(
		l,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetMinsToUnlock() {
	_jsii_.InvokeVoid(
		l,
		"resetMinsToUnlock",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetMultiStatementCount() {
	_jsii_.InvokeVoid(
		l,
		"resetMultiStatementCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetMustChangePassword() {
	_jsii_.InvokeVoid(
		l,
		"resetMustChangePassword",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetNoorderSequenceAsDefault() {
	_jsii_.InvokeVoid(
		l,
		"resetNoorderSequenceAsDefault",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetOdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		l,
		"resetOdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetPassword() {
	_jsii_.InvokeVoid(
		l,
		"resetPassword",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetPreventUnloadToInternalStages() {
	_jsii_.InvokeVoid(
		l,
		"resetPreventUnloadToInternalStages",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetQueryTag() {
	_jsii_.InvokeVoid(
		l,
		"resetQueryTag",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetQuotedIdentifiersIgnoreCase() {
	_jsii_.InvokeVoid(
		l,
		"resetQuotedIdentifiersIgnoreCase",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetRowsPerResultset() {
	_jsii_.InvokeVoid(
		l,
		"resetRowsPerResultset",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetRsaPublicKey() {
	_jsii_.InvokeVoid(
		l,
		"resetRsaPublicKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetRsaPublicKey2() {
	_jsii_.InvokeVoid(
		l,
		"resetRsaPublicKey2",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetS3StageVpceDnsName() {
	_jsii_.InvokeVoid(
		l,
		"resetS3StageVpceDnsName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetSearchPath() {
	_jsii_.InvokeVoid(
		l,
		"resetSearchPath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetSimulatedDataSharingConsumer() {
	_jsii_.InvokeVoid(
		l,
		"resetSimulatedDataSharingConsumer",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetStatementQueuedTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetStatementQueuedTimeoutInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetStatementTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetStatementTimeoutInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetStrictJsonOutput() {
	_jsii_.InvokeVoid(
		l,
		"resetStrictJsonOutput",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimeInputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeInputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimeOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimestampDayIsAlways24H() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampDayIsAlways24H",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimestampInputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampInputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimestampLtzOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampLtzOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimestampNtzOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampNtzOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimestampOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampTypeMapping",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimestampTzOutputFormat() {
	_jsii_.InvokeVoid(
		l,
		"resetTimestampTzOutputFormat",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTimezone() {
	_jsii_.InvokeVoid(
		l,
		"resetTimezone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		l,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTransactionAbortOnError() {
	_jsii_.InvokeVoid(
		l,
		"resetTransactionAbortOnError",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTransactionDefaultIsolationLevel() {
	_jsii_.InvokeVoid(
		l,
		"resetTransactionDefaultIsolationLevel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetTwoDigitCenturyStart() {
	_jsii_.InvokeVoid(
		l,
		"resetTwoDigitCenturyStart",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetUnsupportedDdlAction() {
	_jsii_.InvokeVoid(
		l,
		"resetUnsupportedDdlAction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetUseCachedResult() {
	_jsii_.InvokeVoid(
		l,
		"resetUseCachedResult",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetWeekOfYearPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetWeekOfYearPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) ResetWeekStart() {
	_jsii_.InvokeVoid(
		l,
		"resetWeekStart",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUser) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUser) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

