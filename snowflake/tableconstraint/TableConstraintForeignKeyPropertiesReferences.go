// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tableconstraint


type TableConstraintForeignKeyPropertiesReferences struct {
	// Columns to use in foreign key reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/table_constraint#columns TableConstraint#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// Name of constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/table_constraint#table_id TableConstraint#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

