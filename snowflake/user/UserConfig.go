// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
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
	// Name of the user. Note that if you do not supply login_name this will be used as login_name. [doc](https://docs.snowflake.net/manuals/sql-reference/sql/create-user.html#required-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#name User#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#comment User#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the namespace (database only or database and schema) that is active by default for the user’s session upon login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#default_namespace User#default_namespace}
	DefaultNamespace *string `field:"optional" json:"defaultNamespace" yaml:"defaultNamespace"`
	// Specifies the role that is active by default for the user’s session upon login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#default_role User#default_role}
	DefaultRole *string `field:"optional" json:"defaultRole" yaml:"defaultRole"`
	// Specifies the set of secondary roles that are active for the user’s session upon login.
	//
	// Currently only ["ALL"] value is supported - more information can be found in [doc](https://docs.snowflake.com/en/sql-reference/sql/create-user#optional-object-properties-objectproperties)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#default_secondary_roles User#default_secondary_roles}
	DefaultSecondaryRoles *[]*string `field:"optional" json:"defaultSecondaryRoles" yaml:"defaultSecondaryRoles"`
	// Specifies the virtual warehouse that is active by default for the user’s session upon login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#default_warehouse User#default_warehouse}
	DefaultWarehouse *string `field:"optional" json:"defaultWarehouse" yaml:"defaultWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#disabled User#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Name displayed for the user in the Snowflake web interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#display_name User#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Email address for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#email User#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// First name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#first_name User#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#id User#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Last name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#last_name User#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The name users use to log in. If not supplied, snowflake will use name instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#login_name User#login_name}
	LoginName *string `field:"optional" json:"loginName" yaml:"loginName"`
	// Specifies whether the user is forced to change their password on next login (including their first/initial login) into the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#must_change_password User#must_change_password}
	MustChangePassword interface{} `field:"optional" json:"mustChangePassword" yaml:"mustChangePassword"`
	// **WARNING:** this will put the password in the terraform state file. Use carefully.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#password User#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Specifies the user’s RSA public key; used for key-pair authentication. Must be on 1 line without header and trailer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#rsa_public_key User#rsa_public_key}
	RsaPublicKey *string `field:"optional" json:"rsaPublicKey" yaml:"rsaPublicKey"`
	// Specifies the user’s second RSA public key;
	//
	// used to rotate the public and private keys for key-pair authentication based on an expiration schedule set by your organization. Must be on 1 line without header and trailer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.77.0/docs/resources/user#rsa_public_key_2 User#rsa_public_key_2}
	RsaPublicKey2 *string `field:"optional" json:"rsaPublicKey2" yaml:"rsaPublicKey2"`
}

