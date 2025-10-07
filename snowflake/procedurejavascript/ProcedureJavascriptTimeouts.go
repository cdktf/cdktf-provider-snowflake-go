// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package procedurejavascript


type ProcedureJavascriptTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/procedure_javascript#create ProcedureJavascript#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/procedure_javascript#delete ProcedureJavascript#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/procedure_javascript#read ProcedureJavascript#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/procedure_javascript#update ProcedureJavascript#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

