// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package maskingpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaskingPolicySignatureColumnList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MaskingPolicySignatureColumnList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MaskingPolicySignatureColumnList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicySignatureColumnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicySignatureColumnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicySignatureColumnList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicySignatureColumnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMaskingPolicySignatureColumnListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

