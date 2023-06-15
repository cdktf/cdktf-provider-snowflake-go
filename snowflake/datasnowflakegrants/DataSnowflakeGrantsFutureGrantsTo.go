package datasnowflakegrants


type DataSnowflakeGrantsFutureGrantsTo struct {
	// Lists all privileges on new (i.e. future) objects of a specified type in a database or schema granted to the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.2/docs/data-sources/grants#role DataSnowflakeGrants#role}
	Role *string `field:"required" json:"role" yaml:"role"`
}

