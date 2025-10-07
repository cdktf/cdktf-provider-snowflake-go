// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountauthenticationpolicyattachment


type AccountAuthenticationPolicyAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/account_authentication_policy_attachment#create AccountAuthenticationPolicyAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/account_authentication_policy_attachment#delete AccountAuthenticationPolicyAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/account_authentication_policy_attachment#read AccountAuthenticationPolicyAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/account_authentication_policy_attachment#update AccountAuthenticationPolicyAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

