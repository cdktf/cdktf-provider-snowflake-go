// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package currentorganizationaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/currentorganizationaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CurrentOrganizationAccountShowOutputOutputReference interface {
	cdktf.ComplexObject
	AccountLocator() *string
	AccountLocatorUrl() *string
	AccountName() *string
	AccountOldUrlLastUsed() *string
	AccountOldUrlSavedOn() *string
	AccountUrl() *string
	Comment() *string
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
	ConsumptionBillingEntityName() *string
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Edition() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CurrentOrganizationAccountShowOutput
	SetInternalValue(val *CurrentOrganizationAccountShowOutput)
	IsEventsAccount() cdktf.IResolvable
	IsOrgAdmin() cdktf.IResolvable
	IsOrganizationAccount() cdktf.IResolvable
	ManagedAccounts() *float64
	MarketplaceConsumerBillingEntityName() *string
	MarketplaceProviderBillingEntityName() *string
	OldAccountUrl() *string
	OrganizationName() *string
	OrganizationOldUrl() *string
	OrganizationOldUrlLastUsed() *string
	OrganizationOldUrlSavedOn() *string
	SnowflakeRegion() *string
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

// The jsii proxy struct for CurrentOrganizationAccountShowOutputOutputReference
type jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) AccountLocator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) AccountLocatorUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLocatorUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) AccountOldUrlLastUsed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountOldUrlLastUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) AccountOldUrlSavedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountOldUrlSavedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) AccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) ConsumptionBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumptionBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) InternalValue() *CurrentOrganizationAccountShowOutput {
	var returns *CurrentOrganizationAccountShowOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) IsEventsAccount() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEventsAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) IsOrgAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isOrgAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) IsOrganizationAccount() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isOrganizationAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) ManagedAccounts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) MarketplaceConsumerBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketplaceConsumerBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) MarketplaceProviderBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketplaceProviderBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) OldAccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldAccountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) OrganizationOldUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) OrganizationOldUrlLastUsed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrlLastUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) OrganizationOldUrlSavedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrlSavedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) SnowflakeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowflakeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCurrentOrganizationAccountShowOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CurrentOrganizationAccountShowOutputOutputReference {
	_init_.Initialize()

	if err := validateNewCurrentOrganizationAccountShowOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccountShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCurrentOrganizationAccountShowOutputOutputReference_Override(c CurrentOrganizationAccountShowOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.currentOrganizationAccount.CurrentOrganizationAccountShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference)SetInternalValue(val *CurrentOrganizationAccountShowOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CurrentOrganizationAccountShowOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

