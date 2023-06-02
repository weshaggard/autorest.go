//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package migroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// MultipleInheritanceServiceClient contains the methods for the MultipleInheritanceServiceClient group.
// Don't use this type directly, use a constructor function instead.
type MultipleInheritanceServiceClient struct {
	internal *azcore.Client
}

// GetCat - Get a cat with name 'Whiskers' where likesMilk, meows, and hisses is true
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - options - MultipleInheritanceServiceClientGetCatOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetCat
//     method.
func (client *MultipleInheritanceServiceClient) GetCat(ctx context.Context, options *MultipleInheritanceServiceClientGetCatOptions) (MultipleInheritanceServiceClientGetCatResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.GetCat"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCatCreateRequest(ctx, options)
	if err != nil {
		return MultipleInheritanceServiceClientGetCatResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientGetCatResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientGetCatResponse{}, err
	}
	resp, err := client.getCatHandleResponse(httpResp)
	return resp, err
}

// getCatCreateRequest creates the GetCat request.
func (client *MultipleInheritanceServiceClient) getCatCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetCatOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/cat"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getCatHandleResponse handles the GetCat response.
func (client *MultipleInheritanceServiceClient) getCatHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientGetCatResponse, error) {
	result := MultipleInheritanceServiceClientGetCatResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cat); err != nil {
		return MultipleInheritanceServiceClientGetCatResponse{}, err
	}
	return result, nil
}

// GetFeline - Get a feline where meows and hisses are true
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - options - MultipleInheritanceServiceClientGetFelineOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetFeline
//     method.
func (client *MultipleInheritanceServiceClient) GetFeline(ctx context.Context, options *MultipleInheritanceServiceClientGetFelineOptions) (MultipleInheritanceServiceClientGetFelineResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.GetFeline"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getFelineCreateRequest(ctx, options)
	if err != nil {
		return MultipleInheritanceServiceClientGetFelineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientGetFelineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientGetFelineResponse{}, err
	}
	resp, err := client.getFelineHandleResponse(httpResp)
	return resp, err
}

// getFelineCreateRequest creates the GetFeline request.
func (client *MultipleInheritanceServiceClient) getFelineCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetFelineOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/feline"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFelineHandleResponse handles the GetFeline response.
func (client *MultipleInheritanceServiceClient) getFelineHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientGetFelineResponse, error) {
	result := MultipleInheritanceServiceClientGetFelineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Feline); err != nil {
		return MultipleInheritanceServiceClientGetFelineResponse{}, err
	}
	return result, nil
}

// GetHorse - Get a horse with name 'Fred' and isAShowHorse true
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - options - MultipleInheritanceServiceClientGetHorseOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetHorse
//     method.
func (client *MultipleInheritanceServiceClient) GetHorse(ctx context.Context, options *MultipleInheritanceServiceClientGetHorseOptions) (MultipleInheritanceServiceClientGetHorseResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.GetHorse"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getHorseCreateRequest(ctx, options)
	if err != nil {
		return MultipleInheritanceServiceClientGetHorseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientGetHorseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientGetHorseResponse{}, err
	}
	resp, err := client.getHorseHandleResponse(httpResp)
	return resp, err
}

// getHorseCreateRequest creates the GetHorse request.
func (client *MultipleInheritanceServiceClient) getHorseCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetHorseOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/horse"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHorseHandleResponse handles the GetHorse response.
func (client *MultipleInheritanceServiceClient) getHorseHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientGetHorseResponse, error) {
	result := MultipleInheritanceServiceClientGetHorseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Horse); err != nil {
		return MultipleInheritanceServiceClientGetHorseResponse{}, err
	}
	return result, nil
}

// GetKitten - Get a kitten with name 'Gatito' where likesMilk and meows is true, and hisses and eatsMiceYet is false
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - options - MultipleInheritanceServiceClientGetKittenOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetKitten
//     method.
func (client *MultipleInheritanceServiceClient) GetKitten(ctx context.Context, options *MultipleInheritanceServiceClientGetKittenOptions) (MultipleInheritanceServiceClientGetKittenResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.GetKitten"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getKittenCreateRequest(ctx, options)
	if err != nil {
		return MultipleInheritanceServiceClientGetKittenResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientGetKittenResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientGetKittenResponse{}, err
	}
	resp, err := client.getKittenHandleResponse(httpResp)
	return resp, err
}

