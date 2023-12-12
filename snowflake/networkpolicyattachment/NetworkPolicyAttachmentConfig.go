// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicyattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkPolicyAttachmentConfig struct {
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
	// Specifies the identifier for the network policy;
	//
	// must be unique for the account in which the network policy is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/network_policy_attachment#network_policy_name NetworkPolicyAttachment#network_policy_name}
	NetworkPolicyName *string `field:"required" json:"networkPolicyName" yaml:"networkPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/network_policy_attachment#id NetworkPolicyAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether the network policy should be applied globally to your Snowflake account<br><br>**Note:** The Snowflake user running `terraform apply` must be on an IP address allowed by the network policy to set that policy globally on the Snowflake account.<br><br>Additionally, a Snowflake account can only have one network policy set globally at any given time. This resource does not enforce one-policy-per-account, it is the user's responsibility to enforce this. If multiple network policy resources have `set_for_account: true`, the final policy set on the account will be non-deterministic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/network_policy_attachment#set_for_account NetworkPolicyAttachment#set_for_account}
	SetForAccount interface{} `field:"optional" json:"setForAccount" yaml:"setForAccount"`
	// Specifies which users the network policy should be attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/network_policy_attachment#users NetworkPolicyAttachment#users}
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

