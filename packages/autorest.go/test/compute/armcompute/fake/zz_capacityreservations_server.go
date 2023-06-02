//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"armcompute"
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

// CapacityReservationsServer is a fake server for instances of the armcompute.CapacityReservationsClient type.
type CapacityReservationsServer struct {
	// BeginCreateOrUpdate is the fake for method CapacityReservationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters armcompute.CapacityReservation, options *armcompute.CapacityReservationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.CapacityReservationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CapacityReservationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, options *armcompute.CapacityReservationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.CapacityReservationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CapacityReservationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, options *armcompute.CapacityReservationsClientGetOptions) (resp azfake.Responder[armcompute.CapacityReservationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByCapacityReservationGroupPager is the fake for method CapacityReservationsClient.NewListByCapacityReservationGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCapacityReservationGroupPager func(resourceGroupName string, capacityReservationGroupName string, options *armcompute.CapacityReservationsClientListByCapacityReservationGroupOptions) (resp azfake.PagerResponder[armcompute.CapacityReservationsClientListByCapacityReservationGroupResponse])

	// BeginUpdate is the fake for method CapacityReservationsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters armcompute.CapacityReservationUpdate, options *armcompute.CapacityReservationsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.CapacityReservationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCapacityReservationsServerTransport creates a new instance of CapacityReservationsServerTransport with the provided implementation.
// The returned CapacityReservationsServerTransport instance is connected to an instance of armcompute.CapacityReservationsClient by way of the
// undefined.Transporter field.
func NewCapacityReservationsServerTransport(srv *CapacityReservationsServer) *CapacityReservationsServerTransport {
	return &CapacityReservationsServerTransport{srv: srv}
}

// CapacityReservationsServerTransport connects instances of armcompute.CapacityReservationsClient to instances of CapacityReservationsServer.
// Don't use this type directly, use NewCapacityReservationsServerTransport instead.
type CapacityReservationsServerTransport struct {
	srv                                    *CapacityReservationsServer
	beginCreateOrUpdate                    *azfake.PollerResponder[armcompute.CapacityReservationsClientCreateOrUpdateResponse]
	beginDelete                            *azfake.PollerResponder[armcompute.CapacityReservationsClientDeleteResponse]
	newListByCapacityReservationGroupPager *azfake.PagerResponder[armcompute.CapacityReservationsClientListByCapacityReservationGroupResponse]
	beginUpdate                            *azfake.PollerResponder[armcompute.CapacityReservationsClientUpdateResponse]
}

// Do implements the policy.Transporter interface for CapacityReservationsServerTransport.
func (c *CapacityReservationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CapacityReservationsClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CapacityReservationsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CapacityReservationsClient.Get":
		resp, err = c.dispatchGet(req)
	case "CapacityReservationsClient.NewListByCapacityReservationGroupPager":
		resp, err = c.dispatchNewListByCapacityReservationGroupPager(req)
	case "CapacityReservationsClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CapacityReservationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if c.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/capacityReservationGroups/(?P<capacityReservationGroupName>[a-zA-Z0-9-_]+)/capacityReservations/(?P<capacityReservationName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.CapacityReservation](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("capacityReservationGroupName")], matches[regex.SubexpIndex("capacityReservationName")], body, nil)
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

func (c *CapacityReservationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if c.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/capacityReservationGroups/(?P<capacityReservationGroupName>[a-zA-Z0-9-_]+)/capacityReservations/(?P<capacityReservationName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("capacityReservationGroupName")], matches[regex.SubexpIndex("capacityReservationName")], nil)
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

func (c *CapacityReservationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/capacityReservationGroups/(?P<capacityReservationGroupName>[a-zA-Z0-9-_]+)/capacityReservations/(?P<capacityReservationName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	expandParam := getOptional(armcompute.CapacityReservationInstanceViewTypes(qp.Get("$expand")))
	var options *armcompute.CapacityReservationsClientGetOptions
	if expandParam != nil {
		options = &armcompute.CapacityReservationsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := c.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("capacityReservationGroupName")], matches[regex.SubexpIndex("capacityReservationName")], options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CapacityReservation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CapacityReservationsServerTransport) dispatchNewListByCapacityReservationGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByCapacityReservationGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByCapacityReservationGroupPager not implemented")}
	}
	if c.newListByCapacityReservationGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/capacityReservationGroups/(?P<capacityReservationGroupName>[a-zA-Z0-9-_]+)/capacityReservations"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListByCapacityReservationGroupPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("capacityReservationGroupName")], nil)
		c.newListByCapacityReservationGroupPager = &resp
		server.PagerResponderInjectNextLinks(c.newListByCapacityReservationGroupPager, req, func(page *armcompute.CapacityReservationsClientListByCapacityReservationGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(c.newListByCapacityReservationGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(c.newListByCapacityReservationGroupPager) {
		c.newListByCapacityReservationGroupPager = nil
	}
	return resp, nil
}

func (c *CapacityReservationsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if c.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/capacityReservationGroups/(?P<capacityReservationGroupName>[a-zA-Z0-9-_]+)/capacityReservations/(?P<capacityReservationName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.CapacityReservationUpdate](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("capacityReservationGroupName")], matches[regex.SubexpIndex("capacityReservationName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(c.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginUpdate) {
		c.beginUpdate = nil
	}

	return resp, nil
}
