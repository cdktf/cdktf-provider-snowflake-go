// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestorole


type GrantPrivilegesToRoleOnSchema struct {
	// The fully qualified name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_role#all_schemas_in_database GrantPrivilegesToRole#all_schemas_in_database}
	AllSchemasInDatabase *string `field:"optional" json:"allSchemasInDatabase" yaml:"allSchemasInDatabase"`
	// The fully qualified name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_role#future_schemas_in_database GrantPrivilegesToRole#future_schemas_in_database}
	FutureSchemasInDatabase *string `field:"optional" json:"futureSchemasInDatabase" yaml:"futureSchemasInDatabase"`
	// The fully qualified name of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_role#schema_name GrantPrivilegesToRole#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

