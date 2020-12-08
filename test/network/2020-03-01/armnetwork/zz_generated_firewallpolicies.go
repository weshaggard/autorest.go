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

// FirewallPoliciesClient contains the methods for the FirewallPolicies group.
// Don't use this type directly, use NewFirewallPoliciesClient() instead.
type FirewallPoliciesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFirewallPoliciesClient creates a new instance of FirewallPoliciesClient with the specified values.
func NewFirewallPoliciesClient(con *armcore.Connection, subscriptionID string) FirewallPoliciesClient {
	return FirewallPoliciesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client FirewallPoliciesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates the specified Firewall Policy.
func (client FirewallPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy, options *FirewallPoliciesBeginCreateOrUpdateOptions) (FirewallPolicyPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, firewallPolicyName, parameters, options)
	if err != nil {
		return FirewallPolicyPollerResponse{}, err
	}
	result := FirewallPolicyPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FirewallPoliciesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return FirewallPolicyPollerResponse{}, err
	}
	poller := &firewallPolicyPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (FirewallPolicyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new FirewallPolicyPoller from the specified resume token.
// token - The value must come from a previous call to FirewallPolicyPoller.ResumeToken().
func (client FirewallPoliciesClient) ResumeCreateOrUpdate(token string) (FirewallPolicyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPoliciesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &firewallPolicyPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified Firewall Policy.
func (client FirewallPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy, options *FirewallPoliciesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, parameters, options)
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
func (client FirewallPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy, options *FirewallPoliciesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client FirewallPoliciesClient) createOrUpdateHandleResponse(resp *azcore.Response) (FirewallPolicyResponse, error) {
	var val *FirewallPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FirewallPolicyResponse{}, err
	}
	return FirewallPolicyResponse{RawResponse: resp.Response, FirewallPolicy: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client FirewallPoliciesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified Firewall Policy.
func (client FirewallPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, firewallPolicyName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FirewallPoliciesClient.Delete", "location", resp, client.deleteHandleError)
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
func (client FirewallPoliciesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPoliciesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified Firewall Policy.
func (client FirewallPoliciesClient) delete(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
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
func (client FirewallPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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
func (client FirewallPoliciesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified Firewall Policy.
func (client FirewallPoliciesClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesGetOptions) (FirewallPolicyResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
	if err != nil {
		return FirewallPolicyResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return FirewallPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FirewallPolicyResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client FirewallPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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

// getHandleResponse handles the Get response.
func (client FirewallPoliciesClient) getHandleResponse(resp *azcore.Response) (FirewallPolicyResponse, error) {
	var val *FirewallPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FirewallPolicyResponse{}, err
	}
	return FirewallPolicyResponse{RawResponse: resp.Response, FirewallPolicy: val}, nil
}

// getHandleError handles the Get error response.
func (client FirewallPoliciesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all Firewall Policies in a resource group.
func (client FirewallPoliciesClient) List(resourceGroupName string, options *FirewallPoliciesListOptions) FirewallPolicyListResultPager {
	return &firewallPolicyListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp FirewallPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client FirewallPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *FirewallPoliciesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies"
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

// listHandleResponse handles the List response.
func (client FirewallPoliciesClient) listHandleResponse(resp *azcore.Response) (FirewallPolicyListResultResponse, error) {
	var val *FirewallPolicyListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FirewallPolicyListResultResponse{}, err
	}
	return FirewallPolicyListResultResponse{RawResponse: resp.Response, FirewallPolicyListResult: val}, nil
}

// listHandleError handles the List error response.
func (client FirewallPoliciesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the Firewall Policies in a subscription.
func (client FirewallPoliciesClient) ListAll(options *FirewallPoliciesListAllOptions) FirewallPolicyListResultPager {
	return &firewallPolicyListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp FirewallPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client FirewallPoliciesClient) listAllCreateRequest(ctx context.Context, options *FirewallPoliciesListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/firewallPolicies"
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

// listAllHandleResponse handles the ListAll response.
func (client FirewallPoliciesClient) listAllHandleResponse(resp *azcore.Response) (FirewallPolicyListResultResponse, error) {
	var val *FirewallPolicyListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FirewallPolicyListResultResponse{}, err
	}
	return FirewallPolicyListResultResponse{RawResponse: resp.Response, FirewallPolicyListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client FirewallPoliciesClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
