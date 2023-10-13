// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants


type DataSnowflakeGrantsGrantsTo struct {
	// Lists all privileges and roles granted to the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.73.0/docs/data-sources/grants#role DataSnowflakeGrants#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Lists all the privileges granted to the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.73.0/docs/data-sources/grants#share DataSnowflakeGrants#share}
	Share *string `field:"optional" json:"share" yaml:"share"`
	// Lists all the roles granted to the user.
	//
	// Note that the PUBLIC role, which is automatically available to every user, is not listed
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.73.0/docs/data-sources/grants#user DataSnowflakeGrants#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

