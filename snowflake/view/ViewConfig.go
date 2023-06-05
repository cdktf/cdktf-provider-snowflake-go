package view

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ViewConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The database in which to create the view. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#database View#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the view;
	//
	// must be unique for the schema in which the view is created. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#name View#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the view. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#schema View#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the query used to create the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#statement View#statement}
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// Specifies a comment for the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#comment View#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Retains the access permissions from the original view when a new view is created using the OR REPLACE clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#copy_grants View#copy_grants}
	CopyGrants interface{} `field:"optional" json:"copyGrants" yaml:"copyGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#id View#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies that the view is secure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#is_secure View#is_secure}
	IsSecure interface{} `field:"optional" json:"isSecure" yaml:"isSecure"`
	// Overwrites the View if it exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#or_replace View#or_replace}
	OrReplace interface{} `field:"optional" json:"orReplace" yaml:"orReplace"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.66.1/docs/resources/view#tag View#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

