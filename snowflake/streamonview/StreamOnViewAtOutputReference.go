// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamonview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/streamonview/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamOnViewAtOutputReference interface {
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
	InternalValue() *StreamOnViewAt
	SetInternalValue(val *StreamOnViewAt)
	Offset() *string
	SetOffset(val *string)
	OffsetInput() *string
	Statement() *string
	SetStatement(val *string)
	StatementInput() *string
	Stream() *string
	SetStream(val *string)
	StreamInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timestamp() *string
	SetTimestamp(val *string)
	TimestampInput() *string
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
	ResetOffset()
	ResetStatement()
	ResetStream()
	ResetTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StreamOnViewAtOutputReference
type jsiiProxy_StreamOnViewAtOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) InternalValue() *StreamOnViewAt {
	var returns *StreamOnViewAt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) Offset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) OffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) Statement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) StatementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) Stream() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) StreamInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) Timestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamOnViewAtOutputReference) TimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInput",
		&returns,
	)
	return returns
}


func NewStreamOnViewAtOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StreamOnViewAtOutputReference {
	_init_.Initialize()

	if err := validateNewStreamOnViewAtOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StreamOnViewAtOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.streamOnView.StreamOnViewAtOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStreamOnViewAtOutputReference_Override(s StreamOnViewAtOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.streamOnView.StreamOnViewAtOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetInternalValue(val *StreamOnViewAt) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetOffset(val *string) {
	if err := j.validateSetOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offset",
		val,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetStatement(val *string) {
	if err := j.validateSetStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statement",
		val,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetStream(val *string) {
	if err := j.validateSetStreamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stream",
		val,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StreamOnViewAtOutputReference)SetTimestamp(val *string) {
	if err := j.validateSetTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestamp",
		val,
	)
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) ResetOffset() {
	_jsii_.InvokeVoid(
		s,
		"resetOffset",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) ResetStatement() {
	_jsii_.InvokeVoid(
		s,
		"resetStatement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) ResetStream() {
	_jsii_.InvokeVoid(
		s,
		"resetStream",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) ResetTimestamp() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestamp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamOnViewAtOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

