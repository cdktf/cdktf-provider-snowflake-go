// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersStrictJsonOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersStrictJsonOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersStrictJsonOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersStrictJsonOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersStrictJsonOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersStrictJsonOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersStrictJsonOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

