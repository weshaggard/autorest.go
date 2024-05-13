// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter

// APIDefinitionsClientCreateOrUpdateResponse contains the response from method APIDefinitionsClient.CreateOrUpdate.
type APIDefinitionsClientCreateOrUpdateResponse struct {
	// API definition entity.
	APIDefinition

	// The entity tag for the response.
	ETag *string
}

// APIDefinitionsClientDeleteResponse contains the response from method APIDefinitionsClient.Delete.
type APIDefinitionsClientDeleteResponse struct {
	// placeholder for future response values
}

// APIDefinitionsClientExportSpecificationResponse contains the response from method APIDefinitionsClient.BeginExportSpecification.
type APIDefinitionsClientExportSpecificationResponse struct {
	// The API specification export result.
	APISpecExportResult
}

// APIDefinitionsClientGetResponse contains the response from method APIDefinitionsClient.Get.
type APIDefinitionsClientGetResponse struct {
	// API definition entity.
	APIDefinition

	// The entity tag for the response.
	ETag *string
}

// APIDefinitionsClientHeadResponse contains the response from method APIDefinitionsClient.Head.
type APIDefinitionsClientHeadResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// APIDefinitionsClientImportSpecificationResponse contains the response from method APIDefinitionsClient.BeginImportSpecification.
type APIDefinitionsClientImportSpecificationResponse struct {
	// The API specification was successfully imported.
	APIImportSuccess
}

// APIDefinitionsClientListResponse contains the response from method APIDefinitionsClient.NewListPager.
type APIDefinitionsClientListResponse struct {
	// The response of a ApiDefinition list operation.
	APIDefinitionListResult
}

// APIVersionsClientCreateOrUpdateResponse contains the response from method APIVersionsClient.CreateOrUpdate.
type APIVersionsClientCreateOrUpdateResponse struct {
	// API version entity.
	APIVersion

	// The entity tag for the response.
	ETag *string
}

// APIVersionsClientDeleteResponse contains the response from method APIVersionsClient.Delete.
type APIVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// APIVersionsClientGetResponse contains the response from method APIVersionsClient.Get.
type APIVersionsClientGetResponse struct {
	// API version entity.
	APIVersion

	// The entity tag for the response.
	ETag *string
}

// APIVersionsClientHeadResponse contains the response from method APIVersionsClient.Head.
type APIVersionsClientHeadResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// APIVersionsClientListResponse contains the response from method APIVersionsClient.NewListPager.
type APIVersionsClientListResponse struct {
	// The response of a ApiVersion list operation.
	APIVersionListResult
}

// ApisClientCreateOrUpdateResponse contains the response from method ApisClient.CreateOrUpdate.
type ApisClientCreateOrUpdateResponse struct {
	// API entity.
	API

	// The entity tag for the response.
	ETag *string
}

// ApisClientDeleteResponse contains the response from method ApisClient.Delete.
type ApisClientDeleteResponse struct {
	// placeholder for future response values
}

// ApisClientGetResponse contains the response from method ApisClient.Get.
type ApisClientGetResponse struct {
	// API entity.
	API

	// The entity tag for the response.
	ETag *string
}

// ApisClientHeadResponse contains the response from method ApisClient.Head.
type ApisClientHeadResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// ApisClientListResponse contains the response from method ApisClient.NewListPager.
type ApisClientListResponse struct {
	// The response of a Api list operation.
	APIListResult
}

// DeletedServicesClientDeleteResponse contains the response from method DeletedServicesClient.Delete.
type DeletedServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// DeletedServicesClientGetResponse contains the response from method DeletedServicesClient.Get.
type DeletedServicesClientGetResponse struct {
	// Soft-deleted service entity.
	DeletedService

	// The entity tag for the response.
	ETag *string
}

// DeletedServicesClientListBySubscriptionResponse contains the response from method DeletedServicesClient.NewListBySubscriptionPager.
type DeletedServicesClientListBySubscriptionResponse struct {
	// The response of a DeletedService list operation.
	DeletedServiceListResult
}

// DeletedServicesClientListResponse contains the response from method DeletedServicesClient.NewListPager.
type DeletedServicesClientListResponse struct {
	// The response of a DeletedService list operation.
	DeletedServiceListResult
}

