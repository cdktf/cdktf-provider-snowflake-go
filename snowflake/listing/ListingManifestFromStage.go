// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package listing


type ListingManifestFromStage struct {
	// Identifier of the stage where the manifest file is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/listing#stage Listing#stage}
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// Location of the manifest file in the stage.
	//
	// If not specified, the manifest file will be expected to be at the root of the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/listing#location Listing#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Specifies a comment for the listing version.
	//
	// Whenever a new version is created, this comment will be associated with it. The comment on the version will be visible in the [SHOW VERSIONS IN LISTING](https://docs.snowflake.com/en/sql-reference/sql/show-versions-in-listing) command output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/listing#version_comment Listing#version_comment}
	VersionComment *string `field:"optional" json:"versionComment" yaml:"versionComment"`
	// Represents manifest version name.
	//
	// It's case-sensitive and used in manifest versioning. Version name should be specified or changed whenever any changes in the manifest should be applied to the listing. Later on the versions of the listing can be analyzed by calling the [SHOW VERSIONS IN LISTING](https://docs.snowflake.com/en/sql-reference/sql/show-versions-in-listing) command. The resource does not track the changes on the specified stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/listing#version_name Listing#version_name}
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

