// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalfunction


type ExternalFunctionArg struct {
	// Argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.91.0/docs/resources/external_function#name ExternalFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Argument type, e.g. VARCHAR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.91.0/docs/resources/external_function#type ExternalFunction#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

