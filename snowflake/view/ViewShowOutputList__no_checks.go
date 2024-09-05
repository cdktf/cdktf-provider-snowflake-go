// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package view

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_ViewShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_ViewShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_ViewShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ViewShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ViewShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ViewShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewViewShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

