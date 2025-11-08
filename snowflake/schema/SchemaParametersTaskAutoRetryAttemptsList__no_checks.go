// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schema

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaParametersTaskAutoRetryAttemptsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SchemaParametersTaskAutoRetryAttemptsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SchemaParametersTaskAutoRetryAttemptsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersTaskAutoRetryAttemptsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersTaskAutoRetryAttemptsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersTaskAutoRetryAttemptsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSchemaParametersTaskAutoRetryAttemptsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

