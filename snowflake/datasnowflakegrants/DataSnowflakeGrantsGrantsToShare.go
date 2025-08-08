// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants


type DataSnowflakeGrantsGrantsToShare struct {
	// Lists all of the privileges and roles granted to the specified share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/data-sources/grants#share_name DataSnowflakeGrants#share_name}
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
}

