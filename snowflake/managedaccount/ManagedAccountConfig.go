package managedaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedAccountConfig struct {
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
	// Identifier, as well as login name, for the initial user in the managed account.
	//
	// This user serves as the account administrator for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/managed_account#admin_name ManagedAccount#admin_name}
	AdminName *string `field:"required" json:"adminName" yaml:"adminName"`
	// Password for the initial user in the managed account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/managed_account#admin_password ManagedAccount#admin_password}
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// Identifier for the managed account; must be unique for your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/managed_account#name ManagedAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies a comment for the managed account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/managed_account#comment ManagedAccount#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/managed_account#id ManagedAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the type of managed account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/managed_account#type ManagedAccount#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

