// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package saml2integration


type Saml2IntegrationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/saml2_integration#create Saml2Integration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/saml2_integration#delete Saml2Integration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/saml2_integration#read Saml2Integration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/saml2_integration#update Saml2Integration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

