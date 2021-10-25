//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DdosProtectionPlansClient contains the methods for the DdosProtectionPlans group.
// Don't use this type directly, use NewDdosProtectionPlansClient() instead.
type DdosProtectionPlansClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDdosProtectionPlansClient creates a new instance of DdosProtectionPlansClient with the specified values.
func NewDdosProtectionPlansClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DdosProtectionPlansClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DdosProtectionPlansClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates a DDoS protection plan.
// If the operation fails it returns the *CloudError error type.
func (client *DdosProtectionPlansClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansBeginCreateOrUpdateOptions) (DdosProtectionPlansCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return DdosProtectionPlansCreateOrUpdatePollerResponse{}, err
	}
	result := DdosProtectionPlansCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DdosProtectionPlansClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DdosProtectionPlansCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DdosProtectionPlansCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a DDoS protection plan.
// If the operation fails it returns the *CloudError error type.
func (client *DdosProtectionPlansClient) createOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DdosProtectionPlansClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DdosProtectionPlansClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified DDoS protection plan.
// If the operation fails it returns the *CloudError error type.
func (client *DdosProtectionPlansClient) BeginDelete(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansBeginDeleteOptions) (DdosProtectionPlansDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return DdosProtectionPlansDeletePollerResponse{}, err
	}
	result := DdosProtectionPlansDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DdosProtectionPlansClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DdosProtectionPlansDeletePollerResponse{}, err
	}
	result.Poller = &DdosProtectionPlansDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified DDoS protection plan.
// If the operation fails it returns the *CloudError error type.
func (client *DdosProtectionPlansClient) deleteOperation(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DdosProtectionPlansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DdosProtectionPlansClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets information about the specified DDoS protection plan.
// If the operation fails it returns the *CloudError error type.
func (client *DdosProtectionPlansClient) Get(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansGetOptions) (DdosProtectionPlansGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return DdosProtectionPlansGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DdosProtectionPlansGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DdosProtectionPlansGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DdosProtectionPlansClient) getCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *DdosProtectionPlansClient) getHandleResponse(resp *http.Response) (DdosProtectionPlansGetResponse, error) {
	result := DdosProtectionPlansGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlan); err != nil {
		return DdosProtectionPlansGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DdosProtectionPlansClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all DDoS protection plans in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *DdosProtectionPlansClient) List(options *DdosProtectionPlansListOptions) *DdosProtectionPlansListPager {
	return &DdosProtectionPlansListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DdosProtectionPlansListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DdosProtectionPlanListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DdosProtectionPlansClient) listCreateRequest(ctx context.Context, options *DdosProtectionPlansListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *DdosProtectionPlansClient) listHandleResponse(resp *http.Response) (DdosProtectionPlansListResponse, error) {
	result := DdosProtectionPlansListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlanListResult); err != nil {
		return DdosProtectionPlansListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DdosProtectionPlansClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets all the DDoS protection plans in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *DdosProtectionPlansClient) ListByResourceGroup(resourceGroupName string, options *DdosProtectionPlansListByResourceGroupOptions) *DdosProtectionPlansListByResourceGroupPager {
	return &DdosProtectionPlansListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DdosProtectionPlansListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DdosProtectionPlanListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DdosProtectionPlansClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DdosProtectionPlansListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DdosProtectionPlansClient) listByResourceGroupHandleResponse(resp *http.Response) (DdosProtectionPlansListByResourceGroupResponse, error) {
	result := DdosProtectionPlansListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlanListResult); err != nil {
		return DdosProtectionPlansListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DdosProtectionPlansClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateTags - Update a DDoS protection plan tags.
// If the operation fails it returns the *CloudError error type.
func (client *DdosProtectionPlansClient) UpdateTags(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansUpdateTagsOptions) (DdosProtectionPlansUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return DdosProtectionPlansUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DdosProtectionPlansUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DdosProtectionPlansUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *DdosProtectionPlansClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *DdosProtectionPlansClient) updateTagsHandleResponse(resp *http.Response) (DdosProtectionPlansUpdateTagsResponse, error) {
	result := DdosProtectionPlansUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlan); err != nil {
		return DdosProtectionPlansUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *DdosProtectionPlansClient) updateTagsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
