// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestorole


type GrantPrivilegesToRoleOnSchemaObjectFuture struct {
	// The plural object type of the schema object on which privileges will be granted.
	//
	// Valid values are: ALERTS | DYNAMIC TABLES | EVENT TABLES | FILE FORMATS | FUNCTIONS | ICEBERG TABLES | PROCEDURES | SECRETS | SEQUENCES | PIPES | MASKING POLICIES | PASSWORD POLICIES | ROW ACCESS POLICIES | SESSION POLICIES | TAGS | STAGES | STREAMS | TABLES | EXTERNAL TABLES | TASKS | VIEWS | MATERIALIZED VIEWS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.86.0/docs/resources/grant_privileges_to_role#object_type_plural GrantPrivilegesToRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
	// The fully qualified name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.86.0/docs/resources/grant_privileges_to_role#in_database GrantPrivilegesToRole#in_database}
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
	// The fully qualified name of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.86.0/docs/resources/grant_privileges_to_role#in_schema GrantPrivilegesToRole#in_schema}
	InSchema *string `field:"optional" json:"inSchema" yaml:"inSchema"`
}

