// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TaskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#database Task#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the task;
	//
	// must be unique for the database and schema in which the task is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#name Task#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the task.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#schema Task#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Any single SQL statement, or a call to a stored procedure, executed when the task runs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#sql_statement Task#sql_statement}
	SqlStatement *string `field:"required" json:"sqlStatement" yaml:"sqlStatement"`
	// Specifies the predecessor task in the same database and schema of the current task.
	//
	// When a run of the predecessor task finishes successfully, it triggers this task (after a brief lag). (Conflict with schedule)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#after Task#after}
	After *string `field:"optional" json:"after" yaml:"after"`
	// Specifies a comment for the task.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#comment Task#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies if the task should be started (enabled) after creation or should remain suspended (default).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#enabled Task#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the name of the notification integration used for error notifications.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#error_integration Task#error_integration}
	ErrorIntegration *string `field:"optional" json:"errorIntegration" yaml:"errorIntegration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#id Task#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The schedule for periodically running the task. This can be a cron or interval in minutes. (Conflict with after).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#schedule Task#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Specifies session parameters to set for the session when the task runs. A task supports all session parameters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#session_parameters Task#session_parameters}
	SessionParameters *map[string]*string `field:"optional" json:"sessionParameters" yaml:"sessionParameters"`
	// Specifies the size of the compute resources to provision for the first run of the task, before a task history is available for Snowflake to determine an ideal size.
	//
	// Once a task has successfully completed a few runs, Snowflake ignores this parameter setting. (Conflicts with warehouse)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#user_task_managed_initial_warehouse_size Task#user_task_managed_initial_warehouse_size}
	UserTaskManagedInitialWarehouseSize *string `field:"optional" json:"userTaskManagedInitialWarehouseSize" yaml:"userTaskManagedInitialWarehouseSize"`
	// Specifies the time limit on a single run of the task before it times out (in milliseconds).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#user_task_timeout_ms Task#user_task_timeout_ms}
	UserTaskTimeoutMs *float64 `field:"optional" json:"userTaskTimeoutMs" yaml:"userTaskTimeoutMs"`
	// The warehouse the task will use.
	//
	// Omit this parameter to use Snowflake-managed compute resources for runs of this task. (Conflicts with user_task_managed_initial_warehouse_size)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#warehouse Task#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
	// Specifies a Boolean SQL expression; multiple conditions joined with AND/OR are supported.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/task#when Task#when}
	When *string `field:"optional" json:"when" yaml:"when"`
}

