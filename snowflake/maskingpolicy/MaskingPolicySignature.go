package maskingpolicy


type MaskingPolicySignature struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/masking_policy#column MaskingPolicy#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
}
