// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceUserParametersQueryTagList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersQueryTagList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersQueryTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersQueryTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersQueryTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersQueryTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceUserParametersQueryTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

