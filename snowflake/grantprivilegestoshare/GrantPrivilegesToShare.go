// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoshare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/grantprivilegestoshare/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/grant_privileges_to_share snowflake_grant_privileges_to_share}.
type GrantPrivilegesToShare interface {
	cdktf.TerraformResource
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
	OnAllTablesInSchema() *string
	SetOnAllTablesInSchema(val *string)
	OnAllTablesInSchemaInput() *string
	OnDatabase() *string
	SetOnDatabase(val *string)
	OnDatabaseInput() *string
	OnFunction() *string
	SetOnFunction(val *string)
	OnFunctionInput() *string
	OnSchema() *string
	SetOnSchema(val *string)
	OnSchemaInput() *string
	OnTable() *string
	SetOnTable(val *string)
	OnTableInput() *string
	OnTag() *string
	SetOnTag(val *string)
	OnTagInput() *string
	OnView() *string
	SetOnView(val *string)
	OnViewInput() *string
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
	Timeouts() GrantPrivilegesToShareTimeoutsOutputReference
	TimeoutsInput() interface{}
	ToShare() *string
	SetToShare(val *string)
	ToShareInput() *string
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
	PutTimeouts(value *GrantPrivilegesToShareTimeouts)
	ResetId()
	ResetOnAllTablesInSchema()
	ResetOnDatabase()
	ResetOnFunction()
	ResetOnSchema()
	ResetOnTable()
	ResetOnTag()
	ResetOnView()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for GrantPrivilegesToShare
type jsiiProxy_GrantPrivilegesToShare struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GrantPrivilegesToShare) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnAllTablesInSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onAllTablesInSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnAllTablesInSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onAllTablesInSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) OnViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onViewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Privileges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) PrivilegesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) Timeouts() GrantPrivilegesToShareTimeoutsOutputReference {
	var returns GrantPrivilegesToShareTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) ToShare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToShare) ToShareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toShareInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/grant_privileges_to_share snowflake_grant_privileges_to_share} Resource.
func NewGrantPrivilegesToShare(scope constructs.Construct, id *string, config *GrantPrivilegesToShareConfig) GrantPrivilegesToShare {
	_init_.Initialize()

	if err := validateNewGrantPrivilegesToShareParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrantPrivilegesToShare{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/grant_privileges_to_share snowflake_grant_privileges_to_share} Resource.
func NewGrantPrivilegesToShare_Override(g GrantPrivilegesToShare, scope constructs.Construct, id *string, config *GrantPrivilegesToShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShare",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetOnAllTablesInSchema(val *string) {
	if err := j.validateSetOnAllTablesInSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onAllTablesInSchema",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetOnDatabase(val *string) {
	if err := j.validateSetOnDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDatabase",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetOnFunction(val *string) {
	if err := j.validateSetOnFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFunction",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetOnSchema(val *string) {
	if err := j.validateSetOnSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onSchema",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetOnTable(val *string) {
	if err := j.validateSetOnTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onTable",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetOnTag(val *string) {
	if err := j.validateSetOnTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onTag",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetOnView(val *string) {
	if err := j.validateSetOnViewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onView",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetPrivileges(val *[]*string) {
	if err := j.validateSetPrivilegesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileges",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToShare)SetToShare(val *string) {
	if err := j.validateSetToShareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toShare",
		val,
	)
}

// Generates CDKTF code for importing a GrantPrivilegesToShare resource upon running "cdktf plan <stack-name>".
func GrantPrivilegesToShare_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGrantPrivilegesToShare_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShare",
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
func GrantPrivilegesToShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToShare_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrantPrivilegesToShare_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToShare_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShare",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GrantPrivilegesToShare_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrantPrivilegesToShare_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShare",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GrantPrivilegesToShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.grantPrivilegesToShare.GrantPrivilegesToShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GrantPrivilegesToShare) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GrantPrivilegesToShare) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToShare) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GrantPrivilegesToShare) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GrantPrivilegesToShare) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GrantPrivilegesToShare) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToShare) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToShare) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GrantPrivilegesToShare) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GrantPrivilegesToShare) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToShare) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToShare) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) PutTimeouts(value *GrantPrivilegesToShareTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetOnAllTablesInSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetOnAllTablesInSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetOnDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetOnDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetOnFunction() {
	_jsii_.InvokeVoid(
		g,
		"resetOnFunction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetOnSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetOnSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetOnTable() {
	_jsii_.InvokeVoid(
		g,
		"resetOnTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetOnTag() {
	_jsii_.InvokeVoid(
		g,
		"resetOnTag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetOnView() {
	_jsii_.InvokeVoid(
		g,
		"resetOnView",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToShare) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToShare) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

