// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalfunction


type ExternalFunctionHeader struct {
	// Header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/external_function#name ExternalFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/external_function#value ExternalFunction#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

