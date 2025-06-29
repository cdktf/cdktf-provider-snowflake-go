// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package computepool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputePoolShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComputePoolShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputePoolShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputePoolShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputePoolShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputePoolShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputePoolShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

