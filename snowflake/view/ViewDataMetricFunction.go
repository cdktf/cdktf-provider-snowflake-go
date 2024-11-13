// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package view


type ViewDataMetricFunction struct {
	// Identifier of the data metric function to add to the table or view or drop from the table or view.
	//
	// This function identifier must be provided without arguments in parenthesis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#function_name View#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The table or view columns on which to associate the data metric function.
	//
	// The data types of the columns must match the data types of the columns specified in the data metric function definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#on View#on}
	On *[]*string `field:"required" json:"on" yaml:"on"`
	// The status of the metrics association.
	//
	// Valid values are: `STARTED` | `SUSPENDED`. When status of a data metric function is changed, it is being reassigned with `DROP DATA METRIC FUNCTION` and `ADD DATA METRIC FUNCTION`, and then its status is changed by `MODIFY DATA METRIC FUNCTION`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.98.0/docs/resources/view#schedule_status View#schedule_status}
	ScheduleStatus *string `field:"required" json:"scheduleStatus" yaml:"scheduleStatus"`
}

