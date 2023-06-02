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

// InterfaceTapConfigurationsServer is a fake server for instances of the armnetwork.InterfaceTapConfigurationsClient type.
type InterfaceTapConfigurationsServer struct {
	// BeginCreateOrUpdate is the fake for method InterfaceTapConfigurationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters armnetwork.InterfaceTapConfiguration, options *armnetwork.InterfaceTapConfigurationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method InterfaceTapConfigurationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *armnetwork.InterfaceTapConfigurationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InterfaceTapConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *armnetwork.InterfaceTapConfigurationsClientGetOptions) (resp azfake.Responder[armnetwork.InterfaceTapConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method InterfaceTapConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkInterfaceName string, options *armnetwork.InterfaceTapConfigurationsClientListOptions) (resp azfake.PagerResponder[armnetwork.InterfaceTapConfigurationsClientListResponse])
}

// NewInterfaceTapConfigurationsServerTransport creates a new instance of InterfaceTapConfigurationsServerTransport with the provided implementation.
// The returned InterfaceTapConfigurationsServerTransport instance is connected to an instance of armnetwork.InterfaceTapConfigurationsClient by way of the
// undefined.Transporter field.
func NewInterfaceTapConfigurationsServerTransport(srv *InterfaceTapConfigurationsServer) *InterfaceTapConfigurationsServerTransport {
	return &InterfaceTapConfigurationsServerTransport{srv: srv}
}

// InterfaceTapConfigurationsServerTransport connects instances of armnetwork.InterfaceTapConfigurationsClient to instances of InterfaceTapConfigurationsServer.
// Don't use this type directly, use NewInterfaceTapConfigurationsServerTransport instead.
type InterfaceTapConfigurationsServerTransport struct {
	srv                 *InterfaceTapConfigurationsServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.InterfaceTapConfigurationsClientListResponse]
}

// Do implements the policy.Transporter interface for InterfaceTapConfigurationsServerTransport.
func (i *InterfaceTapConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InterfaceTapConfigurationsClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "InterfaceTapConfigurationsClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "InterfaceTapConfigurationsClient.Get":
		resp, err = i.dispatchGet(req)
	case "InterfaceTapConfigurationsClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InterfaceTapConfigurationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if i.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[a-zA-Z0-9-_]+)/tapConfigurations/(?P<tapConfigurationName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.InterfaceTapConfiguration](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkInterfaceName")], matches[regex.SubexpIndex("tapConfigurationName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		i.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(i.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(i.beginCreateOrUpdate) {
		i.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (i *InterfaceTapConfigurationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if i.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[a-zA-Z0-9-_]+)/tapConfigurations/(?P<tapConfigurationName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkInterfaceName")], matches[regex.SubexpIndex("tapConfigurationName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		i.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(i.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(i.beginDelete) {
		i.beginDelete = nil
	}

	return resp, nil
}

func (i *InterfaceTapConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[a-zA-Z0-9-_]+)/tapConfigurations/(?P<tapConfigurationName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := i.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkInterfaceName")], matches[regex.SubexpIndex("tapConfigurationName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InterfaceTapConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InterfaceTapConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if i.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[a-zA-Z0-9-_]+)/tapConfigurations"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := i.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("networkInterfaceName")], nil)
		i.newListPager = &resp
		server.PagerResponderInjectNextLinks(i.newListPager, req, func(page *armnetwork.InterfaceTapConfigurationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(i.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(i.newListPager) {
		i.newListPager = nil
	}
	return resp, nil
}
