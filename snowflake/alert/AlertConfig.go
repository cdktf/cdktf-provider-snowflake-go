// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertConfig struct {
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
	// The SQL statement that should be executed if the condition returns one or more rows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#action Alert#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The SQL statement that represents the condition for the alert. (SELECT, SHOW, CALL).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#condition Alert#condition}
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// The database in which to create the alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#database Alert#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the alert;
	//
	// must be unique for the database and schema in which the alert is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#name Alert#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#schema Alert#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The warehouse the alert will use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#warehouse Alert#warehouse}
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// alert_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#alert_schedule Alert#alert_schedule}
	AlertSchedule *AlertAlertSchedule `field:"optional" json:"alertSchedule" yaml:"alertSchedule"`
	// Specifies a comment for the alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#comment Alert#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies if an alert should be 'started' (enabled) after creation or should remain 'suspended' (default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#enabled Alert#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.2/docs/resources/alert#id Alert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

