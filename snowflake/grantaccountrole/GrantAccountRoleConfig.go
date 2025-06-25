// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantaccountrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantAccountRoleConfig struct {
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
	// The fully qualified name of the role which will be granted to the user or parent role.
	//
	// For more information about this resource, see [docs](./account_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/grant_account_role#role_name GrantAccountRole#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/grant_account_role#id GrantAccountRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The fully qualified name of the parent role which will create a parent-child relationship between the roles.
	//
	// For more information about this resource, see [docs](./account_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/grant_account_role#parent_role_name GrantAccountRole#parent_role_name}
	ParentRoleName *string `field:"optional" json:"parentRoleName" yaml:"parentRoleName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/grant_account_role#timeouts GrantAccountRole#timeouts}
	Timeouts *GrantAccountRoleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The fully qualified name of the user on which specified role will be granted.
	//
	// For more information about this resource, see [docs](./user).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/grant_account_role#user_name GrantAccountRole#user_name}
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

