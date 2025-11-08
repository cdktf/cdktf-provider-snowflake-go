// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databaserole

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseRoleShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseRoleShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseRoleShowOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseRoleShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseRoleShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseRoleShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseRoleShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

