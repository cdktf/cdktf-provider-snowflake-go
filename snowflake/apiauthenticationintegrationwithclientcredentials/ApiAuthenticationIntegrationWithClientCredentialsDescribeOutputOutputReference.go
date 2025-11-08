// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiauthenticationintegrationwithclientcredentials

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/apiauthenticationintegrationwithclientcredentials/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	AuthType() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputAuthTypeList
	Comment() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputCommentList
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
	Enabled() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputEnabledList
	// Experimental.
	Fqn() *string
	InternalValue() *ApiAuthenticationIntegrationWithClientCredentialsDescribeOutput
	SetInternalValue(val *ApiAuthenticationIntegrationWithClientCredentialsDescribeOutput)
	OauthAccessTokenValidity() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAccessTokenValidityList
	OauthAllowedScopes() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAllowedScopesList
	OauthAuthorizationEndpoint() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAuthorizationEndpointList
	OauthClientAuthMethod() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthClientAuthMethodList
	OauthGrant() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthGrantList
	OauthRefreshTokenValidity() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthRefreshTokenValidityList
	OauthTokenEndpoint() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList
	ParentIntegration() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputParentIntegrationList
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

// The jsii proxy struct for ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference
type jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) AuthType() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputAuthTypeList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputAuthTypeList
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) Comment() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputCommentList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputCommentList
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) Enabled() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputEnabledList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputEnabledList
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) InternalValue() *ApiAuthenticationIntegrationWithClientCredentialsDescribeOutput {
	var returns *ApiAuthenticationIntegrationWithClientCredentialsDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) OauthAccessTokenValidity() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAccessTokenValidityList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAccessTokenValidityList
	_jsii_.Get(
		j,
		"oauthAccessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) OauthAllowedScopes() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAllowedScopesList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAllowedScopesList
	_jsii_.Get(
		j,
		"oauthAllowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) OauthAuthorizationEndpoint() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAuthorizationEndpointList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthAuthorizationEndpointList
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) OauthClientAuthMethod() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthClientAuthMethodList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthClientAuthMethodList
	_jsii_.Get(
		j,
		"oauthClientAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) OauthGrant() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthGrantList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthGrantList
	_jsii_.Get(
		j,
		"oauthGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) OauthRefreshTokenValidity() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthRefreshTokenValidityList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthRefreshTokenValidityList
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) OauthTokenEndpoint() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) ParentIntegration() ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputParentIntegrationList {
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputParentIntegrationList
	_jsii_.Get(
		j,
		"parentIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithClientCredentials.ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference_Override(a ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithClientCredentials.ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference)SetInternalValue(val *ApiAuthenticationIntegrationWithClientCredentialsDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

