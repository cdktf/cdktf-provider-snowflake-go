// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageIntegrationDescribeOutputStorageAwsObjectAclList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputStorageAwsObjectAclList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputStorageAwsObjectAclList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputStorageAwsObjectAclList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputStorageAwsObjectAclList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputStorageAwsObjectAclList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageIntegrationDescribeOutputStorageAwsObjectAclListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

