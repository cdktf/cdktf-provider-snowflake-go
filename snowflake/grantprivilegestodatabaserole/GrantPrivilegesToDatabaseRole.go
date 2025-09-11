// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestodatabaserole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/grantprivilegestodatabaserole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/grant_privileges_to_database_role snowflake_grant_privileges_to_database_role}.
type GrantPrivilegesToDatabaseRole interface {
	cdktf.TerraformResource
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
	DatabaseRoleName() *string
	SetDatabaseRoleName(val *string)
	DatabaseRoleNameInput() *string
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
	OnDatabase() *string
	SetOnDatabase(val *string)
	OnDatabaseInput() *string
	OnSchema() GrantPrivilegesToDatabaseRoleOnSchemaOutputReference
	OnSchemaInput() *GrantPrivilegesToDatabaseRoleOnSchema
	OnSchemaObject() GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference
	OnSchemaObjectInput() *GrantPrivilegesToDatabaseRoleOnSchemaObject
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
	Timeouts() GrantPrivilegesToDatabaseRoleTimeoutsOutputReference
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
	PutOnSchema(value *GrantPrivilegesToDatabaseRoleOnSchema)
	PutOnSchemaObject(value *GrantPrivilegesToDatabaseRoleOnSchemaObject)
	PutTimeouts(value *GrantPrivilegesToDatabaseRoleTimeouts)
	ResetAllPrivileges()
	ResetAlwaysApply()
	ResetAlwaysApplyTrigger()
	ResetId()
	ResetOnDatabase()
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

// The jsii proxy struct for GrantPrivilegesToDatabaseRole
type jsiiProxy_GrantPrivilegesToDatabaseRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) AllPrivileges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPrivileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) AllPrivilegesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPrivilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) AlwaysApply() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysApply",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) AlwaysApplyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysApplyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) AlwaysApplyTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysApplyTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) AlwaysApplyTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alwaysApplyTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) DatabaseRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) DatabaseRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) OnDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) OnDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) OnSchema() GrantPrivilegesToDatabaseRoleOnSchemaOutputReference {
	var returns GrantPrivilegesToDatabaseRoleOnSchemaOutputReference
	_jsii_.Get(
		j,
		"onSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) OnSchemaInput() *GrantPrivilegesToDatabaseRoleOnSchema {
	var returns *GrantPrivilegesToDatabaseRoleOnSchema
	_jsii_.Get(
		j,
		"onSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) OnSchemaObject() GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference {
	var returns GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference
	_jsii_.Get(
		j,
		"onSchemaObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) OnSchemaObjectInput() *GrantPrivilegesToDatabaseRoleOnSchemaObject {
	var returns *GrantPrivilegesToDatabaseRoleOnSchemaObject
	_jsii_.Get(
		j,
		"onSchemaObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Privileges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) PrivilegesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) Timeouts() GrantPrivilegesToDatabaseRoleTimeoutsOutputReference {
	var returns GrantPrivilegesToDatabaseRoleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) WithGrantOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole) WithGrantOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/grant_privileges_to_database_role snowflake_grant_privileges_to_database_role} Resource.
func NewGrantPrivilegesToDatabaseRole(scope constructs.Construct, id *string, config *GrantPrivilegesToDatabaseRoleConfig) GrantPrivilegesToDatabaseRole {
	_init_.Initialize()

	if err := validateNewGrantPrivilegesToDatabaseRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrantPrivilegesToDatabaseRole{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/grant_privileges_to_database_role snowflake_grant_privileges_to_database_role} Resource.
func NewGrantPrivilegesToDatabaseRole_Override(g GrantPrivilegesToDatabaseRole, scope constructs.Construct, id *string, config *GrantPrivilegesToDatabaseRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRole",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetAllPrivileges(val interface{}) {
	if err := j.validateSetAllPrivilegesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allPrivileges",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetAlwaysApply(val interface{}) {
	if err := j.validateSetAlwaysApplyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysApply",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetAlwaysApplyTrigger(val *string) {
	if err := j.validateSetAlwaysApplyTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysApplyTrigger",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetDatabaseRoleName(val *string) {
	if err := j.validateSetDatabaseRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseRoleName",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetOnDatabase(val *string) {
	if err := j.validateSetOnDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDatabase",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetPrivileges(val *[]*string) {
	if err := j.validateSetPrivilegesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileges",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRole)SetWithGrantOption(val interface{}) {
	if err := j.validateSetWithGrantOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withGrantOption",
		val,
	)
}

// Generates CDKTF code for importing a GrantPrivilegesToDatabaseRole resource upon running "cdktf plan <stack-name>".
func GrantPrivilegesToDatabaseRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGrantPrivilegesToDatabaseRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRole",
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
func GrantPrivilegesToDatabaseRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToDatabaseRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrantPrivilegesToDatabaseRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToDatabaseRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrantPrivilegesToDatabaseRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToDatabaseRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GrantPrivilegesToDatabaseRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) PutOnSchema(value *GrantPrivilegesToDatabaseRoleOnSchema) {
	if err := g.validatePutOnSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnSchema",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) PutOnSchemaObject(value *GrantPrivilegesToDatabaseRoleOnSchemaObject) {
	if err := g.validatePutOnSchemaObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnSchemaObject",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) PutTimeouts(value *GrantPrivilegesToDatabaseRoleTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetAllPrivileges() {
	_jsii_.InvokeVoid(
		g,
		"resetAllPrivileges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetAlwaysApply() {
	_jsii_.InvokeVoid(
		g,
		"resetAlwaysApply",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetAlwaysApplyTrigger() {
	_jsii_.InvokeVoid(
		g,
		"resetAlwaysApplyTrigger",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetOnDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetOnDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetOnSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetOnSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetOnSchemaObject() {
	_jsii_.InvokeVoid(
		g,
		"resetOnSchemaObject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetPrivileges() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivileges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ResetWithGrantOption() {
	_jsii_.InvokeVoid(
		g,
		"resetWithGrantOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

