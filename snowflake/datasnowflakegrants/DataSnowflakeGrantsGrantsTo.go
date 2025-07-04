// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants


type DataSnowflakeGrantsGrantsTo struct {
	// Lists all privileges and roles granted to the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/grants#account_role DataSnowflakeGrants#account_role}
	AccountRole *string `field:"optional" json:"accountRole" yaml:"accountRole"`
	// Lists all the privileges and roles granted to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/grants#application DataSnowflakeGrants#application}
	Application *string `field:"optional" json:"application" yaml:"application"`
	// Lists all the privileges and roles granted to the application role. Must be a fully qualified name ("&lt;app_name&gt;"."&lt;app_role_name&gt;").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/grants#application_role DataSnowflakeGrants#application_role}
	ApplicationRole *string `field:"optional" json:"applicationRole" yaml:"applicationRole"`
	// Lists all privileges and roles granted to the database role. Must be a fully qualified name ("&lt;db_name&gt;"."&lt;database_role_name&gt;").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/grants#database_role DataSnowflakeGrants#database_role}
	DatabaseRole *string `field:"optional" json:"databaseRole" yaml:"databaseRole"`
	// share block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/grants#share DataSnowflakeGrants#share}
	Share *DataSnowflakeGrantsGrantsToShare `field:"optional" json:"share" yaml:"share"`
	// Lists all the roles granted to the user.
	//
	// Note that the PUBLIC role, which is automatically available to every user, is not listed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/data-sources/grants#user DataSnowflakeGrants#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

