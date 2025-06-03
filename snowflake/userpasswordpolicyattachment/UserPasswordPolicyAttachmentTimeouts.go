// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userpasswordpolicyattachment


type UserPasswordPolicyAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/user_password_policy_attachment#create UserPasswordPolicyAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/user_password_policy_attachment#delete UserPasswordPolicyAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/user_password_policy_attachment#read UserPasswordPolicyAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/user_password_policy_attachment#update UserPasswordPolicyAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

