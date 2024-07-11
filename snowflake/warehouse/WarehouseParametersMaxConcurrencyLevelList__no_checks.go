// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package warehouse

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WarehouseParametersMaxConcurrencyLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WarehouseParametersMaxConcurrencyLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WarehouseParametersMaxConcurrencyLevelList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WarehouseParametersMaxConcurrencyLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WarehouseParametersMaxConcurrencyLevelList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WarehouseParametersMaxConcurrencyLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWarehouseParametersMaxConcurrencyLevelListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

