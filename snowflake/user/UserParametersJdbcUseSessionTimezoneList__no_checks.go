// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersJdbcUseSessionTimezoneList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersJdbcUseSessionTimezoneList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersJdbcUseSessionTimezoneList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersJdbcUseSessionTimezoneList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersJdbcUseSessionTimezoneList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersJdbcUseSessionTimezoneList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersJdbcUseSessionTimezoneListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

