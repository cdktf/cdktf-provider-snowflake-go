// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantownership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantOwnershipConfig struct {
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
	// on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/grant_ownership#on GrantOwnership#on}
	On *GrantOwnershipOn `field:"required" json:"on" yaml:"on"`
	// The fully qualified name of the account role to which privileges will be granted.
	//
	// For more information about this resource, see [docs](./account_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/grant_ownership#account_role_name GrantOwnership#account_role_name}
	AccountRoleName *string `field:"optional" json:"accountRoleName" yaml:"accountRoleName"`
	// The fully qualified name of the database role to which privileges will be granted.
	//
	// For more information about this resource, see [docs](./database_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/grant_ownership#database_role_name GrantOwnership#database_role_name}
	DatabaseRoleName *string `field:"optional" json:"databaseRoleName" yaml:"databaseRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/grant_ownership#id GrantOwnership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether to remove or transfer all existing outbound privileges on the object when ownership is transferred to a new role.
	//
	// Available options are: REVOKE for removing existing privileges and COPY to transfer them with ownership. For more information head over to [Snowflake documentation](https://docs.snowflake.com/en/sql-reference/sql/grant-ownership#optional-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/grant_ownership#outbound_privileges GrantOwnership#outbound_privileges}
	OutboundPrivileges *string `field:"optional" json:"outboundPrivileges" yaml:"outboundPrivileges"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/grant_ownership#timeouts GrantOwnership#timeouts}
	Timeouts *GrantOwnershipTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

