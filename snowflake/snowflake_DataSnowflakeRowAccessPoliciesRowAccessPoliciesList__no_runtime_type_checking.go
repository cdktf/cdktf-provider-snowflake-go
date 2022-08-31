//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeRowAccessPoliciesRowAccessPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeRowAccessPoliciesRowAccessPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeRowAccessPoliciesRowAccessPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeRowAccessPoliciesRowAccessPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeRowAccessPoliciesRowAccessPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeRowAccessPoliciesRowAccessPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

