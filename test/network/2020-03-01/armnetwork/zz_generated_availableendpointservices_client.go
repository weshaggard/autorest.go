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

// AvailableEndpointServicesClient contains the methods for the AvailableEndpointServices group.
// Don't use this type directly, use NewAvailableEndpointServicesClient() instead.
type AvailableEndpointServicesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAvailableEndpointServicesClient creates a new instance of AvailableEndpointServicesClient with the specified values.
func NewAvailableEndpointServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AvailableEndpointServicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AvailableEndpointServicesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - List what values of endpoint services are available for use.
// If the operation fails it returns the *CloudError error type.
func (client *AvailableEndpointServicesClient) List(location string, options *AvailableEndpointServicesListOptions) *AvailableEndpointServicesListPager {
	return &AvailableEndpointServicesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp AvailableEndpointServicesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EndpointServicesListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AvailableEndpointServicesClient) listCreateRequest(ctx context.Context, location string, options *AvailableEndpointServicesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
func (client *AvailableEndpointServicesClient) listHandleResponse(resp *http.Response) (AvailableEndpointServicesListResponse, error) {
	result := AvailableEndpointServicesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointServicesListResult); err != nil {
		return AvailableEndpointServicesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AvailableEndpointServicesClient) listHandleError(resp *http.Response) error {
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
