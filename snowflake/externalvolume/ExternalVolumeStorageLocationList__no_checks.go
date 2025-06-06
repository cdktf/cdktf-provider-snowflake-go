// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package externalvolume

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalVolumeStorageLocationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_ExternalVolumeStorageLocationList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalVolumeStorageLocationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeStorageLocationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeStorageLocationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeStorageLocationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeStorageLocationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalVolumeStorageLocationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

