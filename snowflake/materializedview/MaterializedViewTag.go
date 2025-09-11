// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package materializedview


type MaterializedViewTag struct {
	// Tag name, e.g. department.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/materialized_view#name MaterializedView#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tag value, e.g. marketing_info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/materialized_view#value MaterializedView#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Name of the database that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/materialized_view#database MaterializedView#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the schema that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/materialized_view#schema MaterializedView#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

