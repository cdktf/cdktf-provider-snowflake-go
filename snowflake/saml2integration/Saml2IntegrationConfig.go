// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package saml2integration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Saml2IntegrationConfig struct {
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
	// Specifies the name of the SAML2 integration.
	//
	// This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#name Saml2Integration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The string containing the IdP EntityID / Issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_issuer Saml2Integration#saml2_issuer}
	Saml2Issuer *string `field:"required" json:"saml2Issuer" yaml:"saml2Issuer"`
	// The string describing the IdP. Valid options are: `OKTA` | `ADFS` | `CUSTOM`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_provider Saml2Integration#saml2_provider}
	Saml2Provider *string `field:"required" json:"saml2Provider" yaml:"saml2Provider"`
	// The string containing the IdP SSO URL, where the user should be redirected by Snowflake (the Service Provider) with a SAML AuthnRequest message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_sso_url Saml2Integration#saml2_sso_url}
	Saml2SsoUrl *string `field:"required" json:"saml2SsoUrl" yaml:"saml2SsoUrl"`
	// The Base64 encoded IdP signing certificate on a single line without the leading -----BEGIN CERTIFICATE----- and ending -----END CERTIFICATE----- markers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_x509_cert Saml2Integration#saml2_x509_cert}
	Saml2X509Cert *string `field:"required" json:"saml2X509Cert" yaml:"saml2X509Cert"`
	// A list of regular expressions that email addresses are matched against to authenticate with a SAML2 security integration.
	//
	// If this field changes value from non-empty to empty, the whole resource is recreated because of Snowflake limitations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#allowed_email_patterns Saml2Integration#allowed_email_patterns}
	AllowedEmailPatterns *[]*string `field:"optional" json:"allowedEmailPatterns" yaml:"allowedEmailPatterns"`
	// A list of email domains that can authenticate with a SAML2 security integration.
	//
	// If this field changes value from non-empty to empty, the whole resource is recreated because of Snowflake limitations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#allowed_user_domains Saml2Integration#allowed_user_domains}
	AllowedUserDomains *[]*string `field:"optional" json:"allowedUserDomains" yaml:"allowedUserDomains"`
	// Specifies a comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#comment Saml2Integration#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether this security integration is enabled or disabled.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#enabled Saml2Integration#enabled}
	Enabled *string `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#id Saml2Integration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The Boolean indicating if the Log In With button will be shown on the login page.
	//
	// TRUE: displays the Log in With button on the login page. FALSE: does not display the Log in With button on the login page. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_enable_sp_initiated Saml2Integration#saml2_enable_sp_initiated}
	Saml2EnableSpInitiated *string `field:"optional" json:"saml2EnableSpInitiated" yaml:"saml2EnableSpInitiated"`
	// The Boolean indicating whether users, during the initial authentication flow, are forced to authenticate again to access Snowflake.
	//
	// When set to TRUE, Snowflake sets the ForceAuthn SAML parameter to TRUE in the outgoing request from Snowflake to the identity provider. TRUE: forces users to authenticate again to access Snowflake, even if a valid session with the identity provider exists. FALSE: does not force users to authenticate again to access Snowflake. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_force_authn Saml2Integration#saml2_force_authn}
	Saml2ForceAuthn *string `field:"optional" json:"saml2ForceAuthn" yaml:"saml2ForceAuthn"`
	// The endpoint to which Snowflake redirects users after clicking the Log Out button in the classic Snowflake web interface.
	//
	// Snowflake terminates the Snowflake session upon redirecting to the specified endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_post_logout_redirect_url Saml2Integration#saml2_post_logout_redirect_url}
	Saml2PostLogoutRedirectUrl *string `field:"optional" json:"saml2PostLogoutRedirectUrl" yaml:"saml2PostLogoutRedirectUrl"`
	// The SAML NameID format allows Snowflake to set an expectation of the identifying attribute of the user (i.e. SAML Subject) in the SAML assertion from the IdP to ensure a valid authentication to Snowflake. Valid options are: `urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified` | `urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress` | `urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName` | `urn:oasis:names:tc:SAML:1.1:nameid-format:WindowsDomainQualifiedName` | `urn:oasis:names:tc:SAML:2.0:nameid-format:kerberos` | `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent` | `urn:oasis:names:tc:SAML:2.0:nameid-format:transient`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_requested_nameid_format Saml2Integration#saml2_requested_nameid_format}
	Saml2RequestedNameidFormat *string `field:"optional" json:"saml2RequestedNameidFormat" yaml:"saml2RequestedNameidFormat"`
	// The Boolean indicating whether SAML requests are signed.
	//
	// TRUE: allows SAML requests to be signed. FALSE: does not allow SAML requests to be signed. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_sign_request Saml2Integration#saml2_sign_request}
	Saml2SignRequest *string `field:"optional" json:"saml2SignRequest" yaml:"saml2SignRequest"`
	// The string containing the Snowflake Assertion Consumer Service URL to which the IdP will send its SAML authentication response back to Snowflake.
	//
	// This property will be set in the SAML authentication request generated by Snowflake when initiating a SAML SSO operation with the IdP. If an incorrect value is specified, Snowflake returns an error message indicating the acceptable values to use. Because Okta does not support underscores in URLs, the underscore in the account name must be converted to a hyphen. See [docs](https://docs.snowflake.com/en/user-guide/organizations-connect#okta-urls).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_snowflake_acs_url Saml2Integration#saml2_snowflake_acs_url}
	Saml2SnowflakeAcsUrl *string `field:"optional" json:"saml2SnowflakeAcsUrl" yaml:"saml2SnowflakeAcsUrl"`
	// The string containing the EntityID / Issuer for the Snowflake service provider.
	//
	// If an incorrect value is specified, Snowflake returns an error message indicating the acceptable values to use. Because Okta does not support underscores in URLs, the underscore in the account name must be converted to a hyphen. See [docs](https://docs.snowflake.com/en/user-guide/organizations-connect#okta-urls).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_snowflake_issuer_url Saml2Integration#saml2_snowflake_issuer_url}
	Saml2SnowflakeIssuerUrl *string `field:"optional" json:"saml2SnowflakeIssuerUrl" yaml:"saml2SnowflakeIssuerUrl"`
	// The string containing the label to display after the Log In With button on the login page.
	//
	// If this field changes value from non-empty to empty, the whole resource is recreated because of Snowflake limitations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.94.0/docs/resources/saml2_integration#saml2_sp_initiated_login_page_label Saml2Integration#saml2_sp_initiated_login_page_label}
	Saml2SpInitiatedLoginPageLabel *string `field:"optional" json:"saml2SpInitiatedLoginPageLabel" yaml:"saml2SpInitiatedLoginPageLabel"`
}

