// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view


type ViewDataMetricSchedule struct {
	// Specifies an interval (in minutes) of wait time inserted between runs of the data metric function.
	//
	// Conflicts with `using_cron`. Valid values are: `5` | `15` | `30` | `60` | `720` | `1440`. Due to Snowflake limitations, changes in this field are not managed by the provider. Please consider using [taint](https://developer.hashicorp.com/terraform/cli/commands/taint) command, `using_cron` field, or [replace_triggered_by](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#replace_triggered_by) metadata argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/view#minutes View#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Specifies a cron expression and time zone for periodically running the data metric function.
	//
	// Supports a subset of standard cron utility syntax. Conflicts with `minutes`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.4/docs/resources/view#using_cron View#using_cron}
	UsingCron *string `field:"optional" json:"usingCron" yaml:"usingCron"`
}

