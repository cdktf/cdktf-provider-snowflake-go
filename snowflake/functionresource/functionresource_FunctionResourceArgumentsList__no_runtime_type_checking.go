//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package functionresource

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionResourceArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionResourceArgumentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionResourceArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionResourceArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionResourceArgumentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionResourceArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionResourceArgumentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
