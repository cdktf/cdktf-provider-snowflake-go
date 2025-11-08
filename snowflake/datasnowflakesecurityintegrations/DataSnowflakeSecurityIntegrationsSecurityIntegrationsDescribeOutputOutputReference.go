// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasnowflakesecurityintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/datasnowflakesecurityintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference interface {
	cdktf.ComplexObject
	AllowedEmailPatterns() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAllowedEmailPatternsList
	AllowedUserDomains() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAllowedUserDomainsList
	AuthType() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAuthTypeList
	BlockedRolesList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputBlockedRolesListStructList
	Comment() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputCommentList
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputEnabledList
	ExternalOauthAllowedRolesList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAllowedRolesListStructList
	ExternalOauthAnyRoleMode() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAnyRoleModeList
	ExternalOauthAudienceList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAudienceListStructList
	ExternalOauthBlockedRolesList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthBlockedRolesListStructList
	ExternalOauthIssuer() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthIssuerList
	ExternalOauthJwsKeysUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthJwsKeysUrlList
	ExternalOauthRsaPublicKey() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKeyList
	ExternalOauthRsaPublicKey2() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List
	ExternalOauthScopeDelimiter() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthScopeDelimiterList
	ExternalOauthSnowflakeUserMappingAttribute() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthSnowflakeUserMappingAttributeList
	ExternalOauthTokenUserMappingClaim() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthTokenUserMappingClaimList
	// Experimental.
	Fqn() *string
	InternalValue() *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutput
	SetInternalValue(val *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutput)
	NetworkPolicy() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputNetworkPolicyList
	OauthAccessTokenValidity() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAccessTokenValidityList
	OauthAllowedAuthorizationEndpoints() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedAuthorizationEndpointsList
	OauthAllowedScopes() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedScopesList
	OauthAllowedTokenEndpoints() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedTokenEndpointsList
	OauthAllowNonTlsRedirectUri() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowNonTlsRedirectUriList
	OauthAuthorizationEndpoint() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAuthorizationEndpointList
	OauthClientAuthMethod() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientAuthMethodList
	OauthClientRsaPublicKey2Fp() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientRsaPublicKey2FpList
	OauthClientRsaPublicKeyFp() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientRsaPublicKeyFpList
	OauthClientType() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientTypeList
	OauthEnforcePkce() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthEnforcePkceList
	OauthGrant() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthGrantList
	OauthIssueRefreshTokens() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthIssueRefreshTokensList
	OauthRefreshTokenValidity() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthRefreshTokenValidityList
	OauthTokenEndpoint() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthTokenEndpointList
	OauthUseSecondaryRoles() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthUseSecondaryRolesList
	ParentIntegration() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputParentIntegrationList
	PreAuthorizedRolesList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputPreAuthorizedRolesListStructList
	RunAsRole() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputRunAsRoleList
	Saml2DigestMethodsUsed() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2DigestMethodsUsedList
	Saml2EnableSpInitiated() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2EnableSpInitiatedList
	Saml2ForceAuthn() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2ForceAuthnList
	Saml2Issuer() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2IssuerList
	Saml2PostLogoutRedirectUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2PostLogoutRedirectUrlList
	Saml2Provider() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2ProviderList
	Saml2RequestedNameidFormat() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2RequestedNameidFormatList
	Saml2SignatureMethodsUsed() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SignatureMethodsUsedList
	Saml2SignRequest() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SignRequestList
	Saml2SnowflakeAcsUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeAcsUrlList
	Saml2SnowflakeIssuerUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeIssuerUrlList
	Saml2SnowflakeMetadata() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeMetadataList
	Saml2SpInitiatedLoginPageLabel() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelList
	Saml2SsoUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SsoUrlList
	SyncPassword() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSyncPasswordList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference
type jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) AllowedEmailPatterns() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAllowedEmailPatternsList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAllowedEmailPatternsList
	_jsii_.Get(
		j,
		"allowedEmailPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) AllowedUserDomains() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAllowedUserDomainsList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAllowedUserDomainsList
	_jsii_.Get(
		j,
		"allowedUserDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) AuthType() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAuthTypeList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputAuthTypeList
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) BlockedRolesList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputBlockedRolesListStructList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputBlockedRolesListStructList
	_jsii_.Get(
		j,
		"blockedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Comment() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputCommentList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputCommentList
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Enabled() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputEnabledList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputEnabledList
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthAllowedRolesList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAllowedRolesListStructList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAllowedRolesListStructList
	_jsii_.Get(
		j,
		"externalOauthAllowedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthAnyRoleMode() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAnyRoleModeList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAnyRoleModeList
	_jsii_.Get(
		j,
		"externalOauthAnyRoleMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthAudienceList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAudienceListStructList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthAudienceListStructList
	_jsii_.Get(
		j,
		"externalOauthAudienceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthBlockedRolesList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthBlockedRolesListStructList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthBlockedRolesListStructList
	_jsii_.Get(
		j,
		"externalOauthBlockedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthIssuer() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthIssuerList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthIssuerList
	_jsii_.Get(
		j,
		"externalOauthIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthJwsKeysUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthJwsKeysUrlList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthJwsKeysUrlList
	_jsii_.Get(
		j,
		"externalOauthJwsKeysUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthRsaPublicKey() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKeyList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKeyList
	_jsii_.Get(
		j,
		"externalOauthRsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthRsaPublicKey2() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthRsaPublicKey2List
	_jsii_.Get(
		j,
		"externalOauthRsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthScopeDelimiter() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthScopeDelimiterList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthScopeDelimiterList
	_jsii_.Get(
		j,
		"externalOauthScopeDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthSnowflakeUserMappingAttribute() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthSnowflakeUserMappingAttributeList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthSnowflakeUserMappingAttributeList
	_jsii_.Get(
		j,
		"externalOauthSnowflakeUserMappingAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ExternalOauthTokenUserMappingClaim() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthTokenUserMappingClaimList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputExternalOauthTokenUserMappingClaimList
	_jsii_.Get(
		j,
		"externalOauthTokenUserMappingClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) InternalValue() *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutput {
	var returns *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) NetworkPolicy() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputNetworkPolicyList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthAccessTokenValidity() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAccessTokenValidityList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAccessTokenValidityList
	_jsii_.Get(
		j,
		"oauthAccessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthAllowedAuthorizationEndpoints() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedAuthorizationEndpointsList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedAuthorizationEndpointsList
	_jsii_.Get(
		j,
		"oauthAllowedAuthorizationEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthAllowedScopes() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedScopesList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedScopesList
	_jsii_.Get(
		j,
		"oauthAllowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthAllowedTokenEndpoints() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedTokenEndpointsList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowedTokenEndpointsList
	_jsii_.Get(
		j,
		"oauthAllowedTokenEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthAllowNonTlsRedirectUri() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowNonTlsRedirectUriList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAllowNonTlsRedirectUriList
	_jsii_.Get(
		j,
		"oauthAllowNonTlsRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthAuthorizationEndpoint() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAuthorizationEndpointList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthAuthorizationEndpointList
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthClientAuthMethod() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientAuthMethodList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientAuthMethodList
	_jsii_.Get(
		j,
		"oauthClientAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthClientRsaPublicKey2Fp() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientRsaPublicKey2FpList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientRsaPublicKey2FpList
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKey2Fp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthClientRsaPublicKeyFp() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientRsaPublicKeyFpList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientRsaPublicKeyFpList
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKeyFp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthClientType() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientTypeList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthClientTypeList
	_jsii_.Get(
		j,
		"oauthClientType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthEnforcePkce() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthEnforcePkceList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthEnforcePkceList
	_jsii_.Get(
		j,
		"oauthEnforcePkce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthGrant() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthGrantList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthGrantList
	_jsii_.Get(
		j,
		"oauthGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthIssueRefreshTokens() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthIssueRefreshTokensList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthIssueRefreshTokensList
	_jsii_.Get(
		j,
		"oauthIssueRefreshTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthRefreshTokenValidity() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthRefreshTokenValidityList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthRefreshTokenValidityList
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthTokenEndpoint() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthTokenEndpointList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthTokenEndpointList
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) OauthUseSecondaryRoles() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthUseSecondaryRolesList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOauthUseSecondaryRolesList
	_jsii_.Get(
		j,
		"oauthUseSecondaryRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ParentIntegration() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputParentIntegrationList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputParentIntegrationList
	_jsii_.Get(
		j,
		"parentIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) PreAuthorizedRolesList() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputPreAuthorizedRolesListStructList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputPreAuthorizedRolesListStructList
	_jsii_.Get(
		j,
		"preAuthorizedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) RunAsRole() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputRunAsRoleList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputRunAsRoleList
	_jsii_.Get(
		j,
		"runAsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2DigestMethodsUsed() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2DigestMethodsUsedList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2DigestMethodsUsedList
	_jsii_.Get(
		j,
		"saml2DigestMethodsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2EnableSpInitiated() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2EnableSpInitiatedList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2EnableSpInitiatedList
	_jsii_.Get(
		j,
		"saml2EnableSpInitiated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2ForceAuthn() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2ForceAuthnList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2ForceAuthnList
	_jsii_.Get(
		j,
		"saml2ForceAuthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2Issuer() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2IssuerList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2IssuerList
	_jsii_.Get(
		j,
		"saml2Issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2PostLogoutRedirectUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2PostLogoutRedirectUrlList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2PostLogoutRedirectUrlList
	_jsii_.Get(
		j,
		"saml2PostLogoutRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2Provider() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2ProviderList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2ProviderList
	_jsii_.Get(
		j,
		"saml2Provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2RequestedNameidFormat() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2RequestedNameidFormatList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2RequestedNameidFormatList
	_jsii_.Get(
		j,
		"saml2RequestedNameidFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2SignatureMethodsUsed() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SignatureMethodsUsedList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SignatureMethodsUsedList
	_jsii_.Get(
		j,
		"saml2SignatureMethodsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2SignRequest() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SignRequestList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SignRequestList
	_jsii_.Get(
		j,
		"saml2SignRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2SnowflakeAcsUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeAcsUrlList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeAcsUrlList
	_jsii_.Get(
		j,
		"saml2SnowflakeAcsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2SnowflakeIssuerUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeIssuerUrlList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeIssuerUrlList
	_jsii_.Get(
		j,
		"saml2SnowflakeIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2SnowflakeMetadata() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeMetadataList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SnowflakeMetadataList
	_jsii_.Get(
		j,
		"saml2SnowflakeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2SpInitiatedLoginPageLabel() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SpInitiatedLoginPageLabelList
	_jsii_.Get(
		j,
		"saml2SpInitiatedLoginPageLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Saml2SsoUrl() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SsoUrlList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSaml2SsoUrlList
	_jsii_.Get(
		j,
		"saml2SsoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) SyncPassword() DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSyncPasswordList {
	var returns DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputSyncPasswordList
	_jsii_.Get(
		j,
		"syncPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeSecurityIntegrations.DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference_Override(d DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.dataSnowflakeSecurityIntegrations.DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference)SetInternalValue(val *DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeSecurityIntegrationsSecurityIntegrationsDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

