// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestorole


type GrantPrivilegesToRoleOnSchemaObject struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/grant_privileges_to_role#all GrantPrivilegesToRole#all}
	All *GrantPrivilegesToRoleOnSchemaObjectAll `field:"optional" json:"all" yaml:"all"`
	// future block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/grant_privileges_to_role#future GrantPrivilegesToRole#future}
	Future *GrantPrivilegesToRoleOnSchemaObjectFuture `field:"optional" json:"future" yaml:"future"`
	// The fully qualified name of the object on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/grant_privileges_to_role#object_name GrantPrivilegesToRole#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// The object type of the schema object on which privileges will be granted.
	//
	// Valid values are: AGGREGATION POLICY | ALERT | AUTHENTICATION POLICY | DATA METRIC FUNCTION | DYNAMIC TABLE | EVENT TABLE | EXTERNAL TABLE | FILE FORMAT | FUNCTION | GIT REPOSITORY | HYBRID TABLE | IMAGE REPOSITORY | ICEBERG TABLE | MASKING POLICY | MATERIALIZED VIEW | MODEL | NETWORK RULE | PACKAGES POLICY | PASSWORD POLICY | PIPE | PROCEDURE | PROJECTION POLICY | ROW ACCESS POLICY | SECRET | SERVICE | SESSION POLICY | SEQUENCE | STAGE | STREAM | TABLE | TAG | TASK | VIEW | STREAMLIT
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/grant_privileges_to_role#object_type GrantPrivilegesToRole#object_type}
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
}

