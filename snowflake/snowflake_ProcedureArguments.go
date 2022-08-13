// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake


type ProcedureArguments struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#name Procedure#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#type Procedure#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

