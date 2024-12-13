// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionscala

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionScalaShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionScalaShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionScalaShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionScalaShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

