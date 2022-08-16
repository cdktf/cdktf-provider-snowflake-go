// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RowAccessPolicyGrantConfig struct {
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
	// The name of the database containing the row access policy on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/row_access_policy_grant#database_name RowAccessPolicyGrant#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the row access policy on which to grant privileges immediately.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/row_access_policy_grant#row_access_policy_name RowAccessPolicyGrant#row_access_policy_name}
	RowAccessPolicyName *string `field:"required" json:"rowAccessPolicyName" yaml:"rowAccessPolicyName"`
	// The name of the schema containing the row access policy on which to grant privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/row_access_policy_grant#schema_name RowAccessPolicyGrant#schema_name}
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// When this is set to true, multiple grants of the same type can be created.
	//
	// This will cause Terraform to not revoke grants applied to roles and objects outside Terraform.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/row_access_policy_grant#enable_multiple_grants RowAccessPolicyGrant#enable_multiple_grants}
	EnableMultipleGrants interface{} `field:"optional" json:"enableMultipleGrants" yaml:"enableMultipleGrants"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/row_access_policy_grant#id RowAccessPolicyGrant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The privilege to grant on the row access policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/row_access_policy_grant#privilege RowAccessPolicyGrant#privilege}
	Privilege *string `field:"optional" json:"privilege" yaml:"privilege"`
	// Grants privilege to these roles.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/row_access_policy_grant#roles RowAccessPolicyGrant#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// When this is set to true, allows the recipient role to grant the privileges to other roles.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/row_access_policy_grant#with_grant_option RowAccessPolicyGrant#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}
