// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants


type DataSnowflakeGrantsGrantsToShare struct {
	// Lists all of the privileges and roles granted to the specified share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/data-sources/grants#share_name DataSnowflakeGrants#share_name}
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
}

