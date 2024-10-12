// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secondarydatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/secondarydatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/secondary_database snowflake_secondary_database}.
type SecondaryDatabase interface {
	cdktf.TerraformResource
	AsReplicaOf() *string
	SetAsReplicaOf(val *string)
	AsReplicaOfInput() *string
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DataRetentionTimeInDays() *float64
	SetDataRetentionTimeInDays(val *float64)
	DataRetentionTimeInDaysInput() *float64
	DefaultDdlCollation() *string
	SetDefaultDdlCollation(val *string)
	DefaultDdlCollationInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableConsoleOutput() interface{}
	SetEnableConsoleOutput(val interface{})
	EnableConsoleOutputInput() interface{}
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
	FullyQualifiedName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsTransient() interface{}
	SetIsTransient(val interface{})
	IsTransientInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	MaxDataExtensionTimeInDays() *float64
	SetMaxDataExtensionTimeInDays(val *float64)
	MaxDataExtensionTimeInDaysInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QuotedIdentifiersIgnoreCase() interface{}
	SetQuotedIdentifiersIgnoreCase(val interface{})
	QuotedIdentifiersIgnoreCaseInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReplaceInvalidCharacters() interface{}
	SetReplaceInvalidCharacters(val interface{})
	ReplaceInvalidCharactersInput() interface{}
	StorageSerializationPolicy() *string
	SetStorageSerializationPolicy(val *string)
	StorageSerializationPolicyInput() *string
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
	TraceLevel() *string
	SetTraceLevel(val *string)
	TraceLevelInput() *string
	UserTaskManagedInitialWarehouseSize() *string
	SetUserTaskManagedInitialWarehouseSize(val *string)
	UserTaskManagedInitialWarehouseSizeInput() *string
	UserTaskMinimumTriggerIntervalInSeconds() *float64
	SetUserTaskMinimumTriggerIntervalInSeconds(val *float64)
	UserTaskMinimumTriggerIntervalInSecondsInput() *float64
	UserTaskTimeoutMs() *float64
	SetUserTaskTimeoutMs(val *float64)
	UserTaskTimeoutMsInput() *float64
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
	ResetCatalog()
	ResetComment()
	ResetDataRetentionTimeInDays()
	ResetDefaultDdlCollation()
	ResetEnableConsoleOutput()
	ResetExternalVolume()
	ResetId()
	ResetIsTransient()
	ResetLogLevel()
	ResetMaxDataExtensionTimeInDays()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQuotedIdentifiersIgnoreCase()
	ResetReplaceInvalidCharacters()
	ResetStorageSerializationPolicy()
	ResetSuspendTaskAfterNumFailures()
	ResetTaskAutoRetryAttempts()
	ResetTraceLevel()
	ResetUserTaskManagedInitialWarehouseSize()
	ResetUserTaskMinimumTriggerIntervalInSeconds()
	ResetUserTaskTimeoutMs()
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

// The jsii proxy struct for SecondaryDatabase
type jsiiProxy_SecondaryDatabase struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecondaryDatabase) AsReplicaOf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asReplicaOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) AsReplicaOfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asReplicaOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) DataRetentionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) DataRetentionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) DefaultDdlCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDdlCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) DefaultDdlCollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDdlCollationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) EnableConsoleOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) EnableConsoleOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) ExternalVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) ExternalVolumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) IsTransient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTransient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) IsTransientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTransientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) MaxDataExtensionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) MaxDataExtensionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) QuotedIdentifiersIgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) QuotedIdentifiersIgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) ReplaceInvalidCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) ReplaceInvalidCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) StorageSerializationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) StorageSerializationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) SuspendTaskAfterNumFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) SuspendTaskAfterNumFailuresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) TaskAutoRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) TaskAutoRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) UserTaskManagedInitialWarehouseSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) UserTaskManagedInitialWarehouseSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) UserTaskMinimumTriggerIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) UserTaskMinimumTriggerIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) UserTaskTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecondaryDatabase) UserTaskTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/secondary_database snowflake_secondary_database} Resource.
func NewSecondaryDatabase(scope constructs.Construct, id *string, config *SecondaryDatabaseConfig) SecondaryDatabase {
	_init_.Initialize()

	if err := validateNewSecondaryDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecondaryDatabase{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.secondaryDatabase.SecondaryDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/secondary_database snowflake_secondary_database} Resource.
func NewSecondaryDatabase_Override(s SecondaryDatabase, scope constructs.Construct, id *string, config *SecondaryDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.secondaryDatabase.SecondaryDatabase",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetAsReplicaOf(val *string) {
	if err := j.validateSetAsReplicaOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asReplicaOf",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetDataRetentionTimeInDays(val *float64) {
	if err := j.validateSetDataRetentionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRetentionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetDefaultDdlCollation(val *string) {
	if err := j.validateSetDefaultDdlCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDdlCollation",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetEnableConsoleOutput(val interface{}) {
	if err := j.validateSetEnableConsoleOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsoleOutput",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetExternalVolume(val *string) {
	if err := j.validateSetExternalVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalVolume",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetIsTransient(val interface{}) {
	if err := j.validateSetIsTransientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isTransient",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetMaxDataExtensionTimeInDays(val *float64) {
	if err := j.validateSetMaxDataExtensionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDataExtensionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetQuotedIdentifiersIgnoreCase(val interface{}) {
	if err := j.validateSetQuotedIdentifiersIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotedIdentifiersIgnoreCase",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetReplaceInvalidCharacters(val interface{}) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetStorageSerializationPolicy(val *string) {
	if err := j.validateSetStorageSerializationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSerializationPolicy",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetSuspendTaskAfterNumFailures(val *float64) {
	if err := j.validateSetSuspendTaskAfterNumFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTaskAfterNumFailures",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetTaskAutoRetryAttempts(val *float64) {
	if err := j.validateSetTaskAutoRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAutoRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetUserTaskManagedInitialWarehouseSize(val *string) {
	if err := j.validateSetUserTaskManagedInitialWarehouseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskManagedInitialWarehouseSize",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetUserTaskMinimumTriggerIntervalInSeconds(val *float64) {
	if err := j.validateSetUserTaskMinimumTriggerIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_SecondaryDatabase)SetUserTaskTimeoutMs(val *float64) {
	if err := j.validateSetUserTaskTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskTimeoutMs",
		val,
	)
}

// Generates CDKTF code for importing a SecondaryDatabase resource upon running "cdktf plan <stack-name>".
func SecondaryDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSecondaryDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.secondaryDatabase.SecondaryDatabase",
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
func SecondaryDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecondaryDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.secondaryDatabase.SecondaryDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecondaryDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecondaryDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.secondaryDatabase.SecondaryDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecondaryDatabase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecondaryDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.secondaryDatabase.SecondaryDatabase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecondaryDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.secondaryDatabase.SecondaryDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecondaryDatabase) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecondaryDatabase) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecondaryDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecondaryDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecondaryDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecondaryDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecondaryDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecondaryDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecondaryDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecondaryDatabase) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecondaryDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecondaryDatabase) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecondaryDatabase) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecondaryDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecondaryDatabase) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecondaryDatabase) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecondaryDatabase) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecondaryDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetCatalog() {
	_jsii_.InvokeVoid(
		s,
		"resetCatalog",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetDataRetentionTimeInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetDataRetentionTimeInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetDefaultDdlCollation() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultDdlCollation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetEnableConsoleOutput() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableConsoleOutput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetExternalVolume() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalVolume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetIsTransient() {
	_jsii_.InvokeVoid(
		s,
		"resetIsTransient",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetLogLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetMaxDataExtensionTimeInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxDataExtensionTimeInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetQuotedIdentifiersIgnoreCase() {
	_jsii_.InvokeVoid(
		s,
		"resetQuotedIdentifiersIgnoreCase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetStorageSerializationPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageSerializationPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetSuspendTaskAfterNumFailures() {
	_jsii_.InvokeVoid(
		s,
		"resetSuspendTaskAfterNumFailures",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetTaskAutoRetryAttempts() {
	_jsii_.InvokeVoid(
		s,
		"resetTaskAutoRetryAttempts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetUserTaskManagedInitialWarehouseSize() {
	_jsii_.InvokeVoid(
		s,
		"resetUserTaskManagedInitialWarehouseSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetUserTaskMinimumTriggerIntervalInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetUserTaskMinimumTriggerIntervalInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) ResetUserTaskTimeoutMs() {
	_jsii_.InvokeVoid(
		s,
		"resetUserTaskTimeoutMs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecondaryDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecondaryDatabase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecondaryDatabase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecondaryDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecondaryDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecondaryDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

