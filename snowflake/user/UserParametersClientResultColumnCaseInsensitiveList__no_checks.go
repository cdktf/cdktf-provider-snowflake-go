// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersClientResultColumnCaseInsensitiveList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersClientResultColumnCaseInsensitiveList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersClientResultColumnCaseInsensitiveList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersClientResultColumnCaseInsensitiveList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersClientResultColumnCaseInsensitiveList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersClientResultColumnCaseInsensitiveList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersClientResultColumnCaseInsensitiveListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

