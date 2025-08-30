// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourceMonitorConfig struct {
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
	// Identifier for the resource monitor;
	//
	// must be unique for your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#name ResourceMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of credits allocated to the resource monitor per frequency interval.
	//
	// When total usage for all warehouses assigned to the monitor reaches this number for the current frequency interval, the resource monitor is considered to be at 100% of quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#credit_quota ResourceMonitor#credit_quota}
	CreditQuota *float64 `field:"optional" json:"creditQuota" yaml:"creditQuota"`
	// The date and time when the resource monitor suspends the assigned warehouses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#end_timestamp ResourceMonitor#end_timestamp}
	EndTimestamp *string `field:"optional" json:"endTimestamp" yaml:"endTimestamp"`
	// The frequency interval at which the credit usage resets to 0.
	//
	// Valid values are (case-insensitive): `MONTHLY` | `DAILY` | `WEEKLY` | `YEARLY` | `NEVER`. If you set a `frequency` for a resource monitor, you must also set `start_timestamp`. If you specify `NEVER` for the frequency, the credit usage for the warehouse does not reset. After removing this field from the config, the previously set value will be preserved on the Snowflake side, not the default value. That's due to Snowflake limitation and the lack of unset functionality for this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#frequency ResourceMonitor#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#id ResourceMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies a list of percentages of the credit quota.
	//
	// After reaching any of the values the users passed in the notify_users field will be notified (to receive the notification they should have notifications enabled). Values over 100 are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#notify_triggers ResourceMonitor#notify_triggers}
	NotifyTriggers *[]*float64 `field:"optional" json:"notifyTriggers" yaml:"notifyTriggers"`
	// Specifies the list of users (their identifiers) to receive email notifications on resource monitors.
	//
	// For more information about this resource, see [docs](./user).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#notify_users ResourceMonitor#notify_users}
	NotifyUsers *[]*string `field:"optional" json:"notifyUsers" yaml:"notifyUsers"`
	// The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
	//
	// If you set a `start_timestamp` for a resource monitor, you must also set `frequency`.  After removing this field from the config, the previously set value will be preserved on the Snowflake side, not the default value. That's due to Snowflake limitation and the lack of unset functionality for this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#start_timestamp ResourceMonitor#start_timestamp}
	StartTimestamp *string `field:"optional" json:"startTimestamp" yaml:"startTimestamp"`
	// Represents a numeric value specified as a percentage of the credit quota.
	//
	// Values over 100 are supported. After reaching this value, all assigned warehouses immediately cancel any currently running queries or statements. In addition, this action sends a notification to all users who have enabled notifications for themselves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#suspend_immediate_trigger ResourceMonitor#suspend_immediate_trigger}
	SuspendImmediateTrigger *float64 `field:"optional" json:"suspendImmediateTrigger" yaml:"suspendImmediateTrigger"`
	// Represents a numeric value specified as a percentage of the credit quota.
	//
	// Values over 100 are supported. After reaching this value, all assigned warehouses while allowing currently running queries to complete will be suspended. No new queries can be executed by the warehouses until the credit quota for the resource monitor is increased. In addition, this action sends a notification to all users who have enabled notifications for themselves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#suspend_trigger ResourceMonitor#suspend_trigger}
	SuspendTrigger *float64 `field:"optional" json:"suspendTrigger" yaml:"suspendTrigger"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/resource_monitor#timeouts ResourceMonitor#timeouts}
	Timeouts *ResourceMonitorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

