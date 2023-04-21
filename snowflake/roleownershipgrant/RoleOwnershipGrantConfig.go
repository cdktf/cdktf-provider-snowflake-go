package roleownershipgrant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleOwnershipGrantConfig struct {
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
	// The name of the role ownership is granted on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.62.0/docs/resources/role_ownership_grant#on_role_name RoleOwnershipGrant#on_role_name}
	OnRoleName *string `field:"required" json:"onRoleName" yaml:"onRoleName"`
	// The name of the role to grant ownership.
	//
	// Please ensure that the role that terraform is using is granted access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.62.0/docs/resources/role_ownership_grant#to_role_name RoleOwnershipGrant#to_role_name}
	ToRoleName *string `field:"required" json:"toRoleName" yaml:"toRoleName"`
	// Specifies whether to remove or transfer all existing outbound privileges on the object when ownership is transferred to a new role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.62.0/docs/resources/role_ownership_grant#current_grants RoleOwnershipGrant#current_grants}
	CurrentGrants *string `field:"optional" json:"currentGrants" yaml:"currentGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.62.0/docs/resources/role_ownership_grant#id RoleOwnershipGrant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the role to revert ownership to on destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.62.0/docs/resources/role_ownership_grant#revert_ownership_to_role_name RoleOwnershipGrant#revert_ownership_to_role_name}
	RevertOwnershipToRoleName *string `field:"optional" json:"revertOwnershipToRoleName" yaml:"revertOwnershipToRoleName"`
}

