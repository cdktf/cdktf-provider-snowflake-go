// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedurepython


type ProcedurePythonArguments struct {
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/procedure_python#arg_data_type ProcedurePython#arg_data_type}
	ArgDataType *string `field:"required" json:"argDataType" yaml:"argDataType"`
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/procedure_python#arg_name ProcedurePython#arg_name}
	ArgName *string `field:"required" json:"argName" yaml:"argName"`
	// Optional default value for the argument.
	//
	// For text values use single quotes. Numeric values can be unquoted. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/procedure_python#arg_default_value ProcedurePython#arg_default_value}
	ArgDefaultValue *string `field:"optional" json:"argDefaultValue" yaml:"argDefaultValue"`
}
