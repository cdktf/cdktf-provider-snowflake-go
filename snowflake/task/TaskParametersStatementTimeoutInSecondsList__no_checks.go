// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersStatementTimeoutInSecondsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersStatementTimeoutInSecondsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersStatementTimeoutInSecondsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersStatementTimeoutInSecondsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersStatementTimeoutInSecondsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersStatementTimeoutInSecondsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersStatementTimeoutInSecondsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

