// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakedatabases

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/datasnowflakedatabases/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/databases snowflake_databases}.
type DataSnowflakeDatabases interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Databases() DataSnowflakeDatabasesDatabasesList
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
	Like() *string
	SetLike(val *string)
	LikeInput() *string
	Limit() DataSnowflakeDatabasesLimitOutputReference
	LimitInput() *DataSnowflakeDatabasesLimit
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	StartsWith() *string
	SetStartsWith(val *string)
	StartsWithInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithDescribe() interface{}
	SetWithDescribe(val interface{})
	WithDescribeInput() interface{}
	WithParameters() interface{}
	SetWithParameters(val interface{})
	WithParametersInput() interface{}
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
	PutLimit(value *DataSnowflakeDatabasesLimit)
	ResetId()
	ResetLike()
	ResetLimit()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartsWith()
	ResetWithDescribe()
	ResetWithParameters()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataSnowflakeDatabases
type jsiiProxy_DataSnowflakeDatabases struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataSnowflakeDatabases) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Databases() DataSnowflakeDatabasesDatabasesList {
	var returns DataSnowflakeDatabasesDatabasesList
	_jsii_.Get(
		j,
		"databases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Like() *string {
	var returns *string
	_jsii_.Get(
		j,
		"like",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) LikeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Limit() DataSnowflakeDatabasesLimitOutputReference {
	var returns DataSnowflakeDatabasesLimitOutputReference
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) LimitInput() *DataSnowflakeDatabasesLimit {
	var returns *DataSnowflakeDatabasesLimit
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) StartsWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) StartsWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) WithDescribe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDescribe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) WithDescribeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDescribeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) WithParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabases) WithParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withParametersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/databases snowflake_databases} Data Source.
func NewDataSnowflakeDatabases(scope constructs.Construct, id *string, config *DataSnowflakeDatabasesConfig) DataSnowflakeDatabases {
	_init_.Initialize()

	if err := validateNewDataSnowflakeDatabasesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeDatabases{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabases",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/databases snowflake_databases} Data Source.
func NewDataSnowflakeDatabases_Override(d DataSnowflakeDatabases, scope constructs.Construct, id *string, config *DataSnowflakeDatabasesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabases",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetLike(val *string) {
	if err := j.validateSetLikeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"like",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetStartsWith(val *string) {
	if err := j.validateSetStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startsWith",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetWithDescribe(val interface{}) {
	if err := j.validateSetWithDescribeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withDescribe",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabases)SetWithParameters(val interface{}) {
	if err := j.validateSetWithParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withParameters",
		val,
	)
}

// Generates CDKTF code for importing a DataSnowflakeDatabases resource upon running "cdktf plan <stack-name>".
func DataSnowflakeDatabases_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataSnowflakeDatabases_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabases",
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
func DataSnowflakeDatabases_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeDatabases_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabases",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataSnowflakeDatabases_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeDatabases_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabases",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataSnowflakeDatabases_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeDatabases_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabases",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataSnowflakeDatabases_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabases",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataSnowflakeDatabases) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeDatabases) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeDatabases) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeDatabases) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeDatabases) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeDatabases) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeDatabases) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeDatabases) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeDatabases) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeDatabases) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeDatabases) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) PutLimit(value *DataSnowflakeDatabasesLimit) {
	if err := d.validatePutLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLimit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) ResetLike() {
	_jsii_.InvokeVoid(
		d,
		"resetLike",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) ResetStartsWith() {
	_jsii_.InvokeVoid(
		d,
		"resetStartsWith",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) ResetWithDescribe() {
	_jsii_.InvokeVoid(
		d,
		"resetWithDescribe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) ResetWithParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetWithParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeDatabases) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeDatabases) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeDatabases) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeDatabases) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeDatabases) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeDatabases) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

