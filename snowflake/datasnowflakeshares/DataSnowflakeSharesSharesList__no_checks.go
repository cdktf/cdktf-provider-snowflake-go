//go:build no_runtime_type_checking

package datasnowflakeshares

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeSharesSharesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeSharesSharesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeSharesSharesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeSharesSharesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeSharesSharesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeSharesSharesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

