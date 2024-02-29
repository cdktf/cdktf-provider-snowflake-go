// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestorole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/grantprivilegestorole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.0/docs/resources/grant_privileges_to_role snowflake_grant_privileges_to_role}.
type GrantPrivilegesToRole interface {
	cdktf.TerraformResource
	AllPrivileges() interface{}
	SetAllPrivileges(val interface{})
	AllPrivilegesInput() interface{}
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
	OnAccountObject() GrantPrivilegesToRoleOnAccountObjectOutputReference
	OnAccountObjectInput() *GrantPrivilegesToRoleOnAccountObject
	OnSchema() GrantPrivilegesToRoleOnSchemaOutputReference
	OnSchemaInput() *GrantPrivilegesToRoleOnSchema
	OnSchemaObject() GrantPrivilegesToRoleOnSchemaObjectOutputReference
	OnSchemaObjectInput() *GrantPrivilegesToRoleOnSchemaObject
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
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutOnAccountObject(value *GrantPrivilegesToRoleOnAccountObject)
	PutOnSchema(value *GrantPrivilegesToRoleOnSchema)
	PutOnSchemaObject(value *GrantPrivilegesToRoleOnSchemaObject)
	ResetAllPrivileges()
	ResetId()
	ResetOnAccount()
	ResetOnAccountObject()
	ResetOnSchema()
	ResetOnSchemaObject()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivileges()
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

// The jsii proxy struct for GrantPrivilegesToRole
type jsiiProxy_GrantPrivilegesToRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GrantPrivilegesToRole) AllPrivileges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPrivileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) AllPrivilegesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPrivilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) OnAccount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) OnAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) OnAccountObject() GrantPrivilegesToRoleOnAccountObjectOutputReference {
	var returns GrantPrivilegesToRoleOnAccountObjectOutputReference
	_jsii_.Get(
		j,
		"onAccountObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) OnAccountObjectInput() *GrantPrivilegesToRoleOnAccountObject {
	var returns *GrantPrivilegesToRoleOnAccountObject
	_jsii_.Get(
		j,
		"onAccountObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) OnSchema() GrantPrivilegesToRoleOnSchemaOutputReference {
	var returns GrantPrivilegesToRoleOnSchemaOutputReference
	_jsii_.Get(
		j,
		"onSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) OnSchemaInput() *GrantPrivilegesToRoleOnSchema {
	var returns *GrantPrivilegesToRoleOnSchema
	_jsii_.Get(
		j,
		"onSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) OnSchemaObject() GrantPrivilegesToRoleOnSchemaObjectOutputReference {
	var returns GrantPrivilegesToRoleOnSchemaObjectOutputReference
	_jsii_.Get(
		j,
		"onSchemaObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) OnSchemaObjectInput() *GrantPrivilegesToRoleOnSchemaObject {
	var returns *GrantPrivilegesToRoleOnSchemaObject
	_jsii_.Get(
		j,
		"onSchemaObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Privileges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) PrivilegesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) WithGrantOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRole) WithGrantOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.0/docs/resources/grant_privileges_to_role snowflake_grant_privileges_to_role} Resource.
func NewGrantPrivilegesToRole(scope constructs.Construct, id *string, config *GrantPrivilegesToRoleConfig) GrantPrivilegesToRole {
	_init_.Initialize()

	if err := validateNewGrantPrivilegesToRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrantPrivilegesToRole{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.0/docs/resources/grant_privileges_to_role snowflake_grant_privileges_to_role} Resource.
func NewGrantPrivilegesToRole_Override(g GrantPrivilegesToRole, scope constructs.Construct, id *string, config *GrantPrivilegesToRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRole",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetAllPrivileges(val interface{}) {
	if err := j.validateSetAllPrivilegesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allPrivileges",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetOnAccount(val interface{}) {
	if err := j.validateSetOnAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onAccount",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetPrivileges(val *[]*string) {
	if err := j.validateSetPrivilegesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileges",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRole)SetWithGrantOption(val interface{}) {
	if err := j.validateSetWithGrantOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withGrantOption",
		val,
	)
}

// Generates CDKTF code for importing a GrantPrivilegesToRole resource upon running "cdktf plan <stack-name>".
func GrantPrivilegesToRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGrantPrivilegesToRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRole",
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
func GrantPrivilegesToRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrantPrivilegesToRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrantPrivilegesToRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GrantPrivilegesToRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRole) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GrantPrivilegesToRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GrantPrivilegesToRole) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GrantPrivilegesToRole) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GrantPrivilegesToRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToRole) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GrantPrivilegesToRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GrantPrivilegesToRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRole) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToRole) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) PutOnAccountObject(value *GrantPrivilegesToRoleOnAccountObject) {
	if err := g.validatePutOnAccountObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnAccountObject",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) PutOnSchema(value *GrantPrivilegesToRoleOnSchema) {
	if err := g.validatePutOnSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnSchema",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) PutOnSchemaObject(value *GrantPrivilegesToRoleOnSchemaObject) {
	if err := g.validatePutOnSchemaObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnSchemaObject",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetAllPrivileges() {
	_jsii_.InvokeVoid(
		g,
		"resetAllPrivileges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetOnAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetOnAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetOnAccountObject() {
	_jsii_.InvokeVoid(
		g,
		"resetOnAccountObject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetOnSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetOnSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetOnSchemaObject() {
	_jsii_.InvokeVoid(
		g,
		"resetOnSchemaObject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetPrivileges() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivileges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) ResetWithGrantOption() {
	_jsii_.InvokeVoid(
		g,
		"resetWithGrantOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