// getKittenCreateRequest creates the GetKitten request.
func (client *MultipleInheritanceServiceClient) getKittenCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetKittenOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/kitten"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getKittenHandleResponse handles the GetKitten response.
func (client *MultipleInheritanceServiceClient) getKittenHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientGetKittenResponse, error) {
	result := MultipleInheritanceServiceClientGetKittenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Kitten); err != nil {
		return MultipleInheritanceServiceClientGetKittenResponse{}, err
	}
	return result, nil
}

// GetPet - Get a pet with name 'Peanut'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - options - MultipleInheritanceServiceClientGetPetOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetPet
//     method.
func (client *MultipleInheritanceServiceClient) GetPet(ctx context.Context, options *MultipleInheritanceServiceClientGetPetOptions) (MultipleInheritanceServiceClientGetPetResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.GetPet"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getPetCreateRequest(ctx, options)
	if err != nil {
		return MultipleInheritanceServiceClientGetPetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientGetPetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientGetPetResponse{}, err
	}
	resp, err := client.getPetHandleResponse(httpResp)
	return resp, err
}

// getPetCreateRequest creates the GetPet request.
func (client *MultipleInheritanceServiceClient) getPetCreateRequest(ctx context.Context, options *MultipleInheritanceServiceClientGetPetOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/pet"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPetHandleResponse handles the GetPet response.
func (client *MultipleInheritanceServiceClient) getPetHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientGetPetResponse, error) {
	result := MultipleInheritanceServiceClientGetPetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pet); err != nil {
		return MultipleInheritanceServiceClientGetPetResponse{}, err
	}
	return result, nil
}

// PutCat - Put a cat with name 'Boots' where likesMilk and hisses is false, meows is true
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - cat - Put a cat with name 'Boots' where likesMilk and hisses is false, meows is true
//   - options - MultipleInheritanceServiceClientPutCatOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutCat
//     method.
func (client *MultipleInheritanceServiceClient) PutCat(ctx context.Context, cat Cat, options *MultipleInheritanceServiceClientPutCatOptions) (MultipleInheritanceServiceClientPutCatResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.PutCat"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCatCreateRequest(ctx, cat, options)
	if err != nil {
		return MultipleInheritanceServiceClientPutCatResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientPutCatResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientPutCatResponse{}, err
	}
	resp, err := client.putCatHandleResponse(httpResp)
	return resp, err
}

// putCatCreateRequest creates the PutCat request.
func (client *MultipleInheritanceServiceClient) putCatCreateRequest(ctx context.Context, cat Cat, options *MultipleInheritanceServiceClientPutCatOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/cat"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, cat); err != nil {
		return nil, err
	}
	return req, nil
}

// putCatHandleResponse handles the PutCat response.
func (client *MultipleInheritanceServiceClient) putCatHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientPutCatResponse, error) {
	result := MultipleInheritanceServiceClientPutCatResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MultipleInheritanceServiceClientPutCatResponse{}, err
	}
	return result, nil
}

// PutFeline - Put a feline who hisses and doesn't meow
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - feline - Put a feline who hisses and doesn't meow
//   - options - MultipleInheritanceServiceClientPutFelineOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutFeline
//     method.
func (client *MultipleInheritanceServiceClient) PutFeline(ctx context.Context, feline Feline, options *MultipleInheritanceServiceClientPutFelineOptions) (MultipleInheritanceServiceClientPutFelineResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.PutFeline"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putFelineCreateRequest(ctx, feline, options)
	if err != nil {
		return MultipleInheritanceServiceClientPutFelineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientPutFelineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientPutFelineResponse{}, err
	}
	resp, err := client.putFelineHandleResponse(httpResp)
	return resp, err
}

// putFelineCreateRequest creates the PutFeline request.
func (client *MultipleInheritanceServiceClient) putFelineCreateRequest(ctx context.Context, feline Feline, options *MultipleInheritanceServiceClientPutFelineOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/feline"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, feline); err != nil {
		return nil, err
	}
	return req, nil
}

// putFelineHandleResponse handles the PutFeline response.
func (client *MultipleInheritanceServiceClient) putFelineHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientPutFelineResponse, error) {
	result := MultipleInheritanceServiceClientPutFelineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MultipleInheritanceServiceClientPutFelineResponse{}, err
	}
	return result, nil
}

