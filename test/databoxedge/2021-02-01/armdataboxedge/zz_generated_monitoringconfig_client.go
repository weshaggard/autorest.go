//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// MonitoringConfigClient contains the methods for the MonitoringConfig group.
// Don't use this type directly, use NewMonitoringConfigClient() instead.
type MonitoringConfigClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMonitoringConfigClient creates a new instance of MonitoringConfigClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMonitoringConfigClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MonitoringConfigClient, error) {
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
	client := &MonitoringConfigClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new metric configuration or updates an existing one for a role.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// roleName - The role name.
// resourceGroupName - The resource group name.
// monitoringMetricConfiguration - The metric configuration.
// options - MonitoringConfigClientBeginCreateOrUpdateOptions contains the optional parameters for the MonitoringConfigClient.BeginCreateOrUpdate
// method.
func (client *MonitoringConfigClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (*armruntime.Poller[MonitoringConfigClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, roleName, resourceGroupName, monitoringMetricConfiguration, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[MonitoringConfigClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[MonitoringConfigClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a new metric configuration or updates an existing one for a role.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MonitoringConfigClient) createOrUpdate(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, roleName, resourceGroupName, monitoringMetricConfiguration, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MonitoringConfigClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, monitoringMetricConfiguration)
}

// BeginDelete - deletes a new metric configuration for a role.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// roleName - The role name.
// resourceGroupName - The resource group name.
// options - MonitoringConfigClientBeginDeleteOptions contains the optional parameters for the MonitoringConfigClient.BeginDelete
// method.
func (client *MonitoringConfigClient) BeginDelete(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (*armruntime.Poller[MonitoringConfigClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, roleName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[MonitoringConfigClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[MonitoringConfigClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - deletes a new metric configuration for a role.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MonitoringConfigClient) deleteOperation(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
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
func (client *MonitoringConfigClient) deleteCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a metric configuration of a role.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// roleName - The role name.
// resourceGroupName - The resource group name.
// options - MonitoringConfigClientGetOptions contains the optional parameters for the MonitoringConfigClient.Get method.
func (client *MonitoringConfigClient) Get(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientGetOptions) (MonitoringConfigClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
	if err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitoringConfigClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MonitoringConfigClient) getCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MonitoringConfigClient) getHandleResponse(resp *http.Response) (MonitoringConfigClientGetResponse, error) {
	result := MonitoringConfigClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringMetricConfiguration); err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists metric configurations in a role.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// roleName - The role name.
// resourceGroupName - The resource group name.
// options - MonitoringConfigClientListOptions contains the optional parameters for the MonitoringConfigClient.List method.
func (client *MonitoringConfigClient) NewListPager(deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientListOptions) *runtime.Pager[MonitoringConfigClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[MonitoringConfigClientListResponse]{
		More: func(page MonitoringConfigClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MonitoringConfigClientListResponse) (MonitoringConfigClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MonitoringConfigClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MonitoringConfigClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MonitoringConfigClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MonitoringConfigClient) listCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MonitoringConfigClient) listHandleResponse(resp *http.Response) (MonitoringConfigClientListResponse, error) {
	result := MonitoringConfigClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringMetricConfigurationList); err != nil {
		return MonitoringConfigClientListResponse{}, err
	}
	return result, nil
}
