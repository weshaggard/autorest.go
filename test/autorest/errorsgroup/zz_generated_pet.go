// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// PetClient contains the methods for the Pet group.
// Don't use this type directly, use NewPetClient() instead.
type PetClient struct {
	con *Connection
}

// NewPetClient creates a new instance of PetClient with the specified values.
func NewPetClient(con *Connection) PetClient {
	return PetClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client PetClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// DoSomething - Asks pet to do something
func (client PetClient) DoSomething(ctx context.Context, whatAction string, options *PetDoSomethingOptions) (PetActionResponse, error) {
	req, err := client.doSomethingCreateRequest(ctx, whatAction, options)
	if err != nil {
		return PetActionResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return PetActionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PetActionResponse{}, client.doSomethingHandleError(resp)
	}
	return client.doSomethingHandleResponse(resp)
}

// doSomethingCreateRequest creates the DoSomething request.
func (client PetClient) doSomethingCreateRequest(ctx context.Context, whatAction string, options *PetDoSomethingOptions) (*azcore.Request, error) {
	urlPath := "/errorStatusCodes/Pets/doSomething/{whatAction}"
	urlPath = strings.ReplaceAll(urlPath, "{whatAction}", url.PathEscape(whatAction))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// doSomethingHandleResponse handles the DoSomething response.
func (client PetClient) doSomethingHandleResponse(resp *azcore.Response) (PetActionResponse, error) {
	var val *PetAction
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PetActionResponse{}, err
	}
	return PetActionResponse{RawResponse: resp.Response, PetAction: val}, nil
}

// doSomethingHandleError handles the DoSomething error response.
func (client PetClient) doSomethingHandleError(resp *azcore.Response) error {
	switch resp.StatusCode {
	case http.StatusInternalServerError:
		var err petActionError
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(err.wrapped, resp.Response)
	default:
		var err petActionError
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(err.wrapped, resp.Response)
	}
}

// GetPetByID - Gets pets by id.
func (client PetClient) GetPetByID(ctx context.Context, petId string, options *PetGetPetByIDOptions) (PetResponse, error) {
	req, err := client.getPetByIdCreateRequest(ctx, petId, options)
	if err != nil {
		return PetResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return PetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return PetResponse{}, client.getPetByIdHandleError(resp)
	}
	return client.getPetByIdHandleResponse(resp)
}

// getPetByIdCreateRequest creates the GetPetByID request.
func (client PetClient) getPetByIdCreateRequest(ctx context.Context, petId string, options *PetGetPetByIDOptions) (*azcore.Request, error) {
	urlPath := "/errorStatusCodes/Pets/{petId}/GetPet"
	urlPath = strings.ReplaceAll(urlPath, "{petId}", url.PathEscape(petId))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPetByIdHandleResponse handles the GetPetByID response.
func (client PetClient) getPetByIdHandleResponse(resp *azcore.Response) (PetResponse, error) {
	var val *Pet
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PetResponse{}, err
	}
	return PetResponse{RawResponse: resp.Response, Pet: val}, nil
}

// getPetByIdHandleError handles the GetPetByID error response.
func (client PetClient) getPetByIdHandleError(resp *azcore.Response) error {
	switch resp.StatusCode {
	case http.StatusBadRequest:
		var err string
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(fmt.Errorf("%v", err), resp.Response)
	case http.StatusNotFound:
		var err notFoundErrorBase
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(err.wrapped, resp.Response)
	case http.StatusNotImplemented:
		var err int32
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(fmt.Errorf("%v", err), resp.Response)
	default:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
		}
		if len(body) == 0 {
			return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
		}
		return azcore.NewResponseError(errors.New(string(body)), resp.Response)
	}
}
