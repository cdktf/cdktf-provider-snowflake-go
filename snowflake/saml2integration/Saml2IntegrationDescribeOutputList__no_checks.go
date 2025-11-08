// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package saml2integration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Saml2IntegrationDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_Saml2IntegrationDescribeOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Saml2IntegrationDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSaml2IntegrationDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

