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
	"strings"
)

// TagsClient contains the methods for the Tags group.
// Don't use this type directly, use NewTagsClient() instead.
type TagsClient struct {
	internal *arm.Client
}

// NewTagsClient creates a new instance of TagsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTagsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TagsClient, error) {
	cl, err := arm.NewClient(moduleName+".TagsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TagsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get all available tag keys for the defined scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - scope - The scope associated with tags operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope,
//     '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
//     resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope,
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
//     scope and
//     '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope..
//   - options - TagsClientGetOptions contains the optional parameters for the TagsClient.Get method.
func (client *TagsClient) Get(ctx context.Context, scope string, options *TagsClientGetOptions) (TagsClientGetResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TagsClient.Get")
	req, err := client.getCreateRequest(ctx, scope, options)
	if err != nil {
		return TagsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TagsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TagsClient) getCreateRequest(ctx context.Context, scope string, options *TagsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/tags"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TagsClient) getHandleResponse(resp *http.Response) (TagsClientGetResponse, error) {
	result := TagsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResult); err != nil {
		return TagsClientGetResponse{}, err
	}
	return result, nil
}
