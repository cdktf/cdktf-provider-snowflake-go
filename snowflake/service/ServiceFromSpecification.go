// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceFromSpecification struct {
	// The file name of the service specification. Example: `spec.yaml`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/service#file Service#file}
	File *string `field:"optional" json:"file" yaml:"file"`
	// The path to the service specification file on the given stage.
	//
	// When the path is specified, the `/` character is automatically added as a path prefix. Example: `path/to/spec`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/service#path Service#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The fully qualified name of the stage containing the service specification file.
	//
	// At symbol (`@`) is added automatically. Example: `"\"<db_name>\".\"<schema_name>\".\"<stage_name>\""`. For more information about this resource, see [docs](./stage).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/service#stage Service#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// The embedded text of the service specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.2.0/docs/resources/service#text Service#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

