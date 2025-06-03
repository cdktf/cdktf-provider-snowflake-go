// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestodatabaserole


type GrantPrivilegesToDatabaseRoleOnSchemaObjectAll struct {
	// The plural object type of the schema object on which privileges will be granted.
	//
	// Valid values are: AGGREGATION POLICIES | ALERTS | AUTHENTICATION POLICIES | CORTEX SEARCH SERVICES | DATA METRIC FUNCTIONS | DYNAMIC TABLES | EVENT TABLES | EXTERNAL TABLES | FILE FORMATS | FUNCTIONS | GIT REPOSITORIES | HYBRID TABLES | IMAGE REPOSITORIES | ICEBERG TABLES | MASKING POLICIES | MATERIALIZED VIEWS | MODELS | NETWORK RULES | NOTEBOOKS | PACKAGES POLICIES | PASSWORD POLICIES | PIPES | PROCEDURES | PROJECTION POLICIES | ROW ACCESS POLICIES | SECRETS | SERVICES | SESSION POLICIES | SEQUENCES | SNAPSHOTS | STAGES | STREAMS | TABLES | TAGS | TASKS | VIEWS | STREAMLITS | DATASETS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/grant_privileges_to_database_role#object_type_plural GrantPrivilegesToDatabaseRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
	// The fully qualified name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/grant_privileges_to_database_role#in_database GrantPrivilegesToDatabaseRole#in_database}
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
	// The fully qualified name of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/grant_privileges_to_database_role#in_schema GrantPrivilegesToDatabaseRole#in_schema}
	InSchema *string `field:"optional" json:"inSchema" yaml:"inSchema"`
}

