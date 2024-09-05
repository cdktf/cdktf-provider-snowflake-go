// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersClientPrefetchThreadsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersClientPrefetchThreadsList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersClientPrefetchThreadsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersClientPrefetchThreadsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersClientPrefetchThreadsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersClientPrefetchThreadsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersClientPrefetchThreadsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

