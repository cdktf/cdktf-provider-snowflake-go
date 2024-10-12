// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package materializedview

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaterializedViewConfig struct {
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
	// The database in which to create the view. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#database MaterializedView#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the view; must be unique for the schema in which the view is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#name MaterializedView#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the view. Don't use the | character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#schema MaterializedView#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the query used to create the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#statement MaterializedView#statement}
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// The warehouse name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#warehouse MaterializedView#warehouse}
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// Specifies a comment for the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#comment MaterializedView#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#id MaterializedView#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies that the view is secure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#is_secure MaterializedView#is_secure}
	IsSecure interface{} `field:"optional" json:"isSecure" yaml:"isSecure"`
	// Overwrites the View if it exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#or_replace MaterializedView#or_replace}
	OrReplace interface{} `field:"optional" json:"orReplace" yaml:"orReplace"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.97.0/docs/resources/materialized_view#tag MaterializedView#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

