// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Task) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateImportFromParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (t *jsiiProxy_Task) validateMoveToIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (t *jsiiProxy_Task) validatePutScheduleParameters(value *TaskSchedule) error {
	return nil
}

func validateTask_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateTask_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTask_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTask_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetAbortDetachedQueryParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetAfterParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetAllowOverlappingExecutionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetAutocommitParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetBinaryInputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetBinaryOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetClientMemoryLimitParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetClientMetadataRequestUseConnectionCtxParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetClientPrefetchThreadsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetClientResultChunkSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetClientResultColumnCaseInsensitiveParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetClientSessionKeepAliveParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetClientSessionKeepAliveHeartbeatFrequencyParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetClientTimestampTypeMappingParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetConfigParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetDateInputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetDateOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetEnableUnloadPhysicalTypeOptimizationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetErrorIntegrationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetErrorOnNondeterministicMergeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetErrorOnNondeterministicUpdateParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetFinalizeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetGeographyOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetGeometryOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetJdbcTreatTimestampNtzAsUtcParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetJdbcUseSessionTimezoneParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetJsonIndentParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetLockTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetLogLevelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetMultiStatementCountParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetNoorderSequenceAsDefaultParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetOdbcTreatDecimalAsIntParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetQueryTagParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetQuotedIdentifiersIgnoreCaseParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetRowsPerResultsetParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetS3StageVpceDnsNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSearchPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSqlStatementParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetStartedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetStatementQueuedTimeoutInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetStatementTimeoutInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetStrictJsonOutputParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetSuspendTaskAfterNumFailuresParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTaskAutoRetryAttemptsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimeInputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimeOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimestampDayIsAlways24HParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimestampInputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimestampLtzOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimestampNtzOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimestampOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimestampTypeMappingParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimestampTzOutputFormatParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTimezoneParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTraceLevelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTransactionAbortOnErrorParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTransactionDefaultIsolationLevelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetTwoDigitCenturyStartParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetUnsupportedDdlActionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetUseCachedResultParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetUserTaskManagedInitialWarehouseSizeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetUserTaskMinimumTriggerIntervalInSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetUserTaskTimeoutMsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetWarehouseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetWeekOfYearPolicyParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetWeekStartParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Task) validateSetWhenParameters(val *string) error {
	return nil
}

func validateNewTaskParameters(scope constructs.Construct, id *string, config *TaskConfig) error {
	return nil
}

