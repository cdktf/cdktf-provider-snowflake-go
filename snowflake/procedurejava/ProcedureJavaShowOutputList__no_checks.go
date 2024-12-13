// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurejava

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureJavaShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavaShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavaShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureJavaShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
