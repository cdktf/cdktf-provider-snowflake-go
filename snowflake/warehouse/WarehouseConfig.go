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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#name Warehouse#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#auto_resume Warehouse#auto_resume}
	AutoResume *string `field:"optional" json:"autoResume" yaml:"autoResume"`
	// Specifies the number of seconds of inactivity after which a warehouse is automatically suspended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#auto_suspend Warehouse#auto_suspend}
	AutoSuspend *float64 `field:"optional" json:"autoSuspend" yaml:"autoSuspend"`
	// Specifies a comment for the warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#comment Warehouse#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether to enable the query acceleration service for queries that rely on this warehouse for compute resources.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#enable_query_acceleration Warehouse#enable_query_acceleration}
	EnableQueryAcceleration *string `field:"optional" json:"enableQueryAcceleration" yaml:"enableQueryAcceleration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#id Warehouse#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether the warehouse is created initially in the ‘Suspended’ state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#initially_suspended Warehouse#initially_suspended}
	InitiallySuspended interface{} `field:"optional" json:"initiallySuspended" yaml:"initiallySuspended"`
	// Specifies the maximum number of server clusters for the warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#max_cluster_count Warehouse#max_cluster_count}
	MaxClusterCount *float64 `field:"optional" json:"maxClusterCount" yaml:"maxClusterCount"`
	// Object parameter that specifies the concurrency level for SQL statements (i.e. queries and DML) executed by a warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#max_concurrency_level Warehouse#max_concurrency_level}
	MaxConcurrencyLevel *float64 `field:"optional" json:"maxConcurrencyLevel" yaml:"maxConcurrencyLevel"`
	// Specifies the minimum number of server clusters for the warehouse (only applies to multi-cluster warehouses).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#min_cluster_count Warehouse#min_cluster_count}
	MinClusterCount *float64 `field:"optional" json:"minClusterCount" yaml:"minClusterCount"`
	// Specifies the maximum scale factor for leasing compute resources for query acceleration.
	//
	// The scale factor is used as a multiplier based on warehouse size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#query_acceleration_max_scale_factor Warehouse#query_acceleration_max_scale_factor}
	QueryAccelerationMaxScaleFactor *float64 `field:"optional" json:"queryAccelerationMaxScaleFactor" yaml:"queryAccelerationMaxScaleFactor"`
	// Specifies the name of a resource monitor that is explicitly assigned to the warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#resource_monitor Warehouse#resource_monitor}
	ResourceMonitor *string `field:"optional" json:"resourceMonitor" yaml:"resourceMonitor"`
	// Specifies the policy for automatically starting and shutting down clusters in a multi-cluster warehouse running in Auto-scale mode.
	//
	// Valid values are (case-insensitive): `STANDARD` | `ECONOMY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#scaling_policy Warehouse#scaling_policy}
	ScalingPolicy *string `field:"optional" json:"scalingPolicy" yaml:"scalingPolicy"`
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on a warehouse before it is canceled by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#statement_queued_timeout_in_seconds Warehouse#statement_queued_timeout_in_seconds}
	StatementQueuedTimeoutInSeconds *float64 `field:"optional" json:"statementQueuedTimeoutInSeconds" yaml:"statementQueuedTimeoutInSeconds"`
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#statement_timeout_in_seconds Warehouse#statement_timeout_in_seconds}
	StatementTimeoutInSeconds *float64 `field:"optional" json:"statementTimeoutInSeconds" yaml:"statementTimeoutInSeconds"`
	// Specifies the size of the virtual warehouse.
	//
	// Valid values are (case-insensitive): `XSMALL` | `X-SMALL` | `SMALL` | `MEDIUM` | `LARGE` | `XLARGE` | `X-LARGE` | `XXLARGE` | `X2LARGE` | `2X-LARGE` | `XXXLARGE` | `X3LARGE` | `3X-LARGE` | `X4LARGE` | `4X-LARGE` | `X5LARGE` | `5X-LARGE` | `X6LARGE` | `6X-LARGE`. Consult [warehouse documentation](https://docs.snowflake.com/en/sql-reference/sql/create-warehouse#optional-properties-objectproperties) for the details. Note: removing the size from config will result in the resource recreation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#warehouse_size Warehouse#warehouse_size}
	WarehouseSize *string `field:"optional" json:"warehouseSize" yaml:"warehouseSize"`
	// Specifies warehouse type.
	//
	// Valid values are (case-insensitive): `STANDARD` | `SNOWPARK-OPTIMIZED`. Warehouse needs to be suspended to change its type. Provider will handle automatic suspension and resumption if needed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.93.0/docs/resources/warehouse#warehouse_type Warehouse#warehouse_type}
	WarehouseType *string `field:"optional" json:"warehouseType" yaml:"warehouseType"`
}

