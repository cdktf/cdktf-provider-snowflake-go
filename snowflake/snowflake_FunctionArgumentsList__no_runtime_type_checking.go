//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionArgumentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionArgumentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionArgumentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

