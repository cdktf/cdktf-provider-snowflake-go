//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalTableTagList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalTableTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalTableTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTableTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTableTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalTableTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalTableTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

