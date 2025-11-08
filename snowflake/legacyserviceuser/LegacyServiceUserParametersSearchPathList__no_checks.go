// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package legacyserviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LegacyServiceUserParametersSearchPathList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersSearchPathList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersSearchPathList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersSearchPathList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersSearchPathList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersSearchPathList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLegacyServiceUserParametersSearchPathListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

