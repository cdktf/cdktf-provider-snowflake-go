// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package warehouse

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WarehouseParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WarehouseParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WarehouseParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WarehouseParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WarehouseParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WarehouseParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWarehouseParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

