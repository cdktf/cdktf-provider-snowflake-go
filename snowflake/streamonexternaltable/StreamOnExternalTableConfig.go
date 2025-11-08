// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamonexternaltable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamOnExternalTableConfig struct {
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
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#database StreamOnExternalTable#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies an identifier for the external table the stream will monitor.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information about this resource, see [docs](./external_table).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#external_table StreamOnExternalTable#external_table}
	ExternalTable *string `field:"required" json:"externalTable" yaml:"externalTable"`
	// Specifies the identifier for the stream;
	//
	// must be unique for the database and schema in which the stream is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#name StreamOnExternalTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the stream.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#schema StreamOnExternalTable#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#at StreamOnExternalTable#at}
	At *StreamOnExternalTableAt `field:"optional" json:"at" yaml:"at"`
	// before block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#before StreamOnExternalTable#before}
	Before *StreamOnExternalTableBefore `field:"optional" json:"before" yaml:"before"`
	// Specifies a comment for the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#comment StreamOnExternalTable#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// (Default: `false`) Retains the access permissions from the original stream when a stream is recreated using the OR REPLACE clause.
	//
	// This is used when the provider detects changes for fields that can not be changed by ALTER. This value will not have any effect during creating a new object with Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#copy_grants StreamOnExternalTable#copy_grants}
	CopyGrants interface{} `field:"optional" json:"copyGrants" yaml:"copyGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#id StreamOnExternalTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether this is an insert-only stream.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#insert_only StreamOnExternalTable#insert_only}
	InsertOnly *string `field:"optional" json:"insertOnly" yaml:"insertOnly"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/stream_on_external_table#timeouts StreamOnExternalTable#timeouts}
	Timeouts *StreamOnExternalTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

