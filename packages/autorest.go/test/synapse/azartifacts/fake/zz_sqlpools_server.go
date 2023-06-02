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
	"net/http"
	"regexp"
)

// SQLPoolsServer is a fake server for instances of the azartifacts.SQLPoolsClient type.
type SQLPoolsServer struct {
	// Get is the fake for method SQLPoolsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, sqlPoolName string, options *azartifacts.SQLPoolsClientGetOptions) (resp azfake.Responder[azartifacts.SQLPoolsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method SQLPoolsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, options *azartifacts.SQLPoolsClientListOptions) (resp azfake.Responder[azartifacts.SQLPoolsClientListResponse], errResp azfake.ErrorResponder)
}

// NewSQLPoolsServerTransport creates a new instance of SQLPoolsServerTransport with the provided implementation.
// The returned SQLPoolsServerTransport instance is connected to an instance of azartifacts.SQLPoolsClient by way of the
// undefined.Transporter field.
func NewSQLPoolsServerTransport(srv *SQLPoolsServer) *SQLPoolsServerTransport {
	return &SQLPoolsServerTransport{srv: srv}
}

// SQLPoolsServerTransport connects instances of azartifacts.SQLPoolsClient to instances of SQLPoolsServer.
// Don't use this type directly, use NewSQLPoolsServerTransport instead.
type SQLPoolsServerTransport struct {
	srv *SQLPoolsServer
}

// Do implements the policy.Transporter interface for SQLPoolsServerTransport.
func (s *SQLPoolsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SQLPoolsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SQLPoolsClient.List":
		resp, err = s.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SQLPoolsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/sqlPools/(?P<sqlPoolName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.Get(req.Context(), matches[regex.SubexpIndex("sqlPoolName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLPool, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLPoolsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if s.srv.List == nil {
		return nil, &nonRetriableError{errors.New("method List not implemented")}
	}
	respr, errRespr := s.srv.List(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLPoolInfoListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
