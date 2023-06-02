//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package nonstringenumgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// IntClient contains the methods for the Int group.
// Don't use this type directly, use a constructor function instead.
type IntClient struct {
	internal *azcore.Client
}

// Get - Get an int enum
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - IntClientGetOptions contains the optional parameters for the IntClient.Get method.
func (client *IntClient) Get(ctx context.Context, options *IntClientGetOptions) (IntClientGetResponse, error) {
	var err error
	const operationName = "IntClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return IntClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntClient) getCreateRequest(ctx context.Context, options *IntClientGetOptions) (*policy.Request, error) {
	urlPath := "/nonStringEnums/int/get"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntClient) getHandleResponse(resp *http.Response) (IntClientGetResponse, error) {
	result := IntClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return IntClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put an int enum
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - input - Input int enum.
//   - options - IntClientPutOptions contains the optional parameters for the IntClient.Put method.
func (client *IntClient) Put(ctx context.Context, input IntEnum, options *IntClientPutOptions) (IntClientPutResponse, error) {
	var err error
	const operationName = "IntClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, input, options)
	if err != nil {
		return IntClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *IntClient) putCreateRequest(ctx context.Context, input IntEnum, options *IntClientPutOptions) (*policy.Request, error) {
	urlPath := "/nonStringEnums/int/put"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *IntClient) putHandleResponse(resp *http.Response) (IntClientPutResponse, error) {
	result := IntClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return IntClientPutResponse{}, err
	}
	return result, nil
}
