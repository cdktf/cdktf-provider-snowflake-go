// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants


type DataSnowflakeGrantsFutureGrantsTo struct {
	// Lists all privileges on new (i.e. future) objects of a specified type in a database or schema granted to the account role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/data-sources/grants#account_role DataSnowflakeGrants#account_role}
	AccountRole *string `field:"optional" json:"accountRole" yaml:"accountRole"`
	// Lists all privileges on new (i.e. future) objects granted to the database role. Must be a fully qualified name ("&lt;db_name&gt;"."&lt;database_role_name&gt;").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/data-sources/grants#database_role DataSnowflakeGrants#database_role}
	DatabaseRole *string `field:"optional" json:"databaseRole" yaml:"databaseRole"`
}

