// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package maskingpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaskingPolicyDescribeOutputSignatureList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MaskingPolicyDescribeOutputSignatureList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MaskingPolicyDescribeOutputSignatureList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicyDescribeOutputSignatureList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicyDescribeOutputSignatureList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MaskingPolicyDescribeOutputSignatureList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMaskingPolicyDescribeOutputSignatureListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

