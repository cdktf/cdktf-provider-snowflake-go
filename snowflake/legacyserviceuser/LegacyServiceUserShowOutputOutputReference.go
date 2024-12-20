// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package legacyserviceuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/legacyserviceuser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LegacyServiceUserShowOutputOutputReference interface {
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
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DaysToExpiry() *string
	DefaultNamespace() *string
	DefaultRole() *string
	DefaultSecondaryRoles() *string
	DefaultWarehouse() *string
	Disabled() cdktf.IResolvable
	DisplayName() *string
	Email() *string
	ExpiresAtTime() *string
	ExtAuthnDuo() cdktf.IResolvable
	ExtAuthnUid() *string
	FirstName() *string
	// Experimental.
	Fqn() *string
	HasMfa() cdktf.IResolvable
	HasPassword() cdktf.IResolvable
	HasRsaPublicKey() cdktf.IResolvable
	InternalValue() *LegacyServiceUserShowOutput
	SetInternalValue(val *LegacyServiceUserShowOutput)
	LastName() *string
	LastSuccessLogin() *string
	LockedUntilTime() *string
	LoginName() *string
	MinsToBypassMfa() *string
	MinsToUnlock() *string
	MustChangePassword() cdktf.IResolvable
	Name() *string
	Owner() *string
	SnowflakeLock() cdktf.IResolvable
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LegacyServiceUserShowOutputOutputReference
type jsiiProxy_LegacyServiceUserShowOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) DaysToExpiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"daysToExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) DefaultNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) DefaultRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) DefaultSecondaryRoles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecondaryRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) DefaultWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) Disabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) ExpiresAtTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresAtTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) ExtAuthnDuo() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"extAuthnDuo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) ExtAuthnUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extAuthnUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) HasMfa() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) HasPassword() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) HasRsaPublicKey() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasRsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) InternalValue() *LegacyServiceUserShowOutput {
	var returns *LegacyServiceUserShowOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) LastSuccessLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSuccessLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) LockedUntilTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockedUntilTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) LoginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) MinsToBypassMfa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minsToBypassMfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) MinsToUnlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minsToUnlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) MustChangePassword() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mustChangePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) SnowflakeLock() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"snowflakeLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewLegacyServiceUserShowOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LegacyServiceUserShowOutputOutputReference {
	_init_.Initialize()

	if err := validateNewLegacyServiceUserShowOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LegacyServiceUserShowOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUserShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLegacyServiceUserShowOutputOutputReference_Override(l LegacyServiceUserShowOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.legacyServiceUser.LegacyServiceUserShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference)SetInternalValue(val *LegacyServiceUserShowOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserShowOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserShowOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

