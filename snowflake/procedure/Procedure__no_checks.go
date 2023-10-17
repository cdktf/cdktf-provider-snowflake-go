// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedure

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Procedure) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateImportFromParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (p *jsiiProxy_Procedure) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (p *jsiiProxy_Procedure) validatePutArgumentsParameters(value interface{}) error {
	return nil
}

func validateProcedure_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateProcedure_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProcedure_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateProcedure_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetExecuteAsParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetHandlerParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetImportsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetLanguageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetNullInputBehaviorParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetPackagesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetReturnBehaviorParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetReturnTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetRuntimeVersionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Procedure) validateSetStatementParameters(val *string) error {
	return nil
}

func validateNewProcedureParameters(scope constructs.Construct, id *string, config *ProcedureConfig) error {
	return nil
}

