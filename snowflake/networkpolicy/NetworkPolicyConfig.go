// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkPolicyConfig struct {
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
	// must be unique for the account in which the network policy is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/network_policy#name NetworkPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies one or more IPv4 addresses (CIDR notation) that are allowed access to your Snowflake account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/network_policy#allowed_ip_list NetworkPolicy#allowed_ip_list}
	AllowedIpList *[]*string `field:"optional" json:"allowedIpList" yaml:"allowedIpList"`
	// Specifies a list of fully qualified network rules that contain the network identifiers that are allowed access to Snowflake.
	//
	// For more information about this resource, see [docs](./network_rule).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/network_policy#allowed_network_rule_list NetworkPolicy#allowed_network_rule_list}
	AllowedNetworkRuleList *[]*string `field:"optional" json:"allowedNetworkRuleList" yaml:"allowedNetworkRuleList"`
	// Specifies one or more IPv4 addresses (CIDR notation) that are denied access to your Snowflake account.
	//
	// **Do not** add `0.0.0.0/0` to `blocked_ip_list`, in order to block all IP addresses except a select list, you only need to add IP addresses to `allowed_ip_list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/network_policy#blocked_ip_list NetworkPolicy#blocked_ip_list}
	BlockedIpList *[]*string `field:"optional" json:"blockedIpList" yaml:"blockedIpList"`
	// Specifies a list of fully qualified network rules that contain the network identifiers that are denied access to Snowflake.
	//
	// For more information about this resource, see [docs](./network_rule).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/network_policy#blocked_network_rule_list NetworkPolicy#blocked_network_rule_list}
	BlockedNetworkRuleList *[]*string `field:"optional" json:"blockedNetworkRuleList" yaml:"blockedNetworkRuleList"`
	// Specifies a comment for the network policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/network_policy#comment NetworkPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/network_policy#id NetworkPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.1/docs/resources/network_policy#timeouts NetworkPolicy#timeouts}
	Timeouts *NetworkPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

