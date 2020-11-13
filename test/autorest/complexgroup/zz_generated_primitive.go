// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PrimitiveOperations contains the methods for the Primitive group.
type PrimitiveOperations interface {
	// GetBool - Get complex types with bool properties
	GetBool(ctx context.Context, options *PrimitiveGetBoolOptions) (*BooleanWrapperResponse, error)
	// GetByte - Get complex types with byte properties
	GetByte(ctx context.Context, options *PrimitiveGetByteOptions) (*ByteWrapperResponse, error)
	// GetDate - Get complex types with date properties
	GetDate(ctx context.Context, options *PrimitiveGetDateOptions) (*DateWrapperResponse, error)
	// GetDateTime - Get complex types with datetime properties
	GetDateTime(ctx context.Context, options *PrimitiveGetDateTimeOptions) (*DatetimeWrapperResponse, error)
	// GetDateTimeRFC1123 - Get complex types with datetimeRfc1123 properties
	GetDateTimeRFC1123(ctx context.Context, options *PrimitiveGetDateTimeRFC1123Options) (*Datetimerfc1123WrapperResponse, error)
	// GetDouble - Get complex types with double properties
	GetDouble(ctx context.Context, options *PrimitiveGetDoubleOptions) (*DoubleWrapperResponse, error)
	// GetDuration - Get complex types with duration properties
	GetDuration(ctx context.Context, options *PrimitiveGetDurationOptions) (*DurationWrapperResponse, error)
	// GetFloat - Get complex types with float properties
	GetFloat(ctx context.Context, options *PrimitiveGetFloatOptions) (*FloatWrapperResponse, error)
	// GetInt - Get complex types with integer properties
	GetInt(ctx context.Context, options *PrimitiveGetIntOptions) (*IntWrapperResponse, error)
	// GetLong - Get complex types with long properties
	GetLong(ctx context.Context, options *PrimitiveGetLongOptions) (*LongWrapperResponse, error)
	// GetString - Get complex types with string properties
	GetString(ctx context.Context, options *PrimitiveGetStringOptions) (*StringWrapperResponse, error)
	// PutBool - Put complex types with bool properties
	PutBool(ctx context.Context, complexBody BooleanWrapper, options *PrimitivePutBoolOptions) (*http.Response, error)
	// PutByte - Put complex types with byte properties
	PutByte(ctx context.Context, complexBody ByteWrapper, options *PrimitivePutByteOptions) (*http.Response, error)
	// PutDate - Put complex types with date properties
	PutDate(ctx context.Context, complexBody DateWrapper, options *PrimitivePutDateOptions) (*http.Response, error)
	// PutDateTime - Put complex types with datetime properties
	PutDateTime(ctx context.Context, complexBody DatetimeWrapper, options *PrimitivePutDateTimeOptions) (*http.Response, error)
	// PutDateTimeRFC1123 - Put complex types with datetimeRfc1123 properties
	PutDateTimeRFC1123(ctx context.Context, complexBody Datetimerfc1123Wrapper, options *PrimitivePutDateTimeRFC1123Options) (*http.Response, error)
	// PutDouble - Put complex types with double properties
	PutDouble(ctx context.Context, complexBody DoubleWrapper, options *PrimitivePutDoubleOptions) (*http.Response, error)
	// PutDuration - Put complex types with duration properties
	PutDuration(ctx context.Context, complexBody DurationWrapper, options *PrimitivePutDurationOptions) (*http.Response, error)
	// PutFloat - Put complex types with float properties
	PutFloat(ctx context.Context, complexBody FloatWrapper, options *PrimitivePutFloatOptions) (*http.Response, error)
	// PutInt - Put complex types with integer properties
	PutInt(ctx context.Context, complexBody IntWrapper, options *PrimitivePutIntOptions) (*http.Response, error)
	// PutLong - Put complex types with long properties
	PutLong(ctx context.Context, complexBody LongWrapper, options *PrimitivePutLongOptions) (*http.Response, error)
	// PutString - Put complex types with string properties
	PutString(ctx context.Context, complexBody StringWrapper, options *PrimitivePutStringOptions) (*http.Response, error)
}

