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
	"strconv"
)

// GroupsServer is a fake server for instances of the armnetwork.GroupsClient type.
type GroupsServer struct {
	// CreateOrUpdate is the fake for method GroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, parameters armnetwork.Group, options *armnetwork.GroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.GroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, options *armnetwork.GroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.GroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, options *armnetwork.GroupsClientGetOptions) (resp azfake.Responder[armnetwork.GroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method GroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkManagerName string, options *armnetwork.GroupsClientListOptions) (resp azfake.PagerResponder[armnetwork.GroupsClientListResponse])
}

// NewGroupsServerTransport creates a new instance of GroupsServerTransport with the provided implementation.
// The returned GroupsServerTransport instance is connected to an instance of armnetwork.GroupsClient by way of the
// undefined.Transporter field.
func NewGroupsServerTransport(srv *GroupsServer) *GroupsServerTransport {
	return &GroupsServerTransport{srv: srv}
}

// GroupsServerTransport connects instances of armnetwork.GroupsClient to instances of GroupsServer.
// Don't use this type directly, use NewGroupsServerTransport instead.
type GroupsServerTransport struct {
	srv          *GroupsServer
	beginDelete  *azfake.PollerResponder[armnetwork.GroupsClientDeleteResponse]
	newListPager *azfake.PagerResponder[armnetwork.GroupsClientListResponse]
}

// Do implements the policy.Transporter interface for GroupsServerTransport.
func (g *GroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GroupsClient.CreateOrUpdate":
		resp, err = g.dispatchCreateOrUpdate(req)
	case "GroupsClient.BeginDelete":
		resp, err = g.dispatchBeginDelete(req)
	case "GroupsClient.Get":
		resp, err = g.dispatchGet(req)
	case "GroupsClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method CreateOrUpdate not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)/networkGroups/(?P<networkGroupName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.Group](req)
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armnetwork.GroupsClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armnetwork.GroupsClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := g.srv.CreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], matches[regex.SubexpIndex("networkGroupName")], body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Group, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (g *GroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if g.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)/networkGroups/(?P<networkGroupName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		forceParam, err := parseOptional(qp.Get("force"), strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armnetwork.GroupsClientBeginDeleteOptions
		if forceParam != nil {
			options = &armnetwork.GroupsClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], matches[regex.SubexpIndex("networkGroupName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		g.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(g.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(g.beginDelete) {
		g.beginDelete = nil
	}

	return resp, nil
}

func (g *GroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)/networkGroups/(?P<networkGroupName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := g.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], matches[regex.SubexpIndex("networkGroupName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Group, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if g.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)/networkGroups"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		topParam, err := parseOptional(qp.Get("$top"), func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(qp.Get("$skipToken"))
		var options *armnetwork.GroupsClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.GroupsClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := g.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], options)
		g.newListPager = &resp
		server.PagerResponderInjectNextLinks(g.newListPager, req, func(page *armnetwork.GroupsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(g.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(g.newListPager) {
		g.newListPager = nil
	}
	return resp, nil
}
