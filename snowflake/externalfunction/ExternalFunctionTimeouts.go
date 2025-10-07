// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalfunction


type ExternalFunctionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/external_function#create ExternalFunction#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/external_function#delete ExternalFunction#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/external_function#read ExternalFunction#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/external_function#update ExternalFunction#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

