package warehouse


type WarehouseTag struct {
	// Tag name, e.g. department.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/warehouse#name Warehouse#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tag value, e.g. marketing_info.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/warehouse#value Warehouse#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Name of the database that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/warehouse#database Warehouse#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the schema that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/warehouse#schema Warehouse#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

