// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sequence

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Sequence) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Sequence) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSequence_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSequence_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSequence_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSequence_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetIncrementParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Sequence) validateSetSchemaParameters(val *string) error {
	return nil
}

func validateNewSequenceParameters(scope constructs.Construct, id *string, config *SequenceConfig) error {
	return nil
}

