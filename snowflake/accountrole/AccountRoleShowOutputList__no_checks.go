// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accountrole

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountRoleShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccountRoleShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccountRoleShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccountRoleShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccountRoleShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccountRoleShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccountRoleShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

