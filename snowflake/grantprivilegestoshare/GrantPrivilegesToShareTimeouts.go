// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoshare


type GrantPrivilegesToShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/grant_privileges_to_share#create GrantPrivilegesToShare#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/grant_privileges_to_share#delete GrantPrivilegesToShare#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/grant_privileges_to_share#read GrantPrivilegesToShare#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/grant_privileges_to_share#update GrantPrivilegesToShare#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

