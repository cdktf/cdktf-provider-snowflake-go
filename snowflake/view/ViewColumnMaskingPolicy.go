// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view


type ViewColumnMaskingPolicy struct {
	// Specifies the masking policy to set on a column. For more information about this resource, see [docs](./masking_policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/view#policy_name View#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Specifies the arguments to pass into the conditional masking policy SQL expression.
	//
	// The first column in the list specifies the column for the policy conditions to mask or tokenize the data and must match the column to which the masking policy is set. The additional columns specify the columns to evaluate to determine whether to mask or tokenize the data in each row of the query result when a query is made on the first column. If the USING clause is omitted, Snowflake treats the conditional masking policy as a normal masking policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/view#using View#using}
	Using *[]*string `field:"optional" json:"using" yaml:"using"`
}

