// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakesecrets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/datasnowflakesecrets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeSecretsInOutputReference interface {
	cdktf.ComplexObject
	Account() interface{}
	SetAccount(val interface{})
	AccountInput() interface{}
	Application() *string
	SetApplication(val *string)
	ApplicationInput() *string
	ApplicationPackage() *string
	SetApplicationPackage(val *string)
	ApplicationPackageInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataSnowflakeSecretsIn
	SetInternalValue(val *DataSnowflakeSecretsIn)
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAccount()
	ResetApplication()
	ResetApplicationPackage()
	ResetDatabase()
	ResetSchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSnowflakeSecretsInOutputReference
type jsiiProxy_DataSnowflakeSecretsInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) Account() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) AccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) Application() *string {
	var returns *string
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) ApplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) ApplicationPackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) ApplicationPackageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationPackageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) InternalValue() *DataSnowflakeSecretsIn {
	var returns *DataSnowflakeSecretsIn
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSnowflakeSecretsInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataSnowflakeSecretsInOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeSecretsInOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeSecretsInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeSecrets.DataSnowflakeSecretsInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataSnowflakeSecretsInOutputReference_Override(d DataSnowflakeSecretsInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeSecrets.DataSnowflakeSecretsInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetAccount(val interface{}) {
	if err := j.validateSetAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetApplication(val *string) {
	if err := j.validateSetApplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"application",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetApplicationPackage(val *string) {
	if err := j.validateSetApplicationPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationPackage",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetInternalValue(val *DataSnowflakeSecretsIn) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecretsInOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) ResetApplication() {
	_jsii_.InvokeVoid(
		d,
		"resetApplication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) ResetApplicationPackage() {
	_jsii_.InvokeVoid(
		d,
		"resetApplicationPackage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecretsInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

