package datasnowflakegrants


type DataSnowflakeGrantsGrantsTo struct {
	// Lists all privileges and roles granted to the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#role DataSnowflakeGrants#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Lists all the privileges granted to the share.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#share DataSnowflakeGrants#share}
	Share *string `field:"optional" json:"share" yaml:"share"`
	// Lists all the roles granted to the user.
	//
	// Note that the PUBLIC role, which is automatically available to every user, is not listed
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#user DataSnowflakeGrants#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

