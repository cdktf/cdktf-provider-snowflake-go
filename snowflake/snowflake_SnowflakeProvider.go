// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-snowflake-go/snowflake/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-snowflake-go/snowflake/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/snowflake snowflake}.
type SnowflakeProvider interface {
	cdktf.TerraformProvider
	Account() *string
	SetAccount(val *string)
	AccountInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	BrowserAuth() interface{}
	SetBrowserAuth(val interface{})
	BrowserAuthInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	OauthAccessToken() *string
	SetOauthAccessToken(val *string)
	OauthAccessTokenInput() *string
	OauthClientId() *string
	SetOauthClientId(val *string)
	OauthClientIdInput() *string
	OauthClientSecret() *string
	SetOauthClientSecret(val *string)
	OauthClientSecretInput() *string
	OauthEndpoint() *string
	SetOauthEndpoint(val *string)
	OauthEndpointInput() *string
	OauthRedirectUrl() *string
	SetOauthRedirectUrl(val *string)
	OauthRedirectUrlInput() *string
	OauthRefreshToken() *string
	SetOauthRefreshToken(val *string)
	OauthRefreshTokenInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	PrivateKeyPassphrase() *string
	SetPrivateKeyPassphrase(val *string)
	PrivateKeyPassphraseInput() *string
	PrivateKeyPath() *string
	SetPrivateKeyPath(val *string)
	PrivateKeyPathInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	Warehouse() *string
	SetWarehouse(val *string)
	WarehouseInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetBrowserAuth()
	ResetHost()
	ResetOauthAccessToken()
	ResetOauthClientId()
	ResetOauthClientSecret()
	ResetOauthEndpoint()
	ResetOauthRedirectUrl()
	ResetOauthRefreshToken()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPort()
	ResetPrivateKey()
	ResetPrivateKeyPassphrase()
	ResetPrivateKeyPath()
	ResetProtocol()
	ResetRegion()
	ResetRole()
	ResetWarehouse()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SnowflakeProvider
type jsiiProxy_SnowflakeProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_SnowflakeProvider) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) AccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) BrowserAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) BrowserAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthRefreshToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthRefreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) OauthRefreshTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthRefreshTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKeyPassphrase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKeyPassphraseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphraseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKeyPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) PrivateKeyPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) Warehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnowflakeProvider) WarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/snowflake snowflake} Resource.
func NewSnowflakeProvider(scope constructs.Construct, id *string, config *SnowflakeProviderConfig) SnowflakeProvider {
	_init_.Initialize()

	if err := validateNewSnowflakeProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnowflakeProvider{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.SnowflakeProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/snowflake snowflake} Resource.
func NewSnowflakeProvider_Override(s SnowflakeProvider, scope constructs.Construct, id *string, config *SnowflakeProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.SnowflakeProvider",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetAccount(val *string) {
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetBrowserAuth(val interface{}) {
	if err := j.validateSetBrowserAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browserAuth",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOauthAccessToken(val *string) {
	_jsii_.Set(
		j,
		"oauthAccessToken",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOauthClientId(val *string) {
	_jsii_.Set(
		j,
		"oauthClientId",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOauthClientSecret(val *string) {
	_jsii_.Set(
		j,
		"oauthClientSecret",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOauthEndpoint(val *string) {
	_jsii_.Set(
		j,
		"oauthEndpoint",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOauthRedirectUrl(val *string) {
	_jsii_.Set(
		j,
		"oauthRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetOauthRefreshToken(val *string) {
	_jsii_.Set(
		j,
		"oauthRefreshToken",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPrivateKeyPassphrase(val *string) {
	_jsii_.Set(
		j,
		"privateKeyPassphrase",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetPrivateKeyPath(val *string) {
	_jsii_.Set(
		j,
		"privateKeyPath",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_SnowflakeProvider)SetWarehouse(val *string) {
	_jsii_.Set(
		j,
		"warehouse",
		val,
	)
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
func SnowflakeProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnowflakeProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.SnowflakeProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnowflakeProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.SnowflakeProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SnowflakeProvider) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SnowflakeProvider) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		s,
		"resetAlias",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetBrowserAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetBrowserAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetHost() {
	_jsii_.InvokeVoid(
		s,
		"resetHost",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOauthAccessToken() {
	_jsii_.InvokeVoid(
		s,
		"resetOauthAccessToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOauthClientId() {
	_jsii_.InvokeVoid(
		s,
		"resetOauthClientId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOauthClientSecret() {
	_jsii_.InvokeVoid(
		s,
		"resetOauthClientSecret",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOauthEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetOauthEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOauthRedirectUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetOauthRedirectUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOauthRefreshToken() {
	_jsii_.InvokeVoid(
		s,
		"resetOauthRefreshToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPort() {
	_jsii_.InvokeVoid(
		s,
		"resetPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPrivateKeyPassphrase() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKeyPassphrase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetPrivateKeyPath() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKeyPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetRole() {
	_jsii_.InvokeVoid(
		s,
		"resetRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) ResetWarehouse() {
	_jsii_.InvokeVoid(
		s,
		"resetWarehouse",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnowflakeProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnowflakeProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnowflakeProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnowflakeProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

