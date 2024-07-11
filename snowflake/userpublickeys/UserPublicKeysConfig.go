// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userpublickeys

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserPublicKeysConfig struct {
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
	// Name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/user_public_keys#name UserPublicKeys#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/user_public_keys#id UserPublicKeys#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/user_public_keys#rsa_public_key UserPublicKeys#rsa_public_key}
	RsaPublicKey *string `field:"optional" json:"rsaPublicKey" yaml:"rsaPublicKey"`
	// Specifies the user’s second RSA public key;
	//
	// used to rotate the public and Public keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/user_public_keys#rsa_public_key_2 UserPublicKeys#rsa_public_key_2}
	RsaPublicKey2 *string `field:"optional" json:"rsaPublicKey2" yaml:"rsaPublicKey2"`
}

