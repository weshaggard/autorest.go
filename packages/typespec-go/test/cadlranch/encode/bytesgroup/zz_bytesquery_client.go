// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bytesgroup

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// BytesQueryClient contains the methods for the Encode.Bytes namespace.
// Don't use this type directly, use [BytesClient.NewBytesQueryClient] instead.
type BytesQueryClient struct {
	internal *azcore.Client
}

// - options - BytesQueryClientBase64Options contains the optional parameters for the BytesQueryClient.Base64 method.
func (client *BytesQueryClient) Base64(ctx context.Context, value []byte, options *BytesQueryClientBase64Options) (BytesQueryClientBase64Response, error) {
	var err error
	const operationName = "BytesQueryClient.Base64"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.base64CreateRequest(ctx, value, options)
	if err != nil {
		return BytesQueryClientBase64Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesQueryClientBase64Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BytesQueryClientBase64Response{}, err
	}
	return BytesQueryClientBase64Response{}, nil
}

// base64CreateRequest creates the Base64 request.
func (client *BytesQueryClient) base64CreateRequest(ctx context.Context, value []byte, options *BytesQueryClientBase64Options) (*policy.Request, error) {
	urlPath := "/encode/bytes/query/base64"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", base64.StdEncoding.EncodeToString(value))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - options - BytesQueryClientBase64URLOptions contains the optional parameters for the BytesQueryClient.Base64URL method.
func (client *BytesQueryClient) Base64URL(ctx context.Context, value []byte, options *BytesQueryClientBase64URLOptions) (BytesQueryClientBase64URLResponse, error) {
	var err error
	const operationName = "BytesQueryClient.Base64URL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.base64URLCreateRequest(ctx, value, options)
	if err != nil {
		return BytesQueryClientBase64URLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesQueryClientBase64URLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BytesQueryClientBase64URLResponse{}, err
	}
	return BytesQueryClientBase64URLResponse{}, nil
}

// base64URLCreateRequest creates the Base64URL request.
func (client *BytesQueryClient) base64URLCreateRequest(ctx context.Context, value []byte, options *BytesQueryClientBase64URLOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/query/base64url"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", base64.RawURLEncoding.EncodeToString(value))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

//   - options - BytesQueryClientBase64URLArrayOptions contains the optional parameters for the BytesQueryClient.Base64URLArray
//     method.
func (client *BytesQueryClient) Base64URLArray(ctx context.Context, value [][]byte, options *BytesQueryClientBase64URLArrayOptions) (BytesQueryClientBase64URLArrayResponse, error) {
	var err error
	const operationName = "BytesQueryClient.Base64URLArray"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.base64URLArrayCreateRequest(ctx, value, options)
	if err != nil {
		return BytesQueryClientBase64URLArrayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesQueryClientBase64URLArrayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BytesQueryClientBase64URLArrayResponse{}, err
	}
	return BytesQueryClientBase64URLArrayResponse{}, nil
}

// base64URLArrayCreateRequest creates the Base64URLArray request.
func (client *BytesQueryClient) base64URLArrayCreateRequest(ctx context.Context, value [][]byte, options *BytesQueryClientBase64URLArrayOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/query/base64url-array"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", strings.Join(func() []string {
		encodedValue := make([]string, len(value))
		for i := 0; i < len(value); i++ {
			encodedValue[i] = base64.RawURLEncoding.EncodeToString(value[i])
		}
		return encodedValue
	}(), ","))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - options - BytesQueryClientDefaultOptions contains the optional parameters for the BytesQueryClient.Default method.
func (client *BytesQueryClient) Default(ctx context.Context, value []byte, options *BytesQueryClientDefaultOptions) (BytesQueryClientDefaultResponse, error) {
	var err error
	const operationName = "BytesQueryClient.Default"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.defaultCreateRequest(ctx, value, options)
	if err != nil {
		return BytesQueryClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BytesQueryClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BytesQueryClientDefaultResponse{}, err
	}
	return BytesQueryClientDefaultResponse{}, nil
}

// defaultCreateRequest creates the Default request.
func (client *BytesQueryClient) defaultCreateRequest(ctx context.Context, value []byte, options *BytesQueryClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/query/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", base64.StdEncoding.EncodeToString(value))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}