// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionsql

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionSqlParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionSqlParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionSqlParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionSqlParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

