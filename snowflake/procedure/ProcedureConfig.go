// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedure

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProcedureConfig struct {
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
	// The database in which to create the procedure. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#database Procedure#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the procedure;
	//
	// does not have to be unique for the schema in which the procedure is created. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#name Procedure#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The return type of the procedure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#return_type Procedure#return_type}
	ReturnType *string `field:"required" json:"returnType" yaml:"returnType"`
	// The schema in which to create the procedure. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#schema Procedure#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the code used to create the procedure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#statement Procedure#statement}
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#arguments Procedure#arguments}
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// Specifies a comment for the procedure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#comment Procedure#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Sets execution context.
	//
	// Allowed values are CALLER and OWNER (consult a proper section in the [docs](https://docs.snowflake.com/en/sql-reference/sql/create-procedure#id1)). For more information see [caller's rights and owner's rights](https://docs.snowflake.com/en/developer-guide/stored-procedure/stored-procedures-rights).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#execute_as Procedure#execute_as}
	ExecuteAs *string `field:"optional" json:"executeAs" yaml:"executeAs"`
	// The handler method for Java / Python procedures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#handler Procedure#handler}
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#id Procedure#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Imports for Java / Python procedures.
	//
	// For Java this a list of jar files, for Python this is a list of Python files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#imports Procedure#imports}
	Imports *[]*string `field:"optional" json:"imports" yaml:"imports"`
	// Specifies the language of the stored procedure code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#language Procedure#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// Specifies the behavior of the procedure when called with null inputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#null_input_behavior Procedure#null_input_behavior}
	NullInputBehavior *string `field:"optional" json:"nullInputBehavior" yaml:"nullInputBehavior"`
	// List of package imports to use for Java / Python procedures.
	//
	// For Java, package imports should be of the form: package_name:version_number, where package_name is snowflake_domain:package. For Python use it should be: ('numpy','pandas','xgboost==1.5.0').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#packages Procedure#packages}
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
	// Specifies the behavior of the function when returning results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#return_behavior Procedure#return_behavior}
	ReturnBehavior *string `field:"optional" json:"returnBehavior" yaml:"returnBehavior"`
	// Required for Python procedures. Specifies Python runtime version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#runtime_version Procedure#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// Specifies that the procedure is secure.
	//
	// For more information about secure procedures, see Protecting Sensitive Information with Secure UDFs and Stored Procedures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/procedure#secure Procedure#secure}
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
}

