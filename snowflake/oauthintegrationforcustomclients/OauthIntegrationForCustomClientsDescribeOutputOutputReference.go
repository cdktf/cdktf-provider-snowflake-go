// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforcustomclients

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/oauthintegrationforcustomclients/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OauthIntegrationForCustomClientsDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	BlockedRolesList() OauthIntegrationForCustomClientsDescribeOutputBlockedRolesListStructList
	Comment() OauthIntegrationForCustomClientsDescribeOutputCommentList
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
	Enabled() OauthIntegrationForCustomClientsDescribeOutputEnabledList
	// Experimental.
	Fqn() *string
	InternalValue() *OauthIntegrationForCustomClientsDescribeOutput
	SetInternalValue(val *OauthIntegrationForCustomClientsDescribeOutput)
	NetworkPolicy() OauthIntegrationForCustomClientsDescribeOutputNetworkPolicyList
	OauthAllowedAuthorizationEndpoints() OauthIntegrationForCustomClientsDescribeOutputOauthAllowedAuthorizationEndpointsList
	OauthAllowedTokenEndpoints() OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList
	OauthAllowNonTlsRedirectUri() OauthIntegrationForCustomClientsDescribeOutputOauthAllowNonTlsRedirectUriList
	OauthAuthorizationEndpoint() OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList
	OauthClientRsaPublicKey2Fp() OauthIntegrationForCustomClientsDescribeOutputOauthClientRsaPublicKey2FpList
	OauthClientRsaPublicKeyFp() OauthIntegrationForCustomClientsDescribeOutputOauthClientRsaPublicKeyFpList
	OauthClientType() OauthIntegrationForCustomClientsDescribeOutputOauthClientTypeList
	OauthEnforcePkce() OauthIntegrationForCustomClientsDescribeOutputOauthEnforcePkceList
	OauthIssueRefreshTokens() OauthIntegrationForCustomClientsDescribeOutputOauthIssueRefreshTokensList
	OauthRefreshTokenValidity() OauthIntegrationForCustomClientsDescribeOutputOauthRefreshTokenValidityList
	OauthTokenEndpoint() OauthIntegrationForCustomClientsDescribeOutputOauthTokenEndpointList
	OauthUseSecondaryRoles() OauthIntegrationForCustomClientsDescribeOutputOauthUseSecondaryRolesList
	PreAuthorizedRolesList() OauthIntegrationForCustomClientsDescribeOutputPreAuthorizedRolesListStructList
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

// The jsii proxy struct for OauthIntegrationForCustomClientsDescribeOutputOutputReference
type jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) BlockedRolesList() OauthIntegrationForCustomClientsDescribeOutputBlockedRolesListStructList {
	var returns OauthIntegrationForCustomClientsDescribeOutputBlockedRolesListStructList
	_jsii_.Get(
		j,
		"blockedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) Comment() OauthIntegrationForCustomClientsDescribeOutputCommentList {
	var returns OauthIntegrationForCustomClientsDescribeOutputCommentList
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) Enabled() OauthIntegrationForCustomClientsDescribeOutputEnabledList {
	var returns OauthIntegrationForCustomClientsDescribeOutputEnabledList
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) InternalValue() *OauthIntegrationForCustomClientsDescribeOutput {
	var returns *OauthIntegrationForCustomClientsDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) NetworkPolicy() OauthIntegrationForCustomClientsDescribeOutputNetworkPolicyList {
	var returns OauthIntegrationForCustomClientsDescribeOutputNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthAllowedAuthorizationEndpoints() OauthIntegrationForCustomClientsDescribeOutputOauthAllowedAuthorizationEndpointsList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthAllowedAuthorizationEndpointsList
	_jsii_.Get(
		j,
		"oauthAllowedAuthorizationEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthAllowedTokenEndpoints() OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList
	_jsii_.Get(
		j,
		"oauthAllowedTokenEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthAllowNonTlsRedirectUri() OauthIntegrationForCustomClientsDescribeOutputOauthAllowNonTlsRedirectUriList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthAllowNonTlsRedirectUriList
	_jsii_.Get(
		j,
		"oauthAllowNonTlsRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthAuthorizationEndpoint() OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthClientRsaPublicKey2Fp() OauthIntegrationForCustomClientsDescribeOutputOauthClientRsaPublicKey2FpList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthClientRsaPublicKey2FpList
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKey2Fp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthClientRsaPublicKeyFp() OauthIntegrationForCustomClientsDescribeOutputOauthClientRsaPublicKeyFpList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthClientRsaPublicKeyFpList
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKeyFp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthClientType() OauthIntegrationForCustomClientsDescribeOutputOauthClientTypeList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthClientTypeList
	_jsii_.Get(
		j,
		"oauthClientType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthEnforcePkce() OauthIntegrationForCustomClientsDescribeOutputOauthEnforcePkceList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthEnforcePkceList
	_jsii_.Get(
		j,
		"oauthEnforcePkce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthIssueRefreshTokens() OauthIntegrationForCustomClientsDescribeOutputOauthIssueRefreshTokensList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthIssueRefreshTokensList
	_jsii_.Get(
		j,
		"oauthIssueRefreshTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthRefreshTokenValidity() OauthIntegrationForCustomClientsDescribeOutputOauthRefreshTokenValidityList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthRefreshTokenValidityList
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthTokenEndpoint() OauthIntegrationForCustomClientsDescribeOutputOauthTokenEndpointList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthTokenEndpointList
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) OauthUseSecondaryRoles() OauthIntegrationForCustomClientsDescribeOutputOauthUseSecondaryRolesList {
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthUseSecondaryRolesList
	_jsii_.Get(
		j,
		"oauthUseSecondaryRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) PreAuthorizedRolesList() OauthIntegrationForCustomClientsDescribeOutputPreAuthorizedRolesListStructList {
	var returns OauthIntegrationForCustomClientsDescribeOutputPreAuthorizedRolesListStructList
	_jsii_.Get(
		j,
		"preAuthorizedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOauthIntegrationForCustomClientsDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OauthIntegrationForCustomClientsDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewOauthIntegrationForCustomClientsDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClientsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOauthIntegrationForCustomClientsDescribeOutputOutputReference_Override(o OauthIntegrationForCustomClientsDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClientsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference)SetInternalValue(val *OauthIntegrationForCustomClientsDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

