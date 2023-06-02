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

// ConnectivityConfigurationsServer is a fake server for instances of the armnetwork.ConnectivityConfigurationsClient type.
type ConnectivityConfigurationsServer struct {
	// CreateOrUpdate is the fake for method ConnectivityConfigurationsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, connectivityConfiguration armnetwork.ConnectivityConfiguration, options *armnetwork.ConnectivityConfigurationsClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.ConnectivityConfigurationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConnectivityConfigurationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, options *armnetwork.ConnectivityConfigurationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ConnectivityConfigurationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConnectivityConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, options *armnetwork.ConnectivityConfigurationsClientGetOptions) (resp azfake.Responder[armnetwork.ConnectivityConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ConnectivityConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkManagerName string, options *armnetwork.ConnectivityConfigurationsClientListOptions) (resp azfake.PagerResponder[armnetwork.ConnectivityConfigurationsClientListResponse])
}

// NewConnectivityConfigurationsServerTransport creates a new instance of ConnectivityConfigurationsServerTransport with the provided implementation.
// The returned ConnectivityConfigurationsServerTransport instance is connected to an instance of armnetwork.ConnectivityConfigurationsClient by way of the
// undefined.Transporter field.
func NewConnectivityConfigurationsServerTransport(srv *ConnectivityConfigurationsServer) *ConnectivityConfigurationsServerTransport {
	return &ConnectivityConfigurationsServerTransport{srv: srv}
}

// ConnectivityConfigurationsServerTransport connects instances of armnetwork.ConnectivityConfigurationsClient to instances of ConnectivityConfigurationsServer.
// Don't use this type directly, use NewConnectivityConfigurationsServerTransport instead.
type ConnectivityConfigurationsServerTransport struct {
	srv          *ConnectivityConfigurationsServer
	beginDelete  *azfake.PollerResponder[armnetwork.ConnectivityConfigurationsClientDeleteResponse]
	newListPager *azfake.PagerResponder[armnetwork.ConnectivityConfigurationsClientListResponse]
}

// Do implements the policy.Transporter interface for ConnectivityConfigurationsServerTransport.
func (c *ConnectivityConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConnectivityConfigurationsClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ConnectivityConfigurationsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ConnectivityConfigurationsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConnectivityConfigurationsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConnectivityConfigurationsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method CreateOrUpdate not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)/connectivityConfigurations/(?P<configurationName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.ConnectivityConfiguration](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], matches[regex.SubexpIndex("configurationName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectivityConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectivityConfigurationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if c.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)/connectivityConfigurations/(?P<configurationName>[a-zA-Z0-9-_]+)"
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
		var options *armnetwork.ConnectivityConfigurationsClientBeginDeleteOptions
		if forceParam != nil {
			options = &armnetwork.ConnectivityConfigurationsClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], matches[regex.SubexpIndex("configurationName")], options)
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

func (c *ConnectivityConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)/connectivityConfigurations/(?P<configurationName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := c.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], matches[regex.SubexpIndex("configurationName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectivityConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectivityConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if c.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[a-zA-Z0-9-_]+)/connectivityConfigurations"
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
		var options *armnetwork.ConnectivityConfigurationsClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.ConnectivityConfigurationsClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := c.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkManagerName")], options)
		c.newListPager = &resp
		server.PagerResponderInjectNextLinks(c.newListPager, req, func(page *armnetwork.ConnectivityConfigurationsClientListResponse, createLink func() string) {
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
