// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforcustomclients

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v11/oauthintegrationforcustomclients/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/oauth_integration_for_custom_clients snowflake_oauth_integration_for_custom_clients}.
type OauthIntegrationForCustomClients interface {
	cdktf.TerraformResource
	BlockedRolesList() *[]*string
	SetBlockedRolesList(val *[]*string)
	BlockedRolesListInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DescribeOutput() OauthIntegrationForCustomClientsDescribeOutputList
	Enabled() *string
	SetEnabled(val *string)
	EnabledInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullyQualifiedName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkPolicy() *string
	SetNetworkPolicy(val *string)
	NetworkPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	OauthAllowNonTlsRedirectUri() *string
	SetOauthAllowNonTlsRedirectUri(val *string)
	OauthAllowNonTlsRedirectUriInput() *string
	OauthClientRsaPublicKey() *string
	SetOauthClientRsaPublicKey(val *string)
	OauthClientRsaPublicKey2() *string
	SetOauthClientRsaPublicKey2(val *string)
	OauthClientRsaPublicKey2Input() *string
	OauthClientRsaPublicKeyInput() *string
	OauthClientType() *string
	SetOauthClientType(val *string)
	OauthClientTypeInput() *string
	OauthEnforcePkce() *string
	SetOauthEnforcePkce(val *string)
	OauthEnforcePkceInput() *string
	OauthIssueRefreshTokens() *string
	SetOauthIssueRefreshTokens(val *string)
	OauthIssueRefreshTokensInput() *string
	OauthRedirectUri() *string
	SetOauthRedirectUri(val *string)
	OauthRedirectUriInput() *string
	OauthRefreshTokenValidity() *float64
	SetOauthRefreshTokenValidity(val *float64)
	OauthRefreshTokenValidityInput() *float64
	OauthUseSecondaryRoles() *string
	SetOauthUseSecondaryRoles(val *string)
	OauthUseSecondaryRolesInput() *string
	PreAuthorizedRolesList() *[]*string
	SetPreAuthorizedRolesList(val *[]*string)
	PreAuthorizedRolesListInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ShowOutput() OauthIntegrationForCustomClientsShowOutputList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetComment()
	ResetEnabled()
	ResetId()
	ResetNetworkPolicy()
	ResetOauthAllowNonTlsRedirectUri()
	ResetOauthClientRsaPublicKey()
	ResetOauthClientRsaPublicKey2()
	ResetOauthEnforcePkce()
	ResetOauthIssueRefreshTokens()
	ResetOauthRefreshTokenValidity()
	ResetOauthUseSecondaryRoles()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreAuthorizedRolesList()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OauthIntegrationForCustomClients
type jsiiProxy_OauthIntegrationForCustomClients struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) BlockedRolesList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) BlockedRolesListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRolesListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) DescribeOutput() OauthIntegrationForCustomClientsDescribeOutputList {
	var returns OauthIntegrationForCustomClientsDescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Enabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) EnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) NetworkPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) NetworkPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthAllowNonTlsRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAllowNonTlsRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthAllowNonTlsRedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAllowNonTlsRedirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthClientRsaPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthClientRsaPublicKey2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthClientRsaPublicKey2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKey2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthClientRsaPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientRsaPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthClientType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthClientTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthEnforcePkce() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthEnforcePkce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthEnforcePkceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthEnforcePkceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthIssueRefreshTokens() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthIssueRefreshTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthIssueRefreshTokensInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthIssueRefreshTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthRedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthRedirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthRefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthRefreshTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthUseSecondaryRoles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthUseSecondaryRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) OauthUseSecondaryRolesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthUseSecondaryRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) PreAuthorizedRolesList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preAuthorizedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) PreAuthorizedRolesListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preAuthorizedRolesListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) ShowOutput() OauthIntegrationForCustomClientsShowOutputList {
	var returns OauthIntegrationForCustomClientsShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClients) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/oauth_integration_for_custom_clients snowflake_oauth_integration_for_custom_clients} Resource.
func NewOauthIntegrationForCustomClients(scope constructs.Construct, id *string, config *OauthIntegrationForCustomClientsConfig) OauthIntegrationForCustomClients {
	_init_.Initialize()

	if err := validateNewOauthIntegrationForCustomClientsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthIntegrationForCustomClients{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClients",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.95.0/docs/resources/oauth_integration_for_custom_clients snowflake_oauth_integration_for_custom_clients} Resource.
func NewOauthIntegrationForCustomClients_Override(o OauthIntegrationForCustomClients, scope constructs.Construct, id *string, config *OauthIntegrationForCustomClientsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClients",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetBlockedRolesList(val *[]*string) {
	if err := j.validateSetBlockedRolesListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedRolesList",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetEnabled(val *string) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetNetworkPolicy(val *string) {
	if err := j.validateSetNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthAllowNonTlsRedirectUri(val *string) {
	if err := j.validateSetOauthAllowNonTlsRedirectUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAllowNonTlsRedirectUri",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthClientRsaPublicKey(val *string) {
	if err := j.validateSetOauthClientRsaPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientRsaPublicKey",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthClientRsaPublicKey2(val *string) {
	if err := j.validateSetOauthClientRsaPublicKey2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientRsaPublicKey2",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthClientType(val *string) {
	if err := j.validateSetOauthClientTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientType",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthEnforcePkce(val *string) {
	if err := j.validateSetOauthEnforcePkceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthEnforcePkce",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthIssueRefreshTokens(val *string) {
	if err := j.validateSetOauthIssueRefreshTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthIssueRefreshTokens",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthRedirectUri(val *string) {
	if err := j.validateSetOauthRedirectUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthRedirectUri",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthRefreshTokenValidity(val *float64) {
	if err := j.validateSetOauthRefreshTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthRefreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetOauthUseSecondaryRoles(val *string) {
	if err := j.validateSetOauthUseSecondaryRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthUseSecondaryRoles",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetPreAuthorizedRolesList(val *[]*string) {
	if err := j.validateSetPreAuthorizedRolesListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preAuthorizedRolesList",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClients)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a OauthIntegrationForCustomClients resource upon running "cdktf plan <stack-name>".
func OauthIntegrationForCustomClients_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOauthIntegrationForCustomClients_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClients",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func OauthIntegrationForCustomClients_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthIntegrationForCustomClients_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClients",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OauthIntegrationForCustomClients_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthIntegrationForCustomClients_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClients",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OauthIntegrationForCustomClients_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthIntegrationForCustomClients_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClients",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OauthIntegrationForCustomClients_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClients",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetComment() {
	_jsii_.InvokeVoid(
		o,
		"resetComment",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetOauthAllowNonTlsRedirectUri() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthAllowNonTlsRedirectUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetOauthClientRsaPublicKey() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthClientRsaPublicKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetOauthClientRsaPublicKey2() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthClientRsaPublicKey2",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetOauthEnforcePkce() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthEnforcePkce",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetOauthIssueRefreshTokens() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthIssueRefreshTokens",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetOauthRefreshTokenValidity() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthRefreshTokenValidity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetOauthUseSecondaryRoles() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthUseSecondaryRoles",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ResetPreAuthorizedRolesList() {
	_jsii_.InvokeVoid(
		o,
		"resetPreAuthorizedRolesList",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClients) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

