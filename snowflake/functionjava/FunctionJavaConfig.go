// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionjava

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionJavaConfig struct {
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
	// The database in which to create the function.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#database FunctionJava#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The name of the handler method or class.
	//
	// If the handler is for a scalar UDF, returning a non-tabular value, the HANDLER value should be a method name, as in the following form: `MyClass.myMethod`. If the handler is for a tabular UDF, the HANDLER value should be the name of a handler class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#handler FunctionJava#handler}
	Handler *string `field:"required" json:"handler" yaml:"handler"`
	// The name of the function;
	//
	// the identifier does not need to be unique for the schema in which the function is created because UDFs are identified and resolved by the combination of the name and argument types. Check the [docs](https://docs.snowflake.com/en/sql-reference/sql/create-function#all-languages). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#name FunctionJava#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the results returned by the UDF, which determines the UDF type.
	//
	// Use `<result_data_type>` to create a scalar UDF that returns a single value with the specified data type. Use `TABLE (col_name col_data_type, ...)` to creates a table UDF that returns tabular results with the specified table column(s) and column type(s). For the details, consult the [docs](https://docs.snowflake.com/en/sql-reference/sql/create-function#all-languages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#return_type FunctionJava#return_type}
	ReturnType *string `field:"required" json:"returnType" yaml:"returnType"`
	// The schema in which to create the function.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#schema FunctionJava#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#arguments FunctionJava#arguments}
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// (Default: `user-defined function`) Specifies a comment for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#comment FunctionJava#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Enable stdout/stderr fast path logging for anonymous stored procs.
	//
	// This is a public parameter (similar to LOG_LEVEL). For more information, check [ENABLE_CONSOLE_OUTPUT docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-console-output).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#enable_console_output FunctionJava#enable_console_output}
	EnableConsoleOutput interface{} `field:"optional" json:"enableConsoleOutput" yaml:"enableConsoleOutput"`
	// The names of [external access integrations](https://docs.snowflake.com/en/sql-reference/sql/create-external-access-integration) needed in order for this function’s handler code to access external networks. An external access integration specifies [network rules](https://docs.snowflake.com/en/sql-reference/sql/create-network-rule) and [secrets](https://docs.snowflake.com/en/sql-reference/sql/create-secret) that specify external locations and credentials (if any) allowed for use by handler code when making requests of an external network, such as an external REST API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#external_access_integrations FunctionJava#external_access_integrations}
	ExternalAccessIntegrations *[]*string `field:"optional" json:"externalAccessIntegrations" yaml:"externalAccessIntegrations"`
	// Defines the handler code executed when the UDF is called.
	//
	// Wrapping `$$` signs are added by the provider automatically; do not include them. The `function_definition` value must be Java source code. For more information, see [Introduction to Java UDFs](https://docs.snowflake.com/en/developer-guide/udf/java/udf-java-introduction). To mitigate permadiff on this field, the provider replaces blank characters with a space. This can lead to false positives in cases where a change in case or run of whitespace is semantically significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#function_definition FunctionJava#function_definition}
	FunctionDefinition *string `field:"optional" json:"functionDefinition" yaml:"functionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#id FunctionJava#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// imports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#imports FunctionJava#imports}
	Imports interface{} `field:"optional" json:"imports" yaml:"imports"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies that the function is secure.
	//
	// By design, the Snowflake's `SHOW FUNCTIONS` command does not provide information about secure functions (consult [function docs](https://docs.snowflake.com/en/sql-reference/sql/create-function#id1) and [Protecting Sensitive Information with Secure UDFs and Stored Procedures](https://docs.snowflake.com/en/developer-guide/secure-udf-procedure)) which is essential to manage/import function with Terraform. Use the role owning the function while managing secure functions. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#is_secure FunctionJava#is_secure}
	IsSecure *string `field:"optional" json:"isSecure" yaml:"isSecure"`
	// LOG_LEVEL to use when filtering events For more information, check [LOG_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#log-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#log_level FunctionJava#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// METRIC_LEVEL value to control whether to emit metrics to Event Table For more information, check [METRIC_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#metric-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#metric_level FunctionJava#metric_level}
	MetricLevel *string `field:"optional" json:"metricLevel" yaml:"metricLevel"`
	// Specifies the behavior of the function when called with null inputs.
	//
	// Valid values are (case-insensitive): `CALLED ON NULL INPUT` | `RETURNS NULL ON NULL INPUT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#null_input_behavior FunctionJava#null_input_behavior}
	NullInputBehavior *string `field:"optional" json:"nullInputBehavior" yaml:"nullInputBehavior"`
	// The name and version number of Snowflake system packages required as dependencies.
	//
	// The value should be of the form `package_name:version_number`, where `package_name` is `snowflake_domain:package`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#packages FunctionJava#packages}
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
	// Specifies the behavior of the function when returning results. Valid values are (case-insensitive): `VOLATILE` | `IMMUTABLE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#return_results_behavior FunctionJava#return_results_behavior}
	ReturnResultsBehavior *string `field:"optional" json:"returnResultsBehavior" yaml:"returnResultsBehavior"`
	// Specifies the Java JDK runtime version to use.
	//
	// The supported versions of Java are 11.x and 17.x. If RUNTIME_VERSION is not set, Java JDK 11 is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#runtime_version FunctionJava#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// secrets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#secrets FunctionJava#secrets}
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// target_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#target_path FunctionJava#target_path}
	TargetPath *FunctionJavaTargetPath `field:"optional" json:"targetPath" yaml:"targetPath"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#timeouts FunctionJava#timeouts}
	Timeouts *FunctionJavaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Trace level value to use when generating/filtering trace events For more information, check [TRACE_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#trace-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_java#trace_level FunctionJava#trace_level}
	TraceLevel *string `field:"optional" json:"traceLevel" yaml:"traceLevel"`
}

