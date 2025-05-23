// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secondaryconnection

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecondaryConnection) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnection) validatePutTimeoutsParameters(value *SecondaryConnectionTimeouts) error {
	return nil
}

func validateSecondaryConnection_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSecondaryConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecondaryConnection_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSecondaryConnection_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnection) validateSetAsReplicaOfParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnection) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnection) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnection) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnection) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnection) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnection) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnection) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewSecondaryConnectionParameters(scope constructs.Construct, id *string, config *SecondaryConnectionConfig) error {
	return nil
}

