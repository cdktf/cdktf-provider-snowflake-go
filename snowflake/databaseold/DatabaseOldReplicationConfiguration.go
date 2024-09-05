// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseold


type DatabaseOldReplicationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/database_old#accounts DatabaseOld#accounts}.
	Accounts *[]*string `field:"required" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/database_old#ignore_edition_check DatabaseOld#ignore_edition_check}.
	IgnoreEditionCheck interface{} `field:"optional" json:"ignoreEditionCheck" yaml:"ignoreEditionCheck"`
}

