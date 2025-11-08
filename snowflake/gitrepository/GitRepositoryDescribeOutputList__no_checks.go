// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gitrepository

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitRepositoryDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GitRepositoryDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GitRepositoryDescribeOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GitRepositoryDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GitRepositoryDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GitRepositoryDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGitRepositoryDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

