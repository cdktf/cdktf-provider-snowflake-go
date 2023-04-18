package stagegrant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StageGrantConfig struct {
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
	// The name of the database containing the current stage on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#database_name StageGrant#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Grants privilege to these roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#roles StageGrant#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// When this is set to true, multiple grants of the same type can be created.
	//
	// This will cause Terraform to not revoke grants applied to roles and objects outside Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#enable_multiple_grants StageGrant#enable_multiple_grants}
	EnableMultipleGrants interface{} `field:"optional" json:"enableMultipleGrants" yaml:"enableMultipleGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#id StageGrant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When this is set to true and a schema_name is provided, apply this grant on all future stages in the given schema.
	//
	// When this is true and no schema_name is provided apply this grant on all future stages in the given database. The stage_name field must be unset in order to use on_future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#on_future StageGrant#on_future}
	OnFuture interface{} `field:"optional" json:"onFuture" yaml:"onFuture"`
	// The privilege to grant on the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#privilege StageGrant#privilege}
	Privilege *string `field:"optional" json:"privilege" yaml:"privilege"`
	// The name of the schema containing the current stage on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#schema_name StageGrant#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// The name of the stage on which to grant privilege (only valid if on_future is false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#stage_name StageGrant#stage_name}
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.61.0/docs/resources/stage_grant#with_grant_option StageGrant#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

