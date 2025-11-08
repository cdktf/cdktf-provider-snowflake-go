// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeusers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/datasnowflakeusers/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeUsersUsersDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	Comment() *string
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
	CustomLandingPageUrl() *string
	CustomLandingPageUrlFlushNextUiLoad() cdktf.IResolvable
	DaysToExpiry() *float64
	DefaultNamespace() *string
	DefaultRole() *string
	DefaultSecondaryRoles() *string
	DefaultWarehouse() *string
	Disabled() cdktf.IResolvable
	DisplayName() *string
	Email() *string
	ExtAuthnDuo() cdktf.IResolvable
	ExtAuthnUid() *string
	FirstName() *string
	// Experimental.
	Fqn() *string
	HasMfa() cdktf.IResolvable
	InternalValue() *DataSnowflakeUsersUsersDescribeOutput
	SetInternalValue(val *DataSnowflakeUsersUsersDescribeOutput)
	LastName() *string
	LoginName() *string
	MiddleName() *string
	MinsToBypassMfa() *float64
	MinsToBypassNetworkPolicy() *float64
	MinsToUnlock() *float64
	MustChangePassword() cdktf.IResolvable
	Name() *string
	PasswordLastSetTime() *string
	RsaPublicKey() *string
	RsaPublicKey2() *string
	RsaPublicKey2Fp() *string
	RsaPublicKeyFp() *string
	SnowflakeLock() cdktf.IResolvable
	SnowflakeSupport() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSnowflakeUsersUsersDescribeOutputOutputReference
type jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) CustomLandingPageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLandingPageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) CustomLandingPageUrlFlushNextUiLoad() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"customLandingPageUrlFlushNextUiLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) DaysToExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysToExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) DefaultNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) DefaultRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) DefaultSecondaryRoles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) DefaultWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) Disabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) ExtAuthnDuo() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"extAuthnDuo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) ExtAuthnUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extAuthnUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) HasMfa() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) InternalValue() *DataSnowflakeUsersUsersDescribeOutput {
	var returns *DataSnowflakeUsersUsersDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) LoginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) MinsToBypassMfa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToBypassMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) MinsToBypassNetworkPolicy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToBypassNetworkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) MinsToUnlock() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minsToUnlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) MustChangePassword() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mustChangePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) PasswordLastSetTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordLastSetTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) RsaPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) RsaPublicKey2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) RsaPublicKey2Fp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2Fp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) RsaPublicKeyFp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKeyFp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) SnowflakeLock() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"snowflakeLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) SnowflakeSupport() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"snowflakeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDataSnowflakeUsersUsersDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeUsersUsersDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeUsersUsersDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeUsers.DataSnowflakeUsersUsersDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeUsersUsersDescribeOutputOutputReference_Override(d DataSnowflakeUsersUsersDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeUsers.DataSnowflakeUsersUsersDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference)SetInternalValue(val *DataSnowflakeUsersUsersDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeUsersUsersDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

