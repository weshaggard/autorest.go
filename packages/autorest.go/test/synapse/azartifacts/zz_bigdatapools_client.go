//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BigDataPoolsClient contains the methods for the BigDataPools group.
// Don't use this type directly, use a constructor function instead.
type BigDataPoolsClient struct {
	internal *azcore.Client
	endpoint string
}

// Get - Get Big Data Pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - bigDataPoolName - The Big Data Pool name
//   - options - BigDataPoolsClientGetOptions contains the optional parameters for the BigDataPoolsClient.Get method.
func (client *BigDataPoolsClient) Get(ctx context.Context, bigDataPoolName string, options *BigDataPoolsClientGetOptions) (BigDataPoolsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, bigDataPoolName, options)
	if err != nil {
		return BigDataPoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BigDataPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BigDataPoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BigDataPoolsClient) getCreateRequest(ctx context.Context, bigDataPoolName string, options *BigDataPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/bigDataPools/{bigDataPoolName}"
	if bigDataPoolName == "" {
		return nil, errors.New("parameter bigDataPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bigDataPoolName}", url.PathEscape(bigDataPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BigDataPoolsClient) getHandleResponse(resp *http.Response) (BigDataPoolsClientGetResponse, error) {
	result := BigDataPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BigDataPoolResourceInfo); err != nil {
		return BigDataPoolsClientGetResponse{}, err
	}
	return result, nil
}

// List - List Big Data Pools
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - options - BigDataPoolsClientListOptions contains the optional parameters for the BigDataPoolsClient.List method.
func (client *BigDataPoolsClient) List(ctx context.Context, options *BigDataPoolsClientListOptions) (BigDataPoolsClientListResponse, error) {
	var err error
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return BigDataPoolsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BigDataPoolsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BigDataPoolsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *BigDataPoolsClient) listCreateRequest(ctx context.Context, options *BigDataPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/bigDataPools"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BigDataPoolsClient) listHandleResponse(resp *http.Response) (BigDataPoolsClientListResponse, error) {
	result := BigDataPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BigDataPoolResourceInfoListResult); err != nil {
		return BigDataPoolsClientListResponse{}, err
	}
	return result, nil
}