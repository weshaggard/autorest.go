// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type sparkSessionClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client sparkSessionClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CancelSparkSession - Cancels a running spark session.
func (client sparkSessionClient) CancelSparkSession(ctx context.Context, sessionId int32, options *SparkSessionCancelSparkSessionOptions) (*http.Response, error) {
	req, err := client.cancelSparkSessionCreateRequest(ctx, sessionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.cancelSparkSessionHandleError(resp)
	}
	return resp.Response, nil
}

// cancelSparkSessionCreateRequest creates the CancelSparkSession request.
func (client sparkSessionClient) cancelSparkSessionCreateRequest(ctx context.Context, sessionId int32, options *SparkSessionCancelSparkSessionOptions) (*azcore.Request, error) {
	urlPath := "/sessions/{sessionId}"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	return req, nil
}

// cancelSparkSessionHandleError handles the CancelSparkSession error response.
func (client sparkSessionClient) cancelSparkSessionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// CancelSparkStatement - Kill a statement within a session.
func (client sparkSessionClient) CancelSparkStatement(ctx context.Context, sessionId int32, statementId int32, options *SparkSessionCancelSparkStatementOptions) (SparkStatementCancellationResultResponse, error) {
	req, err := client.cancelSparkStatementCreateRequest(ctx, sessionId, statementId, options)
	if err != nil {
		return SparkStatementCancellationResultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SparkStatementCancellationResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkStatementCancellationResultResponse{}, client.cancelSparkStatementHandleError(resp)
	}
	return client.cancelSparkStatementHandleResponse(resp)
}

// cancelSparkStatementCreateRequest creates the CancelSparkStatement request.
func (client sparkSessionClient) cancelSparkStatementCreateRequest(ctx context.Context, sessionId int32, statementId int32, options *SparkSessionCancelSparkStatementOptions) (*azcore.Request, error) {
	urlPath := "/sessions/{sessionId}/statements/{statementId}/cancel"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionId), 10)))
	urlPath = strings.ReplaceAll(urlPath, "{statementId}", url.PathEscape(strconv.FormatInt(int64(statementId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// cancelSparkStatementHandleResponse handles the CancelSparkStatement response.
func (client sparkSessionClient) cancelSparkStatementHandleResponse(resp *azcore.Response) (SparkStatementCancellationResultResponse, error) {
	var val *SparkStatementCancellationResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkStatementCancellationResultResponse{}, err
	}
	return SparkStatementCancellationResultResponse{RawResponse: resp.Response, SparkStatementCancellationResult: val}, nil
}

// cancelSparkStatementHandleError handles the CancelSparkStatement error response.
func (client sparkSessionClient) cancelSparkStatementHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// CreateSparkSession - Create new spark session.
func (client sparkSessionClient) CreateSparkSession(ctx context.Context, sparkSessionOptions SparkSessionOptions, options *SparkSessionCreateSparkSessionOptions) (SparkSessionResponse, error) {
	req, err := client.createSparkSessionCreateRequest(ctx, sparkSessionOptions, options)
	if err != nil {
		return SparkSessionResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SparkSessionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkSessionResponse{}, client.createSparkSessionHandleError(resp)
	}
	return client.createSparkSessionHandleResponse(resp)
}

// createSparkSessionCreateRequest creates the CreateSparkSession request.
func (client sparkSessionClient) createSparkSessionCreateRequest(ctx context.Context, sparkSessionOptions SparkSessionOptions, options *SparkSessionCreateSparkSessionOptions) (*azcore.Request, error) {
	urlPath := "/sessions"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Detailed != nil {
		query.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkSessionOptions)
}

// createSparkSessionHandleResponse handles the CreateSparkSession response.
func (client sparkSessionClient) createSparkSessionHandleResponse(resp *azcore.Response) (SparkSessionResponse, error) {
	var val *SparkSession
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkSessionResponse{}, err
	}
	return SparkSessionResponse{RawResponse: resp.Response, SparkSession: val}, nil
}

// createSparkSessionHandleError handles the CreateSparkSession error response.
func (client sparkSessionClient) createSparkSessionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// CreateSparkStatement - Create statement within a spark session.
func (client sparkSessionClient) CreateSparkStatement(ctx context.Context, sessionId int32, sparkStatementOptions SparkStatementOptions, options *SparkSessionCreateSparkStatementOptions) (SparkStatementResponse, error) {
	req, err := client.createSparkStatementCreateRequest(ctx, sessionId, sparkStatementOptions, options)
	if err != nil {
		return SparkStatementResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SparkStatementResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkStatementResponse{}, client.createSparkStatementHandleError(resp)
	}
	return client.createSparkStatementHandleResponse(resp)
}

// createSparkStatementCreateRequest creates the CreateSparkStatement request.
func (client sparkSessionClient) createSparkStatementCreateRequest(ctx context.Context, sessionId int32, sparkStatementOptions SparkStatementOptions, options *SparkSessionCreateSparkStatementOptions) (*azcore.Request, error) {
	urlPath := "/sessions/{sessionId}/statements"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sparkStatementOptions)
}

// createSparkStatementHandleResponse handles the CreateSparkStatement response.
func (client sparkSessionClient) createSparkStatementHandleResponse(resp *azcore.Response) (SparkStatementResponse, error) {
	var val *SparkStatement
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkStatementResponse{}, err
	}
	return SparkStatementResponse{RawResponse: resp.Response, SparkStatement: val}, nil
}

