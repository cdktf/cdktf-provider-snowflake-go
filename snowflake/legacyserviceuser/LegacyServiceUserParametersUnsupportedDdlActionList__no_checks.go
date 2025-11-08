// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package legacyserviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LegacyServiceUserParametersUnsupportedDdlActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersUnsupportedDdlActionList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersUnsupportedDdlActionList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersUnsupportedDdlActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersUnsupportedDdlActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersUnsupportedDdlActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLegacyServiceUserParametersUnsupportedDdlActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

