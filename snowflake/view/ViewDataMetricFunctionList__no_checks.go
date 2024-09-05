// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package view

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_ViewDataMetricFunctionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_ViewDataMetricFunctionList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_ViewDataMetricFunctionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ViewDataMetricFunctionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ViewDataMetricFunctionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ViewDataMetricFunctionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ViewDataMetricFunctionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewViewDataMetricFunctionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

