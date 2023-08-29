// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestorole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v9/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v9/grantprivilegestorole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantPrivilegesToRoleOnSchemaOutputReference interface {
	cdktf.ComplexObject
	AllSchemasInDatabase() *string
	SetAllSchemasInDatabase(val *string)
	AllSchemasInDatabaseInput() *string
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
	FutureSchemasInDatabase() *string
	SetFutureSchemasInDatabase(val *string)
	FutureSchemasInDatabaseInput() *string
	InternalValue() *GrantPrivilegesToRoleOnSchema
	SetInternalValue(val *GrantPrivilegesToRoleOnSchema)
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
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
	ResetAllSchemasInDatabase()
	ResetFutureSchemasInDatabase()
	ResetSchemaName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GrantPrivilegesToRoleOnSchemaOutputReference
type jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) AllSchemasInDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allSchemasInDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) AllSchemasInDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allSchemasInDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) FutureSchemasInDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"futureSchemasInDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) FutureSchemasInDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"futureSchemasInDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) InternalValue() *GrantPrivilegesToRoleOnSchema {
	var returns *GrantPrivilegesToRoleOnSchema
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGrantPrivilegesToRoleOnSchemaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GrantPrivilegesToRoleOnSchemaOutputReference {
	_init_.Initialize()

	if err := validateNewGrantPrivilegesToRoleOnSchemaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRoleOnSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGrantPrivilegesToRoleOnSchemaOutputReference_Override(g GrantPrivilegesToRoleOnSchemaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToRole.GrantPrivilegesToRoleOnSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference)SetAllSchemasInDatabase(val *string) {
	if err := j.validateSetAllSchemasInDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allSchemasInDatabase",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference)SetFutureSchemasInDatabase(val *string) {
	if err := j.validateSetFutureSchemasInDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"futureSchemasInDatabase",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference)SetInternalValue(val *GrantPrivilegesToRoleOnSchema) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) ResetAllSchemasInDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetAllSchemasInDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) ResetFutureSchemasInDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetFutureSchemasInDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) ResetSchemaName() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GrantPrivilegesToRoleOnSchemaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

