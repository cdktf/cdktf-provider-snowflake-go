// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionscala

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionScalaParametersLogLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionScalaParametersLogLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionScalaParametersLogLevelList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaParametersLogLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaParametersLogLevelList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaParametersLogLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionScalaParametersLogLevelListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

