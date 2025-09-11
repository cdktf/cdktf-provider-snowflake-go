// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package table

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TableConfig struct {
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
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#column Table#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The database in which to create the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#database Table#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the table;
	//
	// must be unique for the database and schema in which the table is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#name Table#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#schema Table#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// (Default: `false`) Specifies whether to enable change tracking on the table. Default false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#change_tracking Table#change_tracking}
	ChangeTracking interface{} `field:"optional" json:"changeTracking" yaml:"changeTracking"`
	// A list of one or more table columns/expressions to be used as clustering key(s) for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#cluster_by Table#cluster_by}
	ClusterBy *[]*string `field:"optional" json:"clusterBy" yaml:"clusterBy"`
	// Specifies a comment for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#comment Table#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the retention period for the table so that Time Travel actions (SELECT, CLONE, UNDROP) can be performed on historical data in the table.
	//
	// If you wish to inherit the parent schema setting then pass in the schema attribute to this argument or do not fill this parameter at all; the default value for this field is -1, which is a fallback to use Snowflake default - in this case the schema value
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#data_retention_time_in_days Table#data_retention_time_in_days}
	DataRetentionTimeInDays *float64 `field:"optional" json:"dataRetentionTimeInDays" yaml:"dataRetentionTimeInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#id Table#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// primary_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#primary_key Table#primary_key}
	PrimaryKey *TablePrimaryKey `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#tag Table#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/table#timeouts Table#timeouts}
	Timeouts *TableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

