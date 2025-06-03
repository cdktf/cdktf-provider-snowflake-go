// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforpartnerapplications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v14/oauthintegrationforpartnerapplications/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/oauth_integration_for_partner_applications snowflake_oauth_integration_for_partner_applications}.
type OauthIntegrationForPartnerApplications interface {
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
	DescribeOutput() OauthIntegrationForPartnerApplicationsDescribeOutputList
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
	// The tree node.
	Node() constructs.Node
	OauthClient() *string
	SetOauthClient(val *string)
	OauthClientInput() *string
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
	RelatedParameters() OauthIntegrationForPartnerApplicationsRelatedParametersList
	ShowOutput() OauthIntegrationForPartnerApplicationsShowOutputList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() OauthIntegrationForPartnerApplicationsTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *OauthIntegrationForPartnerApplicationsTimeouts)
	ResetBlockedRolesList()
	ResetComment()
	ResetEnabled()
	ResetId()
	ResetOauthIssueRefreshTokens()
	ResetOauthRedirectUri()
	ResetOauthRefreshTokenValidity()
	ResetOauthUseSecondaryRoles()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
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

// The jsii proxy struct for OauthIntegrationForPartnerApplications
type jsiiProxy_OauthIntegrationForPartnerApplications struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) BlockedRolesList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRolesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) BlockedRolesListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRolesListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) DescribeOutput() OauthIntegrationForPartnerApplicationsDescribeOutputList {
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Enabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) EnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthClient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthIssueRefreshTokens() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthIssueRefreshTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthIssueRefreshTokensInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthIssueRefreshTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthRedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthRedirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthRefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthRefreshTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthUseSecondaryRoles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthUseSecondaryRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) OauthUseSecondaryRolesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthUseSecondaryRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) RelatedParameters() OauthIntegrationForPartnerApplicationsRelatedParametersList {
	var returns OauthIntegrationForPartnerApplicationsRelatedParametersList
	_jsii_.Get(
		j,
		"relatedParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) ShowOutput() OauthIntegrationForPartnerApplicationsShowOutputList {
	var returns OauthIntegrationForPartnerApplicationsShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) Timeouts() OauthIntegrationForPartnerApplicationsTimeoutsOutputReference {
	var returns OauthIntegrationForPartnerApplicationsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/oauth_integration_for_partner_applications snowflake_oauth_integration_for_partner_applications} Resource.
func NewOauthIntegrationForPartnerApplications(scope constructs.Construct, id *string, config *OauthIntegrationForPartnerApplicationsConfig) OauthIntegrationForPartnerApplications {
	_init_.Initialize()

	if err := validateNewOauthIntegrationForPartnerApplicationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthIntegrationForPartnerApplications{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplications",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.1.0/docs/resources/oauth_integration_for_partner_applications snowflake_oauth_integration_for_partner_applications} Resource.
func NewOauthIntegrationForPartnerApplications_Override(o OauthIntegrationForPartnerApplications, scope constructs.Construct, id *string, config *OauthIntegrationForPartnerApplicationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplications",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetBlockedRolesList(val *[]*string) {
	if err := j.validateSetBlockedRolesListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedRolesList",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetEnabled(val *string) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetOauthClient(val *string) {
	if err := j.validateSetOauthClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClient",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetOauthIssueRefreshTokens(val *string) {
	if err := j.validateSetOauthIssueRefreshTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthIssueRefreshTokens",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetOauthRedirectUri(val *string) {
	if err := j.validateSetOauthRedirectUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthRedirectUri",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetOauthRefreshTokenValidity(val *float64) {
	if err := j.validateSetOauthRefreshTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthRefreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetOauthUseSecondaryRoles(val *string) {
	if err := j.validateSetOauthUseSecondaryRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthUseSecondaryRoles",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplications)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a OauthIntegrationForPartnerApplications resource upon running "cdktf plan <stack-name>".
func OauthIntegrationForPartnerApplications_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOauthIntegrationForPartnerApplications_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplications",
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
func OauthIntegrationForPartnerApplications_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthIntegrationForPartnerApplications_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplications",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OauthIntegrationForPartnerApplications_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthIntegrationForPartnerApplications_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplications",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OauthIntegrationForPartnerApplications_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthIntegrationForPartnerApplications_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplications",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OauthIntegrationForPartnerApplications_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplications",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) PutTimeouts(value *OauthIntegrationForPartnerApplicationsTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetBlockedRolesList() {
	_jsii_.InvokeVoid(
		o,
		"resetBlockedRolesList",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetComment() {
	_jsii_.InvokeVoid(
		o,
		"resetComment",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetOauthIssueRefreshTokens() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthIssueRefreshTokens",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetOauthRedirectUri() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthRedirectUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetOauthRefreshTokenValidity() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthRefreshTokenValidity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetOauthUseSecondaryRoles() {
	_jsii_.InvokeVoid(
		o,
		"resetOauthUseSecondaryRoles",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplications) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

