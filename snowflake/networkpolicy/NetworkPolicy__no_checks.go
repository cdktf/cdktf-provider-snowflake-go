// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPolicy) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateImportFromParameters(id *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateMoveToIdParameters(id *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validatePutTimeoutsParameters(value *NetworkPolicyTimeouts) error {
	return nil
}

func validateNetworkPolicy_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateNetworkPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkPolicy_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNetworkPolicy_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetAllowedIpListParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetAllowedNetworkRuleListParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetBlockedIpListParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetBlockedNetworkRuleListParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicy) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewNetworkPolicyParameters(scope constructs.Construct, id *string, config *NetworkPolicyConfig) error {
	return nil
}

