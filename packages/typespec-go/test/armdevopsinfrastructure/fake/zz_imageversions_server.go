// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armdevopsinfrastructure"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
)

// ImageVersionsServer is a fake server for instances of the armdevopsinfrastructure.ImageVersionsClient type.
type ImageVersionsServer struct {
	// NewListByImagePager is the fake for method ImageVersionsClient.NewListByImagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByImagePager func(resourceGroupName string, imageName string, options *armdevopsinfrastructure.ImageVersionsClientListByImageOptions) (resp azfake.PagerResponder[armdevopsinfrastructure.ImageVersionsClientListByImageResponse])
}

// NewImageVersionsServerTransport creates a new instance of ImageVersionsServerTransport with the provided implementation.
// The returned ImageVersionsServerTransport instance is connected to an instance of armdevopsinfrastructure.ImageVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewImageVersionsServerTransport(srv *ImageVersionsServer) *ImageVersionsServerTransport {
	return &ImageVersionsServerTransport{
		srv:                 srv,
		newListByImagePager: newTracker[azfake.PagerResponder[armdevopsinfrastructure.ImageVersionsClientListByImageResponse]](),
	}
}

// ImageVersionsServerTransport connects instances of armdevopsinfrastructure.ImageVersionsClient to instances of ImageVersionsServer.
// Don't use this type directly, use NewImageVersionsServerTransport instead.
type ImageVersionsServerTransport struct {
	srv                 *ImageVersionsServer
	newListByImagePager *tracker[azfake.PagerResponder[armdevopsinfrastructure.ImageVersionsClientListByImageResponse]]
}

// Do implements the policy.Transporter interface for ImageVersionsServerTransport.
func (i *ImageVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *ImageVersionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ImageVersionsClient.NewListByImagePager":
		resp, err = i.dispatchNewListByImagePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (i *ImageVersionsServerTransport) dispatchNewListByImagePager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByImagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByImagePager not implemented")}
	}
	newListByImagePager := i.newListByImagePager.get(req)
	if newListByImagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevOpsInfrastructure/images/(?P<imageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		imageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByImagePager(resourceGroupNameParam, imageNameParam, nil)
		newListByImagePager = &resp
		i.newListByImagePager.add(req, newListByImagePager)
		server.PagerResponderInjectNextLinks(newListByImagePager, req, func(page *armdevopsinfrastructure.ImageVersionsClientListByImageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByImagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByImagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByImagePager) {
		i.newListByImagePager.remove(req)
	}
	return resp, nil
}
