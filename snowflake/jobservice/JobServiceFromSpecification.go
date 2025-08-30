// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobservice


type JobServiceFromSpecification struct {
	// The file name of the service specification. Example: `spec.yaml`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/job_service#file JobService#file}
	File *string `field:"optional" json:"file" yaml:"file"`
	// The path to the service specification file on the given stage.
	//
	// When the path is specified, the `/` character is automatically added as a path prefix. Example: `path/to/spec`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/job_service#path JobService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The fully qualified name of the stage containing the service specification file.
	//
	// At symbol (`@`) is added automatically. Example: `"\"<db_name>\".\"<schema_name>\".\"<stage_name>\""`. For more information about this resource, see [docs](./stage).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/job_service#stage JobService#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// The embedded text of the service specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/job_service#text JobService#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

