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

// SparkConfigurationServer is a fake server for instances of the azartifacts.SparkConfigurationClient type.
type SparkConfigurationServer struct {
	// BeginCreateOrUpdateSparkConfiguration is the fake for method SparkConfigurationClient.BeginCreateOrUpdateSparkConfiguration
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdateSparkConfiguration func(ctx context.Context, sparkConfigurationName string, sparkConfiguration azartifacts.SparkConfigurationResource, options *azartifacts.SparkConfigurationClientBeginCreateOrUpdateSparkConfigurationOptions) (resp azfake.PollerResponder[azartifacts.SparkConfigurationClientCreateOrUpdateSparkConfigurationResponse], errResp azfake.ErrorResponder)

	// BeginDeleteSparkConfiguration is the fake for method SparkConfigurationClient.BeginDeleteSparkConfiguration
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteSparkConfiguration func(ctx context.Context, sparkConfigurationName string, options *azartifacts.SparkConfigurationClientBeginDeleteSparkConfigurationOptions) (resp azfake.PollerResponder[azartifacts.SparkConfigurationClientDeleteSparkConfigurationResponse], errResp azfake.ErrorResponder)

	// GetSparkConfiguration is the fake for method SparkConfigurationClient.GetSparkConfiguration
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	GetSparkConfiguration func(ctx context.Context, sparkConfigurationName string, options *azartifacts.SparkConfigurationClientGetSparkConfigurationOptions) (resp azfake.Responder[azartifacts.SparkConfigurationClientGetSparkConfigurationResponse], errResp azfake.ErrorResponder)

	// NewGetSparkConfigurationsByWorkspacePager is the fake for method SparkConfigurationClient.NewGetSparkConfigurationsByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetSparkConfigurationsByWorkspacePager func(options *azartifacts.SparkConfigurationClientGetSparkConfigurationsByWorkspaceOptions) (resp azfake.PagerResponder[azartifacts.SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse])

	// BeginRenameSparkConfiguration is the fake for method SparkConfigurationClient.BeginRenameSparkConfiguration
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRenameSparkConfiguration func(ctx context.Context, sparkConfigurationName string, request azartifacts.ArtifactRenameRequest, options *azartifacts.SparkConfigurationClientBeginRenameSparkConfigurationOptions) (resp azfake.PollerResponder[azartifacts.SparkConfigurationClientRenameSparkConfigurationResponse], errResp azfake.ErrorResponder)
}

// NewSparkConfigurationServerTransport creates a new instance of SparkConfigurationServerTransport with the provided implementation.
// The returned SparkConfigurationServerTransport instance is connected to an instance of azartifacts.SparkConfigurationClient by way of the
// undefined.Transporter field.
func NewSparkConfigurationServerTransport(srv *SparkConfigurationServer) *SparkConfigurationServerTransport {
	return &SparkConfigurationServerTransport{srv: srv}
}

// SparkConfigurationServerTransport connects instances of azartifacts.SparkConfigurationClient to instances of SparkConfigurationServer.
// Don't use this type directly, use NewSparkConfigurationServerTransport instead.
type SparkConfigurationServerTransport struct {
	srv                                       *SparkConfigurationServer
	beginCreateOrUpdateSparkConfiguration     *azfake.PollerResponder[azartifacts.SparkConfigurationClientCreateOrUpdateSparkConfigurationResponse]
	beginDeleteSparkConfiguration             *azfake.PollerResponder[azartifacts.SparkConfigurationClientDeleteSparkConfigurationResponse]
	newGetSparkConfigurationsByWorkspacePager *azfake.PagerResponder[azartifacts.SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse]
	beginRenameSparkConfiguration             *azfake.PollerResponder[azartifacts.SparkConfigurationClientRenameSparkConfigurationResponse]
}

