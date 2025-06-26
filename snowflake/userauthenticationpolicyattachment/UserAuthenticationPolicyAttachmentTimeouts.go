// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userauthenticationpolicyattachment


type UserAuthenticationPolicyAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/user_authentication_policy_attachment#create UserAuthenticationPolicyAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/user_authentication_policy_attachment#delete UserAuthenticationPolicyAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/user_authentication_policy_attachment#read UserAuthenticationPolicyAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/user_authentication_policy_attachment#update UserAuthenticationPolicyAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

