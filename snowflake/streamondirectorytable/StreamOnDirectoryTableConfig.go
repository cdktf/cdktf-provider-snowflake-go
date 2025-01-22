// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamondirectorytable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamOnDirectoryTableConfig struct {
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
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/stream_on_directory_table#database StreamOnDirectoryTable#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the stream;
	//
	// must be unique for the database and schema in which the stream is created. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/stream_on_directory_table#name StreamOnDirectoryTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the stream.
	//
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/stream_on_directory_table#schema StreamOnDirectoryTable#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies an identifier for the stage the stream will monitor.
	//
	// Due to Snowflake limitations, the provider can not read the stage's database and schema. For stages, Snowflake returns only partially qualified name instead of fully qualified name. Please use stages located in the same schema as the stream. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information about this resource, see [docs](./stage).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/stream_on_directory_table#stage StreamOnDirectoryTable#stage}
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Specifies a comment for the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/stream_on_directory_table#comment StreamOnDirectoryTable#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Retains the access permissions from the original stream when a stream is recreated using the OR REPLACE clause.
	//
	// This is used when the provider detects changes for fields that can not be changed by ALTER. This value will not have any effect during creating a new object with Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/stream_on_directory_table#copy_grants StreamOnDirectoryTable#copy_grants}
	CopyGrants interface{} `field:"optional" json:"copyGrants" yaml:"copyGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/stream_on_directory_table#id StreamOnDirectoryTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

