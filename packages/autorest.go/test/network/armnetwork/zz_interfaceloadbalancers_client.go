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

// InterfaceLoadBalancersClient contains the methods for the NetworkInterfaceLoadBalancers group.
// Don't use this type directly, use NewInterfaceLoadBalancersClient() instead.
type InterfaceLoadBalancersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInterfaceLoadBalancersClient creates a new instance of InterfaceLoadBalancersClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInterfaceLoadBalancersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InterfaceLoadBalancersClient, error) {
	cl, err := arm.NewClient(moduleName+".InterfaceLoadBalancersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InterfaceLoadBalancersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List all load balancers in a network interface.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkInterfaceName - The name of the network interface.
//   - options - InterfaceLoadBalancersClientListOptions contains the optional parameters for the InterfaceLoadBalancersClient.NewListPager
//     method.
func (client *InterfaceLoadBalancersClient) NewListPager(resourceGroupName string, networkInterfaceName string, options *InterfaceLoadBalancersClientListOptions) *runtime.Pager[InterfaceLoadBalancersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InterfaceLoadBalancersClientListResponse]{
		More: func(page InterfaceLoadBalancersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InterfaceLoadBalancersClientListResponse) (InterfaceLoadBalancersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InterfaceLoadBalancersClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InterfaceLoadBalancersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InterfaceLoadBalancersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InterfaceLoadBalancersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *InterfaceLoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *InterfaceLoadBalancersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
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
func (client *InterfaceLoadBalancersClient) listHandleResponse(resp *http.Response) (InterfaceLoadBalancersClientListResponse, error) {
	result := InterfaceLoadBalancersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InterfaceLoadBalancerListResult); err != nil {
		return InterfaceLoadBalancersClientListResponse{}, err
	}
	return result, nil
}
