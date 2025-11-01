// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeauthenticationpolicies


type DataSnowflakeAuthenticationPoliciesIn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/data-sources/authentication_policies#account DataSnowflakeAuthenticationPolicies#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the specified application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/data-sources/authentication_policies#application DataSnowflakeAuthenticationPolicies#application}
	Application *string `field:"optional" json:"application" yaml:"application"`
	// Returns records for the specified application package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/data-sources/authentication_policies#application_package DataSnowflakeAuthenticationPolicies#application_package}
	ApplicationPackage *string `field:"optional" json:"applicationPackage" yaml:"applicationPackage"`
	// Returns records for the current database in use or for a specified database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/data-sources/authentication_policies#database DataSnowflakeAuthenticationPolicies#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Returns records for the current schema in use or a specified schema. Use fully qualified name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/data-sources/authentication_policies#schema DataSnowflakeAuthenticationPolicies#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

