// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/grantprivilegestoaccountrole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/grant_privileges_to_account_role snowflake_grant_privileges_to_account_role}.
type GrantPrivilegesToAccountRole interface {
	cdktf.TerraformResource
	AccountRoleName() *string
	SetAccountRoleName(val *string)
	AccountRoleNameInput() *string
	AllPrivileges() interface{}
	SetAllPrivileges(val interface{})
	AllPrivilegesInput() interface{}
	AlwaysApply() interface{}
	SetAlwaysApply(val interface{})
	AlwaysApplyInput() interface{}
	AlwaysApplyTrigger() *string
	SetAlwaysApplyTrigger(val *string)
	AlwaysApplyTriggerInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OnAccount() interface{}
	SetOnAccount(val interface{})
	OnAccountInput() interface{}
	OnAccountObject() GrantPrivilegesToAccountRoleOnAccountObjectOutputReference
	OnAccountObjectInput() *GrantPrivilegesToAccountRoleOnAccountObject
	OnSchema() GrantPrivilegesToAccountRoleOnSchemaOutputReference
	OnSchemaInput() *GrantPrivilegesToAccountRoleOnSchema
	OnSchemaObject() GrantPrivilegesToAccountRoleOnSchemaObjectOutputReference
	OnSchemaObjectInput() *GrantPrivilegesToAccountRoleOnSchemaObject
	Privileges() *[]*string
	SetPrivileges(val *[]*string)
	PrivilegesInput() *[]*string
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
	Timeouts() GrantPrivilegesToAccountRoleTimeoutsOutputReference
	TimeoutsInput() interface{}
	WithGrantOption() interface{}
	SetWithGrantOption(val interface{})
	WithGrantOptionInput() interface{}
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
	PutOnAccountObject(value *GrantPrivilegesToAccountRoleOnAccountObject)
	PutOnSchema(value *GrantPrivilegesToAccountRoleOnSchema)
	PutOnSchemaObject(value *GrantPrivilegesToAccountRoleOnSchemaObject)
	PutTimeouts(value *GrantPrivilegesToAccountRoleTimeouts)
	ResetAllPrivileges()
	ResetAlwaysApply()
	ResetAlwaysApplyTrigger()
	ResetId()
	ResetOnAccount()
	ResetOnAccountObject()
	ResetOnSchema()
	ResetOnSchemaObject()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivileges()
	ResetTimeouts()
	ResetWithGrantOption()
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

// The jsii proxy struct for GrantPrivilegesToAccountRole
type jsiiProxy_GrantPrivilegesToAccountRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) AccountRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) AccountRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) AllPrivileges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPrivileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) AllPrivilegesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPrivilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) AlwaysApply() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysApply",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) AlwaysApplyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysApplyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) AlwaysApplyTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysApplyTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) AlwaysApplyTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysApplyTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) OnAccount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) OnAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) OnAccountObject() GrantPrivilegesToAccountRoleOnAccountObjectOutputReference {
	var returns GrantPrivilegesToAccountRoleOnAccountObjectOutputReference
	_jsii_.Get(
		j,
		"onAccountObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) OnAccountObjectInput() *GrantPrivilegesToAccountRoleOnAccountObject {
	var returns *GrantPrivilegesToAccountRoleOnAccountObject
	_jsii_.Get(
		j,
		"onAccountObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) OnSchema() GrantPrivilegesToAccountRoleOnSchemaOutputReference {
	var returns GrantPrivilegesToAccountRoleOnSchemaOutputReference
	_jsii_.Get(
		j,
		"onSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) OnSchemaInput() *GrantPrivilegesToAccountRoleOnSchema {
	var returns *GrantPrivilegesToAccountRoleOnSchema
	_jsii_.Get(
		j,
		"onSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) OnSchemaObject() GrantPrivilegesToAccountRoleOnSchemaObjectOutputReference {
	var returns GrantPrivilegesToAccountRoleOnSchemaObjectOutputReference
	_jsii_.Get(
		j,
		"onSchemaObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) OnSchemaObjectInput() *GrantPrivilegesToAccountRoleOnSchemaObject {
	var returns *GrantPrivilegesToAccountRoleOnSchemaObject
	_jsii_.Get(
		j,
		"onSchemaObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Privileges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) PrivilegesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) Timeouts() GrantPrivilegesToAccountRoleTimeoutsOutputReference {
	var returns GrantPrivilegesToAccountRoleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) WithGrantOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole) WithGrantOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/grant_privileges_to_account_role snowflake_grant_privileges_to_account_role} Resource.
func NewGrantPrivilegesToAccountRole(scope constructs.Construct, id *string, config *GrantPrivilegesToAccountRoleConfig) GrantPrivilegesToAccountRole {
	_init_.Initialize()

	if err := validateNewGrantPrivilegesToAccountRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrantPrivilegesToAccountRole{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/grant_privileges_to_account_role snowflake_grant_privileges_to_account_role} Resource.
func NewGrantPrivilegesToAccountRole_Override(g GrantPrivilegesToAccountRole, scope constructs.Construct, id *string, config *GrantPrivilegesToAccountRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRole",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetAccountRoleName(val *string) {
	if err := j.validateSetAccountRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountRoleName",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetAllPrivileges(val interface{}) {
	if err := j.validateSetAllPrivilegesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allPrivileges",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetAlwaysApply(val interface{}) {
	if err := j.validateSetAlwaysApplyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysApply",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetAlwaysApplyTrigger(val *string) {
	if err := j.validateSetAlwaysApplyTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysApplyTrigger",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetOnAccount(val interface{}) {
	if err := j.validateSetOnAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onAccount",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetPrivileges(val *[]*string) {
	if err := j.validateSetPrivilegesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileges",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRole)SetWithGrantOption(val interface{}) {
	if err := j.validateSetWithGrantOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withGrantOption",
		val,
	)
}

// Generates CDKTF code for importing a GrantPrivilegesToAccountRole resource upon running "cdktf plan <stack-name>".
func GrantPrivilegesToAccountRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGrantPrivilegesToAccountRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRole",
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
func GrantPrivilegesToAccountRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToAccountRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrantPrivilegesToAccountRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToAccountRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrantPrivilegesToAccountRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToAccountRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GrantPrivilegesToAccountRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) PutOnAccountObject(value *GrantPrivilegesToAccountRoleOnAccountObject) {
	if err := g.validatePutOnAccountObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnAccountObject",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) PutOnSchema(value *GrantPrivilegesToAccountRoleOnSchema) {
	if err := g.validatePutOnSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnSchema",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) PutOnSchemaObject(value *GrantPrivilegesToAccountRoleOnSchemaObject) {
	if err := g.validatePutOnSchemaObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnSchemaObject",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) PutTimeouts(value *GrantPrivilegesToAccountRoleTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetAllPrivileges() {
	_jsii_.InvokeVoid(
		g,
		"resetAllPrivileges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetAlwaysApply() {
	_jsii_.InvokeVoid(
		g,
		"resetAlwaysApply",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetAlwaysApplyTrigger() {
	_jsii_.InvokeVoid(
		g,
		"resetAlwaysApplyTrigger",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetOnAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetOnAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetOnAccountObject() {
	_jsii_.InvokeVoid(
		g,
		"resetOnAccountObject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetOnSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetOnSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetOnSchemaObject() {
	_jsii_.InvokeVoid(
		g,
		"resetOnSchemaObject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetPrivileges() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivileges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ResetWithGrantOption() {
	_jsii_.InvokeVoid(
		g,
		"resetWithGrantOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

