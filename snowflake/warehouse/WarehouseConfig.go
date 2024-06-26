// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package warehouse

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WarehouseConfig struct {
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
	// Identifier for the virtual warehouse; must be unique for your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#name Warehouse#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#auto_resume Warehouse#auto_resume}
	AutoResume interface{} `field:"optional" json:"autoResume" yaml:"autoResume"`
	// Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#auto_suspend Warehouse#auto_suspend}
	AutoSuspend *float64 `field:"optional" json:"autoSuspend" yaml:"autoSuspend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#comment Warehouse#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether to enable the query acceleration service for queries that rely on this warehouse for compute resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#enable_query_acceleration Warehouse#enable_query_acceleration}
	EnableQueryAcceleration interface{} `field:"optional" json:"enableQueryAcceleration" yaml:"enableQueryAcceleration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#id Warehouse#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether the warehouse is created initially in the ‘Suspended’ state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#initially_suspended Warehouse#initially_suspended}
	InitiallySuspended interface{} `field:"optional" json:"initiallySuspended" yaml:"initiallySuspended"`
	// Specifies the maximum number of server clusters for the warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#max_cluster_count Warehouse#max_cluster_count}
	MaxClusterCount *float64 `field:"optional" json:"maxClusterCount" yaml:"maxClusterCount"`
	// Object parameter that specifies the concurrency level for SQL statements (i.e. queries and DML) executed by a warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#max_concurrency_level Warehouse#max_concurrency_level}
	MaxConcurrencyLevel *float64 `field:"optional" json:"maxConcurrencyLevel" yaml:"maxConcurrencyLevel"`
	// Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#min_cluster_count Warehouse#min_cluster_count}
	MinClusterCount *float64 `field:"optional" json:"minClusterCount" yaml:"minClusterCount"`
	// Specifies the maximum scale factor for leasing compute resources for query acceleration.
	//
	// The scale factor is used as a multiplier based on warehouse size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#query_acceleration_max_scale_factor Warehouse#query_acceleration_max_scale_factor}
	QueryAccelerationMaxScaleFactor *float64 `field:"optional" json:"queryAccelerationMaxScaleFactor" yaml:"queryAccelerationMaxScaleFactor"`
	// Specifies the name of a resource monitor that is explicitly assigned to the warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#resource_monitor Warehouse#resource_monitor}
	ResourceMonitor *string `field:"optional" json:"resourceMonitor" yaml:"resourceMonitor"`
	// Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#scaling_policy Warehouse#scaling_policy}
	ScalingPolicy *string `field:"optional" json:"scalingPolicy" yaml:"scalingPolicy"`
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on a warehouse before it is canceled by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#statement_queued_timeout_in_seconds Warehouse#statement_queued_timeout_in_seconds}
	StatementQueuedTimeoutInSeconds *float64 `field:"optional" json:"statementQueuedTimeoutInSeconds" yaml:"statementQueuedTimeoutInSeconds"`
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#statement_timeout_in_seconds Warehouse#statement_timeout_in_seconds}
	StatementTimeoutInSeconds *float64 `field:"optional" json:"statementTimeoutInSeconds" yaml:"statementTimeoutInSeconds"`
	// Specifies whether the warehouse, after being resized, waits for all the servers to provision before executing any queued or new queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#wait_for_provisioning Warehouse#wait_for_provisioning}
	WaitForProvisioning interface{} `field:"optional" json:"waitForProvisioning" yaml:"waitForProvisioning"`
	// Specifies the size of the virtual warehouse.
	//
	// Larger warehouse sizes 5X-Large and 6X-Large are currently in preview and only available on Amazon Web Services (AWS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#warehouse_size Warehouse#warehouse_size}
	WarehouseSize *string `field:"optional" json:"warehouseSize" yaml:"warehouseSize"`
	// Specifies a STANDARD or SNOWPARK-OPTIMIZED warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/warehouse#warehouse_type Warehouse#warehouse_type}
	WarehouseType *string `field:"optional" json:"warehouseType" yaml:"warehouseType"`
}

