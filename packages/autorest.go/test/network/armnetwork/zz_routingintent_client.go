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

// RoutingIntentClient contains the methods for the RoutingIntent group.
// Don't use this type directly, use NewRoutingIntentClient() instead.
type RoutingIntentClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRoutingIntentClient creates a new instance of RoutingIntentClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoutingIntentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RoutingIntentClient, error) {
	cl, err := arm.NewClient(moduleName+".RoutingIntentClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoutingIntentClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a RoutingIntent resource if it doesn't exist else updates the existing RoutingIntent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the RoutingIntent.
//   - virtualHubName - The name of the VirtualHub.
//   - routingIntentName - The name of the per VirtualHub singleton Routing Intent resource.
//   - routingIntentParameters - Parameters supplied to create or update RoutingIntent.
//   - options - RoutingIntentClientBeginCreateOrUpdateOptions contains the optional parameters for the RoutingIntentClient.BeginCreateOrUpdate
//     method.
func (client *RoutingIntentClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, routingIntentParameters RoutingIntent, options *RoutingIntentClientBeginCreateOrUpdateOptions) (*runtime.Poller[RoutingIntentClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, routingIntentName, routingIntentParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RoutingIntentClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RoutingIntentClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a RoutingIntent resource if it doesn't exist else updates the existing RoutingIntent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *RoutingIntentClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, routingIntentParameters RoutingIntent, options *RoutingIntentClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RoutingIntentClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, routingIntentName, routingIntentParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RoutingIntentClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, routingIntentParameters RoutingIntent, options *RoutingIntentClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent/{routingIntentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routingIntentName == "" {
		return nil, errors.New("parameter routingIntentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routingIntentName}", url.PathEscape(routingIntentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, routingIntentParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a RoutingIntent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the RoutingIntent.
//   - virtualHubName - The name of the VirtualHub.
//   - routingIntentName - The name of the RoutingIntent.
//   - options - RoutingIntentClientBeginDeleteOptions contains the optional parameters for the RoutingIntentClient.BeginDelete
//     method.
func (client *RoutingIntentClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, options *RoutingIntentClientBeginDeleteOptions) (*runtime.Poller[RoutingIntentClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, routingIntentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RoutingIntentClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RoutingIntentClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a RoutingIntent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *RoutingIntentClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, options *RoutingIntentClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RoutingIntentClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, routingIntentName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RoutingIntentClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, options *RoutingIntentClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent/{routingIntentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routingIntentName == "" {
		return nil, errors.New("parameter routingIntentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routingIntentName}", url.PathEscape(routingIntentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a RoutingIntent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the RoutingIntent.
//   - virtualHubName - The name of the VirtualHub.
//   - routingIntentName - The name of the RoutingIntent.
//   - options - RoutingIntentClientGetOptions contains the optional parameters for the RoutingIntentClient.Get method.
func (client *RoutingIntentClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, options *RoutingIntentClientGetOptions) (RoutingIntentClientGetResponse, error) {
	var err error
	const operationName = "RoutingIntentClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, routingIntentName, options)
	if err != nil {
		return RoutingIntentClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoutingIntentClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoutingIntentClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RoutingIntentClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, options *RoutingIntentClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent/{routingIntentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routingIntentName == "" {
		return nil, errors.New("parameter routingIntentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routingIntentName}", url.PathEscape(routingIntentName))
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
func (client *RoutingIntentClient) getHandleResponse(resp *http.Response) (RoutingIntentClientGetResponse, error) {
	result := RoutingIntentClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoutingIntent); err != nil {
		return RoutingIntentClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves the details of all RoutingIntent child resources of the VirtualHub.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - options - RoutingIntentClientListOptions contains the optional parameters for the RoutingIntentClient.NewListPager method.
func (client *RoutingIntentClient) NewListPager(resourceGroupName string, virtualHubName string, options *RoutingIntentClientListOptions) *runtime.Pager[RoutingIntentClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoutingIntentClientListResponse]{
		More: func(page RoutingIntentClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoutingIntentClientListResponse) (RoutingIntentClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RoutingIntentClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoutingIntentClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoutingIntentClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoutingIntentClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RoutingIntentClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *RoutingIntentClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
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
func (client *RoutingIntentClient) listHandleResponse(resp *http.Response) (RoutingIntentClientListResponse, error) {
	result := RoutingIntentClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListRoutingIntentResult); err != nil {
		return RoutingIntentClientListResponse{}, err
	}
	return result, nil
}
