// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurepython

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedurePythonShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedurePythonShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedurePythonShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedurePythonShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

