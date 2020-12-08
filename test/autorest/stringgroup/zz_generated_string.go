// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// StringClient contains the methods for the String group.
// Don't use this type directly, use NewStringClient() instead.
type StringClient struct {
	con *Connection
}

// NewStringClient creates a new instance of StringClient with the specified values.
func NewStringClient(con *Connection) StringClient {
	return StringClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client StringClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetBase64Encoded - Get value that is base64 encoded
func (client StringClient) GetBase64Encoded(ctx context.Context, options *StringGetBase64EncodedOptions) (ByteArrayResponse, error) {
	req, err := client.getBase64EncodedCreateRequest(ctx, options)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ByteArrayResponse{}, client.getBase64EncodedHandleError(resp)
	}
	return client.getBase64EncodedHandleResponse(resp)
}

// getBase64EncodedCreateRequest creates the GetBase64Encoded request.
func (client StringClient) getBase64EncodedCreateRequest(ctx context.Context, options *StringGetBase64EncodedOptions) (*azcore.Request, error) {
	urlPath := "/string/base64Encoding"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getBase64EncodedHandleResponse handles the GetBase64Encoded response.
func (client StringClient) getBase64EncodedHandleResponse(resp *azcore.Response) (ByteArrayResponse, error) {
	var val *[]byte
	if err := resp.UnmarshalAsByteArray(&val, azcore.Base64StdFormat); err != nil {
		return ByteArrayResponse{}, err
	}
	return ByteArrayResponse{RawResponse: resp.Response, Value: val}, nil
}

// getBase64EncodedHandleError handles the GetBase64Encoded error response.
func (client StringClient) getBase64EncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetBase64URLEncoded - Get value that is base64url encoded
func (client StringClient) GetBase64URLEncoded(ctx context.Context, options *StringGetBase64URLEncodedOptions) (ByteArrayResponse, error) {
	req, err := client.getBase64UrlEncodedCreateRequest(ctx, options)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ByteArrayResponse{}, client.getBase64UrlEncodedHandleError(resp)
	}
	return client.getBase64UrlEncodedHandleResponse(resp)
}

