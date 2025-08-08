// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secondaryconnection


type SecondaryConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/secondary_connection#create SecondaryConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/secondary_connection#delete SecondaryConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/secondary_connection#read SecondaryConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/secondary_connection#update SecondaryConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

