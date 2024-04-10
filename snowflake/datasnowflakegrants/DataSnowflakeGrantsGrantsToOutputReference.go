// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/datasnowflakegrants/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeGrantsGrantsToOutputReference interface {
	cdktf.ComplexObject
	AccountRole() *string
	SetAccountRole(val *string)
	AccountRoleInput() *string
	Application() *string
	SetApplication(val *string)
	ApplicationInput() *string
	ApplicationRole() *string
	SetApplicationRole(val *string)
	ApplicationRoleInput() *string
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
	DatabaseRole() *string
	SetDatabaseRole(val *string)
	DatabaseRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataSnowflakeGrantsGrantsTo
	SetInternalValue(val *DataSnowflakeGrantsGrantsTo)
	Share() DataSnowflakeGrantsGrantsToShareOutputReference
	ShareInput() *DataSnowflakeGrantsGrantsToShare
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	PutShare(value *DataSnowflakeGrantsGrantsToShare)
	ResetAccountRole()
	ResetApplication()
	ResetApplicationRole()
	ResetDatabaseRole()
	ResetShare()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSnowflakeGrantsGrantsToOutputReference
type jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) AccountRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) AccountRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) Application() *string {
	var returns *string
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ApplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ApplicationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ApplicationRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) DatabaseRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) DatabaseRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) InternalValue() *DataSnowflakeGrantsGrantsTo {
	var returns *DataSnowflakeGrantsGrantsTo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) Share() DataSnowflakeGrantsGrantsToShareOutputReference {
	var returns DataSnowflakeGrantsGrantsToShareOutputReference
	_jsii_.Get(
		j,
		"share",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ShareInput() *DataSnowflakeGrantsGrantsToShare {
	var returns *DataSnowflakeGrantsGrantsToShare
	_jsii_.Get(
		j,
		"shareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDataSnowflakeGrantsGrantsToOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataSnowflakeGrantsGrantsToOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeGrantsGrantsToOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrantsGrantsToOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataSnowflakeGrantsGrantsToOutputReference_Override(d DataSnowflakeGrantsGrantsToOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrantsGrantsToOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetAccountRole(val *string) {
	if err := j.validateSetAccountRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountRole",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetApplication(val *string) {
	if err := j.validateSetApplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"application",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetApplicationRole(val *string) {
	if err := j.validateSetApplicationRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationRole",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetDatabaseRole(val *string) {
	if err := j.validateSetDatabaseRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseRole",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetInternalValue(val *DataSnowflakeGrantsGrantsTo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) PutShare(value *DataSnowflakeGrantsGrantsToShare) {
	if err := d.validatePutShareParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putShare",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ResetAccountRole() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ResetApplication() {
	_jsii_.InvokeVoid(
		d,
		"resetApplication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ResetApplicationRole() {
	_jsii_.InvokeVoid(
		d,
		"resetApplicationRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ResetDatabaseRole() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ResetShare() {
	_jsii_.InvokeVoid(
		d,
		"resetShare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsToOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

