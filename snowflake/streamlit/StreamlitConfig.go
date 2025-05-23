// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamlit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamlitConfig struct {
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
	// The database in which to create the streamlit Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#database Streamlit#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the filename of the Streamlit Python application. This filename is relative to the value of `directory_location`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#main_file Streamlit#main_file}
	MainFile *string `field:"required" json:"mainFile" yaml:"mainFile"`
	// String that specifies the identifier (i.e. name) for the streamlit; must be unique in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#name Streamlit#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the streamlit.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#schema Streamlit#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The stage in which streamlit files are located. For more information about this resource, see [docs](./stage).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#stage Streamlit#stage}
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Specifies a comment for the streamlit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#comment Streamlit#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the full path to the named stage containing the Streamlit Python files, media files, and the environment.yml file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#directory_location Streamlit#directory_location}
	DirectoryLocation *string `field:"optional" json:"directoryLocation" yaml:"directoryLocation"`
	// External access integrations connected to the Streamlit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#external_access_integrations Streamlit#external_access_integrations}
	ExternalAccessIntegrations *[]*string `field:"optional" json:"externalAccessIntegrations" yaml:"externalAccessIntegrations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#id Streamlit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the warehouse where SQL queries issued by the Streamlit application are run.
	//
	// Due to Snowflake limitations warehouse identifier can consist of only upper-cased letters. For more information about this resource, see [docs](./warehouse).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#query_warehouse Streamlit#query_warehouse}
	QueryWarehouse *string `field:"optional" json:"queryWarehouse" yaml:"queryWarehouse"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#timeouts Streamlit#timeouts}
	Timeouts *StreamlitTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Specifies a title for the Streamlit app to display in Snowsight.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/1.2.1/docs/resources/streamlit#title Streamlit#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

