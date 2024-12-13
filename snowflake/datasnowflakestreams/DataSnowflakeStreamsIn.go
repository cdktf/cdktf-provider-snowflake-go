// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakestreams


type DataSnowflakeStreamsIn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/data-sources/streams#account DataSnowflakeStreams#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the specified application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/data-sources/streams#application DataSnowflakeStreams#application}
	Application *string `field:"optional" json:"application" yaml:"application"`
	// Returns records for the specified application package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/data-sources/streams#application_package DataSnowflakeStreams#application_package}
	ApplicationPackage *string `field:"optional" json:"applicationPackage" yaml:"applicationPackage"`
	// Returns records for the current database in use or for a specified database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/data-sources/streams#database DataSnowflakeStreams#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Returns records for the current schema in use or a specified schema. Use fully qualified name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/data-sources/streams#schema DataSnowflakeStreams#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

