// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userpublickeys


type UserPublicKeysTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/user_public_keys#create UserPublicKeys#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/user_public_keys#delete UserPublicKeys#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/user_public_keys#read UserPublicKeys#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/user_public_keys#update UserPublicKeys#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

