// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

type workspaceClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client workspaceClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Get Workspace
func (client workspaceClient) Get(ctx context.Context, options *WorkspaceGetOptions) (WorkspaceResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return WorkspaceResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return WorkspaceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return WorkspaceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client workspaceClient) getCreateRequest(ctx context.Context, options *WorkspaceGetOptions) (*azcore.Request, error) {
	urlPath := "/workspace"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client workspaceClient) getHandleResponse(resp *azcore.Response) (WorkspaceResponse, error) {
	var val *Workspace
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return WorkspaceResponse{}, err
	}
	return WorkspaceResponse{RawResponse: resp.Response, Workspace: val}, nil
}

// getHandleError handles the Get error response.
func (client workspaceClient) getHandleError(resp *azcore.Response) error {
	var err ErrorContract
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
