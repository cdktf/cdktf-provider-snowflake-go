// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakesecurityintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/datasnowflakesecurityintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2OutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List
type jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List {
	_init_.Initialize()

	if err := validateNewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2ListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeSecurityIntegrations.DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List_Override(d DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeSecurityIntegrations.DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) Get(index *float64) DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2OutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2OutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
