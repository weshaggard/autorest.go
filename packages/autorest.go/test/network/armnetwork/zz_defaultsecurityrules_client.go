//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DefaultSecurityRulesClient contains the methods for the DefaultSecurityRules group.
// Don't use this type directly, use NewDefaultSecurityRulesClient() instead.
type DefaultSecurityRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDefaultSecurityRulesClient creates a new instance of DefaultSecurityRulesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDefaultSecurityRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DefaultSecurityRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".DefaultSecurityRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DefaultSecurityRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the specified default network security rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - defaultSecurityRuleName - The name of the default security rule.
//   - options - DefaultSecurityRulesClientGetOptions contains the optional parameters for the DefaultSecurityRulesClient.Get
//     method.
func (client *DefaultSecurityRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string, options *DefaultSecurityRulesClientGetOptions) (DefaultSecurityRulesClientGetResponse, error) {
	var err error
	const operationName = "DefaultSecurityRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, defaultSecurityRuleName, options)
	if err != nil {
		return DefaultSecurityRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefaultSecurityRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DefaultSecurityRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DefaultSecurityRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string, options *DefaultSecurityRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules/{defaultSecurityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if defaultSecurityRuleName == "" {
		return nil, errors.New("parameter defaultSecurityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{defaultSecurityRuleName}", url.PathEscape(defaultSecurityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DefaultSecurityRulesClient) getHandleResponse(resp *http.Response) (DefaultSecurityRulesClientGetResponse, error) {
	result := DefaultSecurityRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRule); err != nil {
		return DefaultSecurityRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all default security rules in a network security group.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityGroupName - The name of the network security group.
//   - options - DefaultSecurityRulesClientListOptions contains the optional parameters for the DefaultSecurityRulesClient.NewListPager
//     method.
func (client *DefaultSecurityRulesClient) NewListPager(resourceGroupName string, networkSecurityGroupName string, options *DefaultSecurityRulesClientListOptions) *runtime.Pager[DefaultSecurityRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefaultSecurityRulesClientListResponse]{
		More: func(page DefaultSecurityRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefaultSecurityRulesClientListResponse) (DefaultSecurityRulesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DefaultSecurityRulesClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DefaultSecurityRulesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DefaultSecurityRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DefaultSecurityRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DefaultSecurityRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *DefaultSecurityRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DefaultSecurityRulesClient) listHandleResponse(resp *http.Response) (DefaultSecurityRulesClientListResponse, error) {
	result := DefaultSecurityRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityRuleListResult); err != nil {
		return DefaultSecurityRulesClientListResponse{}, err
	}
	return result, nil
}
