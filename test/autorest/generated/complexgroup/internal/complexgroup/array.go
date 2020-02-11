// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

type ArrayOperations struct{}

// GetEmptyCreateRequest creates the GetEmpty request.
func (ArrayOperations) GetEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/array/empty")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (ArrayOperations) GetEmptyHandleResponse(resp *azcore.Response) (*ArrayGetEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ArrayGetEmptyResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// GetNotProvidedCreateRequest creates the GetNotProvided request.
func (ArrayOperations) GetNotProvidedCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/array/notprovided")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNotProvidedHandleResponse handles the GetNotProvided response.
func (ArrayOperations) GetNotProvidedHandleResponse(resp *azcore.Response) (*ArrayGetNotProvidedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ArrayGetNotProvidedResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// GetValidCreateRequest creates the GetValid request.
func (ArrayOperations) GetValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/array/valid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetValidHandleResponse handles the GetValid response.
func (ArrayOperations) GetValidHandleResponse(resp *azcore.Response) (*ArrayGetValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ArrayGetValidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// PutEmptyCreateRequest creates the PutEmpty request.
func (ArrayOperations) PutEmptyCreateRequest(u url.URL, complexBody ArrayWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/array/empty")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutEmptyHandleResponse handles the PutEmpty response.
func (ArrayOperations) PutEmptyHandleResponse(resp *azcore.Response) (*ArrayPutEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &ArrayPutEmptyResponse{StatusCode: resp.StatusCode}, nil
}

// PutValidCreateRequest creates the PutValid request.
func (ArrayOperations) PutValidCreateRequest(u url.URL, complexBody ArrayWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/array/valid")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutValidHandleResponse handles the PutValid response.
func (ArrayOperations) PutValidHandleResponse(resp *azcore.Response) (*ArrayPutValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &ArrayPutValidResponse{StatusCode: resp.StatusCode}, nil
}
