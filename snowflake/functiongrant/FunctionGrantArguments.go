package functiongrant


type FunctionGrantArguments struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/function_grant#name FunctionGrant#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/function_grant#type FunctionGrant#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

