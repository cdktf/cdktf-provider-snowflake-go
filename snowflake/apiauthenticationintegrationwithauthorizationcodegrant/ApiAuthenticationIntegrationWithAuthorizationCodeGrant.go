// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiauthenticationintegrationwithauthorizationcodegrant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/apiauthenticationintegrationwithauthorizationcodegrant/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/api_authentication_integration_with_authorization_code_grant snowflake_api_authentication_integration_with_authorization_code_grant}.
type ApiAuthenticationIntegrationWithAuthorizationCodeGrant interface {
	cdktf.TerraformResource
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
	DescribeOutput() ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputList
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
	OauthAccessTokenValidity() *float64
	SetOauthAccessTokenValidity(val *float64)
	OauthAccessTokenValidityInput() *float64
	OauthAllowedScopes() *[]*string
	SetOauthAllowedScopes(val *[]*string)
	OauthAllowedScopesInput() *[]*string
	OauthAuthorizationEndpoint() *string
	SetOauthAuthorizationEndpoint(val *string)
	OauthAuthorizationEndpointInput() *string
	OauthClientAuthMethod() *string
	SetOauthClientAuthMethod(val *string)
	OauthClientAuthMethodInput() *string
	OauthClientId() *string
	SetOauthClientId(val *string)
	OauthClientIdInput() *string
	OauthClientSecret() *string
	SetOauthClientSecret(val *string)
	OauthClientSecretInput() *string
	OauthRefreshTokenValidity() *float64
	SetOauthRefreshTokenValidity(val *float64)
	OauthRefreshTokenValidityInput() *float64
	OauthTokenEndpoint() *string
	SetOauthTokenEndpoint(val *string)
	OauthTokenEndpointInput() *string
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
	ShowOutput() ApiAuthenticationIntegrationWithAuthorizationCodeGrantShowOutputList
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
	ResetId()
	ResetOauthAccessTokenValidity()
	ResetOauthAllowedScopes()
	ResetOauthAuthorizationEndpoint()
	ResetOauthClientAuthMethod()
	ResetOauthRefreshTokenValidity()
	ResetOauthTokenEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for ApiAuthenticationIntegrationWithAuthorizationCodeGrant
type jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) DescribeOutput() ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputList {
	var returns ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthAccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthAccessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthAccessTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthAccessTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthAllowedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthAllowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthAllowedScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthAllowedScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthAuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthAuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthClientAuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthClientAuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientAuthMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthRefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthRefreshTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthTokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OauthTokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ShowOutput() ApiAuthenticationIntegrationWithAuthorizationCodeGrantShowOutputList {
	var returns ApiAuthenticationIntegrationWithAuthorizationCodeGrantShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/api_authentication_integration_with_authorization_code_grant snowflake_api_authentication_integration_with_authorization_code_grant} Resource.
func NewApiAuthenticationIntegrationWithAuthorizationCodeGrant(scope constructs.Construct, id *string, config *ApiAuthenticationIntegrationWithAuthorizationCodeGrantConfig) ApiAuthenticationIntegrationWithAuthorizationCodeGrant {
	_init_.Initialize()

	if err := validateNewApiAuthenticationIntegrationWithAuthorizationCodeGrantParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithAuthorizationCodeGrant.ApiAuthenticationIntegrationWithAuthorizationCodeGrant",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/api_authentication_integration_with_authorization_code_grant snowflake_api_authentication_integration_with_authorization_code_grant} Resource.
func NewApiAuthenticationIntegrationWithAuthorizationCodeGrant_Override(a ApiAuthenticationIntegrationWithAuthorizationCodeGrant, scope constructs.Construct, id *string, config *ApiAuthenticationIntegrationWithAuthorizationCodeGrantConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithAuthorizationCodeGrant.ApiAuthenticationIntegrationWithAuthorizationCodeGrant",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetOauthAccessTokenValidity(val *float64) {
	if err := j.validateSetOauthAccessTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAccessTokenValidity",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetOauthAllowedScopes(val *[]*string) {
	if err := j.validateSetOauthAllowedScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAllowedScopes",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetOauthAuthorizationEndpoint(val *string) {
	if err := j.validateSetOauthAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAuthorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetOauthClientAuthMethod(val *string) {
	if err := j.validateSetOauthClientAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientAuthMethod",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetOauthClientId(val *string) {
	if err := j.validateSetOauthClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientId",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetOauthClientSecret(val *string) {
	if err := j.validateSetOauthClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientSecret",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetOauthRefreshTokenValidity(val *float64) {
	if err := j.validateSetOauthRefreshTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthRefreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetOauthTokenEndpoint(val *string) {
	if err := j.validateSetOauthTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthTokenEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ApiAuthenticationIntegrationWithAuthorizationCodeGrant resource upon running "cdktf plan <stack-name>".
func ApiAuthenticationIntegrationWithAuthorizationCodeGrant_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApiAuthenticationIntegrationWithAuthorizationCodeGrant_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithAuthorizationCodeGrant.ApiAuthenticationIntegrationWithAuthorizationCodeGrant",
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
func ApiAuthenticationIntegrationWithAuthorizationCodeGrant_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiAuthenticationIntegrationWithAuthorizationCodeGrant_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithAuthorizationCodeGrant.ApiAuthenticationIntegrationWithAuthorizationCodeGrant",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiAuthenticationIntegrationWithAuthorizationCodeGrant_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiAuthenticationIntegrationWithAuthorizationCodeGrant_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithAuthorizationCodeGrant.ApiAuthenticationIntegrationWithAuthorizationCodeGrant",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiAuthenticationIntegrationWithAuthorizationCodeGrant_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiAuthenticationIntegrationWithAuthorizationCodeGrant_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithAuthorizationCodeGrant.ApiAuthenticationIntegrationWithAuthorizationCodeGrant",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiAuthenticationIntegrationWithAuthorizationCodeGrant_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.apiAuthenticationIntegrationWithAuthorizationCodeGrant.ApiAuthenticationIntegrationWithAuthorizationCodeGrant",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetComment() {
	_jsii_.InvokeVoid(
		a,
		"resetComment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetOauthAccessTokenValidity() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthAccessTokenValidity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetOauthAllowedScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthAllowedScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetOauthAuthorizationEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthAuthorizationEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetOauthClientAuthMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthClientAuthMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetOauthRefreshTokenValidity() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthRefreshTokenValidity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetOauthTokenEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthTokenEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrant) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

