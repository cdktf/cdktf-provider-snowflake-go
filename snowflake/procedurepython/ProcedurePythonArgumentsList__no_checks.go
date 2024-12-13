// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurepython

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedurePythonArgumentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedurePythonArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedurePythonArgumentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonArgumentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedurePythonArgumentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

