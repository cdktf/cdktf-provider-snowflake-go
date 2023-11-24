// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnowflakeProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SnowflakeProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSnowflakeProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSnowflakeProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSnowflakeProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSnowflakeProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetBrowserAuthParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetClientRequestMfaTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetClientStoreTemporaryCredentialParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetDisableQueryContextCacheParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetDisableTelemetryParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetInsecureModeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetKeepSessionAliveParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetOcspFailOpenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetPasscodeInPasswordParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetTokenAccessorParameters(val *SnowflakeProviderTokenAccessor) error {
	return nil
}

func (j *jsiiProxy_SnowflakeProvider) validateSetValidateDefaultParametersParameters(val interface{}) error {
	return nil
}

func validateNewSnowflakeProviderParameters(scope constructs.Construct, id *string, config *SnowflakeProviderConfig) error {
	return nil
}