// Do implements the policy.Transporter interface for SparkConfigurationServerTransport.
func (s *SparkConfigurationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SparkConfigurationClient.BeginCreateOrUpdateSparkConfiguration":
		resp, err = s.dispatchBeginCreateOrUpdateSparkConfiguration(req)
	case "SparkConfigurationClient.BeginDeleteSparkConfiguration":
		resp, err = s.dispatchBeginDeleteSparkConfiguration(req)
	case "SparkConfigurationClient.GetSparkConfiguration":
		resp, err = s.dispatchGetSparkConfiguration(req)
	case "SparkConfigurationClient.NewGetSparkConfigurationsByWorkspacePager":
		resp, err = s.dispatchNewGetSparkConfigurationsByWorkspacePager(req)
	case "SparkConfigurationClient.BeginRenameSparkConfiguration":
		resp, err = s.dispatchBeginRenameSparkConfiguration(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SparkConfigurationServerTransport) dispatchBeginCreateOrUpdateSparkConfiguration(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdateSparkConfiguration == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdateSparkConfiguration not implemented")}
	}
	if s.beginCreateOrUpdateSparkConfiguration == nil {
		const regexStr = "/sparkconfigurations/(?P<sparkConfigurationName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.SparkConfigurationResource](req)
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *azartifacts.SparkConfigurationClientBeginCreateOrUpdateSparkConfigurationOptions
		if ifMatchParam != nil {
			options = &azartifacts.SparkConfigurationClientBeginCreateOrUpdateSparkConfigurationOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := s.srv.BeginCreateOrUpdateSparkConfiguration(req.Context(), matches[regex.SubexpIndex("sparkConfigurationName")], body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginCreateOrUpdateSparkConfiguration = &respr
	}

	resp, err := server.PollerResponderNext(s.beginCreateOrUpdateSparkConfiguration, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginCreateOrUpdateSparkConfiguration) {
		s.beginCreateOrUpdateSparkConfiguration = nil
	}

	return resp, nil
}

func (s *SparkConfigurationServerTransport) dispatchBeginDeleteSparkConfiguration(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDeleteSparkConfiguration == nil {
		return nil, &nonRetriableError{errors.New("method BeginDeleteSparkConfiguration not implemented")}
	}
	if s.beginDeleteSparkConfiguration == nil {
		const regexStr = "/sparkconfigurations/(?P<sparkConfigurationName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := s.srv.BeginDeleteSparkConfiguration(req.Context(), matches[regex.SubexpIndex("sparkConfigurationName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginDeleteSparkConfiguration = &respr
	}

	resp, err := server.PollerResponderNext(s.beginDeleteSparkConfiguration, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginDeleteSparkConfiguration) {
		s.beginDeleteSparkConfiguration = nil
	}

	return resp, nil
}

func (s *SparkConfigurationServerTransport) dispatchGetSparkConfiguration(req *http.Request) (*http.Response, error) {
	if s.srv.GetSparkConfiguration == nil {
		return nil, &nonRetriableError{errors.New("method GetSparkConfiguration not implemented")}
	}
	const regexStr = "/sparkconfigurations/(?P<sparkConfigurationName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *azartifacts.SparkConfigurationClientGetSparkConfigurationOptions
	if ifNoneMatchParam != nil {
		options = &azartifacts.SparkConfigurationClientGetSparkConfigurationOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := s.srv.GetSparkConfiguration(req.Context(), matches[regex.SubexpIndex("sparkConfigurationName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SparkConfigurationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SparkConfigurationServerTransport) dispatchNewGetSparkConfigurationsByWorkspacePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewGetSparkConfigurationsByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("method NewGetSparkConfigurationsByWorkspacePager not implemented")}
	}
	if s.newGetSparkConfigurationsByWorkspacePager == nil {
		resp := s.srv.NewGetSparkConfigurationsByWorkspacePager(nil)
		s.newGetSparkConfigurationsByWorkspacePager = &resp
		server.PagerResponderInjectNextLinks(s.newGetSparkConfigurationsByWorkspacePager, req, func(page *azartifacts.SparkConfigurationClientGetSparkConfigurationsByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newGetSparkConfigurationsByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newGetSparkConfigurationsByWorkspacePager) {
		s.newGetSparkConfigurationsByWorkspacePager = nil
	}
	return resp, nil
}

func (s *SparkConfigurationServerTransport) dispatchBeginRenameSparkConfiguration(req *http.Request) (*http.Response, error) {
	if s.srv.BeginRenameSparkConfiguration == nil {
		return nil, &nonRetriableError{errors.New("method BeginRenameSparkConfiguration not implemented")}
	}
	if s.beginRenameSparkConfiguration == nil {
		const regexStr = "/sparkconfigurations/(?P<sparkConfigurationName>[a-zA-Z0-9-_]+)/rename"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[azartifacts.ArtifactRenameRequest](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginRenameSparkConfiguration(req.Context(), matches[regex.SubexpIndex("sparkConfigurationName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginRenameSparkConfiguration = &respr
	}

	resp, err := server.PollerResponderNext(s.beginRenameSparkConfiguration, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginRenameSparkConfiguration) {
		s.beginRenameSparkConfiguration = nil
	}

	return resp, nil
}
