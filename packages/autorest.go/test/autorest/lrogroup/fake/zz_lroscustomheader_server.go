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
	"generatortests/lrogroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// LROsCustomHeaderServer is a fake server for instances of the lrogroup.LROsCustomHeaderClient type.
type LROsCustomHeaderServer struct {
	// BeginPost202Retry200 is the fake for method LROsCustomHeaderClient.BeginPost202Retry200
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginPost202Retry200 func(ctx context.Context, options *lrogroup.LROsCustomHeaderClientBeginPost202Retry200Options) (resp azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPost202Retry200Response], errResp azfake.ErrorResponder)

	// BeginPostAsyncRetrySucceeded is the fake for method LROsCustomHeaderClient.BeginPostAsyncRetrySucceeded
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginPostAsyncRetrySucceeded func(ctx context.Context, options *lrogroup.LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPostAsyncRetrySucceededResponse], errResp azfake.ErrorResponder)

	// BeginPut201CreatingSucceeded200 is the fake for method LROsCustomHeaderClient.BeginPut201CreatingSucceeded200
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPut201CreatingSucceeded200 func(ctx context.Context, product lrogroup.Product, options *lrogroup.LROsCustomHeaderClientBeginPut201CreatingSucceeded200Options) (resp azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPut201CreatingSucceeded200Response], errResp azfake.ErrorResponder)

	// BeginPutAsyncRetrySucceeded is the fake for method LROsCustomHeaderClient.BeginPutAsyncRetrySucceeded
	// HTTP status codes to indicate success: http.StatusOK
	BeginPutAsyncRetrySucceeded func(ctx context.Context, product lrogroup.Product, options *lrogroup.LROsCustomHeaderClientBeginPutAsyncRetrySucceededOptions) (resp azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPutAsyncRetrySucceededResponse], errResp azfake.ErrorResponder)
}

// NewLROsCustomHeaderServerTransport creates a new instance of LROsCustomHeaderServerTransport with the provided implementation.
// The returned LROsCustomHeaderServerTransport instance is connected to an instance of lrogroup.LROsCustomHeaderClient by way of the
// undefined.Transporter field.
func NewLROsCustomHeaderServerTransport(srv *LROsCustomHeaderServer) *LROsCustomHeaderServerTransport {
	return &LROsCustomHeaderServerTransport{srv: srv}
}

// LROsCustomHeaderServerTransport connects instances of lrogroup.LROsCustomHeaderClient to instances of LROsCustomHeaderServer.
// Don't use this type directly, use NewLROsCustomHeaderServerTransport instead.
type LROsCustomHeaderServerTransport struct {
	srv                             *LROsCustomHeaderServer
	beginPost202Retry200            *azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPost202Retry200Response]
	beginPostAsyncRetrySucceeded    *azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPostAsyncRetrySucceededResponse]
	beginPut201CreatingSucceeded200 *azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPut201CreatingSucceeded200Response]
	beginPutAsyncRetrySucceeded     *azfake.PollerResponder[lrogroup.LROsCustomHeaderClientPutAsyncRetrySucceededResponse]
}

// Do implements the policy.Transporter interface for LROsCustomHeaderServerTransport.
func (l *LROsCustomHeaderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LROsCustomHeaderClient.BeginPost202Retry200":
		resp, err = l.dispatchBeginPost202Retry200(req)
	case "LROsCustomHeaderClient.BeginPostAsyncRetrySucceeded":
		resp, err = l.dispatchBeginPostAsyncRetrySucceeded(req)
	case "LROsCustomHeaderClient.BeginPut201CreatingSucceeded200":
		resp, err = l.dispatchBeginPut201CreatingSucceeded200(req)
	case "LROsCustomHeaderClient.BeginPutAsyncRetrySucceeded":
		resp, err = l.dispatchBeginPutAsyncRetrySucceeded(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LROsCustomHeaderServerTransport) dispatchBeginPost202Retry200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPost202Retry200 == nil {
		return nil, &nonRetriableError{errors.New("method BeginPost202Retry200 not implemented")}
	}
	if l.beginPost202Retry200 == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		var options *lrogroup.LROsCustomHeaderClientBeginPost202Retry200Options
		if !reflect.ValueOf(body).IsZero() {
			options = &lrogroup.LROsCustomHeaderClientBeginPost202Retry200Options{
				Product: &body,
			}
		}
		respr, errRespr := l.srv.BeginPost202Retry200(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginPost202Retry200 = &respr
	}

	resp, err := server.PollerResponderNext(l.beginPost202Retry200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginPost202Retry200) {
		l.beginPost202Retry200 = nil
	}

	return resp, nil
}

func (l *LROsCustomHeaderServerTransport) dispatchBeginPostAsyncRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPostAsyncRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("method BeginPostAsyncRetrySucceeded not implemented")}
	}
	if l.beginPostAsyncRetrySucceeded == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		var options *lrogroup.LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &lrogroup.LROsCustomHeaderClientBeginPostAsyncRetrySucceededOptions{
				Product: &body,
			}
		}
		respr, errRespr := l.srv.BeginPostAsyncRetrySucceeded(req.Context(), options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginPostAsyncRetrySucceeded = &respr
	}

	resp, err := server.PollerResponderNext(l.beginPostAsyncRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginPostAsyncRetrySucceeded) {
		l.beginPostAsyncRetrySucceeded = nil
	}

	return resp, nil
}

func (l *LROsCustomHeaderServerTransport) dispatchBeginPut201CreatingSucceeded200(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPut201CreatingSucceeded200 == nil {
		return nil, &nonRetriableError{errors.New("method BeginPut201CreatingSucceeded200 not implemented")}
	}
	if l.beginPut201CreatingSucceeded200 == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPut201CreatingSucceeded200(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginPut201CreatingSucceeded200 = &respr
	}

	resp, err := server.PollerResponderNext(l.beginPut201CreatingSucceeded200, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginPut201CreatingSucceeded200) {
		l.beginPut201CreatingSucceeded200 = nil
	}

	return resp, nil
}

func (l *LROsCustomHeaderServerTransport) dispatchBeginPutAsyncRetrySucceeded(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPutAsyncRetrySucceeded == nil {
		return nil, &nonRetriableError{errors.New("method BeginPutAsyncRetrySucceeded not implemented")}
	}
	if l.beginPutAsyncRetrySucceeded == nil {
		body, err := server.UnmarshalRequestAsJSON[lrogroup.Product](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPutAsyncRetrySucceeded(req.Context(), body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginPutAsyncRetrySucceeded = &respr
	}

	resp, err := server.PollerResponderNext(l.beginPutAsyncRetrySucceeded, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginPutAsyncRetrySucceeded) {
		l.beginPutAsyncRetrySucceeded = nil
	}

	return resp, nil
}
