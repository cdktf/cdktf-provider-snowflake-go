// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake


type TablePrimaryKey struct {
	// Columns to use in primary key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#keys Table#keys}
	Keys *[]*string `field:"required" json:"keys" yaml:"keys"`
	// Name of constraint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#name Table#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

