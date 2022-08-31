//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeExternalTablesExternalTablesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeExternalTablesExternalTablesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeExternalTablesExternalTablesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeExternalTablesExternalTablesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeExternalTablesExternalTablesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeExternalTablesExternalTablesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

