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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#database AuthenticationPolicy#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the authentication policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#name AuthenticationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the authentication policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#schema AuthenticationPolicy#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// A list of authentication methods that are allowed during login.
	//
	// Valid values are (case-insensitive): `ALL` | `SAML` | `PASSWORD` | `OAUTH` | `KEYPAIR` | `PROGRAMMATIC_ACCESS_TOKEN` | `WORKLOAD_IDENTITY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#authentication_methods AuthenticationPolicy#authentication_methods}
	AuthenticationMethods *[]*string `field:"optional" json:"authenticationMethods" yaml:"authenticationMethods"`
	// A list of clients that can authenticate with Snowflake.
	//
	// If a client tries to connect, and the client is not one of the valid `client_types`, then the login attempt fails. Valid values are (case-insensitive): `ALL` | `SNOWFLAKE_UI` | `DRIVERS` | `SNOWSQL` | `SNOWFLAKE_CLI`. The `client_types` property of an authentication policy is a best effort method to block user logins based on specific clients. It should not be used as the sole control to establish a security boundary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#client_types AuthenticationPolicy#client_types}
	ClientTypes *[]*string `field:"optional" json:"clientTypes" yaml:"clientTypes"`
	// Specifies a comment for the authentication policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#comment AuthenticationPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#id AuthenticationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of authentication methods that enforce multi-factor authentication (MFA) during login.
	//
	// Authentication methods not listed in this parameter do not prompt for multi-factor authentication. Allowed values are `ALL` | `SAML` | `PASSWORD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#mfa_authentication_methods AuthenticationPolicy#mfa_authentication_methods}
	MfaAuthenticationMethods *[]*string `field:"optional" json:"mfaAuthenticationMethods" yaml:"mfaAuthenticationMethods"`
	// Determines whether a user must enroll in multi-factor authentication.
	//
	// Valid values are (case-insensitive): `REQUIRED` | `REQUIRED_PASSWORD_ONLY` | `OPTIONAL`. When REQUIRED is specified, Enforces users to enroll in MFA. If this value is used, then the `client_types` parameter must include `snowflake_ui`, because Snowsight is the only place users can enroll in multi-factor authentication (MFA). Note that when you set this value to OPTIONAL, and your account setup forces users to enroll in MFA, then Snowflake may set quietly this value to `REQUIRED_PASSWORD_ONLY`, which may cause permadiff. In this case, you may want to adjust this field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#mfa_enrollment AuthenticationPolicy#mfa_enrollment}
	MfaEnrollment *string `field:"optional" json:"mfaEnrollment" yaml:"mfaEnrollment"`
	// mfa_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#mfa_policy AuthenticationPolicy#mfa_policy}
	MfaPolicy *AuthenticationPolicyMfaPolicy `field:"optional" json:"mfaPolicy" yaml:"mfaPolicy"`
	// pat_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#pat_policy AuthenticationPolicy#pat_policy}
	PatPolicy *AuthenticationPolicyPatPolicy `field:"optional" json:"patPolicy" yaml:"patPolicy"`
	// A list of security integrations the authentication policy is associated with.
	//
	// This parameter has no effect when `saml` or `oauth` are not in the `authentication_methods` list. All values in the `security_integrations` list must be compatible with the values in the `authentication_methods` list. For example, if `security_integrations` contains a SAML security integration, and `authentication_methods` contains OAUTH, then you cannot create the authentication policy. To allow all security integrations use `ALL` as parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#security_integrations AuthenticationPolicy#security_integrations}
	SecurityIntegrations *[]*string `field:"optional" json:"securityIntegrations" yaml:"securityIntegrations"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#timeouts AuthenticationPolicy#timeouts}
	Timeouts *AuthenticationPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// workload_identity_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.10.1/docs/resources/authentication_policy#workload_identity_policy AuthenticationPolicy#workload_identity_policy}
	WorkloadIdentityPolicy *AuthenticationPolicyWorkloadIdentityPolicy `field:"optional" json:"workloadIdentityPolicy" yaml:"workloadIdentityPolicy"`
}

