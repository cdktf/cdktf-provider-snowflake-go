// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersSuspendTaskAfterNumFailuresList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersSuspendTaskAfterNumFailuresList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersSuspendTaskAfterNumFailuresList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersSuspendTaskAfterNumFailuresList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersSuspendTaskAfterNumFailuresList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersSuspendTaskAfterNumFailuresList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersSuspendTaskAfterNumFailuresListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
