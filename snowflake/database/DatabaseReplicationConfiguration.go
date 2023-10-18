// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package database


type DatabaseReplicationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/database#accounts Database#accounts}.
	Accounts *[]*string `field:"required" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.74.0/docs/resources/database#ignore_edition_check Database#ignore_edition_check}.
	IgnoreEditionCheck interface{} `field:"optional" json:"ignoreEditionCheck" yaml:"ignoreEditionCheck"`
}

