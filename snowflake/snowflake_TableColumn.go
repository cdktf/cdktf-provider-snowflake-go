// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake


type TableColumn struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#name Table#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Column type, e.g. VARIANT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#type Table#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Column comment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#comment Table#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// default block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#default Table#default}
	Default *TableColumnDefault `field:"optional" json:"default" yaml:"default"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#identity Table#identity}
	Identity *TableColumnIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Whether this column can contain null values.
	//
	// **Note**: Depending on your Snowflake version, the default value will not suffice if this column is used in a primary key constraint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#nullable Table#nullable}
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
}

