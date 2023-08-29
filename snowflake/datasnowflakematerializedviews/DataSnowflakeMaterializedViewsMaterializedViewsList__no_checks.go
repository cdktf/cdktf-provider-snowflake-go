// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datasnowflakematerializedviews

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeMaterializedViewsMaterializedViewsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeMaterializedViewsMaterializedViewsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeMaterializedViewsMaterializedViewsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeMaterializedViewsMaterializedViewsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeMaterializedViewsMaterializedViewsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeMaterializedViewsMaterializedViewsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

