// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package account

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v13/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v13/account/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountShowOutputOutputReference interface {
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
	DroppedOn() *string
	Edition() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AccountShowOutput
	SetInternalValue(val *AccountShowOutput)
	IsEventsAccount() cdktf.IResolvable
	IsOrgAdmin() cdktf.IResolvable
	IsOrganizationAccount() cdktf.IResolvable
	ManagedAccounts() *float64
	MarketplaceConsumerBillingEntityName() *string
	MarketplaceProviderBillingEntityName() *string
	MovedOn() *string
	MovedToOrganization() *string
	OldAccountUrl() *string
	OrganizationName() *string
	OrganizationOldUrl() *string
	OrganizationOldUrlLastUsed() *string
	OrganizationOldUrlSavedOn() *string
	OrganizationUrlExpirationOn() *string
	RegionGroup() *string
	RestoredOn() *string
	ScheduledDeletionTime() *string
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

// The jsii proxy struct for AccountShowOutputOutputReference
type jsiiProxy_AccountShowOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccountShowOutputOutputReference) AccountLocator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) AccountLocatorUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLocatorUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) AccountOldUrlLastUsed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountOldUrlLastUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) AccountOldUrlSavedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountOldUrlSavedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) AccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) ConsumptionBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumptionBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) DroppedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"droppedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) InternalValue() *AccountShowOutput {
	var returns *AccountShowOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) IsEventsAccount() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEventsAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) IsOrgAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isOrgAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) IsOrganizationAccount() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isOrganizationAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) ManagedAccounts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) MarketplaceConsumerBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketplaceConsumerBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) MarketplaceProviderBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketplaceProviderBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) MovedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"movedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) MovedToOrganization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"movedToOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) OldAccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldAccountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) OrganizationOldUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) OrganizationOldUrlLastUsed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrlLastUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) OrganizationOldUrlSavedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrlSavedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) OrganizationUrlExpirationOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationUrlExpirationOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) RegionGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) RestoredOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoredOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) ScheduledDeletionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledDeletionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) SnowflakeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowflakeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountShowOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountShowOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccountShowOutputOutputReference {
	_init_.Initialize()

	if err := validateNewAccountShowOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountShowOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.account.AccountShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccountShowOutputOutputReference_Override(a AccountShowOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.account.AccountShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccountShowOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountShowOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountShowOutputOutputReference)SetInternalValue(val *AccountShowOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountShowOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountShowOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountShowOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountShowOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

