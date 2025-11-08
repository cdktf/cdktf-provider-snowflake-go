// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cortexsearchservice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CortexSearchServiceDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CortexSearchServiceDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CortexSearchServiceDescribeOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CortexSearchServiceDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CortexSearchServiceDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CortexSearchServiceDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCortexSearchServiceDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

