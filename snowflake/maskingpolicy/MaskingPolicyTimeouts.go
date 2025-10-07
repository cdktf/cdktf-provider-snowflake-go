// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maskingpolicy


type MaskingPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/masking_policy#create MaskingPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/masking_policy#delete MaskingPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/masking_policy#read MaskingPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/masking_policy#update MaskingPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

