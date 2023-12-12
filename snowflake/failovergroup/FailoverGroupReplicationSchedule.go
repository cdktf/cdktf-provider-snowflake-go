// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package failovergroup


type FailoverGroupReplicationSchedule struct {
	// cron block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/failover_group#cron FailoverGroup#cron}
	Cron *FailoverGroupReplicationScheduleCron `field:"optional" json:"cron" yaml:"cron"`
	// Specifies the interval in minutes for the replication schedule.
	//
	// The interval must be greater than 0 and less than 1440 (24 hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.79.1/docs/resources/failover_group#interval FailoverGroup#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

