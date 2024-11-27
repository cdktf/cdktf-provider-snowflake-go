// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersWeekStartList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersWeekStartList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersWeekStartList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersWeekStartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersWeekStartList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersWeekStartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersWeekStartListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

