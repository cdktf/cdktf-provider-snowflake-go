package apiintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v5/apiintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/snowflake/r/api_integration snowflake_api_integration}.
type ApiIntegration interface {
	cdktf.TerraformResource
	ApiAllowedPrefixes() *[]*string
	SetApiAllowedPrefixes(val *[]*string)
	ApiAllowedPrefixesInput() *[]*string
	ApiAwsExternalId() *string
	ApiAwsIamUserArn() *string
	ApiAwsRoleArn() *string
	SetApiAwsRoleArn(val *string)
	ApiAwsRoleArnInput() *string
	ApiBlockedPrefixes() *[]*string
	SetApiBlockedPrefixes(val *[]*string)
	ApiBlockedPrefixesInput() *[]*string
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	ApiProvider() *string
	SetApiProvider(val *string)
	ApiProviderInput() *string
	AzureAdApplicationId() *string
	SetAzureAdApplicationId(val *string)
	AzureAdApplicationIdInput() *string
	AzureConsentUrl() *string
	AzureMultiTenantAppName() *string
	AzureTenantId() *string
	SetAzureTenantId(val *string)
	AzureTenantIdInput() *string
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
	GoogleAudience() *string
	SetGoogleAudience(val *string)
	GoogleAudienceInput() *string
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
	ResetApiAwsRoleArn()
	ResetApiBlockedPrefixes()
	ResetApiKey()
	ResetAzureAdApplicationId()
	ResetAzureTenantId()
	ResetComment()
	ResetEnabled()
	ResetGoogleAudience()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApiIntegration
type jsiiProxy_ApiIntegration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApiIntegration) ApiAllowedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiAllowedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiAllowedPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiAllowedPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiAwsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiAwsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiAwsIamUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiAwsIamUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiAwsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiAwsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiAwsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiAwsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiBlockedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiBlockedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiBlockedPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiBlockedPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ApiProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) AzureAdApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAdApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) AzureAdApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAdApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) AzureConsentUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureConsentUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) AzureMultiTenantAppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureMultiTenantAppName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) AzureTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) GoogleAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) GoogleAudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/api_integration snowflake_api_integration} Resource.
func NewApiIntegration(scope constructs.Construct, id *string, config *ApiIntegrationConfig) ApiIntegration {
	_init_.Initialize()

	if err := validateNewApiIntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiIntegration{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.apiIntegration.ApiIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/api_integration snowflake_api_integration} Resource.
func NewApiIntegration_Override(a ApiIntegration, scope constructs.Construct, id *string, config *ApiIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.apiIntegration.ApiIntegration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiIntegration)SetApiAllowedPrefixes(val *[]*string) {
	if err := j.validateSetApiAllowedPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiAllowedPrefixes",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetApiAwsRoleArn(val *string) {
	if err := j.validateSetApiAwsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiAwsRoleArn",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetApiBlockedPrefixes(val *[]*string) {
	if err := j.validateSetApiBlockedPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiBlockedPrefixes",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetApiKey(val *string) {
	if err := j.validateSetApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetApiProvider(val *string) {
	if err := j.validateSetApiProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiProvider",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetAzureAdApplicationId(val *string) {
	if err := j.validateSetAzureAdApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureAdApplicationId",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetAzureTenantId(val *string) {
	if err := j.validateSetAzureTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureTenantId",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetGoogleAudience(val *string) {
	if err := j.validateSetGoogleAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleAudience",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiIntegration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func ApiIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.apiIntegration.ApiIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiIntegration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiIntegration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.apiIntegration.ApiIntegration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiIntegration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiIntegration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.apiIntegration.ApiIntegration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.apiIntegration.ApiIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiIntegration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiIntegration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiIntegration) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiIntegration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiIntegration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiIntegration) ResetApiAwsRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetApiAwsRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetApiBlockedPrefixes() {
	_jsii_.InvokeVoid(
		a,
		"resetApiBlockedPrefixes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetAzureAdApplicationId() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureAdApplicationId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetAzureTenantId() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureTenantId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetComment() {
	_jsii_.InvokeVoid(
		a,
		"resetComment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetGoogleAudience() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAudience",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

