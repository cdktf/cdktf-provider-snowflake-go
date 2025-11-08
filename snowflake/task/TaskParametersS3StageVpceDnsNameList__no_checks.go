// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersS3StageVpceDnsNameList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersS3StageVpceDnsNameList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersS3StageVpceDnsNameList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersS3StageVpceDnsNameList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersS3StageVpceDnsNameList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersS3StageVpceDnsNameList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersS3StageVpceDnsNameListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

