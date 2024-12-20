// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeaccounts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/datasnowflakeaccounts/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeAccountsAccountsShowOutputOutputReference interface {
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
	InternalValue() *DataSnowflakeAccountsAccountsShowOutput
	SetInternalValue(val *DataSnowflakeAccountsAccountsShowOutput)
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

// The jsii proxy struct for DataSnowflakeAccountsAccountsShowOutputOutputReference
type jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) AccountLocator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) AccountLocatorUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLocatorUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) AccountOldUrlLastUsed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountOldUrlLastUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) AccountOldUrlSavedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountOldUrlSavedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) AccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) ConsumptionBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumptionBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) DroppedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"droppedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) InternalValue() *DataSnowflakeAccountsAccountsShowOutput {
	var returns *DataSnowflakeAccountsAccountsShowOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) IsEventsAccount() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEventsAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) IsOrgAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isOrgAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) IsOrganizationAccount() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isOrganizationAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) ManagedAccounts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) MarketplaceConsumerBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketplaceConsumerBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) MarketplaceProviderBillingEntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketplaceProviderBillingEntityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) MovedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"movedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) MovedToOrganization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"movedToOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) OldAccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldAccountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) OrganizationOldUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) OrganizationOldUrlLastUsed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrlLastUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) OrganizationOldUrlSavedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationOldUrlSavedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) OrganizationUrlExpirationOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationUrlExpirationOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) RegionGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) RestoredOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoredOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) ScheduledDeletionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledDeletionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) SnowflakeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowflakeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSnowflakeAccountsAccountsShowOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeAccountsAccountsShowOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeAccountsAccountsShowOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeAccounts.DataSnowflakeAccountsAccountsShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeAccountsAccountsShowOutputOutputReference_Override(d DataSnowflakeAccountsAccountsShowOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeAccounts.DataSnowflakeAccountsAccountsShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference)SetInternalValue(val *DataSnowflakeAccountsAccountsShowOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsShowOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

