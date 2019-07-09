// Package azurereport implements the Azure ARM Azurereport service API version 1.0.0.
//
// Test Infrastructure for AutoRest
package azurereport

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Azurereport
	DefaultBaseURI = "http://localhost:3000"
)

// BaseClient is the base client for Azurereport.
type BaseClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}

// GetReport get test coverage report
// Parameters:
// qualifier - if specified, qualifies the generated report further (e.g. '2.7' vs '3.5' in for Python). The
// only effect is, that generators that run all tests several times, can distinguish the generated reports.
func (client BaseClient) GetReport(ctx context.Context, qualifier string) (result SetInt32, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetReport")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetReportPreparer(ctx, qualifier)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurereport.BaseClient", "GetReport", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetReportSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurereport.BaseClient", "GetReport", resp, "Failure sending request")
		return
	}

	result, err = client.GetReportResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurereport.BaseClient", "GetReport", resp, "Failure responding to request")
	}

	return
}

// GetReportPreparer prepares the GetReport request.
func (client BaseClient) GetReportPreparer(ctx context.Context, qualifier string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(qualifier) > 0 {
		queryParameters["qualifier"] = autorest.Encode("query", qualifier)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/report/azure"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetReportSender sends the GetReport request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetReportSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetReportResponder handles the response to the GetReport request. The method always
// closes the http.Response Body.
func (client BaseClient) GetReportResponder(resp *http.Response) (result SetInt32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
