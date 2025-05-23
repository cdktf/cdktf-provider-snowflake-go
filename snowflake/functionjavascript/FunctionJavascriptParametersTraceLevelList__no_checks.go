// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionjavascript

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionJavascriptParametersTraceLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionJavascriptParametersTraceLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionJavascriptParametersTraceLevelList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersTraceLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersTraceLevelList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersTraceLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionJavascriptParametersTraceLevelListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

