// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armlargeinstance"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// AzureLargeInstancesServer is a fake server for instances of the armlargeinstance.AzureLargeInstancesClient type.
type AzureLargeInstancesServer struct {
	// Get is the fake for method AzureLargeInstancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, azureLargeInstanceName string, options *armlargeinstance.AzureLargeInstancesClientGetOptions) (resp azfake.Responder[armlargeinstance.AzureLargeInstancesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method AzureLargeInstancesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armlargeinstance.AzureLargeInstancesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armlargeinstance.AzureLargeInstancesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method AzureLargeInstancesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armlargeinstance.AzureLargeInstancesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armlargeinstance.AzureLargeInstancesClientListBySubscriptionResponse])

	// BeginRestart is the fake for method AzureLargeInstancesClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestart func(ctx context.Context, resourceGroupName string, azureLargeInstanceName string, options *armlargeinstance.AzureLargeInstancesClientBeginRestartOptions) (resp azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientRestartResponse], errResp azfake.ErrorResponder)

	// BeginShutdown is the fake for method AzureLargeInstancesClient.BeginShutdown
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginShutdown func(ctx context.Context, resourceGroupName string, azureLargeInstanceName string, body any, options *armlargeinstance.AzureLargeInstancesClientBeginShutdownOptions) (resp azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientShutdownResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method AzureLargeInstancesClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, azureLargeInstanceName string, body any, options *armlargeinstance.AzureLargeInstancesClientBeginStartOptions) (resp azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientStartResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method AzureLargeInstancesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, azureLargeInstanceName string, properties armlargeinstance.TagsUpdate, options *armlargeinstance.AzureLargeInstancesClientUpdateOptions) (resp azfake.Responder[armlargeinstance.AzureLargeInstancesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAzureLargeInstancesServerTransport creates a new instance of AzureLargeInstancesServerTransport with the provided implementation.
// The returned AzureLargeInstancesServerTransport instance is connected to an instance of armlargeinstance.AzureLargeInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAzureLargeInstancesServerTransport(srv *AzureLargeInstancesServer) *AzureLargeInstancesServerTransport {
	return &AzureLargeInstancesServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armlargeinstance.AzureLargeInstancesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armlargeinstance.AzureLargeInstancesClientListBySubscriptionResponse]](),
		beginRestart:                newTracker[azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientRestartResponse]](),
		beginShutdown:               newTracker[azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientShutdownResponse]](),
		beginStart:                  newTracker[azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientStartResponse]](),
	}
}

// AzureLargeInstancesServerTransport connects instances of armlargeinstance.AzureLargeInstancesClient to instances of AzureLargeInstancesServer.
// Don't use this type directly, use NewAzureLargeInstancesServerTransport instead.
type AzureLargeInstancesServerTransport struct {
	srv                         *AzureLargeInstancesServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armlargeinstance.AzureLargeInstancesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armlargeinstance.AzureLargeInstancesClientListBySubscriptionResponse]]
	beginRestart                *tracker[azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientRestartResponse]]
	beginShutdown               *tracker[azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientShutdownResponse]]
	beginStart                  *tracker[azfake.PollerResponder[armlargeinstance.AzureLargeInstancesClientStartResponse]]
}

// Do implements the policy.Transporter interface for AzureLargeInstancesServerTransport.
func (a *AzureLargeInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AzureLargeInstancesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "AzureLargeInstancesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AzureLargeInstancesClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "AzureLargeInstancesClient.NewListBySubscriptionPager":
		resp, err = a.dispatchNewListBySubscriptionPager(req)
	case "AzureLargeInstancesClient.BeginRestart":
		resp, err = a.dispatchBeginRestart(req)
	case "AzureLargeInstancesClient.BeginShutdown":
		resp, err = a.dispatchBeginShutdown(req)
	case "AzureLargeInstancesClient.BeginStart":
		resp, err = a.dispatchBeginStart(req)
	case "AzureLargeInstancesClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (a *AzureLargeInstancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeInstances/(?P<azureLargeInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureLargeInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureLargeInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, azureLargeInstanceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureLargeInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureLargeInstancesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armlargeinstance.AzureLargeInstancesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (a *AzureLargeInstancesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := a.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		a.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armlargeinstance.AzureLargeInstancesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		a.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (a *AzureLargeInstancesServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if a.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	beginRestart := a.beginRestart.get(req)
	if beginRestart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeInstances/(?P<azureLargeInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armlargeinstance.ForceState](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureLargeInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureLargeInstanceName")])
		if err != nil {
			return nil, err
		}
		var options *armlargeinstance.AzureLargeInstancesClientBeginRestartOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armlargeinstance.AzureLargeInstancesClientBeginRestartOptions{
				ForceParameter: &body,
			}
		}
		respr, errRespr := a.srv.BeginRestart(req.Context(), resourceGroupNameParam, azureLargeInstanceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestart = &respr
		a.beginRestart.add(req, beginRestart)
	}

	resp, err := server.PollerResponderNext(beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginRestart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestart) {
		a.beginRestart.remove(req)
	}

	return resp, nil
}

func (a *AzureLargeInstancesServerTransport) dispatchBeginShutdown(req *http.Request) (*http.Response, error) {
	if a.srv.BeginShutdown == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginShutdown not implemented")}
	}
	beginShutdown := a.beginShutdown.get(req)
	if beginShutdown == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeInstances/(?P<azureLargeInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shutdown`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[any](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureLargeInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureLargeInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginShutdown(req.Context(), resourceGroupNameParam, azureLargeInstanceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginShutdown = &respr
		a.beginShutdown.add(req, beginShutdown)
	}

	resp, err := server.PollerResponderNext(beginShutdown, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginShutdown.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginShutdown) {
		a.beginShutdown.remove(req)
	}

	return resp, nil
}

func (a *AzureLargeInstancesServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if a.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := a.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeInstances/(?P<azureLargeInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[any](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureLargeInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureLargeInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginStart(req.Context(), resourceGroupNameParam, azureLargeInstanceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		a.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		a.beginStart.remove(req)
	}

	return resp, nil
}

func (a *AzureLargeInstancesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeInstances/(?P<azureLargeInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlargeinstance.TagsUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureLargeInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureLargeInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameParam, azureLargeInstanceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureLargeInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
