// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcodesigning

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	// The CheckNameAvailability operation response.
	CheckNameAvailabilityResult
}

// AccountsClientCreateResponse contains the response from method AccountsClient.Create.
type AccountsClientCreateResponse struct {
	// Trusted signing account resource.
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	// Trusted signing account resource.
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	// The response of a CodeSigningAccount list operation.
	AccountListResult
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.NewListBySubscriptionPager.
type AccountsClientListBySubscriptionResponse struct {
	// The response of a CodeSigningAccount list operation.
	AccountListResult
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	// Trusted signing account resource.
	Account
}

// CertificateProfilesClientCreateResponse contains the response from method CertificateProfilesClient.Create.
type CertificateProfilesClientCreateResponse struct {
	// Certificate profile resource.
	CertificateProfile
}

// CertificateProfilesClientDeleteResponse contains the response from method CertificateProfilesClient.Delete.
type CertificateProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// CertificateProfilesClientGetResponse contains the response from method CertificateProfilesClient.Get.
type CertificateProfilesClientGetResponse struct {
	// Certificate profile resource.
	CertificateProfile
}

// CertificateProfilesClientListByCodeSigningAccountResponse contains the response from method CertificateProfilesClient.NewListByCodeSigningAccountPager.
type CertificateProfilesClientListByCodeSigningAccountResponse struct {
	// The response of a CertificateProfile list operation.
	CertificateProfileListResult
}

// CertificateProfilesClientRevokeCertificateResponse contains the response from method CertificateProfilesClient.RevokeCertificate.
type CertificateProfilesClientRevokeCertificateResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	PagedOperation
}