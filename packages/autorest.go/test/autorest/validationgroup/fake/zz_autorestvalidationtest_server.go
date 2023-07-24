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
	"generatortests/validationgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// AutoRestValidationTestServer is a fake server for instances of the validationgroup.AutoRestValidationTestClient type.
type AutoRestValidationTestServer struct {
	// GetWithConstantInPath is the fake for method AutoRestValidationTestClient.GetWithConstantInPath
	// HTTP status codes to indicate success: http.StatusOK
	GetWithConstantInPath func(ctx context.Context, options *validationgroup.AutoRestValidationTestClientGetWithConstantInPathOptions) (resp azfake.Responder[validationgroup.AutoRestValidationTestClientGetWithConstantInPathResponse], errResp azfake.ErrorResponder)

	// PostWithConstantInBody is the fake for method AutoRestValidationTestClient.PostWithConstantInBody
	// HTTP status codes to indicate success: http.StatusOK
	PostWithConstantInBody func(ctx context.Context, options *validationgroup.AutoRestValidationTestClientPostWithConstantInBodyOptions) (resp azfake.Responder[validationgroup.AutoRestValidationTestClientPostWithConstantInBodyResponse], errResp azfake.ErrorResponder)

	// ValidationOfBody is the fake for method AutoRestValidationTestClient.ValidationOfBody
	// HTTP status codes to indicate success: http.StatusOK
	ValidationOfBody func(ctx context.Context, resourceGroupName string, id int32, body validationgroup.Product, options *validationgroup.AutoRestValidationTestClientValidationOfBodyOptions) (resp azfake.Responder[validationgroup.AutoRestValidationTestClientValidationOfBodyResponse], errResp azfake.ErrorResponder)

	// ValidationOfMethodParameters is the fake for method AutoRestValidationTestClient.ValidationOfMethodParameters
	// HTTP status codes to indicate success: http.StatusOK
	ValidationOfMethodParameters func(ctx context.Context, resourceGroupName string, id int32, options *validationgroup.AutoRestValidationTestClientValidationOfMethodParametersOptions) (resp azfake.Responder[validationgroup.AutoRestValidationTestClientValidationOfMethodParametersResponse], errResp azfake.ErrorResponder)
}

// NewAutoRestValidationTestServerTransport creates a new instance of AutoRestValidationTestServerTransport with the provided implementation.
// The returned AutoRestValidationTestServerTransport instance is connected to an instance of validationgroup.AutoRestValidationTestClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAutoRestValidationTestServerTransport(srv *AutoRestValidationTestServer) *AutoRestValidationTestServerTransport {
	return &AutoRestValidationTestServerTransport{srv: srv}
}

// AutoRestValidationTestServerTransport connects instances of validationgroup.AutoRestValidationTestClient to instances of AutoRestValidationTestServer.
// Don't use this type directly, use NewAutoRestValidationTestServerTransport instead.
type AutoRestValidationTestServerTransport struct {
	srv *AutoRestValidationTestServer
}

// Do implements the policy.Transporter interface for AutoRestValidationTestServerTransport.
func (a *AutoRestValidationTestServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AutoRestValidationTestClient.GetWithConstantInPath":
		resp, err = a.dispatchGetWithConstantInPath(req)
	case "AutoRestValidationTestClient.PostWithConstantInBody":
		resp, err = a.dispatchPostWithConstantInBody(req)
	case "AutoRestValidationTestClient.ValidationOfBody":
		resp, err = a.dispatchValidationOfBody(req)
	case "AutoRestValidationTestClient.ValidationOfMethodParameters":
		resp, err = a.dispatchValidationOfMethodParameters(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AutoRestValidationTestServerTransport) dispatchGetWithConstantInPath(req *http.Request) (*http.Response, error) {
	if a.srv.GetWithConstantInPath == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetWithConstantInPath not implemented")}
	}
	respr, errRespr := a.srv.GetWithConstantInPath(req.Context(), nil)
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

func (a *AutoRestValidationTestServerTransport) dispatchPostWithConstantInBody(req *http.Request) (*http.Response, error) {
	if a.srv.PostWithConstantInBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method PostWithConstantInBody not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[validationgroup.Product](req)
	if err != nil {
		return nil, err
	}
	var options *validationgroup.AutoRestValidationTestClientPostWithConstantInBodyOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &validationgroup.AutoRestValidationTestClientPostWithConstantInBodyOptions{
			Body: &body,
		}
	}
	respr, errRespr := a.srv.PostWithConstantInBody(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Product, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AutoRestValidationTestServerTransport) dispatchValidationOfBody(req *http.Request) (*http.Response, error) {
	if a.srv.ValidationOfBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidationOfBody not implemented")}
	}
	const regexStr = `/fakepath/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[validationgroup.Product](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.ValidationOfBody(req.Context(), resourceGroupNameUnescaped, int32(idParam), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Product, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AutoRestValidationTestServerTransport) dispatchValidationOfMethodParameters(req *http.Request) (*http.Response, error) {
	if a.srv.ValidationOfMethodParameters == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidationOfMethodParameters not implemented")}
	}
	const regexStr = `/fakepath/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	idUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	idParam, err := parseWithCast(idUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.ValidationOfMethodParameters(req.Context(), resourceGroupNameUnescaped, int32(idParam), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Product, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
