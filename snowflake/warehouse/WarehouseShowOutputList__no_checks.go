// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package warehouse

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WarehouseShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WarehouseShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WarehouseShowOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WarehouseShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WarehouseShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WarehouseShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWarehouseShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

