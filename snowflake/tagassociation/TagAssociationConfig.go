package tagassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TagAssociationConfig struct {
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
	// object_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/tag_association#object_identifier TagAssociation#object_identifier}
	ObjectIdentifier interface{} `field:"required" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the type of object to add a tag to. ex: 'ACCOUNT', 'COLUMN', 'DATABASE', etc. For more information: https://docs.snowflake.com/en/user-guide/object-tagging.html#supported-objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/tag_association#object_type TagAssociation#object_type}
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Specifies the identifier for the tag. Note: format must follow: "databaseName"."schemaName"."tagName" or "databaseName.schemaName.tagName" or "databaseName|schemaName.tagName" (snowflake_tag.tag.id).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/tag_association#tag_id TagAssociation#tag_id}
	TagId *string `field:"required" json:"tagId" yaml:"tagId"`
	// Specifies the value of the tag, (e.g. 'finance' or 'engineering').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/tag_association#tag_value TagAssociation#tag_value}
	TagValue *string `field:"required" json:"tagValue" yaml:"tagValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/tag_association#id TagAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the object identifier for the tag association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/tag_association#object_name TagAssociation#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// If true, skips validation of the tag association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/tag_association#skip_validation TagAssociation#skip_validation}
	SkipValidation interface{} `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.65.0/docs/resources/tag_association#timeouts TagAssociation#timeouts}
	Timeouts *TagAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

