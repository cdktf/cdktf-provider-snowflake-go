// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package database

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Database) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_Database) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_Database) validatePutReplicationParameters(value *DatabaseReplication) error {
	return nil
}

func validateDatabase_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDatabase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabase_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDatabase_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetCatalogParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetDataRetentionTimeInDaysParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetDefaultDdlCollationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetDropPublicSchemaOnCreationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetEnableConsoleOutputParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetExternalVolumeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetIsTransientParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetLogLevelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetMaxDataExtensionTimeInDaysParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetQuotedIdentifiersIgnoreCaseParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetReplaceInvalidCharactersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetStorageSerializationPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetSuspendTaskAfterNumFailuresParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetTaskAutoRetryAttemptsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetTraceLevelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetUserTaskManagedInitialWarehouseSizeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetUserTaskMinimumTriggerIntervalInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Database) validateSetUserTaskTimeoutMsParameters(val *float64) error {
	return nil
}

func validateNewDatabaseParameters(scope constructs.Construct, id *string, config *DatabaseConfig) error {
	return nil
}

