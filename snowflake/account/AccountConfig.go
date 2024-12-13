// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package account

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountConfig struct {
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
	// Login name of the initial administrative user of the account.
	//
	// A new user is created in the new account with this name and password and granted the ACCOUNTADMIN role in the account. A login name can be any string consisting of letters, numbers, and underscores. Login names are always case-insensitive. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#admin_name Account#admin_name}
	AdminName *string `field:"required" json:"adminName" yaml:"adminName"`
	// Snowflake Edition of the account.
	//
	// See more about Snowflake Editions in the [official documentation](https://docs.snowflake.com/en/user-guide/intro-editions). Valid options are: `STANDARD` | `ENTERPRISE` | `BUSINESS_CRITICAL`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#edition Account#edition}
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// Email address of the initial administrative user of the account.
	//
	// This email address is used to send any notifications about the account. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#email Account#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Specifies the number of days during which the account can be restored (“undropped”).
	//
	// The minimum is 3 days and the maximum is 90 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#grace_period_in_days Account#grace_period_in_days}
	GracePeriodInDays *float64 `field:"required" json:"gracePeriodInDays" yaml:"gracePeriodInDays"`
	// Specifies the identifier (i.e. name) for the account. It must be unique within an organization, regardless of which Snowflake Region the account is in and must start with an alphabetic character and cannot contain spaces or special characters except for underscores (_). Note that if the account name includes underscores, features that do not accept account names with underscores (e.g. Okta SSO or SCIM) can reference a version of the account name that substitutes hyphens (-) for the underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#name Account#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Password for the initial administrative user of the account.
	//
	// Either admin_password or admin_rsa_public_key has to be specified. This field cannot be used whenever admin_user_type is set to SERVICE. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#admin_password Account#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Assigns a public key to the initial administrative user of the account.
	//
	// Either admin_password or admin_rsa_public_key has to be specified. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#admin_rsa_public_key Account#admin_rsa_public_key}
	AdminRsaPublicKey *string `field:"optional" json:"adminRsaPublicKey" yaml:"adminRsaPublicKey"`
	// Used for setting the type of the first user that is assigned the ACCOUNTADMIN role during account creation.
	//
	// Valid options are: `PERSON` | `SERVICE` | `LEGACY_SERVICE` External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#admin_user_type Account#admin_user_type}
	AdminUserType *string `field:"optional" json:"adminUserType" yaml:"adminUserType"`
	// Specifies a comment for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#comment Account#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// First name of the initial administrative user of the account.
	//
	// This field cannot be used whenever admin_user_type is set to SERVICE. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#first_name Account#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#id Account#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Sets an account property that determines whether the ORGADMIN role is enabled in the account.
	//
	// Only an organization administrator (i.e. user with the ORGADMIN role) can set the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#is_org_admin Account#is_org_admin}
	IsOrgAdmin *string `field:"optional" json:"isOrgAdmin" yaml:"isOrgAdmin"`
	// Last name of the initial administrative user of the account.
	//
	// This field cannot be used whenever admin_user_type is set to SERVICE. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#last_name Account#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Specifies whether the new user created to administer the account is forced to change their password upon first login into the account.
	//
	// This field cannot be used whenever admin_user_type is set to SERVICE. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#must_change_password Account#must_change_password}
	MustChangePassword *string `field:"optional" json:"mustChangePassword" yaml:"mustChangePassword"`
	// [Snowflake Region ID](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#label-snowflake-region-ids) of the region where the account is created. If no value is provided, Snowflake creates the account in the same Snowflake Region as the current account (i.e. the account in which the CREATE ACCOUNT statement is executed.).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#region Account#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ID of the region group where the account is created.
	//
	// To retrieve the region group ID for existing accounts in your organization, execute the [SHOW REGIONS](https://docs.snowflake.com/en/sql-reference/sql/show-regions) command. For information about when you might need to specify region group, see [Region groups](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#label-region-groups).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/account#region_group Account#region_group}
	RegionGroup *string `field:"optional" json:"regionGroup" yaml:"regionGroup"`
}

