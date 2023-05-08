package functionresource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionResourceConfig struct {
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
	// The database in which to create the function. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#database FunctionResource#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the function;
	//
	// does not have to be unique for the schema in which the function is created. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#name FunctionResource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The return type of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#return_type FunctionResource#return_type}
	ReturnType *string `field:"required" json:"returnType" yaml:"returnType"`
	// The schema in which to create the function. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#schema FunctionResource#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the javascript / java / sql / python code used to create the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#statement FunctionResource#statement}
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#arguments FunctionResource#arguments}
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// Specifies a comment for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#comment FunctionResource#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The handler method for Java / Python function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#handler FunctionResource#handler}
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#id FunctionResource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Imports for Java / Python functions.
	//
	// For Java this a list of jar files, for Python this is a list of Python files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#imports FunctionResource#imports}
	Imports *[]*string `field:"optional" json:"imports" yaml:"imports"`
	// Specifies that the function is secure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#is_secure FunctionResource#is_secure}
	IsSecure interface{} `field:"optional" json:"isSecure" yaml:"isSecure"`
	// The language of the statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#language FunctionResource#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// Specifies the behavior of the function when called with null inputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#null_input_behavior FunctionResource#null_input_behavior}
	NullInputBehavior *string `field:"optional" json:"nullInputBehavior" yaml:"nullInputBehavior"`
	// List of package imports to use for Java / Python functions.
	//
	// For Java, package imports should be of the form: package_name:version_number, where package_name is snowflake_domain:package. For Python use it should be: ('numpy','pandas','xgboost==1.5.0').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#packages FunctionResource#packages}
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
	// Specifies the behavior of the function when returning results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#return_behavior FunctionResource#return_behavior}
	ReturnBehavior *string `field:"optional" json:"returnBehavior" yaml:"returnBehavior"`
	// Required for Python functions. Specifies Python runtime version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#runtime_version FunctionResource#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// The target path for the Java / Python functions.
	//
	// For Java, it is the path of compiled jar files and for the Python it is the path of the Python files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.63.0/docs/resources/function#target_path FunctionResource#target_path}
	TargetPath *string `field:"optional" json:"targetPath" yaml:"targetPath"`
}

