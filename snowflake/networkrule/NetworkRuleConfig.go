// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkRuleConfig struct {
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
	// The database in which to create the network rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/network_rule#database NetworkRule#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies what is restricted by the network rule. Valid values are INGRESS, INTERNAL_STAGE and EGRESS; see https://docs.snowflake.com/en/sql-reference/sql/create-network-rule#required-parameters for details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/network_rule#mode NetworkRule#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Specifies the identifier for the network rule;
	//
	// must be unique for the database and schema in which the network rule is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/network_rule#name NetworkRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the network rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/network_rule#schema NetworkRule#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the type of network identifiers being allowed or blocked.
	//
	// A network rule can have only one type. Allowed values are IPV4, AWSVPCEID, AZURELINKID and HOST_PORT; allowed values are determined by the mode of the network rule; see https://docs.snowflake.com/en/sql-reference/sql/create-network-rule#required-parameters for details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/network_rule#type NetworkRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the network identifiers that will be allowed or blocked.
	//
	// Valid values in the list are determined by the type of network rule, see https://docs.snowflake.com/en/sql-reference/sql/create-network-rule#required-parameters for details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/network_rule#value_list NetworkRule#value_list}
	ValueList *[]*string `field:"required" json:"valueList" yaml:"valueList"`
	// Specifies a comment for the network rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/network_rule#comment NetworkRule#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/network_rule#id NetworkRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

