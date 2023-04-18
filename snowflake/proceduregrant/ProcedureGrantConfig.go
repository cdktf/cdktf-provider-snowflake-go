package proceduregrant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProcedureGrantConfig struct {
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
	// The name of the database containing the current or future procedures on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#database_name ProcedureGrant#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Grants privilege to these roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#roles ProcedureGrant#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// List of the argument data types for the procedure (must be present if procedure has arguments and procedure_name is present).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#argument_data_types ProcedureGrant#argument_data_types}
	ArgumentDataTypes *[]*string `field:"optional" json:"argumentDataTypes" yaml:"argumentDataTypes"`
	// arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#arguments ProcedureGrant#arguments}
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// When this is set to true, multiple grants of the same type can be created.
	//
	// This will cause Terraform to not revoke grants applied to roles and objects outside Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#enable_multiple_grants ProcedureGrant#enable_multiple_grants}
	EnableMultipleGrants interface{} `field:"optional" json:"enableMultipleGrants" yaml:"enableMultipleGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#id ProcedureGrant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When this is set to true and a schema_name is provided, apply this grant on all future procedures in the given schema.
	//
	// When this is true and no schema_name is provided apply this grant on all future procedures in the given database. The procedure_name and shares fields must be unset in order to use on_future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#on_future ProcedureGrant#on_future}
	OnFuture interface{} `field:"optional" json:"onFuture" yaml:"onFuture"`
	// The privilege to grant on the current or future procedure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#privilege ProcedureGrant#privilege}
	Privilege *string `field:"optional" json:"privilege" yaml:"privilege"`
	// The name of the procedure on which to grant privileges immediately (only valid if on_future is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#procedure_name ProcedureGrant#procedure_name}
	ProcedureName *string `field:"optional" json:"procedureName" yaml:"procedureName"`
	// The return type of the procedure (must be present if procedure_name is present).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#return_type ProcedureGrant#return_type}
	ReturnType *string `field:"optional" json:"returnType" yaml:"returnType"`
	// The name of the schema containing the current or future procedures on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#schema_name ProcedureGrant#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Grants privilege to these shares (only valid if on_future is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#shares ProcedureGrant#shares}
	Shares *[]*string `field:"optional" json:"shares" yaml:"shares"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/procedure_grant#with_grant_option ProcedureGrant#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

