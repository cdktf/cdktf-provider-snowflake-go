package proceduregrant


type ProcedureGrantArguments struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#name ProcedureGrant#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#type ProcedureGrant#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

