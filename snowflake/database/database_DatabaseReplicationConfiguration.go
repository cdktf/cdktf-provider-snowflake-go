package database


type DatabaseReplicationConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/database#accounts Database#accounts}.
	Accounts *[]*string `field:"required" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/database#ignore_edition_check Database#ignore_edition_check}.
	IgnoreEditionCheck interface{} `field:"optional" json:"ignoreEditionCheck" yaml:"ignoreEditionCheck"`
}

