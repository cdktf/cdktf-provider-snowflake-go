// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceConfig struct {
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
	// Specifies the name of the compute pool in your account on which to run the service.
	//
	// Identifiers with special or lower-case characters are not supported. This limitation in the provider follows the limitation in Snowflake (see [docs](https://docs.snowflake.com/en/sql-reference/sql/create-compute-pool)). Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#compute_pool Service#compute_pool}
	ComputePool *string `field:"required" json:"computePool" yaml:"computePool"`
	// The database in which to create the service.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#database Service#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the service;
	//
	// must be unique for the schema in which the service is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#name Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the service.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#schema Service#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to automatically resume a service.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#auto_resume Service#auto_resume}
	AutoResume *string `field:"optional" json:"autoResume" yaml:"autoResume"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the number of seconds of inactivity (service is idle) after which Snowflake automatically suspends the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#auto_suspend_secs Service#auto_suspend_secs}
	AutoSuspendSecs *float64 `field:"optional" json:"autoSuspendSecs" yaml:"autoSuspendSecs"`
	// Specifies a comment for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#comment Service#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the names of the external access integrations that allow your service to access external sites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#external_access_integrations Service#external_access_integrations}
	ExternalAccessIntegrations *[]*string `field:"optional" json:"externalAccessIntegrations" yaml:"externalAccessIntegrations"`
	// from_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#from_specification Service#from_specification}
	FromSpecification *ServiceFromSpecification `field:"optional" json:"fromSpecification" yaml:"fromSpecification"`
	// from_specification_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#from_specification_template Service#from_specification_template}
	FromSpecificationTemplate *ServiceFromSpecificationTemplate `field:"optional" json:"fromSpecificationTemplate" yaml:"fromSpecificationTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#id Service#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the maximum number of service instances to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#max_instances Service#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// Specifies the minimum number of service instances to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#min_instances Service#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// Indicates the minimum service instances that must be ready for Snowflake to consider the service is ready to process requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#min_ready_instances Service#min_ready_instances}
	MinReadyInstances *float64 `field:"optional" json:"minReadyInstances" yaml:"minReadyInstances"`
	// Warehouse to use if a service container connects to Snowflake to execute a query but does not explicitly specify a warehouse to use.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#query_warehouse Service#query_warehouse}
	QueryWarehouse *string `field:"optional" json:"queryWarehouse" yaml:"queryWarehouse"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.0/docs/resources/service#timeouts Service#timeouts}
	Timeouts *ServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

