// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package computepool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputePoolDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComputePoolDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputePoolDescribeOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputePoolDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputePoolDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputePoolDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputePoolDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

