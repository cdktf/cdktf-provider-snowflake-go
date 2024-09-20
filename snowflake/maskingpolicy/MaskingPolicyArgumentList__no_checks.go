// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package maskingpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaskingPolicyArgumentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MaskingPolicyArgumentList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MaskingPolicyArgumentList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicyArgumentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicyArgumentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicyArgumentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicyArgumentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMaskingPolicyArgumentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

