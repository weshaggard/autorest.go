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

// VpnConnectionsClient contains the methods for the VpnConnections group.
// Don't use this type directly, use NewVpnConnectionsClient() instead.
type VpnConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVpnConnectionsClient creates a new instance of VpnConnectionsClient with the specified values.
func NewVpnConnectionsClient(con *armcore.Connection, subscriptionID string) VpnConnectionsClient {
	return VpnConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VpnConnectionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
func (client VpnConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection, options *VpnConnectionsBeginCreateOrUpdateOptions) (VpnConnectionPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
	if err != nil {
		return VpnConnectionPollerResponse{}, err
	}
	result := VpnConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VpnConnectionPollerResponse{}, err
	}
	poller := &vpnConnectionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VpnConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VpnConnectionPoller from the specified resume token.
// token - The value must come from a previous call to VpnConnectionPoller.ResumeToken().
func (client VpnConnectionsClient) ResumeCreateOrUpdate(token string) (VpnConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnConnectionsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
func (client VpnConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection, options *VpnConnectionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client VpnConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection, options *VpnConnectionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnConnectionParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client VpnConnectionsClient) createOrUpdateHandleResponse(resp *azcore.Response) (VpnConnectionResponse, error) {
	var val *VpnConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VpnConnectionResponse{}, err
	}
	return VpnConnectionResponse{RawResponse: resp.Response, VpnConnection: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client VpnConnectionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes a vpn connection.
func (client VpnConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VpnConnectionsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
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

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client VpnConnectionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VpnConnectionsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a vpn connection.
func (client VpnConnectionsClient) delete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client VpnConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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

// deleteHandleError handles the Delete error response.
func (client VpnConnectionsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a vpn connection.
func (client VpnConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsGetOptions) (VpnConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return VpnConnectionResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VpnConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VpnConnectionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client VpnConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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

// getHandleResponse handles the Get response.
func (client VpnConnectionsClient) getHandleResponse(resp *azcore.Response) (VpnConnectionResponse, error) {
	var val *VpnConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VpnConnectionResponse{}, err
	}
	return VpnConnectionResponse{RawResponse: resp.Response, VpnConnection: val}, nil
}

// getHandleError handles the Get error response.
func (client VpnConnectionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByVpnGateway - Retrieves all vpn connections for a particular virtual wan vpn gateway.
func (client VpnConnectionsClient) ListByVpnGateway(resourceGroupName string, gatewayName string, options *VpnConnectionsListByVpnGatewayOptions) ListVpnConnectionsResultPager {
	return &listVpnConnectionsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByVpnGatewayCreateRequest(ctx, resourceGroupName, gatewayName, options)
		},
		responder: client.listByVpnGatewayHandleResponse,
		errorer:   client.listByVpnGatewayHandleError,
		advancer: func(ctx context.Context, resp ListVpnConnectionsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnConnectionsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByVpnGatewayCreateRequest creates the ListByVpnGateway request.
func (client VpnConnectionsClient) listByVpnGatewayCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VpnConnectionsListByVpnGatewayOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// listByVpnGatewayHandleResponse handles the ListByVpnGateway response.
func (client VpnConnectionsClient) listByVpnGatewayHandleResponse(resp *azcore.Response) (ListVpnConnectionsResultResponse, error) {
	var val *ListVpnConnectionsResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVpnConnectionsResultResponse{}, err
	}
	return ListVpnConnectionsResultResponse{RawResponse: resp.Response, ListVpnConnectionsResult: val}, nil
}

// listByVpnGatewayHandleError handles the ListByVpnGateway error response.
func (client VpnConnectionsClient) listByVpnGatewayHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
