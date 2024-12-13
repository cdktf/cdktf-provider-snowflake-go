// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiIntegrationConfig struct {
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
	// Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service endpoints and resources within those proxies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#api_allowed_prefixes ApiIntegration#api_allowed_prefixes}
	ApiAllowedPrefixes *[]*string `field:"required" json:"apiAllowedPrefixes" yaml:"apiAllowedPrefixes"`
	// Specifies the HTTPS proxy service type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#api_provider ApiIntegration#api_provider}
	ApiProvider *string `field:"required" json:"apiProvider" yaml:"apiProvider"`
	// Specifies the name of the API integration.
	//
	// This name follows the rules for Object Identifiers. The name should be unique among api integrations in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#name ApiIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ARN of a cloud platform role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#api_aws_role_arn ApiIntegration#api_aws_role_arn}
	ApiAwsRoleArn *string `field:"optional" json:"apiAwsRoleArn" yaml:"apiAwsRoleArn"`
	// Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#api_blocked_prefixes ApiIntegration#api_blocked_prefixes}
	ApiBlockedPrefixes *[]*string `field:"optional" json:"apiBlockedPrefixes" yaml:"apiBlockedPrefixes"`
	// The service account used for communication with the Google API Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#api_gcp_service_account ApiIntegration#api_gcp_service_account}
	ApiGcpServiceAccount *string `field:"optional" json:"apiGcpServiceAccount" yaml:"apiGcpServiceAccount"`
	// The API key (also called a “subscription key”).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#api_key ApiIntegration#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The 'Application (client) id' of the Azure AD app for your remote service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#azure_ad_application_id ApiIntegration#azure_ad_application_id}
	AzureAdApplicationId *string `field:"optional" json:"azureAdApplicationId" yaml:"azureAdApplicationId"`
	// Specifies the ID for your Office 365 tenant that all Azure API Management instances belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#azure_tenant_id ApiIntegration#azure_tenant_id}
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#comment ApiIntegration#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies whether this API integration is enabled or disabled.
	//
	// If the API integration is disabled, any external function that relies on it will not work.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#enabled ApiIntegration#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The audience claim when generating the JWT (JSON Web Token) to authenticate to the Google API Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#google_audience ApiIntegration#google_audience}
	GoogleAudience *string `field:"optional" json:"googleAudience" yaml:"googleAudience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflake-labs/snowflake/0.100.0/docs/resources/api_integration#id ApiIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

