// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package task

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TaskConfig struct {
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
	// The database in which to create the task.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#database Task#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the task;
	//
	// must be unique for the database and schema in which the task is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#name Task#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the task.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#schema Task#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Any single SQL statement, or a call to a stored procedure, executed when the task runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#sql_statement Task#sql_statement}
	SqlStatement *string `field:"required" json:"sqlStatement" yaml:"sqlStatement"`
	// Specifies if the task should be started or suspended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#started Task#started}
	Started interface{} `field:"required" json:"started" yaml:"started"`
	// Specifies the action that Snowflake performs for in-progress queries if connectivity is lost due to abrupt termination of a session (e.g. network outage, browser termination, service interruption). For more information, check [ABORT_DETACHED_QUERY docs](https://docs.snowflake.com/en/sql-reference/parameters#abort-detached-query).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#abort_detached_query Task#abort_detached_query}
	AbortDetachedQuery interface{} `field:"optional" json:"abortDetachedQuery" yaml:"abortDetachedQuery"`
	// Specifies one or more predecessor tasks for the current task.
	//
	// Use this option to [create a DAG](https://docs.snowflake.com/en/user-guide/tasks-graphs.html#label-task-dag) of tasks or add this task to an existing DAG. A DAG is a series of tasks that starts with a scheduled root task and is linked together by dependencies. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#after Task#after}
	After *[]*string `field:"optional" json:"after" yaml:"after"`
	// By default, Snowflake ensures that only one instance of a particular DAG is allowed to run at a time, setting the parameter value to TRUE permits DAG runs to overlap.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#allow_overlapping_execution Task#allow_overlapping_execution}
	AllowOverlappingExecution *string `field:"optional" json:"allowOverlappingExecution" yaml:"allowOverlappingExecution"`
	// Specifies whether autocommit is enabled for the session.
	//
	// Autocommit determines whether a DML statement, when executed without an active transaction, is automatically committed after the statement successfully completes. For more information, see [Transactions](https://docs.snowflake.com/en/sql-reference/transactions). For more information, check [AUTOCOMMIT docs](https://docs.snowflake.com/en/sql-reference/parameters#autocommit).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#autocommit Task#autocommit}
	Autocommit interface{} `field:"optional" json:"autocommit" yaml:"autocommit"`
	// The format of VARCHAR values passed as input to VARCHAR-to-BINARY conversion functions.
	//
	// For more information, see [Binary input and output](https://docs.snowflake.com/en/sql-reference/binary-input-output). For more information, check [BINARY_INPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#binary-input-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#binary_input_format Task#binary_input_format}
	BinaryInputFormat *string `field:"optional" json:"binaryInputFormat" yaml:"binaryInputFormat"`
	// The format for VARCHAR values returned as output by BINARY-to-VARCHAR conversion functions.
	//
	// For more information, see [Binary input and output](https://docs.snowflake.com/en/sql-reference/binary-input-output). For more information, check [BINARY_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#binary-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#binary_output_format Task#binary_output_format}
	BinaryOutputFormat *string `field:"optional" json:"binaryOutputFormat" yaml:"binaryOutputFormat"`
	// Parameter that specifies the maximum amount of memory the JDBC driver or ODBC driver should use for the result set from queries (in MB).
	//
	// For more information, check [CLIENT_MEMORY_LIMIT docs](https://docs.snowflake.com/en/sql-reference/parameters#client-memory-limit).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#client_memory_limit Task#client_memory_limit}
	ClientMemoryLimit *float64 `field:"optional" json:"clientMemoryLimit" yaml:"clientMemoryLimit"`
	// For specific ODBC functions and JDBC methods, this parameter can change the default search scope from all databases/schemas to the current database/schema.
	//
	// The narrower search typically returns fewer rows and executes more quickly. For more information, check [CLIENT_METADATA_REQUEST_USE_CONNECTION_CTX docs](https://docs.snowflake.com/en/sql-reference/parameters#client-metadata-request-use-connection-ctx).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#client_metadata_request_use_connection_ctx Task#client_metadata_request_use_connection_ctx}
	ClientMetadataRequestUseConnectionCtx interface{} `field:"optional" json:"clientMetadataRequestUseConnectionCtx" yaml:"clientMetadataRequestUseConnectionCtx"`
	// Parameter that specifies the number of threads used by the client to pre-fetch large result sets.
	//
	// The driver will attempt to honor the parameter value, but defines the minimum and maximum values (depending on your system’s resources) to improve performance. For more information, check [CLIENT_PREFETCH_THREADS docs](https://docs.snowflake.com/en/sql-reference/parameters#client-prefetch-threads).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#client_prefetch_threads Task#client_prefetch_threads}
	ClientPrefetchThreads *float64 `field:"optional" json:"clientPrefetchThreads" yaml:"clientPrefetchThreads"`
	// Parameter that specifies the maximum size of each set (or chunk) of query results to download (in MB).
	//
	// The JDBC driver downloads query results in chunks. For more information, check [CLIENT_RESULT_CHUNK_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#client-result-chunk-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#client_result_chunk_size Task#client_result_chunk_size}
	ClientResultChunkSize *float64 `field:"optional" json:"clientResultChunkSize" yaml:"clientResultChunkSize"`
	// Parameter that indicates whether to match column name case-insensitively in ResultSet.get* methods in JDBC. For more information, check [CLIENT_RESULT_COLUMN_CASE_INSENSITIVE docs](https://docs.snowflake.com/en/sql-reference/parameters#client-result-column-case-insensitive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#client_result_column_case_insensitive Task#client_result_column_case_insensitive}
	ClientResultColumnCaseInsensitive interface{} `field:"optional" json:"clientResultColumnCaseInsensitive" yaml:"clientResultColumnCaseInsensitive"`
	// Parameter that indicates whether to force a user to log in again after a period of inactivity in the session.
	//
	// For more information, check [CLIENT_SESSION_KEEP_ALIVE docs](https://docs.snowflake.com/en/sql-reference/parameters#client-session-keep-alive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#client_session_keep_alive Task#client_session_keep_alive}
	ClientSessionKeepAlive interface{} `field:"optional" json:"clientSessionKeepAlive" yaml:"clientSessionKeepAlive"`
	// Number of seconds in-between client attempts to update the token for the session. For more information, check [CLIENT_SESSION_KEEP_ALIVE_HEARTBEAT_FREQUENCY docs](https://docs.snowflake.com/en/sql-reference/parameters#client-session-keep-alive-heartbeat-frequency).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#client_session_keep_alive_heartbeat_frequency Task#client_session_keep_alive_heartbeat_frequency}
	ClientSessionKeepAliveHeartbeatFrequency *float64 `field:"optional" json:"clientSessionKeepAliveHeartbeatFrequency" yaml:"clientSessionKeepAliveHeartbeatFrequency"`
	// Specifies the [TIMESTAMP_* variation](https://docs.snowflake.com/en/sql-reference/data-types-datetime.html#label-datatypes-timestamp-variations) to use when binding timestamp variables for JDBC or ODBC applications that use the bind API to load data. For more information, check [CLIENT_TIMESTAMP_TYPE_MAPPING docs](https://docs.snowflake.com/en/sql-reference/parameters#client-timestamp-type-mapping).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#client_timestamp_type_mapping Task#client_timestamp_type_mapping}
	ClientTimestampTypeMapping *string `field:"optional" json:"clientTimestampTypeMapping" yaml:"clientTimestampTypeMapping"`
	// Specifies a comment for the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#comment Task#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies a string representation of key value pairs that can be accessed by all tasks in the task graph.
	//
	// Must be in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#config Task#config}
	Config *string `field:"optional" json:"config" yaml:"config"`
	// Specifies the input format for the DATE data type.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [DATE_INPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#date-input-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#date_input_format Task#date_input_format}
	DateInputFormat *string `field:"optional" json:"dateInputFormat" yaml:"dateInputFormat"`
	// Specifies the display format for the DATE data type.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [DATE_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#date-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#date_output_format Task#date_output_format}
	DateOutputFormat *string `field:"optional" json:"dateOutputFormat" yaml:"dateOutputFormat"`
	// Specifies whether to set the schema for unloaded Parquet files based on the logical column data types (i.e. the types in the unload SQL query or source table) or on the unloaded column values (i.e. the smallest data types and precision that support the values in the output columns of the unload SQL statement or source table). For more information, check [ENABLE_UNLOAD_PHYSICAL_TYPE_OPTIMIZATION docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-unload-physical-type-optimization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#enable_unload_physical_type_optimization Task#enable_unload_physical_type_optimization}
	EnableUnloadPhysicalTypeOptimization interface{} `field:"optional" json:"enableUnloadPhysicalTypeOptimization" yaml:"enableUnloadPhysicalTypeOptimization"`
	// Specifies the name of the notification integration used for error notifications.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information about this resource, see [docs](./notification_integration).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#error_integration Task#error_integration}
	ErrorIntegration *string `field:"optional" json:"errorIntegration" yaml:"errorIntegration"`
	// Specifies whether to return an error when the [MERGE](https://docs.snowflake.com/en/sql-reference/sql/merge) command is used to update or delete a target row that joins multiple source rows and the system cannot determine the action to perform on the target row. For more information, check [ERROR_ON_NONDETERMINISTIC_MERGE docs](https://docs.snowflake.com/en/sql-reference/parameters#error-on-nondeterministic-merge).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#error_on_nondeterministic_merge Task#error_on_nondeterministic_merge}
	ErrorOnNondeterministicMerge interface{} `field:"optional" json:"errorOnNondeterministicMerge" yaml:"errorOnNondeterministicMerge"`
	// Specifies whether to return an error when the [UPDATE](https://docs.snowflake.com/en/sql-reference/sql/update) command is used to update a target row that joins multiple source rows and the system cannot determine the action to perform on the target row. For more information, check [ERROR_ON_NONDETERMINISTIC_UPDATE docs](https://docs.snowflake.com/en/sql-reference/parameters#error-on-nondeterministic-update).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#error_on_nondeterministic_update Task#error_on_nondeterministic_update}
	ErrorOnNondeterministicUpdate interface{} `field:"optional" json:"errorOnNondeterministicUpdate" yaml:"errorOnNondeterministicUpdate"`
	// Specifies the name of a root task that the finalizer task is associated with.
	//
	// Finalizer tasks run after all other tasks in the task graph run to completion. You can define the SQL of a finalizer task to handle notifications and the release and cleanup of resources that a task graph uses. For more information, see [Release and cleanup of task graphs](https://docs.snowflake.com/en/user-guide/tasks-graphs.html#label-finalizer-task). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#finalize Task#finalize}
	Finalize *string `field:"optional" json:"finalize" yaml:"finalize"`
	// Display format for [GEOGRAPHY values](https://docs.snowflake.com/en/sql-reference/data-types-geospatial.html#label-data-types-geography). For more information, check [GEOGRAPHY_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#geography-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#geography_output_format Task#geography_output_format}
	GeographyOutputFormat *string `field:"optional" json:"geographyOutputFormat" yaml:"geographyOutputFormat"`
	// Display format for [GEOMETRY values](https://docs.snowflake.com/en/sql-reference/data-types-geospatial.html#label-data-types-geometry). For more information, check [GEOMETRY_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#geometry-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#geometry_output_format Task#geometry_output_format}
	GeometryOutputFormat *string `field:"optional" json:"geometryOutputFormat" yaml:"geometryOutputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#id Task#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies how JDBC processes TIMESTAMP_NTZ values. For more information, check [JDBC_TREAT_TIMESTAMP_NTZ_AS_UTC docs](https://docs.snowflake.com/en/sql-reference/parameters#jdbc-treat-timestamp-ntz-as-utc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#jdbc_treat_timestamp_ntz_as_utc Task#jdbc_treat_timestamp_ntz_as_utc}
	JdbcTreatTimestampNtzAsUtc interface{} `field:"optional" json:"jdbcTreatTimestampNtzAsUtc" yaml:"jdbcTreatTimestampNtzAsUtc"`
	// Specifies whether the JDBC Driver uses the time zone of the JVM or the time zone of the session (specified by the [TIMEZONE](https://docs.snowflake.com/en/sql-reference/parameters#label-timezone) parameter) for the getDate(), getTime(), and getTimestamp() methods of the ResultSet class. For more information, check [JDBC_USE_SESSION_TIMEZONE docs](https://docs.snowflake.com/en/sql-reference/parameters#jdbc-use-session-timezone).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#jdbc_use_session_timezone Task#jdbc_use_session_timezone}
	JdbcUseSessionTimezone interface{} `field:"optional" json:"jdbcUseSessionTimezone" yaml:"jdbcUseSessionTimezone"`
	// Specifies the number of blank spaces to indent each new element in JSON output in the session.
	//
	// Also specifies whether to insert newline characters after each element. For more information, check [JSON_INDENT docs](https://docs.snowflake.com/en/sql-reference/parameters#json-indent).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#json_indent Task#json_indent}
	JsonIndent *float64 `field:"optional" json:"jsonIndent" yaml:"jsonIndent"`
	// Number of seconds to wait while trying to lock a resource, before timing out and aborting the statement.
	//
	// For more information, check [LOCK_TIMEOUT docs](https://docs.snowflake.com/en/sql-reference/parameters#lock-timeout).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#lock_timeout Task#lock_timeout}
	LockTimeout *float64 `field:"optional" json:"lockTimeout" yaml:"lockTimeout"`
	// Specifies the severity level of messages that should be ingested and made available in the active event table.
	//
	// Messages at the specified level (and at more severe levels) are ingested. For more information about log levels, see [Setting log level](https://docs.snowflake.com/en/developer-guide/logging-tracing/logging-log-level). For more information, check [LOG_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#log-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#log_level Task#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Number of statements to execute when using the multi-statement capability. For more information, check [MULTI_STATEMENT_COUNT docs](https://docs.snowflake.com/en/sql-reference/parameters#multi-statement-count).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#multi_statement_count Task#multi_statement_count}
	MultiStatementCount *float64 `field:"optional" json:"multiStatementCount" yaml:"multiStatementCount"`
	// Specifies whether the ORDER or NOORDER property is set by default when you create a new sequence or add a new table column.
	//
	// The ORDER and NOORDER properties determine whether or not the values are generated for the sequence or auto-incremented column in [increasing or decreasing order](https://docs.snowflake.com/en/user-guide/querying-sequences.html#label-querying-sequences-increasing-values). For more information, check [NOORDER_SEQUENCE_AS_DEFAULT docs](https://docs.snowflake.com/en/sql-reference/parameters#noorder-sequence-as-default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#noorder_sequence_as_default Task#noorder_sequence_as_default}
	NoorderSequenceAsDefault interface{} `field:"optional" json:"noorderSequenceAsDefault" yaml:"noorderSequenceAsDefault"`
	// Specifies how ODBC processes columns that have a scale of zero (0). For more information, check [ODBC_TREAT_DECIMAL_AS_INT docs](https://docs.snowflake.com/en/sql-reference/parameters#odbc-treat-decimal-as-int).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#odbc_treat_decimal_as_int Task#odbc_treat_decimal_as_int}
	OdbcTreatDecimalAsInt interface{} `field:"optional" json:"odbcTreatDecimalAsInt" yaml:"odbcTreatDecimalAsInt"`
	// Optional string that can be used to tag queries and other SQL statements executed within a session.
	//
	// The tags are displayed in the output of the [QUERY_HISTORY, QUERY_HISTORY_BY_*](https://docs.snowflake.com/en/sql-reference/functions/query_history) functions. For more information, check [QUERY_TAG docs](https://docs.snowflake.com/en/sql-reference/parameters#query-tag).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#query_tag Task#query_tag}
	QueryTag *string `field:"optional" json:"queryTag" yaml:"queryTag"`
	// Specifies whether letters in double-quoted object identifiers are stored and resolved as uppercase letters.
	//
	// By default, Snowflake preserves the case of alphabetic characters when storing and resolving double-quoted identifiers (see [Identifier resolution](https://docs.snowflake.com/en/sql-reference/identifiers-syntax.html#label-identifier-casing)). You can use this parameter in situations in which [third-party applications always use double quotes around identifiers](https://docs.snowflake.com/en/sql-reference/identifiers-syntax.html#label-identifier-casing-parameter). For more information, check [QUOTED_IDENTIFIERS_IGNORE_CASE docs](https://docs.snowflake.com/en/sql-reference/parameters#quoted-identifiers-ignore-case).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#quoted_identifiers_ignore_case Task#quoted_identifiers_ignore_case}
	QuotedIdentifiersIgnoreCase interface{} `field:"optional" json:"quotedIdentifiersIgnoreCase" yaml:"quotedIdentifiersIgnoreCase"`
	// Specifies the maximum number of rows returned in a result set.
	//
	// A value of 0 specifies no maximum. For more information, check [ROWS_PER_RESULTSET docs](https://docs.snowflake.com/en/sql-reference/parameters#rows-per-resultset).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#rows_per_resultset Task#rows_per_resultset}
	RowsPerResultset *float64 `field:"optional" json:"rowsPerResultset" yaml:"rowsPerResultset"`
	// Specifies the DNS name of an Amazon S3 interface endpoint.
	//
	// Requests sent to the internal stage of an account via [AWS PrivateLink for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/privatelink-interface-endpoints.html) use this endpoint to connect. For more information, see [Accessing Internal stages with dedicated interface endpoints](https://docs.snowflake.com/en/user-guide/private-internal-stages-aws.html#label-aws-privatelink-internal-stage-network-isolation). For more information, check [S3_STAGE_VPCE_DNS_NAME docs](https://docs.snowflake.com/en/sql-reference/parameters#s3-stage-vpce-dns-name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#s3_stage_vpce_dns_name Task#s3_stage_vpce_dns_name}
	S3StageVpceDnsName *string `field:"optional" json:"s3StageVpceDnsName" yaml:"s3StageVpceDnsName"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#schedule Task#schedule}
	Schedule *TaskSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Specifies the path to search to resolve unqualified object names in queries.
	//
	// For more information, see [Name resolution in queries](https://docs.snowflake.com/en/sql-reference/name-resolution.html#label-object-name-resolution-search-path). Comma-separated list of identifiers. An identifier can be a fully or partially qualified schema name. For more information, check [SEARCH_PATH docs](https://docs.snowflake.com/en/sql-reference/parameters#search-path).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#search_path Task#search_path}
	SearchPath *string `field:"optional" json:"searchPath" yaml:"searchPath"`
	// Amount of time, in seconds, a SQL statement (query, DDL, DML, etc.) remains queued for a warehouse before it is canceled by the system. This parameter can be used in conjunction with the [MAX_CONCURRENCY_LEVEL](https://docs.snowflake.com/en/sql-reference/parameters#label-max-concurrency-level) parameter to ensure a warehouse is never backlogged. For more information, check [STATEMENT_QUEUED_TIMEOUT_IN_SECONDS docs](https://docs.snowflake.com/en/sql-reference/parameters#statement-queued-timeout-in-seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#statement_queued_timeout_in_seconds Task#statement_queued_timeout_in_seconds}
	StatementQueuedTimeoutInSeconds *float64 `field:"optional" json:"statementQueuedTimeoutInSeconds" yaml:"statementQueuedTimeoutInSeconds"`
	// Amount of time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system. For more information, check [STATEMENT_TIMEOUT_IN_SECONDS docs](https://docs.snowflake.com/en/sql-reference/parameters#statement-timeout-in-seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#statement_timeout_in_seconds Task#statement_timeout_in_seconds}
	StatementTimeoutInSeconds *float64 `field:"optional" json:"statementTimeoutInSeconds" yaml:"statementTimeoutInSeconds"`
	// This parameter specifies whether JSON output in a session is compatible with the general standard (as described by [http://json.org](http://json.org)). By design, Snowflake allows JSON input that contains non-standard values; however, these non-standard values might result in Snowflake outputting JSON that is incompatible with other platforms and languages. This parameter, when enabled, ensures that Snowflake outputs valid/compatible JSON. For more information, check [STRICT_JSON_OUTPUT docs](https://docs.snowflake.com/en/sql-reference/parameters#strict-json-output).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#strict_json_output Task#strict_json_output}
	StrictJsonOutput interface{} `field:"optional" json:"strictJsonOutput" yaml:"strictJsonOutput"`
	// Specifies the number of consecutive failed task runs after which the current task is suspended automatically.
	//
	// The default is 0 (no automatic suspension). For more information, check [SUSPEND_TASK_AFTER_NUM_FAILURES docs](https://docs.snowflake.com/en/sql-reference/parameters#suspend-task-after-num-failures).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#suspend_task_after_num_failures Task#suspend_task_after_num_failures}
	SuspendTaskAfterNumFailures *float64 `field:"optional" json:"suspendTaskAfterNumFailures" yaml:"suspendTaskAfterNumFailures"`
	// Specifies the number of automatic task graph retry attempts.
	//
	// If any task graphs complete in a FAILED state, Snowflake can automatically retry the task graphs from the last task in the graph that failed. For more information, check [TASK_AUTO_RETRY_ATTEMPTS docs](https://docs.snowflake.com/en/sql-reference/parameters#task-auto-retry-attempts).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#task_auto_retry_attempts Task#task_auto_retry_attempts}
	TaskAutoRetryAttempts *float64 `field:"optional" json:"taskAutoRetryAttempts" yaml:"taskAutoRetryAttempts"`
	// Specifies the input format for the TIME data type.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). Any valid, supported time format or AUTO (AUTO specifies that Snowflake attempts to automatically detect the format of times stored in the system during the session). For more information, check [TIME_INPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#time-input-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#time_input_format Task#time_input_format}
	TimeInputFormat *string `field:"optional" json:"timeInputFormat" yaml:"timeInputFormat"`
	// Specifies the display format for the TIME data type.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [TIME_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#time-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#time_output_format Task#time_output_format}
	TimeOutputFormat *string `field:"optional" json:"timeOutputFormat" yaml:"timeOutputFormat"`
	// Specifies whether the [DATEADD](https://docs.snowflake.com/en/sql-reference/functions/dateadd) function (and its aliases) always consider a day to be exactly 24 hours for expressions that span multiple days. For more information, check [TIMESTAMP_DAY_IS_ALWAYS_24H docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-day-is-always-24h).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#timestamp_day_is_always_24h Task#timestamp_day_is_always_24h}
	TimestampDayIsAlways24H interface{} `field:"optional" json:"timestampDayIsAlways24H" yaml:"timestampDayIsAlways24H"`
	// Specifies the input format for the TIMESTAMP data type alias.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). Any valid, supported timestamp format or AUTO (AUTO specifies that Snowflake attempts to automatically detect the format of timestamps stored in the system during the session). For more information, check [TIMESTAMP_INPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-input-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#timestamp_input_format Task#timestamp_input_format}
	TimestampInputFormat *string `field:"optional" json:"timestampInputFormat" yaml:"timestampInputFormat"`
	// Specifies the display format for the TIMESTAMP_LTZ data type.
	//
	// If no format is specified, defaults to [TIMESTAMP_OUTPUT_FORMAT](https://docs.snowflake.com/en/sql-reference/parameters#label-timestamp-output-format). For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [TIMESTAMP_LTZ_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-ltz-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#timestamp_ltz_output_format Task#timestamp_ltz_output_format}
	TimestampLtzOutputFormat *string `field:"optional" json:"timestampLtzOutputFormat" yaml:"timestampLtzOutputFormat"`
	// Specifies the display format for the TIMESTAMP_NTZ data type. For more information, check [TIMESTAMP_NTZ_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-ntz-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#timestamp_ntz_output_format Task#timestamp_ntz_output_format}
	TimestampNtzOutputFormat *string `field:"optional" json:"timestampNtzOutputFormat" yaml:"timestampNtzOutputFormat"`
	// Specifies the display format for the TIMESTAMP data type alias.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [TIMESTAMP_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#timestamp_output_format Task#timestamp_output_format}
	TimestampOutputFormat *string `field:"optional" json:"timestampOutputFormat" yaml:"timestampOutputFormat"`
	// Specifies the TIMESTAMP_* variation that the TIMESTAMP data type alias maps to. For more information, check [TIMESTAMP_TYPE_MAPPING docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-type-mapping).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#timestamp_type_mapping Task#timestamp_type_mapping}
	TimestampTypeMapping *string `field:"optional" json:"timestampTypeMapping" yaml:"timestampTypeMapping"`
	// Specifies the display format for the TIMESTAMP_TZ data type.
	//
	// If no format is specified, defaults to [TIMESTAMP_OUTPUT_FORMAT](https://docs.snowflake.com/en/sql-reference/parameters#label-timestamp-output-format). For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [TIMESTAMP_TZ_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-tz-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#timestamp_tz_output_format Task#timestamp_tz_output_format}
	TimestampTzOutputFormat *string `field:"optional" json:"timestampTzOutputFormat" yaml:"timestampTzOutputFormat"`
	// Specifies the time zone for the session.
	//
	// You can specify a [time zone name](https://data.iana.org/time-zones/tzdb-2021a/zone1970.tab) or a [link name](https://data.iana.org/time-zones/tzdb-2021a/backward) from release 2021a of the [IANA Time Zone Database](https://www.iana.org/time-zones) (e.g. America/Los_Angeles, Europe/London, UTC, Etc/GMT, etc.). For more information, check [TIMEZONE docs](https://docs.snowflake.com/en/sql-reference/parameters#timezone).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#timezone Task#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Controls how trace events are ingested into the event table.
	//
	// For more information about trace levels, see [Setting trace level](https://docs.snowflake.com/en/developer-guide/logging-tracing/tracing-trace-level). For more information, check [TRACE_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#trace-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#trace_level Task#trace_level}
	TraceLevel *string `field:"optional" json:"traceLevel" yaml:"traceLevel"`
	// Specifies the action to perform when a statement issued within a non-autocommit transaction returns with an error.
	//
	// For more information, check [TRANSACTION_ABORT_ON_ERROR docs](https://docs.snowflake.com/en/sql-reference/parameters#transaction-abort-on-error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#transaction_abort_on_error Task#transaction_abort_on_error}
	TransactionAbortOnError interface{} `field:"optional" json:"transactionAbortOnError" yaml:"transactionAbortOnError"`
	// Specifies the isolation level for transactions in the user session. For more information, check [TRANSACTION_DEFAULT_ISOLATION_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#transaction-default-isolation-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#transaction_default_isolation_level Task#transaction_default_isolation_level}
	TransactionDefaultIsolationLevel *string `field:"optional" json:"transactionDefaultIsolationLevel" yaml:"transactionDefaultIsolationLevel"`
	// Specifies the “century start” year for 2-digit years (i.e. the earliest year such dates can represent). This parameter prevents ambiguous dates when importing or converting data with the `YY` date format component (i.e. years represented as 2 digits). For more information, check [TWO_DIGIT_CENTURY_START docs](https://docs.snowflake.com/en/sql-reference/parameters#two-digit-century-start).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#two_digit_century_start Task#two_digit_century_start}
	TwoDigitCenturyStart *float64 `field:"optional" json:"twoDigitCenturyStart" yaml:"twoDigitCenturyStart"`
	// Determines if an unsupported (i.e. non-default) value specified for a constraint property returns an error. For more information, check [UNSUPPORTED_DDL_ACTION docs](https://docs.snowflake.com/en/sql-reference/parameters#unsupported-ddl-action).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#unsupported_ddl_action Task#unsupported_ddl_action}
	UnsupportedDdlAction *string `field:"optional" json:"unsupportedDdlAction" yaml:"unsupportedDdlAction"`
	// Specifies whether to reuse persisted query results, if available, when a matching query is submitted.
	//
	// For more information, check [USE_CACHED_RESULT docs](https://docs.snowflake.com/en/sql-reference/parameters#use-cached-result).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#use_cached_result Task#use_cached_result}
	UseCachedResult interface{} `field:"optional" json:"useCachedResult" yaml:"useCachedResult"`
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size.
	//
	// Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. Valid values are (case-insensitive): %s. (Conflicts with warehouse). For more information about warehouses, see [docs](./warehouse). For more information, check [USER_TASK_MANAGED_INITIAL_WAREHOUSE_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-managed-initial-warehouse-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#user_task_managed_initial_warehouse_size Task#user_task_managed_initial_warehouse_size}
	UserTaskManagedInitialWarehouseSize *string `field:"optional" json:"userTaskManagedInitialWarehouseSize" yaml:"userTaskManagedInitialWarehouseSize"`
	// Minimum amount of time between Triggered Task executions in seconds For more information, check [USER_TASK_MINIMUM_TRIGGER_INTERVAL_IN_SECONDS docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-minimum-trigger-interval-in-seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#user_task_minimum_trigger_interval_in_seconds Task#user_task_minimum_trigger_interval_in_seconds}
	UserTaskMinimumTriggerIntervalInSeconds *float64 `field:"optional" json:"userTaskMinimumTriggerIntervalInSeconds" yaml:"userTaskMinimumTriggerIntervalInSeconds"`
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	//
	// For more information, check [USER_TASK_TIMEOUT_MS docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-timeout-ms).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#user_task_timeout_ms Task#user_task_timeout_ms}
	UserTaskTimeoutMs *float64 `field:"optional" json:"userTaskTimeoutMs" yaml:"userTaskTimeoutMs"`
	// The warehouse the task will use.
	//
	// Omit this parameter to use Snowflake-managed compute resources for runs of this task. Due to Snowflake limitations warehouse identifier can consist of only upper-cased letters. (Conflicts with user_task_managed_initial_warehouse_size) For more information about this resource, see [docs](./warehouse).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#warehouse Task#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
	// Specifies how the weeks in a given year are computed.
	//
	// `0`: The semantics used are equivalent to the ISO semantics, in which a week belongs to a given year if at least 4 days of that week are in that year. `1`: January 1 is included in the first week of the year and December 31 is included in the last week of the year. For more information, check [WEEK_OF_YEAR_POLICY docs](https://docs.snowflake.com/en/sql-reference/parameters#week-of-year-policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#week_of_year_policy Task#week_of_year_policy}
	WeekOfYearPolicy *float64 `field:"optional" json:"weekOfYearPolicy" yaml:"weekOfYearPolicy"`
	// Specifies the first day of the week (used by week-related date functions).
	//
	// `0`: Legacy Snowflake behavior is used (i.e. ISO-like semantics). `1` (Monday) to `7` (Sunday): All the week-related functions use weeks that start on the specified day of the week. For more information, check [WEEK_START docs](https://docs.snowflake.com/en/sql-reference/parameters#week-start).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#week_start Task#week_start}
	WeekStart *float64 `field:"optional" json:"weekStart" yaml:"weekStart"`
	// Specifies a Boolean SQL expression;
	//
	// multiple conditions joined with AND/OR are supported. When a task is triggered (based on its SCHEDULE or AFTER setting), it validates the conditions of the expression to determine whether to execute. If the conditions of the expression are not met, then the task skips the current run. Any tasks that identify this task as a predecessor also don’t run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/task#when Task#when}
	When *string `field:"optional" json:"when" yaml:"when"`
}

