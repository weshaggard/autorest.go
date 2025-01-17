// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armlargeinstance

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

// AzureLargeStorageInstancesClient contains the methods for the AzureLargeStorageInstances group.
// Don't use this type directly, use NewAzureLargeStorageInstancesClient() instead.
type AzureLargeStorageInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureLargeStorageInstancesClient creates a new instance of AzureLargeStorageInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureLargeStorageInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureLargeStorageInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureLargeStorageInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an Azure Large Storage instance for the specified subscription, resource
// group, and instance name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeStorageInstanceName - Name of the AzureLargeStorageInstance.
//   - options - AzureLargeStorageInstancesClientGetOptions contains the optional parameters for the AzureLargeStorageInstancesClient.Get
//     method.
func (client *AzureLargeStorageInstancesClient) Get(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, options *AzureLargeStorageInstancesClientGetOptions) (AzureLargeStorageInstancesClientGetResponse, error) {
	var err error
	const operationName = "AzureLargeStorageInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureLargeStorageInstanceName, options)
	if err != nil {
		return AzureLargeStorageInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureLargeStorageInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureLargeStorageInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AzureLargeStorageInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, _ *AzureLargeStorageInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeStorageInstances/{azureLargeStorageInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeStorageInstanceName == "" {
		return nil, errors.New("parameter azureLargeStorageInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeStorageInstanceName}", url.PathEscape(azureLargeStorageInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureLargeStorageInstancesClient) getHandleResponse(resp *http.Response) (AzureLargeStorageInstancesClientGetResponse, error) {
	result := AzureLargeStorageInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeStorageInstance); err != nil {
		return AzureLargeStorageInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of AzureLargeStorageInstances in the specified subscription and
// resource group. The operations returns various properties of each Azure
// LargeStorage instance.
//
// Generated from API version 2023-07-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AzureLargeStorageInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureLargeStorageInstancesClient.NewListByResourceGroupPager
//     method.
func (client *AzureLargeStorageInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *AzureLargeStorageInstancesClientListByResourceGroupOptions) *runtime.Pager[AzureLargeStorageInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureLargeStorageInstancesClientListByResourceGroupResponse]{
		More: func(page AzureLargeStorageInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureLargeStorageInstancesClientListByResourceGroupResponse) (AzureLargeStorageInstancesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureLargeStorageInstancesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AzureLargeStorageInstancesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AzureLargeStorageInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *AzureLargeStorageInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeStorageInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AzureLargeStorageInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (AzureLargeStorageInstancesClientListByResourceGroupResponse, error) {
	result := AzureLargeStorageInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeStorageInstanceListResult); err != nil {
		return AzureLargeStorageInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of AzureLargeStorageInstances in the specified subscription. The
// operations returns various properties of each Azure LargeStorage instance.
//
// Generated from API version 2023-07-20-preview
//   - options - AzureLargeStorageInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureLargeStorageInstancesClient.NewListBySubscriptionPager
//     method.
func (client *AzureLargeStorageInstancesClient) NewListBySubscriptionPager(options *AzureLargeStorageInstancesClientListBySubscriptionOptions) *runtime.Pager[AzureLargeStorageInstancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureLargeStorageInstancesClientListBySubscriptionResponse]{
		More: func(page AzureLargeStorageInstancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureLargeStorageInstancesClientListBySubscriptionResponse) (AzureLargeStorageInstancesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureLargeStorageInstancesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AzureLargeStorageInstancesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AzureLargeStorageInstancesClient) listBySubscriptionCreateRequest(ctx context.Context, _ *AzureLargeStorageInstancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureLargeInstance/azureLargeStorageInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AzureLargeStorageInstancesClient) listBySubscriptionHandleResponse(resp *http.Response) (AzureLargeStorageInstancesClientListBySubscriptionResponse, error) {
	result := AzureLargeStorageInstancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeStorageInstanceListResult); err != nil {
		return AzureLargeStorageInstancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patches the Tags field of a Azure Large Storage Instance for the specified
// subscription, resource group, and instance name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeStorageInstanceName - Name of the AzureLargeStorageInstance.
//   - properties - The resource properties to be updated.
//   - options - AzureLargeStorageInstancesClientUpdateOptions contains the optional parameters for the AzureLargeStorageInstancesClient.Update
//     method.
func (client *AzureLargeStorageInstancesClient) Update(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, properties AzureLargeStorageInstanceTagsUpdate, options *AzureLargeStorageInstancesClientUpdateOptions) (AzureLargeStorageInstancesClientUpdateResponse, error) {
	var err error
	const operationName = "AzureLargeStorageInstancesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, azureLargeStorageInstanceName, properties, options)
	if err != nil {
		return AzureLargeStorageInstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureLargeStorageInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureLargeStorageInstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AzureLargeStorageInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, properties AzureLargeStorageInstanceTagsUpdate, _ *AzureLargeStorageInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeStorageInstances/{azureLargeStorageInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeStorageInstanceName == "" {
		return nil, errors.New("parameter azureLargeStorageInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeStorageInstanceName}", url.PathEscape(azureLargeStorageInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AzureLargeStorageInstancesClient) updateHandleResponse(resp *http.Response) (AzureLargeStorageInstancesClientUpdateResponse, error) {
	result := AzureLargeStorageInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeStorageInstance); err != nil {
		return AzureLargeStorageInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
