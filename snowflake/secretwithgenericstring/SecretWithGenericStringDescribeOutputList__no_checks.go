// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretwithgenericstring

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecretWithGenericStringDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

