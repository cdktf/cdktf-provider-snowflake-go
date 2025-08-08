// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package failovergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FailoverGroupConfig struct {
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
	// Specifies the identifier for the failover group.
	//
	// The identifier must start with an alphabetic character and cannot contain spaces or special characters unless the identifier string is enclosed in double quotes (e.g. "My object"). Identifiers enclosed in double quotes are also case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#name FailoverGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the target account or list of target accounts to which replication and failover of specified objects from the source account is enabled.
	//
	// Secondary failover groups in the target accounts in this list can be promoted to serve as the primary failover group in case of failover. Expected in the form <org_name>.<target_account_name>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#allowed_accounts FailoverGroup#allowed_accounts}
	AllowedAccounts *[]*string `field:"optional" json:"allowedAccounts" yaml:"allowedAccounts"`
	// Specifies the database or list of databases for which you are enabling replication and failover from the source account to the target account.
	//
	// The OBJECT_TYPES list must include DATABASES to set this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#allowed_databases FailoverGroup#allowed_databases}
	AllowedDatabases *[]*string `field:"optional" json:"allowedDatabases" yaml:"allowedDatabases"`
	// Type(s) of integrations for which you are enabling replication and failover from the source account to the target account.
	//
	// This property requires that the OBJECT_TYPES list include INTEGRATIONS to set this parameter. The following integration types are supported: "SECURITY INTEGRATIONS", "API INTEGRATIONS", "STORAGE INTEGRATIONS", "EXTERNAL ACCESS INTEGRATIONS", "NOTIFICATION INTEGRATIONS"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#allowed_integration_types FailoverGroup#allowed_integration_types}
	AllowedIntegrationTypes *[]*string `field:"optional" json:"allowedIntegrationTypes" yaml:"allowedIntegrationTypes"`
	// Specifies the share or list of shares for which you are enabling replication and failover from the source account to the target account.
	//
	// The OBJECT_TYPES list must include SHARES to set this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#allowed_shares FailoverGroup#allowed_shares}
	AllowedShares *[]*string `field:"optional" json:"allowedShares" yaml:"allowedShares"`
	// from_replica block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#from_replica FailoverGroup#from_replica}
	FromReplica *FailoverGroupFromReplica `field:"optional" json:"fromReplica" yaml:"fromReplica"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#id FailoverGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// (Default: `false`) Allows replicating objects to accounts on lower editions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#ignore_edition_check FailoverGroup#ignore_edition_check}
	IgnoreEditionCheck interface{} `field:"optional" json:"ignoreEditionCheck" yaml:"ignoreEditionCheck"`
	// Type(s) of objects for which you are enabling replication and failover from the source account to the target account.
	//
	// The following object types are supported: "ACCOUNT PARAMETERS", "DATABASES", "INTEGRATIONS", "NETWORK POLICIES", "RESOURCE MONITORS", "ROLES", "SHARES", "USERS", "WAREHOUSES"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#object_types FailoverGroup#object_types}
	ObjectTypes *[]*string `field:"optional" json:"objectTypes" yaml:"objectTypes"`
	// replication_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#replication_schedule FailoverGroup#replication_schedule}
	ReplicationSchedule *FailoverGroupReplicationSchedule `field:"optional" json:"replicationSchedule" yaml:"replicationSchedule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.5.0/docs/resources/failover_group#timeouts FailoverGroup#timeouts}
	Timeouts *FailoverGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

