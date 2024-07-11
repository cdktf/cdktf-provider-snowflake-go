// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package saml2integration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/saml2integration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Saml2IntegrationDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	AllowedEmailPatterns() Saml2IntegrationDescribeOutputAllowedEmailPatternsList
	AllowedUserDomains() Saml2IntegrationDescribeOutputAllowedUserDomainsList
	Comment() Saml2IntegrationDescribeOutputCommentList
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
	InternalValue() *Saml2IntegrationDescribeOutput
	SetInternalValue(val *Saml2IntegrationDescribeOutput)
	Saml2DigestMethodsUsed() Saml2IntegrationDescribeOutputSaml2DigestMethodsUsedList
	Saml2EnableSpInitiated() Saml2IntegrationDescribeOutputSaml2EnableSpInitiatedList
	Saml2ForceAuthn() Saml2IntegrationDescribeOutputSaml2ForceAuthnList
	Saml2Issuer() Saml2IntegrationDescribeOutputSaml2IssuerList
	Saml2PostLogoutRedirectUrl() Saml2IntegrationDescribeOutputSaml2PostLogoutRedirectUrlList
	Saml2Provider() Saml2IntegrationDescribeOutputSaml2ProviderList
	Saml2RequestedNameidFormat() Saml2IntegrationDescribeOutputSaml2RequestedNameidFormatList
	Saml2SignatureMethodsUsed() Saml2IntegrationDescribeOutputSaml2SignatureMethodsUsedList
	Saml2SignRequest() Saml2IntegrationDescribeOutputSaml2SignRequestList
	Saml2SnowflakeAcsUrl() Saml2IntegrationDescribeOutputSaml2SnowflakeAcsUrlList
	Saml2SnowflakeIssuerUrl() Saml2IntegrationDescribeOutputSaml2SnowflakeIssuerUrlList
	Saml2SnowflakeMetadata() Saml2IntegrationDescribeOutputSaml2SnowflakeMetadataList
	Saml2SnowflakeX509Cert() Saml2IntegrationDescribeOutputSaml2SnowflakeX509CertList
	Saml2SpInitiatedLoginPageLabel() Saml2IntegrationDescribeOutputSaml2SpInitiatedLoginPageLabelList
	Saml2SsoUrl() Saml2IntegrationDescribeOutputSaml2SsoUrlList
	Saml2X509Cert() Saml2IntegrationDescribeOutputSaml2X509CertList
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

