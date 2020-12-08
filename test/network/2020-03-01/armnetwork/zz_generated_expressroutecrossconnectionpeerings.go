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

// ExpressRouteCrossConnectionPeeringsClient contains the methods for the ExpressRouteCrossConnectionPeerings group.
// Don't use this type directly, use NewExpressRouteCrossConnectionPeeringsClient() instead.
type ExpressRouteCrossConnectionPeeringsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCrossConnectionPeeringsClient creates a new instance of ExpressRouteCrossConnectionPeeringsClient with the specified values.
func NewExpressRouteCrossConnectionPeeringsClient(con *armcore.Connection, subscriptionID string) ExpressRouteCrossConnectionPeeringsClient {
	return ExpressRouteCrossConnectionPeeringsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client ExpressRouteCrossConnectionPeeringsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (ExpressRouteCrossConnectionPeeringPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringPollerResponse{}, err
	}
	result := ExpressRouteCrossConnectionPeeringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringPollerResponse{}, err
	}
	poller := &expressRouteCrossConnectionPeeringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCrossConnectionPeeringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteCrossConnectionPeeringPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCrossConnectionPeeringPoller.ResumeToken().
func (client ExpressRouteCrossConnectionPeeringsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPeeringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionPeeringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters, options)
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
func (client ExpressRouteCrossConnectionPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering, options *ExpressRouteCrossConnectionPeeringsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
	return req, req.MarshalAsJSON(peeringParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client ExpressRouteCrossConnectionPeeringsClient) createOrUpdateHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionPeeringResponse, error) {
	var val *ExpressRouteCrossConnectionPeering
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCrossConnectionPeeringResponse{}, err
	}
	return ExpressRouteCrossConnectionPeeringResponse{RawResponse: resp.Response, ExpressRouteCrossConnectionPeering: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client ExpressRouteCrossConnectionPeeringsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified peering from the ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionPeeringsClient.Delete", "location", resp, client.deleteHandleError)
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
func (client ExpressRouteCrossConnectionPeeringsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionPeeringsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified peering from the ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionPeeringsClient) delete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
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
func (client ExpressRouteCrossConnectionPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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

// deleteHandleError handles the Delete error response.
func (client ExpressRouteCrossConnectionPeeringsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified peering for the ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionPeeringsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (ExpressRouteCrossConnectionPeeringResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, options)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCrossConnectionPeeringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteCrossConnectionPeeringResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client ExpressRouteCrossConnectionPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, options *ExpressRouteCrossConnectionPeeringsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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

// getHandleResponse handles the Get response.
func (client ExpressRouteCrossConnectionPeeringsClient) getHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionPeeringResponse, error) {
	var val *ExpressRouteCrossConnectionPeering
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCrossConnectionPeeringResponse{}, err
	}
	return ExpressRouteCrossConnectionPeeringResponse{RawResponse: resp.Response, ExpressRouteCrossConnectionPeering: val}, nil
}

// getHandleError handles the Get error response.
func (client ExpressRouteCrossConnectionPeeringsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all peerings in a specified ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionPeeringsClient) List(resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) ExpressRouteCrossConnectionPeeringListPager {
	return &expressRouteCrossConnectionPeeringListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ExpressRouteCrossConnectionPeeringListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionPeeringList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client ExpressRouteCrossConnectionPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionPeeringsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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

// listHandleResponse handles the List response.
func (client ExpressRouteCrossConnectionPeeringsClient) listHandleResponse(resp *azcore.Response) (ExpressRouteCrossConnectionPeeringListResponse, error) {
	var val *ExpressRouteCrossConnectionPeeringList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCrossConnectionPeeringListResponse{}, err
	}
	return ExpressRouteCrossConnectionPeeringListResponse{RawResponse: resp.Response, ExpressRouteCrossConnectionPeeringList: val}, nil
}

// listHandleError handles the List error response.
func (client ExpressRouteCrossConnectionPeeringsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
