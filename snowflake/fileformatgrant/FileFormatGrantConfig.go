// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fileformatgrant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FileFormatGrantConfig struct {
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
	// The name of the database containing the current or future file formats on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#database_name FileFormatGrant#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Grants privilege to these roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#roles FileFormatGrant#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// When this is set to true, multiple grants of the same type can be created.
	//
	// This will cause Terraform to not revoke grants applied to roles and objects outside Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#enable_multiple_grants FileFormatGrant#enable_multiple_grants}
	EnableMultipleGrants interface{} `field:"optional" json:"enableMultipleGrants" yaml:"enableMultipleGrants"`
	// The name of the file format on which to grant privileges immediately (only valid if on_future is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#file_format_name FileFormatGrant#file_format_name}
	FileFormatName *string `field:"optional" json:"fileFormatName" yaml:"fileFormatName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#id FileFormatGrant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When this is set to true and a schema_name is provided, apply this grant on all file formats in the given schema.
	//
	// When this is true and no schema_name is provided apply this grant on all file formats in the given database. The file_format_name field must be unset in order to use on_all. Cannot be used together with on_future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#on_all FileFormatGrant#on_all}
	OnAll interface{} `field:"optional" json:"onAll" yaml:"onAll"`
	// When this is set to true and a schema_name is provided, apply this grant on all future file formats in the given schema.
	//
	// When this is true and no schema_name is provided apply this grant on all future file formats in the given database. The file_format_name field must be unset in order to use on_future. Cannot be used together with on_all.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#on_future FileFormatGrant#on_future}
	OnFuture interface{} `field:"optional" json:"onFuture" yaml:"onFuture"`
	// The privilege to grant on the current or future file format.
	//
	// To grant all privileges, use the value `ALL PRIVILEGES`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#privilege FileFormatGrant#privilege}
	Privilege *string `field:"optional" json:"privilege" yaml:"privilege"`
	// The name of the role to revert ownership to on destroy.
	//
	// Has no effect unless `privilege` is set to `OWNERSHIP`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#revert_ownership_to_role_name FileFormatGrant#revert_ownership_to_role_name}
	RevertOwnershipToRoleName *string `field:"optional" json:"revertOwnershipToRoleName" yaml:"revertOwnershipToRoleName"`
	// The name of the schema containing the current or future file formats on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#schema_name FileFormatGrant#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format_grant#with_grant_option FileFormatGrant#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

