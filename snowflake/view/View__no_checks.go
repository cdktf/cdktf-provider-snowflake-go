// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package view

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_View) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (v *jsiiProxy_View) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_View) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateImportFromParameters(id *string) error {
	return nil
}

func (v *jsiiProxy_View) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (v *jsiiProxy_View) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (v *jsiiProxy_View) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (v *jsiiProxy_View) validateMoveToIdParameters(id *string) error {
	return nil
}

func (v *jsiiProxy_View) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (v *jsiiProxy_View) validatePutAggregationPolicyParameters(value *ViewAggregationPolicy) error {
	return nil
}

func (v *jsiiProxy_View) validatePutColumnParameters(value interface{}) error {
	return nil
}

func (v *jsiiProxy_View) validatePutDataMetricFunctionParameters(value interface{}) error {
	return nil
}

func (v *jsiiProxy_View) validatePutDataMetricScheduleParameters(value *ViewDataMetricSchedule) error {
	return nil
}

func (v *jsiiProxy_View) validatePutRowAccessPolicyParameters(value *ViewRowAccessPolicy) error {
	return nil
}

func validateView_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateView_IsConstructParameters(x interface{}) error {
	return nil
}

func validateView_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateView_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_View) validateSetChangeTrackingParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_View) validateSetCopyGrantsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_View) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_View) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetIsRecursiveParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetIsSecureParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetIsTemporaryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_View) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_View) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_View) validateSetStatementParameters(val *string) error {
	return nil
}

func validateNewViewParameters(scope constructs.Construct, id *string, config *ViewConfig) error {
	return nil
}

