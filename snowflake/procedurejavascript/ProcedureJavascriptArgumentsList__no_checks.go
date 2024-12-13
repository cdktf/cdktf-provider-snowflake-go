// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurejavascript

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureJavascriptArgumentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavascriptArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavascriptArgumentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavascriptArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavascriptArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavascriptArgumentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavascriptArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureJavascriptArgumentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

