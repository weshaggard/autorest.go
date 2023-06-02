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

// PeerExpressRouteCircuitConnectionsServer is a fake server for instances of the armnetwork.PeerExpressRouteCircuitConnectionsClient type.
type PeerExpressRouteCircuitConnectionsServer struct {
	// Get is the fake for method PeerExpressRouteCircuitConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *armnetwork.PeerExpressRouteCircuitConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.PeerExpressRouteCircuitConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PeerExpressRouteCircuitConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, circuitName string, peeringName string, options *armnetwork.PeerExpressRouteCircuitConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.PeerExpressRouteCircuitConnectionsClientListResponse])
}

// NewPeerExpressRouteCircuitConnectionsServerTransport creates a new instance of PeerExpressRouteCircuitConnectionsServerTransport with the provided implementation.
// The returned PeerExpressRouteCircuitConnectionsServerTransport instance is connected to an instance of armnetwork.PeerExpressRouteCircuitConnectionsClient by way of the
// undefined.Transporter field.
func NewPeerExpressRouteCircuitConnectionsServerTransport(srv *PeerExpressRouteCircuitConnectionsServer) *PeerExpressRouteCircuitConnectionsServerTransport {
	return &PeerExpressRouteCircuitConnectionsServerTransport{srv: srv}
}

// PeerExpressRouteCircuitConnectionsServerTransport connects instances of armnetwork.PeerExpressRouteCircuitConnectionsClient to instances of PeerExpressRouteCircuitConnectionsServer.
// Don't use this type directly, use NewPeerExpressRouteCircuitConnectionsServerTransport instead.
type PeerExpressRouteCircuitConnectionsServerTransport struct {
	srv          *PeerExpressRouteCircuitConnectionsServer
	newListPager *azfake.PagerResponder[armnetwork.PeerExpressRouteCircuitConnectionsClientListResponse]
}

// Do implements the policy.Transporter interface for PeerExpressRouteCircuitConnectionsServerTransport.
func (p *PeerExpressRouteCircuitConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PeerExpressRouteCircuitConnectionsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PeerExpressRouteCircuitConnectionsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PeerExpressRouteCircuitConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[a-zA-Z0-9-_]+)/peerings/(?P<peeringName>[a-zA-Z0-9-_]+)/peerConnections/(?P<connectionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("circuitName")], matches[regex.SubexpIndex("peeringName")], matches[regex.SubexpIndex("connectionName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PeerExpressRouteCircuitConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PeerExpressRouteCircuitConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if p.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[a-zA-Z0-9-_]+)/peerings/(?P<peeringName>[a-zA-Z0-9-_]+)/peerConnections"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("circuitName")], matches[regex.SubexpIndex("peeringName")], nil)
		p.newListPager = &resp
		server.PagerResponderInjectNextLinks(p.newListPager, req, func(page *armnetwork.PeerExpressRouteCircuitConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListPager) {
		p.newListPager = nil
	}
	return resp, nil
}
