// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/authenticationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthenticationPolicyPatPolicyOutputReference interface {
	cdktf.ComplexObject
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
	DefaultExpiryInDays() *float64
	SetDefaultExpiryInDays(val *float64)
	DefaultExpiryInDaysInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *AuthenticationPolicyPatPolicy
	SetInternalValue(val *AuthenticationPolicyPatPolicy)
	MaxExpiryInDays() *float64
	SetMaxExpiryInDays(val *float64)
	MaxExpiryInDaysInput() *float64
	NetworkPolicyEvaluation() *string
	SetNetworkPolicyEvaluation(val *string)
	NetworkPolicyEvaluationInput() *string
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
	ResetDefaultExpiryInDays()
	ResetMaxExpiryInDays()
	ResetNetworkPolicyEvaluation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AuthenticationPolicyPatPolicyOutputReference
type jsiiProxy_AuthenticationPolicyPatPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) DefaultExpiryInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultExpiryInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) DefaultExpiryInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultExpiryInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) InternalValue() *AuthenticationPolicyPatPolicy {
	var returns *AuthenticationPolicyPatPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) MaxExpiryInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExpiryInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) MaxExpiryInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExpiryInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) NetworkPolicyEvaluation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) NetworkPolicyEvaluationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAuthenticationPolicyPatPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AuthenticationPolicyPatPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAuthenticationPolicyPatPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthenticationPolicyPatPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.authenticationPolicy.AuthenticationPolicyPatPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAuthenticationPolicyPatPolicyOutputReference_Override(a AuthenticationPolicyPatPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.authenticationPolicy.AuthenticationPolicyPatPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference)SetDefaultExpiryInDays(val *float64) {
	if err := j.validateSetDefaultExpiryInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultExpiryInDays",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference)SetInternalValue(val *AuthenticationPolicyPatPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference)SetMaxExpiryInDays(val *float64) {
	if err := j.validateSetMaxExpiryInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxExpiryInDays",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference)SetNetworkPolicyEvaluation(val *string) {
	if err := j.validateSetNetworkPolicyEvaluationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPolicyEvaluation",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) ResetDefaultExpiryInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultExpiryInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) ResetMaxExpiryInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxExpiryInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) ResetNetworkPolicyEvaluation() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkPolicyEvaluation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyPatPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

