package datasnowflakegrants


type DataSnowflakeGrantsGrantsOn struct {
	// Object hierarchy to list privileges on.
	//
	// The only valid value is: ACCOUNT. Setting this attribute lists all the account-level (i.e. global) privileges that have been granted to roles.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#account DataSnowflakeGrants#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Name of object to list privileges on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#object_name DataSnowflakeGrants#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// Type of object to list privileges on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#object_type DataSnowflakeGrants#object_type}
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
}
