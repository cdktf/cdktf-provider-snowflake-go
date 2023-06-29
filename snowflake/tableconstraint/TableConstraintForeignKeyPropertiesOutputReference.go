package tableconstraint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v8/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v8/tableconstraint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TableConstraintForeignKeyPropertiesOutputReference interface {
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
	InternalValue() *TableConstraintForeignKeyProperties
	SetInternalValue(val *TableConstraintForeignKeyProperties)
	Match() *string
	SetMatch(val *string)
	MatchInput() *string
	OnDelete() *string
	SetOnDelete(val *string)
	OnDeleteInput() *string
	OnUpdate() *string
	SetOnUpdate(val *string)
	OnUpdateInput() *string
	References() TableConstraintForeignKeyPropertiesReferencesOutputReference
	ReferencesInput() *TableConstraintForeignKeyPropertiesReferences
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
	PutReferences(value *TableConstraintForeignKeyPropertiesReferences)
	ResetMatch()
	ResetOnDelete()
	ResetOnUpdate()
	ResetReferences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TableConstraintForeignKeyPropertiesOutputReference
type jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) InternalValue() *TableConstraintForeignKeyProperties {
	var returns *TableConstraintForeignKeyProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) Match() *string {
	var returns *string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) MatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) OnDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) OnDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) OnUpdate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) OnUpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) References() TableConstraintForeignKeyPropertiesReferencesOutputReference {
	var returns TableConstraintForeignKeyPropertiesReferencesOutputReference
	_jsii_.Get(
		j,
		"references",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ReferencesInput() *TableConstraintForeignKeyPropertiesReferences {
	var returns *TableConstraintForeignKeyPropertiesReferences
	_jsii_.Get(
		j,
		"referencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTableConstraintForeignKeyPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TableConstraintForeignKeyPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewTableConstraintForeignKeyPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.tableConstraint.TableConstraintForeignKeyPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTableConstraintForeignKeyPropertiesOutputReference_Override(t TableConstraintForeignKeyPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.tableConstraint.TableConstraintForeignKeyPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference)SetInternalValue(val *TableConstraintForeignKeyProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference)SetMatch(val *string) {
	if err := j.validateSetMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference)SetOnDelete(val *string) {
	if err := j.validateSetOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDelete",
		val,
	)
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference)SetOnUpdate(val *string) {
	if err := j.validateSetOnUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onUpdate",
		val,
	)
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) PutReferences(value *TableConstraintForeignKeyPropertiesReferences) {
	if err := t.validatePutReferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReferences",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ResetMatch() {
	_jsii_.InvokeVoid(
		t,
		"resetMatch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ResetOnDelete() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDelete",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ResetOnUpdate() {
	_jsii_.InvokeVoid(
		t,
		"resetOnUpdate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ResetReferences() {
	_jsii_.InvokeVoid(
		t,
		"resetReferences",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableConstraintForeignKeyPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

