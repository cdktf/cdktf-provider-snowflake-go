package functionresource


type FunctionResourceArguments struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/function#name FunctionResource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/function#type FunctionResource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