// PutHorse - Put a horse with name 'General' and isAShowHorse false
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - horse - Put a horse with name 'General' and isAShowHorse false
//   - options - MultipleInheritanceServiceClientPutHorseOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutHorse
//     method.
func (client *MultipleInheritanceServiceClient) PutHorse(ctx context.Context, horse Horse, options *MultipleInheritanceServiceClientPutHorseOptions) (MultipleInheritanceServiceClientPutHorseResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.PutHorse"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putHorseCreateRequest(ctx, horse, options)
	if err != nil {
		return MultipleInheritanceServiceClientPutHorseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientPutHorseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientPutHorseResponse{}, err
	}
	resp, err := client.putHorseHandleResponse(httpResp)
	return resp, err
}

// putHorseCreateRequest creates the PutHorse request.
func (client *MultipleInheritanceServiceClient) putHorseCreateRequest(ctx context.Context, horse Horse, options *MultipleInheritanceServiceClientPutHorseOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/horse"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, horse); err != nil {
		return nil, err
	}
	return req, nil
}

// putHorseHandleResponse handles the PutHorse response.
func (client *MultipleInheritanceServiceClient) putHorseHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientPutHorseResponse, error) {
	result := MultipleInheritanceServiceClientPutHorseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MultipleInheritanceServiceClientPutHorseResponse{}, err
	}
	return result, nil
}

// PutKitten - Put a kitten with name 'Kitty' where likesMilk and hisses is false, meows and eatsMiceYet is true
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - kitten - Put a kitten with name 'Kitty' where likesMilk and hisses is false, meows and eatsMiceYet is true
//   - options - MultipleInheritanceServiceClientPutKittenOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutKitten
//     method.
func (client *MultipleInheritanceServiceClient) PutKitten(ctx context.Context, kitten Kitten, options *MultipleInheritanceServiceClientPutKittenOptions) (MultipleInheritanceServiceClientPutKittenResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.PutKitten"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putKittenCreateRequest(ctx, kitten, options)
	if err != nil {
		return MultipleInheritanceServiceClientPutKittenResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientPutKittenResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientPutKittenResponse{}, err
	}
	resp, err := client.putKittenHandleResponse(httpResp)
	return resp, err
}

// putKittenCreateRequest creates the PutKitten request.
func (client *MultipleInheritanceServiceClient) putKittenCreateRequest(ctx context.Context, kitten Kitten, options *MultipleInheritanceServiceClientPutKittenOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/kitten"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, kitten); err != nil {
		return nil, err
	}
	return req, nil
}

// putKittenHandleResponse handles the PutKitten response.
func (client *MultipleInheritanceServiceClient) putKittenHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientPutKittenResponse, error) {
	result := MultipleInheritanceServiceClientPutKittenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MultipleInheritanceServiceClientPutKittenResponse{}, err
	}
	return result, nil
}

// PutPet - Put a pet with name 'Butter'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 3.0.0
//   - pet - Put a pet with name 'Butter'
//   - options - MultipleInheritanceServiceClientPutPetOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutPet
//     method.
func (client *MultipleInheritanceServiceClient) PutPet(ctx context.Context, pet Pet, options *MultipleInheritanceServiceClientPutPetOptions) (MultipleInheritanceServiceClientPutPetResponse, error) {
	var err error
	const operationName = "MultipleInheritanceServiceClient.PutPet"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putPetCreateRequest(ctx, pet, options)
	if err != nil {
		return MultipleInheritanceServiceClientPutPetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleInheritanceServiceClientPutPetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MultipleInheritanceServiceClientPutPetResponse{}, err
	}
	resp, err := client.putPetHandleResponse(httpResp)
	return resp, err
}

// putPetCreateRequest creates the PutPet request.
func (client *MultipleInheritanceServiceClient) putPetCreateRequest(ctx context.Context, pet Pet, options *MultipleInheritanceServiceClientPutPetOptions) (*policy.Request, error) {
	urlPath := "/multipleInheritance/pet"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, pet); err != nil {
		return nil, err
	}
	return req, nil
}

// putPetHandleResponse handles the PutPet response.
func (client *MultipleInheritanceServiceClient) putPetHandleResponse(resp *http.Response) (MultipleInheritanceServiceClientPutPetResponse, error) {
	result := MultipleInheritanceServiceClientPutPetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MultipleInheritanceServiceClientPutPetResponse{}, err
	}
	return result, nil
}
