// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package scimintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScimIntegrationDescribeOutputSyncPasswordList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScimIntegrationDescribeOutputSyncPasswordList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScimIntegrationDescribeOutputSyncPasswordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputSyncPasswordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputSyncPasswordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputSyncPasswordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScimIntegrationDescribeOutputSyncPasswordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

