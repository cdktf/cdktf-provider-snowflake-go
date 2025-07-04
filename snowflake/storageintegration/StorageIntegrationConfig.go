// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#name StorageIntegration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Explicitly limits external stages that use the integration to reference one or more storage locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#storage_allowed_locations StorageIntegration#storage_allowed_locations}
	StorageAllowedLocations *[]*string `field:"required" json:"storageAllowedLocations" yaml:"storageAllowedLocations"`
	// Specifies the storage provider for the integration. Valid options are: `S3` | `S3GOV` | `S3CHINA` | `GCS` | `AZURE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#storage_provider StorageIntegration#storage_provider}
	StorageProvider *string `field:"required" json:"storageProvider" yaml:"storageProvider"`
	// (Default: ``).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#azure_tenant_id StorageIntegration#azure_tenant_id}
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// (Default: ``).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#comment StorageIntegration#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// (Default: `true`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#enabled StorageIntegration#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#id StorageIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// "bucket-owner-full-control" Enables support for AWS access control lists (ACLs) to grant the bucket owner full control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#storage_aws_object_acl StorageIntegration#storage_aws_object_acl}
	StorageAwsObjectAcl *string `field:"optional" json:"storageAwsObjectAcl" yaml:"storageAwsObjectAcl"`
	// (Default: ``).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#storage_aws_role_arn StorageIntegration#storage_aws_role_arn}
	StorageAwsRoleArn *string `field:"optional" json:"storageAwsRoleArn" yaml:"storageAwsRoleArn"`
	// Explicitly prohibits external stages that use the integration from referencing one or more storage locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#storage_blocked_locations StorageIntegration#storage_blocked_locations}
	StorageBlockedLocations *[]*string `field:"optional" json:"storageBlockedLocations" yaml:"storageBlockedLocations"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#timeouts StorageIntegration#timeouts}
	Timeouts *StorageIntegrationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// (Default: `EXTERNAL_STAGE`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.3.0/docs/resources/storage_integration#type StorageIntegration#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

