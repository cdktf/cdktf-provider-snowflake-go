// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakesecurityintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/datasnowflakesecurityintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference interface {
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
	Default() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabel
	SetInternalValue(val *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabel)
	Name() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	Value() *string
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

// The jsii proxy struct for DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference
type jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) Default() *string {
	var returns *string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) InternalValue() *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabel {
	var returns *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeSecurityIntegrations.DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference_Override(d DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeSecurityIntegrations.DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference)SetInternalValue(val *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

