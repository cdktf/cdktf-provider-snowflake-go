// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package materializedview

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaterializedView) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateImportFromParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedView) validatePutTagParameters(value interface{}) error {
	return nil
}

func validateMaterializedView_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateMaterializedView_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMaterializedView_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateMaterializedView_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetIsSecureParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetOrReplaceParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetStatementParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaterializedView) validateSetWarehouseParameters(val *string) error {
	return nil
}

func validateNewMaterializedViewParameters(scope constructs.Construct, id *string, config *MaterializedViewConfig) error {
	return nil
}

