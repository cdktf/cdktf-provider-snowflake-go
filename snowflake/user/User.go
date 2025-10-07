// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/user snowflake_user}.
type User interface {
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
	DisableMfa() *string
	SetDisableMfa(val *string)
	DisableMfaInput() *string
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
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
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
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
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
	MiddleName() *string
	SetMiddleName(val *string)
	MiddleNameInput() *string
	MinsToBypassMfa() *float64
	SetMinsToBypassMfa(val *float64)
	MinsToBypassMfaInput() *float64
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
	Parameters() UserParametersList
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
	ShowOutput() UserShowOutputList
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
	Timeouts() UserTimeoutsOutputReference
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
	PutTimeouts(value *UserTimeouts)
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
	ResetDisableMfa()
	ResetDisplayName()
	ResetEmail()
	ResetEnableUnloadPhysicalTypeOptimization()
	ResetEnableUnredactedQuerySyntaxError()
	ResetErrorOnNondeterministicMerge()
	ResetErrorOnNondeterministicUpdate()
	ResetFirstName()
	ResetGeographyOutputFormat()
	ResetGeometryOutputFormat()
	ResetId()
	ResetJdbcTreatDecimalAsInt()
	ResetJdbcTreatTimestampNtzAsUtc()
	ResetJdbcUseSessionTimezone()
	ResetJsonIndent()
	ResetLastName()
	ResetLockTimeout()
	ResetLoginName()
	ResetLogLevel()
	ResetMiddleName()
	ResetMinsToBypassMfa()
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

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_User) AbortDetachedQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AbortDetachedQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortDetachedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Autocommit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AutocommitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autocommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) BinaryInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) BinaryInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) BinaryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) BinaryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientMemoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientMemoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientMemoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientMetadataRequestUseConnectionCtx() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientMetadataRequestUseConnectionCtxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientPrefetchThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientPrefetchThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPrefetchThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientResultChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientResultChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientResultChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientResultColumnCaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientResultColumnCaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientSessionKeepAlive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientSessionKeepAliveHeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientSessionKeepAliveHeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientSessionKeepAliveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSessionKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientTimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ClientTimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTimestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DateInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DateInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DateOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DateOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DaysToExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DaysToExpiryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultSecondaryRolesOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryRolesOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultSecondaryRolesOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryRolesOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultWarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Disabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisableMfa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disableMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisableMfaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disableMfaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EnableUnloadPhysicalTypeOptimization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EnableUnloadPhysicalTypeOptimizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EnableUnredactedQuerySyntaxError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EnableUnredactedQuerySyntaxErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ErrorOnNondeterministicMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ErrorOnNondeterministicMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ErrorOnNondeterministicUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ErrorOnNondeterministicUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GeographyOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GeographyOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geographyOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GeometryOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GeometryOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geometryOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JdbcTreatTimestampNtzAsUtc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JdbcTreatTimestampNtzAsUtcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JdbcUseSessionTimezone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JdbcUseSessionTimezoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JsonIndent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JsonIndentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsonIndentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LoginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LoginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MiddleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MinsToBypassMfa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToBypassMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MinsToBypassMfaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToBypassMfaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MinsToUnlock() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToUnlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MinsToUnlockInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToUnlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MultiStatementCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MultiStatementCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiStatementCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MustChangePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mustChangePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MustChangePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mustChangePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NetworkPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NetworkPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NoorderSequenceAsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NoorderSequenceAsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noorderSequenceAsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OdbcTreatDecimalAsInt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OdbcTreatDecimalAsIntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsIntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Parameters() UserParametersList {
	var returns UserParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PreventUnloadToInternalStages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PreventUnloadToInternalStagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventUnloadToInternalStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) QueryTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) QueryTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) QuotedIdentifiersIgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) QuotedIdentifiersIgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RowsPerResultset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RowsPerResultsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsPerResultsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RsaPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RsaPublicKey2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RsaPublicKey2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RsaPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) S3StageVpceDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) S3StageVpceDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3StageVpceDnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SearchPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SearchPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ShowOutput() UserShowOutputList {
	var returns UserShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SimulatedDataSharingConsumer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SimulatedDataSharingConsumerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StatementQueuedTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StatementQueuedTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StatementTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StatementTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StrictJsonOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StrictJsonOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictJsonOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Timeouts() UserTimeoutsOutputReference {
	var returns UserTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampDayIsAlways24H() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampDayIsAlways24HInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestampDayIsAlways24HInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampInputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampInputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampLtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampLtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampLtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampNtzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampNtzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampNtzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampTypeMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampTypeMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTypeMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampTzOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimestampTzOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampTzOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TransactionAbortOnError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TransactionAbortOnErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionAbortOnErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TransactionDefaultIsolationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TransactionDefaultIsolationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TwoDigitCenturyStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TwoDigitCenturyStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"twoDigitCenturyStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UnsupportedDdlAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UnsupportedDdlActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unsupportedDdlActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UseCachedResult() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UseCachedResultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCachedResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) WeekOfYearPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) WeekOfYearPolicyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekOfYearPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) WeekStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) WeekStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekStartInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/user snowflake_user} Resource.
func NewUser(scope constructs.Construct, id *string, config *UserConfig) User {
	_init_.Initialize()

	if err := validateNewUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_User{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.user.User",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/user snowflake_user} Resource.
func NewUser_Override(u User, scope constructs.Construct, id *string, config *UserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.user.User",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_User)SetAbortDetachedQuery(val interface{}) {
	if err := j.validateSetAbortDetachedQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortDetachedQuery",
		val,
	)
}

func (j *jsiiProxy_User)SetAutocommit(val interface{}) {
	if err := j.validateSetAutocommitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autocommit",
		val,
	)
}

func (j *jsiiProxy_User)SetBinaryInputFormat(val *string) {
	if err := j.validateSetBinaryInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryInputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetBinaryOutputFormat(val *string) {
	if err := j.validateSetBinaryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetClientMemoryLimit(val *float64) {
	if err := j.validateSetClientMemoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMemoryLimit",
		val,
	)
}

func (j *jsiiProxy_User)SetClientMetadataRequestUseConnectionCtx(val interface{}) {
	if err := j.validateSetClientMetadataRequestUseConnectionCtxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadataRequestUseConnectionCtx",
		val,
	)
}

func (j *jsiiProxy_User)SetClientPrefetchThreads(val *float64) {
	if err := j.validateSetClientPrefetchThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientPrefetchThreads",
		val,
	)
}

func (j *jsiiProxy_User)SetClientResultChunkSize(val *float64) {
	if err := j.validateSetClientResultChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultChunkSize",
		val,
	)
}

