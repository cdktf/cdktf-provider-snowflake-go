// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersErrorOnNondeterministicUpdateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersErrorOnNondeterministicUpdateList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersErrorOnNondeterministicUpdateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersErrorOnNondeterministicUpdateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersErrorOnNondeterministicUpdateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersErrorOnNondeterministicUpdateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersErrorOnNondeterministicUpdateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

