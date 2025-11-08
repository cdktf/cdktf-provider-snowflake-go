// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersQuotedIdentifiersIgnoreCaseList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersQuotedIdentifiersIgnoreCaseList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersQuotedIdentifiersIgnoreCaseList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersQuotedIdentifiersIgnoreCaseList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersQuotedIdentifiersIgnoreCaseList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersQuotedIdentifiersIgnoreCaseList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersQuotedIdentifiersIgnoreCaseListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

