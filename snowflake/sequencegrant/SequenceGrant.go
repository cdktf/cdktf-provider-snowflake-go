// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sequencegrant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v9/sequencegrant/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.70.1/docs/resources/sequence_grant snowflake_sequence_grant}.
type SequenceGrant interface {
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableMultipleGrants() interface{}
	SetEnableMultipleGrants(val interface{})
	EnableMultipleGrantsInput() interface{}
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
	OnAll() interface{}
	SetOnAll(val interface{})
	OnAllInput() interface{}
	OnFuture() interface{}
	SetOnFuture(val interface{})
	OnFutureInput() interface{}
	Privilege() *string
	SetPrivilege(val *string)
	PrivilegeInput() *string
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
	RevertOwnershipToRoleName() *string
	SetRevertOwnershipToRoleName(val *string)
	RevertOwnershipToRoleNameInput() *string
	Roles() *[]*string
	SetRoles(val *[]*string)
	RolesInput() *[]*string
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
	SequenceName() *string
	SetSequenceName(val *string)
	SequenceNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithGrantOption() interface{}
	SetWithGrantOption(val interface{})
	WithGrantOptionInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetEnableMultipleGrants()
	ResetId()
	ResetOnAll()
	ResetOnFuture()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivilege()
	ResetRevertOwnershipToRoleName()
	ResetSchemaName()
	ResetSequenceName()
	ResetWithGrantOption()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SequenceGrant
type jsiiProxy_SequenceGrant struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SequenceGrant) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) EnableMultipleGrants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) EnableMultipleGrantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleGrantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) OnAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) OnAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) OnFuture() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFuture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) OnFutureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFutureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Privilege() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilege",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) PrivilegeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) RevertOwnershipToRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revertOwnershipToRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) RevertOwnershipToRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revertOwnershipToRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) RolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) SequenceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sequenceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) SequenceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sequenceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) WithGrantOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SequenceGrant) WithGrantOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.70.1/docs/resources/sequence_grant snowflake_sequence_grant} Resource.
func NewSequenceGrant(scope constructs.Construct, id *string, config *SequenceGrantConfig) SequenceGrant {
	_init_.Initialize()

	if err := validateNewSequenceGrantParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SequenceGrant{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.sequenceGrant.SequenceGrant",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.70.1/docs/resources/sequence_grant snowflake_sequence_grant} Resource.
func NewSequenceGrant_Override(s SequenceGrant, scope constructs.Construct, id *string, config *SequenceGrantConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.sequenceGrant.SequenceGrant",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SequenceGrant)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetEnableMultipleGrants(val interface{}) {
	if err := j.validateSetEnableMultipleGrantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMultipleGrants",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetOnAll(val interface{}) {
	if err := j.validateSetOnAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onAll",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetOnFuture(val interface{}) {
	if err := j.validateSetOnFutureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFuture",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetPrivilege(val *string) {
	if err := j.validateSetPrivilegeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilege",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetRevertOwnershipToRoleName(val *string) {
	if err := j.validateSetRevertOwnershipToRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revertOwnershipToRoleName",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetRoles(val *[]*string) {
	if err := j.validateSetRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetSequenceName(val *string) {
	if err := j.validateSetSequenceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sequenceName",
		val,
	)
}

func (j *jsiiProxy_SequenceGrant)SetWithGrantOption(val interface{}) {
	if err := j.validateSetWithGrantOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withGrantOption",
		val,
	)
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
func SequenceGrant_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSequenceGrant_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.sequenceGrant.SequenceGrant",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SequenceGrant_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSequenceGrant_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.sequenceGrant.SequenceGrant",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SequenceGrant_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSequenceGrant_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.sequenceGrant.SequenceGrant",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SequenceGrant_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.sequenceGrant.SequenceGrant",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SequenceGrant) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SequenceGrant) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SequenceGrant) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SequenceGrant) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SequenceGrant) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SequenceGrant) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SequenceGrant) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SequenceGrant) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SequenceGrant) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SequenceGrant) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SequenceGrant) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SequenceGrant) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SequenceGrant) ResetEnableMultipleGrants() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableMultipleGrants",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetOnAll() {
	_jsii_.InvokeVoid(
		s,
		"resetOnAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetOnFuture() {
	_jsii_.InvokeVoid(
		s,
		"resetOnFuture",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetPrivilege() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivilege",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetRevertOwnershipToRoleName() {
	_jsii_.InvokeVoid(
		s,
		"resetRevertOwnershipToRoleName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetSchemaName() {
	_jsii_.InvokeVoid(
		s,
		"resetSchemaName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetSequenceName() {
	_jsii_.InvokeVoid(
		s,
		"resetSequenceName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) ResetWithGrantOption() {
	_jsii_.InvokeVoid(
		s,
		"resetWithGrantOption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SequenceGrant) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SequenceGrant) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SequenceGrant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SequenceGrant) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