func (j *jsiiProxy_User)SetClientResultColumnCaseInsensitive(val interface{}) {
	if err := j.validateSetClientResultColumnCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientResultColumnCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_User)SetClientSessionKeepAlive(val interface{}) {
	if err := j.validateSetClientSessionKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAlive",
		val,
	)
}

func (j *jsiiProxy_User)SetClientSessionKeepAliveHeartbeatFrequency(val *float64) {
	if err := j.validateSetClientSessionKeepAliveHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_User)SetClientTimestampTypeMapping(val *string) {
	if err := j.validateSetClientTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTimestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_User)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_User)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_User)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_User)SetDateInputFormat(val *string) {
	if err := j.validateSetDateInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateInputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetDateOutputFormat(val *string) {
	if err := j.validateSetDateOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetDaysToExpiry(val *float64) {
	if err := j.validateSetDaysToExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysToExpiry",
		val,
	)
}

func (j *jsiiProxy_User)SetDefaultNamespace(val *string) {
	if err := j.validateSetDefaultNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNamespace",
		val,
	)
}

func (j *jsiiProxy_User)SetDefaultRole(val *string) {
	if err := j.validateSetDefaultRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRole",
		val,
	)
}

func (j *jsiiProxy_User)SetDefaultSecondaryRolesOption(val *string) {
	if err := j.validateSetDefaultSecondaryRolesOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSecondaryRolesOption",
		val,
	)
}

func (j *jsiiProxy_User)SetDefaultWarehouse(val *string) {
	if err := j.validateSetDefaultWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultWarehouse",
		val,
	)
}

