package externaltablegrant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v8/externaltablegrant/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/external_table_grant snowflake_external_table_grant}.
type ExternalTableGrant interface {
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableMultipleGrants() interface{}
	SetEnableMultipleGrants(val interface{})
	EnableMultipleGrantsInput() interface{}
	ExternalTableName() *string
	SetExternalTableName(val *string)
	ExternalTableNameInput() *string
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
	// The tree node.
	Node() constructs.Node
	OnAll() interface{}
	SetOnAll(val interface{})
	OnAllInput() interface{}
	OnFuture() interface{}
	SetOnFuture(val interface{})
	OnFutureInput() interface{}
	Privilege() *string
	SetPrivilege(val *string)
	PrivilegeInput() *string
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
	RevertOwnershipToRoleName() *string
	SetRevertOwnershipToRoleName(val *string)
	RevertOwnershipToRoleNameInput() *string
	Roles() *[]*string
	SetRoles(val *[]*string)
	RolesInput() *[]*string
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
	Shares() *[]*string
	SetShares(val *[]*string)
	SharesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithGrantOption() interface{}
	SetWithGrantOption(val interface{})
	WithGrantOptionInput() interface{}
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
	ResetEnableMultipleGrants()
	ResetExternalTableName()
	ResetId()
	ResetOnAll()
	ResetOnFuture()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivilege()
	ResetRevertOwnershipToRoleName()
	ResetSchemaName()
	ResetShares()
	ResetWithGrantOption()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ExternalTableGrant
type jsiiProxy_ExternalTableGrant struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ExternalTableGrant) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) EnableMultipleGrants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) EnableMultipleGrantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleGrantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) ExternalTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) ExternalTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) OnAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) OnAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) OnFuture() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFuture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) OnFutureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFutureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Privilege() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilege",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) PrivilegeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) RevertOwnershipToRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revertOwnershipToRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) RevertOwnershipToRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revertOwnershipToRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) RolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) Shares() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"shares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) SharesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) WithGrantOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalTableGrant) WithGrantOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/external_table_grant snowflake_external_table_grant} Resource.
func NewExternalTableGrant(scope constructs.Construct, id *string, config *ExternalTableGrantConfig) ExternalTableGrant {
	_init_.Initialize()

	if err := validateNewExternalTableGrantParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalTableGrant{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalTableGrant.ExternalTableGrant",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.67.0/docs/resources/external_table_grant snowflake_external_table_grant} Resource.
func NewExternalTableGrant_Override(e ExternalTableGrant, scope constructs.Construct, id *string, config *ExternalTableGrantConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.externalTableGrant.ExternalTableGrant",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetEnableMultipleGrants(val interface{}) {
	if err := j.validateSetEnableMultipleGrantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMultipleGrants",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetExternalTableName(val *string) {
	if err := j.validateSetExternalTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalTableName",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetOnAll(val interface{}) {
	if err := j.validateSetOnAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onAll",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetOnFuture(val interface{}) {
	if err := j.validateSetOnFutureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFuture",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetPrivilege(val *string) {
	if err := j.validateSetPrivilegeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilege",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetRevertOwnershipToRoleName(val *string) {
	if err := j.validateSetRevertOwnershipToRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revertOwnershipToRoleName",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetRoles(val *[]*string) {
	if err := j.validateSetRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetShares(val *[]*string) {
	if err := j.validateSetSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shares",
		val,
	)
}

func (j *jsiiProxy_ExternalTableGrant)SetWithGrantOption(val interface{}) {
	if err := j.validateSetWithGrantOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withGrantOption",
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
func ExternalTableGrant_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalTableGrant_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalTableGrant.ExternalTableGrant",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalTableGrant_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalTableGrant_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalTableGrant.ExternalTableGrant",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalTableGrant_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalTableGrant_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.externalTableGrant.ExternalTableGrant",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExternalTableGrant_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.externalTableGrant.ExternalTableGrant",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExternalTableGrant) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExternalTableGrant) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ExternalTableGrant) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ExternalTableGrant) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ExternalTableGrant) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ExternalTableGrant) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ExternalTableGrant) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ExternalTableGrant) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ExternalTableGrant) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ExternalTableGrant) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ExternalTableGrant) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ExternalTableGrant) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetEnableMultipleGrants() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableMultipleGrants",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetExternalTableName() {
	_jsii_.InvokeVoid(
		e,
		"resetExternalTableName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetOnAll() {
	_jsii_.InvokeVoid(
		e,
		"resetOnAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetOnFuture() {
	_jsii_.InvokeVoid(
		e,
		"resetOnFuture",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetPrivilege() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivilege",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetRevertOwnershipToRoleName() {
	_jsii_.InvokeVoid(
		e,
		"resetRevertOwnershipToRoleName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetSchemaName() {
	_jsii_.InvokeVoid(
		e,
		"resetSchemaName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetShares() {
	_jsii_.InvokeVoid(
		e,
		"resetShares",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) ResetWithGrantOption() {
	_jsii_.InvokeVoid(
		e,
		"resetWithGrantOption",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalTableGrant) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalTableGrant) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalTableGrant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalTableGrant) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

