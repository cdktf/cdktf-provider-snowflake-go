// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package primaryconnection


type PrimaryConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/primary_connection#create PrimaryConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/primary_connection#delete PrimaryConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/primary_connection#read PrimaryConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/primary_connection#update PrimaryConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

