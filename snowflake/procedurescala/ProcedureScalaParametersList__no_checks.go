// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurescala

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureScalaParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureScalaParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureScalaParametersList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureScalaParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureScalaParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureScalaParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureScalaParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

