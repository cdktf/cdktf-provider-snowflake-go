// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	// Specifies the object identifiers for the tag association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/tag_association#object_identifiers TagAssociation#object_identifiers}
	ObjectIdentifiers *[]*string `field:"required" json:"objectIdentifiers" yaml:"objectIdentifiers"`
	// Specifies the type of object to add a tag.
	//
	// Allowed object types: [ACCOUNT APPLICATION APPLICATION PACKAGE DATABASE FAILOVER GROUP INTEGRATION NETWORK POLICY REPLICATION GROUP ROLE SHARE USER WAREHOUSE DATABASE ROLE SCHEMA ALERT SNOWFLAKE.CORE.BUDGET SNOWFLAKE.ML.CLASSIFICATION EXTERNAL FUNCTION EXTERNAL TABLE FUNCTION GIT REPOSITORY ICEBERG TABLE MATERIALIZED VIEW PIPE MASKING POLICY PASSWORD POLICY ROW ACCESS POLICY SESSION POLICY PRIVACY POLICY PROCEDURE STAGE STREAM TABLE TASK VIEW COLUMN EVENT TABLE].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/tag_association#object_type TagAssociation#object_type}
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Specifies the identifier for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/tag_association#tag_id TagAssociation#tag_id}
	TagId *string `field:"required" json:"tagId" yaml:"tagId"`
	// Specifies the value of the tag, (e.g. 'finance' or 'engineering').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/tag_association#tag_value TagAssociation#tag_value}
	TagValue *string `field:"required" json:"tagValue" yaml:"tagValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/tag_association#id TagAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, skips validation of the tag association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/tag_association#skip_validation TagAssociation#skip_validation}
	SkipValidation interface{} `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/tag_association#timeouts TagAssociation#timeouts}
	Timeouts *TagAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

