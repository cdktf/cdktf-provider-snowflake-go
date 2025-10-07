// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputePoolConfig struct {
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
	// Identifies the type of machine you want to provision for the nodes in the compute pool.
	//
	// Valid values are (case-insensitive): `CPU_X64_XS` | `CPU_X64_S` | `CPU_X64_M` | `CPU_X64_L` | `HIGHMEM_X64_S` | `HIGHMEM_X64_M` | `HIGHMEM_X64_L` | `HIGHMEM_X64_SL` | `GPU_NV_S` | `GPU_NV_M` | `GPU_NV_L` | `GPU_NV_XS` | `GPU_NV_SM` | `GPU_NV_2M` | `GPU_NV_3M` | `GPU_NV_SL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#instance_family ComputePool#instance_family}
	InstanceFamily *string `field:"required" json:"instanceFamily" yaml:"instanceFamily"`
	// Specifies the maximum number of nodes for the compute pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#max_nodes ComputePool#max_nodes}
	MaxNodes *float64 `field:"required" json:"maxNodes" yaml:"maxNodes"`
	// Specifies the minimum number of nodes for the compute pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#min_nodes ComputePool#min_nodes}
	MinNodes *float64 `field:"required" json:"minNodes" yaml:"minNodes"`
	// Specifies the identifier for the compute pool;
	//
	// must be unique for the account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#name ComputePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to automatically resume a compute pool when a service or job is submitted to it.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#auto_resume ComputePool#auto_resume}
	AutoResume *string `field:"optional" json:"autoResume" yaml:"autoResume"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Number of seconds of inactivity after which you want Snowflake to automatically suspend the compute pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#auto_suspend_secs ComputePool#auto_suspend_secs}
	AutoSuspendSecs *float64 `field:"optional" json:"autoSuspendSecs" yaml:"autoSuspendSecs"`
	// Specifies a comment for the compute pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#comment ComputePool#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the Snowflake Native App name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#for_application ComputePool#for_application}
	ForApplication *string `field:"optional" json:"forApplication" yaml:"forApplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#id ComputePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether the compute pool is created initially in the suspended state.
	//
	// This field is used only when creating a compute pool. Changes on this field are ignored after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#initially_suspended ComputePool#initially_suspended}
	InitiallySuspended *string `field:"optional" json:"initiallySuspended" yaml:"initiallySuspended"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/compute_pool#timeouts ComputePool#timeouts}
	Timeouts *ComputePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

