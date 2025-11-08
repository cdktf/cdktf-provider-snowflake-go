// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gitrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GitRepositoryConfig struct {
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
	// Identifier of API INTEGRATION containing information about the remote Git repository such as allowed credentials and prefixes for target URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#api_integration GitRepository#api_integration}
	ApiIntegration *string `field:"required" json:"apiIntegration" yaml:"apiIntegration"`
	// The database in which to create the git repository.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#database GitRepository#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the git repository;
	//
	// must be unique for the schema in which the git repository is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#name GitRepository#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the origin URL of the remote Git repository that this Git repository clone represents.
	//
	// The URL must use HTTPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#origin GitRepository#origin}
	Origin *string `field:"required" json:"origin" yaml:"origin"`
	// The schema in which to create the git repository.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#schema GitRepository#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies a comment for the git repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#comment GitRepository#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the Snowflake secret fully qualified name (e.g `"\"<db_name>\".\"<schema_name>\".\"<secret_name>\""`) containing the credentials to use for authenticating with the remote Git repository. Omit this parameter to use the default secret specified by the API integration or if this integration does not require authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#git_credentials GitRepository#git_credentials}
	GitCredentials *string `field:"optional" json:"gitCredentials" yaml:"gitCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#id GitRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/git_repository#timeouts GitRepository#timeouts}
	Timeouts *GitRepositoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

