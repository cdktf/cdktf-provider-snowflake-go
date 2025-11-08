// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package legacyserviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LegacyServiceUserParametersMultiStatementCountList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersMultiStatementCountList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersMultiStatementCountList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersMultiStatementCountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersMultiStatementCountList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersMultiStatementCountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLegacyServiceUserParametersMultiStatementCountListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

