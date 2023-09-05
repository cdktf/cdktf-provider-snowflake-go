// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v9/datasnowflakegrants/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.70.1/docs/data-sources/grants snowflake_grants}.
type DataSnowflakeGrants interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	FutureGrantsIn() DataSnowflakeGrantsFutureGrantsInOutputReference
	FutureGrantsInInput() *DataSnowflakeGrantsFutureGrantsIn
	FutureGrantsTo() DataSnowflakeGrantsFutureGrantsToOutputReference
	FutureGrantsToInput() *DataSnowflakeGrantsFutureGrantsTo
	Grants() DataSnowflakeGrantsGrantsList
	GrantsOf() DataSnowflakeGrantsGrantsOfOutputReference
	GrantsOfInput() *DataSnowflakeGrantsGrantsOf
	GrantsOn() DataSnowflakeGrantsGrantsOnOutputReference
	GrantsOnInput() *DataSnowflakeGrantsGrantsOn
	GrantsTo() DataSnowflakeGrantsGrantsToOutputReference
	GrantsToInput() *DataSnowflakeGrantsGrantsTo
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutFutureGrantsIn(value *DataSnowflakeGrantsFutureGrantsIn)
	PutFutureGrantsTo(value *DataSnowflakeGrantsFutureGrantsTo)
	PutGrantsOf(value *DataSnowflakeGrantsGrantsOf)
	PutGrantsOn(value *DataSnowflakeGrantsGrantsOn)
	PutGrantsTo(value *DataSnowflakeGrantsGrantsTo)
	ResetFutureGrantsIn()
	ResetFutureGrantsTo()
	ResetGrantsOf()
	ResetGrantsOn()
	ResetGrantsTo()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataSnowflakeGrants
type jsiiProxy_DataSnowflakeGrants struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataSnowflakeGrants) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) FutureGrantsIn() DataSnowflakeGrantsFutureGrantsInOutputReference {
	var returns DataSnowflakeGrantsFutureGrantsInOutputReference
	_jsii_.Get(
		j,
		"futureGrantsIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) FutureGrantsInInput() *DataSnowflakeGrantsFutureGrantsIn {
	var returns *DataSnowflakeGrantsFutureGrantsIn
	_jsii_.Get(
		j,
		"futureGrantsInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) FutureGrantsTo() DataSnowflakeGrantsFutureGrantsToOutputReference {
	var returns DataSnowflakeGrantsFutureGrantsToOutputReference
	_jsii_.Get(
		j,
		"futureGrantsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) FutureGrantsToInput() *DataSnowflakeGrantsFutureGrantsTo {
	var returns *DataSnowflakeGrantsFutureGrantsTo
	_jsii_.Get(
		j,
		"futureGrantsToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) Grants() DataSnowflakeGrantsGrantsList {
	var returns DataSnowflakeGrantsGrantsList
	_jsii_.Get(
		j,
		"grants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) GrantsOf() DataSnowflakeGrantsGrantsOfOutputReference {
	var returns DataSnowflakeGrantsGrantsOfOutputReference
	_jsii_.Get(
		j,
		"grantsOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) GrantsOfInput() *DataSnowflakeGrantsGrantsOf {
	var returns *DataSnowflakeGrantsGrantsOf
	_jsii_.Get(
		j,
		"grantsOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) GrantsOn() DataSnowflakeGrantsGrantsOnOutputReference {
	var returns DataSnowflakeGrantsGrantsOnOutputReference
	_jsii_.Get(
		j,
		"grantsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) GrantsOnInput() *DataSnowflakeGrantsGrantsOn {
	var returns *DataSnowflakeGrantsGrantsOn
	_jsii_.Get(
		j,
		"grantsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) GrantsTo() DataSnowflakeGrantsGrantsToOutputReference {
	var returns DataSnowflakeGrantsGrantsToOutputReference
	_jsii_.Get(
		j,
		"grantsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) GrantsToInput() *DataSnowflakeGrantsGrantsTo {
	var returns *DataSnowflakeGrantsGrantsTo
	_jsii_.Get(
		j,
		"grantsToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrants) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.70.1/docs/data-sources/grants snowflake_grants} Data Source.
func NewDataSnowflakeGrants(scope constructs.Construct, id *string, config *DataSnowflakeGrantsConfig) DataSnowflakeGrants {
	_init_.Initialize()

	if err := validateNewDataSnowflakeGrantsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeGrants{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrants",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.70.1/docs/data-sources/grants snowflake_grants} Data Source.
func NewDataSnowflakeGrants_Override(d DataSnowflakeGrants, scope constructs.Construct, id *string, config *DataSnowflakeGrantsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrants",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeGrants)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrants)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrants)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrants)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrants)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrants)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataSnowflakeGrants_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeGrants_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrants",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataSnowflakeGrants_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeGrants_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrants",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataSnowflakeGrants_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeGrants_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrants",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataSnowflakeGrants_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrants",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) PutFutureGrantsIn(value *DataSnowflakeGrantsFutureGrantsIn) {
	if err := d.validatePutFutureGrantsInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFutureGrantsIn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) PutFutureGrantsTo(value *DataSnowflakeGrantsFutureGrantsTo) {
	if err := d.validatePutFutureGrantsToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFutureGrantsTo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) PutGrantsOf(value *DataSnowflakeGrantsGrantsOf) {
	if err := d.validatePutGrantsOfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGrantsOf",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) PutGrantsOn(value *DataSnowflakeGrantsGrantsOn) {
	if err := d.validatePutGrantsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGrantsOn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) PutGrantsTo(value *DataSnowflakeGrantsGrantsTo) {
	if err := d.validatePutGrantsToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGrantsTo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) ResetFutureGrantsIn() {
	_jsii_.InvokeVoid(
		d,
		"resetFutureGrantsIn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) ResetFutureGrantsTo() {
	_jsii_.InvokeVoid(
		d,
		"resetFutureGrantsTo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) ResetGrantsOf() {
	_jsii_.InvokeVoid(
		d,
		"resetGrantsOf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) ResetGrantsOn() {
	_jsii_.InvokeVoid(
		d,
		"resetGrantsOn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) ResetGrantsTo() {
	_jsii_.InvokeVoid(
		d,
		"resetGrantsTo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrants) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrants) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

