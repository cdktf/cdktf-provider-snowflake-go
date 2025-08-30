// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalvolume


type ExternalVolumeStorageLocation struct {
	// Specifies the base URL for your cloud storage location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/external_volume#storage_base_url ExternalVolume#storage_base_url}
	StorageBaseUrl *string `field:"required" json:"storageBaseUrl" yaml:"storageBaseUrl"`
	// Name of the storage location.
	//
	// Must be unique for the external volume. Do not use the name `terraform_provider_sentinel_storage_location` - this is reserved for the provider for performing update operations. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/external_volume#storage_location_name ExternalVolume#storage_location_name}
	StorageLocationName *string `field:"required" json:"storageLocationName" yaml:"storageLocationName"`
	// Specifies the cloud storage provider that stores your data files.
	//
	// Valid values are (case-insensitive): `GCS` | `AZURE` | `S3` | `S3GOV`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/external_volume#storage_provider ExternalVolume#storage_provider}
	StorageProvider *string `field:"required" json:"storageProvider" yaml:"storageProvider"`
	// Specifies the ID for your Office 365 tenant that the allowed and blocked storage accounts belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/external_volume#azure_tenant_id ExternalVolume#azure_tenant_id}
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// Specifies the ID for the KMS-managed key used to encrypt files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/external_volume#encryption_kms_key_id ExternalVolume#encryption_kms_key_id}
	EncryptionKmsKeyId *string `field:"optional" json:"encryptionKmsKeyId" yaml:"encryptionKmsKeyId"`
	// Specifies the encryption type used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/external_volume#encryption_type ExternalVolume#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Specifies the case-sensitive Amazon Resource Name (ARN) of the AWS identity and access management (IAM) role that grants privileges on the S3 bucket containing your data files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/external_volume#storage_aws_role_arn ExternalVolume#storage_aws_role_arn}
	StorageAwsRoleArn *string `field:"optional" json:"storageAwsRoleArn" yaml:"storageAwsRoleArn"`
}

