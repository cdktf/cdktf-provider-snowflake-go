package schema


type SchemaTag struct {
	// Tag name, e.g. department.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/schema#name Schema#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tag value, e.g. marketing_info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/schema#value Schema#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Name of the database that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/schema#database Schema#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the schema that the tag was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/schema#schema Schema#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

