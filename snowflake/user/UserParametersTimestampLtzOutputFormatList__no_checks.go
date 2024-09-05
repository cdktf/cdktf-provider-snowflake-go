// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersTimestampLtzOutputFormatList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersTimestampLtzOutputFormatList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersTimestampLtzOutputFormatList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersTimestampLtzOutputFormatList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersTimestampLtzOutputFormatList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersTimestampLtzOutputFormatList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersTimestampLtzOutputFormatListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

