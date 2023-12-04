// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Task) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateImportFromParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (t *jsiiProxy_Task) validateMoveToIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTask_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateTask_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTask_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTask_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetAfterParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetAllowOverlappingExecutionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetErrorIntegrationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetScheduleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSessionParametersParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSqlStatementParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSuspendTaskAfterNumFailuresParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetUserTaskManagedInitialWarehouseSizeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetUserTaskTimeoutMsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetWarehouseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetWhenParameters(val *string) error {
	return nil
}

func validateNewTaskParameters(scope constructs.Construct, id *string, config *TaskConfig) error {
	return nil
}

