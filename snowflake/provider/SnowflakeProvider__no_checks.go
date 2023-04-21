//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnowflakeProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SnowflakeProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
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

func (j *jsiiProxy_SnowflakeProvider) validateSetInsecureModeParameters(val interface{}) error {
	return nil
}

func validateNewSnowflakeProviderParameters(scope constructs.Construct, id *string, config *SnowflakeProviderConfig) error {
	return nil
}

