// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secondarydatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecondaryDatabaseConfig struct {
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
	// A fully qualified path to a database to create a replica from.
	//
	// A fully qualified path follows the format of `"<organization_name>"."<account_name>"."<database_name>"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#as_replica_of SecondaryDatabase#as_replica_of}
	AsReplicaOf *string `field:"required" json:"asReplicaOf" yaml:"asReplicaOf"`
	// Specifies the identifier for the database;
	//
	// must be unique for your account. As a best practice for [Database Replication and Failover](https://docs.snowflake.com/en/user-guide/db-replication-intro), it is recommended to give each secondary database the same name as its primary database. This practice supports referencing fully-qualified objects (i.e. '<db>.<schema>.<object>') by other objects in the same database, such as querying a fully-qualified table name in a view. If a secondary database has a different name from the primary database, then these object references would break in the secondary database. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#name SecondaryDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The database parameter that specifies the default catalog to use for Iceberg tables. For more information, see [CATALOG](https://docs.snowflake.com/en/sql-reference/parameters#catalog).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#catalog SecondaryDatabase#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Specifies a comment for the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#comment SecondaryDatabase#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the number of days for which Time Travel actions (CLONE and UNDROP) can be performed on the database, as well as specifying the default Time Travel retention time for all schemas created in the database.
	//
	// For more details, see [Understanding & Using Time Travel](https://docs.snowflake.com/en/user-guide/data-time-travel).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#data_retention_time_in_days SecondaryDatabase#data_retention_time_in_days}
	DataRetentionTimeInDays *float64 `field:"optional" json:"dataRetentionTimeInDays" yaml:"dataRetentionTimeInDays"`
	// Specifies a default collation specification for all schemas and tables added to the database.
	//
	// It can be overridden on schema or table level. For more information, see [collation specification](https://docs.snowflake.com/en/sql-reference/collation#label-collation-specification).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#default_ddl_collation SecondaryDatabase#default_ddl_collation}
	DefaultDdlCollation *string `field:"optional" json:"defaultDdlCollation" yaml:"defaultDdlCollation"`
	// If true, enables stdout/stderr fast path logging for anonymous stored procedures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#enable_console_output SecondaryDatabase#enable_console_output}
	EnableConsoleOutput interface{} `field:"optional" json:"enableConsoleOutput" yaml:"enableConsoleOutput"`
	// The database parameter that specifies the default external volume to use for Iceberg tables. For more information, see [EXTERNAL_VOLUME](https://docs.snowflake.com/en/sql-reference/parameters#external-volume).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#external_volume SecondaryDatabase#external_volume}
	ExternalVolume *string `field:"optional" json:"externalVolume" yaml:"externalVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#id SecondaryDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the database as transient.
	//
	// Transient databases do not have a Fail-safe period so they do not incur additional storage costs once they leave Time Travel; however, this means they are also not protected by Fail-safe in the event of a data loss.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#is_transient SecondaryDatabase#is_transient}
	IsTransient interface{} `field:"optional" json:"isTransient" yaml:"isTransient"`
	// Specifies the severity level of messages that should be ingested and made available in the active event table.
	//
	// Valid options are: [TRACE DEBUG INFO WARN ERROR FATAL OFF]. Messages at the specified level (and at more severe levels) are ingested. For more information, see [LOG_LEVEL](https://docs.snowflake.com/en/sql-reference/parameters.html#label-log-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#log_level SecondaryDatabase#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Object parameter that specifies the maximum number of days for which Snowflake can extend the data retention period for tables in the database to prevent streams on the tables from becoming stale.
	//
	// For a detailed description of this parameter, see [MAX_DATA_EXTENSION_TIME_IN_DAYS](https://docs.snowflake.com/en/sql-reference/parameters.html#label-max-data-extension-time-in-days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#max_data_extension_time_in_days SecondaryDatabase#max_data_extension_time_in_days}
	MaxDataExtensionTimeInDays *float64 `field:"optional" json:"maxDataExtensionTimeInDays" yaml:"maxDataExtensionTimeInDays"`
	// If true, the case of quoted identifiers is ignored. For more information, see [QUOTED_IDENTIFIERS_IGNORE_CASE](https://docs.snowflake.com/en/sql-reference/parameters#quoted-identifiers-ignore-case).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#quoted_identifiers_ignore_case SecondaryDatabase#quoted_identifiers_ignore_case}
	QuotedIdentifiersIgnoreCase interface{} `field:"optional" json:"quotedIdentifiersIgnoreCase" yaml:"quotedIdentifiersIgnoreCase"`
	// Specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (�) in query results for an Iceberg table.
	//
	// You can only set this parameter for tables that use an external Iceberg catalog. For more information, see [REPLACE_INVALID_CHARACTERS](https://docs.snowflake.com/en/sql-reference/parameters#replace-invalid-characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#replace_invalid_characters SecondaryDatabase#replace_invalid_characters}
	ReplaceInvalidCharacters interface{} `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// The storage serialization policy for Iceberg tables that use Snowflake as the catalog.
	//
	// Valid options are: [COMPATIBLE OPTIMIZED]. COMPATIBLE: Snowflake performs encoding and compression of data files that ensures interoperability with third-party compute engines. OPTIMIZED: Snowflake performs encoding and compression of data files that ensures the best table performance within Snowflake. For more information, see [STORAGE_SERIALIZATION_POLICY](https://docs.snowflake.com/en/sql-reference/parameters#storage-serialization-policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#storage_serialization_policy SecondaryDatabase#storage_serialization_policy}
	StorageSerializationPolicy *string `field:"optional" json:"storageSerializationPolicy" yaml:"storageSerializationPolicy"`
	// How many times a task must fail in a row before it is automatically suspended.
	//
	// 0 disables auto-suspending. For more information, see [SUSPEND_TASK_AFTER_NUM_FAILURES](https://docs.snowflake.com/en/sql-reference/parameters#suspend-task-after-num-failures).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#suspend_task_after_num_failures SecondaryDatabase#suspend_task_after_num_failures}
	SuspendTaskAfterNumFailures *float64 `field:"optional" json:"suspendTaskAfterNumFailures" yaml:"suspendTaskAfterNumFailures"`
	// Maximum automatic retries allowed for a user task. For more information, see [TASK_AUTO_RETRY_ATTEMPTS](https://docs.snowflake.com/en/sql-reference/parameters#task-auto-retry-attempts).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#task_auto_retry_attempts SecondaryDatabase#task_auto_retry_attempts}
	TaskAutoRetryAttempts *float64 `field:"optional" json:"taskAutoRetryAttempts" yaml:"taskAutoRetryAttempts"`
	// Controls how trace events are ingested into the event table.
	//
	// Valid options are: [ALWAYS ON_EVENT OFF]. For information about levels, see [TRACE_LEVEL](https://docs.snowflake.com/en/sql-reference/parameters.html#label-trace-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#trace_level SecondaryDatabase#trace_level}
	TraceLevel *string `field:"optional" json:"traceLevel" yaml:"traceLevel"`
	// The initial size of warehouse to use for managed warehouses in the absence of history.
	//
	// For more information, see [USER_TASK_MANAGED_INITIAL_WAREHOUSE_SIZE](https://docs.snowflake.com/en/sql-reference/parameters#user-task-managed-initial-warehouse-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#user_task_managed_initial_warehouse_size SecondaryDatabase#user_task_managed_initial_warehouse_size}
	UserTaskManagedInitialWarehouseSize *string `field:"optional" json:"userTaskManagedInitialWarehouseSize" yaml:"userTaskManagedInitialWarehouseSize"`
	// Minimum amount of time between Triggered Task executions in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#user_task_minimum_trigger_interval_in_seconds SecondaryDatabase#user_task_minimum_trigger_interval_in_seconds}
	UserTaskMinimumTriggerIntervalInSeconds *float64 `field:"optional" json:"userTaskMinimumTriggerIntervalInSeconds" yaml:"userTaskMinimumTriggerIntervalInSeconds"`
	// User task execution timeout in milliseconds. For more information, see [USER_TASK_TIMEOUT_MS](https://docs.snowflake.com/en/sql-reference/parameters#user-task-timeout-ms).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/secondary_database#user_task_timeout_ms SecondaryDatabase#user_task_timeout_ms}
	UserTaskTimeoutMs *float64 `field:"optional" json:"userTaskTimeoutMs" yaml:"userTaskTimeoutMs"`
}

