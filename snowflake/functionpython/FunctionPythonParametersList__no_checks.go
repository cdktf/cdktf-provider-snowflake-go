// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionpython

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionPythonParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionPythonParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionPythonParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionPythonParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

