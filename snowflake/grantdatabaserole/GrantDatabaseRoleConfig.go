// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantdatabaserole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantDatabaseRoleConfig struct {
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
	// The fully qualified name of the database role which will be granted to share or parent role.
	//
	// For more information about this resource, see [docs](./database_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_database_role#database_role_name GrantDatabaseRole#database_role_name}
	DatabaseRoleName *string `field:"required" json:"databaseRoleName" yaml:"databaseRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_database_role#id GrantDatabaseRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The fully qualified name of the parent database role which will create a parent-child relationship between the roles.
	//
	// For more information about this resource, see [docs](./database_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_database_role#parent_database_role_name GrantDatabaseRole#parent_database_role_name}
	ParentDatabaseRoleName *string `field:"optional" json:"parentDatabaseRoleName" yaml:"parentDatabaseRoleName"`
	// The fully qualified name of the parent account role which will create a parent-child relationship between the roles.
	//
	// For more information about this resource, see [docs](./account_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_database_role#parent_role_name GrantDatabaseRole#parent_role_name}
	ParentRoleName *string `field:"optional" json:"parentRoleName" yaml:"parentRoleName"`
	// The fully qualified name of the share on which privileges will be granted.
	//
	// For more information about this resource, see [docs](./share).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_database_role#share_name GrantDatabaseRole#share_name}
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
}

