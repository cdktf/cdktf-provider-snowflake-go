package externaltable


type ExternalTableColumn struct {
	// String that specifies the expression for the column. When queried, the column returns results derived from this expression.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/external_table#as ExternalTable#as}
	As *string `field:"required" json:"as" yaml:"as"`
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/external_table#name ExternalTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Column type, e.g. VARIANT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/external_table#type ExternalTable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

