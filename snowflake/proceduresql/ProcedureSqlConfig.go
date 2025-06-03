// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package proceduresql

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProcedureSqlConfig struct {
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
	// The database in which to create the procedure.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#database ProcedureSql#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The name of the procedure;
	//
	// the identifier does not need to be unique for the schema in which the procedure is created because stored procedures are [identified and resolved by the combination of the name and argument types](https://docs.snowflake.com/en/developer-guide/udf-stored-procedure-naming-conventions.html#label-procedure-function-name-overloading). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#name ProcedureSql#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Defines the code executed by the stored procedure.
	//
	// The definition can consist of any valid code. Wrapping `$$` signs are added by the provider automatically; do not include them. The `procedure_definition` value must be SQL source code. For more information, see [Snowflake Scripting](https://docs.snowflake.com/en/developer-guide/snowflake-scripting/index). To mitigate permadiff on this field, the provider replaces blank characters with a space. This can lead to false positives in cases where a change in case or run of whitespace is semantically significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#procedure_definition ProcedureSql#procedure_definition}
	ProcedureDefinition *string `field:"required" json:"procedureDefinition" yaml:"procedureDefinition"`
	// Specifies the type of the result returned by the stored procedure.
	//
	// For `<result_data_type>`, use the Snowflake data type that corresponds to the type of the language that you are using (see [SQL data type](https://docs.snowflake.com/en/sql-reference-data-types)). For `RETURNS TABLE ( [ col_name col_data_type [ , ... ] ] )`, if you know the Snowflake data types of the columns in the returned table, specify the column names and types. Otherwise (e.g. if you are determining the column types during run time), you can omit the column names and types (i.e. `TABLE ()`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#return_type ProcedureSql#return_type}
	ReturnType *string `field:"required" json:"returnType" yaml:"returnType"`
	// The schema in which to create the procedure.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#schema ProcedureSql#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#arguments ProcedureSql#arguments}
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// (Default: `user-defined procedure`) Specifies a comment for the procedure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#comment ProcedureSql#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Enable stdout/stderr fast path logging for anonyous stored procs.
	//
	// This is a public parameter (similar to LOG_LEVEL). For more information, check [ENABLE_CONSOLE_OUTPUT docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-console-output).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#enable_console_output ProcedureSql#enable_console_output}
	EnableConsoleOutput interface{} `field:"optional" json:"enableConsoleOutput" yaml:"enableConsoleOutput"`
	// Specifies whether the stored procedure executes with the privileges of the owner (an “owner’s rights” stored procedure) or with the privileges of the caller (a “caller’s rights” stored procedure).
	//
	// If you execute the statement CREATE PROCEDURE … EXECUTE AS CALLER, then in the future the procedure will execute as a caller’s rights procedure. If you execute CREATE PROCEDURE … EXECUTE AS OWNER, then the procedure will execute as an owner’s rights procedure. For more information, see [Understanding caller’s rights and owner’s rights stored procedures](https://docs.snowflake.com/en/developer-guide/stored-procedure/stored-procedures-rights). Valid values are (case-insensitive): `CALLER` | `OWNER`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#execute_as ProcedureSql#execute_as}
	ExecuteAs *string `field:"optional" json:"executeAs" yaml:"executeAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#id ProcedureSql#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies that the procedure is secure.
	//
	// For more information about secure procedures, see [Protecting Sensitive Information with Secure UDFs and Stored Procedures](https://docs.snowflake.com/en/developer-guide/secure-udf-procedure). Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#is_secure ProcedureSql#is_secure}
	IsSecure *string `field:"optional" json:"isSecure" yaml:"isSecure"`
	// LOG_LEVEL to use when filtering events For more information, check [LOG_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#log-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#log_level ProcedureSql#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// METRIC_LEVEL value to control whether to emit metrics to Event Table For more information, check [METRIC_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#metric-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#metric_level ProcedureSql#metric_level}
	MetricLevel *string `field:"optional" json:"metricLevel" yaml:"metricLevel"`
	// Specifies the behavior of the procedure when called with null inputs.
	//
	// Valid values are (case-insensitive): `CALLED ON NULL INPUT` | `RETURNS NULL ON NULL INPUT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#null_input_behavior ProcedureSql#null_input_behavior}
	NullInputBehavior *string `field:"optional" json:"nullInputBehavior" yaml:"nullInputBehavior"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#timeouts ProcedureSql#timeouts}
	Timeouts *ProcedureSqlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Trace level value to use when generating/filtering trace events For more information, check [TRACE_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#trace-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#trace_level ProcedureSql#trace_level}
	TraceLevel *string `field:"optional" json:"traceLevel" yaml:"traceLevel"`
}

