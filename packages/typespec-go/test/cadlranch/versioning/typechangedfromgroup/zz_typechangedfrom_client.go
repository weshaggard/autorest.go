// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package typechangedfromgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// TypeChangedFromClient - Test for the `@typeChangedFrom` decorator.
// Don't use this type directly, use a constructor function instead.
type TypeChangedFromClient struct {
	internal *azcore.Client
	endpoint string
	version  Versions
}

// Test -
// If the operation fails it returns an *azcore.ResponseError type.
//   - options - TypeChangedFromClientTestOptions contains the optional parameters for the TypeChangedFromClient.Test method.
func (client *TypeChangedFromClient) Test(ctx context.Context, param string, body TestModel, options *TypeChangedFromClientTestOptions) (TypeChangedFromClientTestResponse, error) {
	var err error
	const operationName = "TypeChangedFromClient.Test"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.testCreateRequest(ctx, param, body, options)
	if err != nil {
		return TypeChangedFromClientTestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TypeChangedFromClientTestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TypeChangedFromClientTestResponse{}, err
	}
	resp, err := client.testHandleResponse(httpResp)
	return resp, err
}

// testCreateRequest creates the Test request.
func (client *TypeChangedFromClient) testCreateRequest(ctx context.Context, param string, body TestModel, _ *TypeChangedFromClientTestOptions) (*policy.Request, error) {
	host := "{endpoint}/versioning/type-changed-from/api-version:{version}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{version}", string(client.version))
	urlPath := "/test"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("param", param)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// testHandleResponse handles the Test response.
func (client *TypeChangedFromClient) testHandleResponse(resp *http.Response) (TypeChangedFromClientTestResponse, error) {
	result := TypeChangedFromClientTestResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestModel); err != nil {
		return TypeChangedFromClientTestResponse{}, err
	}
	return result, nil
}
