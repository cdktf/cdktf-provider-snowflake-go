package datasnowflakeexternalfunctions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeExternalFunctionsConfig struct {
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
	// The database from which to return the schemas from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/external_functions#database DataSnowflakeExternalFunctions#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The schema from which to return the external functions from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/external_functions#schema DataSnowflakeExternalFunctions#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/d/external_functions#id DataSnowflakeExternalFunctions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
