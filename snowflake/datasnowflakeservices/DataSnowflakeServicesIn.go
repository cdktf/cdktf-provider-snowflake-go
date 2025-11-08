// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeservices


type DataSnowflakeServicesIn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/data-sources/services#account DataSnowflakeServices#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the specified compute pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/data-sources/services#compute_pool DataSnowflakeServices#compute_pool}
	ComputePool *string `field:"optional" json:"computePool" yaml:"computePool"`
	// Returns records for the current database in use or for a specified database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/data-sources/services#database DataSnowflakeServices#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Returns records for the current schema in use or a specified schema. Use fully qualified name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/data-sources/services#schema DataSnowflakeServices#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

