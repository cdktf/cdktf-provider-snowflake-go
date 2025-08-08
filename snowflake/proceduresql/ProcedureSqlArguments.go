// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package proceduresql


type ProcedureSqlArguments struct {
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/procedure_sql#arg_data_type ProcedureSql#arg_data_type}
	ArgDataType *string `field:"required" json:"argDataType" yaml:"argDataType"`
	// The argument name.
	//
	// The provider wraps it in double quotes by default, so be aware of that while referencing the argument in the procedure definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/procedure_sql#arg_name ProcedureSql#arg_name}
	ArgName *string `field:"required" json:"argName" yaml:"argName"`
	// Optional default value for the argument.
	//
	// For text values use single quotes. Numeric values can be unquoted. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/procedure_sql#arg_default_value ProcedureSql#arg_default_value}
	ArgDefaultValue *string `field:"optional" json:"argDefaultValue" yaml:"argDefaultValue"`
}

