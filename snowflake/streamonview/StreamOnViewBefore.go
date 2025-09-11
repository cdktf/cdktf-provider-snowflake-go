// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamonview


type StreamOnViewBefore struct {
	// Specifies the difference in seconds from the current time to use for Time Travel, in the form -N where N can be an integer or arithmetic expression (e.g. -120 is 120 seconds, -30*60 is 1800 seconds or 30 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/stream_on_view#offset StreamOnView#offset}
	Offset *string `field:"optional" json:"offset" yaml:"offset"`
	// Specifies the query ID of a statement to use as the reference point for Time Travel.
	//
	// This parameter supports any statement of one of the following types: DML (e.g. INSERT, UPDATE, DELETE), TCL (BEGIN, COMMIT transaction), SELECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/stream_on_view#statement StreamOnView#statement}
	Statement *string `field:"optional" json:"statement" yaml:"statement"`
	// Specifies the identifier (i.e. name) for an existing stream on the queried table or view. The current offset in the stream is used as the AT point in time for returning change data for the source object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/stream_on_view#stream StreamOnView#stream}
	Stream *string `field:"optional" json:"stream" yaml:"stream"`
	// Specifies an exact date and time to use for Time Travel.
	//
	// The value must be explicitly cast to a TIMESTAMP, TIMESTAMP_LTZ, TIMESTAMP_NTZ, or TIMESTAMP_TZ data type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.7.0/docs/resources/stream_on_view#timestamp StreamOnView#timestamp}
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

