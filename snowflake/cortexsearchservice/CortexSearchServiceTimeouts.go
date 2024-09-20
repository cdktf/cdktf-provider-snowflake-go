// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cortexsearchservice


type CortexSearchServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/cortex_search_service#create CortexSearchService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.96.0/docs/resources/cortex_search_service#update CortexSearchService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

