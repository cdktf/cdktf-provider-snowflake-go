package objectparameter


type ObjectParameterObjectIdentifier struct {
	// Name of the object to set the parameter for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/object_parameter#name ObjectParameter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the database that the object was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/object_parameter#database ObjectParameter#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the schema that the object was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/object_parameter#schema ObjectParameter#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

