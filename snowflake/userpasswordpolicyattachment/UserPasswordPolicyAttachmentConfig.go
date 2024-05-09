// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userpasswordpolicyattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserPasswordPolicyAttachmentConfig struct {
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
	// Fully qualified name of the password policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/user_password_policy_attachment#password_policy_name UserPasswordPolicyAttachment#password_policy_name}
	PasswordPolicyName *string `field:"required" json:"passwordPolicyName" yaml:"passwordPolicyName"`
	// User name of the user you want to attach the password policy to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/user_password_policy_attachment#user_name UserPasswordPolicyAttachment#user_name}
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.90.0/docs/resources/user_password_policy_attachment#id UserPasswordPolicyAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

