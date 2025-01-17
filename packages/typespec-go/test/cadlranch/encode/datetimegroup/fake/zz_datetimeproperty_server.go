// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"datetimegroup"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DatetimePropertyServer is a fake server for instances of the datetimegroup.DatetimePropertyClient type.
type DatetimePropertyServer struct {
	// Default is the fake for method DatetimePropertyClient.Default
	// HTTP status codes to indicate success: http.StatusOK
	Default func(ctx context.Context, body datetimegroup.DefaultDatetimeProperty, options *datetimegroup.DatetimePropertyClientDefaultOptions) (resp azfake.Responder[datetimegroup.DatetimePropertyClientDefaultResponse], errResp azfake.ErrorResponder)

	// RFC3339 is the fake for method DatetimePropertyClient.RFC3339
	// HTTP status codes to indicate success: http.StatusOK
	RFC3339 func(ctx context.Context, body datetimegroup.RFC3339DatetimeProperty, options *datetimegroup.DatetimePropertyClientRFC3339Options) (resp azfake.Responder[datetimegroup.DatetimePropertyClientRFC3339Response], errResp azfake.ErrorResponder)

	// RFC7231 is the fake for method DatetimePropertyClient.RFC7231
	// HTTP status codes to indicate success: http.StatusOK
	RFC7231 func(ctx context.Context, body datetimegroup.RFC7231DatetimeProperty, options *datetimegroup.DatetimePropertyClientRFC7231Options) (resp azfake.Responder[datetimegroup.DatetimePropertyClientRFC7231Response], errResp azfake.ErrorResponder)

	// UnixTimestamp is the fake for method DatetimePropertyClient.UnixTimestamp
	// HTTP status codes to indicate success: http.StatusOK
	UnixTimestamp func(ctx context.Context, body datetimegroup.UnixTimestampDatetimeProperty, options *datetimegroup.DatetimePropertyClientUnixTimestampOptions) (resp azfake.Responder[datetimegroup.DatetimePropertyClientUnixTimestampResponse], errResp azfake.ErrorResponder)

	// UnixTimestampArray is the fake for method DatetimePropertyClient.UnixTimestampArray
	// HTTP status codes to indicate success: http.StatusOK
	UnixTimestampArray func(ctx context.Context, body datetimegroup.UnixTimestampArrayDatetimeProperty, options *datetimegroup.DatetimePropertyClientUnixTimestampArrayOptions) (resp azfake.Responder[datetimegroup.DatetimePropertyClientUnixTimestampArrayResponse], errResp azfake.ErrorResponder)
}

// NewDatetimePropertyServerTransport creates a new instance of DatetimePropertyServerTransport with the provided implementation.
// The returned DatetimePropertyServerTransport instance is connected to an instance of datetimegroup.DatetimePropertyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatetimePropertyServerTransport(srv *DatetimePropertyServer) *DatetimePropertyServerTransport {
	return &DatetimePropertyServerTransport{srv: srv}
}

// DatetimePropertyServerTransport connects instances of datetimegroup.DatetimePropertyClient to instances of DatetimePropertyServer.
// Don't use this type directly, use NewDatetimePropertyServerTransport instead.
type DatetimePropertyServerTransport struct {
	srv *DatetimePropertyServer
}

// Do implements the policy.Transporter interface for DatetimePropertyServerTransport.
func (d *DatetimePropertyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DatetimePropertyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "DatetimePropertyClient.Default":
		resp, err = d.dispatchDefault(req)
	case "DatetimePropertyClient.RFC3339":
		resp, err = d.dispatchRFC3339(req)
	case "DatetimePropertyClient.RFC7231":
		resp, err = d.dispatchRFC7231(req)
	case "DatetimePropertyClient.UnixTimestamp":
		resp, err = d.dispatchUnixTimestamp(req)
	case "DatetimePropertyClient.UnixTimestampArray":
		resp, err = d.dispatchUnixTimestampArray(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (d *DatetimePropertyServerTransport) dispatchDefault(req *http.Request) (*http.Response, error) {
	if d.srv.Default == nil {
		return nil, &nonRetriableError{errors.New("fake for method Default not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[datetimegroup.DefaultDatetimeProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Default(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DefaultDatetimeProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimePropertyServerTransport) dispatchRFC3339(req *http.Request) (*http.Response, error) {
	if d.srv.RFC3339 == nil {
		return nil, &nonRetriableError{errors.New("fake for method RFC3339 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[datetimegroup.RFC3339DatetimeProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.RFC3339(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RFC3339DatetimeProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimePropertyServerTransport) dispatchRFC7231(req *http.Request) (*http.Response, error) {
	if d.srv.RFC7231 == nil {
		return nil, &nonRetriableError{errors.New("fake for method RFC7231 not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[datetimegroup.RFC7231DatetimeProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.RFC7231(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RFC7231DatetimeProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimePropertyServerTransport) dispatchUnixTimestamp(req *http.Request) (*http.Response, error) {
	if d.srv.UnixTimestamp == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnixTimestamp not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[datetimegroup.UnixTimestampDatetimeProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.UnixTimestamp(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UnixTimestampDatetimeProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatetimePropertyServerTransport) dispatchUnixTimestampArray(req *http.Request) (*http.Response, error) {
	if d.srv.UnixTimestampArray == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnixTimestampArray not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[datetimegroup.UnixTimestampArrayDatetimeProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.UnixTimestampArray(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UnixTimestampArrayDatetimeProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
