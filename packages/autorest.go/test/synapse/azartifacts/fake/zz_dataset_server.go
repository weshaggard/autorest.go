//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"azartifacts"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"regexp"
)

// DatasetServer is a fake server for instances of the azartifacts.DatasetClient type.
type DatasetServer struct {
	// BeginCreateOrUpdateDataset is the fake for method DatasetClient.BeginCreateOrUpdateDataset
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdateDataset func(ctx context.Context, datasetName string, dataset azartifacts.DatasetResource, options *azartifacts.DatasetClientBeginCreateOrUpdateDatasetOptions) (resp azfake.PollerResponder[azartifacts.DatasetClientCreateOrUpdateDatasetResponse], errResp azfake.ErrorResponder)

	// BeginDeleteDataset is the fake for method DatasetClient.BeginDeleteDataset
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteDataset func(ctx context.Context, datasetName string, options *azartifacts.DatasetClientBeginDeleteDatasetOptions) (resp azfake.PollerResponder[azartifacts.DatasetClientDeleteDatasetResponse], errResp azfake.ErrorResponder)

	// GetDataset is the fake for method DatasetClient.GetDataset
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	GetDataset func(ctx context.Context, datasetName string, options *azartifacts.DatasetClientGetDatasetOptions) (resp azfake.Responder[azartifacts.DatasetClientGetDatasetResponse], errResp azfake.ErrorResponder)

	// NewGetDatasetsByWorkspacePager is the fake for method DatasetClient.NewGetDatasetsByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetDatasetsByWorkspacePager func(options *azartifacts.DatasetClientGetDatasetsByWorkspaceOptions) (resp azfake.PagerResponder[azartifacts.DatasetClientGetDatasetsByWorkspaceResponse])

	// BeginRenameDataset is the fake for method DatasetClient.BeginRenameDataset
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRenameDataset func(ctx context.Context, datasetName string, request azartifacts.ArtifactRenameRequest, options *azartifacts.DatasetClientBeginRenameDatasetOptions) (resp azfake.PollerResponder[azartifacts.DatasetClientRenameDatasetResponse], errResp azfake.ErrorResponder)
}

// NewDatasetServerTransport creates a new instance of DatasetServerTransport with the provided implementation.
// The returned DatasetServerTransport instance is connected to an instance of azartifacts.DatasetClient by way of the
// undefined.Transporter field.
func NewDatasetServerTransport(srv *DatasetServer) *DatasetServerTransport {
	return &DatasetServerTransport{srv: srv}
}

// DatasetServerTransport connects instances of azartifacts.DatasetClient to instances of DatasetServer.
// Don't use this type directly, use NewDatasetServerTransport instead.
type DatasetServerTransport struct {
	srv                            *DatasetServer
	beginCreateOrUpdateDataset     *azfake.PollerResponder[azartifacts.DatasetClientCreateOrUpdateDatasetResponse]
	beginDeleteDataset             *azfake.PollerResponder[azartifacts.DatasetClientDeleteDatasetResponse]
	newGetDatasetsByWorkspacePager *azfake.PagerResponder[azartifacts.DatasetClientGetDatasetsByWorkspaceResponse]
	beginRenameDataset             *azfake.PollerResponder[azartifacts.DatasetClientRenameDatasetResponse]
}

// Do implements the policy.Transporter interface for DatasetServerTransport.
func (d *DatasetServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DatasetClient.BeginCreateOrUpdateDataset":
		resp, err = d.dispatchBeginCreateOrUpdateDataset(req)
	case "DatasetClient.BeginDeleteDataset":
		resp, err = d.dispatchBeginDeleteDataset(req)
	case "DatasetClient.GetDataset":
		resp, err = d.dispatchGetDataset(req)
	case "DatasetClient.NewGetDatasetsByWorkspacePager":
		resp, err = d.dispatchNewGetDatasetsByWorkspacePager(req)
	case "DatasetClient.BeginRenameDataset":
		resp, err = d.dispatchBeginRenameDataset(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DatasetServerTransport) dispatchBeginCreateOrUpdateDataset(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdateDataset == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdateDataset not implemented")}
	}
	if d.beginCreateOrUpdateDataset == nil {
		const regexStr = "/datasets/(?P<datasetName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.DatasetResource](req)
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *azartifacts.DatasetClientBeginCreateOrUpdateDatasetOptions
		if ifMatchParam != nil {
			options = &azartifacts.DatasetClientBeginCreateOrUpdateDatasetOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := d.srv.BeginCreateOrUpdateDataset(req.Context(), matches[regex.SubexpIndex("datasetName")], body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginCreateOrUpdateDataset = &respr
	}

	resp, err := server.PollerResponderNext(d.beginCreateOrUpdateDataset, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginCreateOrUpdateDataset) {
		d.beginCreateOrUpdateDataset = nil
	}

	return resp, nil
}

func (d *DatasetServerTransport) dispatchBeginDeleteDataset(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDeleteDataset == nil {
		return nil, &nonRetriableError{errors.New("method BeginDeleteDataset not implemented")}
	}
	if d.beginDeleteDataset == nil {
		const regexStr = "/datasets/(?P<datasetName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := d.srv.BeginDeleteDataset(req.Context(), matches[regex.SubexpIndex("datasetName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginDeleteDataset = &respr
	}

	resp, err := server.PollerResponderNext(d.beginDeleteDataset, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginDeleteDataset) {
		d.beginDeleteDataset = nil
	}

	return resp, nil
}

func (d *DatasetServerTransport) dispatchGetDataset(req *http.Request) (*http.Response, error) {
	if d.srv.GetDataset == nil {
		return nil, &nonRetriableError{errors.New("method GetDataset not implemented")}
	}
	const regexStr = "/datasets/(?P<datasetName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *azartifacts.DatasetClientGetDatasetOptions
	if ifNoneMatchParam != nil {
		options = &azartifacts.DatasetClientGetDatasetOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := d.srv.GetDataset(req.Context(), matches[regex.SubexpIndex("datasetName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatasetResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatasetServerTransport) dispatchNewGetDatasetsByWorkspacePager(req *http.Request) (*http.Response, error) {
	if d.srv.NewGetDatasetsByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("method NewGetDatasetsByWorkspacePager not implemented")}
	}
	if d.newGetDatasetsByWorkspacePager == nil {
		resp := d.srv.NewGetDatasetsByWorkspacePager(nil)
		d.newGetDatasetsByWorkspacePager = &resp
		server.PagerResponderInjectNextLinks(d.newGetDatasetsByWorkspacePager, req, func(page *azartifacts.DatasetClientGetDatasetsByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newGetDatasetsByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newGetDatasetsByWorkspacePager) {
		d.newGetDatasetsByWorkspacePager = nil
	}
	return resp, nil
}

func (d *DatasetServerTransport) dispatchBeginRenameDataset(req *http.Request) (*http.Response, error) {
	if d.srv.BeginRenameDataset == nil {
		return nil, &nonRetriableError{errors.New("method BeginRenameDataset not implemented")}
	}
	if d.beginRenameDataset == nil {
		const regexStr = "/datasets/(?P<datasetName>[a-zA-Z0-9-_]+)/rename"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.ArtifactRenameRequest](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginRenameDataset(req.Context(), matches[regex.SubexpIndex("datasetName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		d.beginRenameDataset = &respr
	}

	resp, err := server.PollerResponderNext(d.beginRenameDataset, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(d.beginRenameDataset) {
		d.beginRenameDataset = nil
	}

	return resp, nil
}
