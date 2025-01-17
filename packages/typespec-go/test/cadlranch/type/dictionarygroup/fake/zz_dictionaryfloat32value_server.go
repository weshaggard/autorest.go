// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"dictionarygroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DictionaryFloat32ValueServer is a fake server for instances of the dictionarygroup.DictionaryFloat32ValueClient type.
type DictionaryFloat32ValueServer struct {
	// Get is the fake for method DictionaryFloat32ValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *dictionarygroup.DictionaryFloat32ValueClientGetOptions) (resp azfake.Responder[dictionarygroup.DictionaryFloat32ValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method DictionaryFloat32ValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body map[string]*float32, options *dictionarygroup.DictionaryFloat32ValueClientPutOptions) (resp azfake.Responder[dictionarygroup.DictionaryFloat32ValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewDictionaryFloat32ValueServerTransport creates a new instance of DictionaryFloat32ValueServerTransport with the provided implementation.
// The returned DictionaryFloat32ValueServerTransport instance is connected to an instance of dictionarygroup.DictionaryFloat32ValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDictionaryFloat32ValueServerTransport(srv *DictionaryFloat32ValueServer) *DictionaryFloat32ValueServerTransport {
	return &DictionaryFloat32ValueServerTransport{srv: srv}
}

// DictionaryFloat32ValueServerTransport connects instances of dictionarygroup.DictionaryFloat32ValueClient to instances of DictionaryFloat32ValueServer.
// Don't use this type directly, use NewDictionaryFloat32ValueServerTransport instead.
type DictionaryFloat32ValueServerTransport struct {
	srv *DictionaryFloat32ValueServer
}

// Do implements the policy.Transporter interface for DictionaryFloat32ValueServerTransport.
func (d *DictionaryFloat32ValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DictionaryFloat32ValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "DictionaryFloat32ValueClient.Get":
		resp, err = d.dispatchGet(req)
	case "DictionaryFloat32ValueClient.Put":
		resp, err = d.dispatchPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (d *DictionaryFloat32ValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := d.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DictionaryFloat32ValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if d.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[map[string]*float32](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Put(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
