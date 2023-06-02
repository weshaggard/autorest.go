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

// SubnetsServer is a fake server for instances of the armnetwork.SubnetsClient type.
type SubnetsServer struct {
	// BeginCreateOrUpdate is the fake for method SubnetsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters armnetwork.Subnet, options *armnetwork.SubnetsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.SubnetsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SubnetsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *armnetwork.SubnetsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.SubnetsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SubnetsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *armnetwork.SubnetsClientGetOptions) (resp azfake.Responder[armnetwork.SubnetsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SubnetsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualNetworkName string, options *armnetwork.SubnetsClientListOptions) (resp azfake.PagerResponder[armnetwork.SubnetsClientListResponse])

	// BeginPrepareNetworkPolicies is the fake for method SubnetsClient.BeginPrepareNetworkPolicies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPrepareNetworkPolicies func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters armnetwork.PrepareNetworkPoliciesRequest, options *armnetwork.SubnetsClientBeginPrepareNetworkPoliciesOptions) (resp azfake.PollerResponder[armnetwork.SubnetsClientPrepareNetworkPoliciesResponse], errResp azfake.ErrorResponder)

	// BeginUnprepareNetworkPolicies is the fake for method SubnetsClient.BeginUnprepareNetworkPolicies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUnprepareNetworkPolicies func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters armnetwork.UnprepareNetworkPoliciesRequest, options *armnetwork.SubnetsClientBeginUnprepareNetworkPoliciesOptions) (resp azfake.PollerResponder[armnetwork.SubnetsClientUnprepareNetworkPoliciesResponse], errResp azfake.ErrorResponder)
}

// NewSubnetsServerTransport creates a new instance of SubnetsServerTransport with the provided implementation.
// The returned SubnetsServerTransport instance is connected to an instance of armnetwork.SubnetsClient by way of the
// undefined.Transporter field.
func NewSubnetsServerTransport(srv *SubnetsServer) *SubnetsServerTransport {
	return &SubnetsServerTransport{srv: srv}
}

// SubnetsServerTransport connects instances of armnetwork.SubnetsClient to instances of SubnetsServer.
// Don't use this type directly, use NewSubnetsServerTransport instead.
type SubnetsServerTransport struct {
	srv                           *SubnetsServer
	beginCreateOrUpdate           *azfake.PollerResponder[armnetwork.SubnetsClientCreateOrUpdateResponse]
	beginDelete                   *azfake.PollerResponder[armnetwork.SubnetsClientDeleteResponse]
	newListPager                  *azfake.PagerResponder[armnetwork.SubnetsClientListResponse]
	beginPrepareNetworkPolicies   *azfake.PollerResponder[armnetwork.SubnetsClientPrepareNetworkPoliciesResponse]
	beginUnprepareNetworkPolicies *azfake.PollerResponder[armnetwork.SubnetsClientUnprepareNetworkPoliciesResponse]
}

// Do implements the policy.Transporter interface for SubnetsServerTransport.
func (s *SubnetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SubnetsClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "SubnetsClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SubnetsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SubnetsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "SubnetsClient.BeginPrepareNetworkPolicies":
		resp, err = s.dispatchBeginPrepareNetworkPolicies(req)
	case "SubnetsClient.BeginUnprepareNetworkPolicies":
		resp, err = s.dispatchBeginUnprepareNetworkPolicies(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SubnetsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if s.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[a-zA-Z0-9-_]+)/subnets/(?P<subnetName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.Subnet](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualNetworkName")], matches[regex.SubexpIndex("subnetName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(s.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginCreateOrUpdate) {
		s.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (s *SubnetsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if s.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[a-zA-Z0-9-_]+)/subnets/(?P<subnetName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualNetworkName")], matches[regex.SubexpIndex("subnetName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(s.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginDelete) {
		s.beginDelete = nil
	}

	return resp, nil
}

func (s *SubnetsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[a-zA-Z0-9-_]+)/subnets/(?P<subnetName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(qp.Get("$expand"))
	var options *armnetwork.SubnetsClientGetOptions
	if expandParam != nil {
		options = &armnetwork.SubnetsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := s.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualNetworkName")], matches[regex.SubexpIndex("subnetName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Subnet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubnetsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if s.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[a-zA-Z0-9-_]+)/subnets"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualNetworkName")], nil)
		s.newListPager = &resp
		server.PagerResponderInjectNextLinks(s.newListPager, req, func(page *armnetwork.SubnetsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newListPager) {
		s.newListPager = nil
	}
	return resp, nil
}

func (s *SubnetsServerTransport) dispatchBeginPrepareNetworkPolicies(req *http.Request) (*http.Response, error) {
	if s.srv.BeginPrepareNetworkPolicies == nil {
		return nil, &nonRetriableError{errors.New("method BeginPrepareNetworkPolicies not implemented")}
	}
	if s.beginPrepareNetworkPolicies == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[a-zA-Z0-9-_]+)/subnets/(?P<subnetName>[a-zA-Z0-9-_]+)/PrepareNetworkPolicies"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PrepareNetworkPoliciesRequest](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginPrepareNetworkPolicies(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualNetworkName")], matches[regex.SubexpIndex("subnetName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginPrepareNetworkPolicies = &respr
	}

	resp, err := server.PollerResponderNext(s.beginPrepareNetworkPolicies, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginPrepareNetworkPolicies) {
		s.beginPrepareNetworkPolicies = nil
	}

	return resp, nil
}

func (s *SubnetsServerTransport) dispatchBeginUnprepareNetworkPolicies(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUnprepareNetworkPolicies == nil {
		return nil, &nonRetriableError{errors.New("method BeginUnprepareNetworkPolicies not implemented")}
	}
	if s.beginUnprepareNetworkPolicies == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[a-zA-Z0-9-_]+)/subnets/(?P<subnetName>[a-zA-Z0-9-_]+)/UnprepareNetworkPolicies"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.UnprepareNetworkPoliciesRequest](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUnprepareNetworkPolicies(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("virtualNetworkName")], matches[regex.SubexpIndex("subnetName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		s.beginUnprepareNetworkPolicies = &respr
	}

	resp, err := server.PollerResponderNext(s.beginUnprepareNetworkPolicies, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(s.beginUnprepareNetworkPolicies) {
		s.beginUnprepareNetworkPolicies = nil
	}

	return resp, nil
}
