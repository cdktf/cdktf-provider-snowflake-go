// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakedatabases

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/datasnowflakedatabases/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeDatabasesDatabasesParametersOutputReference interface {
	cdktf.ComplexObject
	Catalog() DataSnowflakeDatabasesDatabasesParametersCatalogList
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
	DataRetentionTimeInDays() DataSnowflakeDatabasesDatabasesParametersDataRetentionTimeInDaysList
	DefaultDdlCollation() DataSnowflakeDatabasesDatabasesParametersDefaultDdlCollationList
	EnableConsoleOutput() DataSnowflakeDatabasesDatabasesParametersEnableConsoleOutputList
	ExternalVolume() DataSnowflakeDatabasesDatabasesParametersExternalVolumeList
	// Experimental.
	Fqn() *string
	InternalValue() *DataSnowflakeDatabasesDatabasesParameters
	SetInternalValue(val *DataSnowflakeDatabasesDatabasesParameters)
	LogLevel() DataSnowflakeDatabasesDatabasesParametersLogLevelList
	MaxDataExtensionTimeInDays() DataSnowflakeDatabasesDatabasesParametersMaxDataExtensionTimeInDaysList
	QuotedIdentifiersIgnoreCase() DataSnowflakeDatabasesDatabasesParametersQuotedIdentifiersIgnoreCaseList
	ReplaceInvalidCharacters() DataSnowflakeDatabasesDatabasesParametersReplaceInvalidCharactersList
	StorageSerializationPolicy() DataSnowflakeDatabasesDatabasesParametersStorageSerializationPolicyList
	SuspendTaskAfterNumFailures() DataSnowflakeDatabasesDatabasesParametersSuspendTaskAfterNumFailuresList
	TaskAutoRetryAttempts() DataSnowflakeDatabasesDatabasesParametersTaskAutoRetryAttemptsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TraceLevel() DataSnowflakeDatabasesDatabasesParametersTraceLevelList
	UserTaskManagedInitialWarehouseSize() DataSnowflakeDatabasesDatabasesParametersUserTaskManagedInitialWarehouseSizeList
	UserTaskMinimumTriggerIntervalInSeconds() DataSnowflakeDatabasesDatabasesParametersUserTaskMinimumTriggerIntervalInSecondsList
	UserTaskTimeoutMs() DataSnowflakeDatabasesDatabasesParametersUserTaskTimeoutMsList
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

// The jsii proxy struct for DataSnowflakeDatabasesDatabasesParametersOutputReference
type jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) Catalog() DataSnowflakeDatabasesDatabasesParametersCatalogList {
	var returns DataSnowflakeDatabasesDatabasesParametersCatalogList
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) DataRetentionTimeInDays() DataSnowflakeDatabasesDatabasesParametersDataRetentionTimeInDaysList {
	var returns DataSnowflakeDatabasesDatabasesParametersDataRetentionTimeInDaysList
	_jsii_.Get(
		j,
		"dataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) DefaultDdlCollation() DataSnowflakeDatabasesDatabasesParametersDefaultDdlCollationList {
	var returns DataSnowflakeDatabasesDatabasesParametersDefaultDdlCollationList
	_jsii_.Get(
		j,
		"defaultDdlCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) EnableConsoleOutput() DataSnowflakeDatabasesDatabasesParametersEnableConsoleOutputList {
	var returns DataSnowflakeDatabasesDatabasesParametersEnableConsoleOutputList
	_jsii_.Get(
		j,
		"enableConsoleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) ExternalVolume() DataSnowflakeDatabasesDatabasesParametersExternalVolumeList {
	var returns DataSnowflakeDatabasesDatabasesParametersExternalVolumeList
	_jsii_.Get(
		j,
		"externalVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) InternalValue() *DataSnowflakeDatabasesDatabasesParameters {
	var returns *DataSnowflakeDatabasesDatabasesParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) LogLevel() DataSnowflakeDatabasesDatabasesParametersLogLevelList {
	var returns DataSnowflakeDatabasesDatabasesParametersLogLevelList
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) MaxDataExtensionTimeInDays() DataSnowflakeDatabasesDatabasesParametersMaxDataExtensionTimeInDaysList {
	var returns DataSnowflakeDatabasesDatabasesParametersMaxDataExtensionTimeInDaysList
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) QuotedIdentifiersIgnoreCase() DataSnowflakeDatabasesDatabasesParametersQuotedIdentifiersIgnoreCaseList {
	var returns DataSnowflakeDatabasesDatabasesParametersQuotedIdentifiersIgnoreCaseList
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) ReplaceInvalidCharacters() DataSnowflakeDatabasesDatabasesParametersReplaceInvalidCharactersList {
	var returns DataSnowflakeDatabasesDatabasesParametersReplaceInvalidCharactersList
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) StorageSerializationPolicy() DataSnowflakeDatabasesDatabasesParametersStorageSerializationPolicyList {
	var returns DataSnowflakeDatabasesDatabasesParametersStorageSerializationPolicyList
	_jsii_.Get(
		j,
		"storageSerializationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) SuspendTaskAfterNumFailures() DataSnowflakeDatabasesDatabasesParametersSuspendTaskAfterNumFailuresList {
	var returns DataSnowflakeDatabasesDatabasesParametersSuspendTaskAfterNumFailuresList
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) TaskAutoRetryAttempts() DataSnowflakeDatabasesDatabasesParametersTaskAutoRetryAttemptsList {
	var returns DataSnowflakeDatabasesDatabasesParametersTaskAutoRetryAttemptsList
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) TraceLevel() DataSnowflakeDatabasesDatabasesParametersTraceLevelList {
	var returns DataSnowflakeDatabasesDatabasesParametersTraceLevelList
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) UserTaskManagedInitialWarehouseSize() DataSnowflakeDatabasesDatabasesParametersUserTaskManagedInitialWarehouseSizeList {
	var returns DataSnowflakeDatabasesDatabasesParametersUserTaskManagedInitialWarehouseSizeList
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) UserTaskMinimumTriggerIntervalInSeconds() DataSnowflakeDatabasesDatabasesParametersUserTaskMinimumTriggerIntervalInSecondsList {
	var returns DataSnowflakeDatabasesDatabasesParametersUserTaskMinimumTriggerIntervalInSecondsList
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) UserTaskTimeoutMs() DataSnowflakeDatabasesDatabasesParametersUserTaskTimeoutMsList {
	var returns DataSnowflakeDatabasesDatabasesParametersUserTaskTimeoutMsList
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}


func NewDataSnowflakeDatabasesDatabasesParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeDatabasesDatabasesParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeDatabasesDatabasesParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabasesDatabasesParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeDatabasesDatabasesParametersOutputReference_Override(d DataSnowflakeDatabasesDatabasesParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeDatabases.DataSnowflakeDatabasesDatabasesParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference)SetInternalValue(val *DataSnowflakeDatabasesDatabasesParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeDatabasesDatabasesParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

