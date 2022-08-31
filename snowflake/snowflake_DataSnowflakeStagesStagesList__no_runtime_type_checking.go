//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeStagesStagesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeStagesStagesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeStagesStagesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeStagesStagesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeStagesStagesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeStagesStagesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

