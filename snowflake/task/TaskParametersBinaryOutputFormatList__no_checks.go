// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersBinaryOutputFormatList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersBinaryOutputFormatList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersBinaryOutputFormatList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersBinaryOutputFormatList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersBinaryOutputFormatList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersBinaryOutputFormatList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersBinaryOutputFormatListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