// DeploymentsClientCreateOrUpdateResponse contains the response from method DeploymentsClient.CreateOrUpdate.
type DeploymentsClientCreateOrUpdateResponse struct {
	// API deployment entity.
	Deployment

	// The entity tag for the response.
	ETag *string
}

// DeploymentsClientDeleteResponse contains the response from method DeploymentsClient.Delete.
type DeploymentsClientDeleteResponse struct {
	// placeholder for future response values
}

// DeploymentsClientGetResponse contains the response from method DeploymentsClient.Get.
type DeploymentsClientGetResponse struct {
	// API deployment entity.
	Deployment

	// The entity tag for the response.
	ETag *string
}

// DeploymentsClientHeadResponse contains the response from method DeploymentsClient.Head.
type DeploymentsClientHeadResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// DeploymentsClientListResponse contains the response from method DeploymentsClient.NewListPager.
type DeploymentsClientListResponse struct {
	// The response of a Deployment list operation.
	DeploymentListResult
}

// EnvironmentsClientCreateOrUpdateResponse contains the response from method EnvironmentsClient.CreateOrUpdate.
type EnvironmentsClientCreateOrUpdateResponse struct {
	// Environment entity.
	Environment

	// The entity tag for the response.
	ETag *string
}

// EnvironmentsClientDeleteResponse contains the response from method EnvironmentsClient.Delete.
type EnvironmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// EnvironmentsClientGetResponse contains the response from method EnvironmentsClient.Get.
type EnvironmentsClientGetResponse struct {
	// Environment entity.
	Environment

	// The entity tag for the response.
	ETag *string
}

// EnvironmentsClientHeadResponse contains the response from method EnvironmentsClient.Head.
type EnvironmentsClientHeadResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// EnvironmentsClientListResponse contains the response from method EnvironmentsClient.NewListPager.
type EnvironmentsClientListResponse struct {
	// The response of a Environment list operation.
	EnvironmentListResult
}

// MetadataSchemasClientCreateOrUpdateResponse contains the response from method MetadataSchemasClient.CreateOrUpdate.
type MetadataSchemasClientCreateOrUpdateResponse struct {
	// Metadata schema entity. Used to define metadata for the entities in API catalog.
	MetadataSchema

	// The entity tag for the response.
	ETag *string
}

// MetadataSchemasClientDeleteResponse contains the response from method MetadataSchemasClient.Delete.
type MetadataSchemasClientDeleteResponse struct {
	// placeholder for future response values
}

// MetadataSchemasClientGetResponse contains the response from method MetadataSchemasClient.Get.
type MetadataSchemasClientGetResponse struct {
	// Metadata schema entity. Used to define metadata for the entities in API catalog.
	MetadataSchema

	// The entity tag for the response.
	ETag *string
}

// MetadataSchemasClientHeadResponse contains the response from method MetadataSchemasClient.Head.
type MetadataSchemasClientHeadResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// MetadataSchemasClientListResponse contains the response from method MetadataSchemasClient.NewListPager.
type MetadataSchemasClientListResponse struct {
	// The response of a MetadataSchema list operation.
	MetadataSchemaListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	PagedOperation
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// The service entity.
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientExportMetadataSchemaResponse contains the response from method ServicesClient.BeginExportMetadataSchema.
type ServicesClientExportMetadataSchemaResponse struct {
	// The metadata schema export result.
	MetadataSchemaExportResult
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// The service entity.
	Service
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.NewListByResourceGroupPager.
type ServicesClientListByResourceGroupResponse struct {
	// The response of a Service list operation.
	ServiceListResult
}

// ServicesClientListBySubscriptionResponse contains the response from method ServicesClient.NewListBySubscriptionPager.
type ServicesClientListBySubscriptionResponse struct {
	// The response of a Service list operation.
	ServiceListResult
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	// The service entity.
	Service
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.CreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	// Workspace entity.
	Workspace

	// The entity tag for the response.
	ETag *string
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.Delete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	// Workspace entity.
	Workspace

	// The entity tag for the response.
	ETag *string
}

// WorkspacesClientHeadResponse contains the response from method WorkspacesClient.Head.
type WorkspacesClientHeadResponse struct {
	// Success indicates if the operation succeeded or failed.
	Success bool
}

// WorkspacesClientListResponse contains the response from method WorkspacesClient.NewListPager.
type WorkspacesClientListResponse struct {
	// The response of a Workspace list operation.
	WorkspaceListResult
}
