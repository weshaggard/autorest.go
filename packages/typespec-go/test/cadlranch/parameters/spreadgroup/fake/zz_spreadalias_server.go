// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
	"spreadgroup"
)

// SpreadAliasServer is a fake server for instances of the spreadgroup.SpreadAliasClient type.
type SpreadAliasServer struct {
	// SpreadAsRequestBody is the fake for method SpreadAliasClient.SpreadAsRequestBody
	// HTTP status codes to indicate success: http.StatusNoContent
	SpreadAsRequestBody func(ctx context.Context, name string, options *spreadgroup.SpreadAliasClientSpreadAsRequestBodyOptions) (resp azfake.Responder[spreadgroup.SpreadAliasClientSpreadAsRequestBodyResponse], errResp azfake.ErrorResponder)

	// SpreadAsRequestParameter is the fake for method SpreadAliasClient.SpreadAsRequestParameter
	// HTTP status codes to indicate success: http.StatusNoContent
	SpreadAsRequestParameter func(ctx context.Context, id string, xMSTestHeader string, name string, options *spreadgroup.SpreadAliasClientSpreadAsRequestParameterOptions) (resp azfake.Responder[spreadgroup.SpreadAliasClientSpreadAsRequestParameterResponse], errResp azfake.ErrorResponder)

	// SpreadWithMultipleParameters is the fake for method SpreadAliasClient.SpreadWithMultipleParameters
	// HTTP status codes to indicate success: http.StatusNoContent
	SpreadWithMultipleParameters func(ctx context.Context, id string, xMSTestHeader string, prop1 string, prop2 string, prop3 string, prop4 string, prop5 string, prop6 string, options *spreadgroup.SpreadAliasClientSpreadWithMultipleParametersOptions) (resp azfake.Responder[spreadgroup.SpreadAliasClientSpreadWithMultipleParametersResponse], errResp azfake.ErrorResponder)
}

// NewSpreadAliasServerTransport creates a new instance of SpreadAliasServerTransport with the provided implementation.
// The returned SpreadAliasServerTransport instance is connected to an instance of spreadgroup.SpreadAliasClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSpreadAliasServerTransport(srv *SpreadAliasServer) *SpreadAliasServerTransport {
	return &SpreadAliasServerTransport{srv: srv}
}

// SpreadAliasServerTransport connects instances of spreadgroup.SpreadAliasClient to instances of SpreadAliasServer.
// Don't use this type directly, use NewSpreadAliasServerTransport instead.
type SpreadAliasServerTransport struct {
	srv *SpreadAliasServer
}

// Do implements the policy.Transporter interface for SpreadAliasServerTransport.
func (s *SpreadAliasServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SpreadAliasServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "SpreadAliasClient.SpreadAsRequestBody":
		resp, err = s.dispatchSpreadAsRequestBody(req)
	case "SpreadAliasClient.SpreadAsRequestParameter":
		resp, err = s.dispatchSpreadAsRequestParameter(req)
	case "SpreadAliasClient.SpreadWithMultipleParameters":
		resp, err = s.dispatchSpreadWithMultipleParameters(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (s *SpreadAliasServerTransport) dispatchSpreadAsRequestBody(req *http.Request) (*http.Response, error) {
	if s.srv.SpreadAsRequestBody == nil {
		return nil, &nonRetriableError{errors.New("fake for method SpreadAsRequestBody not implemented")}
	}
	type partialBodyParams struct {
		Name string `json:"name"`
	}
	body, err := server.UnmarshalRequestAsJSON[partialBodyParams](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.SpreadAsRequestBody(req.Context(), body.Name, nil)
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

func (s *SpreadAliasServerTransport) dispatchSpreadAsRequestParameter(req *http.Request) (*http.Response, error) {
	if s.srv.SpreadAsRequestParameter == nil {
		return nil, &nonRetriableError{errors.New("fake for method SpreadAsRequestParameter not implemented")}
	}
	const regexStr = `/parameters/spread/alias/request-parameter/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	type partialBodyParams struct {
		Name string `json:"name"`
	}
	body, err := server.UnmarshalRequestAsJSON[partialBodyParams](req)
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.SpreadAsRequestParameter(req.Context(), idParam, getHeaderValue(req.Header, "x-ms-test-header"), body.Name, nil)
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

func (s *SpreadAliasServerTransport) dispatchSpreadWithMultipleParameters(req *http.Request) (*http.Response, error) {
	if s.srv.SpreadWithMultipleParameters == nil {
		return nil, &nonRetriableError{errors.New("fake for method SpreadWithMultipleParameters not implemented")}
	}
	const regexStr = `/parameters/spread/alias/multiple-parameters/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	type partialBodyParams struct {
		Prop1 string `json:"prop1"`
		Prop2 string `json:"prop2"`
		Prop3 string `json:"prop3"`
		Prop4 string `json:"prop4"`
		Prop5 string `json:"prop5"`
		Prop6 string `json:"prop6"`
	}
	body, err := server.UnmarshalRequestAsJSON[partialBodyParams](req)
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.SpreadWithMultipleParameters(req.Context(), idParam, getHeaderValue(req.Header, "x-ms-test-header"), body.Prop1, body.Prop2, body.Prop3, body.Prop4, body.Prop5, body.Prop6, nil)
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
