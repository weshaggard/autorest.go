// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ServiceEndpointPoliciesClient contains the methods for the ServiceEndpointPolicies group.
// Don't use this type directly, use NewServiceEndpointPoliciesClient() instead.
type ServiceEndpointPoliciesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewServiceEndpointPoliciesClient creates a new instance of ServiceEndpointPoliciesClient with the specified values.
func NewServiceEndpointPoliciesClient(con *armcore.Connection, subscriptionID string) *ServiceEndpointPoliciesClient {
	return &ServiceEndpointPoliciesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a service Endpoint Policies.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceEndpointPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesBeginCreateOrUpdateOptions) (ServiceEndpointPolicyPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
	if err != nil {
		return ServiceEndpointPolicyPollerResponse{}, err
	}
	result := ServiceEndpointPolicyPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServiceEndpointPoliciesClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ServiceEndpointPolicyPollerResponse{}, err
	}
	poller := &serviceEndpointPolicyPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServiceEndpointPolicyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ServiceEndpointPolicyPoller from the specified resume token.
// token - The value must come from a previous call to ServiceEndpointPolicyPoller.ResumeToken().
func (client *ServiceEndpointPoliciesClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ServiceEndpointPolicyPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServiceEndpointPoliciesClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ServiceEndpointPolicyPollerResponse{}, err
	}
	poller := &serviceEndpointPolicyPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServiceEndpointPolicyPollerResponse{}, err
	}
	result := ServiceEndpointPolicyPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServiceEndpointPolicyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a service Endpoint Policies.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceEndpointPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceEndpointPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters ServiceEndpointPolicy, options *ServiceEndpointPoliciesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServiceEndpointPoliciesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes the specified service endpoint policy.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceEndpointPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serviceEndpointPolicyName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServiceEndpointPoliciesClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ServiceEndpointPoliciesClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServiceEndpointPoliciesClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified service endpoint policy.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceEndpointPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceEndpointPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServiceEndpointPoliciesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the specified service Endpoint Policies in a specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceEndpointPoliciesClient) Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesGetOptions) (ServiceEndpointPolicyResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, options)
	if err != nil {
		return ServiceEndpointPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceEndpointPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceEndpointPolicyResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceEndpointPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceEndpointPoliciesClient) getHandleResponse(resp *azcore.Response) (ServiceEndpointPolicyResponse, error) {
	var val *ServiceEndpointPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceEndpointPolicyResponse{}, err
	}
	return ServiceEndpointPolicyResponse{RawResponse: resp.Response, ServiceEndpointPolicy: val}, nil
}

// getHandleError handles the Get error response.
func (client *ServiceEndpointPoliciesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all the service endpoint policies in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceEndpointPoliciesClient) List(options *ServiceEndpointPoliciesListOptions) ServiceEndpointPolicyListResultPager {
	return &serviceEndpointPolicyListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ServiceEndpointPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServiceEndpointPolicyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ServiceEndpointPoliciesClient) listCreateRequest(ctx context.Context, options *ServiceEndpointPoliciesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ServiceEndpointPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceEndpointPoliciesClient) listHandleResponse(resp *azcore.Response) (ServiceEndpointPolicyListResultResponse, error) {
	var val *ServiceEndpointPolicyListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceEndpointPolicyListResultResponse{}, err
	}
	return ServiceEndpointPolicyListResultResponse{RawResponse: resp.Response, ServiceEndpointPolicyListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ServiceEndpointPoliciesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - Gets all service endpoint Policies in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceEndpointPoliciesClient) ListByResourceGroup(resourceGroupName string, options *ServiceEndpointPoliciesListByResourceGroupOptions) ServiceEndpointPolicyListResultPager {
	return &serviceEndpointPolicyListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ServiceEndpointPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServiceEndpointPolicyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServiceEndpointPoliciesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServiceEndpointPoliciesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServiceEndpointPoliciesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ServiceEndpointPolicyListResultResponse, error) {
	var val *ServiceEndpointPolicyListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceEndpointPolicyListResultResponse{}, err
	}
	return ServiceEndpointPolicyListResultResponse{RawResponse: resp.Response, ServiceEndpointPolicyListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ServiceEndpointPoliciesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UpdateTags - Updates tags of a service endpoint policy.
// If the operation fails it returns the *CloudError error type.
func (client *ServiceEndpointPoliciesClient) UpdateTags(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters TagsObject, options *ServiceEndpointPoliciesUpdateTagsOptions) (ServiceEndpointPolicyResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, parameters, options)
	if err != nil {
		return ServiceEndpointPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceEndpointPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceEndpointPolicyResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ServiceEndpointPoliciesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, parameters TagsObject, options *ServiceEndpointPoliciesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ServiceEndpointPoliciesClient) updateTagsHandleResponse(resp *azcore.Response) (ServiceEndpointPolicyResponse, error) {
	var val *ServiceEndpointPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceEndpointPolicyResponse{}, err
	}
	return ServiceEndpointPolicyResponse{RawResponse: resp.Response, ServiceEndpointPolicy: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ServiceEndpointPoliciesClient) updateTagsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
