//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"azkeyvault"
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

// RoleAssignmentsServer is a fake server for instances of the azkeyvault.RoleAssignmentsClient type.
type RoleAssignmentsServer struct {
	// Create is the fake for method RoleAssignmentsClient.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, host string, scope string, roleAssignmentName string, parameters azkeyvault.RoleAssignmentCreateParameters, options *azkeyvault.RoleAssignmentsClientCreateOptions) (resp azfake.Responder[azkeyvault.RoleAssignmentsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method RoleAssignmentsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, host string, scope string, roleAssignmentName string, options *azkeyvault.RoleAssignmentsClientDeleteOptions) (resp azfake.Responder[azkeyvault.RoleAssignmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RoleAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, host string, scope string, roleAssignmentName string, options *azkeyvault.RoleAssignmentsClientGetOptions) (resp azfake.Responder[azkeyvault.RoleAssignmentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListForScopePager is the fake for method RoleAssignmentsClient.NewListForScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForScopePager func(host string, scope string, options *azkeyvault.RoleAssignmentsClientListForScopeOptions) (resp azfake.PagerResponder[azkeyvault.RoleAssignmentsClientListForScopeResponse])
}

// NewRoleAssignmentsServerTransport creates a new instance of RoleAssignmentsServerTransport with the provided implementation.
// The returned RoleAssignmentsServerTransport instance is connected to an instance of azkeyvault.RoleAssignmentsClient by way of the
// undefined.Transporter field.
func NewRoleAssignmentsServerTransport(srv *RoleAssignmentsServer) *RoleAssignmentsServerTransport {
	return &RoleAssignmentsServerTransport{srv: srv}
}

// RoleAssignmentsServerTransport connects instances of azkeyvault.RoleAssignmentsClient to instances of RoleAssignmentsServer.
// Don't use this type directly, use NewRoleAssignmentsServerTransport instead.
type RoleAssignmentsServerTransport struct {
	srv                  *RoleAssignmentsServer
	newListForScopePager *azfake.PagerResponder[azkeyvault.RoleAssignmentsClientListForScopeResponse]
}

// Do implements the policy.Transporter interface for RoleAssignmentsServerTransport.
func (r *RoleAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RoleAssignmentsClient.Create":
		resp, err = r.dispatchCreate(req)
	case "RoleAssignmentsClient.Delete":
		resp, err = r.dispatchDelete(req)
	case "RoleAssignmentsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RoleAssignmentsClient.NewListForScopePager":
		resp, err = r.dispatchNewListForScopePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if r.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("method Create not implemented")}
	}
	const regexStr = "/(?P<scope>[a-zA-Z0-9-_]+)/providers/Microsoft.Authorization/roleAssignments/(?P<roleAssignmentName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azkeyvault.RoleAssignmentCreateParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Create(req.Context(), req.URL.Host, matches[regex.SubexpIndex("scope")], matches[regex.SubexpIndex("roleAssignmentName")], body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if r.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("method Delete not implemented")}
	}
	const regexStr = "/(?P<scope>[a-zA-Z0-9-_]+)/providers/Microsoft.Authorization/roleAssignments/(?P<roleAssignmentName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := r.srv.Delete(req.Context(), req.URL.Host, matches[regex.SubexpIndex("scope")], matches[regex.SubexpIndex("roleAssignmentName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/(?P<scope>[a-zA-Z0-9-_]+)/providers/Microsoft.Authorization/roleAssignments/(?P<roleAssignmentName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := r.srv.Get(req.Context(), req.URL.Host, matches[regex.SubexpIndex("scope")], matches[regex.SubexpIndex("roleAssignmentName")], nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleAssignmentsServerTransport) dispatchNewListForScopePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListForScopePager == nil {
		return nil, &nonRetriableError{errors.New("method NewListForScopePager not implemented")}
	}
	if r.newListForScopePager == nil {
		const regexStr = "/(?P<scope>[a-zA-Z0-9-_]+)/providers/Microsoft.Authorization/roleAssignments"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterParam := getOptional(qp.Get("$filter"))
		var options *azkeyvault.RoleAssignmentsClientListForScopeOptions
		if filterParam != nil {
			options = &azkeyvault.RoleAssignmentsClientListForScopeOptions{
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListForScopePager(req.URL.Host, matches[regex.SubexpIndex("scope")], options)
		r.newListForScopePager = &resp
		server.PagerResponderInjectNextLinks(r.newListForScopePager, req, func(page *azkeyvault.RoleAssignmentsClientListForScopeResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(r.newListForScopePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(r.newListForScopePager) {
		r.newListForScopePager = nil
	}
	return resp, nil
}
