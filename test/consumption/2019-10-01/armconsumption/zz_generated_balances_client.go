//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BalancesClient contains the methods for the Balances group.
// Don't use this type directly, use NewBalancesClient() instead.
type BalancesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewBalancesClient creates a new instance of BalancesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBalancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BalancesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &BalancesClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// GetByBillingAccount - Gets the balances for a scope by billingAccountId. Balances are available via this API only for May
// 1, 2014 or later.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountID - BillingAccount ID
// options - BalancesClientGetByBillingAccountOptions contains the optional parameters for the BalancesClient.GetByBillingAccount
// method.
func (client *BalancesClient) GetByBillingAccount(ctx context.Context, billingAccountID string, options *BalancesClientGetByBillingAccountOptions) (BalancesClientGetByBillingAccountResponse, error) {
	req, err := client.getByBillingAccountCreateRequest(ctx, billingAccountID, options)
	if err != nil {
		return BalancesClientGetByBillingAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BalancesClientGetByBillingAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BalancesClientGetByBillingAccountResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByBillingAccountHandleResponse(resp)
}

// getByBillingAccountCreateRequest creates the GetByBillingAccount request.
func (client *BalancesClient) getByBillingAccountCreateRequest(ctx context.Context, billingAccountID string, options *BalancesClientGetByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/balances"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByBillingAccountHandleResponse handles the GetByBillingAccount response.
func (client *BalancesClient) getByBillingAccountHandleResponse(resp *http.Response) (BalancesClientGetByBillingAccountResponse, error) {
	result := BalancesClientGetByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Balance); err != nil {
		return BalancesClientGetByBillingAccountResponse{}, err
	}
	return result, nil
}

// GetForBillingPeriodByBillingAccount - Gets the balances for a scope by billing period and billingAccountId. Balances are
// available via this API only for May 1, 2014 or later.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountID - BillingAccount ID
// billingPeriodName - Billing Period Name.
// options - BalancesClientGetForBillingPeriodByBillingAccountOptions contains the optional parameters for the BalancesClient.GetForBillingPeriodByBillingAccount
// method.
func (client *BalancesClient) GetForBillingPeriodByBillingAccount(ctx context.Context, billingAccountID string, billingPeriodName string, options *BalancesClientGetForBillingPeriodByBillingAccountOptions) (BalancesClientGetForBillingPeriodByBillingAccountResponse, error) {
	req, err := client.getForBillingPeriodByBillingAccountCreateRequest(ctx, billingAccountID, billingPeriodName, options)
	if err != nil {
		return BalancesClientGetForBillingPeriodByBillingAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BalancesClientGetForBillingPeriodByBillingAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BalancesClientGetForBillingPeriodByBillingAccountResponse{}, runtime.NewResponseError(resp)
	}
	return client.getForBillingPeriodByBillingAccountHandleResponse(resp)
}

// getForBillingPeriodByBillingAccountCreateRequest creates the GetForBillingPeriodByBillingAccount request.
func (client *BalancesClient) getForBillingPeriodByBillingAccountCreateRequest(ctx context.Context, billingAccountID string, billingPeriodName string, options *BalancesClientGetForBillingPeriodByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/balances"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if billingPeriodName == "" {
		return nil, errors.New("parameter billingPeriodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingPeriodName}", url.PathEscape(billingPeriodName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getForBillingPeriodByBillingAccountHandleResponse handles the GetForBillingPeriodByBillingAccount response.
func (client *BalancesClient) getForBillingPeriodByBillingAccountHandleResponse(resp *http.Response) (BalancesClientGetForBillingPeriodByBillingAccountResponse, error) {
	result := BalancesClientGetForBillingPeriodByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Balance); err != nil {
		return BalancesClientGetForBillingPeriodByBillingAccountResponse{}, err
	}
	return result, nil
}
