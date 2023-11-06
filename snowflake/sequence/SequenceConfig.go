// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sequence

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SequenceConfig struct {
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
	// The database in which to create the sequence. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/sequence#database Sequence#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the name for the sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/sequence#name Sequence#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the sequence. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/sequence#schema Sequence#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies a comment for the sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/sequence#comment Sequence#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/sequence#id Sequence#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The amount the sequence will increase by each time it is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.75.0/docs/resources/sequence#increment Sequence#increment}
	Increment *float64 `field:"optional" json:"increment" yaml:"increment"`
}