// getBase64UrlEncodedCreateRequest creates the GetBase64URLEncoded request.
func (client StringClient) getBase64UrlEncodedCreateRequest(ctx context.Context, options *StringGetBase64URLEncodedOptions) (*azcore.Request, error) {
	urlPath := "/string/base64UrlEncoding"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getBase64UrlEncodedHandleResponse handles the GetBase64URLEncoded response.
func (client StringClient) getBase64UrlEncodedHandleResponse(resp *azcore.Response) (ByteArrayResponse, error) {
	var val *[]byte
	if err := resp.UnmarshalAsByteArray(&val, azcore.Base64URLFormat); err != nil {
		return ByteArrayResponse{}, err
	}
	return ByteArrayResponse{RawResponse: resp.Response, Value: val}, nil
}

// getBase64UrlEncodedHandleError handles the GetBase64URLEncoded error response.
func (client StringClient) getBase64UrlEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetEmpty - Get empty string value value ''
func (client StringClient) GetEmpty(ctx context.Context, options *StringGetEmptyOptions) (StringResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.getEmptyHandleError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client StringClient) getEmptyCreateRequest(ctx context.Context, options *StringGetEmptyOptions) (*azcore.Request, error) {
	urlPath := "/string/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client StringClient) getEmptyHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client StringClient) getEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetMBCS - Get mbcs string value '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
func (client StringClient) GetMBCS(ctx context.Context, options *StringGetMBCSOptions) (StringResponse, error) {
	req, err := client.getMbcsCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.getMbcsHandleError(resp)
	}
	return client.getMbcsHandleResponse(resp)
}

// getMbcsCreateRequest creates the GetMBCS request.
func (client StringClient) getMbcsCreateRequest(ctx context.Context, options *StringGetMBCSOptions) (*azcore.Request, error) {
	urlPath := "/string/mbcs"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getMbcsHandleResponse handles the GetMBCS response.
func (client StringClient) getMbcsHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// getMbcsHandleError handles the GetMBCS error response.
func (client StringClient) getMbcsHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotProvided - Get String value when no string value is sent in response payload
func (client StringClient) GetNotProvided(ctx context.Context, options *StringGetNotProvidedOptions) (StringResponse, error) {
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.getNotProvidedHandleError(resp)
	}
	return client.getNotProvidedHandleResponse(resp)
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client StringClient) getNotProvidedCreateRequest(ctx context.Context, options *StringGetNotProvidedOptions) (*azcore.Request, error) {
	urlPath := "/string/notProvided"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client StringClient) getNotProvidedHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// getNotProvidedHandleError handles the GetNotProvided error response.
func (client StringClient) getNotProvidedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get null string value value
func (client StringClient) GetNull(ctx context.Context, options *StringGetNullOptions) (StringResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client StringClient) getNullCreateRequest(ctx context.Context, options *StringGetNullOptions) (*azcore.Request, error) {
	urlPath := "/string/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client StringClient) getNullHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// getNullHandleError handles the GetNull error response.
func (client StringClient) getNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNullBase64URLEncoded - Get null value that is expected to be base64url encoded
func (client StringClient) GetNullBase64URLEncoded(ctx context.Context, options *StringGetNullBase64URLEncodedOptions) (ByteArrayResponse, error) {
	req, err := client.getNullBase64UrlEncodedCreateRequest(ctx, options)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ByteArrayResponse{}, client.getNullBase64UrlEncodedHandleError(resp)
	}
	return client.getNullBase64UrlEncodedHandleResponse(resp)
}

// getNullBase64UrlEncodedCreateRequest creates the GetNullBase64URLEncoded request.
func (client StringClient) getNullBase64UrlEncodedCreateRequest(ctx context.Context, options *StringGetNullBase64URLEncodedOptions) (*azcore.Request, error) {
	urlPath := "/string/nullBase64UrlEncoding"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullBase64UrlEncodedHandleResponse handles the GetNullBase64URLEncoded response.
func (client StringClient) getNullBase64UrlEncodedHandleResponse(resp *azcore.Response) (ByteArrayResponse, error) {
	var val *[]byte
	if err := resp.UnmarshalAsByteArray(&val, azcore.Base64URLFormat); err != nil {
		return ByteArrayResponse{}, err
	}
	return ByteArrayResponse{RawResponse: resp.Response, Value: val}, nil
}

// getNullBase64UrlEncodedHandleError handles the GetNullBase64URLEncoded error response.
func (client StringClient) getNullBase64UrlEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetWhitespace - Get string value with leading and trailing whitespace 'Now is the time for all good men to come to the aid of their country'
func (client StringClient) GetWhitespace(ctx context.Context, options *StringGetWhitespaceOptions) (StringResponse, error) {
	req, err := client.getWhitespaceCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.getWhitespaceHandleError(resp)
	}
	return client.getWhitespaceHandleResponse(resp)
}

// getWhitespaceCreateRequest creates the GetWhitespace request.
func (client StringClient) getWhitespaceCreateRequest(ctx context.Context, options *StringGetWhitespaceOptions) (*azcore.Request, error) {
	urlPath := "/string/whitespace"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getWhitespaceHandleResponse handles the GetWhitespace response.
func (client StringClient) getWhitespaceHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// getWhitespaceHandleError handles the GetWhitespace error response.
func (client StringClient) getWhitespaceHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutBase64URLEncoded - Put value that is base64url encoded
func (client StringClient) PutBase64URLEncoded(ctx context.Context, stringBody []byte, options *StringPutBase64URLEncodedOptions) (*http.Response, error) {
	req, err := client.putBase64UrlEncodedCreateRequest(ctx, stringBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putBase64UrlEncodedHandleError(resp)
	}
	return resp.Response, nil
}

// putBase64UrlEncodedCreateRequest creates the PutBase64URLEncoded request.
func (client StringClient) putBase64UrlEncodedCreateRequest(ctx context.Context, stringBody []byte, options *StringPutBase64URLEncodedOptions) (*azcore.Request, error) {
	urlPath := "/string/base64UrlEncoding"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsByteArray(stringBody, azcore.Base64URLFormat)
}

// putBase64UrlEncodedHandleError handles the PutBase64URLEncoded error response.
func (client StringClient) putBase64UrlEncodedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutEmpty - Set string value empty ''
func (client StringClient) PutEmpty(ctx context.Context, options *StringPutEmptyOptions) (*http.Response, error) {
	req, err := client.putEmptyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client StringClient) putEmptyCreateRequest(ctx context.Context, options *StringPutEmptyOptions) (*azcore.Request, error) {
	urlPath := "/string/empty"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON("")
}

// putEmptyHandleError handles the PutEmpty error response.
func (client StringClient) putEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutMBCS - Set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
func (client StringClient) PutMBCS(ctx context.Context, options *StringPutMBCSOptions) (*http.Response, error) {
	req, err := client.putMbcsCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putMbcsHandleError(resp)
	}
	return resp.Response, nil
}

// putMbcsCreateRequest creates the PutMBCS request.
func (client StringClient) putMbcsCreateRequest(ctx context.Context, options *StringPutMBCSOptions) (*azcore.Request, error) {
	urlPath := "/string/mbcs"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON("啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€")
}

// putMbcsHandleError handles the PutMBCS error response.
func (client StringClient) putMbcsHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutNull - Set string value null
func (client StringClient) PutNull(ctx context.Context, options *StringPutNullOptions) (*http.Response, error) {
	req, err := client.putNullCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putNullHandleError(resp)
	}
	return resp.Response, nil
}

// putNullCreateRequest creates the PutNull request.
func (client StringClient) putNullCreateRequest(ctx context.Context, options *StringPutNullOptions) (*azcore.Request, error) {
	urlPath := "/string/null"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.StringBody)
	}
	return req, nil
}

// putNullHandleError handles the PutNull error response.
func (client StringClient) putNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutWhitespace - Set String value with leading and trailing whitespace 'Now is the time for all good men to come to the aid of their country'
func (client StringClient) PutWhitespace(ctx context.Context, options *StringPutWhitespaceOptions) (*http.Response, error) {
	req, err := client.putWhitespaceCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putWhitespaceHandleError(resp)
	}
	return resp.Response, nil
}

// putWhitespaceCreateRequest creates the PutWhitespace request.
func (client StringClient) putWhitespaceCreateRequest(ctx context.Context, options *StringPutWhitespaceOptions) (*azcore.Request, error) {
	urlPath := "/string/whitespace"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON("    Now is the time for all good men to come to the aid of their country    ")
}

// putWhitespaceHandleError handles the PutWhitespace error response.
func (client StringClient) putWhitespaceHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
