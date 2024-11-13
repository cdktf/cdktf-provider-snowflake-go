// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakestreamlits


type DataSnowflakeStreamlitsIn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/data-sources/streamlits#account DataSnowflakeStreamlits#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the current database in use or for a specified database (db_name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/data-sources/streamlits#database DataSnowflakeStreamlits#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Returns records for the current schema in use or a specified schema (schema_name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/data-sources/streamlits#schema DataSnowflakeStreamlits#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

