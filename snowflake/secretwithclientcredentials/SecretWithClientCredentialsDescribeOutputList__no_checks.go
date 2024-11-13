// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretwithclientcredentials

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecretWithClientCredentialsDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

