package datasnowflakegrants


type DataSnowflakeGrantsGrantsOf struct {
	// Lists all users and roles to which the role has been granted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#role DataSnowflakeGrants#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Lists all the accounts for the share and indicates the accounts that are using the share.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/grants#share DataSnowflakeGrants#share}
	Share *string `field:"optional" json:"share" yaml:"share"`
}

