// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package share


type ShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/share#create Share#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/share#delete Share#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/share#read Share#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/share#update Share#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

