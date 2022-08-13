// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProcedureConfig struct {
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
	// The database in which to create the procedure. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#database Procedure#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the procedure;
	//
	// does not have to be unique for the schema in which the procedure is created. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#name Procedure#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The return type of the procedure.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#return_type Procedure#return_type}
	ReturnType *string `field:"required" json:"returnType" yaml:"returnType"`
	// The schema in which to create the procedure. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#schema Procedure#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the javascript code used to create the procedure.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#statement Procedure#statement}
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// arguments block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#arguments Procedure#arguments}
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// Specifies a comment for the procedure.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#comment Procedure#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Sets execute context - see caller's rights and owner's rights.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#execute_as Procedure#execute_as}
	ExecuteAs *string `field:"optional" json:"executeAs" yaml:"executeAs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#id Procedure#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the language of the stored procedure code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#language Procedure#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// Specifies the behavior of the procedure when called with null inputs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#null_input_behavior Procedure#null_input_behavior}
	NullInputBehavior *string `field:"optional" json:"nullInputBehavior" yaml:"nullInputBehavior"`
	// Specifies the behavior of the function when returning results.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/procedure#return_behavior Procedure#return_behavior}
	ReturnBehavior *string `field:"optional" json:"returnBehavior" yaml:"returnBehavior"`
}

