// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersPreventUnloadToInternalStagesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersPreventUnloadToInternalStagesList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersPreventUnloadToInternalStagesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersPreventUnloadToInternalStagesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersPreventUnloadToInternalStagesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersPreventUnloadToInternalStagesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersPreventUnloadToInternalStagesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

