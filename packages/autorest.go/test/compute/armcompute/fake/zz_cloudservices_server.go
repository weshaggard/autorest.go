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
	"reflect"
	"regexp"
)

// CloudServicesServer is a fake server for instances of the armcompute.CloudServicesClient type.
type CloudServicesServer struct {
	// BeginCreateOrUpdate is the fake for method CloudServicesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, cloudServiceName string, parameters armcompute.CloudService, options *armcompute.CloudServicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CloudServicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeleteInstances is the fake for method CloudServicesClient.BeginDeleteInstances
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDeleteInstances func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginDeleteInstancesOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientDeleteInstancesResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CloudServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientGetOptions) (resp azfake.Responder[armcompute.CloudServicesClientGetResponse], errResp azfake.ErrorResponder)

	// GetInstanceView is the fake for method CloudServicesClient.GetInstanceView
	// HTTP status codes to indicate success: http.StatusOK
	GetInstanceView func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientGetInstanceViewOptions) (resp azfake.Responder[armcompute.CloudServicesClientGetInstanceViewResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CloudServicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armcompute.CloudServicesClientListOptions) (resp azfake.PagerResponder[armcompute.CloudServicesClientListResponse])

	// NewListAllPager is the fake for method CloudServicesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armcompute.CloudServicesClientListAllOptions) (resp azfake.PagerResponder[armcompute.CloudServicesClientListAllResponse])

	// BeginPowerOff is the fake for method CloudServicesClient.BeginPowerOff
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPowerOff func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginPowerOffOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientPowerOffResponse], errResp azfake.ErrorResponder)

	// BeginRebuild is the fake for method CloudServicesClient.BeginRebuild
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRebuild func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginRebuildOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientRebuildResponse], errResp azfake.ErrorResponder)

	// BeginReimage is the fake for method CloudServicesClient.BeginReimage
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReimage func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginReimageOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientReimageResponse], errResp azfake.ErrorResponder)

	// BeginRestart is the fake for method CloudServicesClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestart func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginRestartOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientRestartResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method CloudServicesClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginStartOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientStartResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method CloudServicesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, cloudServiceName string, parameters armcompute.CloudServiceUpdate, options *armcompute.CloudServicesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCloudServicesServerTransport creates a new instance of CloudServicesServerTransport with the provided implementation.
// The returned CloudServicesServerTransport instance is connected to an instance of armcompute.CloudServicesClient by way of the
// undefined.Transporter field.
func NewCloudServicesServerTransport(srv *CloudServicesServer) *CloudServicesServerTransport {
	return &CloudServicesServerTransport{srv: srv}
}

// CloudServicesServerTransport connects instances of armcompute.CloudServicesClient to instances of CloudServicesServer.
// Don't use this type directly, use NewCloudServicesServerTransport instead.
type CloudServicesServerTransport struct {
	srv                  *CloudServicesServer
	beginCreateOrUpdate  *azfake.PollerResponder[armcompute.CloudServicesClientCreateOrUpdateResponse]
	beginDelete          *azfake.PollerResponder[armcompute.CloudServicesClientDeleteResponse]
	beginDeleteInstances *azfake.PollerResponder[armcompute.CloudServicesClientDeleteInstancesResponse]
	newListPager         *azfake.PagerResponder[armcompute.CloudServicesClientListResponse]
	newListAllPager      *azfake.PagerResponder[armcompute.CloudServicesClientListAllResponse]
	beginPowerOff        *azfake.PollerResponder[armcompute.CloudServicesClientPowerOffResponse]
	beginRebuild         *azfake.PollerResponder[armcompute.CloudServicesClientRebuildResponse]
	beginReimage         *azfake.PollerResponder[armcompute.CloudServicesClientReimageResponse]
	beginRestart         *azfake.PollerResponder[armcompute.CloudServicesClientRestartResponse]
	beginStart           *azfake.PollerResponder[armcompute.CloudServicesClientStartResponse]
	beginUpdate          *azfake.PollerResponder[armcompute.CloudServicesClientUpdateResponse]
}

