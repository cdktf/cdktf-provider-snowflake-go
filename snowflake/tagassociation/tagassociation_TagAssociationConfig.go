package tagassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TagAssociationConfig struct {
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
	// Specifies the object identifier for the tag association.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#object_name TagAssociation#object_name}
	ObjectName *string `field:"required" json:"objectName" yaml:"objectName"`
	// Specifies the type of object to add a tag to. ex: 'ACCOUNT', 'COLUMN', 'DATABASE', etc. For more information: https://docs.snowflake.com/en/user-guide/object-tagging.html#supported-objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#object_type TagAssociation#object_type}
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Specifies the identifier for the tag. Note: format must follow: "databaseName"."schemaName"."tagName" or "databaseName.schemaName.tagName" or "databaseName|schemaName.tagName" (snowflake_tag.tag.id).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#tag_id TagAssociation#tag_id}
	TagId *string `field:"required" json:"tagId" yaml:"tagId"`
	// Specifies the value of the tag, (e.g. 'finance' or 'engineering').
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#tag_value TagAssociation#tag_value}
	TagValue *string `field:"required" json:"tagValue" yaml:"tagValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#id TagAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, skips validation of the tag association.
	//
	// It can take up to an hour for the SNOWFLAKE.TAG_REFERENCES table to update, and also requires ACCOUNT_ADMIN role to read from. https://docs.snowflake.com/en/sql-reference/account-usage/tag_references.html
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#skip_validation TagAssociation#skip_validation}
	SkipValidation interface{} `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/tag_association#timeouts TagAssociation#timeouts}
	Timeouts *TagAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

