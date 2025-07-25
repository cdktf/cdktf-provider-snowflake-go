// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userprogrammaticaccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserProgrammaticAccessTokenConfig struct {
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
	// Specifies the name for the programmatic access token;
	//
	// must be unique for the user. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#name UserProgrammaticAccessToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the user that the token is associated with.
	//
	// A user cannot use another user's programmatic access token to authenticate. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#user UserProgrammaticAccessToken#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// Descriptive comment about the programmatic access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#comment UserProgrammaticAccessToken#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The number of days that the programmatic access token can be used for authentication.
	//
	// This field cannot be altered after the token is created. Instead, you must rotate the token with the `keeper` field. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#days_to_expiry UserProgrammaticAccessToken#days_to_expiry}
	DaysToExpiry *float64 `field:"optional" json:"daysToExpiry" yaml:"daysToExpiry"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Disables or enables the programmatic access token.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#disabled UserProgrammaticAccessToken#disabled}
	Disabled *string `field:"optional" json:"disabled" yaml:"disabled"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) This field is only used when the token is rotated by changing the `keeper` field.
	//
	// Sets the expiration time of the existing token secret to expire after the specified number of hours. You can set this to a value of 0 to expire the current token secret immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#expire_rotated_token_after_hours UserProgrammaticAccessToken#expire_rotated_token_after_hours}
	ExpireRotatedTokenAfterHours *float64 `field:"optional" json:"expireRotatedTokenAfterHours" yaml:"expireRotatedTokenAfterHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#id UserProgrammaticAccessToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Arbitrary string that, if and only if, changed from a non-empty to a different non-empty value (or known after apply), will trigger a key to be rotated.
	//
	// When you add this field to the configuration, or remove it from the configuration, the rotation is not triggered. When the token is rotated, the `token` and `rotated_token_name` fields are marked as computed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#keeper UserProgrammaticAccessToken#keeper}
	Keeper *string `field:"optional" json:"keeper" yaml:"keeper"`
	// The number of minutes during which a user can use this token to access Snowflake without being subject to an active network policy.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#mins_to_bypass_network_policy_requirement UserProgrammaticAccessToken#mins_to_bypass_network_policy_requirement}
	MinsToBypassNetworkPolicyRequirement *float64 `field:"optional" json:"minsToBypassNetworkPolicyRequirement" yaml:"minsToBypassNetworkPolicyRequirement"`
	// The name of the role used for privilege evaluation and object creation.
	//
	// This must be one of the roles that has already been granted to the user. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#role_restriction UserProgrammaticAccessToken#role_restriction}
	RoleRestriction *string `field:"optional" json:"roleRestriction" yaml:"roleRestriction"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.4.0/docs/resources/user_programmatic_access_token#timeouts UserProgrammaticAccessToken#timeouts}
	Timeouts *UserProgrammaticAccessTokenTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

