// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package task


type TaskSchedule struct {
	// Specifies an interval (in minutes) of wait time inserted between runs of the task.
	//
	// Accepts positive integers only. (conflicts with `using_cron`)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/task#minutes Task#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Specifies a cron expression and time zone for periodically running the task.
	//
	// Supports a subset of standard cron utility syntax. (conflicts with `minutes`)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/task#using_cron Task#using_cron}
	UsingCron *string `field:"optional" json:"usingCron" yaml:"usingCron"`
}

