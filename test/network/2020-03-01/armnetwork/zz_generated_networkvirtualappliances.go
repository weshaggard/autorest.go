// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// NetworkVirtualAppliancesOperations contains the methods for the NetworkVirtualAppliances group.
type NetworkVirtualAppliancesOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters NetworkVirtualAppliance, options *NetworkVirtualAppliancesCreateOrUpdateOptions) (*NetworkVirtualAppliancePollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (NetworkVirtualAppliancePoller, error)
	// BeginDelete - Deletes the specified Network Virtual Appliance.
	BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified Network Virtual Appliance.
	Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesGetOptions) (*NetworkVirtualApplianceResponse, error)
	// List - Gets all Network Virtual Appliances in a subscription.
	List(options *NetworkVirtualAppliancesListOptions) NetworkVirtualApplianceListResultPager
	// ListByResourceGroup - Lists all Network Virtual Appliances in a resource group.
	ListByResourceGroup(resourceGroupName string, options *NetworkVirtualAppliancesListByResourceGroupOptions) NetworkVirtualApplianceListResultPager
	// UpdateTags - Updates a Network Virtual Appliance.
	UpdateTags(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *NetworkVirtualAppliancesUpdateTagsOptions) (*NetworkVirtualApplianceResponse, error)
}

// NetworkVirtualAppliancesClient implements the NetworkVirtualAppliancesOperations interface.
// Don't use this type directly, use NewNetworkVirtualAppliancesClient() instead.
type NetworkVirtualAppliancesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkVirtualAppliancesClient creates a new instance of NetworkVirtualAppliancesClient with the specified values.
func NewNetworkVirtualAppliancesClient(con *armcore.Connection, subscriptionID string) NetworkVirtualAppliancesOperations {
	return &NetworkVirtualAppliancesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *NetworkVirtualAppliancesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *NetworkVirtualAppliancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters NetworkVirtualAppliance, options *NetworkVirtualAppliancesCreateOrUpdateOptions) (*NetworkVirtualAppliancePollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &NetworkVirtualAppliancePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkVirtualAppliancesClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &networkVirtualAppliancePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*NetworkVirtualApplianceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *NetworkVirtualAppliancesClient) ResumeCreateOrUpdate(token string) (NetworkVirtualAppliancePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkVirtualAppliancesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &networkVirtualAppliancePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters NetworkVirtualAppliance, options *NetworkVirtualAppliancesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NetworkVirtualAppliancesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters NetworkVirtualAppliance, options *NetworkVirtualAppliancesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NetworkVirtualAppliancesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*NetworkVirtualApplianceResponse, error) {
	result := NetworkVirtualApplianceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkVirtualAppliance)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NetworkVirtualAppliancesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *NetworkVirtualAppliancesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, networkVirtualApplianceName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkVirtualAppliancesClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *NetworkVirtualAppliancesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkVirtualAppliancesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) Delete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *NetworkVirtualAppliancesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *NetworkVirtualAppliancesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesGetOptions) (*NetworkVirtualApplianceResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *NetworkVirtualAppliancesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *NetworkVirtualAppliancesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *NetworkVirtualAppliancesClient) GetHandleResponse(resp *azcore.Response) (*NetworkVirtualApplianceResponse, error) {
	result := NetworkVirtualApplianceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkVirtualAppliance)
}

// GetHandleError handles the Get error response.
func (client *NetworkVirtualAppliancesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all Network Virtual Appliances in a subscription.
func (client *NetworkVirtualAppliancesClient) List(options *NetworkVirtualAppliancesListOptions) NetworkVirtualApplianceListResultPager {
	return &networkVirtualApplianceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *NetworkVirtualApplianceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkVirtualApplianceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *NetworkVirtualAppliancesClient) ListCreateRequest(ctx context.Context, options *NetworkVirtualAppliancesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualAppliances"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *NetworkVirtualAppliancesClient) ListHandleResponse(resp *azcore.Response) (*NetworkVirtualApplianceListResultResponse, error) {
	result := NetworkVirtualApplianceListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkVirtualApplianceListResult)
}

// ListHandleError handles the List error response.
func (client *NetworkVirtualAppliancesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all Network Virtual Appliances in a resource group.
func (client *NetworkVirtualAppliancesClient) ListByResourceGroup(resourceGroupName string, options *NetworkVirtualAppliancesListByResourceGroupOptions) NetworkVirtualApplianceListResultPager {
	return &networkVirtualApplianceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *NetworkVirtualApplianceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkVirtualApplianceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkVirtualAppliancesClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkVirtualAppliancesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkVirtualAppliancesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*NetworkVirtualApplianceListResultResponse, error) {
	result := NetworkVirtualApplianceListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkVirtualApplianceListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *NetworkVirtualAppliancesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates a Network Virtual Appliance.
func (client *NetworkVirtualAppliancesClient) UpdateTags(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *NetworkVirtualAppliancesUpdateTagsOptions) (*NetworkVirtualApplianceResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *NetworkVirtualAppliancesClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters TagsObject, options *NetworkVirtualAppliancesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *NetworkVirtualAppliancesClient) UpdateTagsHandleResponse(resp *azcore.Response) (*NetworkVirtualApplianceResponse, error) {
	result := NetworkVirtualApplianceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkVirtualAppliance)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *NetworkVirtualAppliancesClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}