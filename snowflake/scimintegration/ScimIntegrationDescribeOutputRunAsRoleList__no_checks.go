// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package scimintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScimIntegrationDescribeOutputRunAsRoleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScimIntegrationDescribeOutputRunAsRoleList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScimIntegrationDescribeOutputRunAsRoleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputRunAsRoleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputRunAsRoleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputRunAsRoleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScimIntegrationDescribeOutputRunAsRoleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

