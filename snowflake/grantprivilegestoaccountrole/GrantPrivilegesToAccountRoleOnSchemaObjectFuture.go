// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole


type GrantPrivilegesToAccountRoleOnSchemaObjectFuture struct {
	// The plural object type of the schema object on which privileges will be granted.
	//
	// Valid values are: ALERTS | AUTHENTICATION POLICIES | CORTEX SEARCH SERVICES | DATA METRIC FUNCTIONS | DATASETS | DBT PROJECTS | DYNAMIC TABLES | EVENT TABLES | EXTERNAL TABLES | FILE FORMATS | FUNCTIONS | GIT REPOSITORIES | HYBRID TABLES | ICEBERG TABLES | JOIN POLICIES | MATERIALIZED VIEWS | MODELS | MODEL MONITORS | NETWORK RULES | NOTEBOOKS | PASSWORD POLICIES | PIPES | PRIVACY POLICIES | PROCEDURES | SECRETS | SEMANTIC VIEWS | SERVICES | SEQUENCES | SNAPSHOTS | SNAPSHOT POLICIES | SNAPSHOT SETS | STAGES | STREAMS | STREAMLITS | TABLES | TASKS | VIEWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/grant_privileges_to_account_role#object_type_plural GrantPrivilegesToAccountRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/grant_privileges_to_account_role#in_database GrantPrivilegesToAccountRole#in_database}.
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/grant_privileges_to_account_role#in_schema GrantPrivilegesToAccountRole#in_schema}.
	InSchema *string `field:"optional" json:"inSchema" yaml:"inSchema"`
}

