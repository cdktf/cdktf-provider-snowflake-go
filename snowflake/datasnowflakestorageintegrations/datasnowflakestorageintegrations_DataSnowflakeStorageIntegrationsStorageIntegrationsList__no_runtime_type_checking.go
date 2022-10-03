//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datasnowflakestorageintegrations

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeStorageIntegrationsStorageIntegrationsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeStorageIntegrationsStorageIntegrationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeStorageIntegrationsStorageIntegrationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeStorageIntegrationsStorageIntegrationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeStorageIntegrationsStorageIntegrationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeStorageIntegrationsStorageIntegrationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

