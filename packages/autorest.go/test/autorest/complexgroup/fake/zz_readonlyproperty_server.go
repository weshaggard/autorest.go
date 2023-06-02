//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/complexgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ReadonlypropertyServer is a fake server for instances of the complexgroup.ReadonlypropertyClient type.
type ReadonlypropertyServer struct {
	// GetValid is the fake for method ReadonlypropertyClient.GetValid
	// HTTP status codes to indicate success: http.StatusOK
	GetValid func(ctx context.Context, options *complexgroup.ReadonlypropertyClientGetValidOptions) (resp azfake.Responder[complexgroup.ReadonlypropertyClientGetValidResponse], errResp azfake.ErrorResponder)

	// PutValid is the fake for method ReadonlypropertyClient.PutValid
	// HTTP status codes to indicate success: http.StatusOK
	PutValid func(ctx context.Context, complexBody complexgroup.ReadonlyObj, options *complexgroup.ReadonlypropertyClientPutValidOptions) (resp azfake.Responder[complexgroup.ReadonlypropertyClientPutValidResponse], errResp azfake.ErrorResponder)
}

// NewReadonlypropertyServerTransport creates a new instance of ReadonlypropertyServerTransport with the provided implementation.
// The returned ReadonlypropertyServerTransport instance is connected to an instance of complexgroup.ReadonlypropertyClient by way of the
// undefined.Transporter field.
func NewReadonlypropertyServerTransport(srv *ReadonlypropertyServer) *ReadonlypropertyServerTransport {
	return &ReadonlypropertyServerTransport{srv: srv}
}

// ReadonlypropertyServerTransport connects instances of complexgroup.ReadonlypropertyClient to instances of ReadonlypropertyServer.
// Don't use this type directly, use NewReadonlypropertyServerTransport instead.
type ReadonlypropertyServerTransport struct {
	srv *ReadonlypropertyServer
}

// Do implements the policy.Transporter interface for ReadonlypropertyServerTransport.
func (r *ReadonlypropertyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReadonlypropertyClient.GetValid":
		resp, err = r.dispatchGetValid(req)
	case "ReadonlypropertyClient.PutValid":
		resp, err = r.dispatchPutValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReadonlypropertyServerTransport) dispatchGetValid(req *http.Request) (*http.Response, error) {
	if r.srv.GetValid == nil {
		return nil, &nonRetriableError{errors.New("method GetValid not implemented")}
	}
	respr, errRespr := r.srv.GetValid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReadonlyObj, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReadonlypropertyServerTransport) dispatchPutValid(req *http.Request) (*http.Response, error) {
	if r.srv.PutValid == nil {
		return nil, &nonRetriableError{errors.New("method PutValid not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.ReadonlyObj](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.PutValid(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
