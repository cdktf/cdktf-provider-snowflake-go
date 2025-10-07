// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userprogrammaticaccesstoken


type UserProgrammaticAccessTokenTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/user_programmatic_access_token#create UserProgrammaticAccessToken#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/user_programmatic_access_token#delete UserProgrammaticAccessToken#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/user_programmatic_access_token#read UserProgrammaticAccessToken#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/user_programmatic_access_token#update UserProgrammaticAccessToken#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

