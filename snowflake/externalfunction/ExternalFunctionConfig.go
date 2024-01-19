// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalfunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExternalFunctionConfig struct {
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
	// The name of the API integration object that should be used to authenticate the call to the proxy service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#api_integration ExternalFunction#api_integration}
	ApiIntegration *string `field:"required" json:"apiIntegration" yaml:"apiIntegration"`
	// The database in which to create the external function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#database ExternalFunction#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the external function.
	//
	// The identifier can contain the schema name and database name, as well as the function name. The function's signature (name and argument data types) must be unique within the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#name ExternalFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the behavior of the function when returning results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#return_behavior ExternalFunction#return_behavior}
	ReturnBehavior *string `field:"required" json:"returnBehavior" yaml:"returnBehavior"`
	// Specifies the data type returned by the external function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#return_type ExternalFunction#return_type}
	ReturnType *string `field:"required" json:"returnType" yaml:"returnType"`
	// The schema in which to create the external function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#schema ExternalFunction#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// This is the invocation URL of the proxy service and resource through which Snowflake calls the remote service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#url_of_proxy_and_resource ExternalFunction#url_of_proxy_and_resource}
	UrlOfProxyAndResource *string `field:"required" json:"urlOfProxyAndResource" yaml:"urlOfProxyAndResource"`
	// arg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#arg ExternalFunction#arg}
	Arg interface{} `field:"optional" json:"arg" yaml:"arg"`
	// A description of the external function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#comment ExternalFunction#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// If specified, the JSON payload is compressed when sent from Snowflake to the proxy service, and when sent back from the proxy service to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#compression ExternalFunction#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Binds Snowflake context function results to HTTP headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#context_headers ExternalFunction#context_headers}
	ContextHeaders *[]*string `field:"optional" json:"contextHeaders" yaml:"contextHeaders"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#header ExternalFunction#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#id ExternalFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This specifies the maximum number of rows in each batch sent to the proxy service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#max_batch_rows ExternalFunction#max_batch_rows}
	MaxBatchRows *float64 `field:"optional" json:"maxBatchRows" yaml:"maxBatchRows"`
	// Specifies the behavior of the external function when called with null inputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#null_input_behavior ExternalFunction#null_input_behavior}
	NullInputBehavior *string `field:"optional" json:"nullInputBehavior" yaml:"nullInputBehavior"`
	// This specifies the name of the request translator function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#request_translator ExternalFunction#request_translator}
	RequestTranslator *string `field:"optional" json:"requestTranslator" yaml:"requestTranslator"`
	// This specifies the name of the response translator function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#response_translator ExternalFunction#response_translator}
	ResponseTranslator *string `field:"optional" json:"responseTranslator" yaml:"responseTranslator"`
	// Indicates whether the function can return NULL values or must return only NON-NULL values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.84.0/docs/resources/external_function#return_null_allowed ExternalFunction#return_null_allowed}
	ReturnNullAllowed interface{} `field:"optional" json:"returnNullAllowed" yaml:"returnNullAllowed"`
}

