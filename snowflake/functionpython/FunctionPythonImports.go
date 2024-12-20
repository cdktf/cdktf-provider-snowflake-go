// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionpython


type FunctionPythonImports struct {
	// Path for import on stage, without the leading `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/resources/function_python#path_on_stage FunctionPython#path_on_stage}
	PathOnStage *string `field:"required" json:"pathOnStage" yaml:"pathOnStage"`
	// Stage location without leading `@`.
	//
	// To use your user's stage set this to `~`, otherwise pass fully qualified name of the stage (with every part contained in double quotes or use `snowflake_stage.<your stage's resource name>.fully_qualified_name` if you manage this stage through terraform).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/resources/function_python#stage_location FunctionPython#stage_location}
	StageLocation *string `field:"required" json:"stageLocation" yaml:"stageLocation"`
}

