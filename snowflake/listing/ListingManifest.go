// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package listing


type ListingManifest struct {
	// from_stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/listing#from_stage Listing#from_stage}
	FromStage *ListingManifestFromStage `field:"optional" json:"fromStage" yaml:"fromStage"`
	// Manifest provided as a string.
	//
	// Wrapping `$$` signs are added by the provider automatically; do not include them. For more information on manifest syntax, see [Listing manifest reference](https://docs.snowflake.com/en/progaccess/listing-manifest-reference). Also, the [multiline string syntax](https://developer.hashicorp.com/terraform/language/expressions/strings#heredoc-strings) is a must here. A proper YAML indentation (2 spaces) is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/listing#from_string Listing#from_string}
	FromString *string `field:"optional" json:"fromString" yaml:"fromString"`
}

