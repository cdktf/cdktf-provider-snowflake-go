// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersNoorderSequenceAsDefaultList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersNoorderSequenceAsDefaultList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersNoorderSequenceAsDefaultList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersNoorderSequenceAsDefaultList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersNoorderSequenceAsDefaultList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersNoorderSequenceAsDefaultList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersNoorderSequenceAsDefaultListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

