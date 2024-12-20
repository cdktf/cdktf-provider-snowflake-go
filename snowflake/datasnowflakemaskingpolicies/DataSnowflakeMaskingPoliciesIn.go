// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakemaskingpolicies


type DataSnowflakeMaskingPoliciesIn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/data-sources/masking_policies#account DataSnowflakeMaskingPolicies#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the specified application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/data-sources/masking_policies#application DataSnowflakeMaskingPolicies#application}
	Application *string `field:"optional" json:"application" yaml:"application"`
	// Returns records for the specified application package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/data-sources/masking_policies#application_package DataSnowflakeMaskingPolicies#application_package}
	ApplicationPackage *string `field:"optional" json:"applicationPackage" yaml:"applicationPackage"`
	// Returns records for the current database in use or for a specified database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/data-sources/masking_policies#database DataSnowflakeMaskingPolicies#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Returns records for the current schema in use or a specified schema. Use fully qualified name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/data-sources/masking_policies#schema DataSnowflakeMaskingPolicies#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

