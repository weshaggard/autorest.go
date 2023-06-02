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
	"io"
	"net/http"
	"regexp"
	"strconv"
)

// LibraryServer is a fake server for instances of the azartifacts.LibraryClient type.
type LibraryServer struct {
	// Append is the fake for method LibraryClient.Append
	// HTTP status codes to indicate success: http.StatusCreated
	Append func(ctx context.Context, comp azartifacts.Enum9, libraryName string, content io.ReadSeekCloser, options *azartifacts.LibraryClientAppendOptions) (resp azfake.Responder[azartifacts.LibraryClientAppendResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method LibraryClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, libraryName string, options *azartifacts.LibraryClientBeginCreateOptions) (resp azfake.PollerResponder[azartifacts.LibraryClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LibraryClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusConflict
	BeginDelete func(ctx context.Context, libraryName string, options *azartifacts.LibraryClientBeginDeleteOptions) (resp azfake.PollerResponder[azartifacts.LibraryClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginFlush is the fake for method LibraryClient.BeginFlush
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFlush func(ctx context.Context, libraryName string, options *azartifacts.LibraryClientBeginFlushOptions) (resp azfake.PollerResponder[azartifacts.LibraryClientFlushResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LibraryClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	Get func(ctx context.Context, libraryName string, options *azartifacts.LibraryClientGetOptions) (resp azfake.Responder[azartifacts.LibraryClientGetResponse], errResp azfake.ErrorResponder)

	// GetOperationResult is the fake for method LibraryClient.GetOperationResult
	// HTTP status codes to indicate success:
	//   - http.StatusOK (returns azartifacts.LibraryResource)
	//   - http.StatusAccepted (returns azartifacts.OperationResult)
	GetOperationResult func(ctx context.Context, operationID string, options *azartifacts.LibraryClientGetOperationResultOptions) (resp azfake.Responder[azartifacts.LibraryClientGetOperationResultResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LibraryClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *azartifacts.LibraryClientListOptions) (resp azfake.PagerResponder[azartifacts.LibraryClientListResponse])
}

// NewLibraryServerTransport creates a new instance of LibraryServerTransport with the provided implementation.
// The returned LibraryServerTransport instance is connected to an instance of azartifacts.LibraryClient by way of the
// undefined.Transporter field.
func NewLibraryServerTransport(srv *LibraryServer) *LibraryServerTransport {
	return &LibraryServerTransport{srv: srv}
}

// LibraryServerTransport connects instances of azartifacts.LibraryClient to instances of LibraryServer.
// Don't use this type directly, use NewLibraryServerTransport instead.
type LibraryServerTransport struct {
	srv          *LibraryServer
	beginCreate  *azfake.PollerResponder[azartifacts.LibraryClientCreateResponse]
	beginDelete  *azfake.PollerResponder[azartifacts.LibraryClientDeleteResponse]
	beginFlush   *azfake.PollerResponder[azartifacts.LibraryClientFlushResponse]
	newListPager *azfake.PagerResponder[azartifacts.LibraryClientListResponse]
}

// Do implements the policy.Transporter interface for LibraryServerTransport.
func (l *LibraryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LibraryClient.Append":
		resp, err = l.dispatchAppend(req)
	case "LibraryClient.BeginCreate":
		resp, err = l.dispatchBeginCreate(req)
	case "LibraryClient.BeginDelete":
		resp, err = l.dispatchBeginDelete(req)
	case "LibraryClient.BeginFlush":
		resp, err = l.dispatchBeginFlush(req)
	case "LibraryClient.Get":
		resp, err = l.dispatchGet(req)
	case "LibraryClient.GetOperationResult":
		resp, err = l.dispatchGetOperationResult(req)
	case "LibraryClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LibraryServerTransport) dispatchAppend(req *http.Request) (*http.Response, error) {
	if l.srv.Append == nil {
		return nil, &nonRetriableError{errors.New("method Append not implemented")}
	}
	const regexStr = "/libraries/(?P<libraryName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	blobConditionAppendPositionParam, err := parseOptional(getHeaderValue(req.Header, "x-ms-blob-condition-appendpos"), func(v string) (int64, error) {
		p, parseErr := strconv.ParseInt(v, 10, 64)
		if parseErr != nil {
			return 0, parseErr
		}
		return p, nil
	})
	if err != nil {
		return nil, err
	}
	var options *azartifacts.LibraryClientAppendOptions
	if blobConditionAppendPositionParam != nil {
		options = &azartifacts.LibraryClientAppendOptions{
			BlobConditionAppendPosition: blobConditionAppendPositionParam,
		}
	}
	respr, errRespr := l.srv.Append(req.Context(), azartifacts.Enum9(qp.Get("comp")), matches[regex.SubexpIndex("libraryName")], req.Body.(io.ReadSeekCloser), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LibraryServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreate not implemented")}
	}
	if l.beginCreate == nil {
		const regexStr = "/libraries/(?P<libraryName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := l.srv.BeginCreate(req.Context(), matches[regex.SubexpIndex("libraryName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginCreate = &respr
	}

	resp, err := server.PollerResponderNext(l.beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginCreate) {
		l.beginCreate = nil
	}

	return resp, nil
}

func (l *LibraryServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if l.beginDelete == nil {
		const regexStr = "/libraries/(?P<libraryName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("libraryName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(l.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusConflict}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusConflict", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginDelete) {
		l.beginDelete = nil
	}

	return resp, nil
}

func (l *LibraryServerTransport) dispatchBeginFlush(req *http.Request) (*http.Response, error) {
	if l.srv.BeginFlush == nil {
		return nil, &nonRetriableError{errors.New("method BeginFlush not implemented")}
	}
	if l.beginFlush == nil {
		const regexStr = "/libraries/(?P<libraryName>[a-zA-Z0-9-_]+)/flush"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := l.srv.BeginFlush(req.Context(), matches[regex.SubexpIndex("libraryName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		l.beginFlush = &respr
	}

	resp, err := server.PollerResponderNext(l.beginFlush, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(l.beginFlush) {
		l.beginFlush = nil
	}

	return resp, nil
}

func (l *LibraryServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/libraries/(?P<libraryName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.Get(req.Context(), matches[regex.SubexpIndex("libraryName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LibraryResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LibraryServerTransport) dispatchGetOperationResult(req *http.Request) (*http.Response, error) {
	if l.srv.GetOperationResult == nil {
		return nil, &nonRetriableError{errors.New("method GetOperationResult not implemented")}
	}
	const regexStr = "/libraryOperationResults/(?P<operationId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := l.srv.GetOperationResult(req.Context(), matches[regex.SubexpIndex("operationId")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LibraryServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if l.newListPager == nil {
		resp := l.srv.NewListPager(nil)
		l.newListPager = &resp
		server.PagerResponderInjectNextLinks(l.newListPager, req, func(page *azartifacts.LibraryClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(l.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(l.newListPager) {
		l.newListPager = nil
	}
	return resp, nil
}
