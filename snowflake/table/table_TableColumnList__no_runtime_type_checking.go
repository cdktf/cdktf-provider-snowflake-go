//go:build no_runtime_type_checking

package table

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TableColumnList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TableColumnList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TableColumnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TableColumnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TableColumnList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TableColumnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTableColumnListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

