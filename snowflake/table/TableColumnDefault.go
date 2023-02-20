package table


type TableColumnDefault struct {
	// The default constant value for the column.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#constant Table#constant}
	Constant *string `field:"optional" json:"constant" yaml:"constant"`
	// The default expression value for the column.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#expression Table#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The default sequence to use for the column.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#sequence Table#sequence}
	Sequence *string `field:"optional" json:"sequence" yaml:"sequence"`
}

