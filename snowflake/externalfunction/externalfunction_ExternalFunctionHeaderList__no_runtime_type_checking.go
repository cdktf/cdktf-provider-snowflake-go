//go:build no_runtime_type_checking

package externalfunction

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalFunctionHeaderList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalFunctionHeaderList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalFunctionHeaderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalFunctionHeaderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalFunctionHeaderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalFunctionHeaderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalFunctionHeaderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

