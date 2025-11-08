// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package streamondirectorytable

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamOnDirectoryTableShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StreamOnDirectoryTableShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StreamOnDirectoryTableShowOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StreamOnDirectoryTableShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StreamOnDirectoryTableShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StreamOnDirectoryTableShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStreamOnDirectoryTableShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

