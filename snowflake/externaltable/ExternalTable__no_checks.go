// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package externaltable

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalTable) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateImportFromParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateMoveToIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validatePutColumnParameters(value interface{}) error {
	return nil
}

func (e *jsiiProxy_ExternalTable) validatePutTagParameters(value interface{}) error {
	return nil
}

func validateExternalTable_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateExternalTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalTable_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateExternalTable_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetAutoRefreshParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetAwsSnsTopicParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetCopyGrantsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetFileFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetPartitionByParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetPatternParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetRefreshOnCreateParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTable) validateSetSchemaParameters(val *string) error {
	return nil
}

func validateNewExternalTableParameters(scope constructs.Construct, id *string, config *ExternalTableConfig) error {
	return nil
}

