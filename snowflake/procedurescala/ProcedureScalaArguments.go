// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedurescala


type ProcedureScalaArguments struct {
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/resources/procedure_scala#arg_data_type ProcedureScala#arg_data_type}
	ArgDataType *string `field:"required" json:"argDataType" yaml:"argDataType"`
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/resources/procedure_scala#arg_name ProcedureScala#arg_name}
	ArgName *string `field:"required" json:"argName" yaml:"argName"`
	// Optional default value for the argument.
	//
	// For text values use single quotes. Numeric values can be unquoted. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.0/docs/resources/procedure_scala#arg_default_value ProcedureScala#arg_default_value}
	ArgDefaultValue *string `field:"optional" json:"argDefaultValue" yaml:"argDefaultValue"`
}

