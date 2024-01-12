// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alert


type AlertAlertSchedule struct {
	// cron block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/alert#cron Alert#cron}
	Cron *AlertAlertScheduleCron `field:"optional" json:"cron" yaml:"cron"`
	// Specifies the interval in minutes for the alert schedule.
	//
	// The interval must be greater than 0 and less than 1440 (24 hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.83.1/docs/resources/alert#interval Alert#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

