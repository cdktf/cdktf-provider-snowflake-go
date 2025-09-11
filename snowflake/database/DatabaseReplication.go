// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package database


type DatabaseReplication struct {
	// enable_to_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/database#enable_to_account Database#enable_to_account}
	EnableToAccount interface{} `field:"required" json:"enableToAccount" yaml:"enableToAccount"`
	// Allows replicating data to accounts on lower editions in either of the following scenarios: 1.
	//
	// The primary database is in a Business Critical (or higher) account but one or more of the accounts approved for replication are on lower editions. Business Critical Edition is intended for Snowflake accounts with extremely sensitive data. 2. The primary database is in a Business Critical (or higher) account and a signed business associate agreement is in place to store PHI data in the account per HIPAA and HITRUST regulations, but no such agreement is in place for one or more of the accounts approved for replication, regardless if they are Business Critical (or higher) accounts. Both scenarios are prohibited by default in an effort to help prevent account administrators for Business Critical (or higher) accounts from inadvertently replicating sensitive data to accounts on lower editions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/database#ignore_edition_check Database#ignore_edition_check}
	IgnoreEditionCheck interface{} `field:"optional" json:"ignoreEditionCheck" yaml:"ignoreEditionCheck"`
}

