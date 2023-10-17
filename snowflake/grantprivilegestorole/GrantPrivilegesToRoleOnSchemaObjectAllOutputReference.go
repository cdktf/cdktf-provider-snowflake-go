// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestorole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v10/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v10/grantprivilegestorole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantPrivilegesToRoleOnSchemaObjectAllOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InDatabase() *string
	SetInDatabase(val *string)
	InDatabaseInput() *string
	InSchema() *string
	SetInSchema(val *string)
	InSchemaInput() *string
	InternalValue() *GrantPrivilegesToRoleOnSchemaObjectAll
	SetInternalValue(val *GrantPrivilegesToRoleOnSchemaObjectAll)
	ObjectTypePlural() *string
	SetObjectTypePlural(val *string)
	ObjectTypePluralInput() *string
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
	ResetInDatabase()
	ResetInSchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GrantPrivilegesToRoleOnSchemaObjectAllOutputReference
type jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) InDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) InDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) InSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) InSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) InternalValue() *GrantPrivilegesToRoleOnSchemaObjectAll {
	var returns *GrantPrivilegesToRoleOnSchemaObjectAll
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) ObjectTypePlural() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypePlural",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) ObjectTypePluralInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypePluralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGrantPrivilegesToRoleOnSchemaObjectAllOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GrantPrivilegesToRoleOnSchemaObjectAllOutputReference {
	_init_.Initialize()

	if err := validateNewGrantPrivilegesToRoleOnSchemaObjectAllOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRoleOnSchemaObjectAllOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGrantPrivilegesToRoleOnSchemaObjectAllOutputReference_Override(g GrantPrivilegesToRoleOnSchemaObjectAllOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRoleOnSchemaObjectAllOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference)SetInDatabase(val *string) {
	if err := j.validateSetInDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inDatabase",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference)SetInSchema(val *string) {
	if err := j.validateSetInSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inSchema",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference)SetInternalValue(val *GrantPrivilegesToRoleOnSchemaObjectAll) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference)SetObjectTypePlural(val *string) {
	if err := j.validateSetObjectTypePluralParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectTypePlural",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) ResetInDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetInDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) ResetInSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetInSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaObjectAllOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

