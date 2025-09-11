// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakecortexsearchservices


type DataSnowflakeCortexSearchServicesIn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/data-sources/cortex_search_services#account DataSnowflakeCortexSearchServices#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the current database in use or for a specified database (db_name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/data-sources/cortex_search_services#database DataSnowflakeCortexSearchServices#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Returns records for the current schema in use or a specified schema (schema_name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/data-sources/cortex_search_services#schema DataSnowflakeCortexSearchServices#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

