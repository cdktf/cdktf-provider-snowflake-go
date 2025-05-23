// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package objectparameter


type ObjectParameterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/object_parameter#create ObjectParameter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/object_parameter#delete ObjectParameter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/object_parameter#read ObjectParameter#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/object_parameter#update ObjectParameter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

