// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package saml2integration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Saml2IntegrationShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_Saml2IntegrationShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_Saml2IntegrationShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Saml2IntegrationShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Saml2IntegrationShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Saml2IntegrationShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSaml2IntegrationShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

