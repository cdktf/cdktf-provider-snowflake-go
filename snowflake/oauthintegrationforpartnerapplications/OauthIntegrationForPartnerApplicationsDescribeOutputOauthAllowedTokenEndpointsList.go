// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforpartnerapplications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/oauthintegrationforpartnerapplications/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList interface {
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
	Get(index *float64) OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList
type jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList {
	_init_.Initialize()

	if err := validateNewOauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList_Override(o OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) Get(index *float64) OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthAllowedTokenEndpointsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

