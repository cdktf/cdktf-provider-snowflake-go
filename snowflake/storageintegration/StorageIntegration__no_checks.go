// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageIntegration) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegration) validatePutTimeoutsParameters(value *StorageIntegrationTimeouts) error {
	return nil
}

func validateStorageIntegration_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateStorageIntegration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStorageIntegration_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateStorageIntegration_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetAzureTenantIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetStorageAllowedLocationsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetStorageAwsObjectAclParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetStorageAwsRoleArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetStorageBlockedLocationsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetStorageProviderParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegration) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewStorageIntegrationParameters(scope constructs.Construct, id *string, config *StorageIntegrationConfig) error {
	return nil
}

