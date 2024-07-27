// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schema

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaParametersDataRetentionTimeInDaysList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SchemaParametersDataRetentionTimeInDaysList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SchemaParametersDataRetentionTimeInDaysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersDataRetentionTimeInDaysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersDataRetentionTimeInDaysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersDataRetentionTimeInDaysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSchemaParametersDataRetentionTimeInDaysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

