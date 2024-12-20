// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionscala


type FunctionScalaImports struct {
	// Path for import on stage, without the leading `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/resources/function_scala#path_on_stage FunctionScala#path_on_stage}
	PathOnStage *string `field:"required" json:"pathOnStage" yaml:"pathOnStage"`
	// Stage location without leading `@`.
	//
	// To use your user's stage set this to `~`, otherwise pass fully qualified name of the stage (with every part contained in double quotes or use `snowflake_stage.<your stage's resource name>.fully_qualified_name` if you manage this stage through terraform).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/resources/function_scala#stage_location FunctionScala#stage_location}
	StageLocation *string `field:"required" json:"stageLocation" yaml:"stageLocation"`
}

