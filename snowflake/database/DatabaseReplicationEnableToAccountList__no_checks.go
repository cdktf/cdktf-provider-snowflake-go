// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package database

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseReplicationEnableToAccountList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseReplicationEnableToAccountList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseReplicationEnableToAccountList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseReplicationEnableToAccountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseReplicationEnableToAccountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseReplicationEnableToAccountList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseReplicationEnableToAccountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseReplicationEnableToAccountListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