func (j *jsiiProxy_User)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_User)SetDisabled(val *string) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_User)SetDisableMfa(val *string) {
	if err := j.validateSetDisableMfaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableMfa",
		val,
	)
}

func (j *jsiiProxy_User)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_User)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_User)SetEnableUnloadPhysicalTypeOptimization(val interface{}) {
	if err := j.validateSetEnableUnloadPhysicalTypeOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnloadPhysicalTypeOptimization",
		val,
	)
}

func (j *jsiiProxy_User)SetEnableUnredactedQuerySyntaxError(val interface{}) {
	if err := j.validateSetEnableUnredactedQuerySyntaxErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUnredactedQuerySyntaxError",
		val,
	)
}

func (j *jsiiProxy_User)SetErrorOnNondeterministicMerge(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicMerge",
		val,
	)
}

func (j *jsiiProxy_User)SetErrorOnNondeterministicUpdate(val interface{}) {
	if err := j.validateSetErrorOnNondeterministicUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnNondeterministicUpdate",
		val,
	)
}

func (j *jsiiProxy_User)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_User)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_User)SetGeographyOutputFormat(val *string) {
	if err := j.validateSetGeographyOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geographyOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetGeometryOutputFormat(val *string) {
	if err := j.validateSetGeometryOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geometryOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_User)SetJdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetJdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_User)SetJdbcTreatTimestampNtzAsUtc(val interface{}) {
	if err := j.validateSetJdbcTreatTimestampNtzAsUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		val,
	)
}

func (j *jsiiProxy_User)SetJdbcUseSessionTimezone(val interface{}) {
	if err := j.validateSetJdbcUseSessionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcUseSessionTimezone",
		val,
	)
}

func (j *jsiiProxy_User)SetJsonIndent(val *float64) {
	if err := j.validateSetJsonIndentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonIndent",
		val,
	)
}

func (j *jsiiProxy_User)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_User)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_User)SetLockTimeout(val *float64) {
	if err := j.validateSetLockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockTimeout",
		val,
	)
}

func (j *jsiiProxy_User)SetLoginName(val *string) {
	if err := j.validateSetLoginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginName",
		val,
	)
}

func (j *jsiiProxy_User)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_User)SetMiddleName(val *string) {
	if err := j.validateSetMiddleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"middleName",
		val,
	)
}

func (j *jsiiProxy_User)SetMinsToBypassMfa(val *float64) {
	if err := j.validateSetMinsToBypassMfaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minsToBypassMfa",
		val,
	)
}

func (j *jsiiProxy_User)SetMinsToUnlock(val *float64) {
	if err := j.validateSetMinsToUnlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minsToUnlock",
		val,
	)
}

func (j *jsiiProxy_User)SetMultiStatementCount(val *float64) {
	if err := j.validateSetMultiStatementCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiStatementCount",
		val,
	)
}

func (j *jsiiProxy_User)SetMustChangePassword(val *string) {
	if err := j.validateSetMustChangePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mustChangePassword",
		val,
	)
}

func (j *jsiiProxy_User)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_User)SetNetworkPolicy(val *string) {
	if err := j.validateSetNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_User)SetNoorderSequenceAsDefault(val interface{}) {
	if err := j.validateSetNoorderSequenceAsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noorderSequenceAsDefault",
		val,
	)
}

func (j *jsiiProxy_User)SetOdbcTreatDecimalAsInt(val interface{}) {
	if err := j.validateSetOdbcTreatDecimalAsIntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbcTreatDecimalAsInt",
		val,
	)
}

func (j *jsiiProxy_User)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_User)SetPreventUnloadToInternalStages(val interface{}) {
	if err := j.validateSetPreventUnloadToInternalStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventUnloadToInternalStages",
		val,
	)
}

func (j *jsiiProxy_User)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_User)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_User)SetQueryTag(val *string) {
	if err := j.validateSetQueryTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryTag",
		val,
	)
}

func (j *jsiiProxy_User)SetQuotedIdentifiersIgnoreCase(val interface{}) {
	if err := j.validateSetQuotedIdentifiersIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotedIdentifiersIgnoreCase",
		val,
	)
}

