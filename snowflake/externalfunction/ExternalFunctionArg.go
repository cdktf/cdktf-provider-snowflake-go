package externalfunction


type ExternalFunctionArg struct {
	// Argument name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/external_function#name ExternalFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Argument type, e.g. VARCHAR.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/external_function#type ExternalFunction#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