// The jsii proxy struct for Saml2IntegrationDescribeOutputOutputReference
type jsiiProxy_Saml2IntegrationDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) AllowedEmailPatterns() Saml2IntegrationDescribeOutputAllowedEmailPatternsList {
	var returns Saml2IntegrationDescribeOutputAllowedEmailPatternsList
	_jsii_.Get(
		j,
		"allowedEmailPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) AllowedUserDomains() Saml2IntegrationDescribeOutputAllowedUserDomainsList {
	var returns Saml2IntegrationDescribeOutputAllowedUserDomainsList
	_jsii_.Get(
		j,
		"allowedUserDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Comment() Saml2IntegrationDescribeOutputCommentList {
	var returns Saml2IntegrationDescribeOutputCommentList
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) InternalValue() *Saml2IntegrationDescribeOutput {
	var returns *Saml2IntegrationDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2DigestMethodsUsed() Saml2IntegrationDescribeOutputSaml2DigestMethodsUsedList {
	var returns Saml2IntegrationDescribeOutputSaml2DigestMethodsUsedList
	_jsii_.Get(
		j,
		"saml2DigestMethodsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2EnableSpInitiated() Saml2IntegrationDescribeOutputSaml2EnableSpInitiatedList {
	var returns Saml2IntegrationDescribeOutputSaml2EnableSpInitiatedList
	_jsii_.Get(
		j,
		"saml2EnableSpInitiated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2ForceAuthn() Saml2IntegrationDescribeOutputSaml2ForceAuthnList {
	var returns Saml2IntegrationDescribeOutputSaml2ForceAuthnList
	_jsii_.Get(
		j,
		"saml2ForceAuthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2Issuer() Saml2IntegrationDescribeOutputSaml2IssuerList {
	var returns Saml2IntegrationDescribeOutputSaml2IssuerList
	_jsii_.Get(
		j,
		"saml2Issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2PostLogoutRedirectUrl() Saml2IntegrationDescribeOutputSaml2PostLogoutRedirectUrlList {
	var returns Saml2IntegrationDescribeOutputSaml2PostLogoutRedirectUrlList
	_jsii_.Get(
		j,
		"saml2PostLogoutRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2Provider() Saml2IntegrationDescribeOutputSaml2ProviderList {
	var returns Saml2IntegrationDescribeOutputSaml2ProviderList
	_jsii_.Get(
		j,
		"saml2Provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2RequestedNameidFormat() Saml2IntegrationDescribeOutputSaml2RequestedNameidFormatList {
	var returns Saml2IntegrationDescribeOutputSaml2RequestedNameidFormatList
	_jsii_.Get(
		j,
		"saml2RequestedNameidFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2SignatureMethodsUsed() Saml2IntegrationDescribeOutputSaml2SignatureMethodsUsedList {
	var returns Saml2IntegrationDescribeOutputSaml2SignatureMethodsUsedList
	_jsii_.Get(
		j,
		"saml2SignatureMethodsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2SignRequest() Saml2IntegrationDescribeOutputSaml2SignRequestList {
	var returns Saml2IntegrationDescribeOutputSaml2SignRequestList
	_jsii_.Get(
		j,
		"saml2SignRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2SnowflakeAcsUrl() Saml2IntegrationDescribeOutputSaml2SnowflakeAcsUrlList {
	var returns Saml2IntegrationDescribeOutputSaml2SnowflakeAcsUrlList
	_jsii_.Get(
		j,
		"saml2SnowflakeAcsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2SnowflakeIssuerUrl() Saml2IntegrationDescribeOutputSaml2SnowflakeIssuerUrlList {
	var returns Saml2IntegrationDescribeOutputSaml2SnowflakeIssuerUrlList
	_jsii_.Get(
		j,
		"saml2SnowflakeIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2SnowflakeMetadata() Saml2IntegrationDescribeOutputSaml2SnowflakeMetadataList {
	var returns Saml2IntegrationDescribeOutputSaml2SnowflakeMetadataList
	_jsii_.Get(
		j,
		"saml2SnowflakeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2SnowflakeX509Cert() Saml2IntegrationDescribeOutputSaml2SnowflakeX509CertList {
	var returns Saml2IntegrationDescribeOutputSaml2SnowflakeX509CertList
	_jsii_.Get(
		j,
		"saml2SnowflakeX509Cert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2SpInitiatedLoginPageLabel() Saml2IntegrationDescribeOutputSaml2SpInitiatedLoginPageLabelList {
	var returns Saml2IntegrationDescribeOutputSaml2SpInitiatedLoginPageLabelList
	_jsii_.Get(
		j,
		"saml2SpInitiatedLoginPageLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2SsoUrl() Saml2IntegrationDescribeOutputSaml2SsoUrlList {
	var returns Saml2IntegrationDescribeOutputSaml2SsoUrlList
	_jsii_.Get(
		j,
		"saml2SsoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Saml2X509Cert() Saml2IntegrationDescribeOutputSaml2X509CertList {
	var returns Saml2IntegrationDescribeOutputSaml2X509CertList
	_jsii_.Get(
		j,
		"saml2X509Cert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSaml2IntegrationDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Saml2IntegrationDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewSaml2IntegrationDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Saml2IntegrationDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.saml2Integration.Saml2IntegrationDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSaml2IntegrationDescribeOutputOutputReference_Override(s Saml2IntegrationDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.saml2Integration.Saml2IntegrationDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference)SetInternalValue(val *Saml2IntegrationDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

