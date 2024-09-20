// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants


type DataSnowflakeGrantsFutureGrantsIn struct {
	// Lists all privileges on new (i.e. future) objects of a specified type in the database granted to a role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/data-sources/grants#database DataSnowflakeGrants#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Lists all privileges on new (i.e. future) objects of a specified type in the schema granted to a role. Schema must be a fully qualified name ("&lt;db_name&gt;"."&lt;schema_name&gt;").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/data-sources/grants#schema DataSnowflakeGrants#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

