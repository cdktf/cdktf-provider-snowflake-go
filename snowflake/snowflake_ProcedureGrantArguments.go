// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake


type ProcedureGrantArguments struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure_grant#name ProcedureGrant#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure_grant#type ProcedureGrant#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

