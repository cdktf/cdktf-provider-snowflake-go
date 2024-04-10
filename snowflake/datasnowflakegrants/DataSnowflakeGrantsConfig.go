// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeGrantsConfig struct {
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
	// future_grants_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.88.0/docs/data-sources/grants#future_grants_in DataSnowflakeGrants#future_grants_in}
	FutureGrantsIn *DataSnowflakeGrantsFutureGrantsIn `field:"optional" json:"futureGrantsIn" yaml:"futureGrantsIn"`
	// future_grants_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.88.0/docs/data-sources/grants#future_grants_to DataSnowflakeGrants#future_grants_to}
	FutureGrantsTo *DataSnowflakeGrantsFutureGrantsTo `field:"optional" json:"futureGrantsTo" yaml:"futureGrantsTo"`
	// grants_of block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.88.0/docs/data-sources/grants#grants_of DataSnowflakeGrants#grants_of}
	GrantsOf *DataSnowflakeGrantsGrantsOf `field:"optional" json:"grantsOf" yaml:"grantsOf"`
	// grants_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.88.0/docs/data-sources/grants#grants_on DataSnowflakeGrants#grants_on}
	GrantsOn *DataSnowflakeGrantsGrantsOn `field:"optional" json:"grantsOn" yaml:"grantsOn"`
	// grants_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.88.0/docs/data-sources/grants#grants_to DataSnowflakeGrants#grants_to}
	GrantsTo *DataSnowflakeGrantsGrantsTo `field:"optional" json:"grantsTo" yaml:"grantsTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.88.0/docs/data-sources/grants#id DataSnowflakeGrants#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

