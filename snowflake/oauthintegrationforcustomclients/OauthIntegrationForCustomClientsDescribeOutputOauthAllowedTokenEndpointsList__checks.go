// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package oauthintegrationforcustomclients

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewOauthIntegrationForCustomClientsDescribeOutputOauthAllowedTokenEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

