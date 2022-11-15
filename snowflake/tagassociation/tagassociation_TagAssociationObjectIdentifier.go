package tagassociation


type TagAssociationObjectIdentifier struct {
	// Name of the object to associate the tag with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#name TagAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the database that the object was created in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#database TagAssociation#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Name of the schema that the object was created in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#schema TagAssociation#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

