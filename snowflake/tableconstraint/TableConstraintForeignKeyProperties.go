// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tableconstraint


type TableConstraintForeignKeyProperties struct {
	// The match type for the foreign key. Not applicable for primary/unique keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/table_constraint#match TableConstraint#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Specifies the action performed when the primary/unique key for the foreign key is deleted. Not applicable for primary/unique keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/table_constraint#on_delete TableConstraint#on_delete}
	OnDelete *string `field:"optional" json:"onDelete" yaml:"onDelete"`
	// Specifies the action performed when the primary/unique key for the foreign key is updated. Not applicable for primary/unique keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/table_constraint#on_update TableConstraint#on_update}
	OnUpdate *string `field:"optional" json:"onUpdate" yaml:"onUpdate"`
	// references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/table_constraint#references TableConstraint#references}
	References *TableConstraintForeignKeyPropertiesReferences `field:"optional" json:"references" yaml:"references"`
}

