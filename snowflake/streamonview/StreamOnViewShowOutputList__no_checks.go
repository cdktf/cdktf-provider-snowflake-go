// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package streamonview

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamOnViewShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StreamOnViewShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StreamOnViewShowOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StreamOnViewShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StreamOnViewShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StreamOnViewShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStreamOnViewShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

