//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StageTagList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StageTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StageTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StageTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StageTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StageTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStageTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

