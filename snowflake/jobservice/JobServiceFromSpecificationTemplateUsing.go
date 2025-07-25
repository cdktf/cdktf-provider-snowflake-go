// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobservice


type JobServiceFromSpecificationTemplateUsing struct {
	// The name of the template variable.
	//
	// The provider wraps it in double quotes by default, so be aware of that while referencing the argument in the spec definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/job_service#key JobService#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value to assign to the variable in the template.
	//
	// The provider wraps it in `$$` by default, so be aware of that while referencing the argument in the spec definition. The value must either be alphanumeric or valid JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/job_service#value JobService#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

