// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alert


type AlertAlertScheduleCron struct {
	// Specifies the cron expression for the alert.
	//
	// The cron expression must be in the following format: "minute hour day-of-month month day-of-week". The following values are supported: minute: 0-59 hour: 0-23 day-of-month: 1-31 month: 1-12 day-of-week: 0-6 (0 is Sunday)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/alert#expression Alert#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Specifies the time zone for alert refresh.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/alert#time_zone Alert#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

