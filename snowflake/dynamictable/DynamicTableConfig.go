// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamictable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamicTableConfig struct {
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
	// The database in which to create the dynamic table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#database DynamicTable#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier (i.e. name) for the dynamic table; must be unique for the schema in which the dynamic table is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#name DynamicTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the query to use to populate the dynamic table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#query DynamicTable#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The schema in which to create the dynamic table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#schema DynamicTable#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// target_lag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#target_lag DynamicTable#target_lag}
	TargetLag *DynamicTableTargetLag `field:"required" json:"targetLag" yaml:"targetLag"`
	// The warehouse in which to create the dynamic table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#warehouse DynamicTable#warehouse}
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// Specifies a comment for the dynamic table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#comment DynamicTable#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#id DynamicTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Initialize trigger for the dynamic table. Can only be set on creation. Available options are ON_CREATE and ON_SCHEDULE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#initialize DynamicTable#initialize}
	Initialize *string `field:"optional" json:"initialize" yaml:"initialize"`
	// Specifies whether to replace the dynamic table if it already exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#or_replace DynamicTable#or_replace}
	OrReplace interface{} `field:"optional" json:"orReplace" yaml:"orReplace"`
	// INCREMENTAL to use incremental refreshes, FULL to recompute the whole table on every refresh, or AUTO to let Snowflake decide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/dynamic_table#refresh_mode DynamicTable#refresh_mode}
	RefreshMode *string `field:"optional" json:"refreshMode" yaml:"refreshMode"`
}

