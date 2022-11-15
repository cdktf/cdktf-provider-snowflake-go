//go:build no_runtime_type_checking

package datasnowflaketasks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeTasksTasksList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeTasksTasksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeTasksTasksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeTasksTasksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeTasksTasksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeTasksTasksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

