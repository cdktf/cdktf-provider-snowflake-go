//go:build no_runtime_type_checking

package proceduregrant

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureGrantArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureGrantArgumentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureGrantArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProcedureGrantArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureGrantArgumentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureGrantArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureGrantArgumentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

