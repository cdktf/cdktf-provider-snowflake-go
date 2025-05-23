// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestodatabaserole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v13/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v13/grantprivilegestodatabaserole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference interface {
	cdktf.ComplexObject
	All() GrantPrivilegesToDatabaseRoleOnSchemaObjectAllOutputReference
	AllInput() *GrantPrivilegesToDatabaseRoleOnSchemaObjectAll
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
	Future() GrantPrivilegesToDatabaseRoleOnSchemaObjectFutureOutputReference
	FutureInput() *GrantPrivilegesToDatabaseRoleOnSchemaObjectFuture
	InternalValue() *GrantPrivilegesToDatabaseRoleOnSchemaObject
	SetInternalValue(val *GrantPrivilegesToDatabaseRoleOnSchemaObject)
	ObjectName() *string
	SetObjectName(val *string)
	ObjectNameInput() *string
	ObjectType() *string
	SetObjectType(val *string)
	ObjectTypeInput() *string
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
	PutAll(value *GrantPrivilegesToDatabaseRoleOnSchemaObjectAll)
	PutFuture(value *GrantPrivilegesToDatabaseRoleOnSchemaObjectFuture)
	ResetAll()
	ResetFuture()
	ResetObjectName()
	ResetObjectType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference
type jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) All() GrantPrivilegesToDatabaseRoleOnSchemaObjectAllOutputReference {
	var returns GrantPrivilegesToDatabaseRoleOnSchemaObjectAllOutputReference
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) AllInput() *GrantPrivilegesToDatabaseRoleOnSchemaObjectAll {
	var returns *GrantPrivilegesToDatabaseRoleOnSchemaObjectAll
	_jsii_.Get(
		j,
		"allInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) Future() GrantPrivilegesToDatabaseRoleOnSchemaObjectFutureOutputReference {
	var returns GrantPrivilegesToDatabaseRoleOnSchemaObjectFutureOutputReference
	_jsii_.Get(
		j,
		"future",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) FutureInput() *GrantPrivilegesToDatabaseRoleOnSchemaObjectFuture {
	var returns *GrantPrivilegesToDatabaseRoleOnSchemaObjectFuture
	_jsii_.Get(
		j,
		"futureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) InternalValue() *GrantPrivilegesToDatabaseRoleOnSchemaObject {
	var returns *GrantPrivilegesToDatabaseRoleOnSchemaObject
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference {
	_init_.Initialize()

	if err := validateNewGrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference_Override(g GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.grantPrivilegesToDatabaseRole.GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference)SetInternalValue(val *GrantPrivilegesToDatabaseRoleOnSchemaObject) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference)SetObjectName(val *string) {
	if err := j.validateSetObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectName",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference)SetObjectType(val *string) {
	if err := j.validateSetObjectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectType",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) PutAll(value *GrantPrivilegesToDatabaseRoleOnSchemaObjectAll) {
	if err := g.validatePutAllParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAll",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) PutFuture(value *GrantPrivilegesToDatabaseRoleOnSchemaObjectFuture) {
	if err := g.validatePutFutureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFuture",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ResetAll() {
	_jsii_.InvokeVoid(
		g,
		"resetAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ResetFuture() {
	_jsii_.InvokeVoid(
		g,
		"resetFuture",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ResetObjectName() {
	_jsii_.InvokeVoid(
		g,
		"resetObjectName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ResetObjectType() {
	_jsii_.InvokeVoid(
		g,
		"resetObjectType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GrantPrivilegesToDatabaseRoleOnSchemaObjectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

