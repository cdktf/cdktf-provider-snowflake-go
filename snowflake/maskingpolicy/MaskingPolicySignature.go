// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maskingpolicy


type MaskingPolicySignature struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/masking_policy#column MaskingPolicy#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
}

