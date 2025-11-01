// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tag


type TagTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/tag#create Tag#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/tag#delete Tag#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/tag#read Tag#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/tag#update Tag#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

