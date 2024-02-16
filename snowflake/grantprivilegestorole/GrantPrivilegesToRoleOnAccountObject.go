// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestorole


type GrantPrivilegesToRoleOnAccountObject struct {
	// The fully qualified name of the object on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.86.0/docs/resources/grant_privileges_to_role#object_name GrantPrivilegesToRole#object_name}
	ObjectName *string `field:"required" json:"objectName" yaml:"objectName"`
	// The object type of the account object on which privileges will be granted.
	//
	// Valid values are: USER | RESOURCE MONITOR | WAREHOUSE | DATABASE | INTEGRATION | FAILOVER GROUP | REPLICATION GROUP | EXTERNAL VOLUME
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.86.0/docs/resources/grant_privileges_to_role#object_type GrantPrivilegesToRole#object_type}
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
}

