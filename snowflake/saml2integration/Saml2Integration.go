// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package saml2integration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/saml2integration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/saml2_integration snowflake_saml2_integration}.
type Saml2Integration interface {
	cdktf.TerraformResource
	AllowedEmailPatterns() *[]*string
	SetAllowedEmailPatterns(val *[]*string)
	AllowedEmailPatternsInput() *[]*string
	AllowedUserDomains() *[]*string
	SetAllowedUserDomains(val *[]*string)
	AllowedUserDomainsInput() *[]*string
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
	DescribeOutput() Saml2IntegrationDescribeOutputList
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
	Saml2EnableSpInitiated() *string
	SetSaml2EnableSpInitiated(val *string)
	Saml2EnableSpInitiatedInput() *string
	Saml2ForceAuthn() *string
	SetSaml2ForceAuthn(val *string)
	Saml2ForceAuthnInput() *string
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
	Saml2SignRequest() *string
	SetSaml2SignRequest(val *string)
	Saml2SignRequestInput() *string
	Saml2SnowflakeAcsUrl() *string
	SetSaml2SnowflakeAcsUrl(val *string)
	Saml2SnowflakeAcsUrlInput() *string
	Saml2SnowflakeIssuerUrl() *string
	SetSaml2SnowflakeIssuerUrl(val *string)
	Saml2SnowflakeIssuerUrlInput() *string
	Saml2SpInitiatedLoginPageLabel() *string
	SetSaml2SpInitiatedLoginPageLabel(val *string)
	Saml2SpInitiatedLoginPageLabelInput() *string
	Saml2SsoUrl() *string
	SetSaml2SsoUrl(val *string)
	Saml2SsoUrlInput() *string
	Saml2X509Cert() *string
	SetSaml2X509Cert(val *string)
	Saml2X509CertInput() *string
	ShowOutput() Saml2IntegrationShowOutputList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Saml2IntegrationTimeoutsOutputReference
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
	PutTimeouts(value *Saml2IntegrationTimeouts)
	ResetAllowedEmailPatterns()
	ResetAllowedUserDomains()
	ResetComment()
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
	ResetSaml2SpInitiatedLoginPageLabel()
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

// The jsii proxy struct for Saml2Integration
type jsiiProxy_Saml2Integration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Saml2Integration) AllowedEmailPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEmailPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) AllowedEmailPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEmailPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) AllowedUserDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUserDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) AllowedUserDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUserDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) DescribeOutput() Saml2IntegrationDescribeOutputList {
	var returns Saml2IntegrationDescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Enabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) EnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2EnableSpInitiated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2EnableSpInitiated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2EnableSpInitiatedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2EnableSpInitiatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2ForceAuthn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2ForceAuthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2ForceAuthnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2ForceAuthnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2Issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2IssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2PostLogoutRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2PostLogoutRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2PostLogoutRedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2PostLogoutRedirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2Provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2ProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2RequestedNameidFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2RequestedNameidFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2RequestedNameidFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2RequestedNameidFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SignRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SignRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SignRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SignRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SnowflakeAcsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeAcsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SnowflakeAcsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeAcsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SnowflakeIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SnowflakeIssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SnowflakeIssuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SpInitiatedLoginPageLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SpInitiatedLoginPageLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SpInitiatedLoginPageLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SpInitiatedLoginPageLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SsoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2SsoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2X509Cert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2X509Cert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Saml2X509CertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saml2X509CertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) ShowOutput() Saml2IntegrationShowOutputList {
	var returns Saml2IntegrationShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) Timeouts() Saml2IntegrationTimeoutsOutputReference {
	var returns Saml2IntegrationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Saml2Integration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/saml2_integration snowflake_saml2_integration} Resource.
func NewSaml2Integration(scope constructs.Construct, id *string, config *Saml2IntegrationConfig) Saml2Integration {
	_init_.Initialize()

	if err := validateNewSaml2IntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Saml2Integration{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.saml2Integration.Saml2Integration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.6.0/docs/resources/saml2_integration snowflake_saml2_integration} Resource.
func NewSaml2Integration_Override(s Saml2Integration, scope constructs.Construct, id *string, config *Saml2IntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.saml2Integration.Saml2Integration",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Saml2Integration)SetAllowedEmailPatterns(val *[]*string) {
	if err := j.validateSetAllowedEmailPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedEmailPatterns",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetAllowedUserDomains(val *[]*string) {
	if err := j.validateSetAllowedUserDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUserDomains",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetEnabled(val *string) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2EnableSpInitiated(val *string) {
	if err := j.validateSetSaml2EnableSpInitiatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2EnableSpInitiated",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2ForceAuthn(val *string) {
	if err := j.validateSetSaml2ForceAuthnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2ForceAuthn",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2Issuer(val *string) {
	if err := j.validateSetSaml2IssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2Issuer",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2PostLogoutRedirectUrl(val *string) {
	if err := j.validateSetSaml2PostLogoutRedirectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2PostLogoutRedirectUrl",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2Provider(val *string) {
	if err := j.validateSetSaml2ProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2Provider",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2RequestedNameidFormat(val *string) {
	if err := j.validateSetSaml2RequestedNameidFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2RequestedNameidFormat",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2SignRequest(val *string) {
	if err := j.validateSetSaml2SignRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2SignRequest",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2SnowflakeAcsUrl(val *string) {
	if err := j.validateSetSaml2SnowflakeAcsUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2SnowflakeAcsUrl",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2SnowflakeIssuerUrl(val *string) {
	if err := j.validateSetSaml2SnowflakeIssuerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2SnowflakeIssuerUrl",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2SpInitiatedLoginPageLabel(val *string) {
	if err := j.validateSetSaml2SpInitiatedLoginPageLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2SpInitiatedLoginPageLabel",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2SsoUrl(val *string) {
	if err := j.validateSetSaml2SsoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2SsoUrl",
		val,
	)
}

func (j *jsiiProxy_Saml2Integration)SetSaml2X509Cert(val *string) {
	if err := j.validateSetSaml2X509CertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saml2X509Cert",
		val,
	)
}

// Generates CDKTF code for importing a Saml2Integration resource upon running "cdktf plan <stack-name>".
func Saml2Integration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSaml2Integration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.saml2Integration.Saml2Integration",
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
func Saml2Integration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSaml2Integration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.saml2Integration.Saml2Integration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Saml2Integration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSaml2Integration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.saml2Integration.Saml2Integration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Saml2Integration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSaml2Integration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.saml2Integration.Saml2Integration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Saml2Integration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.saml2Integration.Saml2Integration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Saml2Integration) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_Saml2Integration) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Saml2Integration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_Saml2Integration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Saml2Integration) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_Saml2Integration) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Saml2Integration) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Saml2Integration) PutTimeouts(value *Saml2IntegrationTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Saml2Integration) ResetAllowedEmailPatterns() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedEmailPatterns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetAllowedUserDomains() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedUserDomains",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetSaml2EnableSpInitiated() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2EnableSpInitiated",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetSaml2ForceAuthn() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2ForceAuthn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetSaml2PostLogoutRedirectUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2PostLogoutRedirectUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetSaml2RequestedNameidFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2RequestedNameidFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetSaml2SignRequest() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SignRequest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetSaml2SnowflakeAcsUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SnowflakeAcsUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetSaml2SnowflakeIssuerUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SnowflakeIssuerUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetSaml2SpInitiatedLoginPageLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetSaml2SpInitiatedLoginPageLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Saml2Integration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Saml2Integration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

