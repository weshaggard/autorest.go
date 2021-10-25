//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineScaleSetExtensionsClient contains the methods for the VirtualMachineScaleSetExtensions group.
// Don't use this type directly, use NewVirtualMachineScaleSetExtensionsClient() instead.
type VirtualMachineScaleSetExtensionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualMachineScaleSetExtensionsClient creates a new instance of VirtualMachineScaleSetExtensionsClient with the specified values.
func NewVirtualMachineScaleSetExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineScaleSetExtensionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VirtualMachineScaleSetExtensionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - The operation to create or update an extension.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsBeginCreateOrUpdateOptions) (VirtualMachineScaleSetExtensionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetExtensionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetExtensionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineScaleSetExtensionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetExtensionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update an extension.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
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
func (client *VirtualMachineScaleSetExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extensionParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineScaleSetExtensionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - The operation to delete the extension.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsBeginDeleteOptions) (VirtualMachineScaleSetExtensionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsDeletePollerResponse{}, err
	}
	result := VirtualMachineScaleSetExtensionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetExtensionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualMachineScaleSetExtensionsDeletePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetExtensionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The operation to delete the extension.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
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
func (client *VirtualMachineScaleSetExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualMachineScaleSetExtensionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - The operation to get the extension.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsGetOptions) (VirtualMachineScaleSetExtensionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineScaleSetExtensionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineScaleSetExtensionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetExtensionsClient) getHandleResponse(resp *http.Response) (VirtualMachineScaleSetExtensionsGetResponse, error) {
	result := VirtualMachineScaleSetExtensionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineScaleSetExtension); err != nil {
		return VirtualMachineScaleSetExtensionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineScaleSetExtensionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Gets a list of all extensions in a VM scale set.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetExtensionsClient) List(resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetExtensionsListOptions) *VirtualMachineScaleSetExtensionsListPager {
	return &VirtualMachineScaleSetExtensionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
		},
		advancer: func(ctx context.Context, resp VirtualMachineScaleSetExtensionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualMachineScaleSetExtensionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetExtensionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetExtensionsClient) listHandleResponse(resp *http.Response) (VirtualMachineScaleSetExtensionsListResponse, error) {
	result := VirtualMachineScaleSetExtensionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineScaleSetExtensionListResult); err != nil {
		return VirtualMachineScaleSetExtensionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineScaleSetExtensionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - The operation to update an extension.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsBeginUpdateOptions) (VirtualMachineScaleSetExtensionsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetExtensionsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetExtensionsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return VirtualMachineScaleSetExtensionsUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetExtensionsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The operation to update an extension.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetExtensionsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extensionParameters)
}

// updateHandleError handles the Update error response.
func (client *VirtualMachineScaleSetExtensionsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
