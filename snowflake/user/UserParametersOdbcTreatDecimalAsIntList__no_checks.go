// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersOdbcTreatDecimalAsIntList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersOdbcTreatDecimalAsIntList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersOdbcTreatDecimalAsIntList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersOdbcTreatDecimalAsIntList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersOdbcTreatDecimalAsIntList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersOdbcTreatDecimalAsIntList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersOdbcTreatDecimalAsIntListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

