// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gitrepository

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitRepositoryShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GitRepositoryShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GitRepositoryShowOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GitRepositoryShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GitRepositoryShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GitRepositoryShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGitRepositoryShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

