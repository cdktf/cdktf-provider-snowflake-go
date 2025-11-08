// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/storageintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageIntegrationDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	AzureConsentUrl() StorageIntegrationDescribeOutputAzureConsentUrlList
	AzureMultiTenantAppName() StorageIntegrationDescribeOutputAzureMultiTenantAppNameList
	Comment() StorageIntegrationDescribeOutputCommentList
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
	Enabled() StorageIntegrationDescribeOutputEnabledList
	// Experimental.
	Fqn() *string
	InternalValue() *StorageIntegrationDescribeOutput
	SetInternalValue(val *StorageIntegrationDescribeOutput)
	StorageAllowedLocations() StorageIntegrationDescribeOutputStorageAllowedLocationsList
	StorageAwsExternalId() StorageIntegrationDescribeOutputStorageAwsExternalIdList
	StorageAwsIamUserArn() StorageIntegrationDescribeOutputStorageAwsIamUserArnList
	StorageAwsObjectAcl() StorageIntegrationDescribeOutputStorageAwsObjectAclList
	StorageAwsRoleArn() StorageIntegrationDescribeOutputStorageAwsRoleArnList
	StorageBlockedLocations() StorageIntegrationDescribeOutputStorageBlockedLocationsList
	StorageGcpServiceAccount() StorageIntegrationDescribeOutputStorageGcpServiceAccountList
	StorageProvider() StorageIntegrationDescribeOutputStorageProviderList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsePrivatelinkEndpoint() StorageIntegrationDescribeOutputUsePrivatelinkEndpointList
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

// The jsii proxy struct for StorageIntegrationDescribeOutputOutputReference
type jsiiProxy_StorageIntegrationDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) AzureConsentUrl() StorageIntegrationDescribeOutputAzureConsentUrlList {
	var returns StorageIntegrationDescribeOutputAzureConsentUrlList
	_jsii_.Get(
		j,
		"azureConsentUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) AzureMultiTenantAppName() StorageIntegrationDescribeOutputAzureMultiTenantAppNameList {
	var returns StorageIntegrationDescribeOutputAzureMultiTenantAppNameList
	_jsii_.Get(
		j,
		"azureMultiTenantAppName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) Comment() StorageIntegrationDescribeOutputCommentList {
	var returns StorageIntegrationDescribeOutputCommentList
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) Enabled() StorageIntegrationDescribeOutputEnabledList {
	var returns StorageIntegrationDescribeOutputEnabledList
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) InternalValue() *StorageIntegrationDescribeOutput {
	var returns *StorageIntegrationDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) StorageAllowedLocations() StorageIntegrationDescribeOutputStorageAllowedLocationsList {
	var returns StorageIntegrationDescribeOutputStorageAllowedLocationsList
	_jsii_.Get(
		j,
		"storageAllowedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) StorageAwsExternalId() StorageIntegrationDescribeOutputStorageAwsExternalIdList {
	var returns StorageIntegrationDescribeOutputStorageAwsExternalIdList
	_jsii_.Get(
		j,
		"storageAwsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) StorageAwsIamUserArn() StorageIntegrationDescribeOutputStorageAwsIamUserArnList {
	var returns StorageIntegrationDescribeOutputStorageAwsIamUserArnList
	_jsii_.Get(
		j,
		"storageAwsIamUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) StorageAwsObjectAcl() StorageIntegrationDescribeOutputStorageAwsObjectAclList {
	var returns StorageIntegrationDescribeOutputStorageAwsObjectAclList
	_jsii_.Get(
		j,
		"storageAwsObjectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) StorageAwsRoleArn() StorageIntegrationDescribeOutputStorageAwsRoleArnList {
	var returns StorageIntegrationDescribeOutputStorageAwsRoleArnList
	_jsii_.Get(
		j,
		"storageAwsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) StorageBlockedLocations() StorageIntegrationDescribeOutputStorageBlockedLocationsList {
	var returns StorageIntegrationDescribeOutputStorageBlockedLocationsList
	_jsii_.Get(
		j,
		"storageBlockedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) StorageGcpServiceAccount() StorageIntegrationDescribeOutputStorageGcpServiceAccountList {
	var returns StorageIntegrationDescribeOutputStorageGcpServiceAccountList
	_jsii_.Get(
		j,
		"storageGcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) StorageProvider() StorageIntegrationDescribeOutputStorageProviderList {
	var returns StorageIntegrationDescribeOutputStorageProviderList
	_jsii_.Get(
		j,
		"storageProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) UsePrivatelinkEndpoint() StorageIntegrationDescribeOutputUsePrivatelinkEndpointList {
	var returns StorageIntegrationDescribeOutputUsePrivatelinkEndpointList
	_jsii_.Get(
		j,
		"usePrivatelinkEndpoint",
		&returns,
	)
	return returns
}


func NewStorageIntegrationDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StorageIntegrationDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewStorageIntegrationDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageIntegrationDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.storageIntegration.StorageIntegrationDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStorageIntegrationDescribeOutputOutputReference_Override(s StorageIntegrationDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.storageIntegration.StorageIntegrationDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference)SetInternalValue(val *StorageIntegrationDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

