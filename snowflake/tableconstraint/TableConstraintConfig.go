package tableconstraint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TableConstraintConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Columns to use in constraint key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#columns TableConstraint#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// Name of constraint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#name TableConstraint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Idenfifier for table to create constraint on.
	//
	// Must be of the form Note: format must follow: "<db_name>"."<schema_name>"."<table_name>" or "<db_name>.<schema_name>.<table_name>" or "<db_name>|<schema_name>.<table_name>" (snowflake_table.my_table.id)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#table_id TableConstraint#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// Type of constraint, one of 'UNIQUE', 'PRIMARY KEY', 'FOREIGN KEY', or 'NOT NULL'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#type TableConstraint#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Comment for the table constraint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#comment TableConstraint#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Whether the constraint is deferrable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#deferrable TableConstraint#deferrable}
	Deferrable interface{} `field:"optional" json:"deferrable" yaml:"deferrable"`
	// Specifies whether the constraint is enabled or disabled. These properties are provided for compatibility with Oracle.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#enable TableConstraint#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Whether the constraint is enforced.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#enforced TableConstraint#enforced}
	Enforced interface{} `field:"optional" json:"enforced" yaml:"enforced"`
	// foreign_key_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#foreign_key_properties TableConstraint#foreign_key_properties}
	ForeignKeyProperties *TableConstraintForeignKeyProperties `field:"optional" json:"foreignKeyProperties" yaml:"foreignKeyProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#id TableConstraint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the constraint is initially deferred or immediate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#initially TableConstraint#initially}
	Initially *string `field:"optional" json:"initially" yaml:"initially"`
	// Specifies whether a constraint in NOVALIDATE mode is taken into account during query rewrite.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#rely TableConstraint#rely}
	Rely interface{} `field:"optional" json:"rely" yaml:"rely"`
	// Specifies whether to validate existing data on the table when a constraint is created.
	//
	// Only used in conjunction with the ENABLE property.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table_constraint#validate TableConstraint#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
}
