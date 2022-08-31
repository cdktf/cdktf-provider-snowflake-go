//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

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

func (j *jsiiProxy_SnowflakeProvider) validateSetBrowserAuthParameters(val interface{}) error {
	return nil
}

func validateNewSnowflakeProviderParameters(scope constructs.Construct, id *string, config *SnowflakeProviderConfig) error {
	return nil
}

