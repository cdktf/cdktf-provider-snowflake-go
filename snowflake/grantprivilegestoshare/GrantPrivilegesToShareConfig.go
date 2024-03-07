// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantPrivilegesToShareConfig struct {
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
	// The privileges to grant on the share. See available list of privileges: https://docs.snowflake.com/en/sql-reference/sql/grant-privilege-share#syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#privileges GrantPrivilegesToShare#privileges}
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
	// The fully qualified name of the share on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#to_share GrantPrivilegesToShare#to_share}
	ToShare *string `field:"required" json:"toShare" yaml:"toShare"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#id GrantPrivilegesToShare#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The fully qualified identifier for the schema for which the specified privilege will be granted for all tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#on_all_tables_in_schema GrantPrivilegesToShare#on_all_tables_in_schema}
	OnAllTablesInSchema *string `field:"optional" json:"onAllTablesInSchema" yaml:"onAllTablesInSchema"`
	// The fully qualified name of the database on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#on_database GrantPrivilegesToShare#on_database}
	OnDatabase *string `field:"optional" json:"onDatabase" yaml:"onDatabase"`
	// The fully qualified name of the schema on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#on_schema GrantPrivilegesToShare#on_schema}
	OnSchema *string `field:"optional" json:"onSchema" yaml:"onSchema"`
	// The fully qualified name of the table on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#on_table GrantPrivilegesToShare#on_table}
	OnTable *string `field:"optional" json:"onTable" yaml:"onTable"`
	// The fully qualified name of the tag on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#on_tag GrantPrivilegesToShare#on_tag}
	OnTag *string `field:"optional" json:"onTag" yaml:"onTag"`
	// The fully qualified name of the view on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/grant_privileges_to_share#on_view GrantPrivilegesToShare#on_view}
	OnView *string `field:"optional" json:"onView" yaml:"onView"`
}

