// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tableconstraint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TableConstraintConfig struct {
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
	// Columns to use in constraint key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#columns TableConstraint#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// Name of constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#name TableConstraint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identifier for table to create constraint on. Format must follow: "\"&lt;db_name&gt;\".\"&lt;schema_name&gt;\".\"&lt;table_name&gt;\"" or "&lt;db_name&gt;.&lt;schema_name&gt;.&lt;table_name&gt;" (snowflake_table.my_table.id).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#table_id TableConstraint#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// Type of constraint, one of 'UNIQUE', 'PRIMARY KEY', or 'FOREIGN KEY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#type TableConstraint#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Comment for the table constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#comment TableConstraint#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Whether the constraint is deferrable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#deferrable TableConstraint#deferrable}
	Deferrable interface{} `field:"optional" json:"deferrable" yaml:"deferrable"`
	// Specifies whether the constraint is enabled or disabled. These properties are provided for compatibility with Oracle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#enable TableConstraint#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Whether the constraint is enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#enforced TableConstraint#enforced}
	Enforced interface{} `field:"optional" json:"enforced" yaml:"enforced"`
	// foreign_key_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#foreign_key_properties TableConstraint#foreign_key_properties}
	ForeignKeyProperties *TableConstraintForeignKeyProperties `field:"optional" json:"foreignKeyProperties" yaml:"foreignKeyProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#id TableConstraint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the constraint is initially deferred or immediate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#initially TableConstraint#initially}
	Initially *string `field:"optional" json:"initially" yaml:"initially"`
	// Specifies whether a constraint in NOVALIDATE mode is taken into account during query rewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#rely TableConstraint#rely}
	Rely interface{} `field:"optional" json:"rely" yaml:"rely"`
	// Specifies whether to validate existing data on the table when a constraint is created.
	//
	// Only used in conjunction with the ENABLE property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.87.1/docs/resources/table_constraint#validate TableConstraint#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
}

