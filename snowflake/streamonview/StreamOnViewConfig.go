// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamonview

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamOnViewConfig struct {
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
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#database StreamOnView#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the stream;
	//
	// must be unique for the database and schema in which the stream is created. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#name StreamOnView#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the stream.
	//
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#schema StreamOnView#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies an identifier for the view the stream will monitor.
	//
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#view StreamOnView#view}
	View *string `field:"required" json:"view" yaml:"view"`
	// Specifies whether this is an append-only stream.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#append_only StreamOnView#append_only}
	AppendOnly *string `field:"optional" json:"appendOnly" yaml:"appendOnly"`
	// at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#at StreamOnView#at}
	At *StreamOnViewAt `field:"optional" json:"at" yaml:"at"`
	// before block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#before StreamOnView#before}
	Before *StreamOnViewBefore `field:"optional" json:"before" yaml:"before"`
	// Specifies a comment for the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#comment StreamOnView#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Retains the access permissions from the original stream when a stream is recreated using the OR REPLACE clause.
	//
	// That is sometimes used when the provider detects changes for fields that can not be changed by ALTER. This value will not have any effect when creating a new stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#copy_grants StreamOnView#copy_grants}
	CopyGrants interface{} `field:"optional" json:"copyGrants" yaml:"copyGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#id StreamOnView#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether to return all existing rows in the source table as row inserts the first time the stream is consumed.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/stream_on_view#show_initial_rows StreamOnView#show_initial_rows}
	ShowInitialRows *string `field:"optional" json:"showInitialRows" yaml:"showInitialRows"`
}

