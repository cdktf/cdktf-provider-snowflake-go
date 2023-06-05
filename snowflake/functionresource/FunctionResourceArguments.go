package functionresource


type FunctionResourceArguments struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/function#name FunctionResource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/function#type FunctionResource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

