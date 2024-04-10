// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/datasnowflakegrants/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeGrantsGrantsOfOutputReference interface {
	cdktf.ComplexObject
	AccountRole() *string
	SetAccountRole(val *string)
	AccountRoleInput() *string
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
	InternalValue() *DataSnowflakeGrantsGrantsOf
	SetInternalValue(val *DataSnowflakeGrantsGrantsOf)
	Share() *string
	SetShare(val *string)
	ShareInput() *string
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
	ResetAccountRole()
	ResetApplicationRole()
	ResetDatabaseRole()
	ResetShare()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSnowflakeGrantsGrantsOfOutputReference
type jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) AccountRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) AccountRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ApplicationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ApplicationRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) DatabaseRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) DatabaseRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) InternalValue() *DataSnowflakeGrantsGrantsOf {
	var returns *DataSnowflakeGrantsGrantsOf
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) Share() *string {
	var returns *string
	_jsii_.Get(
		j,
		"share",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ShareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSnowflakeGrantsGrantsOfOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataSnowflakeGrantsGrantsOfOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeGrantsGrantsOfOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrantsGrantsOfOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataSnowflakeGrantsGrantsOfOutputReference_Override(d DataSnowflakeGrantsGrantsOfOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeGrants.DataSnowflakeGrantsGrantsOfOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetAccountRole(val *string) {
	if err := j.validateSetAccountRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountRole",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetApplicationRole(val *string) {
	if err := j.validateSetApplicationRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationRole",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetDatabaseRole(val *string) {
	if err := j.validateSetDatabaseRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseRole",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetInternalValue(val *DataSnowflakeGrantsGrantsOf) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetShare(val *string) {
	if err := j.validateSetShareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"share",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ResetAccountRole() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ResetApplicationRole() {
	_jsii_.InvokeVoid(
		d,
		"resetApplicationRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ResetDatabaseRole() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ResetShare() {
	_jsii_.InvokeVoid(
		d,
		"resetShare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeGrantsGrantsOfOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

