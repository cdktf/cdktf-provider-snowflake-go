// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceUserParametersWeekOfYearPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersWeekOfYearPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersWeekOfYearPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersWeekOfYearPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersWeekOfYearPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersWeekOfYearPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceUserParametersWeekOfYearPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

