// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package stage

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stage) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validatePutTagParameters(value interface{}) error {
	return nil
}

func validateStage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateStage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateStage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetAwsExternalIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetCopyOptionsParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetCredentialsParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetDirectoryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetEncryptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetFileFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetSnowflakeIamUserParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetStorageIntegrationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stage) validateSetUrlParameters(val *string) error {
	return nil
}

func validateNewStageParameters(scope constructs.Construct, id *string, config *StageConfig) error {
	return nil
}

