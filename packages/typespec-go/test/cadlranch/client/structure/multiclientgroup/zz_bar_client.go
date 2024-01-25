//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package multiclientgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BarClient contains the methods for the Client.Structure.Service group.
// Don't use this type directly, use a constructor function instead.
type BarClient struct {
	internal *azcore.Client
}

func (client *BarClient) Five(ctx context.Context, options *BarClientFiveOptions) (BarClientFiveResponse, error) {
	var err error
	req, err := client.fiveCreateRequest(ctx, options)
	if err != nil {
		return BarClientFiveResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BarClientFiveResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BarClientFiveResponse{}, err
	}
	return BarClientFiveResponse{}, nil
}

// fiveCreateRequest creates the Five request.
func (client *BarClient) fiveCreateRequest(ctx context.Context, options *BarClientFiveOptions) (*policy.Request, error) {
	urlPath := "/five"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *BarClient) Six(ctx context.Context, options *BarClientSixOptions) (BarClientSixResponse, error) {
	var err error
	req, err := client.sixCreateRequest(ctx, options)
	if err != nil {
		return BarClientSixResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BarClientSixResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BarClientSixResponse{}, err
	}
	return BarClientSixResponse{}, nil
}

// sixCreateRequest creates the Six request.
func (client *BarClient) sixCreateRequest(ctx context.Context, options *BarClientSixOptions) (*policy.Request, error) {
	urlPath := "/six"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}