package externalfunction


type ExternalFunctionHeader struct {
	// Header name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/external_function#name ExternalFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Header value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/external_function#value ExternalFunction#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

