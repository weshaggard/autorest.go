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

// CustomIPPrefixesServer is a fake server for instances of the armnetwork.CustomIPPrefixesClient type.
type CustomIPPrefixesServer struct {
	// BeginCreateOrUpdate is the fake for method CustomIPPrefixesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters armnetwork.CustomIPPrefix, options *armnetwork.CustomIPPrefixesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.CustomIPPrefixesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CustomIPPrefixesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *armnetwork.CustomIPPrefixesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.CustomIPPrefixesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CustomIPPrefixesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, customIPPrefixName string, options *armnetwork.CustomIPPrefixesClientGetOptions) (resp azfake.Responder[armnetwork.CustomIPPrefixesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CustomIPPrefixesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.CustomIPPrefixesClientListOptions) (resp azfake.PagerResponder[armnetwork.CustomIPPrefixesClientListResponse])

	// NewListAllPager is the fake for method CustomIPPrefixesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.CustomIPPrefixesClientListAllOptions) (resp azfake.PagerResponder[armnetwork.CustomIPPrefixesClientListAllResponse])

	// UpdateTags is the fake for method CustomIPPrefixesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, customIPPrefixName string, parameters armnetwork.TagsObject, options *armnetwork.CustomIPPrefixesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.CustomIPPrefixesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewCustomIPPrefixesServerTransport creates a new instance of CustomIPPrefixesServerTransport with the provided implementation.
// The returned CustomIPPrefixesServerTransport instance is connected to an instance of armnetwork.CustomIPPrefixesClient by way of the
// undefined.Transporter field.
func NewCustomIPPrefixesServerTransport(srv *CustomIPPrefixesServer) *CustomIPPrefixesServerTransport {
	return &CustomIPPrefixesServerTransport{srv: srv}
}

// CustomIPPrefixesServerTransport connects instances of armnetwork.CustomIPPrefixesClient to instances of CustomIPPrefixesServer.
// Don't use this type directly, use NewCustomIPPrefixesServerTransport instead.
type CustomIPPrefixesServerTransport struct {
	srv                 *CustomIPPrefixesServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.CustomIPPrefixesClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.CustomIPPrefixesClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.CustomIPPrefixesClientListResponse]
	newListAllPager     *azfake.PagerResponder[armnetwork.CustomIPPrefixesClientListAllResponse]
}

// Do implements the policy.Transporter interface for CustomIPPrefixesServerTransport.
func (c *CustomIPPrefixesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CustomIPPrefixesClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CustomIPPrefixesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CustomIPPrefixesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CustomIPPrefixesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "CustomIPPrefixesClient.NewListAllPager":
		resp, err = c.dispatchNewListAllPager(req)
	case "CustomIPPrefixesClient.UpdateTags":
		resp, err = c.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CustomIPPrefixesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if c.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/customIpPrefixes/(?P<customIpPrefixName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.CustomIPPrefix](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("customIpPrefixName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(c.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginCreateOrUpdate) {
		c.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (c *CustomIPPrefixesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if c.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/customIpPrefixes/(?P<customIpPrefixName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("customIpPrefixName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(c.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginDelete) {
		c.beginDelete = nil
	}

	return resp, nil
}

func (c *CustomIPPrefixesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/customIpPrefixes/(?P<customIpPrefixName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(qp.Get("$expand"))
	var options *armnetwork.CustomIPPrefixesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.CustomIPPrefixesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := c.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("customIpPrefixName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CustomIPPrefix, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CustomIPPrefixesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if c.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/customIpPrefixes"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		c.newListPager = &resp
		server.PagerResponderInjectNextLinks(c.newListPager, req, func(page *armnetwork.CustomIPPrefixesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(c.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newListPager) {
		c.newListPager = nil
	}
	return resp, nil
}

func (c *CustomIPPrefixesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListAllPager not implemented")}
	}
	if c.newListAllPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/customIpPrefixes"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListAllPager(nil)
		c.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(c.newListAllPager, req, func(page *armnetwork.CustomIPPrefixesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(c.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newListAllPager) {
		c.newListAllPager = nil
	}
	return resp, nil
}

func (c *CustomIPPrefixesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if c.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("method UpdateTags not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/customIpPrefixes/(?P<customIpPrefixName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.UpdateTags(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("customIpPrefixName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CustomIPPrefix, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
