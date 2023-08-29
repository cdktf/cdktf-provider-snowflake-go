// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package objectparameter

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObjectParameterObjectIdentifierList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObjectParameterObjectIdentifierList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObjectParameterObjectIdentifierList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObjectParameterObjectIdentifierList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObjectParameterObjectIdentifierList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObjectParameterObjectIdentifierList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObjectParameterObjectIdentifierListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

