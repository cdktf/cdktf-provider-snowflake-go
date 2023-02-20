package datasnowflakegrants


type DataSnowflakeGrantsFutureGrantsIn struct {
	// Lists all privileges on new (i.e. future) objects of a specified type in the database granted to a role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#database DataSnowflakeGrants#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#schema DataSnowflakeGrants#schema}
	Schema *DataSnowflakeGrantsFutureGrantsInSchema `field:"optional" json:"schema" yaml:"schema"`
}
