//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserTagList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

