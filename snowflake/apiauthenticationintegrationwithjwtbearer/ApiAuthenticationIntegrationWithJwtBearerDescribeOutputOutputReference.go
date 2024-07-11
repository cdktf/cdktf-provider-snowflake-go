// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiauthenticationintegrationwithjwtbearer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/apiauthenticationintegrationwithjwtbearer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	AuthType() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputAuthTypeList
	Comment() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputCommentList
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
	Enabled() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputEnabledList
	// Experimental.
	Fqn() *string
	InternalValue() *ApiAuthenticationIntegrationWithJwtBearerDescribeOutput
	SetInternalValue(val *ApiAuthenticationIntegrationWithJwtBearerDescribeOutput)
	OauthAccessTokenValidity() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAccessTokenValidityList
	OauthAllowedScopes() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAllowedScopesList
	OauthAuthorizationEndpoint() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAuthorizationEndpointList
	OauthClientAuthMethod() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthClientAuthMethodList
	OauthClientId() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthClientIdList
	OauthGrant() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthGrantList
	OauthRefreshTokenValidity() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthRefreshTokenValidityList
	OauthTokenEndpoint() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthTokenEndpointList
	ParentIntegration() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputParentIntegrationList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference
type jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) AuthType() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputAuthTypeList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputAuthTypeList
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) Comment() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputCommentList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputCommentList
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) Enabled() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputEnabledList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputEnabledList
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) InternalValue() *ApiAuthenticationIntegrationWithJwtBearerDescribeOutput {
	var returns *ApiAuthenticationIntegrationWithJwtBearerDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) OauthAccessTokenValidity() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAccessTokenValidityList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAccessTokenValidityList
	_jsii_.Get(
		j,
		"oauthAccessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) OauthAllowedScopes() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAllowedScopesList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAllowedScopesList
	_jsii_.Get(
		j,
		"oauthAllowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) OauthAuthorizationEndpoint() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAuthorizationEndpointList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthAuthorizationEndpointList
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) OauthClientAuthMethod() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthClientAuthMethodList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthClientAuthMethodList
	_jsii_.Get(
		j,
		"oauthClientAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) OauthClientId() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthClientIdList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthClientIdList
	_jsii_.Get(
		j,
		"oauthClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) OauthGrant() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthGrantList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthGrantList
	_jsii_.Get(
		j,
		"oauthGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) OauthRefreshTokenValidity() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthRefreshTokenValidityList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthRefreshTokenValidityList
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) OauthTokenEndpoint() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthTokenEndpointList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOauthTokenEndpointList
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) ParentIntegration() ApiAuthenticationIntegrationWithJwtBearerDescribeOutputParentIntegrationList {
	var returns ApiAuthenticationIntegrationWithJwtBearerDescribeOutputParentIntegrationList
	_jsii_.Get(
		j,
		"parentIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithJwtBearer.ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference_Override(a ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithJwtBearer.ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference)SetInternalValue(val *ApiAuthenticationIntegrationWithJwtBearerDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApiAuthenticationIntegrationWithJwtBearerDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

