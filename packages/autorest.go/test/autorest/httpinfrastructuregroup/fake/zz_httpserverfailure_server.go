//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/httpinfrastructuregroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// HTTPServerFailureServer is a fake server for instances of the httpinfrastructuregroup.HTTPServerFailureClient type.
type HTTPServerFailureServer struct {
	// Delete505 is the fake for method HTTPServerFailureClient.Delete505
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Delete505 func(ctx context.Context, options *httpinfrastructuregroup.HTTPServerFailureClientDelete505Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPServerFailureClientDelete505Response], errResp azfake.ErrorResponder)

	// Get501 is the fake for method HTTPServerFailureClient.Get501
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Get501 func(ctx context.Context, options *httpinfrastructuregroup.HTTPServerFailureClientGet501Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPServerFailureClientGet501Response], errResp azfake.ErrorResponder)

	// Head501 is the fake for method HTTPServerFailureClient.Head501
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Head501 func(ctx context.Context, options *httpinfrastructuregroup.HTTPServerFailureClientHead501Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPServerFailureClientHead501Response], errResp azfake.ErrorResponder)

	// Post505 is the fake for method HTTPServerFailureClient.Post505
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent
	Post505 func(ctx context.Context, options *httpinfrastructuregroup.HTTPServerFailureClientPost505Options) (resp azfake.Responder[httpinfrastructuregroup.HTTPServerFailureClientPost505Response], errResp azfake.ErrorResponder)
}

// NewHTTPServerFailureServerTransport creates a new instance of HTTPServerFailureServerTransport with the provided implementation.
// The returned HTTPServerFailureServerTransport instance is connected to an instance of httpinfrastructuregroup.HTTPServerFailureClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHTTPServerFailureServerTransport(srv *HTTPServerFailureServer) *HTTPServerFailureServerTransport {
	return &HTTPServerFailureServerTransport{srv: srv}
}

// HTTPServerFailureServerTransport connects instances of httpinfrastructuregroup.HTTPServerFailureClient to instances of HTTPServerFailureServer.
// Don't use this type directly, use NewHTTPServerFailureServerTransport instead.
type HTTPServerFailureServerTransport struct {
	srv *HTTPServerFailureServer
}

// Do implements the policy.Transporter interface for HTTPServerFailureServerTransport.
func (h *HTTPServerFailureServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HTTPServerFailureClient.Delete505":
		resp, err = h.dispatchDelete505(req)
	case "HTTPServerFailureClient.Get501":
		resp, err = h.dispatchGet501(req)
	case "HTTPServerFailureClient.Head501":
		resp, err = h.dispatchHead501(req)
	case "HTTPServerFailureClient.Post505":
		resp, err = h.dispatchPost505(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HTTPServerFailureServerTransport) dispatchDelete505(req *http.Request) (*http.Response, error) {
	if h.srv.Delete505 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete505 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPServerFailureClientDelete505Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPServerFailureClientDelete505Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Delete505(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPServerFailureServerTransport) dispatchGet501(req *http.Request) (*http.Response, error) {
	if h.srv.Get501 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get501 not implemented")}
	}
	respr, errRespr := h.srv.Get501(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPServerFailureServerTransport) dispatchHead501(req *http.Request) (*http.Response, error) {
	if h.srv.Head501 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head501 not implemented")}
	}
	respr, errRespr := h.srv.Head501(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPServerFailureServerTransport) dispatchPost505(req *http.Request) (*http.Response, error) {
	if h.srv.Post505 == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post505 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[bool](req)
	if err != nil {
		return nil, err
	}
	var options *httpinfrastructuregroup.HTTPServerFailureClientPost505Options
	if !reflect.ValueOf(body).IsZero() {
		options = &httpinfrastructuregroup.HTTPServerFailureClientPost505Options{
			BooleanValue: &body,
		}
	}
	respr, errRespr := h.srv.Post505(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
