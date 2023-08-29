// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schema

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaTagList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SchemaTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SchemaTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SchemaTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchemaTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SchemaTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSchemaTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

