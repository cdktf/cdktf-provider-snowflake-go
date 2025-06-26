// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package jobservice

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobService) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateImportFromParameters(id *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (j *jsiiProxy_JobService) validateMoveToIdParameters(id *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validatePutFromSpecificationParameters(value *JobServiceFromSpecification) error {
	return nil
}

func (j *jsiiProxy_JobService) validatePutFromSpecificationTemplateParameters(value *JobServiceFromSpecificationTemplate) error {
	return nil
}

func (j *jsiiProxy_JobService) validatePutTimeoutsParameters(value *JobServiceTimeouts) error {
	return nil
}

func validateJobService_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateJobService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJobService_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateJobService_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetComputePoolParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetExternalAccessIntegrationsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetQueryWarehouseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobService) validateSetSchemaParameters(val *string) error {
	return nil
}

func validateNewJobServiceParameters(scope constructs.Construct, id *string, config *JobServiceConfig) error {
	return nil
}

