// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secondaryconnection

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecondaryConnectionShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnectionShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecondaryConnectionShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnectionShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnectionShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecondaryConnectionShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecondaryConnectionShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

