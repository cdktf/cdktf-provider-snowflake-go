// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package proceduresql

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureSqlParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureSqlParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureSqlParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureSqlParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureSqlParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureSqlParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureSqlParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

