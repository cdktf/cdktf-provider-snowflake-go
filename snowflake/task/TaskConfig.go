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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#database Task#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the task;
	//
	// must be unique for the database and schema in which the task is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#name Task#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#schema Task#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Any single SQL statement, or a call to a stored procedure, executed when the task runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#sql_statement Task#sql_statement}
	SqlStatement *string `field:"required" json:"sqlStatement" yaml:"sqlStatement"`
	// Specifies one or more predecessor tasks for the current task.
	//
	// Use this option to create a DAG of tasks or add this task to an existing DAG. A DAG is a series of tasks that starts with a scheduled root task and is linked together by dependencies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#after Task#after}
	After *[]*string `field:"optional" json:"after" yaml:"after"`
	// By default, Snowflake ensures that only one instance of a particular DAG is allowed to run at a time, setting the parameter value to TRUE permits DAG runs to overlap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#allow_overlapping_execution Task#allow_overlapping_execution}
	AllowOverlappingExecution interface{} `field:"optional" json:"allowOverlappingExecution" yaml:"allowOverlappingExecution"`
	// Specifies a comment for the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#comment Task#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies if the task should be started (enabled) after creation or should remain suspended (default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#enabled Task#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the name of the notification integration used for error notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#error_integration Task#error_integration}
	ErrorIntegration *string `field:"optional" json:"errorIntegration" yaml:"errorIntegration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#id Task#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The schedule for periodically running the task. This can be a cron or interval in minutes. (Conflict with after).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#schedule Task#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Specifies session parameters to set for the session when the task runs. A task supports all session parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#session_parameters Task#session_parameters}
	SessionParameters *map[string]*string `field:"optional" json:"sessionParameters" yaml:"sessionParameters"`
	// Specifies the number of consecutive failed task runs after which the current task is suspended automatically.
	//
	// The default is 0 (no automatic suspension).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#suspend_task_after_num_failures Task#suspend_task_after_num_failures}
	SuspendTaskAfterNumFailures *float64 `field:"optional" json:"suspendTaskAfterNumFailures" yaml:"suspendTaskAfterNumFailures"`
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size.
	//
	// Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. (Conflicts with warehouse)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#user_task_managed_initial_warehouse_size Task#user_task_managed_initial_warehouse_size}
	UserTaskManagedInitialWarehouseSize *string `field:"optional" json:"userTaskManagedInitialWarehouseSize" yaml:"userTaskManagedInitialWarehouseSize"`
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#user_task_timeout_ms Task#user_task_timeout_ms}
	UserTaskTimeoutMs *float64 `field:"optional" json:"userTaskTimeoutMs" yaml:"userTaskTimeoutMs"`
	// The warehouse the task will use.
	//
	// Omit this parameter to use Snowflake-managed compute resources for runs of this task. (Conflicts with user_task_managed_initial_warehouse_size)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#warehouse Task#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
	// Specifies a Boolean SQL expression; multiple conditions joined with AND/OR are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/task#when Task#when}
	When *string `field:"optional" json:"when" yaml:"when"`
}

