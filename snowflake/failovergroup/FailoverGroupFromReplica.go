package failovergroup


type FailoverGroupFromReplica struct {
	// Identifier for the primary failover group in the source account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/failover_group#name FailoverGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of your Snowflake organization.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/failover_group#organization_name FailoverGroup#organization_name}
	OrganizationName *string `field:"required" json:"organizationName" yaml:"organizationName"`
	// Source account from which you are enabling replication and failover of the specified objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/failover_group#source_account_name FailoverGroup#source_account_name}
	SourceAccountName *string `field:"required" json:"sourceAccountName" yaml:"sourceAccountName"`
}

