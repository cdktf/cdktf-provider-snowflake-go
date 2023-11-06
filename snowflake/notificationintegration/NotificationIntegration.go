// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v10/notificationintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/notification_integration snowflake_notification_integration}.
type NotificationIntegration interface {
	cdktf.TerraformResource
	AwsSnsExternalId() *string
	AwsSnsIamUserArn() *string
	AwsSnsRoleArn() *string
	SetAwsSnsRoleArn(val *string)
	AwsSnsRoleArnInput() *string
	AwsSnsTopicArn() *string
	SetAwsSnsTopicArn(val *string)
	AwsSnsTopicArnInput() *string
	AwsSqsArn() *string
	SetAwsSqsArn(val *string)
	AwsSqsArnInput() *string
	AwsSqsExternalId() *string
	AwsSqsIamUserArn() *string
	AwsSqsRoleArn() *string
	SetAwsSqsRoleArn(val *string)
	AwsSqsRoleArnInput() *string
	AzureStorageQueuePrimaryUri() *string
	SetAzureStorageQueuePrimaryUri(val *string)
	AzureStorageQueuePrimaryUriInput() *string
	AzureTenantId() *string
	SetAzureTenantId(val *string)
	AzureTenantIdInput() *string
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
	CreatedOn() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpPubsubServiceAccount() *string
	GcpPubsubSubscriptionName() *string
	SetGcpPubsubSubscriptionName(val *string)
	GcpPubsubSubscriptionNameInput() *string
	GcpPubsubTopicName() *string
	SetGcpPubsubTopicName(val *string)
	GcpPubsubTopicNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationProvider() *string
	SetNotificationProvider(val *string)
	NotificationProviderInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAwsSnsRoleArn()
	ResetAwsSnsTopicArn()
	ResetAwsSqsArn()
	ResetAwsSqsRoleArn()
	ResetAzureStorageQueuePrimaryUri()
	ResetAzureTenantId()
	ResetComment()
	ResetDirection()
	ResetEnabled()
	ResetGcpPubsubSubscriptionName()
	ResetGcpPubsubTopicName()
	ResetId()
	ResetNotificationProvider()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NotificationIntegration
type jsiiProxy_NotificationIntegration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NotificationIntegration) AwsSnsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSnsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSnsIamUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSnsIamUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSnsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSnsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSnsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSnsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSnsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSnsTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSnsTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSqsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSqsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSqsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSqsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSqsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSqsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSqsIamUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSqsIamUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSqsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSqsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AwsSqsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSqsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AzureStorageQueuePrimaryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageQueuePrimaryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AzureStorageQueuePrimaryUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageQueuePrimaryUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) AzureTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) GcpPubsubServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpPubsubServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) GcpPubsubSubscriptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpPubsubSubscriptionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) GcpPubsubSubscriptionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpPubsubSubscriptionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) GcpPubsubTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpPubsubTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) GcpPubsubTopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpPubsubTopicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) NotificationProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) NotificationProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationIntegration) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/notification_integration snowflake_notification_integration} Resource.
func NewNotificationIntegration(scope constructs.Construct, id *string, config *NotificationIntegrationConfig) NotificationIntegration {
	_init_.Initialize()

	if err := validateNewNotificationIntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationIntegration{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/notification_integration snowflake_notification_integration} Resource.
func NewNotificationIntegration_Override(n NotificationIntegration, scope constructs.Construct, id *string, config *NotificationIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegration",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetAwsSnsRoleArn(val *string) {
	if err := j.validateSetAwsSnsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSnsRoleArn",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetAwsSnsTopicArn(val *string) {
	if err := j.validateSetAwsSnsTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSnsTopicArn",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetAwsSqsArn(val *string) {
	if err := j.validateSetAwsSqsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSqsArn",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetAwsSqsRoleArn(val *string) {
	if err := j.validateSetAwsSqsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSqsRoleArn",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetAzureStorageQueuePrimaryUri(val *string) {
	if err := j.validateSetAzureStorageQueuePrimaryUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageQueuePrimaryUri",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetAzureTenantId(val *string) {
	if err := j.validateSetAzureTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureTenantId",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetGcpPubsubSubscriptionName(val *string) {
	if err := j.validateSetGcpPubsubSubscriptionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpPubsubSubscriptionName",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetGcpPubsubTopicName(val *string) {
	if err := j.validateSetGcpPubsubTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpPubsubTopicName",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetNotificationProvider(val *string) {
	if err := j.validateSetNotificationProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationProvider",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NotificationIntegration)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a NotificationIntegration resource upon running "cdktf plan <stack-name>".
func NotificationIntegration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNotificationIntegration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegration",
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
func NotificationIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotificationIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NotificationIntegration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotificationIntegration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NotificationIntegration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotificationIntegration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NotificationIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.notificationIntegration.NotificationIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NotificationIntegration) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NotificationIntegration) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NotificationIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NotificationIntegration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NotificationIntegration) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetAwsSnsRoleArn() {
	_jsii_.InvokeVoid(
		n,
		"resetAwsSnsRoleArn",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetAwsSnsTopicArn() {
	_jsii_.InvokeVoid(
		n,
		"resetAwsSnsTopicArn",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetAwsSqsArn() {
	_jsii_.InvokeVoid(
		n,
		"resetAwsSqsArn",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetAwsSqsRoleArn() {
	_jsii_.InvokeVoid(
		n,
		"resetAwsSqsRoleArn",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetAzureStorageQueuePrimaryUri() {
	_jsii_.InvokeVoid(
		n,
		"resetAzureStorageQueuePrimaryUri",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetAzureTenantId() {
	_jsii_.InvokeVoid(
		n,
		"resetAzureTenantId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetComment() {
	_jsii_.InvokeVoid(
		n,
		"resetComment",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetDirection() {
	_jsii_.InvokeVoid(
		n,
		"resetDirection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetGcpPubsubSubscriptionName() {
	_jsii_.InvokeVoid(
		n,
		"resetGcpPubsubSubscriptionName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetGcpPubsubTopicName() {
	_jsii_.InvokeVoid(
		n,
		"resetGcpPubsubTopicName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetNotificationProvider() {
	_jsii_.InvokeVoid(
		n,
		"resetNotificationProvider",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) ResetType() {
	_jsii_.InvokeVoid(
		n,
		"resetType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

