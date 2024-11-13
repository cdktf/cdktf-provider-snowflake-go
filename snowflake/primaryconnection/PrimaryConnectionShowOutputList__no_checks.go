// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package primaryconnection

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrimaryConnectionShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrimaryConnectionShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrimaryConnectionShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrimaryConnectionShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrimaryConnectionShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrimaryConnectionShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrimaryConnectionShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

