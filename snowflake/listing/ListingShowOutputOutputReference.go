// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package listing

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/listing/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ListingShowOutputOutputReference interface {
	cdktf.ComplexObject
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
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DetailedTargetAccounts() *string
	Distribution() *string
	// Experimental.
	Fqn() *string
	GlobalName() *string
	InternalValue() *ListingShowOutput
	SetInternalValue(val *ListingShowOutput)
	IsApplication() cdktf.IResolvable
	IsByRequest() cdktf.IResolvable
	IsLimitedTrial() cdktf.IResolvable
	IsMonetized() cdktf.IResolvable
	IsMountlessQueryable() cdktf.IResolvable
	IsTargeted() cdktf.IResolvable
	Name() *string
	OrganizationProfileName() *string
	Owner() *string
	OwnerRoleType() *string
	Profile() *string
	PublishedOn() *string
	Regions() *string
	RejectedOn() *string
	ReviewState() *string
	State() *string
	Subtitle() *string
	TargetAccounts() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	UniformListingLocator() *string
	UpdatedOn() *string
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

// The jsii proxy struct for ListingShowOutputOutputReference
type jsiiProxy_ListingShowOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) DetailedTargetAccounts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detailedTargetAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Distribution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) GlobalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) InternalValue() *ListingShowOutput {
	var returns *ListingShowOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) IsApplication() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) IsByRequest() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isByRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) IsLimitedTrial() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isLimitedTrial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) IsMonetized() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isMonetized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) IsMountlessQueryable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isMountlessQueryable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) IsTargeted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isTargeted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) OrganizationProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) OwnerRoleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerRoleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) PublishedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Regions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) RejectedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rejectedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) ReviewState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Subtitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) TargetAccounts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) UniformListingLocator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniformListingLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListingShowOutputOutputReference) UpdatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedOn",
		&returns,
	)
	return returns
}


func NewListingShowOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ListingShowOutputOutputReference {
	_init_.Initialize()

	if err := validateNewListingShowOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ListingShowOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.listing.ListingShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewListingShowOutputOutputReference_Override(l ListingShowOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.listing.ListingShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_ListingShowOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ListingShowOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ListingShowOutputOutputReference)SetInternalValue(val *ListingShowOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ListingShowOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ListingShowOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_ListingShowOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListingShowOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

