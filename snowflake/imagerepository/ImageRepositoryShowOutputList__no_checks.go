// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package imagerepository

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImageRepositoryShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_ImageRepositoryShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_ImageRepositoryShowOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ImageRepositoryShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ImageRepositoryShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ImageRepositoryShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewImageRepositoryShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

