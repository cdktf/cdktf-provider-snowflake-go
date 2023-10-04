// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamConfig struct {
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
	// The database in which to create the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#database Stream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the stream;
	//
	// must be unique for the database and schema in which the stream is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#name Stream#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#schema Stream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Type of the stream that will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#append_only Stream#append_only}
	AppendOnly interface{} `field:"optional" json:"appendOnly" yaml:"appendOnly"`
	// Specifies a comment for the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#comment Stream#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#id Stream#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Create an insert only stream type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#insert_only Stream#insert_only}
	InsertOnly interface{} `field:"optional" json:"insertOnly" yaml:"insertOnly"`
	// Name of the stage the stream will monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#on_stage Stream#on_stage}
	OnStage *string `field:"optional" json:"onStage" yaml:"onStage"`
	// Name of the table the stream will monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#on_table Stream#on_table}
	OnTable *string `field:"optional" json:"onTable" yaml:"onTable"`
	// Name of the view the stream will monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#on_view Stream#on_view}
	OnView *string `field:"optional" json:"onView" yaml:"onView"`
	// Specifies whether to return all existing rows in the source table as row inserts the first time the stream is consumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.72.0/docs/resources/stream#show_initial_rows Stream#show_initial_rows}
	ShowInitialRows interface{} `field:"optional" json:"showInitialRows" yaml:"showInitialRows"`
}

