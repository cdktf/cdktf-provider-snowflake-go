// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view


type ViewAggregationPolicy struct {
	// Aggregation policy name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/view#policy_name View#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Defines which columns uniquely identify an entity within the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/view#entity_key View#entity_key}
	EntityKey *[]*string `field:"optional" json:"entityKey" yaml:"entityKey"`
}

