// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package materializedview

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaterializedViewTagList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MaterializedViewTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MaterializedViewTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaterializedViewTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaterializedViewTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MaterializedViewTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMaterializedViewTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

