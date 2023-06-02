//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// MarketplacesClient contains the methods for the Marketplaces group.
// Don't use this type directly, use NewMarketplacesClient() instead.
type MarketplacesClient struct {
	internal *arm.Client
}

// NewMarketplacesClient creates a new instance of MarketplacesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMarketplacesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*MarketplacesClient, error) {
	cl, err := arm.NewClient(moduleName+".MarketplacesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MarketplacesClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Lists the marketplaces for a scope at the defined scope. Marketplaces are available via this API only for
// May 1, 2014 or later.
//
// Generated from API version 2019-10-01
//   - scope - The scope associated with marketplace operations. This includes '/subscriptions/{subscriptionId}/' for subscription
//     scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing
//     Account scope, '/providers/Microsoft.Billing/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountId}'
//     for EnrollmentAccount scope and
//     '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope. For subscription, billing
//     account, department, enrollment account and ManagementGroup, you can also
//     add billing period to the scope using '/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'. For e.g. to specify
//     billing period at department scope use
//     '/providers/Microsoft.Billing/departments/{departmentId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'
//   - options - MarketplacesClientListOptions contains the optional parameters for the MarketplacesClient.NewListPager method.
func (client *MarketplacesClient) NewListPager(scope string, options *MarketplacesClientListOptions) *runtime.Pager[MarketplacesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MarketplacesClientListResponse]{
		More: func(page MarketplacesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MarketplacesClientListResponse) (MarketplacesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MarketplacesClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MarketplacesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return MarketplacesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MarketplacesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MarketplacesClient) listCreateRequest(ctx context.Context, scope string, options *MarketplacesClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/marketplaces"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplacesClient) listHandleResponse(resp *http.Response) (MarketplacesClientListResponse, error) {
	result := MarketplacesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplacesListResult); err != nil {
		return MarketplacesClientListResponse{}, err
	}
	return result, nil
}
