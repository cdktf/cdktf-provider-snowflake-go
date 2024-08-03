// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package database


type DatabaseReplicationEnableToAccount struct {
	// Specifies account identifier for which replication should be enabled. The account identifiers should be in the form of `"<organization_name>"."<account_name>"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/resources/database#account_identifier Database#account_identifier}
	AccountIdentifier *string `field:"required" json:"accountIdentifier" yaml:"accountIdentifier"`
	// Specifies if failover should be enabled for the specified account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.1/docs/resources/database#with_failover Database#with_failover}
	WithFailover interface{} `field:"optional" json:"withFailover" yaml:"withFailover"`
}