// createSparkStatementHandleError handles the CreateSparkStatement error response.
func (client sparkSessionClient) createSparkStatementHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetSparkSession - Gets a single spark session.
func (client sparkSessionClient) GetSparkSession(ctx context.Context, sessionId int32, options *SparkSessionGetSparkSessionOptions) (SparkSessionResponse, error) {
	req, err := client.getSparkSessionCreateRequest(ctx, sessionId, options)
	if err != nil {
		return SparkSessionResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SparkSessionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkSessionResponse{}, client.getSparkSessionHandleError(resp)
	}
	return client.getSparkSessionHandleResponse(resp)
}

// getSparkSessionCreateRequest creates the GetSparkSession request.
func (client sparkSessionClient) getSparkSessionCreateRequest(ctx context.Context, sessionId int32, options *SparkSessionGetSparkSessionOptions) (*azcore.Request, error) {
	urlPath := "/sessions/{sessionId}"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Detailed != nil {
		query.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkSessionHandleResponse handles the GetSparkSession response.
func (client sparkSessionClient) getSparkSessionHandleResponse(resp *azcore.Response) (SparkSessionResponse, error) {
	var val *SparkSession
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkSessionResponse{}, err
	}
	return SparkSessionResponse{RawResponse: resp.Response, SparkSession: val}, nil
}

// getSparkSessionHandleError handles the GetSparkSession error response.
func (client sparkSessionClient) getSparkSessionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetSparkSessions - List all spark sessions which are running under a particular spark pool.
func (client sparkSessionClient) GetSparkSessions(ctx context.Context, options *SparkSessionGetSparkSessionsOptions) (SparkSessionCollectionResponse, error) {
	req, err := client.getSparkSessionsCreateRequest(ctx, options)
	if err != nil {
		return SparkSessionCollectionResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SparkSessionCollectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkSessionCollectionResponse{}, client.getSparkSessionsHandleError(resp)
	}
	return client.getSparkSessionsHandleResponse(resp)
}

// getSparkSessionsCreateRequest creates the GetSparkSessions request.
func (client sparkSessionClient) getSparkSessionsCreateRequest(ctx context.Context, options *SparkSessionGetSparkSessionsOptions) (*azcore.Request, error) {
	urlPath := "/sessions"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.From != nil {
		query.Set("from", strconv.FormatInt(int64(*options.From), 10))
	}
	if options != nil && options.Size != nil {
		query.Set("size", strconv.FormatInt(int64(*options.Size), 10))
	}
	if options != nil && options.Detailed != nil {
		query.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkSessionsHandleResponse handles the GetSparkSessions response.
func (client sparkSessionClient) getSparkSessionsHandleResponse(resp *azcore.Response) (SparkSessionCollectionResponse, error) {
	var val *SparkSessionCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkSessionCollectionResponse{}, err
	}
	return SparkSessionCollectionResponse{RawResponse: resp.Response, SparkSessionCollection: val}, nil
}

// getSparkSessionsHandleError handles the GetSparkSessions error response.
func (client sparkSessionClient) getSparkSessionsHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetSparkStatement - Gets a single statement within a spark session.
func (client sparkSessionClient) GetSparkStatement(ctx context.Context, sessionId int32, statementId int32, options *SparkSessionGetSparkStatementOptions) (SparkStatementResponse, error) {
	req, err := client.getSparkStatementCreateRequest(ctx, sessionId, statementId, options)
	if err != nil {
		return SparkStatementResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SparkStatementResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkStatementResponse{}, client.getSparkStatementHandleError(resp)
	}
	return client.getSparkStatementHandleResponse(resp)
}

// getSparkStatementCreateRequest creates the GetSparkStatement request.
func (client sparkSessionClient) getSparkStatementCreateRequest(ctx context.Context, sessionId int32, statementId int32, options *SparkSessionGetSparkStatementOptions) (*azcore.Request, error) {
	urlPath := "/sessions/{sessionId}/statements/{statementId}"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionId), 10)))
	urlPath = strings.ReplaceAll(urlPath, "{statementId}", url.PathEscape(strconv.FormatInt(int64(statementId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkStatementHandleResponse handles the GetSparkStatement response.
func (client sparkSessionClient) getSparkStatementHandleResponse(resp *azcore.Response) (SparkStatementResponse, error) {
	var val *SparkStatement
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkStatementResponse{}, err
	}
	return SparkStatementResponse{RawResponse: resp.Response, SparkStatement: val}, nil
}

// getSparkStatementHandleError handles the GetSparkStatement error response.
func (client sparkSessionClient) getSparkStatementHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetSparkStatements - Gets a list of statements within a spark session.
func (client sparkSessionClient) GetSparkStatements(ctx context.Context, sessionId int32, options *SparkSessionGetSparkStatementsOptions) (SparkStatementCollectionResponse, error) {
	req, err := client.getSparkStatementsCreateRequest(ctx, sessionId, options)
	if err != nil {
		return SparkStatementCollectionResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SparkStatementCollectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SparkStatementCollectionResponse{}, client.getSparkStatementsHandleError(resp)
	}
	return client.getSparkStatementsHandleResponse(resp)
}

// getSparkStatementsCreateRequest creates the GetSparkStatements request.
func (client sparkSessionClient) getSparkStatementsCreateRequest(ctx context.Context, sessionId int32, options *SparkSessionGetSparkStatementsOptions) (*azcore.Request, error) {
	urlPath := "/sessions/{sessionId}/statements"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSparkStatementsHandleResponse handles the GetSparkStatements response.
func (client sparkSessionClient) getSparkStatementsHandleResponse(resp *azcore.Response) (SparkStatementCollectionResponse, error) {
	var val *SparkStatementCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SparkStatementCollectionResponse{}, err
	}
	return SparkStatementCollectionResponse{RawResponse: resp.Response, SparkStatementCollection: val}, nil
}

// getSparkStatementsHandleError handles the GetSparkStatements error response.
func (client sparkSessionClient) getSparkStatementsHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ResetSparkSessionTimeout - Sends a keep alive call to the current session to reset the session timeout.
func (client sparkSessionClient) ResetSparkSessionTimeout(ctx context.Context, sessionId int32, options *SparkSessionResetSparkSessionTimeoutOptions) (*http.Response, error) {
	req, err := client.resetSparkSessionTimeoutCreateRequest(ctx, sessionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.resetSparkSessionTimeoutHandleError(resp)
	}
	return resp.Response, nil
}

// resetSparkSessionTimeoutCreateRequest creates the ResetSparkSessionTimeout request.
func (client sparkSessionClient) resetSparkSessionTimeoutCreateRequest(ctx context.Context, sessionId int32, options *SparkSessionResetSparkSessionTimeoutOptions) (*azcore.Request, error) {
	urlPath := "/sessions/{sessionId}/reset-timeout"
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionId), 10)))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	return req, nil
}

// resetSparkSessionTimeoutHandleError handles the ResetSparkSessionTimeout error response.
func (client sparkSessionClient) resetSparkSessionTimeoutHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
