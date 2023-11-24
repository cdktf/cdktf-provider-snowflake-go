// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view


type ViewTag struct {
	// Tag name, e.g. department.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/view#name View#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tag value, e.g. marketing_info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/view#value View#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Name of the database that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/view#database View#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the schema that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/view#schema View#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

