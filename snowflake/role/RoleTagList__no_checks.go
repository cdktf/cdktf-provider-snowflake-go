// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package role

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleTagList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RoleTagList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RoleTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RoleTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RoleTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RoleTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoleTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

