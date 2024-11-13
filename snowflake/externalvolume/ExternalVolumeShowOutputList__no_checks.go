// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package externalvolume

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalVolumeShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_ExternalVolumeShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalVolumeShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalVolumeShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

