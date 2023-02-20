package maskingpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaskingPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The database in which to create the masking policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/masking_policy#database MaskingPolicy#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the SQL expression that transforms the data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/masking_policy#masking_expression MaskingPolicy#masking_expression}
	MaskingExpression *string `field:"required" json:"maskingExpression" yaml:"maskingExpression"`
	// Specifies the identifier for the masking policy;
	//
	// must be unique for the database and schema in which the masking policy is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/masking_policy#name MaskingPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the data type to return.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/masking_policy#return_data_type MaskingPolicy#return_data_type}
	ReturnDataType *string `field:"required" json:"returnDataType" yaml:"returnDataType"`
	// The schema in which to create the masking policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/masking_policy#schema MaskingPolicy#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the data type to mask.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/masking_policy#value_data_type MaskingPolicy#value_data_type}
	ValueDataType *string `field:"required" json:"valueDataType" yaml:"valueDataType"`
	// Specifies a comment for the masking policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/masking_policy#comment MaskingPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/masking_policy#id MaskingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

