// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeauthenticationpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/datasnowflakeauthenticationpolicies/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList interface {
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
	Get(index *float64) DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList
type jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList {
	_init_.Initialize()

	if err := validateNewDataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeAuthenticationPolicies.DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList_Override(d DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeAuthenticationPolicies.DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
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

func (d *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) Get(index *float64) DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeAuthenticationPoliciesAuthenticationPoliciesDescribeOutputList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

