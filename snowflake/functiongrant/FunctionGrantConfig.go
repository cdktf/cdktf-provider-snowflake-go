// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functiongrant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionGrantConfig struct {
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
	// The name of the database containing the current or future functions on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#database_name FunctionGrant#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Grants privilege to these roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#roles FunctionGrant#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// List of the argument data types for the function (must be present if function has arguments and function_name is present).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#argument_data_types FunctionGrant#argument_data_types}
	ArgumentDataTypes *[]*string `field:"optional" json:"argumentDataTypes" yaml:"argumentDataTypes"`
	// When this is set to true, multiple grants of the same type can be created.
	//
	// This will cause Terraform to not revoke grants applied to roles and objects outside Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#enable_multiple_grants FunctionGrant#enable_multiple_grants}
	EnableMultipleGrants interface{} `field:"optional" json:"enableMultipleGrants" yaml:"enableMultipleGrants"`
	// The name of the function on which to grant privileges immediately (only valid if on_future is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#function_name FunctionGrant#function_name}
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#id FunctionGrant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When this is set to true and a schema_name is provided, apply this grant on all functions in the given schema.
	//
	// When this is true and no schema_name is provided apply this grant on all functions in the given database. The function_name, arguments, return_type, and shares fields must be unset in order to use on_all. Cannot be used together with on_future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#on_all FunctionGrant#on_all}
	OnAll interface{} `field:"optional" json:"onAll" yaml:"onAll"`
	// When this is set to true and a schema_name is provided, apply this grant on all future functions in the given schema.
	//
	// When this is true and no schema_name is provided apply this grant on all future functions in the given database. The function_name, arguments, return_type, and shares fields must be unset in order to use on_future. Cannot be used together with on_all.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#on_future FunctionGrant#on_future}
	OnFuture interface{} `field:"optional" json:"onFuture" yaml:"onFuture"`
	// The privilege to grant on the current or future function.
	//
	// Must be one of `USAGE` or `OWNERSHIP`. To grant all privileges, use the value `ALL PRIVILEGES`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#privilege FunctionGrant#privilege}
	Privilege *string `field:"optional" json:"privilege" yaml:"privilege"`
	// The name of the role to revert ownership to on destroy.
	//
	// Has no effect unless `privilege` is set to `OWNERSHIP`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#revert_ownership_to_role_name FunctionGrant#revert_ownership_to_role_name}
	RevertOwnershipToRoleName *string `field:"optional" json:"revertOwnershipToRoleName" yaml:"revertOwnershipToRoleName"`
	// The name of the schema containing the current or future functions on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#schema_name FunctionGrant#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Grants privilege to these shares (only valid if on_future is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#shares FunctionGrant#shares}
	Shares *[]*string `field:"optional" json:"shares" yaml:"shares"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/function_grant#with_grant_option FunctionGrant#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

