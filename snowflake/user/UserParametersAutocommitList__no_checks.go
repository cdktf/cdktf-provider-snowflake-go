// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersAutocommitList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersAutocommitList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersAutocommitList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersAutocommitList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersAutocommitList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersAutocommitList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersAutocommitListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

