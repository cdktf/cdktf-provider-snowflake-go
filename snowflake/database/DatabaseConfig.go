// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package database

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#name Database#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#comment Database#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Number of days for which Snowflake retains historical data for performing Time Travel actions (SELECT, CLONE, UNDROP) on the object.
	//
	// A value of 0 effectively disables Time Travel for the specified database, schema, or table. For more information, see Understanding & Using Time Travel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#data_retention_time_in_days Database#data_retention_time_in_days}
	DataRetentionTimeInDays *float64 `field:"optional" json:"dataRetentionTimeInDays" yaml:"dataRetentionTimeInDays"`
	// Specify a database to create a clone from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#from_database Database#from_database}
	FromDatabase *string `field:"optional" json:"fromDatabase" yaml:"fromDatabase"`
	// Specify a fully-qualified path to a database to create a replica from.
	//
	// A fully qualified path follows the format of "<organization_name>"."<account_name>"."<db_name>". An example would be: "myorg1"."account1"."db1"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#from_replica Database#from_replica}
	FromReplica *string `field:"optional" json:"fromReplica" yaml:"fromReplica"`
	// Specify a provider and a share in this map to create a database from a share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#from_share Database#from_share}
	FromShare *map[string]*string `field:"optional" json:"fromShare" yaml:"fromShare"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#id Database#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies a database as transient.
	//
	// Transient databases do not have a Fail-safe period so they do not incur additional storage costs once they leave Time Travel; however, this means they are also not protected by Fail-safe in the event of a data loss.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#is_transient Database#is_transient}
	IsTransient interface{} `field:"optional" json:"isTransient" yaml:"isTransient"`
	// replication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.71.0/docs/resources/database#replication_configuration Database#replication_configuration}
	ReplicationConfiguration *DatabaseReplicationConfiguration `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
}

