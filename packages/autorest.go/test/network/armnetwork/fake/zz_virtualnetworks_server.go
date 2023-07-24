//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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
	"net/url"
	"regexp"
	"strconv"
)

// VirtualNetworksServer is a fake server for instances of the armnetwork.VirtualNetworksClient type.
type VirtualNetworksServer struct {
	// CheckIPAddressAvailability is the fake for method VirtualNetworksClient.CheckIPAddressAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckIPAddressAvailability func(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string, options *armnetwork.VirtualNetworksClientCheckIPAddressAvailabilityOptions) (resp azfake.Responder[armnetwork.VirtualNetworksClientCheckIPAddressAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method VirtualNetworksClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters armnetwork.VirtualNetwork, options *armnetwork.VirtualNetworksClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualNetworksClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *armnetwork.VirtualNetworksClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworksClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualNetworksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *armnetwork.VirtualNetworksClientGetOptions) (resp azfake.Responder[armnetwork.VirtualNetworksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualNetworksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.VirtualNetworksClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualNetworksClientListResponse])

	// NewListAllPager is the fake for method VirtualNetworksClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.VirtualNetworksClientListAllOptions) (resp azfake.PagerResponder[armnetwork.VirtualNetworksClientListAllResponse])

	// BeginListDdosProtectionStatus is the fake for method VirtualNetworksClient.BeginListDdosProtectionStatus
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListDdosProtectionStatus func(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *armnetwork.VirtualNetworksClientBeginListDdosProtectionStatusOptions) (resp azfake.PollerResponder[azfake.PagerResponder[armnetwork.VirtualNetworksClientListDdosProtectionStatusResponse]], errResp azfake.ErrorResponder)

	// NewListUsagePager is the fake for method VirtualNetworksClient.NewListUsagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListUsagePager func(resourceGroupName string, virtualNetworkName string, options *armnetwork.VirtualNetworksClientListUsageOptions) (resp azfake.PagerResponder[armnetwork.VirtualNetworksClientListUsageResponse])

	// UpdateTags is the fake for method VirtualNetworksClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters armnetwork.TagsObject, options *armnetwork.VirtualNetworksClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.VirtualNetworksClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewVirtualNetworksServerTransport creates a new instance of VirtualNetworksServerTransport with the provided implementation.
// The returned VirtualNetworksServerTransport instance is connected to an instance of armnetwork.VirtualNetworksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualNetworksServerTransport(srv *VirtualNetworksServer) *VirtualNetworksServerTransport {
	return &VirtualNetworksServerTransport{
		srv:                           srv,
		beginCreateOrUpdate:           newTracker[azfake.PollerResponder[armnetwork.VirtualNetworksClientCreateOrUpdateResponse]](),
		beginDelete:                   newTracker[azfake.PollerResponder[armnetwork.VirtualNetworksClientDeleteResponse]](),
		newListPager:                  newTracker[azfake.PagerResponder[armnetwork.VirtualNetworksClientListResponse]](),
		newListAllPager:               newTracker[azfake.PagerResponder[armnetwork.VirtualNetworksClientListAllResponse]](),
		beginListDdosProtectionStatus: newTracker[azfake.PollerResponder[azfake.PagerResponder[armnetwork.VirtualNetworksClientListDdosProtectionStatusResponse]]](),
		newListUsagePager:             newTracker[azfake.PagerResponder[armnetwork.VirtualNetworksClientListUsageResponse]](),
	}
}

// VirtualNetworksServerTransport connects instances of armnetwork.VirtualNetworksClient to instances of VirtualNetworksServer.
// Don't use this type directly, use NewVirtualNetworksServerTransport instead.
type VirtualNetworksServerTransport struct {
	srv                           *VirtualNetworksServer
	beginCreateOrUpdate           *tracker[azfake.PollerResponder[armnetwork.VirtualNetworksClientCreateOrUpdateResponse]]
	beginDelete                   *tracker[azfake.PollerResponder[armnetwork.VirtualNetworksClientDeleteResponse]]
	newListPager                  *tracker[azfake.PagerResponder[armnetwork.VirtualNetworksClientListResponse]]
	newListAllPager               *tracker[azfake.PagerResponder[armnetwork.VirtualNetworksClientListAllResponse]]
	beginListDdosProtectionStatus *tracker[azfake.PollerResponder[azfake.PagerResponder[armnetwork.VirtualNetworksClientListDdosProtectionStatusResponse]]]
	newListUsagePager             *tracker[azfake.PagerResponder[armnetwork.VirtualNetworksClientListUsageResponse]]
}

// Do implements the policy.Transporter interface for VirtualNetworksServerTransport.
func (v *VirtualNetworksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualNetworksClient.CheckIPAddressAvailability":
		resp, err = v.dispatchCheckIPAddressAvailability(req)
	case "VirtualNetworksClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualNetworksClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualNetworksClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualNetworksClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualNetworksClient.NewListAllPager":
		resp, err = v.dispatchNewListAllPager(req)
	case "VirtualNetworksClient.BeginListDdosProtectionStatus":
		resp, err = v.dispatchBeginListDdosProtectionStatus(req)
	case "VirtualNetworksClient.NewListUsagePager":
		resp, err = v.dispatchNewListUsagePager(req)
	case "VirtualNetworksClient.UpdateTags":
		resp, err = v.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchCheckIPAddressAvailability(req *http.Request) (*http.Response, error) {
	if v.srv.CheckIPAddressAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckIPAddressAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/CheckIPAddressAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
	if err != nil {
		return nil, err
	}
	ipAddressUnescaped, err := url.QueryUnescape(qp.Get("ipAddress"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.CheckIPAddressAvailability(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, ipAddressUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IPAddressAvailabilityResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualNetwork](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.VirtualNetworksClientGetOptions
	if expandParam != nil {
		options = &armnetwork.VirtualNetworksClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualNetwork, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameUnescaped, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualNetworksClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := v.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListAllPager(nil)
		newListAllPager = &resp
		v.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armnetwork.VirtualNetworksClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		v.newListAllPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchBeginListDdosProtectionStatus(req *http.Request) (*http.Response, error) {
	if v.srv.BeginListDdosProtectionStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListDdosProtectionStatus not implemented")}
	}
	beginListDdosProtectionStatus := v.beginListDdosProtectionStatus.get(req)
	if beginListDdosProtectionStatus == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ddosProtectionStatus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armnetwork.VirtualNetworksClientBeginListDdosProtectionStatusOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.VirtualNetworksClientBeginListDdosProtectionStatusOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		respr, errRespr := v.srv.BeginListDdosProtectionStatus(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginListDdosProtectionStatus = &respr
		v.beginListDdosProtectionStatus.add(req, beginListDdosProtectionStatus)
	}

	resp, err := server.PollerResponderNext(beginListDdosProtectionStatus, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginListDdosProtectionStatus.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginListDdosProtectionStatus) {
		v.beginListDdosProtectionStatus.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchNewListUsagePager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListUsagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListUsagePager not implemented")}
	}
	newListUsagePager := v.newListUsagePager.get(req)
	if newListUsagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListUsagePager(resourceGroupNameUnescaped, virtualNetworkNameUnescaped, nil)
		newListUsagePager = &resp
		v.newListUsagePager.add(req, newListUsagePager)
		server.PagerResponderInjectNextLinks(newListUsagePager, req, func(page *armnetwork.VirtualNetworksClientListUsageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListUsagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListUsagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListUsagePager) {
		v.newListUsagePager.remove(req)
	}
	return resp, nil
}

func (v *VirtualNetworksServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if v.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualNetwork, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}