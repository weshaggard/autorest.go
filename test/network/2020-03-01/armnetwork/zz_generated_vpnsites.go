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

// VpnSitesOperations contains the methods for the VpnSites group.
type VpnSitesOperations interface {
	// BeginCreateOrUpdate - Creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VpnSite, options *VpnSitesCreateOrUpdateOptions) (*VpnSitePollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VpnSitePoller, error)
	// BeginDelete - Deletes a VpnSite.
	BeginDelete(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VpnSitesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the details of a VPN site.
	Get(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VpnSitesGetOptions) (*VpnSiteResponse, error)
	// List - Lists all the VpnSites in a subscription.
	List(options *VpnSitesListOptions) ListVpnSitesResultPager
	// ListByResourceGroup - Lists all the vpnSites in a resource group.
	ListByResourceGroup(resourceGroupName string, options *VpnSitesListByResourceGroupOptions) ListVpnSitesResultPager
	// UpdateTags - Updates VpnSite tags.
	UpdateTags(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject, options *VpnSitesUpdateTagsOptions) (*VpnSiteResponse, error)
}

// VpnSitesClient implements the VpnSitesOperations interface.
// Don't use this type directly, use NewVpnSitesClient() instead.
type VpnSitesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVpnSitesClient creates a new instance of VpnSitesClient with the specified values.
func NewVpnSitesClient(con *armcore.Connection, subscriptionID string) VpnSitesOperations {
	return &VpnSitesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VpnSitesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *VpnSitesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VpnSite, options *VpnSitesCreateOrUpdateOptions) (*VpnSitePollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
	if err != nil {
		return nil, err
	}
	result := &VpnSitePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnSitesClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &vpnSitePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VpnSiteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VpnSitesClient) ResumeCreateOrUpdate(token string) (VpnSitePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnSitesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnSitePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
func (client *VpnSitesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VpnSite, options *VpnSitesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
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
func (client *VpnSitesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VpnSite, options *VpnSitesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnSiteParameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VpnSitesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VpnSiteResponse, error) {
	result := VpnSiteResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnSite)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VpnSitesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VpnSitesClient) BeginDelete(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VpnSitesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, vpnSiteName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnSitesClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *VpnSitesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnSitesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a VpnSite.
func (client *VpnSitesClient) Delete(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VpnSitesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
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
func (client *VpnSitesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VpnSitesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
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
func (client *VpnSitesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a VPN site.
func (client *VpnSitesClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VpnSitesGetOptions) (*VpnSiteResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
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
func (client *VpnSitesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VpnSitesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
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

// GetHandleResponse handles the Get response.
func (client *VpnSitesClient) GetHandleResponse(resp *azcore.Response) (*VpnSiteResponse, error) {
	result := VpnSiteResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnSite)
}

// GetHandleError handles the Get error response.
func (client *VpnSitesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the VpnSites in a subscription.
func (client *VpnSitesClient) List(options *VpnSitesListOptions) ListVpnSitesResultPager {
	return &listVpnSitesResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ListVpnSitesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnSitesResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *VpnSitesClient) ListCreateRequest(ctx context.Context, options *VpnSitesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites"
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
func (client *VpnSitesClient) ListHandleResponse(resp *azcore.Response) (*ListVpnSitesResultResponse, error) {
	result := ListVpnSitesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnSitesResult)
}

// ListHandleError handles the List error response.
func (client *VpnSitesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all the vpnSites in a resource group.
func (client *VpnSitesClient) ListByResourceGroup(resourceGroupName string, options *VpnSitesListByResourceGroupOptions) ListVpnSitesResultPager {
	return &listVpnSitesResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *ListVpnSitesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnSitesResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VpnSitesClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VpnSitesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *VpnSitesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*ListVpnSitesResultResponse, error) {
	result := ListVpnSitesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnSitesResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VpnSitesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates VpnSite tags.
func (client *VpnSitesClient) UpdateTags(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject, options *VpnSitesUpdateTagsOptions) (*VpnSiteResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
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
func (client *VpnSitesClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject, options *VpnSitesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnSiteParameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *VpnSitesClient) UpdateTagsHandleResponse(resp *azcore.Response) (*VpnSiteResponse, error) {
	result := VpnSiteResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnSite)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *VpnSitesClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}