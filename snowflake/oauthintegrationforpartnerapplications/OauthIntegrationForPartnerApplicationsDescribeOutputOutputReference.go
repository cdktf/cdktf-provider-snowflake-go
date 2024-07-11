// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforpartnerapplications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/oauthintegrationforpartnerapplications/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	BlockedRolesList() OauthIntegrationForPartnerApplicationsDescribeOutputBlockedRolesListStructList
	Comment() OauthIntegrationForPartnerApplicationsDescribeOutputCommentList
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
	Enabled() OauthIntegrationForPartnerApplicationsDescribeOutputEnabledList
	// Experimental.
	Fqn() *string
	InternalValue() *OauthIntegrationForPartnerApplicationsDescribeOutput
	SetInternalValue(val *OauthIntegrationForPartnerApplicationsDescribeOutput)
	NetworkPolicy() OauthIntegrationForPartnerApplicationsDescribeOutputNetworkPolicyList
	OauthAllowedAuthorizationEndpoints() OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedAuthorizationEndpointsList
	OauthAllowedTokenEndpoints() OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList
	OauthAllowNonTlsRedirectUri() OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowNonTlsRedirectUriList
	OauthAuthorizationEndpoint() OauthIntegrationForPartnerApplicationsDescribeOutputOauthAuthorizationEndpointList
	OauthClientId() OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientIdList
	OauthClientRsaPublicKey2Fp() OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientRsaPublicKey2FpList
	OauthClientRsaPublicKeyFp() OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientRsaPublicKeyFpList
	OauthClientType() OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList
	OauthEnforcePkce() OauthIntegrationForPartnerApplicationsDescribeOutputOauthEnforcePkceList
	OauthIssueRefreshTokens() OauthIntegrationForPartnerApplicationsDescribeOutputOauthIssueRefreshTokensList
	OauthRedirectUri() OauthIntegrationForPartnerApplicationsDescribeOutputOauthRedirectUriList
	OauthRefreshTokenValidity() OauthIntegrationForPartnerApplicationsDescribeOutputOauthRefreshTokenValidityList
	OauthTokenEndpoint() OauthIntegrationForPartnerApplicationsDescribeOutputOauthTokenEndpointList
	OauthUseSecondaryRoles() OauthIntegrationForPartnerApplicationsDescribeOutputOauthUseSecondaryRolesList
	PreAuthorizedRolesList() OauthIntegrationForPartnerApplicationsDescribeOutputPreAuthorizedRolesListStructList
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

// The jsii proxy struct for OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference
type jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) BlockedRolesList() OauthIntegrationForPartnerApplicationsDescribeOutputBlockedRolesListStructList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputBlockedRolesListStructList
	_jsii_.Get(
		j,
		"blockedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) Comment() OauthIntegrationForPartnerApplicationsDescribeOutputCommentList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputCommentList
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) Enabled() OauthIntegrationForPartnerApplicationsDescribeOutputEnabledList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputEnabledList
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) InternalValue() *OauthIntegrationForPartnerApplicationsDescribeOutput {
	var returns *OauthIntegrationForPartnerApplicationsDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) NetworkPolicy() OauthIntegrationForPartnerApplicationsDescribeOutputNetworkPolicyList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthAllowedAuthorizationEndpoints() OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedAuthorizationEndpointsList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedAuthorizationEndpointsList
	_jsii_.Get(
		j,
		"oauthAllowedAuthorizationEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthAllowedTokenEndpoints() OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList
	_jsii_.Get(
		j,
		"oauthAllowedTokenEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthAllowNonTlsRedirectUri() OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowNonTlsRedirectUriList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowNonTlsRedirectUriList
	_jsii_.Get(
		j,
		"oauthAllowNonTlsRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthAuthorizationEndpoint() OauthIntegrationForPartnerApplicationsDescribeOutputOauthAuthorizationEndpointList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthAuthorizationEndpointList
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthClientId() OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientIdList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientIdList
	_jsii_.Get(
		j,
		"oauthClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthClientRsaPublicKey2Fp() OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientRsaPublicKey2FpList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientRsaPublicKey2FpList
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKey2Fp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthClientRsaPublicKeyFp() OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientRsaPublicKeyFpList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientRsaPublicKeyFpList
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKeyFp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthClientType() OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList
	_jsii_.Get(
		j,
		"oauthClientType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthEnforcePkce() OauthIntegrationForPartnerApplicationsDescribeOutputOauthEnforcePkceList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthEnforcePkceList
	_jsii_.Get(
		j,
		"oauthEnforcePkce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthIssueRefreshTokens() OauthIntegrationForPartnerApplicationsDescribeOutputOauthIssueRefreshTokensList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthIssueRefreshTokensList
	_jsii_.Get(
		j,
		"oauthIssueRefreshTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthRedirectUri() OauthIntegrationForPartnerApplicationsDescribeOutputOauthRedirectUriList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthRedirectUriList
	_jsii_.Get(
		j,
		"oauthRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthRefreshTokenValidity() OauthIntegrationForPartnerApplicationsDescribeOutputOauthRefreshTokenValidityList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthRefreshTokenValidityList
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthTokenEndpoint() OauthIntegrationForPartnerApplicationsDescribeOutputOauthTokenEndpointList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthTokenEndpointList
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) OauthUseSecondaryRoles() OauthIntegrationForPartnerApplicationsDescribeOutputOauthUseSecondaryRolesList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthUseSecondaryRolesList
	_jsii_.Get(
		j,
		"oauthUseSecondaryRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) PreAuthorizedRolesList() OauthIntegrationForPartnerApplicationsDescribeOutputPreAuthorizedRolesListStructList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputPreAuthorizedRolesListStructList
	_jsii_.Get(
		j,
		"preAuthorizedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOauthIntegrationForPartnerApplicationsDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewOauthIntegrationForPartnerApplicationsDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOauthIntegrationForPartnerApplicationsDescribeOutputOutputReference_Override(o OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference)SetInternalValue(val *OauthIntegrationForPartnerApplicationsDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

