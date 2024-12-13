// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package execute

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Execute) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateImportFromParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (e *jsiiProxy_Execute) validateMoveToIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Execute) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateExecute_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateExecute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExecute_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateExecute_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Execute) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Execute) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Execute) validateSetExecuteParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Execute) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Execute) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Execute) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Execute) validateSetQueryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Execute) validateSetRevertParameters(val *string) error {
	return nil
}

func validateNewExecuteParameters(scope constructs.Construct, id *string, config *ExecuteConfig) error {
	return nil
}

