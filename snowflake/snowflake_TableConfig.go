// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TableConfig struct {
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
	// column block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#column Table#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The database in which to create the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#database Table#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the table;
	//
	// must be unique for the database and schema in which the table is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#name Table#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#schema Table#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies whether to enable change tracking on the table. Default false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#change_tracking Table#change_tracking}
	ChangeTracking interface{} `field:"optional" json:"changeTracking" yaml:"changeTracking"`
	// A list of one or more table columns/expressions to be used as clustering key(s) for the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#cluster_by Table#cluster_by}
	ClusterBy *[]*string `field:"optional" json:"clusterBy" yaml:"clusterBy"`
	// Specifies a comment for the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#comment Table#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the retention period for the table so that Time Travel actions (SELECT, CLONE, UNDROP) can be performed on historical data in the table.
	//
	// Default value is 1, if you wish to inherit the parent schema setting then pass in the schema attribute to this argument.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#data_retention_days Table#data_retention_days}
	DataRetentionDays *float64 `field:"optional" json:"dataRetentionDays" yaml:"dataRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#id Table#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// primary_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#primary_key Table#primary_key}
	PrimaryKey *TablePrimaryKey `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/table#tag Table#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

