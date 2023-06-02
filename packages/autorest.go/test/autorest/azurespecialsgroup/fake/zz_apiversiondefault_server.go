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
	"generatortests/azurespecialsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// APIVersionDefaultServer is a fake server for instances of the azurespecialsgroup.APIVersionDefaultClient type.
type APIVersionDefaultServer struct {
	// GetMethodGlobalNotProvidedValid is the fake for method APIVersionDefaultClient.GetMethodGlobalNotProvidedValid
	// HTTP status codes to indicate success: http.StatusOK
	GetMethodGlobalNotProvidedValid func(ctx context.Context, options *azurespecialsgroup.APIVersionDefaultClientGetMethodGlobalNotProvidedValidOptions) (resp azfake.Responder[azurespecialsgroup.APIVersionDefaultClientGetMethodGlobalNotProvidedValidResponse], errResp azfake.ErrorResponder)

	// GetMethodGlobalValid is the fake for method APIVersionDefaultClient.GetMethodGlobalValid
	// HTTP status codes to indicate success: http.StatusOK
	GetMethodGlobalValid func(ctx context.Context, options *azurespecialsgroup.APIVersionDefaultClientGetMethodGlobalValidOptions) (resp azfake.Responder[azurespecialsgroup.APIVersionDefaultClientGetMethodGlobalValidResponse], errResp azfake.ErrorResponder)

	// GetPathGlobalValid is the fake for method APIVersionDefaultClient.GetPathGlobalValid
	// HTTP status codes to indicate success: http.StatusOK
	GetPathGlobalValid func(ctx context.Context, options *azurespecialsgroup.APIVersionDefaultClientGetPathGlobalValidOptions) (resp azfake.Responder[azurespecialsgroup.APIVersionDefaultClientGetPathGlobalValidResponse], errResp azfake.ErrorResponder)

	// GetSwaggerGlobalValid is the fake for method APIVersionDefaultClient.GetSwaggerGlobalValid
	// HTTP status codes to indicate success: http.StatusOK
	GetSwaggerGlobalValid func(ctx context.Context, options *azurespecialsgroup.APIVersionDefaultClientGetSwaggerGlobalValidOptions) (resp azfake.Responder[azurespecialsgroup.APIVersionDefaultClientGetSwaggerGlobalValidResponse], errResp azfake.ErrorResponder)
}

// NewAPIVersionDefaultServerTransport creates a new instance of APIVersionDefaultServerTransport with the provided implementation.
// The returned APIVersionDefaultServerTransport instance is connected to an instance of azurespecialsgroup.APIVersionDefaultClient by way of the
// undefined.Transporter field.
func NewAPIVersionDefaultServerTransport(srv *APIVersionDefaultServer) *APIVersionDefaultServerTransport {
	return &APIVersionDefaultServerTransport{srv: srv}
}

// APIVersionDefaultServerTransport connects instances of azurespecialsgroup.APIVersionDefaultClient to instances of APIVersionDefaultServer.
// Don't use this type directly, use NewAPIVersionDefaultServerTransport instead.
type APIVersionDefaultServerTransport struct {
	srv *APIVersionDefaultServer
}

// Do implements the policy.Transporter interface for APIVersionDefaultServerTransport.
func (a *APIVersionDefaultServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "APIVersionDefaultClient.GetMethodGlobalNotProvidedValid":
		resp, err = a.dispatchGetMethodGlobalNotProvidedValid(req)
	case "APIVersionDefaultClient.GetMethodGlobalValid":
		resp, err = a.dispatchGetMethodGlobalValid(req)
	case "APIVersionDefaultClient.GetPathGlobalValid":
		resp, err = a.dispatchGetPathGlobalValid(req)
	case "APIVersionDefaultClient.GetSwaggerGlobalValid":
		resp, err = a.dispatchGetSwaggerGlobalValid(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *APIVersionDefaultServerTransport) dispatchGetMethodGlobalNotProvidedValid(req *http.Request) (*http.Response, error) {
	if a.srv.GetMethodGlobalNotProvidedValid == nil {
		return nil, &nonRetriableError{errors.New("method GetMethodGlobalNotProvidedValid not implemented")}
	}
	respr, errRespr := a.srv.GetMethodGlobalNotProvidedValid(req.Context(), nil)
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

func (a *APIVersionDefaultServerTransport) dispatchGetMethodGlobalValid(req *http.Request) (*http.Response, error) {
	if a.srv.GetMethodGlobalValid == nil {
		return nil, &nonRetriableError{errors.New("method GetMethodGlobalValid not implemented")}
	}
	respr, errRespr := a.srv.GetMethodGlobalValid(req.Context(), nil)
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

func (a *APIVersionDefaultServerTransport) dispatchGetPathGlobalValid(req *http.Request) (*http.Response, error) {
	if a.srv.GetPathGlobalValid == nil {
		return nil, &nonRetriableError{errors.New("method GetPathGlobalValid not implemented")}
	}
	respr, errRespr := a.srv.GetPathGlobalValid(req.Context(), nil)
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

func (a *APIVersionDefaultServerTransport) dispatchGetSwaggerGlobalValid(req *http.Request) (*http.Response, error) {
	if a.srv.GetSwaggerGlobalValid == nil {
		return nil, &nonRetriableError{errors.New("method GetSwaggerGlobalValid not implemented")}
	}
	respr, errRespr := a.srv.GetSwaggerGlobalValid(req.Context(), nil)
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
