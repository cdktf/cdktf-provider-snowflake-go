// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ViewConfig struct {
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
	// The database in which to create the view.
	//
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#database View#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the view;
	//
	// must be unique for the schema in which the view is created. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#name View#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the view.
	//
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#schema View#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the query used to create the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#statement View#statement}
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// aggregation_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#aggregation_policy View#aggregation_policy}
	AggregationPolicy *ViewAggregationPolicy `field:"optional" json:"aggregationPolicy" yaml:"aggregationPolicy"`
	// Specifies to enable or disable change tracking on the table.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#change_tracking View#change_tracking}
	ChangeTracking *string `field:"optional" json:"changeTracking" yaml:"changeTracking"`
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#column View#column}
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// Specifies a comment for the view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#comment View#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Retains the access permissions from the original view when a new view is created using the OR REPLACE clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#copy_grants View#copy_grants}
	CopyGrants interface{} `field:"optional" json:"copyGrants" yaml:"copyGrants"`
	// data_metric_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#data_metric_function View#data_metric_function}
	DataMetricFunction interface{} `field:"optional" json:"dataMetricFunction" yaml:"dataMetricFunction"`
	// data_metric_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#data_metric_schedule View#data_metric_schedule}
	DataMetricSchedule *ViewDataMetricSchedule `field:"optional" json:"dataMetricSchedule" yaml:"dataMetricSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#id View#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies that the view can refer to itself using recursive syntax without necessarily using a CTE (common table expression).
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#is_recursive View#is_recursive}
	IsRecursive *string `field:"optional" json:"isRecursive" yaml:"isRecursive"`
	// Specifies that the view is secure.
	//
	// By design, the Snowflake's `SHOW VIEWS` command does not provide information about secure views (consult [view usage notes](https://docs.snowflake.com/en/sql-reference/sql/create-view#usage-notes)) which is essential to manage/import view with Terraform. Use the role owning the view while managing secure views. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#is_secure View#is_secure}
	IsSecure *string `field:"optional" json:"isSecure" yaml:"isSecure"`
	// Specifies that the view persists only for the duration of the session that you created it in.
	//
	// A temporary view and all its contents are dropped at the end of the session. In context of this provider, it means that it's dropped after a Terraform operation. This results in a permanent plan with object creation. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#is_temporary View#is_temporary}
	IsTemporary *string `field:"optional" json:"isTemporary" yaml:"isTemporary"`
	// row_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#row_access_policy View#row_access_policy}
	RowAccessPolicy *ViewRowAccessPolicy `field:"optional" json:"rowAccessPolicy" yaml:"rowAccessPolicy"`
}

