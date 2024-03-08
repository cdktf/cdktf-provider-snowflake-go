// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole


type GrantPrivilegesToAccountRoleOnSchemaObject struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_account_role#all GrantPrivilegesToAccountRole#all}
	All *GrantPrivilegesToAccountRoleOnSchemaObjectAll `field:"optional" json:"all" yaml:"all"`
	// future block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_account_role#future GrantPrivilegesToAccountRole#future}
	Future *GrantPrivilegesToAccountRoleOnSchemaObjectFuture `field:"optional" json:"future" yaml:"future"`
	// The fully qualified name of the object on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_account_role#object_name GrantPrivilegesToAccountRole#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// The object type of the schema object on which privileges will be granted.
	//
	// Valid values are: ALERT | DYNAMIC TABLE | EVENT TABLE | FILE FORMAT | FUNCTION | PROCEDURE | SECRET | SEQUENCE | PIPE | MASKING POLICY | PASSWORD POLICY | ROW ACCESS POLICY | SESSION POLICY | TAG | STAGE | STREAM | TABLE | EXTERNAL TABLE | TASK | VIEW | MATERIALIZED VIEW | NETWORK RULE | PACKAGES POLICY | ICEBERG TABLE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.2/docs/resources/grant_privileges_to_account_role#object_type GrantPrivilegesToAccountRole#object_type}
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
}

