// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maskingpolicy


type MaskingPolicySignatureColumn struct {
	// Specifies the column name to mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/masking_policy#name MaskingPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the column type to mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/masking_policy#type MaskingPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

