//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"armnetwork"
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

// WebCategoriesServer is a fake server for instances of the armnetwork.WebCategoriesClient type.
type WebCategoriesServer struct {
	// Get is the fake for method WebCategoriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, name string, options *armnetwork.WebCategoriesClientGetOptions) (resp azfake.Responder[armnetwork.WebCategoriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySubscriptionPager is the fake for method WebCategoriesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armnetwork.WebCategoriesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armnetwork.WebCategoriesClientListBySubscriptionResponse])
}

// NewWebCategoriesServerTransport creates a new instance of WebCategoriesServerTransport with the provided implementation.
// The returned WebCategoriesServerTransport instance is connected to an instance of armnetwork.WebCategoriesClient by way of the
// undefined.Transporter field.
func NewWebCategoriesServerTransport(srv *WebCategoriesServer) *WebCategoriesServerTransport {
	return &WebCategoriesServerTransport{srv: srv}
}

// WebCategoriesServerTransport connects instances of armnetwork.WebCategoriesClient to instances of WebCategoriesServer.
// Don't use this type directly, use NewWebCategoriesServerTransport instead.
type WebCategoriesServerTransport struct {
	srv                        *WebCategoriesServer
	newListBySubscriptionPager *azfake.PagerResponder[armnetwork.WebCategoriesClientListBySubscriptionResponse]
}

// Do implements the policy.Transporter interface for WebCategoriesServerTransport.
func (w *WebCategoriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WebCategoriesClient.Get":
		resp, err = w.dispatchGet(req)
	case "WebCategoriesClient.NewListBySubscriptionPager":
		resp, err = w.dispatchNewListBySubscriptionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WebCategoriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/azureWebCategories/(?P<name>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(qp.Get("$expand"))
	var options *armnetwork.WebCategoriesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.WebCategoriesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := w.srv.Get(req.Context(), matches[regex.SubexpIndex("name")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureWebCategory, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebCategoriesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListBySubscriptionPager not implemented")}
	}
	if w.newListBySubscriptionPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/azureWebCategories"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := w.srv.NewListBySubscriptionPager(nil)
		w.newListBySubscriptionPager = &resp
		server.PagerResponderInjectNextLinks(w.newListBySubscriptionPager, req, func(page *armnetwork.WebCategoriesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(w.newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(w.newListBySubscriptionPager) {
		w.newListBySubscriptionPager = nil
	}
	return resp, nil
}
