// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionpython


type FunctionPythonSecrets struct {
	// Fully qualified name of the allowed [secret](https://docs.snowflake.com/en/sql-reference/sql/create-secret). You will receive an error if you specify a SECRETS value whose secret isn’t also included in an integration specified by the EXTERNAL_ACCESS_INTEGRATIONS parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_python#secret_id FunctionPython#secret_id}
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// The variable that will be used in handler code when retrieving information from the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/function_python#secret_variable_name FunctionPython#secret_variable_name}
	SecretVariableName *string `field:"required" json:"secretVariableName" yaml:"secretVariableName"`
}

