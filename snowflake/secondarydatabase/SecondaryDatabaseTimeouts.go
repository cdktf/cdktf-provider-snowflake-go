// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secondarydatabase


type SecondaryDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/secondary_database#create SecondaryDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/secondary_database#delete SecondaryDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/secondary_database#read SecondaryDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/secondary_database#update SecondaryDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

