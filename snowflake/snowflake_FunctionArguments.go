// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake


type FunctionArguments struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/function#name Function#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/function#type Function#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

