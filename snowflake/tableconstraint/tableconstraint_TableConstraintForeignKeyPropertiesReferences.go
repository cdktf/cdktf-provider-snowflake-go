package tableconstraint


type TableConstraintForeignKeyPropertiesReferences struct {
	// Columns to use in foreign key reference.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#columns TableConstraint#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// Name of constraint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#table_id TableConstraint#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

