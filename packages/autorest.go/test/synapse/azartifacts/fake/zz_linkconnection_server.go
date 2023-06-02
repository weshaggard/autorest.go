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

// LinkConnectionServer is a fake server for instances of the azartifacts.LinkConnectionClient type.
type LinkConnectionServer struct {
	// CreateOrUpdate is the fake for method LinkConnectionClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, linkConnectionName string, linkConnection azartifacts.LinkConnectionResource, options *azartifacts.LinkConnectionClientCreateOrUpdateOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method LinkConnectionClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, linkConnectionName string, options *azartifacts.LinkConnectionClientDeleteOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientDeleteResponse], errResp azfake.ErrorResponder)

	// EditTables is the fake for method LinkConnectionClient.EditTables
	// HTTP status codes to indicate success: http.StatusOK
	EditTables func(ctx context.Context, linkConnectionName string, editTablesRequest azartifacts.EditTablesRequest, options *azartifacts.LinkConnectionClientEditTablesOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientEditTablesResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LinkConnectionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, linkConnectionName string, options *azartifacts.LinkConnectionClientGetOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientGetResponse], errResp azfake.ErrorResponder)

	// GetDetailedStatus is the fake for method LinkConnectionClient.GetDetailedStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetDetailedStatus func(ctx context.Context, linkConnectionName string, options *azartifacts.LinkConnectionClientGetDetailedStatusOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientGetDetailedStatusResponse], errResp azfake.ErrorResponder)

	// NewListByWorkspacePager is the fake for method LinkConnectionClient.NewListByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWorkspacePager func(options *azartifacts.LinkConnectionClientListByWorkspaceOptions) (resp azfake.PagerResponder[azartifacts.LinkConnectionClientListByWorkspaceResponse])

	// ListLinkTables is the fake for method LinkConnectionClient.ListLinkTables
	// HTTP status codes to indicate success: http.StatusOK
	ListLinkTables func(ctx context.Context, linkConnectionName string, options *azartifacts.LinkConnectionClientListLinkTablesOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientListLinkTablesResponse], errResp azfake.ErrorResponder)

	// Pause is the fake for method LinkConnectionClient.Pause
	// HTTP status codes to indicate success: http.StatusOK
	Pause func(ctx context.Context, linkConnectionName string, options *azartifacts.LinkConnectionClientPauseOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientPauseResponse], errResp azfake.ErrorResponder)

	// QueryTableStatus is the fake for method LinkConnectionClient.QueryTableStatus
	// HTTP status codes to indicate success: http.StatusOK
	QueryTableStatus func(ctx context.Context, linkConnectionName string, queryTableStatusRequest azartifacts.QueryTableStatusRequest, options *azartifacts.LinkConnectionClientQueryTableStatusOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientQueryTableStatusResponse], errResp azfake.ErrorResponder)

	// Resume is the fake for method LinkConnectionClient.Resume
	// HTTP status codes to indicate success: http.StatusOK
	Resume func(ctx context.Context, linkConnectionName string, options *azartifacts.LinkConnectionClientResumeOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientResumeResponse], errResp azfake.ErrorResponder)

	// Start is the fake for method LinkConnectionClient.Start
	// HTTP status codes to indicate success: http.StatusOK
	Start func(ctx context.Context, linkConnectionName string, options *azartifacts.LinkConnectionClientStartOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientStartResponse], errResp azfake.ErrorResponder)

	// Stop is the fake for method LinkConnectionClient.Stop
	// HTTP status codes to indicate success: http.StatusOK
	Stop func(ctx context.Context, linkConnectionName string, options *azartifacts.LinkConnectionClientStopOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientStopResponse], errResp azfake.ErrorResponder)

	// UpdateLandingZoneCredential is the fake for method LinkConnectionClient.UpdateLandingZoneCredential
	// HTTP status codes to indicate success: http.StatusOK
	UpdateLandingZoneCredential func(ctx context.Context, linkConnectionName string, updateLandingZoneCredentialRequest azartifacts.UpdateLandingZoneCredential, options *azartifacts.LinkConnectionClientUpdateLandingZoneCredentialOptions) (resp azfake.Responder[azartifacts.LinkConnectionClientUpdateLandingZoneCredentialResponse], errResp azfake.ErrorResponder)
}

// NewLinkConnectionServerTransport creates a new instance of LinkConnectionServerTransport with the provided implementation.
// The returned LinkConnectionServerTransport instance is connected to an instance of azartifacts.LinkConnectionClient by way of the
// undefined.Transporter field.
func NewLinkConnectionServerTransport(srv *LinkConnectionServer) *LinkConnectionServerTransport {
	return &LinkConnectionServerTransport{srv: srv}
}

// LinkConnectionServerTransport connects instances of azartifacts.LinkConnectionClient to instances of LinkConnectionServer.
// Don't use this type directly, use NewLinkConnectionServerTransport instead.
type LinkConnectionServerTransport struct {
	srv                     *LinkConnectionServer
	newListByWorkspacePager *azfake.PagerResponder[azartifacts.LinkConnectionClientListByWorkspaceResponse]
}

