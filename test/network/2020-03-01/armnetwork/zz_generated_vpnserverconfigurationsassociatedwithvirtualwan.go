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

// VpnServerConfigurationsAssociatedWithVirtualWanClient contains the methods for the VpnServerConfigurationsAssociatedWithVirtualWan group.
// Don't use this type directly, use NewVpnServerConfigurationsAssociatedWithVirtualWanClient() instead.
type VpnServerConfigurationsAssociatedWithVirtualWanClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVpnServerConfigurationsAssociatedWithVirtualWanClient creates a new instance of VpnServerConfigurationsAssociatedWithVirtualWanClient with the specified values.
func NewVpnServerConfigurationsAssociatedWithVirtualWanClient(con *armcore.Connection, subscriptionID string) VpnServerConfigurationsAssociatedWithVirtualWanClient {
	return VpnServerConfigurationsAssociatedWithVirtualWanClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginList - Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) BeginList(ctx context.Context, resourceGroupName string, virtualWanName string, options *VpnServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (VpnServerConfigurationsResponsePollerResponse, error) {
	resp, err := client.list(ctx, resourceGroupName, virtualWanName, options)
	if err != nil {
		return VpnServerConfigurationsResponsePollerResponse{}, err
	}
	result := VpnServerConfigurationsResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnServerConfigurationsAssociatedWithVirtualWanClient.List", "location", resp, client.listHandleError)
	if err != nil {
		return VpnServerConfigurationsResponsePollerResponse{}, err
	}
	poller := &vpnServerConfigurationsResponsePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VpnServerConfigurationsResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeList creates a new VpnServerConfigurationsResponsePoller from the specified resume token.
// token - The value must come from a previous call to VpnServerConfigurationsResponsePoller.ResumeToken().
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) ResumeList(token string) (VpnServerConfigurationsResponsePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnServerConfigurationsAssociatedWithVirtualWanClient.List", token, client.listHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnServerConfigurationsResponsePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// List - Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) list(ctx context.Context, resourceGroupName string, virtualWanName string, options *VpnServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (*azcore.Response, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, virtualWanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listHandleError(resp)
	}
	return resp, nil
}

// listCreateRequest creates the List request.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualWanName string, options *VpnServerConfigurationsAssociatedWithVirtualWanBeginListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnServerConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWanName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// listHandleResponse handles the List response.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) listHandleResponse(resp *azcore.Response) (VpnServerConfigurationsResponseResponse, error) {
	var val *VpnServerConfigurationsResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VpnServerConfigurationsResponseResponse{}, err
	}
	return VpnServerConfigurationsResponseResponse{RawResponse: resp.Response, VpnServerConfigurationsResponse: val}, nil
}

// listHandleError handles the List error response.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
