// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/authenticationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthenticationPolicyWorkloadIdentityPolicyOutputReference interface {
	cdktf.ComplexObject
	AllowedAwsAccounts() *[]*string
	SetAllowedAwsAccounts(val *[]*string)
	AllowedAwsAccountsInput() *[]*string
	AllowedAzureIssuers() *[]*string
	SetAllowedAzureIssuers(val *[]*string)
	AllowedAzureIssuersInput() *[]*string
	AllowedOidcIssuers() *[]*string
	SetAllowedOidcIssuers(val *[]*string)
	AllowedOidcIssuersInput() *[]*string
	AllowedProviders() *[]*string
	SetAllowedProviders(val *[]*string)
	AllowedProvidersInput() *[]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() *AuthenticationPolicyWorkloadIdentityPolicy
	SetInternalValue(val *AuthenticationPolicyWorkloadIdentityPolicy)
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
	ResetAllowedAwsAccounts()
	ResetAllowedAzureIssuers()
	ResetAllowedOidcIssuers()
	ResetAllowedProviders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AuthenticationPolicyWorkloadIdentityPolicyOutputReference
type jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) AllowedAwsAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAwsAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) AllowedAwsAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAwsAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) AllowedAzureIssuers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAzureIssuers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) AllowedAzureIssuersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAzureIssuersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) AllowedOidcIssuers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOidcIssuers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) AllowedOidcIssuersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOidcIssuersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) AllowedProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) AllowedProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) InternalValue() *AuthenticationPolicyWorkloadIdentityPolicy {
	var returns *AuthenticationPolicyWorkloadIdentityPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAuthenticationPolicyWorkloadIdentityPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AuthenticationPolicyWorkloadIdentityPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAuthenticationPolicyWorkloadIdentityPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.authenticationPolicy.AuthenticationPolicyWorkloadIdentityPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAuthenticationPolicyWorkloadIdentityPolicyOutputReference_Override(a AuthenticationPolicyWorkloadIdentityPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.authenticationPolicy.AuthenticationPolicyWorkloadIdentityPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetAllowedAwsAccounts(val *[]*string) {
	if err := j.validateSetAllowedAwsAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAwsAccounts",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetAllowedAzureIssuers(val *[]*string) {
	if err := j.validateSetAllowedAzureIssuersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAzureIssuers",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetAllowedOidcIssuers(val *[]*string) {
	if err := j.validateSetAllowedOidcIssuersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOidcIssuers",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetAllowedProviders(val *[]*string) {
	if err := j.validateSetAllowedProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedProviders",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetInternalValue(val *AuthenticationPolicyWorkloadIdentityPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) ResetAllowedAwsAccounts() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedAwsAccounts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) ResetAllowedAzureIssuers() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedAzureIssuers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) ResetAllowedOidcIssuers() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedOidcIssuers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) ResetAllowedProviders() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedProviders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AuthenticationPolicyWorkloadIdentityPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

