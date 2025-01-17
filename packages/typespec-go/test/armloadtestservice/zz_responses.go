// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armloadtestservice

// LoadTestsClientCreateOrUpdateResponse contains the response from method LoadTestsClient.BeginCreateOrUpdate.
type LoadTestsClientCreateOrUpdateResponse struct {
	// LoadTest details.
	LoadTestResource
}

// LoadTestsClientDeleteResponse contains the response from method LoadTestsClient.BeginDelete.
type LoadTestsClientDeleteResponse struct {
	// placeholder for future response values
}

// LoadTestsClientGetResponse contains the response from method LoadTestsClient.Get.
type LoadTestsClientGetResponse struct {
	// LoadTest details.
	LoadTestResource
}

// LoadTestsClientListByResourceGroupResponse contains the response from method LoadTestsClient.NewListByResourceGroupPager.
type LoadTestsClientListByResourceGroupResponse struct {
	// The response of a LoadTestResource list operation.
	LoadTestResourceListResult
}

// LoadTestsClientListBySubscriptionResponse contains the response from method LoadTestsClient.NewListBySubscriptionPager.
type LoadTestsClientListBySubscriptionResponse struct {
	// The response of a LoadTestResource list operation.
	LoadTestResourceListResult
}

// LoadTestsClientOutboundNetworkDependenciesEndpointsResponse contains the response from method LoadTestsClient.NewOutboundNetworkDependenciesEndpointsPager.
type LoadTestsClientOutboundNetworkDependenciesEndpointsResponse struct {
	// Values returned by the List operation.
	PagedOutboundEnvironmentEndpoint
}

// LoadTestsClientUpdateResponse contains the response from method LoadTestsClient.BeginUpdate.
type LoadTestsClientUpdateResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// QuotasClientCheckAvailabilityResponse contains the response from method QuotasClient.CheckAvailability.
type QuotasClientCheckAvailabilityResponse struct {
	// Check quota availability response object.
	CheckQuotaAvailabilityResponse
}

// QuotasClientGetResponse contains the response from method QuotasClient.Get.
type QuotasClientGetResponse struct {
	// Quota bucket details object.
	QuotaResource
}

// QuotasClientListResponse contains the response from method QuotasClient.NewListPager.
type QuotasClientListResponse struct {
	// The response of a QuotaResource list operation.
	QuotaResourceListResult
}
