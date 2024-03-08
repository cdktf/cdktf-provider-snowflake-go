// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestodatabaserole


type GrantPrivilegesToDatabaseRoleOnSchemaObjectAll struct {
	// The plural object type of the schema object on which privileges will be granted.
	//
	// Valid values are: ALERTS | DYNAMIC TABLES | EVENT TABLES | FILE FORMATS | FUNCTIONS | PROCEDURES | SECRETS | SEQUENCES | PIPES | MASKING POLICIES | PASSWORD POLICIES | ROW ACCESS POLICIES | SESSION POLICIES | TAGS | STAGES | STREAMS | TABLES | EXTERNAL TABLES | TASKS | VIEWS | MATERIALIZED VIEWS | NETWORK RULES | PACKAGES POLICIES | ICEBERG TABLES
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_database_role#object_type_plural GrantPrivilegesToDatabaseRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_database_role#in_database GrantPrivilegesToDatabaseRole#in_database}.
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_database_role#in_schema GrantPrivilegesToDatabaseRole#in_schema}.
	InSchema *string `field:"optional" json:"inSchema" yaml:"inSchema"`
}

