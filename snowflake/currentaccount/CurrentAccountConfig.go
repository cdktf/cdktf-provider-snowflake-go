// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package currentaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CurrentAccountConfig struct {
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
	// Specifies the action that Snowflake performs for in-progress queries if connectivity is lost due to abrupt termination of a session (e.g. network outage, browser termination, service interruption). For more information, check [ABORT_DETACHED_QUERY docs](https://docs.snowflake.com/en/sql-reference/parameters#abort-detached-query).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#abort_detached_query CurrentAccount#abort_detached_query}
	AbortDetachedQuery interface{} `field:"optional" json:"abortDetachedQuery" yaml:"abortDetachedQuery"`
	// Sets the profiler to use for the session when [profiling Python handler code](https://docs.snowflake.com/en/developer-guide/stored-procedure/python/procedure-python-profiler). Valid values are (case-insensitive): `LINE` | `MEMORY`. For more information, check [ACTIVE_PYTHON_PROFILER docs](https://docs.snowflake.com/en/sql-reference/parameters#active-python-profiler).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#active_python_profiler CurrentAccount#active_python_profiler}
	ActivePythonProfiler *string `field:"optional" json:"activePythonProfiler" yaml:"activePythonProfiler"`
	// Specifies whether an MFA token can be saved in the client-side operating system keystore to promote continuous, secure connectivity without users needing to respond to an MFA prompt at the start of each connection attempt to Snowflake.
	//
	// For details and the list of supported Snowflake-provided clients, see [Using MFA token caching to minimize the number of prompts during authentication — optional.](https://docs.snowflake.com/en/user-guide/security-mfa.html#label-mfa-token-caching) For more information, check [ALLOW_CLIENT_MFA_CACHING docs](https://docs.snowflake.com/en/sql-reference/parameters#allow-client-mfa-caching).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#allow_client_mfa_caching CurrentAccount#allow_client_mfa_caching}
	AllowClientMfaCaching interface{} `field:"optional" json:"allowClientMfaCaching" yaml:"allowClientMfaCaching"`
	// Specifies whether a connection token can be saved in the client-side operating system keystore to promote continuous, secure connectivity without users needing to enter login credentials at the start of each connection attempt to Snowflake.
	//
	// For details and the list of supported Snowflake-provided clients, see [Using connection caching to minimize the number of prompts for authentication — optional.](https://docs.snowflake.com/en/user-guide/admin-security-fed-auth-use.html#label-browser-based-sso-connection-caching) For more information, check [ALLOW_ID_TOKEN docs](https://docs.snowflake.com/en/sql-reference/parameters#allow-id-token).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#allow_id_token CurrentAccount#allow_id_token}
	AllowIdToken interface{} `field:"optional" json:"allowIdToken" yaml:"allowIdToken"`
	// Specifies [authentication policy](https://docs.snowflake.com/en/user-guide/authentication-policies) for the current account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#authentication_policy CurrentAccount#authentication_policy}
	AuthenticationPolicy *string `field:"optional" json:"authenticationPolicy" yaml:"authenticationPolicy"`
	// Specifies whether autocommit is enabled for the session.
	//
	// Autocommit determines whether a DML statement, when executed without an active transaction, is automatically committed after the statement successfully completes. For more information, see [Transactions](https://docs.snowflake.com/en/sql-reference/transactions). For more information, check [AUTOCOMMIT docs](https://docs.snowflake.com/en/sql-reference/parameters#autocommit).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#autocommit CurrentAccount#autocommit}
	Autocommit interface{} `field:"optional" json:"autocommit" yaml:"autocommit"`
	// Specifies a prefix for Snowflake to use in the write path for Snowflake-managed Apache Iceberg™ tables.
	//
	// For more information, see [data and metadata directories for Iceberg tables](https://docs.snowflake.com/en/user-guide/tables-iceberg-storage.html#label-tables-iceberg-configure-external-volume-base-location). For more information, check [BASE_LOCATION_PREFIX docs](https://docs.snowflake.com/en/sql-reference/parameters#base-location-prefix).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#base_location_prefix CurrentAccount#base_location_prefix}
	BaseLocationPrefix *string `field:"optional" json:"baseLocationPrefix" yaml:"baseLocationPrefix"`
	// The format of VARCHAR values passed as input to VARCHAR-to-BINARY conversion functions.
	//
	// For more information, see [Binary input and output](https://docs.snowflake.com/en/sql-reference/binary-input-output). Valid values are (case-insensitive): `HEX` | `BASE64` | `UTF8`. For more information, check [BINARY_INPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#binary-input-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#binary_input_format CurrentAccount#binary_input_format}
	BinaryInputFormat *string `field:"optional" json:"binaryInputFormat" yaml:"binaryInputFormat"`
	// The format for VARCHAR values returned as output by BINARY-to-VARCHAR conversion functions.
	//
	// For more information, see [Binary input and output](https://docs.snowflake.com/en/sql-reference/binary-input-output). Valid values are (case-insensitive): `HEX` | `BASE64`. For more information, check [BINARY_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#binary-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#binary_output_format CurrentAccount#binary_output_format}
	BinaryOutputFormat *string `field:"optional" json:"binaryOutputFormat" yaml:"binaryOutputFormat"`
	// Specifies the catalog for Apache Iceberg™ tables.
	//
	// For more information, see the [Iceberg table documentation](https://docs.snowflake.com/en/user-guide/tables-iceberg.html#label-tables-iceberg-catalog-def). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information, check [CATALOG docs](https://docs.snowflake.com/en/sql-reference/parameters#catalog).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#catalog CurrentAccount#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Specifies the name of your catalog integration for [Snowflake Open Catalog](https://other-docs.snowflake.com/en/opencatalog/overview). Snowflake syncs tables that use the specified catalog integration with your Snowflake Open Catalog account. For more information, see [Sync a Snowflake-managed table with Snowflake Open Catalog](https://docs.snowflake.com/en/user-guide/tables-iceberg-open-catalog-sync). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information, check [CATALOG_SYNC docs](https://docs.snowflake.com/en/sql-reference/parameters#catalog-sync).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#catalog_sync CurrentAccount#catalog_sync}
	CatalogSync *string `field:"optional" json:"catalogSync" yaml:"catalogSync"`
	// Enables users to log the data values bound to [PreparedStatements](https://docs.snowflake.com/en/developer-guide/jdbc/jdbc-api.html#label-jdbc-api-preparedstatement) ([more details](https://docs.snowflake.com/en/sql-reference/parameters#client-enable-log-info-statement-parameters)). For more information, check [CLIENT_ENABLE_LOG_INFO_STATEMENT_PARAMETERS docs](https://docs.snowflake.com/en/sql-reference/parameters#client-enable-log-info-statement-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_enable_log_info_statement_parameters CurrentAccount#client_enable_log_info_statement_parameters}
	ClientEnableLogInfoStatementParameters interface{} `field:"optional" json:"clientEnableLogInfoStatementParameters" yaml:"clientEnableLogInfoStatementParameters"`
	// Specifies the AES encryption key size, in bits, used by Snowflake to encrypt/decrypt files stored on internal stages (for loading/unloading data) when you use the SNOWFLAKE_FULL encryption type.
	//
	// For more information, check [CLIENT_ENCRYPTION_KEY_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#client-encryption-key-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_encryption_key_size CurrentAccount#client_encryption_key_size}
	ClientEncryptionKeySize *float64 `field:"optional" json:"clientEncryptionKeySize" yaml:"clientEncryptionKeySize"`
	// Parameter that specifies the maximum amount of memory the JDBC driver or ODBC driver should use for the result set from queries (in MB).
	//
	// For more information, check [CLIENT_MEMORY_LIMIT docs](https://docs.snowflake.com/en/sql-reference/parameters#client-memory-limit).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_memory_limit CurrentAccount#client_memory_limit}
	ClientMemoryLimit *float64 `field:"optional" json:"clientMemoryLimit" yaml:"clientMemoryLimit"`
	// For specific ODBC functions and JDBC methods, this parameter can change the default search scope from all databases/schemas to the current database/schema.
	//
	// The narrower search typically returns fewer rows and executes more quickly. For more information, check [CLIENT_METADATA_REQUEST_USE_CONNECTION_CTX docs](https://docs.snowflake.com/en/sql-reference/parameters#client-metadata-request-use-connection-ctx).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_metadata_request_use_connection_ctx CurrentAccount#client_metadata_request_use_connection_ctx}
	ClientMetadataRequestUseConnectionCtx interface{} `field:"optional" json:"clientMetadataRequestUseConnectionCtx" yaml:"clientMetadataRequestUseConnectionCtx"`
	// For specific ODBC functions and JDBC methods, this parameter can change the default search scope from all databases to the current database.
	//
	// The narrower search typically returns fewer rows and executes more quickly ([more details on the usage](https://docs.snowflake.com/en/sql-reference/parameters#client-metadata-use-session-database)). For more information, check [CLIENT_METADATA_USE_SESSION_DATABASE docs](https://docs.snowflake.com/en/sql-reference/parameters#client-metadata-use-session-database).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_metadata_use_session_database CurrentAccount#client_metadata_use_session_database}
	ClientMetadataUseSessionDatabase interface{} `field:"optional" json:"clientMetadataUseSessionDatabase" yaml:"clientMetadataUseSessionDatabase"`
	// Parameter that specifies the number of threads used by the client to pre-fetch large result sets.
	//
	// The driver will attempt to honor the parameter value, but defines the minimum and maximum values (depending on your system’s resources) to improve performance. For more information, check [CLIENT_PREFETCH_THREADS docs](https://docs.snowflake.com/en/sql-reference/parameters#client-prefetch-threads).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_prefetch_threads CurrentAccount#client_prefetch_threads}
	ClientPrefetchThreads *float64 `field:"optional" json:"clientPrefetchThreads" yaml:"clientPrefetchThreads"`
	// Parameter that specifies the maximum size of each set (or chunk) of query results to download (in MB).
	//
	// The JDBC driver downloads query results in chunks. For more information, check [CLIENT_RESULT_CHUNK_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#client-result-chunk-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_result_chunk_size CurrentAccount#client_result_chunk_size}
	ClientResultChunkSize *float64 `field:"optional" json:"clientResultChunkSize" yaml:"clientResultChunkSize"`
	// Parameter that indicates whether to match column name case-insensitively in ResultSet.get* methods in JDBC. For more information, check [CLIENT_RESULT_COLUMN_CASE_INSENSITIVE docs](https://docs.snowflake.com/en/sql-reference/parameters#client-result-column-case-insensitive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_result_column_case_insensitive CurrentAccount#client_result_column_case_insensitive}
	ClientResultColumnCaseInsensitive interface{} `field:"optional" json:"clientResultColumnCaseInsensitive" yaml:"clientResultColumnCaseInsensitive"`
	// Parameter that indicates whether to force a user to log in again after a period of inactivity in the session.
	//
	// For more information, check [CLIENT_SESSION_KEEP_ALIVE docs](https://docs.snowflake.com/en/sql-reference/parameters#client-session-keep-alive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_session_keep_alive CurrentAccount#client_session_keep_alive}
	ClientSessionKeepAlive interface{} `field:"optional" json:"clientSessionKeepAlive" yaml:"clientSessionKeepAlive"`
	// Number of seconds in-between client attempts to update the token for the session. For more information, check [CLIENT_SESSION_KEEP_ALIVE_HEARTBEAT_FREQUENCY docs](https://docs.snowflake.com/en/sql-reference/parameters#client-session-keep-alive-heartbeat-frequency).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_session_keep_alive_heartbeat_frequency CurrentAccount#client_session_keep_alive_heartbeat_frequency}
	ClientSessionKeepAliveHeartbeatFrequency *float64 `field:"optional" json:"clientSessionKeepAliveHeartbeatFrequency" yaml:"clientSessionKeepAliveHeartbeatFrequency"`
	// Specifies the [TIMESTAMP_* variation](https://docs.snowflake.com/en/sql-reference/data-types-datetime.html#label-datatypes-timestamp-variations) to use when binding timestamp variables for JDBC or ODBC applications that use the bind API to load data. Valid values are (case-insensitive): `TIMESTAMP_LTZ` | `TIMESTAMP_NTZ`. For more information, check [CLIENT_TIMESTAMP_TYPE_MAPPING docs](https://docs.snowflake.com/en/sql-reference/parameters#client-timestamp-type-mapping).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#client_timestamp_type_mapping CurrentAccount#client_timestamp_type_mapping}
	ClientTimestampTypeMapping *string `field:"optional" json:"clientTimestampTypeMapping" yaml:"clientTimestampTypeMapping"`
	// Specifies the regions where an inference request may be processed in case the request cannot be processed in the region where request is originally placed.
	//
	// Specifying DISABLED disables cross-region inferencing. For examples and details, see [Cross-region inference](https://docs.snowflake.com/en/user-guide/snowflake-cortex/cross-region-inference). For more information, check [CORTEX_ENABLED_CROSS_REGION docs](https://docs.snowflake.com/en/sql-reference/parameters#cortex-enabled-cross-region).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#cortex_enabled_cross_region CurrentAccount#cortex_enabled_cross_region}
	CortexEnabledCrossRegion *string `field:"optional" json:"cortexEnabledCrossRegion" yaml:"cortexEnabledCrossRegion"`
	// Specifies the models that users in the account can access.
	//
	// Use this parameter to allowlist models for all users in the account. If you need to provide specific users with access beyond what you’ve specified in the allowlist, use role-based access control instead. For more information, see [Model allowlist](https://docs.snowflake.com/en/user-guide/snowflake-cortex/aisql.html#label-cortex-llm-allowlist). For more information, check [CORTEX_MODELS_ALLOWLIST docs](https://docs.snowflake.com/en/sql-reference/parameters#cortex-models-allowlist).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#cortex_models_allowlist CurrentAccount#cortex_models_allowlist}
	CortexModelsAllowlist *string `field:"optional" json:"cortexModelsAllowlist" yaml:"cortexModelsAllowlist"`
	// Specifies the format for TIMESTAMP values in CSV files downloaded from Snowsight.
	//
	// If this parameter is not set, [TIMESTAMP_LTZ_OUTPUT_FORMAT](https://docs.snowflake.com/en/sql-reference/parameters#label-timestamp-ltz-output-format) will be used for TIMESTAMP_LTZ values, [TIMESTAMP_TZ_OUTPUT_FORMAT](https://docs.snowflake.com/en/sql-reference/parameters#label-timestamp-tz-output-format) will be used for TIMESTAMP_TZ and [TIMESTAMP_NTZ_OUTPUT_FORMAT](https://docs.snowflake.com/en/sql-reference/parameters#label-timestamp-ntz-output-format) for TIMESTAMP_NTZ values. For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output) or [Download your query results](https://docs.snowflake.com/en/user-guide/ui-snowsight-query.html#label-snowsight-download-query-results). For more information, check [CSV_TIMESTAMP_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#csv-timestamp-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#csv_timestamp_format CurrentAccount#csv_timestamp_format}
	CsvTimestampFormat *string `field:"optional" json:"csvTimestampFormat" yaml:"csvTimestampFormat"`
	// Number of days for which Snowflake retains historical data for performing Time Travel actions (SELECT, CLONE, UNDROP) on the object.
	//
	// A value of 0 effectively disables Time Travel for the specified database, schema, or table. For more information, see [Understanding & using Time Travel](https://docs.snowflake.com/en/user-guide/data-time-travel). For more information, check [DATA_RETENTION_TIME_IN_DAYS docs](https://docs.snowflake.com/en/sql-reference/parameters#data-retention-time-in-days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#data_retention_time_in_days CurrentAccount#data_retention_time_in_days}
	DataRetentionTimeInDays *float64 `field:"optional" json:"dataRetentionTimeInDays" yaml:"dataRetentionTimeInDays"`
	// Specifies the input format for the DATE data type.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [DATE_INPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#date-input-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#date_input_format CurrentAccount#date_input_format}
	DateInputFormat *string `field:"optional" json:"dateInputFormat" yaml:"dateInputFormat"`
	// Specifies the display format for the DATE data type.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [DATE_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#date-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#date_output_format CurrentAccount#date_output_format}
	DateOutputFormat *string `field:"optional" json:"dateOutputFormat" yaml:"dateOutputFormat"`
	// Sets the default collation used for the following DDL operations: [CREATE TABLE](https://docs.snowflake.com/en/sql-reference/sql/create-table), [ALTER TABLE](https://docs.snowflake.com/en/sql-reference/sql/alter-table) … ADD COLUMN. Setting this parameter forces all subsequently-created columns in the affected objects (table, schema, database, or account) to have the specified collation as the default, unless the collation for the column is explicitly defined in the DDL. For more information, check [DEFAULT_DDL_COLLATION docs](https://docs.snowflake.com/en/sql-reference/parameters#default-ddl-collation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#default_ddl_collation CurrentAccount#default_ddl_collation}
	DefaultDdlCollation *string `field:"optional" json:"defaultDdlCollation" yaml:"defaultDdlCollation"`
	// Sets the preferred CPU compute pool used for [Notebooks on CPU Container Runtime](https://docs.snowflake.com/en/developer-guide/snowflake-ml/notebooks-on-spcs). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information, check [DEFAULT_NOTEBOOK_COMPUTE_POOL_CPU docs](https://docs.snowflake.com/en/sql-reference/parameters#default-notebook-compute-pool-cpu).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#default_notebook_compute_pool_cpu CurrentAccount#default_notebook_compute_pool_cpu}
	DefaultNotebookComputePoolCpu *string `field:"optional" json:"defaultNotebookComputePoolCpu" yaml:"defaultNotebookComputePoolCpu"`
	// Sets the preferred GPU compute pool used for [Notebooks on GPU Container Runtime](https://docs.snowflake.com/en/developer-guide/snowflake-ml/notebooks-on-spcs). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information, check [DEFAULT_NOTEBOOK_COMPUTE_POOL_GPU docs](https://docs.snowflake.com/en/sql-reference/parameters#default-notebook-compute-pool-gpu).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#default_notebook_compute_pool_gpu CurrentAccount#default_notebook_compute_pool_gpu}
	DefaultNotebookComputePoolGpu *string `field:"optional" json:"defaultNotebookComputePoolGpu" yaml:"defaultNotebookComputePoolGpu"`
	// Specifies the default ordering of NULL values in a result set ([more details](https://docs.snowflake.com/en/sql-reference/parameters#default-null-ordering)). Valid values are (case-insensitive): `FIRST` | `LAST`. For more information, check [DEFAULT_NULL_ORDERING docs](https://docs.snowflake.com/en/sql-reference/parameters#default-null-ordering).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#default_null_ordering CurrentAccount#default_null_ordering}
	DefaultNullOrdering *string `field:"optional" json:"defaultNullOrdering" yaml:"defaultNullOrdering"`
	// Specifies the name of the default warehouse to use when creating a notebook.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information, check [DEFAULT_STREAMLIT_NOTEBOOK_WAREHOUSE docs](https://docs.snowflake.com/en/sql-reference/parameters#default-streamlit-notebook-warehouse).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#default_streamlit_notebook_warehouse CurrentAccount#default_streamlit_notebook_warehouse}
	DefaultStreamlitNotebookWarehouse *string `field:"optional" json:"defaultStreamlitNotebookWarehouse" yaml:"defaultStreamlitNotebookWarehouse"`
	// Controls whether users in an account see a button to download data in Snowsight or the Classic Console, such as a table returned from running a query in a worksheet.
	//
	// If the button to download is hidden in Snowsight or the Classic Console, users can still download or export data using [third-party software](https://docs.snowflake.com/en/user-guide/ecosystem). For more information, check [DISABLE_UI_DOWNLOAD_BUTTON docs](https://docs.snowflake.com/en/sql-reference/parameters#disable-ui-download-button).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#disable_ui_download_button CurrentAccount#disable_ui_download_button}
	DisableUiDownloadButton interface{} `field:"optional" json:"disableUiDownloadButton" yaml:"disableUiDownloadButton"`
	// Controls whether users in an account can grant privileges directly to other users.
	//
	// Disabling user privilege grants (that is, setting DISABLE_USER_PRIVILEGE_GRANTS to TRUE) does not affect existing grants to users. Existing grants to users continue to confer privileges to those users. For more information, see [GRANT <privileges> … TO USER](https://docs.snowflake.com/en/sql-reference/sql/grant-privilege-user). For more information, check [DISABLE_USER_PRIVILEGE_GRANTS docs](https://docs.snowflake.com/en/sql-reference/parameters#disable-user-privilege-grants).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#disable_user_privilege_grants CurrentAccount#disable_user_privilege_grants}
	DisableUserPrivilegeGrants interface{} `field:"optional" json:"disableUserPrivilegeGrants" yaml:"disableUserPrivilegeGrants"`
	// Controls whether events from [automatic sensitive data classification](https://docs.snowflake.com/en/user-guide/classify-auto) are logged in the user event table. For more information, check [ENABLE_AUTOMATIC_SENSITIVE_DATA_CLASSIFICATION_LOG docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-automatic-sensitive-data-classification-log).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_automatic_sensitive_data_classification_log CurrentAccount#enable_automatic_sensitive_data_classification_log}
	EnableAutomaticSensitiveDataClassificationLog interface{} `field:"optional" json:"enableAutomaticSensitiveDataClassificationLog" yaml:"enableAutomaticSensitiveDataClassificationLog"`
	// Enables or disables the Listing Cross-cloud auto-fulfillment Egress cost optimizer. For more information, check [ENABLE_EGRESS_COST_OPTIMIZER docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-egress-cost-optimizer).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_egress_cost_optimizer CurrentAccount#enable_egress_cost_optimizer}
	EnableEgressCostOptimizer interface{} `field:"optional" json:"enableEgressCostOptimizer" yaml:"enableEgressCostOptimizer"`
	// Determines the login flow for users.
	//
	// When enabled, Snowflake prompts users for their username or email address before presenting authentication methods. For details, see [Identifier-first login](https://docs.snowflake.com/en/user-guide/identifier-first-login). For more information, check [ENABLE_IDENTIFIER_FIRST_LOGIN docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-identifier-first-login).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_identifier_first_login CurrentAccount#enable_identifier_first_login}
	EnableIdentifierFirstLogin interface{} `field:"optional" json:"enableIdentifierFirstLogin" yaml:"enableIdentifierFirstLogin"`
	// Specifies whether the [SYSTEM$GET_PRIVATELINK_CONFIG](https://docs.snowflake.com/en/sql-reference/functions/system_get_privatelink_config) function returns the private-internal-stages key in the query result. The corresponding value in the query result is used during the configuration process for private connectivity to internal stages. For more information, check [ENABLE_INTERNAL_STAGES_PRIVATELINK docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-internal-stages-privatelink).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_internal_stages_privatelink CurrentAccount#enable_internal_stages_privatelink}
	EnableInternalStagesPrivatelink interface{} `field:"optional" json:"enableInternalStagesPrivatelink" yaml:"enableInternalStagesPrivatelink"`
	// Specifies choice for the [image repository](https://docs.snowflake.com/en/developer-guide/snowpark-container-services/working-with-registry-repository.html#label-registry-and-repository-image-repository) to opt out of Tri-Secret Secure and [Periodic rekeying](https://docs.snowflake.com/en/user-guide/security-encryption-manage.html#label-periodic-rekeying). For more information, check [ENABLE_TRI_SECRET_AND_REKEY_OPT_OUT_FOR_IMAGE_REPOSITORY docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-tri-secret-and-rekey-opt-out-for-image-repository).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_tri_secret_and_rekey_opt_out_for_image_repository CurrentAccount#enable_tri_secret_and_rekey_opt_out_for_image_repository}
	EnableTriSecretAndRekeyOptOutForImageRepository interface{} `field:"optional" json:"enableTriSecretAndRekeyOptOutForImageRepository" yaml:"enableTriSecretAndRekeyOptOutForImageRepository"`
	// Specifies the choice for the [Snowpark Container Services block storage volume](https://docs.snowflake.com/en/developer-guide/snowpark-container-services/block-storage-volume) to opt out of Tri-Secret Secure and [Periodic rekeying](https://docs.snowflake.com/en/user-guide/security-encryption-manage.html#label-periodic-rekeying). For more information, check [ENABLE_TRI_SECRET_AND_REKEY_OPT_OUT_FOR_SPCS_BLOCK_STORAGE docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-tri-secret-and-rekey-opt-out-for-spcs-block-storage).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_tri_secret_and_rekey_opt_out_for_spcs_block_storage CurrentAccount#enable_tri_secret_and_rekey_opt_out_for_spcs_block_storage}
	EnableTriSecretAndRekeyOptOutForSpcsBlockStorage interface{} `field:"optional" json:"enableTriSecretAndRekeyOptOutForSpcsBlockStorage" yaml:"enableTriSecretAndRekeyOptOutForSpcsBlockStorage"`
	// Specifies whether Snowflake may capture – in an event table – log messages or trace event data for unhandled exceptions in procedure or UDF handler code.
	//
	// For more information, see [Capturing messages from unhandled exceptions](https://docs.snowflake.com/en/developer-guide/logging-tracing/unhandled-exception-messages). For more information, check [ENABLE_UNHANDLED_EXCEPTIONS_REPORTING docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-unhandled-exceptions-reporting).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_unhandled_exceptions_reporting CurrentAccount#enable_unhandled_exceptions_reporting}
	EnableUnhandledExceptionsReporting interface{} `field:"optional" json:"enableUnhandledExceptionsReporting" yaml:"enableUnhandledExceptionsReporting"`
	// Specifies whether to set the schema for unloaded Parquet files based on the logical column data types (i.e. the types in the unload SQL query or source table) or on the unloaded column values (i.e. the smallest data types and precision that support the values in the output columns of the unload SQL statement or source table). For more information, check [ENABLE_UNLOAD_PHYSICAL_TYPE_OPTIMIZATION docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-unload-physical-type-optimization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_unload_physical_type_optimization CurrentAccount#enable_unload_physical_type_optimization}
	EnableUnloadPhysicalTypeOptimization interface{} `field:"optional" json:"enableUnloadPhysicalTypeOptimization" yaml:"enableUnloadPhysicalTypeOptimization"`
	// Controls whether query text is redacted if a SQL query fails due to a syntax or parsing error.
	//
	// If FALSE, the content of a failed query is redacted in the views, pages, and functions that provide a query history. Only users with a role that is granted or inherits the AUDIT privilege can set the ENABLE_UNREDACTED_QUERY_SYNTAX_ERROR parameter. When using the ALTER USER command to set the parameter to TRUE for a particular user, modify the user that you want to see the query text, not the user who executed the query (if those are different users). For more information, check [ENABLE_UNREDACTED_QUERY_SYNTAX_ERROR docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-unredacted-query-syntax-error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_unredacted_query_syntax_error CurrentAccount#enable_unredacted_query_syntax_error}
	EnableUnredactedQuerySyntaxError interface{} `field:"optional" json:"enableUnredactedQuerySyntaxError" yaml:"enableUnredactedQuerySyntaxError"`
	// Controls whether error messages related to secure objects are redacted in metadata.
	//
	// For more information, see [Secure objects: Redaction of information in error messages](https://docs.snowflake.com/en/release-notes/bcr-bundles/un-bundled/bcr-1858). Only users with a role that is granted or inherits the AUDIT privilege can set the ENABLE_UNREDACTED_SECURE_OBJECT_ERROR parameter. When using the ALTER USER command to set the parameter to TRUE for a particular user, modify the user that you want to see the redacted error messages in metadata, not the user who caused the error. For more information, check [ENABLE_UNREDACTED_SECURE_OBJECT_ERROR docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-unredacted-secure-object-error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enable_unredacted_secure_object_error CurrentAccount#enable_unredacted_secure_object_error}
	EnableUnredactedSecureObjectError interface{} `field:"optional" json:"enableUnredactedSecureObjectError" yaml:"enableUnredactedSecureObjectError"`
	// Specifies whether a network policy that uses network rules can restrict access to AWS internal stages.
	//
	// This parameter has no effect on network policies that do not use network rules. This account-level parameter affects both account-level and user-level network policies. For details about using network policies and network rules to restrict access to AWS internal stages, including the use of this parameter, see [Protecting internal stages on AWS](https://docs.snowflake.com/en/user-guide/network-policies.html#label-network-policies-rules-stages). For more information, check [ENFORCE_NETWORK_RULES_FOR_INTERNAL_STAGES docs](https://docs.snowflake.com/en/sql-reference/parameters#enforce-network-rules-for-internal-stages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#enforce_network_rules_for_internal_stages CurrentAccount#enforce_network_rules_for_internal_stages}
	EnforceNetworkRulesForInternalStages interface{} `field:"optional" json:"enforceNetworkRulesForInternalStages" yaml:"enforceNetworkRulesForInternalStages"`
	// Specifies whether to return an error when the [MERGE](https://docs.snowflake.com/en/sql-reference/sql/merge) command is used to update or delete a target row that joins multiple source rows and the system cannot determine the action to perform on the target row. For more information, check [ERROR_ON_NONDETERMINISTIC_MERGE docs](https://docs.snowflake.com/en/sql-reference/parameters#error-on-nondeterministic-merge).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#error_on_nondeterministic_merge CurrentAccount#error_on_nondeterministic_merge}
	ErrorOnNondeterministicMerge interface{} `field:"optional" json:"errorOnNondeterministicMerge" yaml:"errorOnNondeterministicMerge"`
	// Specifies whether to return an error when the [UPDATE](https://docs.snowflake.com/en/sql-reference/sql/update) command is used to update a target row that joins multiple source rows and the system cannot determine the action to perform on the target row. For more information, check [ERROR_ON_NONDETERMINISTIC_UPDATE docs](https://docs.snowflake.com/en/sql-reference/parameters#error-on-nondeterministic-update).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#error_on_nondeterministic_update CurrentAccount#error_on_nondeterministic_update}
	ErrorOnNondeterministicUpdate interface{} `field:"optional" json:"errorOnNondeterministicUpdate" yaml:"errorOnNondeterministicUpdate"`
	// Specifies the name of the event table for logging messages from stored procedures and UDFs contained by the object with which the event table is associated.
	//
	// Associating an event table with a database is available in [Enterprise Edition or higher](https://docs.snowflake.com/en/user-guide/intro-editions). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information, check [EVENT_TABLE docs](https://docs.snowflake.com/en/sql-reference/parameters#event-table).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#event_table CurrentAccount#event_table}
	EventTable *string `field:"optional" json:"eventTable" yaml:"eventTable"`
	// Determines whether the ACCOUNTADMIN, ORGADMIN, GLOBALORGADMIN, and SECURITYADMIN roles can be used as the primary role when creating a Snowflake session based on the access token from the External OAuth authorization server.
	//
	// For more information, check [EXTERNAL_OAUTH_ADD_PRIVILEGED_ROLES_TO_BLOCKED_LIST docs](https://docs.snowflake.com/en/sql-reference/parameters#external-oauth-add-privileged-roles-to-blocked-list).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#external_oauth_add_privileged_roles_to_blocked_list CurrentAccount#external_oauth_add_privileged_roles_to_blocked_list}
	ExternalOauthAddPrivilegedRolesToBlockedList interface{} `field:"optional" json:"externalOauthAddPrivilegedRolesToBlockedList" yaml:"externalOauthAddPrivilegedRolesToBlockedList"`
	// Specifies the external volume for Apache Iceberg™ tables.
	//
	// For more information, see the [Iceberg table documentation](https://docs.snowflake.com/en/user-guide/tables-iceberg.html#label-tables-iceberg-external-volume-def). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information, check [EXTERNAL_VOLUME docs](https://docs.snowflake.com/en/sql-reference/parameters#external-volume).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#external_volume CurrentAccount#external_volume}
	ExternalVolume *string `field:"optional" json:"externalVolume" yaml:"externalVolume"`
	// Specifies [feature policy](https://docs.snowflake.com/en/developer-guide/native-apps/ui-consumer-feature-policies) for the current account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#feature_policy CurrentAccount#feature_policy}
	FeaturePolicy *string `field:"optional" json:"featurePolicy" yaml:"featurePolicy"`
	// Display format for [GEOGRAPHY values](https://docs.snowflake.com/en/sql-reference/data-types-geospatial.html#label-data-types-geography). Valid values are (case-insensitive): `GeoJSON` | `WKT` | `WKB` | `EWKT` | `EWKB`. For more information, check [GEOGRAPHY_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#geography-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#geography_output_format CurrentAccount#geography_output_format}
	GeographyOutputFormat *string `field:"optional" json:"geographyOutputFormat" yaml:"geographyOutputFormat"`
	// Display format for [GEOMETRY values](https://docs.snowflake.com/en/sql-reference/data-types-geospatial.html#label-data-types-geometry). Valid values are (case-insensitive): `GeoJSON` | `WKT` | `WKB` | `EWKT` | `EWKB`. For more information, check [GEOMETRY_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#geometry-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#geometry_output_format CurrentAccount#geometry_output_format}
	GeometryOutputFormat *string `field:"optional" json:"geometryOutputFormat" yaml:"geometryOutputFormat"`
	// Number of seconds to wait while trying to acquire row-level locks on a hybrid table, before timing out and aborting the statement.
	//
	// For more information, check [HYBRID_TABLE_LOCK_TIMEOUT docs](https://docs.snowflake.com/en/sql-reference/parameters#hybrid-table-lock-timeout).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#hybrid_table_lock_timeout CurrentAccount#hybrid_table_lock_timeout}
	HybridTableLockTimeout *float64 `field:"optional" json:"hybridTableLockTimeout" yaml:"hybridTableLockTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#id CurrentAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Sets the maximum estimated size limit for the initial replication of a primary database to a secondary database (in TB).
	//
	// Set this parameter on any account that stores a secondary database. This size limit helps prevent accounts from accidentally incurring large database replication charges. To remove the size limit, set the value to 0.0. It is required to pass numbers with scale of at least 1 (e.g. 20.5, 32.25, 33.333, etc.). For more information, check [INITIAL_REPLICATION_SIZE_LIMIT_IN_TB docs](https://docs.snowflake.com/en/sql-reference/parameters#initial-replication-size-limit-in-tb).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#initial_replication_size_limit_in_tb CurrentAccount#initial_replication_size_limit_in_tb}
	InitialReplicationSizeLimitInTb *string `field:"optional" json:"initialReplicationSizeLimitInTb" yaml:"initialReplicationSizeLimitInTb"`
	// Specifies how JDBC processes columns that have a scale of zero (0). For more information, check [JDBC_TREAT_DECIMAL_AS_INT docs](https://docs.snowflake.com/en/sql-reference/parameters#jdbc-treat-decimal-as-int).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#jdbc_treat_decimal_as_int CurrentAccount#jdbc_treat_decimal_as_int}
	JdbcTreatDecimalAsInt interface{} `field:"optional" json:"jdbcTreatDecimalAsInt" yaml:"jdbcTreatDecimalAsInt"`
	// Specifies how JDBC processes TIMESTAMP_NTZ values ([more details](https://docs.snowflake.com/en/sql-reference/parameters#jdbc-treat-timestamp-ntz-as-utc)). For more information, check [JDBC_TREAT_TIMESTAMP_NTZ_AS_UTC docs](https://docs.snowflake.com/en/sql-reference/parameters#jdbc-treat-timestamp-ntz-as-utc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#jdbc_treat_timestamp_ntz_as_utc CurrentAccount#jdbc_treat_timestamp_ntz_as_utc}
	JdbcTreatTimestampNtzAsUtc interface{} `field:"optional" json:"jdbcTreatTimestampNtzAsUtc" yaml:"jdbcTreatTimestampNtzAsUtc"`
	// Specifies whether the JDBC Driver uses the time zone of the JVM or the time zone of the session (specified by the [TIMEZONE](https://docs.snowflake.com/en/sql-reference/parameters#label-timezone) parameter) for the getDate(), getTime(), and getTimestamp() methods of the ResultSet class. For more information, check [JDBC_USE_SESSION_TIMEZONE docs](https://docs.snowflake.com/en/sql-reference/parameters#jdbc-use-session-timezone).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#jdbc_use_session_timezone CurrentAccount#jdbc_use_session_timezone}
	JdbcUseSessionTimezone interface{} `field:"optional" json:"jdbcUseSessionTimezone" yaml:"jdbcUseSessionTimezone"`
	// Specifies the number of blank spaces to indent each new element in JSON output in the session.
	//
	// Also specifies whether to insert newline characters after each element. For more information, check [JSON_INDENT docs](https://docs.snowflake.com/en/sql-reference/parameters#json-indent).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#json_indent CurrentAccount#json_indent}
	JsonIndent *float64 `field:"optional" json:"jsonIndent" yaml:"jsonIndent"`
	// Specifies how the Snowflake Node.js Driver processes numeric columns that have a scale of zero (0), for example INTEGER or NUMBER(p, 0). For more information, check [JS_TREAT_INTEGER_AS_BIGINT docs](https://docs.snowflake.com/en/sql-reference/parameters#js-treat-integer-as-bigint).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#js_treat_integer_as_bigint CurrentAccount#js_treat_integer_as_bigint}
	JsTreatIntegerAsBigint interface{} `field:"optional" json:"jsTreatIntegerAsBigint" yaml:"jsTreatIntegerAsBigint"`
	// Sets the time interval used to refresh the application package based data products to other regions.
	//
	// For more information, check [LISTING_AUTO_FULFILLMENT_REPLICATION_REFRESH_SCHEDULE docs](https://docs.snowflake.com/en/sql-reference/parameters#listing-auto-fulfillment-replication-refresh-schedule).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#listing_auto_fulfillment_replication_refresh_schedule CurrentAccount#listing_auto_fulfillment_replication_refresh_schedule}
	ListingAutoFulfillmentReplicationRefreshSchedule *string `field:"optional" json:"listingAutoFulfillmentReplicationRefreshSchedule" yaml:"listingAutoFulfillmentReplicationRefreshSchedule"`
	// Number of seconds to wait while trying to lock a resource, before timing out and aborting the statement.
	//
	// For more information, check [LOCK_TIMEOUT docs](https://docs.snowflake.com/en/sql-reference/parameters#lock-timeout).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#lock_timeout CurrentAccount#lock_timeout}
	LockTimeout *float64 `field:"optional" json:"lockTimeout" yaml:"lockTimeout"`
	// Specifies the severity level of messages that should be ingested and made available in the active event table.
	//
	// Messages at the specified level (and at more severe levels) are ingested. For more information about log levels, see [Setting levels for logging, metrics, and tracing](https://docs.snowflake.com/en/developer-guide/logging-tracing/telemetry-levels). Valid values are (case-insensitive): `TRACE` | `DEBUG` | `INFO` | `WARN` | `ERROR` | `FATAL` | `OFF`. For more information, check [LOG_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#log-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#log_level CurrentAccount#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Specifies the concurrency level for SQL statements (that is, queries and DML) executed by a warehouse ([more details](https://docs.snowflake.com/en/sql-reference/parameters#max-concurrency-level)). For more information, check [MAX_CONCURRENCY_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#max-concurrency-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#max_concurrency_level CurrentAccount#max_concurrency_level}
	MaxConcurrencyLevel *float64 `field:"optional" json:"maxConcurrencyLevel" yaml:"maxConcurrencyLevel"`
	// Maximum number of days Snowflake can extend the data retention period for tables to prevent streams on the tables from becoming stale.
	//
	// By default, if the [DATA_RETENTION_TIME_IN_DAYS](https://docs.snowflake.com/en/sql-reference/parameters#data-retention-time-in-days) setting for a source table is less than 14 days, and a stream has not been consumed, Snowflake temporarily extends this period to the stream’s offset, up to a maximum of 14 days, regardless of the [Snowflake Edition](https://docs.snowflake.com/en/user-guide/intro-editions) for your account. The MAX_DATA_EXTENSION_TIME_IN_DAYS parameter enables you to limit this automatic extension period to control storage costs for data retention or for compliance reasons. For more information, check [MAX_DATA_EXTENSION_TIME_IN_DAYS docs](https://docs.snowflake.com/en/sql-reference/parameters#max-data-extension-time-in-days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#max_data_extension_time_in_days CurrentAccount#max_data_extension_time_in_days}
	MaxDataExtensionTimeInDays *float64 `field:"optional" json:"maxDataExtensionTimeInDays" yaml:"maxDataExtensionTimeInDays"`
	// Controls how metrics data is ingested into the event table.
	//
	// For more information about metric levels, see [Setting levels for logging, metrics, and tracing](https://docs.snowflake.com/en/developer-guide/logging-tracing/telemetry-levels). Valid values are (case-insensitive): `ALL` | `NONE`. For more information, check [METRIC_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#metric-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#metric_level CurrentAccount#metric_level}
	MetricLevel *string `field:"optional" json:"metricLevel" yaml:"metricLevel"`
	// Minimum number of days for which Snowflake retains historical data for performing Time Travel actions (SELECT, CLONE, UNDROP) on an object.
	//
	// If a minimum number of days for data retention is set on an account, the data retention period for an object is determined by MAX([DATA_RETENTION_TIME_IN_DAYS](https://docs.snowflake.com/en/sql-reference/parameters#label-data-retention-time-in-days), MIN_DATA_RETENTION_TIME_IN_DAYS). For more information, check [MIN_DATA_RETENTION_TIME_IN_DAYS docs](https://docs.snowflake.com/en/sql-reference/parameters#min-data-retention-time-in-days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#min_data_retention_time_in_days CurrentAccount#min_data_retention_time_in_days}
	MinDataRetentionTimeInDays *float64 `field:"optional" json:"minDataRetentionTimeInDays" yaml:"minDataRetentionTimeInDays"`
	// Number of statements to execute when using the multi-statement capability. For more information, check [MULTI_STATEMENT_COUNT docs](https://docs.snowflake.com/en/sql-reference/parameters#multi-statement-count).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#multi_statement_count CurrentAccount#multi_statement_count}
	MultiStatementCount *float64 `field:"optional" json:"multiStatementCount" yaml:"multiStatementCount"`
	// Specifies the network policy to enforce for your account.
	//
	// Network policies enable restricting access to your account based on users’ IP address. For more details, see [Controlling network traffic with network policies](https://docs.snowflake.com/en/user-guide/network-policies). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`. For more information, check [NETWORK_POLICY docs](https://docs.snowflake.com/en/sql-reference/parameters#network-policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#network_policy CurrentAccount#network_policy}
	NetworkPolicy *string `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
	// Specifies whether the ORDER or NOORDER property is set by default when you create a new sequence or add a new table column.
	//
	// The ORDER and NOORDER properties determine whether or not the values are generated for the sequence or auto-incremented column in [increasing or decreasing order](https://docs.snowflake.com/en/user-guide/querying-sequences.html#label-querying-sequences-increasing-values). For more information, check [NOORDER_SEQUENCE_AS_DEFAULT docs](https://docs.snowflake.com/en/sql-reference/parameters#noorder-sequence-as-default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#noorder_sequence_as_default CurrentAccount#noorder_sequence_as_default}
	NoorderSequenceAsDefault interface{} `field:"optional" json:"noorderSequenceAsDefault" yaml:"noorderSequenceAsDefault"`
	// Determines whether the ACCOUNTADMIN, ORGADMIN, GLOBALORGADMIN, and SECURITYADMIN roles can be used as the primary role when creating a Snowflake session based on the access token from Snowflake’s authorization server.
	//
	// For more information, check [OAUTH_ADD_PRIVILEGED_ROLES_TO_BLOCKED_LIST docs](https://docs.snowflake.com/en/sql-reference/parameters#oauth-add-privileged-roles-to-blocked-list).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#oauth_add_privileged_roles_to_blocked_list CurrentAccount#oauth_add_privileged_roles_to_blocked_list}
	OauthAddPrivilegedRolesToBlockedList interface{} `field:"optional" json:"oauthAddPrivilegedRolesToBlockedList" yaml:"oauthAddPrivilegedRolesToBlockedList"`
	// Specifies how ODBC processes columns that have a scale of zero (0). For more information, check [ODBC_TREAT_DECIMAL_AS_INT docs](https://docs.snowflake.com/en/sql-reference/parameters#odbc-treat-decimal-as-int).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#odbc_treat_decimal_as_int CurrentAccount#odbc_treat_decimal_as_int}
	OdbcTreatDecimalAsInt interface{} `field:"optional" json:"odbcTreatDecimalAsInt" yaml:"odbcTreatDecimalAsInt"`
	// Specifies [packages policy](https://docs.snowflake.com/en/developer-guide/udf/python/packages-policy) for the current account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#packages_policy CurrentAccount#packages_policy}
	PackagesPolicy *string `field:"optional" json:"packagesPolicy" yaml:"packagesPolicy"`
	// Specifies [password policy](https://docs.snowflake.com/en/user-guide/password-authentication#label-using-password-policies) for the current account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#password_policy CurrentAccount#password_policy}
	PasswordPolicy *string `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// It enables/disables re-encryption of table data with new keys on a yearly basis to provide additional levels of data protection ([more details](https://docs.snowflake.com/en/sql-reference/parameters#periodic-data-rekeying)). For more information, check [PERIODIC_DATA_REKEYING docs](https://docs.snowflake.com/en/sql-reference/parameters#periodic-data-rekeying).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#periodic_data_rekeying CurrentAccount#periodic_data_rekeying}
	PeriodicDataRekeying interface{} `field:"optional" json:"periodicDataRekeying" yaml:"periodicDataRekeying"`
	// Specifies whether to pause a running pipe, primarily in preparation for transferring ownership of the pipe to a different role ([more details](https://docs.snowflake.com/en/sql-reference/parameters#pipe-execution-paused)). For more information, check [PIPE_EXECUTION_PAUSED docs](https://docs.snowflake.com/en/sql-reference/parameters#pipe-execution-paused).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#pipe_execution_paused CurrentAccount#pipe_execution_paused}
	PipeExecutionPaused interface{} `field:"optional" json:"pipeExecutionPaused" yaml:"pipeExecutionPaused"`
	// Specifies whether to prevent ad hoc data unload operations to external cloud storage locations (that is, [COPY INTO location](https://docs.snowflake.com/en/sql-reference/sql/copy-into-location) statements that specify the cloud storage URL and access settings directly in the statement). For an example, see [Unloading data from a table directly to files in an external location](https://docs.snowflake.com/en/sql-reference/sql/copy-into-location.html#label-copy-into-location-ad-hoc). For more information, check [PREVENT_UNLOAD_TO_INLINE_URL docs](https://docs.snowflake.com/en/sql-reference/parameters#prevent-unload-to-inline-url).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#prevent_unload_to_inline_url CurrentAccount#prevent_unload_to_inline_url}
	PreventUnloadToInlineUrl interface{} `field:"optional" json:"preventUnloadToInlineUrl" yaml:"preventUnloadToInlineUrl"`
	// Specifies whether to prevent data unload operations to internal (Snowflake) stages using [COPY INTO location](https://docs.snowflake.com/en/sql-reference/sql/copy-into-location) statements. For more information, check [PREVENT_UNLOAD_TO_INTERNAL_STAGES docs](https://docs.snowflake.com/en/sql-reference/parameters#prevent-unload-to-internal-stages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#prevent_unload_to_internal_stages CurrentAccount#prevent_unload_to_internal_stages}
	PreventUnloadToInternalStages interface{} `field:"optional" json:"preventUnloadToInternalStages" yaml:"preventUnloadToInternalStages"`
	// Specifies the list of Python modules to include in a report when [profiling Python handler code](https://docs.snowflake.com/en/developer-guide/stored-procedure/python/procedure-python-profiler). For more information, check [PYTHON_PROFILER_MODULES docs](https://docs.snowflake.com/en/sql-reference/parameters#python-profiler-modules).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#python_profiler_modules CurrentAccount#python_profiler_modules}
	PythonProfilerModules *string `field:"optional" json:"pythonProfilerModules" yaml:"pythonProfilerModules"`
	// Specifies the fully-qualified name of the stage in which to save a report when [profiling Python handler code](https://docs.snowflake.com/en/developer-guide/stored-procedure/python/procedure-python-profiler). For more information, check [PYTHON_PROFILER_TARGET_STAGE docs](https://docs.snowflake.com/en/sql-reference/parameters#python-profiler-target-stage).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#python_profiler_target_stage CurrentAccount#python_profiler_target_stage}
	PythonProfilerTargetStage *string `field:"optional" json:"pythonProfilerTargetStage" yaml:"pythonProfilerTargetStage"`
	// Optional string that can be used to tag queries and other SQL statements executed within a session.
	//
	// The tags are displayed in the output of the [QUERY_HISTORY, QUERY_HISTORY_BY_*](https://docs.snowflake.com/en/sql-reference/functions/query_history) functions. For more information, check [QUERY_TAG docs](https://docs.snowflake.com/en/sql-reference/parameters#query-tag).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#query_tag CurrentAccount#query_tag}
	QueryTag *string `field:"optional" json:"queryTag" yaml:"queryTag"`
	// Specifies whether letters in double-quoted object identifiers are stored and resolved as uppercase letters.
	//
	// By default, Snowflake preserves the case of alphabetic characters when storing and resolving double-quoted identifiers (see [Identifier resolution](https://docs.snowflake.com/en/sql-reference/identifiers-syntax.html#label-identifier-casing)). You can use this parameter in situations in which [third-party applications always use double quotes around identifiers](https://docs.snowflake.com/en/sql-reference/identifiers-syntax.html#label-identifier-casing-parameter). For more information, check [QUOTED_IDENTIFIERS_IGNORE_CASE docs](https://docs.snowflake.com/en/sql-reference/parameters#quoted-identifiers-ignore-case).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#quoted_identifiers_ignore_case CurrentAccount#quoted_identifiers_ignore_case}
	QuotedIdentifiersIgnoreCase interface{} `field:"optional" json:"quotedIdentifiersIgnoreCase" yaml:"quotedIdentifiersIgnoreCase"`
	// Specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (�) in query results for [Apache Iceberg™ tables](https://docs.snowflake.com/en/sql-reference/sql/create-iceberg-table) that use an external catalog. For more information, check [REPLACE_INVALID_CHARACTERS docs](https://docs.snowflake.com/en/sql-reference/parameters#replace-invalid-characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#replace_invalid_characters CurrentAccount#replace_invalid_characters}
	ReplaceInvalidCharacters interface{} `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// Specifies whether to require a storage integration object as cloud credentials when creating a named external stage (using [CREATE STAGE](https://docs.snowflake.com/en/sql-reference/sql/create-stage)) to access a private cloud storage location. For more information, check [REQUIRE_STORAGE_INTEGRATION_FOR_STAGE_CREATION docs](https://docs.snowflake.com/en/sql-reference/parameters#require-storage-integration-for-stage-creation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#require_storage_integration_for_stage_creation CurrentAccount#require_storage_integration_for_stage_creation}
	RequireStorageIntegrationForStageCreation interface{} `field:"optional" json:"requireStorageIntegrationForStageCreation" yaml:"requireStorageIntegrationForStageCreation"`
	// Specifies whether to require using a named external stage that references a storage integration object as cloud credentials when loading data from or unloading data to a private cloud storage location.
	//
	// For more information, check [REQUIRE_STORAGE_INTEGRATION_FOR_STAGE_OPERATION docs](https://docs.snowflake.com/en/sql-reference/parameters#require-storage-integration-for-stage-operation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#require_storage_integration_for_stage_operation CurrentAccount#require_storage_integration_for_stage_operation}
	RequireStorageIntegrationForStageOperation interface{} `field:"optional" json:"requireStorageIntegrationForStageOperation" yaml:"requireStorageIntegrationForStageOperation"`
	// Parameter that specifies the name of the resource monitor used to control all virtual warehouses created in the account.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#resource_monitor CurrentAccount#resource_monitor}
	ResourceMonitor *string `field:"optional" json:"resourceMonitor" yaml:"resourceMonitor"`
	// Specifies the maximum number of rows returned in a result set.
	//
	// A value of 0 specifies no maximum. For more information, check [ROWS_PER_RESULTSET docs](https://docs.snowflake.com/en/sql-reference/parameters#rows-per-resultset).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#rows_per_resultset CurrentAccount#rows_per_resultset}
	RowsPerResultset *float64 `field:"optional" json:"rowsPerResultset" yaml:"rowsPerResultset"`
	// Specifies the DNS name of an Amazon S3 interface endpoint.
	//
	// Requests sent to the internal stage of an account via [AWS PrivateLink for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/privatelink-interface-endpoints.html) use this endpoint to connect. For more information, see [Accessing Internal stages with dedicated interface endpoints](https://docs.snowflake.com/en/user-guide/private-internal-stages-aws.html#label-aws-privatelink-internal-stage-network-isolation). For more information, check [S3_STAGE_VPCE_DNS_NAME docs](https://docs.snowflake.com/en/sql-reference/parameters#s3-stage-vpce-dns-name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#s3_stage_vpce_dns_name CurrentAccount#s3_stage_vpce_dns_name}
	S3StageVpceDnsName *string `field:"optional" json:"s3StageVpceDnsName" yaml:"s3StageVpceDnsName"`
	// Enables federated authentication. This deprecated parameter enables federated authentication ([more details](https://docs.snowflake.com/en/sql-reference/parameters#saml-identity-provider)). For more information, check [SAML_IDENTITY_PROVIDER docs](https://docs.snowflake.com/en/sql-reference/parameters#saml-identity-provider).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#saml_identity_provider CurrentAccount#saml_identity_provider}
	SamlIdentityProvider *string `field:"optional" json:"samlIdentityProvider" yaml:"samlIdentityProvider"`
	// Specifies the path to search to resolve unqualified object names in queries.
	//
	// For more information, see [Name resolution in queries](https://docs.snowflake.com/en/sql-reference/name-resolution.html#label-object-name-resolution-search-path). Comma-separated list of identifiers. An identifier can be a fully or partially qualified schema name. For more information, check [SEARCH_PATH docs](https://docs.snowflake.com/en/sql-reference/parameters#search-path).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#search_path CurrentAccount#search_path}
	SearchPath *string `field:"optional" json:"searchPath" yaml:"searchPath"`
	// Specifies the maximum allowed warehouse size for [Serverless tasks](https://docs.snowflake.com/en/user-guide/tasks-intro.html#label-tasks-compute-resources-serverless). Valid values are (case-insensitive): `XSMALL` | `X-SMALL` | `SMALL` | `MEDIUM` | `LARGE` | `XLARGE` | `X-LARGE` | `XXLARGE` | `X2LARGE` | `2X-LARGE` | `XXXLARGE` | `X3LARGE` | `3X-LARGE` | `X4LARGE` | `4X-LARGE` | `X5LARGE` | `5X-LARGE` | `X6LARGE` | `6X-LARGE`. For more information, check [SERVERLESS_TASK_MAX_STATEMENT_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#serverless-task-max-statement-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#serverless_task_max_statement_size CurrentAccount#serverless_task_max_statement_size}
	ServerlessTaskMaxStatementSize *string `field:"optional" json:"serverlessTaskMaxStatementSize" yaml:"serverlessTaskMaxStatementSize"`
	// Specifies the minimum allowed warehouse size for [Serverless tasks](https://docs.snowflake.com/en/user-guide/tasks-intro.html#label-tasks-compute-resources-serverless). Valid values are (case-insensitive): `XSMALL` | `X-SMALL` | `SMALL` | `MEDIUM` | `LARGE` | `XLARGE` | `X-LARGE` | `XXLARGE` | `X2LARGE` | `2X-LARGE` | `XXXLARGE` | `X3LARGE` | `3X-LARGE` | `X4LARGE` | `4X-LARGE` | `X5LARGE` | `5X-LARGE` | `X6LARGE` | `6X-LARGE`. For more information, check [SERVERLESS_TASK_MIN_STATEMENT_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#serverless-task-min-statement-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#serverless_task_min_statement_size CurrentAccount#serverless_task_min_statement_size}
	ServerlessTaskMinStatementSize *string `field:"optional" json:"serverlessTaskMinStatementSize" yaml:"serverlessTaskMinStatementSize"`
	// Specifies [session policy](https://docs.snowflake.com/en/user-guide/session-policies-using) for the current account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#session_policy CurrentAccount#session_policy}
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// Specifies the name of a consumer account to simulate for testing/validating shared data, particularly shared secure views.
	//
	// When this parameter is set in a session, shared views return rows as if executed in the specified consumer account rather than the provider account. For more information, check [SIMULATED_DATA_SHARING_CONSUMER docs](https://docs.snowflake.com/en/sql-reference/parameters#simulated-data-sharing-consumer).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#simulated_data_sharing_consumer CurrentAccount#simulated_data_sharing_consumer}
	SimulatedDataSharingConsumer *string `field:"optional" json:"simulatedDataSharingConsumer" yaml:"simulatedDataSharingConsumer"`
	// This deprecated parameter disables preview mode for testing SSO (after enabling federated authentication) before rolling it out to users.
	//
	// For more information, check [SSO_LOGIN_PAGE docs](https://docs.snowflake.com/en/sql-reference/parameters#sso-login-page).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#sso_login_page CurrentAccount#sso_login_page}
	SsoLoginPage interface{} `field:"optional" json:"ssoLoginPage" yaml:"ssoLoginPage"`
	// Amount of time, in seconds, a SQL statement (query, DDL, DML, etc.) remains queued for a warehouse before it is canceled by the system. This parameter can be used in conjunction with the [MAX_CONCURRENCY_LEVEL](https://docs.snowflake.com/en/sql-reference/parameters#label-max-concurrency-level) parameter to ensure a warehouse is never backlogged. For more information, check [STATEMENT_QUEUED_TIMEOUT_IN_SECONDS docs](https://docs.snowflake.com/en/sql-reference/parameters#statement-queued-timeout-in-seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#statement_queued_timeout_in_seconds CurrentAccount#statement_queued_timeout_in_seconds}
	StatementQueuedTimeoutInSeconds *float64 `field:"optional" json:"statementQueuedTimeoutInSeconds" yaml:"statementQueuedTimeoutInSeconds"`
	// Amount of time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system. For more information, check [STATEMENT_TIMEOUT_IN_SECONDS docs](https://docs.snowflake.com/en/sql-reference/parameters#statement-timeout-in-seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#statement_timeout_in_seconds CurrentAccount#statement_timeout_in_seconds}
	StatementTimeoutInSeconds *float64 `field:"optional" json:"statementTimeoutInSeconds" yaml:"statementTimeoutInSeconds"`
	// Specifies the storage serialization policy for Snowflake-managed [Apache Iceberg™ tables](https://docs.snowflake.com/en/user-guide/tables-iceberg). Valid values are (case-insensitive): `COMPATIBLE` | `OPTIMIZED`. For more information, check [STORAGE_SERIALIZATION_POLICY docs](https://docs.snowflake.com/en/sql-reference/parameters#storage-serialization-policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#storage_serialization_policy CurrentAccount#storage_serialization_policy}
	StorageSerializationPolicy *string `field:"optional" json:"storageSerializationPolicy" yaml:"storageSerializationPolicy"`
	// This parameter specifies whether JSON output in a session is compatible with the general standard (as described by [http://json.org](http://json.org)). By design, Snowflake allows JSON input that contains non-standard values; however, these non-standard values might result in Snowflake outputting JSON that is incompatible with other platforms and languages. This parameter, when enabled, ensures that Snowflake outputs valid/compatible JSON. For more information, check [STRICT_JSON_OUTPUT docs](https://docs.snowflake.com/en/sql-reference/parameters#strict-json-output).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#strict_json_output CurrentAccount#strict_json_output}
	StrictJsonOutput interface{} `field:"optional" json:"strictJsonOutput" yaml:"strictJsonOutput"`
	// Specifies the number of consecutive failed task runs after which the current task is suspended automatically.
	//
	// The default is 0 (no automatic suspension). For more information, check [SUSPEND_TASK_AFTER_NUM_FAILURES docs](https://docs.snowflake.com/en/sql-reference/parameters#suspend-task-after-num-failures).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#suspend_task_after_num_failures CurrentAccount#suspend_task_after_num_failures}
	SuspendTaskAfterNumFailures *float64 `field:"optional" json:"suspendTaskAfterNumFailures" yaml:"suspendTaskAfterNumFailures"`
	// Specifies the number of automatic task graph retry attempts.
	//
	// If any task graphs complete in a FAILED state, Snowflake can automatically retry the task graphs from the last task in the graph that failed. For more information, check [TASK_AUTO_RETRY_ATTEMPTS docs](https://docs.snowflake.com/en/sql-reference/parameters#task-auto-retry-attempts).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#task_auto_retry_attempts CurrentAccount#task_auto_retry_attempts}
	TaskAutoRetryAttempts *float64 `field:"optional" json:"taskAutoRetryAttempts" yaml:"taskAutoRetryAttempts"`
	// Specifies the input format for the TIME data type.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). Any valid, supported time format or AUTO (AUTO specifies that Snowflake attempts to automatically detect the format of times stored in the system during the session). For more information, check [TIME_INPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#time-input-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#time_input_format CurrentAccount#time_input_format}
	TimeInputFormat *string `field:"optional" json:"timeInputFormat" yaml:"timeInputFormat"`
	// Specifies the display format for the TIME data type.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [TIME_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#time-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#time_output_format CurrentAccount#time_output_format}
	TimeOutputFormat *string `field:"optional" json:"timeOutputFormat" yaml:"timeOutputFormat"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timeouts CurrentAccount#timeouts}
	Timeouts *CurrentAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Specifies whether the [DATEADD](https://docs.snowflake.com/en/sql-reference/functions/dateadd) function (and its aliases) always consider a day to be exactly 24 hours for expressions that span multiple days. For more information, check [TIMESTAMP_DAY_IS_ALWAYS_24H docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-day-is-always-24h).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timestamp_day_is_always_24h CurrentAccount#timestamp_day_is_always_24h}
	TimestampDayIsAlways24H interface{} `field:"optional" json:"timestampDayIsAlways24H" yaml:"timestampDayIsAlways24H"`
	// Specifies the input format for the TIMESTAMP data type alias.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). Any valid, supported timestamp format or AUTO (AUTO specifies that Snowflake attempts to automatically detect the format of timestamps stored in the system during the session). For more information, check [TIMESTAMP_INPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-input-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timestamp_input_format CurrentAccount#timestamp_input_format}
	TimestampInputFormat *string `field:"optional" json:"timestampInputFormat" yaml:"timestampInputFormat"`
	// Specifies the display format for the TIMESTAMP_LTZ data type.
	//
	// If no format is specified, defaults to [TIMESTAMP_OUTPUT_FORMAT](https://docs.snowflake.com/en/sql-reference/parameters#label-timestamp-output-format). For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [TIMESTAMP_LTZ_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-ltz-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timestamp_ltz_output_format CurrentAccount#timestamp_ltz_output_format}
	TimestampLtzOutputFormat *string `field:"optional" json:"timestampLtzOutputFormat" yaml:"timestampLtzOutputFormat"`
	// Specifies the display format for the TIMESTAMP_NTZ data type. For more information, check [TIMESTAMP_NTZ_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-ntz-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timestamp_ntz_output_format CurrentAccount#timestamp_ntz_output_format}
	TimestampNtzOutputFormat *string `field:"optional" json:"timestampNtzOutputFormat" yaml:"timestampNtzOutputFormat"`
	// Specifies the display format for the TIMESTAMP data type alias.
	//
	// For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [TIMESTAMP_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timestamp_output_format CurrentAccount#timestamp_output_format}
	TimestampOutputFormat *string `field:"optional" json:"timestampOutputFormat" yaml:"timestampOutputFormat"`
	// Specifies the TIMESTAMP_* variation that the TIMESTAMP data type alias maps to.
	//
	// Valid values are (case-insensitive): `TIMESTAMP_LTZ` | `TIMESTAMP_NTZ` | `TIMESTAMP_TZ`. For more information, check [TIMESTAMP_TYPE_MAPPING docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-type-mapping).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timestamp_type_mapping CurrentAccount#timestamp_type_mapping}
	TimestampTypeMapping *string `field:"optional" json:"timestampTypeMapping" yaml:"timestampTypeMapping"`
	// Specifies the display format for the TIMESTAMP_TZ data type.
	//
	// If no format is specified, defaults to [TIMESTAMP_OUTPUT_FORMAT](https://docs.snowflake.com/en/sql-reference/parameters#label-timestamp-output-format). For more information, see [Date and time input and output formats](https://docs.snowflake.com/en/sql-reference/date-time-input-output). For more information, check [TIMESTAMP_TZ_OUTPUT_FORMAT docs](https://docs.snowflake.com/en/sql-reference/parameters#timestamp-tz-output-format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timestamp_tz_output_format CurrentAccount#timestamp_tz_output_format}
	TimestampTzOutputFormat *string `field:"optional" json:"timestampTzOutputFormat" yaml:"timestampTzOutputFormat"`
	// Specifies the time zone for the session.
	//
	// You can specify a [time zone name](https://data.iana.org/time-zones/tzdb-2021a/zone1970.tab) or a [link name](https://data.iana.org/time-zones/tzdb-2021a/backward) from release 2021a of the [IANA Time Zone Database](https://www.iana.org/time-zones) (e.g. America/Los_Angeles, Europe/London, UTC, Etc/GMT, etc.). For more information, check [TIMEZONE docs](https://docs.snowflake.com/en/sql-reference/parameters#timezone).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#timezone CurrentAccount#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Controls how trace events are ingested into the event table.
	//
	// For more information about trace levels, see [Setting trace level](https://docs.snowflake.com/en/developer-guide/logging-tracing/tracing-trace-level). Valid values are (case-insensitive): `ALWAYS` | `ON_EVENT` | `PROPAGATE` | `OFF`. For more information, check [TRACE_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#trace-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#trace_level CurrentAccount#trace_level}
	TraceLevel *string `field:"optional" json:"traceLevel" yaml:"traceLevel"`
	// Specifies the action to perform when a statement issued within a non-autocommit transaction returns with an error.
	//
	// For more information, check [TRANSACTION_ABORT_ON_ERROR docs](https://docs.snowflake.com/en/sql-reference/parameters#transaction-abort-on-error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#transaction_abort_on_error CurrentAccount#transaction_abort_on_error}
	TransactionAbortOnError interface{} `field:"optional" json:"transactionAbortOnError" yaml:"transactionAbortOnError"`
	// Specifies the isolation level for transactions in the user session.
	//
	// Valid values are (case-insensitive): `READ COMMITTED`. For more information, check [TRANSACTION_DEFAULT_ISOLATION_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#transaction-default-isolation-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#transaction_default_isolation_level CurrentAccount#transaction_default_isolation_level}
	TransactionDefaultIsolationLevel *string `field:"optional" json:"transactionDefaultIsolationLevel" yaml:"transactionDefaultIsolationLevel"`
	// Specifies the “century start” year for 2-digit years (i.e. the earliest year such dates can represent). This parameter prevents ambiguous dates when importing or converting data with the `YY` date format component (i.e. years represented as 2 digits). For more information, check [TWO_DIGIT_CENTURY_START docs](https://docs.snowflake.com/en/sql-reference/parameters#two-digit-century-start).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#two_digit_century_start CurrentAccount#two_digit_century_start}
	TwoDigitCenturyStart *float64 `field:"optional" json:"twoDigitCenturyStart" yaml:"twoDigitCenturyStart"`
	// Determines if an unsupported (i.e. non-default) value specified for a constraint property returns an error. For more information, check [UNSUPPORTED_DDL_ACTION docs](https://docs.snowflake.com/en/sql-reference/parameters#unsupported-ddl-action).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#unsupported_ddl_action CurrentAccount#unsupported_ddl_action}
	UnsupportedDdlAction *string `field:"optional" json:"unsupportedDdlAction" yaml:"unsupportedDdlAction"`
	// Specifies whether to reuse persisted query results, if available, when a matching query is submitted.
	//
	// For more information, check [USE_CACHED_RESULT docs](https://docs.snowflake.com/en/sql-reference/parameters#use-cached-result).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#use_cached_result CurrentAccount#use_cached_result}
	UseCachedResult interface{} `field:"optional" json:"useCachedResult" yaml:"useCachedResult"`
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size.
	//
	// Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. Valid values are (case-insensitive): `XSMALL` | `X-SMALL` | `SMALL` | `MEDIUM` | `LARGE` | `XLARGE` | `X-LARGE` | `XXLARGE` | `X2LARGE` | `2X-LARGE` | `XXXLARGE` | `X3LARGE` | `3X-LARGE` | `X4LARGE` | `4X-LARGE` | `X5LARGE` | `5X-LARGE` | `X6LARGE` | `6X-LARGE`. For more information, check [USER_TASK_MANAGED_INITIAL_WAREHOUSE_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-managed-initial-warehouse-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#user_task_managed_initial_warehouse_size CurrentAccount#user_task_managed_initial_warehouse_size}
	UserTaskManagedInitialWarehouseSize *string `field:"optional" json:"userTaskManagedInitialWarehouseSize" yaml:"userTaskManagedInitialWarehouseSize"`
	// Minimum amount of time between Triggered Task executions in seconds For more information, check [USER_TASK_MINIMUM_TRIGGER_INTERVAL_IN_SECONDS docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-minimum-trigger-interval-in-seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#user_task_minimum_trigger_interval_in_seconds CurrentAccount#user_task_minimum_trigger_interval_in_seconds}
	UserTaskMinimumTriggerIntervalInSeconds *float64 `field:"optional" json:"userTaskMinimumTriggerIntervalInSeconds" yaml:"userTaskMinimumTriggerIntervalInSeconds"`
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	//
	// For more information, check [USER_TASK_TIMEOUT_MS docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-timeout-ms).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#user_task_timeout_ms CurrentAccount#user_task_timeout_ms}
	UserTaskTimeoutMs *float64 `field:"optional" json:"userTaskTimeoutMs" yaml:"userTaskTimeoutMs"`
	// Specifies how the weeks in a given year are computed.
	//
	// `0`: The semantics used are equivalent to the ISO semantics, in which a week belongs to a given year if at least 4 days of that week are in that year. `1`: January 1 is included in the first week of the year and December 31 is included in the last week of the year. For more information, check [WEEK_OF_YEAR_POLICY docs](https://docs.snowflake.com/en/sql-reference/parameters#week-of-year-policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#week_of_year_policy CurrentAccount#week_of_year_policy}
	WeekOfYearPolicy *float64 `field:"optional" json:"weekOfYearPolicy" yaml:"weekOfYearPolicy"`
	// Specifies the first day of the week (used by week-related date functions).
	//
	// `0`: Legacy Snowflake behavior is used (i.e. ISO-like semantics). `1` (Monday) to `7` (Sunday): All the week-related functions use weeks that start on the specified day of the week. For more information, check [WEEK_START docs](https://docs.snowflake.com/en/sql-reference/parameters#week-start).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/current_account#week_start CurrentAccount#week_start}
	WeekStart *float64 `field:"optional" json:"weekStart" yaml:"weekStart"`
}

