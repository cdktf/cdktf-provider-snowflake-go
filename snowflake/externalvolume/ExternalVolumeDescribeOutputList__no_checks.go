// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package externalvolume

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalVolumeDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_ExternalVolumeDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalVolumeDescribeOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalVolumeDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

