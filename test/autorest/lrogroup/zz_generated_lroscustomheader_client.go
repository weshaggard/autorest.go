// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// LROsCustomHeaderClient contains the methods for the LROsCustomHeader group.
// Don't use this type directly, use NewLROsCustomHeaderClient() instead.
type LROsCustomHeaderClient struct {
	con *Connection
}

// NewLROsCustomHeaderClient creates a new instance of LROsCustomHeaderClient with the specified values.
func NewLROsCustomHeaderClient(con *Connection) *LROsCustomHeaderClient {
	return &LROsCustomHeaderClient{con: con}
}

// BeginPost202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request,
// service returns a 202 to the initial request, with 'Location' and
// 'Retry-After' headers, Polls return a 200 with a response body after success
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) BeginPost202Retry200(ctx context.Context, options *LROsCustomHeaderBeginPost202Retry200Options) (HTTPPollerResponse, error) {
	resp, err := client.post202Retry200(ctx, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LROsCustomHeaderClient.Post202Retry200", "", resp, client.con.Pipeline(), client.post202Retry200HandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePost202Retry200 creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *LROsCustomHeaderClient) ResumePost202Retry200(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LROsCustomHeaderClient.Post202Retry200", token, client.con.Pipeline(), client.post202Retry200HandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Post202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request,
// service returns a 202 to the initial request, with 'Location' and
// 'Retry-After' headers, Polls return a 200 with a response body after success
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) post202Retry200(ctx context.Context, options *LROsCustomHeaderBeginPost202Retry200Options) (*azcore.Response, error) {
	req, err := client.post202Retry200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.post202Retry200HandleError(resp)
	}
	return resp, nil
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *LROsCustomHeaderClient) post202Retry200CreateRequest(ctx context.Context, options *LROsCustomHeaderBeginPost202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/customheader/post/202/retry/200"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, req.MarshalAsJSON(*options.Product)
	}
	return req, nil
}

// post202Retry200HandleError handles the Post202Retry200 error response.
func (client *LROsCustomHeaderClient) post202Retry200HandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginPostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// post request, service returns a 202 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) BeginPostAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderBeginPostAsyncRetrySucceededOptions) (HTTPPollerResponse, error) {
	resp, err := client.postAsyncRetrySucceeded(ctx, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LROsCustomHeaderClient.PostAsyncRetrySucceeded", "", resp, client.con.Pipeline(), client.postAsyncRetrySucceededHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePostAsyncRetrySucceeded creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *LROsCustomHeaderClient) ResumePostAsyncRetrySucceeded(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LROsCustomHeaderClient.PostAsyncRetrySucceeded", token, client.con.Pipeline(), client.postAsyncRetrySucceededHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// PostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post
// request, service returns a 202 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) postAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderBeginPostAsyncRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.postAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.postAsyncRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// postAsyncRetrySucceededCreateRequest creates the PostAsyncRetrySucceeded request.
func (client *LROsCustomHeaderClient) postAsyncRetrySucceededCreateRequest(ctx context.Context, options *LROsCustomHeaderBeginPostAsyncRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/customheader/postasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, req.MarshalAsJSON(*options.Product)
	}
	return req, nil
}

// postAsyncRetrySucceededHandleError handles the PostAsyncRetrySucceeded error response.
func (client *LROsCustomHeaderClient) postAsyncRetrySucceededHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginPut201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) BeginPut201CreatingSucceeded200(ctx context.Context, options *LROsCustomHeaderBeginPut201CreatingSucceeded200Options) (ProductPollerResponse, error) {
	resp, err := client.put201CreatingSucceeded200(ctx, options)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LROsCustomHeaderClient.Put201CreatingSucceeded200", "", resp, client.con.Pipeline(), client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	poller := &productPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePut201CreatingSucceeded200 creates a new ProductPoller from the specified resume token.
// token - The value must come from a previous call to ProductPoller.ResumeToken().
func (client *LROsCustomHeaderClient) ResumePut201CreatingSucceeded200(ctx context.Context, token string) (ProductPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LROsCustomHeaderClient.Put201CreatingSucceeded200", token, client.con.Pipeline(), client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	poller := &productPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Put201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 201 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) put201CreatingSucceeded200(ctx context.Context, options *LROsCustomHeaderBeginPut201CreatingSucceeded200Options) (*azcore.Response, error) {
	req, err := client.put201CreatingSucceeded200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.put201CreatingSucceeded200HandleError(resp)
	}
	return resp, nil
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *LROsCustomHeaderClient) put201CreatingSucceeded200CreateRequest(ctx context.Context, options *LROsCustomHeaderBeginPut201CreatingSucceeded200Options) (*azcore.Request, error) {
	urlPath := "/lro/customheader/put/201/creating/succeeded/200"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, req.MarshalAsJSON(*options.Product)
	}
	return req, nil
}

// put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client *LROsCustomHeaderClient) put201CreatingSucceeded200HandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginPutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running
// put request, service returns a 200 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) BeginPutAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderBeginPutAsyncRetrySucceededOptions) (ProductPollerResponse, error) {
	resp, err := client.putAsyncRetrySucceeded(ctx, options)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LROsCustomHeaderClient.PutAsyncRetrySucceeded", "", resp, client.con.Pipeline(), client.putAsyncRetrySucceededHandleError)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	poller := &productPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePutAsyncRetrySucceeded creates a new ProductPoller from the specified resume token.
// token - The value must come from a previous call to ProductPoller.ResumeToken().
func (client *LROsCustomHeaderClient) ResumePutAsyncRetrySucceeded(ctx context.Context, token string) (ProductPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("LROsCustomHeaderClient.PutAsyncRetrySucceeded", token, client.con.Pipeline(), client.putAsyncRetrySucceededHandleError)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	poller := &productPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ProductPollerResponse{}, err
	}
	result := ProductPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// PutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put
// request, service returns a 200 to the initial request, with an entity that
// contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
// If the operation fails it returns the *CloudError error type.
func (client *LROsCustomHeaderClient) putAsyncRetrySucceeded(ctx context.Context, options *LROsCustomHeaderBeginPutAsyncRetrySucceededOptions) (*azcore.Response, error) {
	req, err := client.putAsyncRetrySucceededCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putAsyncRetrySucceededHandleError(resp)
	}
	return resp, nil
}

// putAsyncRetrySucceededCreateRequest creates the PutAsyncRetrySucceeded request.
func (client *LROsCustomHeaderClient) putAsyncRetrySucceededCreateRequest(ctx context.Context, options *LROsCustomHeaderBeginPutAsyncRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/customheader/putasync/retry/succeeded"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Product != nil {
		return req, req.MarshalAsJSON(*options.Product)
	}
	return req, nil
}

// putAsyncRetrySucceededHandleError handles the PutAsyncRetrySucceeded error response.
func (client *LROsCustomHeaderClient) putAsyncRetrySucceededHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
