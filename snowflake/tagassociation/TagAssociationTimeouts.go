// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagassociation


type TagAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/tag_association#create TagAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

