// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package legacyserviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LegacyServiceUserParametersTimezoneList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersTimezoneList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersTimezoneList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersTimezoneList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersTimezoneList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersTimezoneList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLegacyServiceUserParametersTimezoneListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

