// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurejava

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureJavaImportsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavaImportsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavaImportsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaImportsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaImportsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaImportsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaImportsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureJavaImportsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

