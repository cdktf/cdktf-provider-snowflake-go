// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scimintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ScimIntegrationConfig struct {
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
	// Specifies the name of the SCIM integration.
	//
	// This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/resources/scim_integration#name ScimIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/resources/scim_integration#provisioner_role ScimIntegration#provisioner_role}
	ProvisionerRole *string `field:"required" json:"provisionerRole" yaml:"provisionerRole"`
	// Specifies the client type for the scim integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/resources/scim_integration#scim_client ScimIntegration#scim_client}
	ScimClient *string `field:"required" json:"scimClient" yaml:"scimClient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/resources/scim_integration#id ScimIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies an existing network policy active for your account.
	//
	// The network policy restricts the list of user IP addresses when exchanging an authorization code for an access or refresh token and when using a refresh token to obtain a new access token. If this parameter is not set, the network policy for the account (if any) is used instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/resources/scim_integration#network_policy ScimIntegration#network_policy}
	NetworkPolicy *string `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
}

