// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package resourcemonitor

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceMonitorShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ResourceMonitorShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResourceMonitorShowOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResourceMonitorShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceMonitorShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResourceMonitorShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResourceMonitorShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

