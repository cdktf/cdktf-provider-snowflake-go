// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package jobservice

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobServiceDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobServiceDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobServiceDescribeOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobServiceDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobServiceDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobServiceDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobServiceDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

