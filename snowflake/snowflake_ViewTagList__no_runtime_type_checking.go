//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_ViewTagList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_ViewTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ViewTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ViewTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ViewTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ViewTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewViewTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

