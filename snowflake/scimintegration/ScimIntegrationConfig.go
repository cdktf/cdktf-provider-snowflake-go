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
	// Specify whether the security integration is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/scim_integration#enabled ScimIntegration#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// String that specifies the identifier (i.e. name) for the integration; must be unique in your account. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/scim_integration#name ScimIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify the SCIM role in Snowflake that owns any users and roles that are imported from the identity provider into Snowflake using SCIM.
	//
	// Provider assumes that the specified role is already provided. Valid options are: `OKTA_PROVISIONER` | `AAD_PROVISIONER` | `GENERIC_SCIM_PROVISIONER`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/scim_integration#run_as_role ScimIntegration#run_as_role}
	RunAsRole *string `field:"required" json:"runAsRole" yaml:"runAsRole"`
	// Specifies the client type for the scim integration. Valid options are: `OKTA` | `AZURE` | `GENERIC`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/scim_integration#scim_client ScimIntegration#scim_client}
	ScimClient *string `field:"required" json:"scimClient" yaml:"scimClient"`
	// Specifies a comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/scim_integration#comment ScimIntegration#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/scim_integration#id ScimIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies an existing network policy that controls SCIM network traffic. For more information about this resource, see [docs](./network_policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/scim_integration#network_policy ScimIntegration#network_policy}
	NetworkPolicy *string `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
	// Specifies whether to enable or disable the synchronization of a user password from an Okta SCIM client as part of the API request to Snowflake.
	//
	// This property is not supported for Azure SCIM. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/scim_integration#sync_password ScimIntegration#sync_password}
	SyncPassword *string `field:"optional" json:"syncPassword" yaml:"syncPassword"`
}

