// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedurejava


type ProcedureJavaTargetPath struct {
	// Path for import on stage, without the leading `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/procedure_java#path_on_stage ProcedureJava#path_on_stage}
	PathOnStage *string `field:"required" json:"pathOnStage" yaml:"pathOnStage"`
	// Stage location without leading `@`.
	//
	// To use your user's stage set this to `~`, otherwise pass fully qualified name of the stage (with every part contained in double quotes or use `snowflake_stage.<your stage's resource name>.fully_qualified_name` if you manage this stage through terraform).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/procedure_java#stage_location ProcedureJava#stage_location}
	StageLocation *string `field:"required" json:"stageLocation" yaml:"stageLocation"`
}

