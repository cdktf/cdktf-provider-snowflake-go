// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externaltable


type ExternalTableTag struct {
	// Tag name, e.g. department.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/external_table#name ExternalTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tag value, e.g. marketing_info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/external_table#value ExternalTable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Name of the database that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/external_table#database ExternalTable#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the schema that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/external_table#schema ExternalTable#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

