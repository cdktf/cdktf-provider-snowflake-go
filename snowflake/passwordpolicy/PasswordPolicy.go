// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package passwordpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v12/passwordpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/password_policy snowflake_password_policy}.
type PasswordPolicy interface {
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullyQualifiedName() *string
	History() *float64
	SetHistory(val *float64)
	HistoryInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	IfNotExists() interface{}
	SetIfNotExists(val interface{})
	IfNotExistsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LockoutTimeMins() *float64
	SetLockoutTimeMins(val *float64)
	LockoutTimeMinsInput() *float64
	MaxAgeDays() *float64
	SetMaxAgeDays(val *float64)
	MaxAgeDaysInput() *float64
	MaxLength() *float64
	SetMaxLength(val *float64)
	MaxLengthInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MinAgeDays() *float64
	SetMinAgeDays(val *float64)
	MinAgeDaysInput() *float64
	MinLength() *float64
	SetMinLength(val *float64)
	MinLengthInput() *float64
	MinLowerCaseChars() *float64
	SetMinLowerCaseChars(val *float64)
	MinLowerCaseCharsInput() *float64
	MinNumericChars() *float64
	SetMinNumericChars(val *float64)
	MinNumericCharsInput() *float64
	MinSpecialChars() *float64
	SetMinSpecialChars(val *float64)
	MinSpecialCharsInput() *float64
	MinUpperCaseChars() *float64
	SetMinUpperCaseChars(val *float64)
	MinUpperCaseCharsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OrReplace() interface{}
	SetOrReplace(val interface{})
	OrReplaceInput() interface{}
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
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
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
	ResetHistory()
	ResetId()
	ResetIfNotExists()
	ResetLockoutTimeMins()
	ResetMaxAgeDays()
	ResetMaxLength()
	ResetMaxRetries()
	ResetMinAgeDays()
	ResetMinLength()
	ResetMinLowerCaseChars()
	ResetMinNumericChars()
	ResetMinSpecialChars()
	ResetMinUpperCaseChars()
	ResetOrReplace()
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

// The jsii proxy struct for PasswordPolicy
type jsiiProxy_PasswordPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PasswordPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) History() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"history",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) HistoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"historyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) IfNotExists() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ifNotExists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) IfNotExistsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ifNotExistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) LockoutTimeMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockoutTimeMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) LockoutTimeMinsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockoutTimeMinsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MaxAgeDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MaxAgeDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MaxLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinAgeDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAgeDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinAgeDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAgeDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinLowerCaseChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLowerCaseChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinLowerCaseCharsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLowerCaseCharsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinNumericChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumericChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinNumericCharsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumericCharsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinSpecialChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSpecialChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinSpecialCharsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSpecialCharsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinUpperCaseChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpperCaseChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) MinUpperCaseCharsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpperCaseCharsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) OrReplace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orReplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) OrReplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orReplaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/password_policy snowflake_password_policy} Resource.
func NewPasswordPolicy(scope constructs.Construct, id *string, config *PasswordPolicyConfig) PasswordPolicy {
	_init_.Initialize()

	if err := validateNewPasswordPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PasswordPolicy{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/1.0.5/docs/resources/password_policy snowflake_password_policy} Resource.
func NewPasswordPolicy_Override(p PasswordPolicy, scope constructs.Construct, id *string, config *PasswordPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicy",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetHistory(val *float64) {
	if err := j.validateSetHistoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"history",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetIfNotExists(val interface{}) {
	if err := j.validateSetIfNotExistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ifNotExists",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetLockoutTimeMins(val *float64) {
	if err := j.validateSetLockoutTimeMinsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockoutTimeMins",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMaxAgeDays(val *float64) {
	if err := j.validateSetMaxAgeDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAgeDays",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMaxLength(val *float64) {
	if err := j.validateSetMaxLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMinAgeDays(val *float64) {
	if err := j.validateSetMinAgeDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minAgeDays",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMinLength(val *float64) {
	if err := j.validateSetMinLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMinLowerCaseChars(val *float64) {
	if err := j.validateSetMinLowerCaseCharsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLowerCaseChars",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMinNumericChars(val *float64) {
	if err := j.validateSetMinNumericCharsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumericChars",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMinSpecialChars(val *float64) {
	if err := j.validateSetMinSpecialCharsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSpecialChars",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetMinUpperCaseChars(val *float64) {
	if err := j.validateSetMinUpperCaseCharsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minUpperCaseChars",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetOrReplace(val interface{}) {
	if err := j.validateSetOrReplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orReplace",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicy)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

// Generates CDKTF code for importing a PasswordPolicy resource upon running "cdktf plan <stack-name>".
func PasswordPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePasswordPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicy",
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
func PasswordPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePasswordPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PasswordPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePasswordPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PasswordPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePasswordPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PasswordPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.passwordPolicy.PasswordPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PasswordPolicy) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PasswordPolicy) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PasswordPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PasswordPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PasswordPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PasswordPolicy) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PasswordPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetComment() {
	_jsii_.InvokeVoid(
		p,
		"resetComment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetHistory() {
	_jsii_.InvokeVoid(
		p,
		"resetHistory",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetIfNotExists() {
	_jsii_.InvokeVoid(
		p,
		"resetIfNotExists",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetLockoutTimeMins() {
	_jsii_.InvokeVoid(
		p,
		"resetLockoutTimeMins",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMaxAgeDays() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxAgeDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMaxLength() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxLength",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMinAgeDays() {
	_jsii_.InvokeVoid(
		p,
		"resetMinAgeDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMinLength() {
	_jsii_.InvokeVoid(
		p,
		"resetMinLength",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMinLowerCaseChars() {
	_jsii_.InvokeVoid(
		p,
		"resetMinLowerCaseChars",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMinNumericChars() {
	_jsii_.InvokeVoid(
		p,
		"resetMinNumericChars",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMinSpecialChars() {
	_jsii_.InvokeVoid(
		p,
		"resetMinSpecialChars",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetMinUpperCaseChars() {
	_jsii_.InvokeVoid(
		p,
		"resetMinUpperCaseChars",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetOrReplace() {
	_jsii_.InvokeVoid(
		p,
		"resetOrReplace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PasswordPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

