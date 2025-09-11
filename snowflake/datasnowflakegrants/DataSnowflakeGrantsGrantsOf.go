// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants


type DataSnowflakeGrantsGrantsOf struct {
	// Lists all users and roles to which the account role has been granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/data-sources/grants#account_role DataSnowflakeGrants#account_role}
	AccountRole *string `field:"optional" json:"accountRole" yaml:"accountRole"`
	// Lists all the users and roles to which the application role has been granted.
	//
	// Must be a fully qualified name ("&lt;db_name&gt;"."&lt;database_role_name&gt;").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/data-sources/grants#application_role DataSnowflakeGrants#application_role}
	ApplicationRole *string `field:"optional" json:"applicationRole" yaml:"applicationRole"`
	// Lists all users and roles to which the database role has been granted.
	//
	// Must be a fully qualified name ("&lt;db_name&gt;"."&lt;database_role_name&gt;").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/data-sources/grants#database_role DataSnowflakeGrants#database_role}
	DatabaseRole *string `field:"optional" json:"databaseRole" yaml:"databaseRole"`
	// Lists all the accounts for the share and indicates the accounts that are using the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/data-sources/grants#share DataSnowflakeGrants#share}
	Share *string `field:"optional" json:"share" yaml:"share"`
}

