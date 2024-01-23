// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedure


type ProcedureArguments struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.1/docs/resources/procedure#name Procedure#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.1/docs/resources/procedure#type Procedure#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

