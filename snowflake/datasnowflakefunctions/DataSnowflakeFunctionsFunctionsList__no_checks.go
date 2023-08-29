// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datasnowflakefunctions

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeFunctionsFunctionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeFunctionsFunctionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeFunctionsFunctionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeFunctionsFunctionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeFunctionsFunctionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeFunctionsFunctionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

