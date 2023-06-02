//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"azspark"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"regexp"
	"strconv"
)

// BatchServer is a fake server for instances of the azspark.BatchClient type.
type BatchServer struct {
	// CancelSparkBatchJob is the fake for method BatchClient.CancelSparkBatchJob
	// HTTP status codes to indicate success: http.StatusOK
	CancelSparkBatchJob func(ctx context.Context, batchID int32, options *azspark.BatchClientCancelSparkBatchJobOptions) (resp azfake.Responder[azspark.BatchClientCancelSparkBatchJobResponse], errResp azfake.ErrorResponder)

	// CreateSparkBatchJob is the fake for method BatchClient.CreateSparkBatchJob
	// HTTP status codes to indicate success: http.StatusOK
	CreateSparkBatchJob func(ctx context.Context, sparkBatchJobOptions azspark.BatchJobOptions, options *azspark.BatchClientCreateSparkBatchJobOptions) (resp azfake.Responder[azspark.BatchClientCreateSparkBatchJobResponse], errResp azfake.ErrorResponder)

	// GetSparkBatchJob is the fake for method BatchClient.GetSparkBatchJob
	// HTTP status codes to indicate success: http.StatusOK
	GetSparkBatchJob func(ctx context.Context, batchID int32, options *azspark.BatchClientGetSparkBatchJobOptions) (resp azfake.Responder[azspark.BatchClientGetSparkBatchJobResponse], errResp azfake.ErrorResponder)

	// GetSparkBatchJobs is the fake for method BatchClient.GetSparkBatchJobs
	// HTTP status codes to indicate success: http.StatusOK
	GetSparkBatchJobs func(ctx context.Context, options *azspark.BatchClientGetSparkBatchJobsOptions) (resp azfake.Responder[azspark.BatchClientGetSparkBatchJobsResponse], errResp azfake.ErrorResponder)
}

// NewBatchServerTransport creates a new instance of BatchServerTransport with the provided implementation.
// The returned BatchServerTransport instance is connected to an instance of azspark.BatchClient by way of the
// undefined.Transporter field.
func NewBatchServerTransport(srv *BatchServer) *BatchServerTransport {
	return &BatchServerTransport{srv: srv}
}

// BatchServerTransport connects instances of azspark.BatchClient to instances of BatchServer.
// Don't use this type directly, use NewBatchServerTransport instead.
type BatchServerTransport struct {
	srv *BatchServer
}

// Do implements the policy.Transporter interface for BatchServerTransport.
func (b *BatchServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BatchClient.CancelSparkBatchJob":
		resp, err = b.dispatchCancelSparkBatchJob(req)
	case "BatchClient.CreateSparkBatchJob":
		resp, err = b.dispatchCreateSparkBatchJob(req)
	case "BatchClient.GetSparkBatchJob":
		resp, err = b.dispatchGetSparkBatchJob(req)
	case "BatchClient.GetSparkBatchJobs":
		resp, err = b.dispatchGetSparkBatchJobs(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BatchServerTransport) dispatchCancelSparkBatchJob(req *http.Request) (*http.Response, error) {
	if b.srv.CancelSparkBatchJob == nil {
		return nil, &nonRetriableError{errors.New("method CancelSparkBatchJob not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/batches/(?P<batchId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	batchIDParam, err := parseWithCast(matches[regex.SubexpIndex("batchId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.CancelSparkBatchJob(req.Context(), int32(batchIDParam), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BatchServerTransport) dispatchCreateSparkBatchJob(req *http.Request) (*http.Response, error) {
	if b.srv.CreateSparkBatchJob == nil {
		return nil, &nonRetriableError{errors.New("method CreateSparkBatchJob not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/batches"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[azspark.BatchJobOptions](req)
	if err != nil {
		return nil, err
	}
	detailedParam, err := parseOptional(qp.Get("detailed"), strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *azspark.BatchClientCreateSparkBatchJobOptions
	if detailedParam != nil {
		options = &azspark.BatchClientCreateSparkBatchJobOptions{
			Detailed: detailedParam,
		}
	}
	respr, errRespr := b.srv.CreateSparkBatchJob(req.Context(), body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BatchJob, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BatchServerTransport) dispatchGetSparkBatchJob(req *http.Request) (*http.Response, error) {
	if b.srv.GetSparkBatchJob == nil {
		return nil, &nonRetriableError{errors.New("method GetSparkBatchJob not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/batches/(?P<batchId>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	batchIDParam, err := parseWithCast(matches[regex.SubexpIndex("batchId")], func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	detailedParam, err := parseOptional(qp.Get("detailed"), strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *azspark.BatchClientGetSparkBatchJobOptions
	if detailedParam != nil {
		options = &azspark.BatchClientGetSparkBatchJobOptions{
			Detailed: detailedParam,
		}
	}
	respr, errRespr := b.srv.GetSparkBatchJob(req.Context(), int32(batchIDParam), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BatchJob, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BatchServerTransport) dispatchGetSparkBatchJobs(req *http.Request) (*http.Response, error) {
	if b.srv.GetSparkBatchJobs == nil {
		return nil, &nonRetriableError{errors.New("method GetSparkBatchJobs not implemented")}
	}
	const regexStr = "/livyApi/versions/(?P<livyApiVersion>[a-zA-Z0-9-_]+)/sparkPools/(?P<sparkPoolName>[a-zA-Z0-9-_]+)/batches"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	fromParam, err := parseOptional(qp.Get("from"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	sizeParam, err := parseOptional(qp.Get("size"), func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	detailedParam, err := parseOptional(qp.Get("detailed"), strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *azspark.BatchClientGetSparkBatchJobsOptions
	if fromParam != nil || sizeParam != nil || detailedParam != nil {
		options = &azspark.BatchClientGetSparkBatchJobsOptions{
			From:     fromParam,
			Size:     sizeParam,
			Detailed: detailedParam,
		}
	}
	respr, errRespr := b.srv.GetSparkBatchJobs(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BatchJobCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
