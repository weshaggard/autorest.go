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
	"net/url"
	"strings"
)

type notebookClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *notebookClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateOrUpdateNotebook - Creates or updates a Note Book.
func (client *notebookClient) CreateOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookCreateOrUpdateNotebookOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateNotebookCreateRequest(ctx, notebookName, notebook, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.CreateOrUpdateNotebookHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateNotebookCreateRequest creates the CreateOrUpdateNotebook request.
func (client *notebookClient) CreateOrUpdateNotebookCreateRequest(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookCreateOrUpdateNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(notebook)
}

// CreateOrUpdateNotebookHandleResponse handles the CreateOrUpdateNotebook response.
func (client *notebookClient) CreateOrUpdateNotebookHandleResponse(resp *azcore.Response) (*NotebookResourceResponse, error) {
	result := NotebookResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NotebookResource)
}

// CreateOrUpdateNotebookHandleError handles the CreateOrUpdateNotebook error response.
func (client *notebookClient) CreateOrUpdateNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteNotebook - Deletes a Note book.
func (client *notebookClient) DeleteNotebook(ctx context.Context, notebookName string, options *NotebookDeleteNotebookOptions) (*azcore.Response, error) {
	req, err := client.DeleteNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteNotebookHandleError(resp)
	}
	return resp, nil
}

// DeleteNotebookCreateRequest creates the DeleteNotebook request.
func (client *notebookClient) DeleteNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookDeleteNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// DeleteNotebookHandleError handles the DeleteNotebook error response.
func (client *notebookClient) DeleteNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebook - Gets a Note Book.
func (client *notebookClient) GetNotebook(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (*NotebookResourceResponse, error) {
	req, err := client.GetNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return nil, client.GetNotebookHandleError(resp)
	}
	result, err := client.GetNotebookHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNotebookCreateRequest creates the GetNotebook request.
func (client *notebookClient) GetNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNotebookHandleResponse handles the GetNotebook response.
func (client *notebookClient) GetNotebookHandleResponse(resp *azcore.Response) (*NotebookResourceResponse, error) {
	result := NotebookResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NotebookResource)
}

// GetNotebookHandleError handles the GetNotebook error response.
func (client *notebookClient) GetNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebookSummaryByWorkSpace - Lists a summary of Notebooks.
func (client *notebookClient) GetNotebookSummaryByWorkSpace(options *NotebookGetNotebookSummaryByWorkSpaceOptions) NotebookListResponsePager {
	return &notebookListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetNotebookSummaryByWorkSpaceCreateRequest(ctx, options)
		},
		responder: client.GetNotebookSummaryByWorkSpaceHandleResponse,
		errorer:   client.GetNotebookSummaryByWorkSpaceHandleError,
		advancer: func(ctx context.Context, resp *NotebookListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// GetNotebookSummaryByWorkSpaceCreateRequest creates the GetNotebookSummaryByWorkSpace request.
func (client *notebookClient) GetNotebookSummaryByWorkSpaceCreateRequest(ctx context.Context, options *NotebookGetNotebookSummaryByWorkSpaceOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/summary"
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

// GetNotebookSummaryByWorkSpaceHandleResponse handles the GetNotebookSummaryByWorkSpace response.
func (client *notebookClient) GetNotebookSummaryByWorkSpaceHandleResponse(resp *azcore.Response) (*NotebookListResponseResponse, error) {
	result := NotebookListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NotebookListResponse)
}

// GetNotebookSummaryByWorkSpaceHandleError handles the GetNotebookSummaryByWorkSpace error response.
func (client *notebookClient) GetNotebookSummaryByWorkSpaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebooksByWorkspace - Lists Notebooks.
func (client *notebookClient) GetNotebooksByWorkspace(options *NotebookGetNotebooksByWorkspaceOptions) NotebookListResponsePager {
	return &notebookListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetNotebooksByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.GetNotebooksByWorkspaceHandleResponse,
		errorer:   client.GetNotebooksByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *NotebookListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// GetNotebooksByWorkspaceCreateRequest creates the GetNotebooksByWorkspace request.
func (client *notebookClient) GetNotebooksByWorkspaceCreateRequest(ctx context.Context, options *NotebookGetNotebooksByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/notebooks"
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

// GetNotebooksByWorkspaceHandleResponse handles the GetNotebooksByWorkspace response.
func (client *notebookClient) GetNotebooksByWorkspaceHandleResponse(resp *azcore.Response) (*NotebookListResponseResponse, error) {
	result := NotebookListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NotebookListResponse)
}

// GetNotebooksByWorkspaceHandleError handles the GetNotebooksByWorkspace error response.
func (client *notebookClient) GetNotebooksByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RenameNotebook - Renames a notebook.
func (client *notebookClient) RenameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookRenameNotebookOptions) (*azcore.Response, error) {
	req, err := client.RenameNotebookCreateRequest(ctx, notebookName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.RenameNotebookHandleError(resp)
	}
	return resp, nil
}

// RenameNotebookCreateRequest creates the RenameNotebook request.
func (client *notebookClient) RenameNotebookCreateRequest(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookRenameNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}/rename"
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// RenameNotebookHandleError handles the RenameNotebook error response.
func (client *notebookClient) RenameNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}