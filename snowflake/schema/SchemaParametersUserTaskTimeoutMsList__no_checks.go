// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schema

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaParametersUserTaskTimeoutMsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SchemaParametersUserTaskTimeoutMsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SchemaParametersUserTaskTimeoutMsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersUserTaskTimeoutMsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersUserTaskTimeoutMsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersUserTaskTimeoutMsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSchemaParametersUserTaskTimeoutMsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

