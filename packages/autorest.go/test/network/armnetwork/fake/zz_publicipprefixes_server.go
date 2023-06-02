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

// PublicIPPrefixesServer is a fake server for instances of the armnetwork.PublicIPPrefixesClient type.
type PublicIPPrefixesServer struct {
	// BeginCreateOrUpdate is the fake for method PublicIPPrefixesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters armnetwork.PublicIPPrefix, options *armnetwork.PublicIPPrefixesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.PublicIPPrefixesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PublicIPPrefixesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *armnetwork.PublicIPPrefixesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.PublicIPPrefixesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PublicIPPrefixesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *armnetwork.PublicIPPrefixesClientGetOptions) (resp azfake.Responder[armnetwork.PublicIPPrefixesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PublicIPPrefixesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.PublicIPPrefixesClientListOptions) (resp azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListResponse])

	// NewListAllPager is the fake for method PublicIPPrefixesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.PublicIPPrefixesClientListAllOptions) (resp azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListAllResponse])

	// UpdateTags is the fake for method PublicIPPrefixesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters armnetwork.TagsObject, options *armnetwork.PublicIPPrefixesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.PublicIPPrefixesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewPublicIPPrefixesServerTransport creates a new instance of PublicIPPrefixesServerTransport with the provided implementation.
// The returned PublicIPPrefixesServerTransport instance is connected to an instance of armnetwork.PublicIPPrefixesClient by way of the
// undefined.Transporter field.
func NewPublicIPPrefixesServerTransport(srv *PublicIPPrefixesServer) *PublicIPPrefixesServerTransport {
	return &PublicIPPrefixesServerTransport{srv: srv}
}

// PublicIPPrefixesServerTransport connects instances of armnetwork.PublicIPPrefixesClient to instances of PublicIPPrefixesServer.
// Don't use this type directly, use NewPublicIPPrefixesServerTransport instead.
type PublicIPPrefixesServerTransport struct {
	srv                 *PublicIPPrefixesServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.PublicIPPrefixesClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.PublicIPPrefixesClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListResponse]
	newListAllPager     *azfake.PagerResponder[armnetwork.PublicIPPrefixesClientListAllResponse]
}

// Do implements the policy.Transporter interface for PublicIPPrefixesServerTransport.
func (p *PublicIPPrefixesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PublicIPPrefixesClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PublicIPPrefixesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PublicIPPrefixesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PublicIPPrefixesClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PublicIPPrefixesClient.NewListAllPager":
		resp, err = p.dispatchNewListAllPager(req)
	case "PublicIPPrefixesClient.UpdateTags":
		resp, err = p.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if p.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/publicIPPrefixes/(?P<publicIpPrefixName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PublicIPPrefix](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("publicIpPrefixName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(p.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginCreateOrUpdate) {
		p.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if p.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/publicIPPrefixes/(?P<publicIpPrefixName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("publicIpPrefixName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(p.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginDelete) {
		p.beginDelete = nil
	}

	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/publicIPPrefixes/(?P<publicIpPrefixName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(qp.Get("$expand"))
	var options *armnetwork.PublicIPPrefixesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.PublicIPPrefixesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("publicIpPrefixName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicIPPrefix, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if p.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/publicIPPrefixes"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		p.newListPager = &resp
		server.PagerResponderInjectNextLinks(p.newListPager, req, func(page *armnetwork.PublicIPPrefixesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListPager) {
		p.newListPager = nil
	}
	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListAllPager not implemented")}
	}
	if p.newListAllPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/publicIPPrefixes"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListAllPager(nil)
		p.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(p.newListAllPager, req, func(page *armnetwork.PublicIPPrefixesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListAllPager) {
		p.newListAllPager = nil
	}
	return resp, nil
}

func (p *PublicIPPrefixesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if p.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method UpdateTags not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/publicIPPrefixes/(?P<publicIpPrefixName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("publicIpPrefixName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublicIPPrefix, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
