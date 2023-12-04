// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dynamictable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamicTable) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_DynamicTable) validatePutTargetLagParameters(value *DynamicTableTargetLag) error {
	return nil
}

func validateDynamicTable_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDynamicTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDynamicTable_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDynamicTable_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetOrReplaceParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetQueryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamicTable) validateSetWarehouseParameters(val *string) error {
	return nil
}

func validateNewDynamicTableParameters(scope constructs.Construct, id *string, config *DynamicTableConfig) error {
	return nil
}

