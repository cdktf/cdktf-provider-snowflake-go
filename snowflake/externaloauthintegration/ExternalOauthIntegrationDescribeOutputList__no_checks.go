// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package externaloauthintegration

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalOauthIntegrationDescribeOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalOauthIntegrationDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalOauthIntegrationDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

