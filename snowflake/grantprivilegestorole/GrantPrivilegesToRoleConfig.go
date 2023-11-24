// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestorole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantPrivilegesToRoleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The fully qualified name of the role to which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#role_name GrantPrivilegesToRole#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Grant all privileges on the account role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#all_privileges GrantPrivilegesToRole#all_privileges}
	AllPrivileges interface{} `field:"optional" json:"allPrivileges" yaml:"allPrivileges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#id GrantPrivilegesToRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, the privileges will be granted on the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#on_account GrantPrivilegesToRole#on_account}
	OnAccount interface{} `field:"optional" json:"onAccount" yaml:"onAccount"`
	// on_account_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#on_account_object GrantPrivilegesToRole#on_account_object}
	OnAccountObject *GrantPrivilegesToRoleOnAccountObject `field:"optional" json:"onAccountObject" yaml:"onAccountObject"`
	// on_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#on_schema GrantPrivilegesToRole#on_schema}
	OnSchema *GrantPrivilegesToRoleOnSchema `field:"optional" json:"onSchema" yaml:"onSchema"`
	// on_schema_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#on_schema_object GrantPrivilegesToRole#on_schema_object}
	OnSchemaObject *GrantPrivilegesToRoleOnSchemaObject `field:"optional" json:"onSchemaObject" yaml:"onSchemaObject"`
	// The privileges to grant on the account role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#privileges GrantPrivilegesToRole#privileges}
	Privileges *[]*string `field:"optional" json:"privileges" yaml:"privileges"`
	// Specifies whether the grantee can grant the privileges to other users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.76.0/docs/resources/grant_privileges_to_role#with_grant_option GrantPrivilegesToRole#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

