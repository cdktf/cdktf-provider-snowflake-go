// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/externalvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExternalVolumeStorageLocationOutputReference interface {
	cdktf.ComplexObject
	AzureTenantId() *string
	SetAzureTenantId(val *string)
	AzureTenantIdInput() *string
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
	EncryptionKmsKeyId() *string
	SetEncryptionKmsKeyId(val *string)
	EncryptionKmsKeyIdInput() *string
	EncryptionType() *string
	SetEncryptionType(val *string)
	EncryptionTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageAwsExternalId() *string
	StorageAwsRoleArn() *string
	SetStorageAwsRoleArn(val *string)
	StorageAwsRoleArnInput() *string
	StorageBaseUrl() *string
	SetStorageBaseUrl(val *string)
	StorageBaseUrlInput() *string
	StorageLocationName() *string
	SetStorageLocationName(val *string)
	StorageLocationNameInput() *string
	StorageProvider() *string
	SetStorageProvider(val *string)
	StorageProviderInput() *string
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
	ResetAzureTenantId()
	ResetEncryptionKmsKeyId()
	ResetEncryptionType()
	ResetStorageAwsRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ExternalVolumeStorageLocationOutputReference
type jsiiProxy_ExternalVolumeStorageLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) AzureTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) EncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) EncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) EncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageAwsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageAwsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageAwsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageBaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBaseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageBaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBaseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageLocationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageLocationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) StorageProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExternalVolumeStorageLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ExternalVolumeStorageLocationOutputReference {
	_init_.Initialize()

	if err := validateNewExternalVolumeStorageLocationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalVolumeStorageLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalVolume.ExternalVolumeStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewExternalVolumeStorageLocationOutputReference_Override(e ExternalVolumeStorageLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalVolume.ExternalVolumeStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetAzureTenantId(val *string) {
	if err := j.validateSetAzureTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureTenantId",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetEncryptionKmsKeyId(val *string) {
	if err := j.validateSetEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetEncryptionType(val *string) {
	if err := j.validateSetEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionType",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetStorageAwsRoleArn(val *string) {
	if err := j.validateSetStorageAwsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAwsRoleArn",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetStorageBaseUrl(val *string) {
	if err := j.validateSetStorageBaseUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageBaseUrl",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetStorageLocationName(val *string) {
	if err := j.validateSetStorageLocationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageLocationName",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetStorageProvider(val *string) {
	if err := j.validateSetStorageProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProvider",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExternalVolumeStorageLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) ResetAzureTenantId() {
	_jsii_.InvokeVoid(
		e,
		"resetAzureTenantId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) ResetEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) ResetEncryptionType() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) ResetStorageAwsRoleArn() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageAwsRoleArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_ExternalVolumeStorageLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