// Do implements the policy.Transporter interface for CloudServicesServerTransport.
func (c *CloudServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudServicesClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CloudServicesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CloudServicesClient.BeginDeleteInstances":
		resp, err = c.dispatchBeginDeleteInstances(req)
	case "CloudServicesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CloudServicesClient.GetInstanceView":
		resp, err = c.dispatchGetInstanceView(req)
	case "CloudServicesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "CloudServicesClient.NewListAllPager":
		resp, err = c.dispatchNewListAllPager(req)
	case "CloudServicesClient.BeginPowerOff":
		resp, err = c.dispatchBeginPowerOff(req)
	case "CloudServicesClient.BeginRebuild":
		resp, err = c.dispatchBeginRebuild(req)
	case "CloudServicesClient.BeginReimage":
		resp, err = c.dispatchBeginReimage(req)
	case "CloudServicesClient.BeginRestart":
		resp, err = c.dispatchBeginRestart(req)
	case "CloudServicesClient.BeginStart":
		resp, err = c.dispatchBeginStart(req)
	case "CloudServicesClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if c.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.CloudService](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], body, nil)
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

func (c *CloudServicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if c.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
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

func (c *CloudServicesServerTransport) dispatchBeginDeleteInstances(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDeleteInstances == nil {
		return nil, &nonRetriableError{errors.New("method BeginDeleteInstances not implemented")}
	}
	if c.beginDeleteInstances == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/delete"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		var options *armcompute.CloudServicesClientBeginDeleteInstancesOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armcompute.CloudServicesClientBeginDeleteInstancesOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginDeleteInstances(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginDeleteInstances = &respr
	}

	resp, err := server.PollerResponderNext(c.beginDeleteInstances, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginDeleteInstances) {
		c.beginDeleteInstances = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := c.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudService, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchGetInstanceView(req *http.Request) (*http.Response, error) {
	if c.srv.GetInstanceView == nil {
		return nil, &nonRetriableError{errors.New("method GetInstanceView not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/instanceView"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := c.srv.GetInstanceView(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudServiceInstanceView, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if c.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		c.newListPager = &resp
		server.PagerResponderInjectNextLinks(c.newListPager, req, func(page *armcompute.CloudServicesClientListResponse, createLink func() string) {
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

func (c *CloudServicesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListAllPager not implemented")}
	}
	if c.newListAllPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListAllPager(nil)
		c.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(c.newListAllPager, req, func(page *armcompute.CloudServicesClientListAllResponse, createLink func() string) {
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

func (c *CloudServicesServerTransport) dispatchBeginPowerOff(req *http.Request) (*http.Response, error) {
	if c.srv.BeginPowerOff == nil {
		return nil, &nonRetriableError{errors.New("method BeginPowerOff not implemented")}
	}
	if c.beginPowerOff == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/poweroff"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginPowerOff(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginPowerOff = &respr
	}

	resp, err := server.PollerResponderNext(c.beginPowerOff, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginPowerOff) {
		c.beginPowerOff = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginRebuild(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRebuild == nil {
		return nil, &nonRetriableError{errors.New("method BeginRebuild not implemented")}
	}
	if c.beginRebuild == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/rebuild"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		var options *armcompute.CloudServicesClientBeginRebuildOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armcompute.CloudServicesClientBeginRebuildOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginRebuild(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginRebuild = &respr
	}

	resp, err := server.PollerResponderNext(c.beginRebuild, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginRebuild) {
		c.beginRebuild = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginReimage(req *http.Request) (*http.Response, error) {
	if c.srv.BeginReimage == nil {
		return nil, &nonRetriableError{errors.New("method BeginReimage not implemented")}
	}
	if c.beginReimage == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/reimage"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		var options *armcompute.CloudServicesClientBeginReimageOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armcompute.CloudServicesClientBeginReimageOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginReimage(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginReimage = &respr
	}

	resp, err := server.PollerResponderNext(c.beginReimage, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginReimage) {
		c.beginReimage = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("method BeginRestart not implemented")}
	}
	if c.beginRestart == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/restart"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		var options *armcompute.CloudServicesClientBeginRestartOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armcompute.CloudServicesClientBeginRestartOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginRestart(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginRestart = &respr
	}

	resp, err := server.PollerResponderNext(c.beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginRestart) {
		c.beginRestart = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("method BeginStart not implemented")}
	}
	if c.beginStart == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/start"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		respr, errRespr := c.srv.BeginStart(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginStart = &respr
	}

	resp, err := server.PollerResponderNext(c.beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginStart) {
		c.beginStart = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if c.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.CloudServiceUpdate](req)
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		c.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(c.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(c.beginUpdate) {
		c.beginUpdate = nil
	}

	return resp, nil
}
