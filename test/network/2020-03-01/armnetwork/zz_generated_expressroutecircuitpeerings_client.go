//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// ExpressRouteCircuitPeeringsClient contains the methods for the ExpressRouteCircuitPeerings group.
// Don't use this type directly, use NewExpressRouteCircuitPeeringsClient() instead.
type ExpressRouteCircuitPeeringsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteCircuitPeeringsClient creates a new instance of ExpressRouteCircuitPeeringsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteCircuitPeeringsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteCircuitPeeringsClient, error) {
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
	client := &ExpressRouteCircuitPeeringsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a peering in the specified express route circuits.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// peeringParameters - Parameters supplied to the create or update express route circuit peering operation.
// options - ExpressRouteCircuitPeeringsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteCircuitPeeringsClient.BeginCreateOrUpdate
// method.
func (client *ExpressRouteCircuitPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering, options *ExpressRouteCircuitPeeringsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ExpressRouteCircuitPeeringsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, circuitName, peeringName, peeringParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ExpressRouteCircuitPeeringsClientCreateOrUpdateResponse]("ExpressRouteCircuitPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ExpressRouteCircuitPeeringsClientCreateOrUpdateResponse]("ExpressRouteCircuitPeeringsClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a peering in the specified express route circuits.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCircuitPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering, options *ExpressRouteCircuitPeeringsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, peeringName, peeringParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteCircuitPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering, options *ExpressRouteCircuitPeeringsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, peeringParameters)
}

// BeginDelete - Deletes the specified peering from the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// options - ExpressRouteCircuitPeeringsClientBeginDeleteOptions contains the optional parameters for the ExpressRouteCircuitPeeringsClient.BeginDelete
// method.
func (client *ExpressRouteCircuitPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsClientBeginDeleteOptions) (*armruntime.Poller[ExpressRouteCircuitPeeringsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, circuitName, peeringName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ExpressRouteCircuitPeeringsClientDeleteResponse]("ExpressRouteCircuitPeeringsClient.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ExpressRouteCircuitPeeringsClientDeleteResponse]("ExpressRouteCircuitPeeringsClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified peering from the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRouteCircuitPeeringsClient) deleteOperation(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExpressRouteCircuitPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified peering for the express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// options - ExpressRouteCircuitPeeringsClientGetOptions contains the optional parameters for the ExpressRouteCircuitPeeringsClient.Get
// method.
func (client *ExpressRouteCircuitPeeringsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsClientGetOptions) (ExpressRouteCircuitPeeringsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
	if err != nil {
		return ExpressRouteCircuitPeeringsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCircuitPeeringsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCircuitPeeringsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCircuitPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitPeeringsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteCircuitPeeringsClient) getHandleResponse(resp *http.Response) (ExpressRouteCircuitPeeringsClientGetResponse, error) {
	result := ExpressRouteCircuitPeeringsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitPeering); err != nil {
		return ExpressRouteCircuitPeeringsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all peerings in a specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// options - ExpressRouteCircuitPeeringsClientListOptions contains the optional parameters for the ExpressRouteCircuitPeeringsClient.List
// method.
func (client *ExpressRouteCircuitPeeringsClient) List(resourceGroupName string, circuitName string, options *ExpressRouteCircuitPeeringsClientListOptions) *runtime.Pager[ExpressRouteCircuitPeeringsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ExpressRouteCircuitPeeringsClientListResponse]{
		More: func(page ExpressRouteCircuitPeeringsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteCircuitPeeringsClientListResponse) (ExpressRouteCircuitPeeringsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, circuitName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRouteCircuitPeeringsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExpressRouteCircuitPeeringsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRouteCircuitPeeringsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCircuitPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitPeeringsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteCircuitPeeringsClient) listHandleResponse(resp *http.Response) (ExpressRouteCircuitPeeringsClientListResponse, error) {
	result := ExpressRouteCircuitPeeringsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitPeeringListResult); err != nil {
		return ExpressRouteCircuitPeeringsClientListResponse{}, err
	}
	return result, nil
}
