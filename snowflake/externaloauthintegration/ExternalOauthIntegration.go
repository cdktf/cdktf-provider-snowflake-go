// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externaloauthintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v10/externaloauthintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/resources/external_oauth_integration snowflake_external_oauth_integration}.
type ExternalOauthIntegration interface {
	cdktf.TerraformResource
	AllowedRoles() *[]*string
	SetAllowedRoles(val *[]*string)
	AllowedRolesInput() *[]*string
	AnyRoleMode() *string
	SetAnyRoleMode(val *string)
	AnyRoleModeInput() *string
	AudienceUrls() *[]*string
	SetAudienceUrls(val *[]*string)
	AudienceUrlsInput() *[]*string
	BlockedRoles() *[]*string
	SetBlockedRoles(val *[]*string)
	BlockedRolesInput() *[]*string
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
	CreatedOn() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	JwsKeysUrls() *[]*string
	SetJwsKeysUrls(val *[]*string)
	JwsKeysUrlsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	RsaPublicKey() *string
	SetRsaPublicKey(val *string)
	RsaPublicKey2() *string
	SetRsaPublicKey2(val *string)
	RsaPublicKey2Input() *string
	RsaPublicKeyInput() *string
	ScopeDelimiter() *string
	SetScopeDelimiter(val *string)
	ScopeDelimiterInput() *string
	ScopeMappingAttribute() *string
	SetScopeMappingAttribute(val *string)
	ScopeMappingAttributeInput() *string
	SnowflakeUserMappingAttribute() *string
	SetSnowflakeUserMappingAttribute(val *string)
	SnowflakeUserMappingAttributeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenUserMappingClaims() *[]*string
	SetTokenUserMappingClaims(val *[]*string)
	TokenUserMappingClaimsInput() *[]*string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAllowedRoles()
	ResetAnyRoleMode()
	ResetAudienceUrls()
	ResetBlockedRoles()
	ResetComment()
	ResetId()
	ResetJwsKeysUrls()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRsaPublicKey()
	ResetRsaPublicKey2()
	ResetScopeDelimiter()
	ResetScopeMappingAttribute()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ExternalOauthIntegration
type jsiiProxy_ExternalOauthIntegration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ExternalOauthIntegration) AllowedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) AllowedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) AnyRoleMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyRoleMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) AnyRoleModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyRoleModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) AudienceUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audienceUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) AudienceUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audienceUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) BlockedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) BlockedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) JwsKeysUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwsKeysUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) JwsKeysUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwsKeysUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) RsaPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) RsaPublicKey2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) RsaPublicKey2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) RsaPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) ScopeDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) ScopeDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) ScopeMappingAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeMappingAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) ScopeMappingAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeMappingAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) SnowflakeUserMappingAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowflakeUserMappingAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) SnowflakeUserMappingAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowflakeUserMappingAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) TokenUserMappingClaims() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenUserMappingClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) TokenUserMappingClaimsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenUserMappingClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalOauthIntegration) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/resources/external_oauth_integration snowflake_external_oauth_integration} Resource.
func NewExternalOauthIntegration(scope constructs.Construct, id *string, config *ExternalOauthIntegrationConfig) ExternalOauthIntegration {
	_init_.Initialize()

	if err := validateNewExternalOauthIntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalOauthIntegration{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.80.0/docs/resources/external_oauth_integration snowflake_external_oauth_integration} Resource.
func NewExternalOauthIntegration_Override(e ExternalOauthIntegration, scope constructs.Construct, id *string, config *ExternalOauthIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegration",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetAllowedRoles(val *[]*string) {
	if err := j.validateSetAllowedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRoles",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetAnyRoleMode(val *string) {
	if err := j.validateSetAnyRoleModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyRoleMode",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetAudienceUrls(val *[]*string) {
	if err := j.validateSetAudienceUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audienceUrls",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetBlockedRoles(val *[]*string) {
	if err := j.validateSetBlockedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedRoles",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetJwsKeysUrls(val *[]*string) {
	if err := j.validateSetJwsKeysUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwsKeysUrls",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetRsaPublicKey(val *string) {
	if err := j.validateSetRsaPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetRsaPublicKey2(val *string) {
	if err := j.validateSetRsaPublicKey2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey2",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetScopeDelimiter(val *string) {
	if err := j.validateSetScopeDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeDelimiter",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetScopeMappingAttribute(val *string) {
	if err := j.validateSetScopeMappingAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeMappingAttribute",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetSnowflakeUserMappingAttribute(val *string) {
	if err := j.validateSetSnowflakeUserMappingAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowflakeUserMappingAttribute",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetTokenUserMappingClaims(val *[]*string) {
	if err := j.validateSetTokenUserMappingClaimsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUserMappingClaims",
		val,
	)
}

func (j *jsiiProxy_ExternalOauthIntegration)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a ExternalOauthIntegration resource upon running "cdktf plan <stack-name>".
func ExternalOauthIntegration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateExternalOauthIntegration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegration",
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
func ExternalOauthIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalOauthIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalOauthIntegration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalOauthIntegration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalOauthIntegration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalOauthIntegration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExternalOauthIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.externalOauthIntegration.ExternalOauthIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetAllowedRoles() {
	_jsii_.InvokeVoid(
		e,
		"resetAllowedRoles",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetAnyRoleMode() {
	_jsii_.InvokeVoid(
		e,
		"resetAnyRoleMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetAudienceUrls() {
	_jsii_.InvokeVoid(
		e,
		"resetAudienceUrls",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetBlockedRoles() {
	_jsii_.InvokeVoid(
		e,
		"resetBlockedRoles",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetComment() {
	_jsii_.InvokeVoid(
		e,
		"resetComment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetJwsKeysUrls() {
	_jsii_.InvokeVoid(
		e,
		"resetJwsKeysUrls",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetRsaPublicKey() {
	_jsii_.InvokeVoid(
		e,
		"resetRsaPublicKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetRsaPublicKey2() {
	_jsii_.InvokeVoid(
		e,
		"resetRsaPublicKey2",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetScopeDelimiter() {
	_jsii_.InvokeVoid(
		e,
		"resetScopeDelimiter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) ResetScopeMappingAttribute() {
	_jsii_.InvokeVoid(
		e,
		"resetScopeMappingAttribute",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalOauthIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalOauthIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

