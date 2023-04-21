package fileformat

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v6/fileformat/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.62.0/docs/resources/file_format snowflake_file_format}.
type FileFormat interface {
	cdktf.TerraformResource
	AllowDuplicate() interface{}
	SetAllowDuplicate(val interface{})
	AllowDuplicateInput() interface{}
	BinaryAsText() interface{}
	SetBinaryAsText(val interface{})
	BinaryAsTextInput() interface{}
	BinaryFormat() *string
	SetBinaryFormat(val *string)
	BinaryFormatInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
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
	DateFormat() *string
	SetDateFormat(val *string)
	DateFormatInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableAutoConvert() interface{}
	SetDisableAutoConvert(val interface{})
	DisableAutoConvertInput() interface{}
	DisableSnowflakeData() interface{}
	SetDisableSnowflakeData(val interface{})
	DisableSnowflakeDataInput() interface{}
	EmptyFieldAsNull() interface{}
	SetEmptyFieldAsNull(val interface{})
	EmptyFieldAsNullInput() interface{}
	EnableOctal() interface{}
	SetEnableOctal(val interface{})
	EnableOctalInput() interface{}
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	ErrorOnColumnCountMismatch() interface{}
	SetErrorOnColumnCountMismatch(val interface{})
	ErrorOnColumnCountMismatchInput() interface{}
	Escape() *string
	SetEscape(val *string)
	EscapeInput() *string
	EscapeUnenclosedField() *string
	SetEscapeUnenclosedField(val *string)
	EscapeUnenclosedFieldInput() *string
	FieldDelimiter() *string
	SetFieldDelimiter(val *string)
	FieldDelimiterInput() *string
	FieldOptionallyEnclosedBy() *string
	SetFieldOptionallyEnclosedBy(val *string)
	FieldOptionallyEnclosedByInput() *string
	FileExtension() *string
	SetFileExtension(val *string)
	FileExtensionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	FormatType() *string
	SetFormatType(val *string)
	FormatTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreUtf8Errors() interface{}
	SetIgnoreUtf8Errors(val interface{})
	IgnoreUtf8ErrorsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NullIf() *[]*string
	SetNullIf(val *[]*string)
	NullIfInput() *[]*string
	PreserveSpace() interface{}
	SetPreserveSpace(val interface{})
	PreserveSpaceInput() interface{}
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
	RecordDelimiter() *string
	SetRecordDelimiter(val *string)
	RecordDelimiterInput() *string
	ReplaceInvalidCharacters() interface{}
	SetReplaceInvalidCharacters(val interface{})
	ReplaceInvalidCharactersInput() interface{}
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	SkipBlankLines() interface{}
	SetSkipBlankLines(val interface{})
	SkipBlankLinesInput() interface{}
	SkipByteOrderMark() interface{}
	SetSkipByteOrderMark(val interface{})
	SkipByteOrderMarkInput() interface{}
	SkipHeader() *float64
	SetSkipHeader(val *float64)
	SkipHeaderInput() *float64
	StripNullValues() interface{}
	SetStripNullValues(val interface{})
	StripNullValuesInput() interface{}
	StripOuterArray() interface{}
	SetStripOuterArray(val interface{})
	StripOuterArrayInput() interface{}
	StripOuterElement() interface{}
	SetStripOuterElement(val interface{})
	StripOuterElementInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeFormat() *string
	SetTimeFormat(val *string)
	TimeFormatInput() *string
	TimestampFormat() *string
	SetTimestampFormat(val *string)
	TimestampFormatInput() *string
	TrimSpace() interface{}
	SetTrimSpace(val interface{})
	TrimSpaceInput() interface{}
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
	ResetAllowDuplicate()
	ResetBinaryAsText()
	ResetBinaryFormat()
	ResetComment()
	ResetCompression()
	ResetDateFormat()
	ResetDisableAutoConvert()
	ResetDisableSnowflakeData()
	ResetEmptyFieldAsNull()
	ResetEnableOctal()
	ResetEncoding()
	ResetErrorOnColumnCountMismatch()
	ResetEscape()
	ResetEscapeUnenclosedField()
	ResetFieldDelimiter()
	ResetFieldOptionallyEnclosedBy()
	ResetFileExtension()
	ResetId()
	ResetIgnoreUtf8Errors()
	ResetNullIf()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreserveSpace()
	ResetRecordDelimiter()
	ResetReplaceInvalidCharacters()
	ResetSkipBlankLines()
	ResetSkipByteOrderMark()
	ResetSkipHeader()
	ResetStripNullValues()
	ResetStripOuterArray()
	ResetStripOuterElement()
	ResetTimeFormat()
	ResetTimestampFormat()
	ResetTrimSpace()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FileFormat
type jsiiProxy_FileFormat struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FileFormat) AllowDuplicate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) AllowDuplicateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) BinaryAsText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"binaryAsText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) BinaryAsTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"binaryAsTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) BinaryFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) BinaryFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) DateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) DisableAutoConvert() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoConvert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) DisableAutoConvertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoConvertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) DisableSnowflakeData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSnowflakeData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) DisableSnowflakeDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSnowflakeDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) EmptyFieldAsNull() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emptyFieldAsNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) EmptyFieldAsNullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emptyFieldAsNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) EnableOctal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOctal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) EnableOctalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOctalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) ErrorOnColumnCountMismatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnColumnCountMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) ErrorOnColumnCountMismatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorOnColumnCountMismatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Escape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) EscapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) EscapeUnenclosedField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeUnenclosedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) EscapeUnenclosedFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeUnenclosedFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FieldOptionallyEnclosedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldOptionallyEnclosedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FieldOptionallyEnclosedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldOptionallyEnclosedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FileExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FormatType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FormatTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) IgnoreUtf8Errors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUtf8Errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) IgnoreUtf8ErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUtf8ErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) NullIf() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) NullIfInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) PreserveSpace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) PreserveSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) RecordDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) RecordDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) ReplaceInvalidCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) ReplaceInvalidCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) SkipBlankLines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBlankLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) SkipBlankLinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipBlankLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) SkipByteOrderMark() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipByteOrderMark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) SkipByteOrderMarkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipByteOrderMarkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) SkipHeader() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) SkipHeaderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) StripNullValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripNullValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) StripNullValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripNullValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) StripOuterArray() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripOuterArray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) StripOuterArrayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripOuterArrayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) StripOuterElement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripOuterElement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) StripOuterElementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripOuterElementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TrimSpace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileFormat) TrimSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimSpaceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.62.0/docs/resources/file_format snowflake_file_format} Resource.
func NewFileFormat(scope constructs.Construct, id *string, config *FileFormatConfig) FileFormat {
	_init_.Initialize()

	if err := validateNewFileFormatParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FileFormat{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.fileFormat.FileFormat",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.62.0/docs/resources/file_format snowflake_file_format} Resource.
func NewFileFormat_Override(f FileFormat, scope constructs.Construct, id *string, config *FileFormatConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.fileFormat.FileFormat",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FileFormat)SetAllowDuplicate(val interface{}) {
	if err := j.validateSetAllowDuplicateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicate",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetBinaryAsText(val interface{}) {
	if err := j.validateSetBinaryAsTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryAsText",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetBinaryFormat(val *string) {
	if err := j.validateSetBinaryFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryFormat",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetDateFormat(val *string) {
	if err := j.validateSetDateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateFormat",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetDisableAutoConvert(val interface{}) {
	if err := j.validateSetDisableAutoConvertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoConvert",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetDisableSnowflakeData(val interface{}) {
	if err := j.validateSetDisableSnowflakeDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableSnowflakeData",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetEmptyFieldAsNull(val interface{}) {
	if err := j.validateSetEmptyFieldAsNullParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyFieldAsNull",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetEnableOctal(val interface{}) {
	if err := j.validateSetEnableOctalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOctal",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetErrorOnColumnCountMismatch(val interface{}) {
	if err := j.validateSetErrorOnColumnCountMismatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnColumnCountMismatch",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetEscape(val *string) {
	if err := j.validateSetEscapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escape",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetEscapeUnenclosedField(val *string) {
	if err := j.validateSetEscapeUnenclosedFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escapeUnenclosedField",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetFieldOptionallyEnclosedBy(val *string) {
	if err := j.validateSetFieldOptionallyEnclosedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldOptionallyEnclosedBy",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetFileExtension(val *string) {
	if err := j.validateSetFileExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileExtension",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetFormatType(val *string) {
	if err := j.validateSetFormatTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatType",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetIgnoreUtf8Errors(val interface{}) {
	if err := j.validateSetIgnoreUtf8ErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUtf8Errors",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetNullIf(val *[]*string) {
	if err := j.validateSetNullIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullIf",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetPreserveSpace(val interface{}) {
	if err := j.validateSetPreserveSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveSpace",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetRecordDelimiter(val *string) {
	if err := j.validateSetRecordDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordDelimiter",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetReplaceInvalidCharacters(val interface{}) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetSkipBlankLines(val interface{}) {
	if err := j.validateSetSkipBlankLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipBlankLines",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetSkipByteOrderMark(val interface{}) {
	if err := j.validateSetSkipByteOrderMarkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipByteOrderMark",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetSkipHeader(val *float64) {
	if err := j.validateSetSkipHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipHeader",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetStripNullValues(val interface{}) {
	if err := j.validateSetStripNullValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripNullValues",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetStripOuterArray(val interface{}) {
	if err := j.validateSetStripOuterArrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripOuterArray",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetStripOuterElement(val interface{}) {
	if err := j.validateSetStripOuterElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripOuterElement",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetTimeFormat(val *string) {
	if err := j.validateSetTimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeFormat",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetTimestampFormat(val *string) {
	if err := j.validateSetTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampFormat",
		val,
	)
}

func (j *jsiiProxy_FileFormat)SetTrimSpace(val interface{}) {
	if err := j.validateSetTrimSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trimSpace",
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
func FileFormat_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFileFormat_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.fileFormat.FileFormat",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FileFormat_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFileFormat_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.fileFormat.FileFormat",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FileFormat_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFileFormat_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-snowflake.fileFormat.FileFormat",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FileFormat_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-snowflake.fileFormat.FileFormat",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FileFormat) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FileFormat) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FileFormat) ResetAllowDuplicate() {
	_jsii_.InvokeVoid(
		f,
		"resetAllowDuplicate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetBinaryAsText() {
	_jsii_.InvokeVoid(
		f,
		"resetBinaryAsText",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetBinaryFormat() {
	_jsii_.InvokeVoid(
		f,
		"resetBinaryFormat",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetComment() {
	_jsii_.InvokeVoid(
		f,
		"resetComment",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetCompression() {
	_jsii_.InvokeVoid(
		f,
		"resetCompression",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetDateFormat() {
	_jsii_.InvokeVoid(
		f,
		"resetDateFormat",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetDisableAutoConvert() {
	_jsii_.InvokeVoid(
		f,
		"resetDisableAutoConvert",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetDisableSnowflakeData() {
	_jsii_.InvokeVoid(
		f,
		"resetDisableSnowflakeData",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetEmptyFieldAsNull() {
	_jsii_.InvokeVoid(
		f,
		"resetEmptyFieldAsNull",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetEnableOctal() {
	_jsii_.InvokeVoid(
		f,
		"resetEnableOctal",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetEncoding() {
	_jsii_.InvokeVoid(
		f,
		"resetEncoding",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetErrorOnColumnCountMismatch() {
	_jsii_.InvokeVoid(
		f,
		"resetErrorOnColumnCountMismatch",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetEscape() {
	_jsii_.InvokeVoid(
		f,
		"resetEscape",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetEscapeUnenclosedField() {
	_jsii_.InvokeVoid(
		f,
		"resetEscapeUnenclosedField",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		f,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetFieldOptionallyEnclosedBy() {
	_jsii_.InvokeVoid(
		f,
		"resetFieldOptionallyEnclosedBy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetFileExtension() {
	_jsii_.InvokeVoid(
		f,
		"resetFileExtension",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetIgnoreUtf8Errors() {
	_jsii_.InvokeVoid(
		f,
		"resetIgnoreUtf8Errors",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetNullIf() {
	_jsii_.InvokeVoid(
		f,
		"resetNullIf",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetPreserveSpace() {
	_jsii_.InvokeVoid(
		f,
		"resetPreserveSpace",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetRecordDelimiter() {
	_jsii_.InvokeVoid(
		f,
		"resetRecordDelimiter",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		f,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetSkipBlankLines() {
	_jsii_.InvokeVoid(
		f,
		"resetSkipBlankLines",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetSkipByteOrderMark() {
	_jsii_.InvokeVoid(
		f,
		"resetSkipByteOrderMark",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetSkipHeader() {
	_jsii_.InvokeVoid(
		f,
		"resetSkipHeader",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetStripNullValues() {
	_jsii_.InvokeVoid(
		f,
		"resetStripNullValues",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetStripOuterArray() {
	_jsii_.InvokeVoid(
		f,
		"resetStripOuterArray",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetStripOuterElement() {
	_jsii_.InvokeVoid(
		f,
		"resetStripOuterElement",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetTimeFormat() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeFormat",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetTimestampFormat() {
	_jsii_.InvokeVoid(
		f,
		"resetTimestampFormat",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) ResetTrimSpace() {
	_jsii_.InvokeVoid(
		f,
		"resetTrimSpace",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileFormat) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileFormat) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

