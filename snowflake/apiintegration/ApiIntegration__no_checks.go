// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apiintegration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiIntegration) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (a *jsiiProxy_ApiIntegration) validatePutTimeoutsParameters(value *ApiIntegrationTimeouts) error {
	return nil
}

func validateApiIntegration_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateApiIntegration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApiIntegration_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateApiIntegration_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetApiAllowedPrefixesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetApiAwsRoleArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetApiBlockedPrefixesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetApiGcpServiceAccountParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetApiKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetApiProviderParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetAzureAdApplicationIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetAzureTenantIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetGoogleAudienceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiIntegration) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewApiIntegrationParameters(scope constructs.Construct, id *string, config *ApiIntegrationConfig) error {
	return nil
}

