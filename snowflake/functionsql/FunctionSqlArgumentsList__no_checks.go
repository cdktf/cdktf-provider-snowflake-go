// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionsql

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionSqlArgumentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionSqlArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionSqlArgumentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlArgumentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionSqlArgumentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

