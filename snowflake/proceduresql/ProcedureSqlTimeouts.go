// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package proceduresql


type ProcedureSqlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#create ProcedureSql#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#delete ProcedureSql#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#read ProcedureSql#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/procedure_sql#update ProcedureSql#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