func (j *jsiiProxy_User)SetRowsPerResultset(val *float64) {
	if err := j.validateSetRowsPerResultsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsPerResultset",
		val,
	)
}

func (j *jsiiProxy_User)SetRsaPublicKey(val *string) {
	if err := j.validateSetRsaPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey",
		val,
	)
}

func (j *jsiiProxy_User)SetRsaPublicKey2(val *string) {
	if err := j.validateSetRsaPublicKey2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey2",
		val,
	)
}

func (j *jsiiProxy_User)SetS3StageVpceDnsName(val *string) {
	if err := j.validateSetS3StageVpceDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3StageVpceDnsName",
		val,
	)
}

func (j *jsiiProxy_User)SetSearchPath(val *string) {
	if err := j.validateSetSearchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchPath",
		val,
	)
}

func (j *jsiiProxy_User)SetSimulatedDataSharingConsumer(val *string) {
	if err := j.validateSetSimulatedDataSharingConsumerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"simulatedDataSharingConsumer",
		val,
	)
}

func (j *jsiiProxy_User)SetStatementQueuedTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementQueuedTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementQueuedTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_User)SetStatementTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_User)SetStrictJsonOutput(val interface{}) {
	if err := j.validateSetStrictJsonOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictJsonOutput",
		val,
	)
}

func (j *jsiiProxy_User)SetTimeInputFormat(val *string) {
	if err := j.validateSetTimeInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeInputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetTimeOutputFormat(val *string) {
	if err := j.validateSetTimeOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetTimestampDayIsAlways24H(val interface{}) {
	if err := j.validateSetTimestampDayIsAlways24HParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDayIsAlways24H",
		val,
	)
}

func (j *jsiiProxy_User)SetTimestampInputFormat(val *string) {
	if err := j.validateSetTimestampInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampInputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetTimestampLtzOutputFormat(val *string) {
	if err := j.validateSetTimestampLtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampLtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetTimestampNtzOutputFormat(val *string) {
	if err := j.validateSetTimestampNtzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampNtzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetTimestampOutputFormat(val *string) {
	if err := j.validateSetTimestampOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetTimestampTypeMapping(val *string) {
	if err := j.validateSetTimestampTypeMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTypeMapping",
		val,
	)
}

func (j *jsiiProxy_User)SetTimestampTzOutputFormat(val *string) {
	if err := j.validateSetTimestampTzOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampTzOutputFormat",
		val,
	)
}

func (j *jsiiProxy_User)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_User)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

func (j *jsiiProxy_User)SetTransactionAbortOnError(val interface{}) {
	if err := j.validateSetTransactionAbortOnErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionAbortOnError",
		val,
	)
}

func (j *jsiiProxy_User)SetTransactionDefaultIsolationLevel(val *string) {
	if err := j.validateSetTransactionDefaultIsolationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionDefaultIsolationLevel",
		val,
	)
}

func (j *jsiiProxy_User)SetTwoDigitCenturyStart(val *float64) {
	if err := j.validateSetTwoDigitCenturyStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoDigitCenturyStart",
		val,
	)
}

func (j *jsiiProxy_User)SetUnsupportedDdlAction(val *string) {
	if err := j.validateSetUnsupportedDdlActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unsupportedDdlAction",
		val,
	)
}

func (j *jsiiProxy_User)SetUseCachedResult(val interface{}) {
	if err := j.validateSetUseCachedResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCachedResult",
		val,
	)
}

func (j *jsiiProxy_User)SetWeekOfYearPolicy(val *float64) {
	if err := j.validateSetWeekOfYearPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekOfYearPolicy",
		val,
	)
}

func (j *jsiiProxy_User)SetWeekStart(val *float64) {
	if err := j.validateSetWeekStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekStart",
		val,
	)
}

