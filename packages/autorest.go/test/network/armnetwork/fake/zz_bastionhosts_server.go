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

// BastionHostsServer is a fake server for instances of the armnetwork.BastionHostsClient type.
type BastionHostsServer struct {
	// BeginCreateOrUpdate is the fake for method BastionHostsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, bastionHostName string, parameters armnetwork.BastionHost, options *armnetwork.BastionHostsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.BastionHostsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BastionHostsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, bastionHostName string, options *armnetwork.BastionHostsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.BastionHostsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BastionHostsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, bastionHostName string, options *armnetwork.BastionHostsClientGetOptions) (resp azfake.Responder[armnetwork.BastionHostsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BastionHostsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.BastionHostsClientListOptions) (resp azfake.PagerResponder[armnetwork.BastionHostsClientListResponse])

	// NewListByResourceGroupPager is the fake for method BastionHostsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.BastionHostsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.BastionHostsClientListByResourceGroupResponse])

	// BeginUpdateTags is the fake for method BastionHostsClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, bastionHostName string, parameters armnetwork.TagsObject, options *armnetwork.BastionHostsClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armnetwork.BastionHostsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewBastionHostsServerTransport creates a new instance of BastionHostsServerTransport with the provided implementation.
// The returned BastionHostsServerTransport instance is connected to an instance of armnetwork.BastionHostsClient by way of the
// undefined.Transporter field.
func NewBastionHostsServerTransport(srv *BastionHostsServer) *BastionHostsServerTransport {
	return &BastionHostsServerTransport{srv: srv}
}

// BastionHostsServerTransport connects instances of armnetwork.BastionHostsClient to instances of BastionHostsServer.
// Don't use this type directly, use NewBastionHostsServerTransport instead.
type BastionHostsServerTransport struct {
	srv                         *BastionHostsServer
	beginCreateOrUpdate         *azfake.PollerResponder[armnetwork.BastionHostsClientCreateOrUpdateResponse]
	beginDelete                 *azfake.PollerResponder[armnetwork.BastionHostsClientDeleteResponse]
	newListPager                *azfake.PagerResponder[armnetwork.BastionHostsClientListResponse]
	newListByResourceGroupPager *azfake.PagerResponder[armnetwork.BastionHostsClientListByResourceGroupResponse]
	beginUpdateTags             *azfake.PollerResponder[armnetwork.BastionHostsClientUpdateTagsResponse]
}

// Do implements the policy.Transporter interface for BastionHostsServerTransport.
func (b *BastionHostsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BastionHostsClient.BeginCreateOrUpdate":
		resp, err = b.dispatchBeginCreateOrUpdate(req)
	case "BastionHostsClient.BeginDelete":
		resp, err = b.dispatchBeginDelete(req)
	case "BastionHostsClient.Get":
		resp, err = b.dispatchGet(req)
	case "BastionHostsClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	case "BastionHostsClient.NewListByResourceGroupPager":
		resp, err = b.dispatchNewListByResourceGroupPager(req)
	case "BastionHostsClient.BeginUpdateTags":
		resp, err = b.dispatchBeginUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BastionHostsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if b.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/bastionHosts/(?P<bastionHostName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.BastionHost](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("bastionHostName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		b.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(b.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(b.beginCreateOrUpdate) {
		b.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (b *BastionHostsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if b.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/bastionHosts/(?P<bastionHostName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("bastionHostName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		b.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(b.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(b.beginDelete) {
		b.beginDelete = nil
	}

	return resp, nil
}

func (b *BastionHostsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/bastionHosts/(?P<bastionHostName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := b.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("bastionHostName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BastionHost, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BastionHostsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if b.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/bastionHosts"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := b.srv.NewListPager(nil)
		b.newListPager = &resp
		server.PagerResponderInjectNextLinks(b.newListPager, req, func(page *armnetwork.BastionHostsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(b.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(b.newListPager) {
		b.newListPager = nil
	}
	return resp, nil
}

func (b *BastionHostsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if b.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/bastionHosts"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := b.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		b.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(b.newListByResourceGroupPager, req, func(page *armnetwork.BastionHostsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(b.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(b.newListByResourceGroupPager) {
		b.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (b *BastionHostsServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if b.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdateTags not implemented")}
	}
	if b.beginUpdateTags == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/bastionHosts/(?P<bastionHostName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginUpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("bastionHostName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		b.beginUpdateTags = &respr
	}

	resp, err := server.PollerResponderNext(b.beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(b.beginUpdateTags) {
		b.beginUpdateTags = nil
	}

	return resp, nil
}
