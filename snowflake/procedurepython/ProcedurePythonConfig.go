// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedurepython

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProcedurePythonConfig struct {
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
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#database ProcedurePython#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Use the name of the stored procedure’s function or method.
	//
	// This can differ depending on whether the code is in-line or referenced at a stage. When the code is in-line, you can specify just the function name. When the code is imported from a stage, specify the fully-qualified handler function name as `<module_name>.<function_name>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#handler ProcedurePython#handler}
	Handler *string `field:"required" json:"handler" yaml:"handler"`
	// The name of the procedure;
	//
	// the identifier does not need to be unique for the schema in which the procedure is created because stored procedures are [identified and resolved by the combination of the name and argument types](https://docs.snowflake.com/en/developer-guide/udf-stored-procedure-naming-conventions.html#label-procedure-function-name-overloading). Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#name ProcedurePython#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the type of the result returned by the stored procedure.
	//
	// For `<result_data_type>`, use the Snowflake data type that corresponds to the type of the language that you are using (see [SQL-Python Data Type Mappings](https://docs.snowflake.com/en/developer-guide/udf-stored-procedure-data-type-mapping.html#label-sql-python-data-type-mappings)). For `RETURNS TABLE ( [ col_name col_data_type [ , ... ] ] )`, if you know the Snowflake data types of the columns in the returned table, specify the column names and types. Otherwise (e.g. if you are determining the column types during run time), you can omit the column names and types (i.e. `TABLE ()`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#return_type ProcedurePython#return_type}
	ReturnType *string `field:"required" json:"returnType" yaml:"returnType"`
	// The language runtime version to use. Currently, the supported versions are: 3.9, 3.10, and 3.11.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#runtime_version ProcedurePython#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// The schema in which to create the procedure.
	//
	// Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#schema ProcedurePython#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The Snowpark package is required for stored procedures, so it must always be present.
	//
	// For more information about Snowpark, see [Snowpark API](https://docs.snowflake.com/en/developer-guide/snowpark/index).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#snowpark_package ProcedurePython#snowpark_package}
	SnowparkPackage *string `field:"required" json:"snowparkPackage" yaml:"snowparkPackage"`
	// arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#arguments ProcedurePython#arguments}
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// Specifies a comment for the procedure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#comment ProcedurePython#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Enable stdout/stderr fast path logging for anonyous stored procs.
	//
	// This is a public parameter (similar to LOG_LEVEL). For more information, check [ENABLE_CONSOLE_OUTPUT docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-console-output).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#enable_console_output ProcedurePython#enable_console_output}
	EnableConsoleOutput interface{} `field:"optional" json:"enableConsoleOutput" yaml:"enableConsoleOutput"`
	// Specifies whether the stored procedure executes with the privileges of the owner (an “owner’s rights” stored procedure) or with the privileges of the caller (a “caller’s rights” stored procedure).
	//
	// If you execute the statement CREATE PROCEDURE … EXECUTE AS CALLER, then in the future the procedure will execute as a caller’s rights procedure. If you execute CREATE PROCEDURE … EXECUTE AS OWNER, then the procedure will execute as an owner’s rights procedure. For more information, see [Understanding caller’s rights and owner’s rights stored procedures](https://docs.snowflake.com/en/developer-guide/stored-procedure/stored-procedures-rights). Valid values are (case-insensitive): `CALLER` | `OWNER`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#execute_as ProcedurePython#execute_as}
	ExecuteAs *string `field:"optional" json:"executeAs" yaml:"executeAs"`
	// The names of [external access integrations](https://docs.snowflake.com/en/sql-reference/sql/create-external-access-integration) needed in order for this procedure’s handler code to access external networks. An external access integration specifies [network rules](https://docs.snowflake.com/en/sql-reference/sql/create-network-rule) and [secrets](https://docs.snowflake.com/en/sql-reference/sql/create-secret) that specify external locations and credentials (if any) allowed for use by handler code when making requests of an external network, such as an external REST API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#external_access_integrations ProcedurePython#external_access_integrations}
	ExternalAccessIntegrations *[]*string `field:"optional" json:"externalAccessIntegrations" yaml:"externalAccessIntegrations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#id ProcedurePython#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// imports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#imports ProcedurePython#imports}
	Imports interface{} `field:"optional" json:"imports" yaml:"imports"`
	// Specifies that the procedure is secure.
	//
	// For more information about secure procedures, see [Protecting Sensitive Information with Secure UDFs and Stored Procedures](https://docs.snowflake.com/en/developer-guide/secure-udf-procedure). Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#is_secure ProcedurePython#is_secure}
	IsSecure *string `field:"optional" json:"isSecure" yaml:"isSecure"`
	// LOG_LEVEL to use when filtering events For more information, check [LOG_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#log-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#log_level ProcedurePython#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// METRIC_LEVEL value to control whether to emit metrics to Event Table For more information, check [METRIC_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#metric-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#metric_level ProcedurePython#metric_level}
	MetricLevel *string `field:"optional" json:"metricLevel" yaml:"metricLevel"`
	// Specifies the behavior of the procedure when called with null inputs.
	//
	// Valid values are (case-insensitive): `CALLED ON NULL INPUT` | `RETURNS NULL ON NULL INPUT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#null_input_behavior ProcedurePython#null_input_behavior}
	NullInputBehavior *string `field:"optional" json:"nullInputBehavior" yaml:"nullInputBehavior"`
	// List of the names of packages deployed in Snowflake that should be included in the handler code’s execution environment.
	//
	// The Snowpark package is required for stored procedures, but is specified in the `snowpark_package` attribute. For more information about Snowpark, see [Snowpark API](https://docs.snowflake.com/en/developer-guide/snowpark/index).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#packages ProcedurePython#packages}
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
	// Defines the code executed by the stored procedure.
	//
	// The definition can consist of any valid code. Wrapping `$$` signs are added by the provider automatically; do not include them. The `procedure_definition` value must be Python source code. For more information, see [Python (using Snowpark)](https://docs.snowflake.com/en/developer-guide/stored-procedure/python/procedure-python-overview). To mitigate permadiff on this field, the provider replaces blank characters with a space. This can lead to false positives in cases where a change in case or run of whitespace is semantically significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#procedure_definition ProcedurePython#procedure_definition}
	ProcedureDefinition *string `field:"optional" json:"procedureDefinition" yaml:"procedureDefinition"`
	// secrets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#secrets ProcedurePython#secrets}
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// Trace level value to use when generating/filtering trace events For more information, check [TRACE_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#trace-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.1/docs/resources/procedure_python#trace_level ProcedurePython#trace_level}
	TraceLevel *string `field:"optional" json:"traceLevel" yaml:"traceLevel"`
}

