// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fileformat

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FileFormatConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The database in which to create the file format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#database FileFormat#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the format of the input files (for data loading) or output files (for data unloading).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#format_type FileFormat#format_type}
	FormatType *string `field:"required" json:"formatType" yaml:"formatType"`
	// Specifies the identifier for the file format;
	//
	// must be unique for the database and schema in which the file format is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#name FileFormat#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the file format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#schema FileFormat#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Boolean that specifies to allow duplicate object field names (only the last one will be preserved).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#allow_duplicate FileFormat#allow_duplicate}
	AllowDuplicate interface{} `field:"optional" json:"allowDuplicate" yaml:"allowDuplicate"`
	// Boolean that specifies whether to interpret columns with no defined logical data type as UTF-8 text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#binary_as_text FileFormat#binary_as_text}
	BinaryAsText interface{} `field:"optional" json:"binaryAsText" yaml:"binaryAsText"`
	// Defines the encoding format for binary input or output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#binary_format FileFormat#binary_format}
	BinaryFormat *string `field:"optional" json:"binaryFormat" yaml:"binaryFormat"`
	// Specifies a comment for the file format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#comment FileFormat#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the current compression algorithm for the data file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#compression FileFormat#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Defines the format of date values in the data files (data loading) or table (data unloading).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#date_format FileFormat#date_format}
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// Boolean that specifies whether the XML parser disables automatic conversion of numeric and Boolean values from text to native representation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#disable_auto_convert FileFormat#disable_auto_convert}
	DisableAutoConvert interface{} `field:"optional" json:"disableAutoConvert" yaml:"disableAutoConvert"`
	// Boolean that specifies whether the XML parser disables recognition of Snowflake semi-structured data tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#disable_snowflake_data FileFormat#disable_snowflake_data}
	DisableSnowflakeData interface{} `field:"optional" json:"disableSnowflakeData" yaml:"disableSnowflakeData"`
	// Specifies whether to insert SQL NULL for empty fields in an input file, which are represented by two successive delimiters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#empty_field_as_null FileFormat#empty_field_as_null}
	EmptyFieldAsNull interface{} `field:"optional" json:"emptyFieldAsNull" yaml:"emptyFieldAsNull"`
	// Boolean that enables parsing of octal numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#enable_octal FileFormat#enable_octal}
	EnableOctal interface{} `field:"optional" json:"enableOctal" yaml:"enableOctal"`
	// String (constant) that specifies the character set of the source data when loading data into a table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#encoding FileFormat#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Boolean that specifies whether to generate a parsing error if the number of delimited columns (i.e. fields) in an input file does not match the number of columns in the corresponding table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#error_on_column_count_mismatch FileFormat#error_on_column_count_mismatch}
	ErrorOnColumnCountMismatch interface{} `field:"optional" json:"errorOnColumnCountMismatch" yaml:"errorOnColumnCountMismatch"`
	// Single character string used as the escape character for field values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#escape FileFormat#escape}
	Escape *string `field:"optional" json:"escape" yaml:"escape"`
	// Single character string used as the escape character for unenclosed field values only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#escape_unenclosed_field FileFormat#escape_unenclosed_field}
	EscapeUnenclosedField *string `field:"optional" json:"escapeUnenclosedField" yaml:"escapeUnenclosedField"`
	// Specifies one or more singlebyte or multibyte characters that separate fields in an input file (data loading) or unloaded file (data unloading).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#field_delimiter FileFormat#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Character used to enclose strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#field_optionally_enclosed_by FileFormat#field_optionally_enclosed_by}
	FieldOptionallyEnclosedBy *string `field:"optional" json:"fieldOptionallyEnclosedBy" yaml:"fieldOptionallyEnclosedBy"`
	// Specifies the extension for files unloaded to a stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#file_extension FileFormat#file_extension}
	FileExtension *string `field:"optional" json:"fileExtension" yaml:"fileExtension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#id FileFormat#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Boolean that specifies whether UTF-8 encoding errors produce error conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#ignore_utf8_errors FileFormat#ignore_utf8_errors}
	IgnoreUtf8Errors interface{} `field:"optional" json:"ignoreUtf8Errors" yaml:"ignoreUtf8Errors"`
	// String used to convert to and from SQL NULL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#null_if FileFormat#null_if}
	NullIf *[]*string `field:"optional" json:"nullIf" yaml:"nullIf"`
	// Boolean that specifies whether to use the first row headers in the data files to determine column names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#parse_header FileFormat#parse_header}
	ParseHeader interface{} `field:"optional" json:"parseHeader" yaml:"parseHeader"`
	// Boolean that specifies whether the XML parser preserves leading and trailing spaces in element content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#preserve_space FileFormat#preserve_space}
	PreserveSpace interface{} `field:"optional" json:"preserveSpace" yaml:"preserveSpace"`
	// Specifies one or more singlebyte or multibyte characters that separate records in an input file (data loading) or unloaded file (data unloading).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#record_delimiter FileFormat#record_delimiter}
	RecordDelimiter *string `field:"optional" json:"recordDelimiter" yaml:"recordDelimiter"`
	// Boolean that specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (�).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#replace_invalid_characters FileFormat#replace_invalid_characters}
	ReplaceInvalidCharacters interface{} `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// Boolean that specifies to skip any blank lines encountered in the data files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#skip_blank_lines FileFormat#skip_blank_lines}
	SkipBlankLines interface{} `field:"optional" json:"skipBlankLines" yaml:"skipBlankLines"`
	// Boolean that specifies whether to skip the BOM (byte order mark), if present in a data file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#skip_byte_order_mark FileFormat#skip_byte_order_mark}
	SkipByteOrderMark interface{} `field:"optional" json:"skipByteOrderMark" yaml:"skipByteOrderMark"`
	// Number of lines at the start of the file to skip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#skip_header FileFormat#skip_header}
	SkipHeader *float64 `field:"optional" json:"skipHeader" yaml:"skipHeader"`
	// Boolean that instructs the JSON parser to remove object fields or array elements containing null values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#strip_null_values FileFormat#strip_null_values}
	StripNullValues interface{} `field:"optional" json:"stripNullValues" yaml:"stripNullValues"`
	// Boolean that instructs the JSON parser to remove outer brackets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#strip_outer_array FileFormat#strip_outer_array}
	StripOuterArray interface{} `field:"optional" json:"stripOuterArray" yaml:"stripOuterArray"`
	// Boolean that specifies whether the XML parser strips out the outer XML element, exposing 2nd level elements as separate documents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#strip_outer_element FileFormat#strip_outer_element}
	StripOuterElement interface{} `field:"optional" json:"stripOuterElement" yaml:"stripOuterElement"`
	// Defines the format of time values in the data files (data loading) or table (data unloading).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#time_format FileFormat#time_format}
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// Defines the format of timestamp values in the data files (data loading) or table (data unloading).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#timestamp_format FileFormat#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// Boolean that specifies whether to remove white space from fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.92.0/docs/resources/file_format#trim_space FileFormat#trim_space}
	TrimSpace interface{} `field:"optional" json:"trimSpace" yaml:"trimSpace"`
}

