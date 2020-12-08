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
func (client notebookClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateOrUpdateNotebook - Creates or updates a Note Book.
func (client notebookClient) createOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateNotebookCreateRequest(ctx, notebookName, notebook, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateNotebookHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateNotebookCreateRequest creates the CreateOrUpdateNotebook request.
func (client notebookClient) createOrUpdateNotebookCreateRequest(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (*azcore.Request, error) {
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

// createOrUpdateNotebookHandleResponse handles the CreateOrUpdateNotebook response.
func (client notebookClient) createOrUpdateNotebookHandleResponse(resp *azcore.Response) (NotebookResourceResponse, error) {
	var val *NotebookResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NotebookResourceResponse{}, err
	}
	return NotebookResourceResponse{RawResponse: resp.Response, NotebookResource: val}, nil
}

// createOrUpdateNotebookHandleError handles the CreateOrUpdateNotebook error response.
func (client notebookClient) createOrUpdateNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteNotebook - Deletes a Note book.
func (client notebookClient) deleteNotebook(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (*azcore.Response, error) {
	req, err := client.deleteNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteNotebookHandleError(resp)
	}
	return resp, nil
}

// deleteNotebookCreateRequest creates the DeleteNotebook request.
func (client notebookClient) deleteNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (*azcore.Request, error) {
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

// deleteNotebookHandleError handles the DeleteNotebook error response.
func (client notebookClient) deleteNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebook - Gets a Note Book.
func (client notebookClient) GetNotebook(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (NotebookResourceResponse, error) {
	req, err := client.getNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return NotebookResourceResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return NotebookResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return NotebookResourceResponse{}, client.getNotebookHandleError(resp)
	}
	return client.getNotebookHandleResponse(resp)
}

// getNotebookCreateRequest creates the GetNotebook request.
func (client notebookClient) getNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (*azcore.Request, error) {
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

// getNotebookHandleResponse handles the GetNotebook response.
func (client notebookClient) getNotebookHandleResponse(resp *azcore.Response) (NotebookResourceResponse, error) {
	var val *NotebookResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NotebookResourceResponse{}, err
	}
	return NotebookResourceResponse{RawResponse: resp.Response, NotebookResource: val}, nil
}

// getNotebookHandleError handles the GetNotebook error response.
func (client notebookClient) getNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebookSummaryByWorkSpace - Lists a summary of Notebooks.
func (client notebookClient) GetNotebookSummaryByWorkSpace(options *NotebookGetNotebookSummaryByWorkSpaceOptions) NotebookListResponsePager {
	return &notebookListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getNotebookSummaryByWorkSpaceCreateRequest(ctx, options)
		},
		responder: client.getNotebookSummaryByWorkSpaceHandleResponse,
		errorer:   client.getNotebookSummaryByWorkSpaceHandleError,
		advancer: func(ctx context.Context, resp NotebookListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getNotebookSummaryByWorkSpaceCreateRequest creates the GetNotebookSummaryByWorkSpace request.
func (client notebookClient) getNotebookSummaryByWorkSpaceCreateRequest(ctx context.Context, options *NotebookGetNotebookSummaryByWorkSpaceOptions) (*azcore.Request, error) {
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

// getNotebookSummaryByWorkSpaceHandleResponse handles the GetNotebookSummaryByWorkSpace response.
func (client notebookClient) getNotebookSummaryByWorkSpaceHandleResponse(resp *azcore.Response) (NotebookListResponseResponse, error) {
	var val *NotebookListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NotebookListResponseResponse{}, err
	}
	return NotebookListResponseResponse{RawResponse: resp.Response, NotebookListResponse: val}, nil
}

// getNotebookSummaryByWorkSpaceHandleError handles the GetNotebookSummaryByWorkSpace error response.
func (client notebookClient) getNotebookSummaryByWorkSpaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebooksByWorkspace - Lists Notebooks.
func (client notebookClient) GetNotebooksByWorkspace(options *NotebookGetNotebooksByWorkspaceOptions) NotebookListResponsePager {
	return &notebookListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getNotebooksByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getNotebooksByWorkspaceHandleResponse,
		errorer:   client.getNotebooksByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp NotebookListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getNotebooksByWorkspaceCreateRequest creates the GetNotebooksByWorkspace request.
func (client notebookClient) getNotebooksByWorkspaceCreateRequest(ctx context.Context, options *NotebookGetNotebooksByWorkspaceOptions) (*azcore.Request, error) {
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

// getNotebooksByWorkspaceHandleResponse handles the GetNotebooksByWorkspace response.
func (client notebookClient) getNotebooksByWorkspaceHandleResponse(resp *azcore.Response) (NotebookListResponseResponse, error) {
	var val *NotebookListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NotebookListResponseResponse{}, err
	}
	return NotebookListResponseResponse{RawResponse: resp.Response, NotebookListResponse: val}, nil
}

// getNotebooksByWorkspaceHandleError handles the GetNotebooksByWorkspace error response.
func (client notebookClient) getNotebooksByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RenameNotebook - Renames a notebook.
func (client notebookClient) renameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (*azcore.Response, error) {
	req, err := client.renameNotebookCreateRequest(ctx, notebookName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameNotebookHandleError(resp)
	}
	return resp, nil
}

// renameNotebookCreateRequest creates the RenameNotebook request.
func (client notebookClient) renameNotebookCreateRequest(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (*azcore.Request, error) {
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

// renameNotebookHandleError handles the RenameNotebook error response.
func (client notebookClient) renameNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
