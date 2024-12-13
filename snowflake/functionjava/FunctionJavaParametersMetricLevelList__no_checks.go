// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionjava

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionJavaParametersMetricLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionJavaParametersMetricLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionJavaParametersMetricLevelList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaParametersMetricLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaParametersMetricLevelList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaParametersMetricLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionJavaParametersMetricLevelListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

