package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-snowflake-go/snowflake/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-snowflake-go/snowflake/v3/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/snowflake/r/user snowflake_user}.
type User interface {
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultNamespace() *string
	SetDefaultNamespace(val *string)
	DefaultNamespaceInput() *string
	DefaultRole() *string
	SetDefaultRole(val *string)
	DefaultRoleInput() *string
	DefaultSecondaryRoles() *[]*string
	SetDefaultSecondaryRoles(val *[]*string)
	DefaultSecondaryRolesInput() *[]*string
	DefaultWarehouse() *string
	SetDefaultWarehouse(val *string)
	DefaultWarehouseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HasRsaPublicKey() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoginName() *string
	SetLoginName(val *string)
	LoginNameInput() *string
	MustChangePassword() interface{}
	SetMustChangePassword(val interface{})
	MustChangePasswordInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	Tag() UserTagList
	TagInput() interface{}
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
	PutTag(value interface{})
	ResetComment()
	ResetDefaultNamespace()
	ResetDefaultRole()
	ResetDefaultSecondaryRoles()
	ResetDefaultWarehouse()
	ResetDisabled()
	ResetDisplayName()
	ResetEmail()
	ResetFirstName()
	ResetId()
	ResetLastName()
	ResetLoginName()
	ResetMustChangePassword()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetRsaPublicKey()
	ResetRsaPublicKey2()
	ResetTag()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_User) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultSecondaryRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultSecondaryRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultSecondaryRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultSecondaryRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DefaultWarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HasRsaPublicKey() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasRsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LoginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LoginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MustChangePassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mustChangePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MustChangePasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mustChangePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RsaPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RsaPublicKey2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RsaPublicKey2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKey2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RsaPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Tag() UserTagList {
	var returns UserTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/user snowflake_user} Resource.
func NewUser(scope constructs.Construct, id *string, config *UserConfig) User {
	_init_.Initialize()

	if err := validateNewUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_User{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.user.User",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/snowflake/r/user snowflake_user} Resource.
func NewUser_Override(u User, scope constructs.Construct, id *string, config *UserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.user.User",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_User)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_User)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_User)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_User)SetDefaultNamespace(val *string) {
	if err := j.validateSetDefaultNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNamespace",
		val,
	)
}

func (j *jsiiProxy_User)SetDefaultRole(val *string) {
	if err := j.validateSetDefaultRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRole",
		val,
	)
}

func (j *jsiiProxy_User)SetDefaultSecondaryRoles(val *[]*string) {
	if err := j.validateSetDefaultSecondaryRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSecondaryRoles",
		val,
	)
}

func (j *jsiiProxy_User)SetDefaultWarehouse(val *string) {
	if err := j.validateSetDefaultWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultWarehouse",
		val,
	)
}

func (j *jsiiProxy_User)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_User)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_User)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_User)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_User)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_User)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_User)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_User)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_User)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_User)SetLoginName(val *string) {
	if err := j.validateSetLoginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginName",
		val,
	)
}

func (j *jsiiProxy_User)SetMustChangePassword(val interface{}) {
	if err := j.validateSetMustChangePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mustChangePassword",
		val,
	)
}

func (j *jsiiProxy_User)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_User)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_User)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_User)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_User)SetRsaPublicKey(val *string) {
	if err := j.validateSetRsaPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey",
		val,
	)
}

func (j *jsiiProxy_User)SetRsaPublicKey2(val *string) {
	if err := j.validateSetRsaPublicKey2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaPublicKey2",
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
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.user.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func User_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.user.User",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_User) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_User) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_User) PutTag(value interface{}) {
	if err := u.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putTag",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) ResetComment() {
	_jsii_.InvokeVoid(
		u,
		"resetComment",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDefaultNamespace() {
	_jsii_.InvokeVoid(
		u,
		"resetDefaultNamespace",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDefaultRole() {
	_jsii_.InvokeVoid(
		u,
		"resetDefaultRole",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDefaultSecondaryRoles() {
	_jsii_.InvokeVoid(
		u,
		"resetDefaultSecondaryRoles",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDefaultWarehouse() {
	_jsii_.InvokeVoid(
		u,
		"resetDefaultWarehouse",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDisabled() {
	_jsii_.InvokeVoid(
		u,
		"resetDisabled",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDisplayName() {
	_jsii_.InvokeVoid(
		u,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmail() {
	_jsii_.InvokeVoid(
		u,
		"resetEmail",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetFirstName() {
	_jsii_.InvokeVoid(
		u,
		"resetFirstName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLastName() {
	_jsii_.InvokeVoid(
		u,
		"resetLastName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLoginName() {
	_jsii_.InvokeVoid(
		u,
		"resetLoginName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMustChangePassword() {
	_jsii_.InvokeVoid(
		u,
		"resetMustChangePassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRsaPublicKey() {
	_jsii_.InvokeVoid(
		u,
		"resetRsaPublicKey",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRsaPublicKey2() {
	_jsii_.InvokeVoid(
		u,
		"resetRsaPublicKey2",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTag() {
	_jsii_.InvokeVoid(
		u,
		"resetTag",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

