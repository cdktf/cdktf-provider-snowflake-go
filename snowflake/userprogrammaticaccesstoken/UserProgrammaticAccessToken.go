// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userprogrammaticaccesstoken

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/userprogrammaticaccesstoken/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/user_programmatic_access_token snowflake_user_programmatic_access_token}.
type UserProgrammaticAccessToken interface {
	cdktf.TerraformResource
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
	DaysToExpiry() *float64
	SetDaysToExpiry(val *float64)
	DaysToExpiryInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() *string
	SetDisabled(val *string)
	DisabledInput() *string
	ExpireRotatedTokenAfterHours() *float64
	SetExpireRotatedTokenAfterHours(val *float64)
	ExpireRotatedTokenAfterHoursInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Keeper() *string
	SetKeeper(val *string)
	KeeperInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinsToBypassNetworkPolicyRequirement() *float64
	SetMinsToBypassNetworkPolicyRequirement(val *float64)
	MinsToBypassNetworkPolicyRequirementInput() *float64
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
	// Experimental.
	RawOverrides() interface{}
	RoleRestriction() *string
	SetRoleRestriction(val *string)
	RoleRestrictionInput() *string
	RotatedTokenName() *string
	ShowOutput() UserProgrammaticAccessTokenShowOutputList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() UserProgrammaticAccessTokenTimeoutsOutputReference
	TimeoutsInput() interface{}
	Token() *string
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	PutTimeouts(value *UserProgrammaticAccessTokenTimeouts)
	ResetComment()
	ResetDaysToExpiry()
	ResetDisabled()
	ResetExpireRotatedTokenAfterHours()
	ResetId()
	ResetKeeper()
	ResetMinsToBypassNetworkPolicyRequirement()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoleRestriction()
	ResetTimeouts()
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

// The jsii proxy struct for UserProgrammaticAccessToken
type jsiiProxy_UserProgrammaticAccessToken struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_UserProgrammaticAccessToken) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) DaysToExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) DaysToExpiryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Disabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) DisabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) ExpireRotatedTokenAfterHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireRotatedTokenAfterHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) ExpireRotatedTokenAfterHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireRotatedTokenAfterHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Keeper() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keeper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) KeeperInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keeperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) MinsToBypassNetworkPolicyRequirement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToBypassNetworkPolicyRequirement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) MinsToBypassNetworkPolicyRequirementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToBypassNetworkPolicyRequirementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) RoleRestriction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) RoleRestrictionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) RotatedTokenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotatedTokenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) ShowOutput() UserProgrammaticAccessTokenShowOutputList {
	var returns UserProgrammaticAccessTokenShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Timeouts() UserProgrammaticAccessTokenTimeoutsOutputReference {
	var returns UserProgrammaticAccessTokenTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserProgrammaticAccessToken) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/user_programmatic_access_token snowflake_user_programmatic_access_token} Resource.
func NewUserProgrammaticAccessToken(scope constructs.Construct, id *string, config *UserProgrammaticAccessTokenConfig) UserProgrammaticAccessToken {
	_init_.Initialize()

	if err := validateNewUserProgrammaticAccessTokenParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserProgrammaticAccessToken{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.userProgrammaticAccessToken.UserProgrammaticAccessToken",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/user_programmatic_access_token snowflake_user_programmatic_access_token} Resource.
func NewUserProgrammaticAccessToken_Override(u UserProgrammaticAccessToken, scope constructs.Construct, id *string, config *UserProgrammaticAccessTokenConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.userProgrammaticAccessToken.UserProgrammaticAccessToken",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetDaysToExpiry(val *float64) {
	if err := j.validateSetDaysToExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysToExpiry",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetDisabled(val *string) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetExpireRotatedTokenAfterHours(val *float64) {
	if err := j.validateSetExpireRotatedTokenAfterHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireRotatedTokenAfterHours",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetKeeper(val *string) {
	if err := j.validateSetKeeperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keeper",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetMinsToBypassNetworkPolicyRequirement(val *float64) {
	if err := j.validateSetMinsToBypassNetworkPolicyRequirementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minsToBypassNetworkPolicyRequirement",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetRoleRestriction(val *string) {
	if err := j.validateSetRoleRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleRestriction",
		val,
	)
}

func (j *jsiiProxy_UserProgrammaticAccessToken)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

// Generates CDKTF code for importing a UserProgrammaticAccessToken resource upon running "cdktf plan <stack-name>".
func UserProgrammaticAccessToken_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateUserProgrammaticAccessToken_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.userProgrammaticAccessToken.UserProgrammaticAccessToken",
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
func UserProgrammaticAccessToken_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserProgrammaticAccessToken_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.userProgrammaticAccessToken.UserProgrammaticAccessToken",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func UserProgrammaticAccessToken_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserProgrammaticAccessToken_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.userProgrammaticAccessToken.UserProgrammaticAccessToken",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func UserProgrammaticAccessToken_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserProgrammaticAccessToken_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.userProgrammaticAccessToken.UserProgrammaticAccessToken",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func UserProgrammaticAccessToken_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.userProgrammaticAccessToken.UserProgrammaticAccessToken",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UserProgrammaticAccessToken) AddMoveTarget(moveTarget *string) {
	if err := u.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) GetStringAttribute(terraformAttribute *string) *string {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := u.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserProgrammaticAccessToken) MoveFromId(id *string) {
	if err := u.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveFromId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) MoveTo(moveTarget *string, index interface{}) {
	if err := u.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) MoveToId(id *string) {
	if err := u.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveToId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) PutTimeouts(value *UserProgrammaticAccessTokenTimeouts) {
	if err := u.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetComment() {
	_jsii_.InvokeVoid(
		u,
		"resetComment",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetDaysToExpiry() {
	_jsii_.InvokeVoid(
		u,
		"resetDaysToExpiry",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetDisabled() {
	_jsii_.InvokeVoid(
		u,
		"resetDisabled",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetExpireRotatedTokenAfterHours() {
	_jsii_.InvokeVoid(
		u,
		"resetExpireRotatedTokenAfterHours",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetKeeper() {
	_jsii_.InvokeVoid(
		u,
		"resetKeeper",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetMinsToBypassNetworkPolicyRequirement() {
	_jsii_.InvokeVoid(
		u,
		"resetMinsToBypassNetworkPolicyRequirement",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetRoleRestriction() {
	_jsii_.InvokeVoid(
		u,
		"resetRoleRestriction",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ResetTimeouts() {
	_jsii_.InvokeVoid(
		u,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserProgrammaticAccessToken) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserProgrammaticAccessToken) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserProgrammaticAccessToken) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

