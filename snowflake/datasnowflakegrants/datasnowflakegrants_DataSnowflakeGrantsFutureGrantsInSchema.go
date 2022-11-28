package datasnowflakegrants


type DataSnowflakeGrantsFutureGrantsInSchema struct {
	// The name of the schema to list all privileges of new (ie. future) objects granted to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#schema_name DataSnowflakeGrants#schema_name}
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// The database in which the scehma resides. Optional when querying a schema in the current database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#database_name DataSnowflakeGrants#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

