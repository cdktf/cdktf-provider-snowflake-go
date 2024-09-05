// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package view

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_ViewColumnList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_ViewColumnList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_ViewColumnList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ViewColumnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ViewColumnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ViewColumnList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ViewColumnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewViewColumnListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

