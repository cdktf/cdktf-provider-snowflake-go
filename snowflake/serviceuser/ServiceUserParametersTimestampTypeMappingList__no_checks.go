// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceUserParametersTimestampTypeMappingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersTimestampTypeMappingList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersTimestampTypeMappingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersTimestampTypeMappingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersTimestampTypeMappingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersTimestampTypeMappingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceUserParametersTimestampTypeMappingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

