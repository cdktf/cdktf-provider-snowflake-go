// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantPrivilegesToAccountRoleConfig struct {
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
	// The fully qualified name of the account role to which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#account_role_name GrantPrivilegesToAccountRole#account_role_name}
	AccountRoleName *string `field:"required" json:"accountRoleName" yaml:"accountRoleName"`
	// Grant all privileges on the account role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#all_privileges GrantPrivilegesToAccountRole#all_privileges}
	AllPrivileges interface{} `field:"optional" json:"allPrivileges" yaml:"allPrivileges"`
	// If true, the resource will always produce a “plan” and on “apply” it will re-grant defined privileges.
	//
	// It is supposed to be used only in “grant privileges on all X’s in database / schema Y” or “grant all privileges to X” scenarios to make sure that every new object in a given database / schema is granted by the account role and every new privilege is granted to the database role. Important note: this flag is not compliant with the Terraform assumptions of the config being eventually convergent (producing an empty plan).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#always_apply GrantPrivilegesToAccountRole#always_apply}
	AlwaysApply interface{} `field:"optional" json:"alwaysApply" yaml:"alwaysApply"`
	// This is a helper field and should not be set.
	//
	// Its main purpose is to help to achieve the functionality described by the always_apply field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#always_apply_trigger GrantPrivilegesToAccountRole#always_apply_trigger}
	AlwaysApplyTrigger *string `field:"optional" json:"alwaysApplyTrigger" yaml:"alwaysApplyTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#id GrantPrivilegesToAccountRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, the privileges will be granted on the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#on_account GrantPrivilegesToAccountRole#on_account}
	OnAccount interface{} `field:"optional" json:"onAccount" yaml:"onAccount"`
	// on_account_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#on_account_object GrantPrivilegesToAccountRole#on_account_object}
	OnAccountObject *GrantPrivilegesToAccountRoleOnAccountObject `field:"optional" json:"onAccountObject" yaml:"onAccountObject"`
	// on_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#on_schema GrantPrivilegesToAccountRole#on_schema}
	OnSchema *GrantPrivilegesToAccountRoleOnSchema `field:"optional" json:"onSchema" yaml:"onSchema"`
	// on_schema_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#on_schema_object GrantPrivilegesToAccountRole#on_schema_object}
	OnSchemaObject *GrantPrivilegesToAccountRoleOnSchemaObject `field:"optional" json:"onSchemaObject" yaml:"onSchemaObject"`
	// The privileges to grant on the account role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#privileges GrantPrivilegesToAccountRole#privileges}
	Privileges *[]*string `field:"optional" json:"privileges" yaml:"privileges"`
	// Specifies whether the grantee can grant the privileges to other users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_account_role#with_grant_option GrantPrivilegesToAccountRole#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

