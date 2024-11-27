// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseold

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseOldConfig struct {
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
	// Specifies the identifier for the database; must be unique for your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#name DatabaseOld#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies a comment for the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#comment DatabaseOld#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Number of days for which Snowflake retains historical data for performing Time Travel actions (SELECT, CLONE, UNDROP) on the object.
	//
	// A value of 0 effectively disables Time Travel for the specified database. Default value for this field is set to -1, which is a fallback to use Snowflake default. For more information, see [Understanding & Using Time Travel](https://docs.snowflake.com/en/user-guide/data-time-travel).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#data_retention_time_in_days DatabaseOld#data_retention_time_in_days}
	DataRetentionTimeInDays *float64 `field:"optional" json:"dataRetentionTimeInDays" yaml:"dataRetentionTimeInDays"`
	// Specify a database to create a clone from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#from_database DatabaseOld#from_database}
	FromDatabase *string `field:"optional" json:"fromDatabase" yaml:"fromDatabase"`
	// Specify a fully-qualified path to a database to create a replica from.
	//
	// A fully qualified path follows the format of `"<organization_name>"."<account_name>"."<db_name>"`. An example would be: `"myorg1"."account1"."db1"`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#from_replica DatabaseOld#from_replica}
	FromReplica *string `field:"optional" json:"fromReplica" yaml:"fromReplica"`
	// Specify a provider and a share in this map to create a database from a share.
	//
	// As of version 0.87.0, the provider field is the account locator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#from_share DatabaseOld#from_share}
	FromShare *map[string]*string `field:"optional" json:"fromShare" yaml:"fromShare"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#id DatabaseOld#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies a database as transient.
	//
	// Transient databases do not have a Fail-safe period so they do not incur additional storage costs once they leave Time Travel; however, this means they are also not protected by Fail-safe in the event of a data loss.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#is_transient DatabaseOld#is_transient}
	IsTransient interface{} `field:"optional" json:"isTransient" yaml:"isTransient"`
	// replication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.99.0/docs/resources/database_old#replication_configuration DatabaseOld#replication_configuration}
	ReplicationConfiguration *DatabaseOldReplicationConfiguration `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
}

