// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthenticationPolicyConfig struct {
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
	// The database in which to create the authentication policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#database AuthenticationPolicy#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the authentication policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#name AuthenticationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the authentication policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#schema AuthenticationPolicy#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// A list of authentication methods that are allowed during login.
	//
	// This parameter accepts one or more of the following values: `ALL` | `SAML` | `PASSWORD` | `OAUTH` | `KEYPAIR`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#authentication_methods AuthenticationPolicy#authentication_methods}
	AuthenticationMethods *[]*string `field:"optional" json:"authenticationMethods" yaml:"authenticationMethods"`
	// A list of clients that can authenticate with Snowflake.
	//
	// If a client tries to connect, and the client is not one of the valid CLIENT_TYPES, then the login attempt fails. Allowed values are `ALL` | `SNOWFLAKE_UI` | `DRIVERS` | `SNOWSQL`. The CLIENT_TYPES property of an authentication policy is a best effort method to block user logins based on specific clients. It should not be used as the sole control to establish a security boundary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#client_types AuthenticationPolicy#client_types}
	ClientTypes *[]*string `field:"optional" json:"clientTypes" yaml:"clientTypes"`
	// Specifies a comment for the authentication policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#comment AuthenticationPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#id AuthenticationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of authentication methods that enforce multi-factor authentication (MFA) during login.
	//
	// Authentication methods not listed in this parameter do not prompt for multi-factor authentication. Allowed values are `ALL` | `SAML` | `PASSWORD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#mfa_authentication_methods AuthenticationPolicy#mfa_authentication_methods}
	MfaAuthenticationMethods *[]*string `field:"optional" json:"mfaAuthenticationMethods" yaml:"mfaAuthenticationMethods"`
	// (Default: `OPTIONAL`) Determines whether a user must enroll in multi-factor authentication.
	//
	// Allowed values are REQUIRED and OPTIONAL. When REQUIRED is specified, Enforces users to enroll in MFA. If this value is used, then the CLIENT_TYPES parameter must include SNOWFLAKE_UI, because Snowsight is the only place users can enroll in multi-factor authentication (MFA).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#mfa_enrollment AuthenticationPolicy#mfa_enrollment}
	MfaEnrollment *string `field:"optional" json:"mfaEnrollment" yaml:"mfaEnrollment"`
	// A list of security integrations the authentication policy is associated with.
	//
	// This parameter has no effect when SAML or OAUTH are not in the AUTHENTICATION_METHODS list. All values in the SECURITY_INTEGRATIONS list must be compatible with the values in the AUTHENTICATION_METHODS list. For example, if SECURITY_INTEGRATIONS contains a SAML security integration, and AUTHENTICATION_METHODS contains OAUTH, then you cannot create the authentication policy. To allow all security integrations use ALL as parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#security_integrations AuthenticationPolicy#security_integrations}
	SecurityIntegrations *[]*string `field:"optional" json:"securityIntegrations" yaml:"securityIntegrations"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.8.0/docs/resources/authentication_policy#timeouts AuthenticationPolicy#timeouts}
	Timeouts *AuthenticationPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

