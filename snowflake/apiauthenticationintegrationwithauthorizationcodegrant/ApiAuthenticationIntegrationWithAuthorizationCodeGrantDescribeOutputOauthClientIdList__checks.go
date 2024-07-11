// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package apiauthenticationintegrationwithauthorizationcodegrant

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputOauthClientIdList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputOauthClientIdList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputOauthClientIdList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputOauthClientIdList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputOauthClientIdList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputOauthClientIdList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewApiAuthenticationIntegrationWithAuthorizationCodeGrantDescribeOutputOauthClientIdListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

