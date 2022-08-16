// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-snowflake-go/snowflake/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-snowflake-go/snowflake/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/snowflake/r/saml_integration snowflake_saml_integration}.
type SamlIntegration interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	Saml2DigestMethodsUsed() *string
	Saml2EnableSpInitiated() interface{}
	SetSaml2EnableSpInitiated(val interface{})
	Saml2EnableSpInitiatedInput() interface{}
	Saml2ForceAuthn() interface{}
	SetSaml2ForceAuthn(val interface{})
	Saml2ForceAuthnInput() interface{}
	Saml2Issuer() *string
	SetSaml2Issuer(val *string)
	Saml2IssuerInput() *string
	Saml2PostLogoutRedirectUrl() *string
	SetSaml2PostLogoutRedirectUrl(val *string)
	Saml2PostLogoutRedirectUrlInput() *string
	Saml2Provider() *string
	SetSaml2Provider(val *string)
	Saml2ProviderInput() *string
	Saml2RequestedNameidFormat() *string
	SetSaml2RequestedNameidFormat(val *string)
	Saml2RequestedNameidFormatInput() *string
	Saml2SignatureMethodsUsed() *string
	Saml2SignRequest() interface{}
	SetSaml2SignRequest(val interface{})
	Saml2SignRequestInput() interface{}
	Saml2SnowflakeAcsUrl() *string
	SetSaml2SnowflakeAcsUrl(val *string)
	Saml2SnowflakeAcsUrlInput() *string
	Saml2SnowflakeIssuerUrl() *string
	SetSaml2SnowflakeIssuerUrl(val *string)
	Saml2SnowflakeIssuerUrlInput() *string
	Saml2SnowflakeMetadata() *string
	Saml2SnowflakeX509Cert() *string
	SetSaml2SnowflakeX509Cert(val *string)
	Saml2SnowflakeX509CertInput() *string
	Saml2SpInitiatedLoginPageLabel() *string
	SetSaml2SpInitiatedLoginPageLabel(val *string)
	Saml2SpInitiatedLoginPageLabelInput() *string
	Saml2SsoUrl() *string
	SetSaml2SsoUrl(val *string)
	Saml2SsoUrlInput() *string
	Saml2X509Cert() *string
	SetSaml2X509Cert(val *string)
	Saml2X509CertInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSaml2EnableSpInitiated()
	ResetSaml2ForceAuthn()
	ResetSaml2PostLogoutRedirectUrl()
	ResetSaml2RequestedNameidFormat()
	ResetSaml2SignRequest()
	ResetSaml2SnowflakeAcsUrl()
	ResetSaml2SnowflakeIssuerUrl()
	ResetSaml2SnowflakeX509Cert()
	ResetSaml2SpInitiatedLoginPageLabel()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SamlIntegration
type jsiiProxy_SamlIntegration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SamlIntegration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2DigestMethodsUsed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2DigestMethodsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2EnableSpInitiated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2EnableSpInitiated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2EnableSpInitiatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2EnableSpInitiatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2ForceAuthn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2ForceAuthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2ForceAuthnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2ForceAuthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2Issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2IssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2PostLogoutRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2PostLogoutRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2PostLogoutRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2PostLogoutRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2Provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2ProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2RequestedNameidFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2RequestedNameidFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2RequestedNameidFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2RequestedNameidFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SignatureMethodsUsed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SignatureMethodsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SignRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2SignRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SignRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2SignRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SnowflakeAcsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeAcsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SnowflakeAcsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeAcsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SnowflakeIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SnowflakeIssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeIssuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SnowflakeMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SnowflakeX509Cert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeX509Cert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SnowflakeX509CertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeX509CertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SpInitiatedLoginPageLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SpInitiatedLoginPageLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SpInitiatedLoginPageLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SpInitiatedLoginPageLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SsoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SsoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2X509Cert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2X509Cert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) Saml2X509CertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2X509CertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/saml_integration snowflake_saml_integration} Resource.
func NewSamlIntegration(scope constructs.Construct, id *string, config *SamlIntegrationConfig) SamlIntegration {
	_init_.Initialize()

	j := jsiiProxy_SamlIntegration{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.SamlIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/saml_integration snowflake_saml_integration} Resource.
func NewSamlIntegration_Override(s SamlIntegration, scope constructs.Construct, id *string, config *SamlIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.SamlIntegration",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SamlIntegration) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2EnableSpInitiated(val interface{}) {
	_jsii_.Set(
		j,
		"saml2EnableSpInitiated",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2ForceAuthn(val interface{}) {
	_jsii_.Set(
		j,
		"saml2ForceAuthn",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2Issuer(val *string) {
	_jsii_.Set(
		j,
		"saml2Issuer",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2PostLogoutRedirectUrl(val *string) {
	_jsii_.Set(
		j,
		"saml2PostLogoutRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2Provider(val *string) {
	_jsii_.Set(
		j,
		"saml2Provider",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2RequestedNameidFormat(val *string) {
	_jsii_.Set(
		j,
		"saml2RequestedNameidFormat",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2SignRequest(val interface{}) {
	_jsii_.Set(
		j,
		"saml2SignRequest",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2SnowflakeAcsUrl(val *string) {
	_jsii_.Set(
		j,
		"saml2SnowflakeAcsUrl",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2SnowflakeIssuerUrl(val *string) {
	_jsii_.Set(
		j,
		"saml2SnowflakeIssuerUrl",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2SnowflakeX509Cert(val *string) {
	_jsii_.Set(
		j,
		"saml2SnowflakeX509Cert",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2SpInitiatedLoginPageLabel(val *string) {
	_jsii_.Set(
		j,
		"saml2SpInitiatedLoginPageLabel",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2SsoUrl(val *string) {
	_jsii_.Set(
		j,
		"saml2SsoUrl",
		val,
	)
}

func (j *jsiiProxy_SamlIntegration) SetSaml2X509Cert(val *string) {
	_jsii_.Set(
		j,
		"saml2X509Cert",
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
func SamlIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.SamlIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SamlIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.SamlIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SamlIntegration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SamlIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SamlIntegration) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2EnableSpInitiated() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2EnableSpInitiated",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2ForceAuthn() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2ForceAuthn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2PostLogoutRedirectUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2PostLogoutRedirectUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2RequestedNameidFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2RequestedNameidFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2SignRequest() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SignRequest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2SnowflakeAcsUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SnowflakeAcsUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2SnowflakeIssuerUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SnowflakeIssuerUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2SnowflakeX509Cert() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SnowflakeX509Cert",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) ResetSaml2SpInitiatedLoginPageLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SpInitiatedLoginPageLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
