// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maskingpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaskingPolicyConfig struct {
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
	// argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#argument MaskingPolicy#argument}
	Argument interface{} `field:"required" json:"argument" yaml:"argument"`
	// Specifies the SQL expression that transforms the data.
	//
	// To mitigate permadiff on this field, the provider replaces blank characters with a space. This can lead to false positives in cases where a change in case or run of whitespace is semantically significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#body MaskingPolicy#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// The database in which to create the masking policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#database MaskingPolicy#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the masking policy;
	//
	// must be unique for the database and schema in which the masking policy is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#name MaskingPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The return data type must match the input data type of the first column that is specified as an input column.
	//
	// For more information about data types, check [Snowflake docs](https://docs.snowflake.com/en/sql-reference/intro-summary-data-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#return_data_type MaskingPolicy#return_data_type}
	ReturnDataType *string `field:"required" json:"returnDataType" yaml:"returnDataType"`
	// The schema in which to create the masking policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#schema MaskingPolicy#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies a comment for the masking policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#comment MaskingPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether the row access policy or conditional masking policy can reference a column that is already protected by a masking policy.
	//
	// Due to Snowflake limitations, when value is chenged, the resource is recreated. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#exempt_other_policies MaskingPolicy#exempt_other_policies}
	ExemptOtherPolicies *string `field:"optional" json:"exemptOtherPolicies" yaml:"exemptOtherPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.3/docs/resources/masking_policy#id MaskingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

