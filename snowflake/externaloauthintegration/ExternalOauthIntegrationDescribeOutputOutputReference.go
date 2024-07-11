// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externaloauthintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/externaloauthintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExternalOauthIntegrationDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	Comment() ExternalOauthIntegrationDescribeOutputCommentList
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
	Enabled() ExternalOauthIntegrationDescribeOutputEnabledList
	ExternalOauthAllowedRolesList() ExternalOauthIntegrationDescribeOutputExternalOauthAllowedRolesListStructList
	ExternalOauthAnyRoleMode() ExternalOauthIntegrationDescribeOutputExternalOauthAnyRoleModeList
	ExternalOauthAudienceList() ExternalOauthIntegrationDescribeOutputExternalOauthAudienceListStructList
	ExternalOauthBlockedRolesList() ExternalOauthIntegrationDescribeOutputExternalOauthBlockedRolesListStructList
	ExternalOauthIssuer() ExternalOauthIntegrationDescribeOutputExternalOauthIssuerList
	ExternalOauthJwsKeysUrl() ExternalOauthIntegrationDescribeOutputExternalOauthJwsKeysUrlList
	ExternalOauthRsaPublicKey() ExternalOauthIntegrationDescribeOutputExternalOauthRsaPublicKeyList
	ExternalOauthRsaPublicKey2() ExternalOauthIntegrationDescribeOutputExternalOauthRsaPublicKey2List
	ExternalOauthScopeDelimiter() ExternalOauthIntegrationDescribeOutputExternalOauthScopeDelimiterList
	ExternalOauthSnowflakeUserMappingAttribute() ExternalOauthIntegrationDescribeOutputExternalOauthSnowflakeUserMappingAttributeList
	ExternalOauthTokenUserMappingClaim() ExternalOauthIntegrationDescribeOutputExternalOauthTokenUserMappingClaimList
	// Experimental.
	Fqn() *string
	InternalValue() *ExternalOauthIntegrationDescribeOutput
	SetInternalValue(val *ExternalOauthIntegrationDescribeOutput)
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

// The jsii proxy struct for ExternalOauthIntegrationDescribeOutputOutputReference
type jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) Comment() ExternalOauthIntegrationDescribeOutputCommentList {
	var returns ExternalOauthIntegrationDescribeOutputCommentList
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) Enabled() ExternalOauthIntegrationDescribeOutputEnabledList {
	var returns ExternalOauthIntegrationDescribeOutputEnabledList
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthAllowedRolesList() ExternalOauthIntegrationDescribeOutputExternalOauthAllowedRolesListStructList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthAllowedRolesListStructList
	_jsii_.Get(
		j,
		"externalOauthAllowedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthAnyRoleMode() ExternalOauthIntegrationDescribeOutputExternalOauthAnyRoleModeList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthAnyRoleModeList
	_jsii_.Get(
		j,
		"externalOauthAnyRoleMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthAudienceList() ExternalOauthIntegrationDescribeOutputExternalOauthAudienceListStructList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthAudienceListStructList
	_jsii_.Get(
		j,
		"externalOauthAudienceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthBlockedRolesList() ExternalOauthIntegrationDescribeOutputExternalOauthBlockedRolesListStructList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthBlockedRolesListStructList
	_jsii_.Get(
		j,
		"externalOauthBlockedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthIssuer() ExternalOauthIntegrationDescribeOutputExternalOauthIssuerList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthIssuerList
	_jsii_.Get(
		j,
		"externalOauthIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthJwsKeysUrl() ExternalOauthIntegrationDescribeOutputExternalOauthJwsKeysUrlList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthJwsKeysUrlList
	_jsii_.Get(
		j,
		"externalOauthJwsKeysUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthRsaPublicKey() ExternalOauthIntegrationDescribeOutputExternalOauthRsaPublicKeyList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthRsaPublicKeyList
	_jsii_.Get(
		j,
		"externalOauthRsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthRsaPublicKey2() ExternalOauthIntegrationDescribeOutputExternalOauthRsaPublicKey2List {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthRsaPublicKey2List
	_jsii_.Get(
		j,
		"externalOauthRsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthScopeDelimiter() ExternalOauthIntegrationDescribeOutputExternalOauthScopeDelimiterList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthScopeDelimiterList
	_jsii_.Get(
		j,
		"externalOauthScopeDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthSnowflakeUserMappingAttribute() ExternalOauthIntegrationDescribeOutputExternalOauthSnowflakeUserMappingAttributeList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthSnowflakeUserMappingAttributeList
	_jsii_.Get(
		j,
		"externalOauthSnowflakeUserMappingAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ExternalOauthTokenUserMappingClaim() ExternalOauthIntegrationDescribeOutputExternalOauthTokenUserMappingClaimList {
	var returns ExternalOauthIntegrationDescribeOutputExternalOauthTokenUserMappingClaimList
	_jsii_.Get(
		j,
		"externalOauthTokenUserMappingClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) InternalValue() *ExternalOauthIntegrationDescribeOutput {
	var returns *ExternalOauthIntegrationDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExternalOauthIntegrationDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ExternalOauthIntegrationDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewExternalOauthIntegrationDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegrationDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewExternalOauthIntegrationDescribeOutputOutputReference_Override(e ExternalOauthIntegrationDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegrationDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference)SetInternalValue(val *ExternalOauthIntegrationDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