// Generates CDKTF code for importing a User resource upon running "cdktf plan <stack-name>".
func User_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateUser_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.user.User",
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
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.user.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func User_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.user.User",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func User_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.user.User",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func User_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.user.User",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_User) AddMoveTarget(moveTarget *string) {
	if err := u.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (u *jsiiProxy_User) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_User) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (u *jsiiProxy_User) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (u *jsiiProxy_User) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (u *jsiiProxy_User) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (u *jsiiProxy_User) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (u *jsiiProxy_User) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (u *jsiiProxy_User) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (u *jsiiProxy_User) GetStringAttribute(terraformAttribute *string) *string {
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

func (u *jsiiProxy_User) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (u *jsiiProxy_User) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := u.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (u *jsiiProxy_User) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) MoveFromId(id *string) {
	if err := u.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveFromId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_User) MoveTo(moveTarget *string, index interface{}) {
	if err := u.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (u *jsiiProxy_User) MoveToId(id *string) {
	if err := u.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveToId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_User) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_User) PutTimeouts(value *UserTimeouts) {
	if err := u.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) ResetAbortDetachedQuery() {
	_jsii_.InvokeVoid(
		u,
		"resetAbortDetachedQuery",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetAutocommit() {
	_jsii_.InvokeVoid(
		u,
		"resetAutocommit",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetBinaryInputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetBinaryInputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetBinaryOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetBinaryOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetClientMemoryLimit() {
	_jsii_.InvokeVoid(
		u,
		"resetClientMemoryLimit",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetClientMetadataRequestUseConnectionCtx() {
	_jsii_.InvokeVoid(
		u,
		"resetClientMetadataRequestUseConnectionCtx",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetClientPrefetchThreads() {
	_jsii_.InvokeVoid(
		u,
		"resetClientPrefetchThreads",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetClientResultChunkSize() {
	_jsii_.InvokeVoid(
		u,
		"resetClientResultChunkSize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetClientResultColumnCaseInsensitive() {
	_jsii_.InvokeVoid(
		u,
		"resetClientResultColumnCaseInsensitive",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetClientSessionKeepAlive() {
	_jsii_.InvokeVoid(
		u,
		"resetClientSessionKeepAlive",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetClientSessionKeepAliveHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		u,
		"resetClientSessionKeepAliveHeartbeatFrequency",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetClientTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		u,
		"resetClientTimestampTypeMapping",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetComment() {
	_jsii_.InvokeVoid(
		u,
		"resetComment",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDateInputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetDateInputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDateOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetDateOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDaysToExpiry() {
	_jsii_.InvokeVoid(
		u,
		"resetDaysToExpiry",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDefaultNamespace() {
	_jsii_.InvokeVoid(
		u,
		"resetDefaultNamespace",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDefaultRole() {
	_jsii_.InvokeVoid(
		u,
		"resetDefaultRole",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDefaultSecondaryRolesOption() {
	_jsii_.InvokeVoid(
		u,
		"resetDefaultSecondaryRolesOption",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDefaultWarehouse() {
	_jsii_.InvokeVoid(
		u,
		"resetDefaultWarehouse",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDisabled() {
	_jsii_.InvokeVoid(
		u,
		"resetDisabled",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDisableMfa() {
	_jsii_.InvokeVoid(
		u,
		"resetDisableMfa",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDisplayName() {
	_jsii_.InvokeVoid(
		u,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmail() {
	_jsii_.InvokeVoid(
		u,
		"resetEmail",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEnableUnloadPhysicalTypeOptimization() {
	_jsii_.InvokeVoid(
		u,
		"resetEnableUnloadPhysicalTypeOptimization",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEnableUnredactedQuerySyntaxError() {
	_jsii_.InvokeVoid(
		u,
		"resetEnableUnredactedQuerySyntaxError",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetErrorOnNondeterministicMerge() {
	_jsii_.InvokeVoid(
		u,
		"resetErrorOnNondeterministicMerge",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetErrorOnNondeterministicUpdate() {
	_jsii_.InvokeVoid(
		u,
		"resetErrorOnNondeterministicUpdate",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetFirstName() {
	_jsii_.InvokeVoid(
		u,
		"resetFirstName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetGeographyOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetGeographyOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetGeometryOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetGeometryOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetJdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		u,
		"resetJdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetJdbcTreatTimestampNtzAsUtc() {
	_jsii_.InvokeVoid(
		u,
		"resetJdbcTreatTimestampNtzAsUtc",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetJdbcUseSessionTimezone() {
	_jsii_.InvokeVoid(
		u,
		"resetJdbcUseSessionTimezone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetJsonIndent() {
	_jsii_.InvokeVoid(
		u,
		"resetJsonIndent",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLastName() {
	_jsii_.InvokeVoid(
		u,
		"resetLastName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLockTimeout() {
	_jsii_.InvokeVoid(
		u,
		"resetLockTimeout",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLoginName() {
	_jsii_.InvokeVoid(
		u,
		"resetLoginName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLogLevel() {
	_jsii_.InvokeVoid(
		u,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMiddleName() {
	_jsii_.InvokeVoid(
		u,
		"resetMiddleName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMinsToBypassMfa() {
	_jsii_.InvokeVoid(
		u,
		"resetMinsToBypassMfa",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMinsToUnlock() {
	_jsii_.InvokeVoid(
		u,
		"resetMinsToUnlock",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMultiStatementCount() {
	_jsii_.InvokeVoid(
		u,
		"resetMultiStatementCount",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMustChangePassword() {
	_jsii_.InvokeVoid(
		u,
		"resetMustChangePassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		u,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetNoorderSequenceAsDefault() {
	_jsii_.InvokeVoid(
		u,
		"resetNoorderSequenceAsDefault",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOdbcTreatDecimalAsInt() {
	_jsii_.InvokeVoid(
		u,
		"resetOdbcTreatDecimalAsInt",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPreventUnloadToInternalStages() {
	_jsii_.InvokeVoid(
		u,
		"resetPreventUnloadToInternalStages",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetQueryTag() {
	_jsii_.InvokeVoid(
		u,
		"resetQueryTag",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetQuotedIdentifiersIgnoreCase() {
	_jsii_.InvokeVoid(
		u,
		"resetQuotedIdentifiersIgnoreCase",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRowsPerResultset() {
	_jsii_.InvokeVoid(
		u,
		"resetRowsPerResultset",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRsaPublicKey() {
	_jsii_.InvokeVoid(
		u,
		"resetRsaPublicKey",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRsaPublicKey2() {
	_jsii_.InvokeVoid(
		u,
		"resetRsaPublicKey2",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetS3StageVpceDnsName() {
	_jsii_.InvokeVoid(
		u,
		"resetS3StageVpceDnsName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSearchPath() {
	_jsii_.InvokeVoid(
		u,
		"resetSearchPath",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSimulatedDataSharingConsumer() {
	_jsii_.InvokeVoid(
		u,
		"resetSimulatedDataSharingConsumer",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetStatementQueuedTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		u,
		"resetStatementQueuedTimeoutInSeconds",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetStatementTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		u,
		"resetStatementTimeoutInSeconds",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetStrictJsonOutput() {
	_jsii_.InvokeVoid(
		u,
		"resetStrictJsonOutput",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimeInputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetTimeInputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimeOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetTimeOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimeouts() {
	_jsii_.InvokeVoid(
		u,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimestampDayIsAlways24H() {
	_jsii_.InvokeVoid(
		u,
		"resetTimestampDayIsAlways24H",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimestampInputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetTimestampInputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimestampLtzOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetTimestampLtzOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimestampNtzOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetTimestampNtzOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimestampOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetTimestampOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimestampTypeMapping() {
	_jsii_.InvokeVoid(
		u,
		"resetTimestampTypeMapping",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimestampTzOutputFormat() {
	_jsii_.InvokeVoid(
		u,
		"resetTimestampTzOutputFormat",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimezone() {
	_jsii_.InvokeVoid(
		u,
		"resetTimezone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		u,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTransactionAbortOnError() {
	_jsii_.InvokeVoid(
		u,
		"resetTransactionAbortOnError",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTransactionDefaultIsolationLevel() {
	_jsii_.InvokeVoid(
		u,
		"resetTransactionDefaultIsolationLevel",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTwoDigitCenturyStart() {
	_jsii_.InvokeVoid(
		u,
		"resetTwoDigitCenturyStart",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetUnsupportedDdlAction() {
	_jsii_.InvokeVoid(
		u,
		"resetUnsupportedDdlAction",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetUseCachedResult() {
	_jsii_.InvokeVoid(
		u,
		"resetUseCachedResult",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetWeekOfYearPolicy() {
	_jsii_.InvokeVoid(
		u,
		"resetWeekOfYearPolicy",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetWeekStart() {
	_jsii_.InvokeVoid(
		u,
		"resetWeekStart",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

