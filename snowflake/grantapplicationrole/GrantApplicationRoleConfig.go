// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantapplicationrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantApplicationRoleConfig struct {
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
	// Specifies the identifier for the application role to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_application_role#application_role_name GrantApplicationRole#application_role_name}
	ApplicationRoleName *string `field:"required" json:"applicationRoleName" yaml:"applicationRoleName"`
	// The fully qualified name of the application on which application role will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_application_role#application_name GrantApplicationRole#application_name}
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_application_role#id GrantApplicationRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The fully qualified name of the account role on which application role will be granted.
	//
	// For more information about this resource, see [docs](./account_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/grant_application_role#parent_account_role_name GrantApplicationRole#parent_account_role_name}
	ParentAccountRoleName *string `field:"optional" json:"parentAccountRoleName" yaml:"parentAccountRoleName"`
}