// PrimitiveClient implements the PrimitiveOperations interface.
// Don't use this type directly, use NewPrimitiveClient() instead.
type PrimitiveClient struct {
	con *Connection
}

// NewPrimitiveClient creates a new instance of PrimitiveClient with the specified values.
func NewPrimitiveClient(con *Connection) PrimitiveOperations {
	return &PrimitiveClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *PrimitiveClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetBool - Get complex types with bool properties
func (client *PrimitiveClient) GetBool(ctx context.Context, options *PrimitiveGetBoolOptions) (*BooleanWrapperResponse, error) {
	req, err := client.GetBoolCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetBoolHandleError(resp)
	}
	result, err := client.GetBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBoolCreateRequest creates the GetBool request.
func (client *PrimitiveClient) GetBoolCreateRequest(ctx context.Context, options *PrimitiveGetBoolOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/bool"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetBoolHandleResponse handles the GetBool response.
func (client *PrimitiveClient) GetBoolHandleResponse(resp *azcore.Response) (*BooleanWrapperResponse, error) {
	result := BooleanWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// GetBoolHandleError handles the GetBool error response.
func (client *PrimitiveClient) GetBoolHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetByte - Get complex types with byte properties
func (client *PrimitiveClient) GetByte(ctx context.Context, options *PrimitiveGetByteOptions) (*ByteWrapperResponse, error) {
	req, err := client.GetByteCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetByteHandleError(resp)
	}
	result, err := client.GetByteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByteCreateRequest creates the GetByte request.
func (client *PrimitiveClient) GetByteCreateRequest(ctx context.Context, options *PrimitiveGetByteOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/byte"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetByteHandleResponse handles the GetByte response.
func (client *PrimitiveClient) GetByteHandleResponse(resp *azcore.Response) (*ByteWrapperResponse, error) {
	result := ByteWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ByteWrapper)
}

// GetByteHandleError handles the GetByte error response.
func (client *PrimitiveClient) GetByteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDate - Get complex types with date properties
func (client *PrimitiveClient) GetDate(ctx context.Context, options *PrimitiveGetDateOptions) (*DateWrapperResponse, error) {
	req, err := client.GetDateCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetDateHandleError(resp)
	}
	result, err := client.GetDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDateCreateRequest creates the GetDate request.
func (client *PrimitiveClient) GetDateCreateRequest(ctx context.Context, options *PrimitiveGetDateOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/date"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDateHandleResponse handles the GetDate response.
func (client *PrimitiveClient) GetDateHandleResponse(resp *azcore.Response) (*DateWrapperResponse, error) {
	result := DateWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DateWrapper)
}

// GetDateHandleError handles the GetDate error response.
func (client *PrimitiveClient) GetDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDateTime - Get complex types with datetime properties
func (client *PrimitiveClient) GetDateTime(ctx context.Context, options *PrimitiveGetDateTimeOptions) (*DatetimeWrapperResponse, error) {
	req, err := client.GetDateTimeCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetDateTimeHandleError(resp)
	}
	result, err := client.GetDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDateTimeCreateRequest creates the GetDateTime request.
func (client *PrimitiveClient) GetDateTimeCreateRequest(ctx context.Context, options *PrimitiveGetDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetime"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDateTimeHandleResponse handles the GetDateTime response.
func (client *PrimitiveClient) GetDateTimeHandleResponse(resp *azcore.Response) (*DatetimeWrapperResponse, error) {
	result := DatetimeWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DatetimeWrapper)
}

// GetDateTimeHandleError handles the GetDateTime error response.
func (client *PrimitiveClient) GetDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDateTimeRFC1123 - Get complex types with datetimeRfc1123 properties
func (client *PrimitiveClient) GetDateTimeRFC1123(ctx context.Context, options *PrimitiveGetDateTimeRFC1123Options) (*Datetimerfc1123WrapperResponse, error) {
	req, err := client.GetDateTimeRFC1123CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetDateTimeRFC1123HandleError(resp)
	}
	result, err := client.GetDateTimeRFC1123HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDateTimeRFC1123CreateRequest creates the GetDateTimeRFC1123 request.
func (client *PrimitiveClient) GetDateTimeRFC1123CreateRequest(ctx context.Context, options *PrimitiveGetDateTimeRFC1123Options) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetimerfc1123"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDateTimeRFC1123HandleResponse handles the GetDateTimeRFC1123 response.
func (client *PrimitiveClient) GetDateTimeRFC1123HandleResponse(resp *azcore.Response) (*Datetimerfc1123WrapperResponse, error) {
	result := Datetimerfc1123WrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Datetimerfc1123Wrapper)
}

// GetDateTimeRFC1123HandleError handles the GetDateTimeRFC1123 error response.
func (client *PrimitiveClient) GetDateTimeRFC1123HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDouble - Get complex types with double properties
func (client *PrimitiveClient) GetDouble(ctx context.Context, options *PrimitiveGetDoubleOptions) (*DoubleWrapperResponse, error) {
	req, err := client.GetDoubleCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetDoubleHandleError(resp)
	}
	result, err := client.GetDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDoubleCreateRequest creates the GetDouble request.
func (client *PrimitiveClient) GetDoubleCreateRequest(ctx context.Context, options *PrimitiveGetDoubleOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/double"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDoubleHandleResponse handles the GetDouble response.
func (client *PrimitiveClient) GetDoubleHandleResponse(resp *azcore.Response) (*DoubleWrapperResponse, error) {
	result := DoubleWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// GetDoubleHandleError handles the GetDouble error response.
func (client *PrimitiveClient) GetDoubleHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDuration - Get complex types with duration properties
func (client *PrimitiveClient) GetDuration(ctx context.Context, options *PrimitiveGetDurationOptions) (*DurationWrapperResponse, error) {
	req, err := client.GetDurationCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetDurationHandleError(resp)
	}
	result, err := client.GetDurationHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDurationCreateRequest creates the GetDuration request.
func (client *PrimitiveClient) GetDurationCreateRequest(ctx context.Context, options *PrimitiveGetDurationOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/duration"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDurationHandleResponse handles the GetDuration response.
func (client *PrimitiveClient) GetDurationHandleResponse(resp *azcore.Response) (*DurationWrapperResponse, error) {
	result := DurationWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DurationWrapper)
}

// GetDurationHandleError handles the GetDuration error response.
func (client *PrimitiveClient) GetDurationHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetFloat - Get complex types with float properties
func (client *PrimitiveClient) GetFloat(ctx context.Context, options *PrimitiveGetFloatOptions) (*FloatWrapperResponse, error) {
	req, err := client.GetFloatCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetFloatHandleError(resp)
	}
	result, err := client.GetFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFloatCreateRequest creates the GetFloat request.
func (client *PrimitiveClient) GetFloatCreateRequest(ctx context.Context, options *PrimitiveGetFloatOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/float"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetFloatHandleResponse handles the GetFloat response.
func (client *PrimitiveClient) GetFloatHandleResponse(resp *azcore.Response) (*FloatWrapperResponse, error) {
	result := FloatWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// GetFloatHandleError handles the GetFloat error response.
func (client *PrimitiveClient) GetFloatHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetInt - Get complex types with integer properties
func (client *PrimitiveClient) GetInt(ctx context.Context, options *PrimitiveGetIntOptions) (*IntWrapperResponse, error) {
	req, err := client.GetIntCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetIntHandleError(resp)
	}
	result, err := client.GetIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetIntCreateRequest creates the GetInt request.
func (client *PrimitiveClient) GetIntCreateRequest(ctx context.Context, options *PrimitiveGetIntOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/integer"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetIntHandleResponse handles the GetInt response.
func (client *PrimitiveClient) GetIntHandleResponse(resp *azcore.Response) (*IntWrapperResponse, error) {
	result := IntWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}

// GetIntHandleError handles the GetInt error response.
func (client *PrimitiveClient) GetIntHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetLong - Get complex types with long properties
func (client *PrimitiveClient) GetLong(ctx context.Context, options *PrimitiveGetLongOptions) (*LongWrapperResponse, error) {
	req, err := client.GetLongCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetLongHandleError(resp)
	}
	result, err := client.GetLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetLongCreateRequest creates the GetLong request.
func (client *PrimitiveClient) GetLongCreateRequest(ctx context.Context, options *PrimitiveGetLongOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/long"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetLongHandleResponse handles the GetLong response.
func (client *PrimitiveClient) GetLongHandleResponse(resp *azcore.Response) (*LongWrapperResponse, error) {
	result := LongWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// GetLongHandleError handles the GetLong error response.
func (client *PrimitiveClient) GetLongHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetString - Get complex types with string properties
func (client *PrimitiveClient) GetString(ctx context.Context, options *PrimitiveGetStringOptions) (*StringWrapperResponse, error) {
	req, err := client.GetStringCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetStringHandleError(resp)
	}
	result, err := client.GetStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStringCreateRequest creates the GetString request.
func (client *PrimitiveClient) GetStringCreateRequest(ctx context.Context, options *PrimitiveGetStringOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/string"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetStringHandleResponse handles the GetString response.
func (client *PrimitiveClient) GetStringHandleResponse(resp *azcore.Response) (*StringWrapperResponse, error) {
	result := StringWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}

// GetStringHandleError handles the GetString error response.
func (client *PrimitiveClient) GetStringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutBool - Put complex types with bool properties
func (client *PrimitiveClient) PutBool(ctx context.Context, complexBody BooleanWrapper, options *PrimitivePutBoolOptions) (*http.Response, error) {
	req, err := client.PutBoolCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutBoolHandleError(resp)
	}
	return resp.Response, nil
}

// PutBoolCreateRequest creates the PutBool request.
func (client *PrimitiveClient) PutBoolCreateRequest(ctx context.Context, complexBody BooleanWrapper, options *PrimitivePutBoolOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/bool"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutBoolHandleError handles the PutBool error response.
func (client *PrimitiveClient) PutBoolHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutByte - Put complex types with byte properties
func (client *PrimitiveClient) PutByte(ctx context.Context, complexBody ByteWrapper, options *PrimitivePutByteOptions) (*http.Response, error) {
	req, err := client.PutByteCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutByteHandleError(resp)
	}
	return resp.Response, nil
}

// PutByteCreateRequest creates the PutByte request.
func (client *PrimitiveClient) PutByteCreateRequest(ctx context.Context, complexBody ByteWrapper, options *PrimitivePutByteOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/byte"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutByteHandleError handles the PutByte error response.
func (client *PrimitiveClient) PutByteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDate - Put complex types with date properties
func (client *PrimitiveClient) PutDate(ctx context.Context, complexBody DateWrapper, options *PrimitivePutDateOptions) (*http.Response, error) {
	req, err := client.PutDateCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutDateHandleError(resp)
	}
	return resp.Response, nil
}

// PutDateCreateRequest creates the PutDate request.
func (client *PrimitiveClient) PutDateCreateRequest(ctx context.Context, complexBody DateWrapper, options *PrimitivePutDateOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/date"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutDateHandleError handles the PutDate error response.
func (client *PrimitiveClient) PutDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDateTime - Put complex types with datetime properties
func (client *PrimitiveClient) PutDateTime(ctx context.Context, complexBody DatetimeWrapper, options *PrimitivePutDateTimeOptions) (*http.Response, error) {
	req, err := client.PutDateTimeCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// PutDateTimeCreateRequest creates the PutDateTime request.
func (client *PrimitiveClient) PutDateTimeCreateRequest(ctx context.Context, complexBody DatetimeWrapper, options *PrimitivePutDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetime"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutDateTimeHandleError handles the PutDateTime error response.
func (client *PrimitiveClient) PutDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDateTimeRFC1123 - Put complex types with datetimeRfc1123 properties
func (client *PrimitiveClient) PutDateTimeRFC1123(ctx context.Context, complexBody Datetimerfc1123Wrapper, options *PrimitivePutDateTimeRFC1123Options) (*http.Response, error) {
	req, err := client.PutDateTimeRFC1123CreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutDateTimeRFC1123HandleError(resp)
	}
	return resp.Response, nil
}

// PutDateTimeRFC1123CreateRequest creates the PutDateTimeRFC1123 request.
func (client *PrimitiveClient) PutDateTimeRFC1123CreateRequest(ctx context.Context, complexBody Datetimerfc1123Wrapper, options *PrimitivePutDateTimeRFC1123Options) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetimerfc1123"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutDateTimeRFC1123HandleError handles the PutDateTimeRFC1123 error response.
func (client *PrimitiveClient) PutDateTimeRFC1123HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDouble - Put complex types with double properties
func (client *PrimitiveClient) PutDouble(ctx context.Context, complexBody DoubleWrapper, options *PrimitivePutDoubleOptions) (*http.Response, error) {
	req, err := client.PutDoubleCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutDoubleHandleError(resp)
	}
	return resp.Response, nil
}

// PutDoubleCreateRequest creates the PutDouble request.
func (client *PrimitiveClient) PutDoubleCreateRequest(ctx context.Context, complexBody DoubleWrapper, options *PrimitivePutDoubleOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/double"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutDoubleHandleError handles the PutDouble error response.
func (client *PrimitiveClient) PutDoubleHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDuration - Put complex types with duration properties
func (client *PrimitiveClient) PutDuration(ctx context.Context, complexBody DurationWrapper, options *PrimitivePutDurationOptions) (*http.Response, error) {
	req, err := client.PutDurationCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutDurationHandleError(resp)
	}
	return resp.Response, nil
}

// PutDurationCreateRequest creates the PutDuration request.
func (client *PrimitiveClient) PutDurationCreateRequest(ctx context.Context, complexBody DurationWrapper, options *PrimitivePutDurationOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/duration"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutDurationHandleError handles the PutDuration error response.
func (client *PrimitiveClient) PutDurationHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutFloat - Put complex types with float properties
func (client *PrimitiveClient) PutFloat(ctx context.Context, complexBody FloatWrapper, options *PrimitivePutFloatOptions) (*http.Response, error) {
	req, err := client.PutFloatCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutFloatHandleError(resp)
	}
	return resp.Response, nil
}

// PutFloatCreateRequest creates the PutFloat request.
func (client *PrimitiveClient) PutFloatCreateRequest(ctx context.Context, complexBody FloatWrapper, options *PrimitivePutFloatOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/float"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutFloatHandleError handles the PutFloat error response.
func (client *PrimitiveClient) PutFloatHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutInt - Put complex types with integer properties
func (client *PrimitiveClient) PutInt(ctx context.Context, complexBody IntWrapper, options *PrimitivePutIntOptions) (*http.Response, error) {
	req, err := client.PutIntCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutIntHandleError(resp)
	}
	return resp.Response, nil
}

// PutIntCreateRequest creates the PutInt request.
func (client *PrimitiveClient) PutIntCreateRequest(ctx context.Context, complexBody IntWrapper, options *PrimitivePutIntOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/integer"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutIntHandleError handles the PutInt error response.
func (client *PrimitiveClient) PutIntHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutLong - Put complex types with long properties
func (client *PrimitiveClient) PutLong(ctx context.Context, complexBody LongWrapper, options *PrimitivePutLongOptions) (*http.Response, error) {
	req, err := client.PutLongCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutLongHandleError(resp)
	}
	return resp.Response, nil
}

// PutLongCreateRequest creates the PutLong request.
func (client *PrimitiveClient) PutLongCreateRequest(ctx context.Context, complexBody LongWrapper, options *PrimitivePutLongOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/long"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutLongHandleError handles the PutLong error response.
func (client *PrimitiveClient) PutLongHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutString - Put complex types with string properties
func (client *PrimitiveClient) PutString(ctx context.Context, complexBody StringWrapper, options *PrimitivePutStringOptions) (*http.Response, error) {
	req, err := client.PutStringCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutStringHandleError(resp)
	}
	return resp.Response, nil
}

// PutStringCreateRequest creates the PutString request.
func (client *PrimitiveClient) PutStringCreateRequest(ctx context.Context, complexBody StringWrapper, options *PrimitivePutStringOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/string"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutStringHandleError handles the PutString error response.
func (client *PrimitiveClient) PutStringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}