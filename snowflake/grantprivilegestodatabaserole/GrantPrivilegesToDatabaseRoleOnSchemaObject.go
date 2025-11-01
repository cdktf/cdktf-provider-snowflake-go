// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestodatabaserole


type GrantPrivilegesToDatabaseRoleOnSchemaObject struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/grant_privileges_to_database_role#all GrantPrivilegesToDatabaseRole#all}
	All *GrantPrivilegesToDatabaseRoleOnSchemaObjectAll `field:"optional" json:"all" yaml:"all"`
	// future block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/grant_privileges_to_database_role#future GrantPrivilegesToDatabaseRole#future}
	Future *GrantPrivilegesToDatabaseRoleOnSchemaObjectFuture `field:"optional" json:"future" yaml:"future"`
	// The fully qualified name of the object on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/grant_privileges_to_database_role#object_name GrantPrivilegesToDatabaseRole#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// The object type of the schema object on which privileges will be granted.
	//
	// Valid values are: AGGREGATION POLICY | ALERT | AUTHENTICATION POLICY | CORTEX SEARCH SERVICE | DATA METRIC FUNCTION | DATASET | DBT PROJECT | DYNAMIC TABLE | EVENT TABLE | EXTERNAL TABLE | FILE FORMAT | FUNCTION | GIT REPOSITORY | HYBRID TABLE | IMAGE REPOSITORY | ICEBERG TABLE | JOIN POLICY | MASKING POLICY | MATERIALIZED VIEW | MODEL | MODEL MONITOR | NETWORK RULE | NOTEBOOK | PACKAGES POLICY | PASSWORD POLICY | PIPE | PRIVACY POLICY | PROCEDURE | PROJECTION POLICY | ROW ACCESS POLICY | SECRET | SEMANTIC VIEW | SERVICE | SESSION POLICY | SEQUENCE | SNAPSHOT | SNAPSHOT POLICY | SNAPSHOT SET | STAGE | STREAM | STREAMLIT | TABLE | TAG | TASK | VIEW
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/grant_privileges_to_database_role#object_type GrantPrivilegesToDatabaseRole#object_type}
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
}

