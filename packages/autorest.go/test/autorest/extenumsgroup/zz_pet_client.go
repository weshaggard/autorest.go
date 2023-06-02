//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package extenumsgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PetClient contains the methods for the Pet group.
// Don't use this type directly, use a constructor function instead.
type PetClient struct {
	internal *azcore.Client
}

// AddPet - add pet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-07-07
//   - options - PetClientAddPetOptions contains the optional parameters for the PetClient.AddPet method.
func (client *PetClient) AddPet(ctx context.Context, options *PetClientAddPetOptions) (PetClientAddPetResponse, error) {
	var err error
	const operationName = "PetClient.AddPet"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.addPetCreateRequest(ctx, options)
	if err != nil {
		return PetClientAddPetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetClientAddPetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetClientAddPetResponse{}, err
	}
	resp, err := client.addPetHandleResponse(httpResp)
	return resp, err
}

// addPetCreateRequest creates the AddPet request.
func (client *PetClient) addPetCreateRequest(ctx context.Context, options *PetClientAddPetOptions) (*policy.Request, error) {
	urlPath := "/extensibleenums/pet/addPet"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.PetParam != nil {
		if err := runtime.MarshalAsJSON(req, *options.PetParam); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// addPetHandleResponse handles the AddPet response.
func (client *PetClient) addPetHandleResponse(resp *http.Response) (PetClientAddPetResponse, error) {
	result := PetClientAddPetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pet); err != nil {
		return PetClientAddPetResponse{}, err
	}
	return result, nil
}

// GetByPetID - get pet by id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-07-07
//   - petID - Pet id
//   - options - PetClientGetByPetIDOptions contains the optional parameters for the PetClient.GetByPetID method.
func (client *PetClient) GetByPetID(ctx context.Context, petID string, options *PetClientGetByPetIDOptions) (PetClientGetByPetIDResponse, error) {
	var err error
	const operationName = "PetClient.GetByPetID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByPetIDCreateRequest(ctx, petID, options)
	if err != nil {
		return PetClientGetByPetIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetClientGetByPetIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetClientGetByPetIDResponse{}, err
	}
	resp, err := client.getByPetIDHandleResponse(httpResp)
	return resp, err
}

// getByPetIDCreateRequest creates the GetByPetID request.
func (client *PetClient) getByPetIDCreateRequest(ctx context.Context, petID string, options *PetClientGetByPetIDOptions) (*policy.Request, error) {
	urlPath := "/extensibleenums/pet/{petId}"
	if petID == "" {
		return nil, errors.New("parameter petID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{petId}", url.PathEscape(petID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByPetIDHandleResponse handles the GetByPetID response.
func (client *PetClient) getByPetIDHandleResponse(resp *http.Response) (PetClientGetByPetIDResponse, error) {
	result := PetClientGetByPetIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pet); err != nil {
		return PetClientGetByPetIDResponse{}, err
	}
	return result, nil
}