// Do implements the policy.Transporter interface for LinkConnectionServerTransport.
func (l *LinkConnectionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LinkConnectionClient.CreateOrUpdate":
		resp, err = l.dispatchCreateOrUpdate(req)
	case "LinkConnectionClient.Delete":
		resp, err = l.dispatchDelete(req)
	case "LinkConnectionClient.EditTables":
		resp, err = l.dispatchEditTables(req)
	case "LinkConnectionClient.Get":
		resp, err = l.dispatchGet(req)
	case "LinkConnectionClient.GetDetailedStatus":
		resp, err = l.dispatchGetDetailedStatus(req)
	case "LinkConnectionClient.NewListByWorkspacePager":
		resp, err = l.dispatchNewListByWorkspacePager(req)
	case "LinkConnectionClient.ListLinkTables":
		resp, err = l.dispatchListLinkTables(req)
	case "LinkConnectionClient.Pause":
		resp, err = l.dispatchPause(req)
	case "LinkConnectionClient.QueryTableStatus":
		resp, err = l.dispatchQueryTableStatus(req)
	case "LinkConnectionClient.Resume":
		resp, err = l.dispatchResume(req)
	case "LinkConnectionClient.Start":
		resp, err = l.dispatchStart(req)
	case "LinkConnectionClient.Stop":
		resp, err = l.dispatchStop(req)
	case "LinkConnectionClient.UpdateLandingZoneCredential":
		resp, err = l.dispatchUpdateLandingZoneCredential(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LinkConnectionServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method CreateOrUpdate not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azartifacts.LinkConnectionResource](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.CreateOrUpdate(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkConnectionResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkConnectionServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if l.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("method Delete not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.Delete(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkConnectionServerTransport) dispatchEditTables(req *http.Request) (*http.Response, error) {
	if l.srv.EditTables == nil {
		return nil, &nonRetriableError{errors.New("method EditTables not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/edittables"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azartifacts.EditTablesRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.EditTables(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], body, nil)
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

func (l *LinkConnectionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.Get(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkConnectionResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkConnectionServerTransport) dispatchGetDetailedStatus(req *http.Request) (*http.Response, error) {
	if l.srv.GetDetailedStatus == nil {
		return nil, &nonRetriableError{errors.New("method GetDetailedStatus not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/detailedstatus"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.GetDetailedStatus(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkConnectionDetailedStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkConnectionServerTransport) dispatchNewListByWorkspacePager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByWorkspacePager not implemented")}
	}
	if l.newListByWorkspacePager == nil {
		resp := l.srv.NewListByWorkspacePager(nil)
		l.newListByWorkspacePager = &resp
		server.PagerResponderInjectNextLinks(l.newListByWorkspacePager, req, func(page *azartifacts.LinkConnectionClientListByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(l.newListByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(l.newListByWorkspacePager) {
		l.newListByWorkspacePager = nil
	}
	return resp, nil
}

func (l *LinkConnectionServerTransport) dispatchListLinkTables(req *http.Request) (*http.Response, error) {
	if l.srv.ListLinkTables == nil {
		return nil, &nonRetriableError{errors.New("method ListLinkTables not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/linktables"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.ListLinkTables(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkTableListResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkConnectionServerTransport) dispatchPause(req *http.Request) (*http.Response, error) {
	if l.srv.Pause == nil {
		return nil, &nonRetriableError{errors.New("method Pause not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/pause"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.Pause(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], nil)
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

func (l *LinkConnectionServerTransport) dispatchQueryTableStatus(req *http.Request) (*http.Response, error) {
	if l.srv.QueryTableStatus == nil {
		return nil, &nonRetriableError{errors.New("method QueryTableStatus not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/querytablestatus"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azartifacts.QueryTableStatusRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.QueryTableStatus(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkConnectionQueryTableStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkConnectionServerTransport) dispatchResume(req *http.Request) (*http.Response, error) {
	if l.srv.Resume == nil {
		return nil, &nonRetriableError{errors.New("method Resume not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/resume"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.Resume(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], nil)
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

func (l *LinkConnectionServerTransport) dispatchStart(req *http.Request) (*http.Response, error) {
	if l.srv.Start == nil {
		return nil, &nonRetriableError{errors.New("method Start not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/start"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.Start(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], nil)
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

func (l *LinkConnectionServerTransport) dispatchStop(req *http.Request) (*http.Response, error) {
	if l.srv.Stop == nil {
		return nil, &nonRetriableError{errors.New("method Stop not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/stop"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.Stop(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], nil)
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

func (l *LinkConnectionServerTransport) dispatchUpdateLandingZoneCredential(req *http.Request) (*http.Response, error) {
	if l.srv.UpdateLandingZoneCredential == nil {
		return nil, &nonRetriableError{errors.New("method UpdateLandingZoneCredential not implemented")}
	}
	const regexStr = "/linkconnections/(?P<linkConnectionName>[a-zA-Z0-9-_]+)/updateLandingZoneCredential"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azartifacts.UpdateLandingZoneCredential](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.UpdateLandingZoneCredential(req.Context(), matches[regex.SubexpIndex("linkConnectionName")], body, nil)
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
