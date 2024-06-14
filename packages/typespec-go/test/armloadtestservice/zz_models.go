// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armloadtestservice

import "time"

// CheckQuotaAvailabilityResponse - Check quota availability response object.
type CheckQuotaAvailabilityResponse struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The name of the resource.
	Name *string

	// Check quota availability response properties.
	Properties *CheckQuotaAvailabilityResponseProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData
}

// CheckQuotaAvailabilityResponseProperties - Check quota availability response properties.
type CheckQuotaAvailabilityResponseProperties struct {
	// Message indicating additional details to add to quota support request.
	AvailabilityStatus *string

	// True/False indicating whether the quota request be granted based on availability.
	IsAvailable *bool
}

// EncryptionProperties - Key and identity details for Customer Managed Key encryption of load test resource.
type EncryptionProperties struct {
	// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
	Identity *EncryptionPropertiesIdentity

	// key encryption key Url, versioned. Ex: https://contosovault.vault.azure.net/keys/contosokek/562a4bb76b524a1493a6afe8e536ee78
	// or https://contosovault.vault.azure.net/keys/contosokek.
	KeyURL *string
}

// EncryptionPropertiesIdentity - All identity configuration for Customer-managed key settings defining which identity should
// be used to auth to Key Vault.
type EncryptionPropertiesIdentity struct {
	// User assigned identity to use for accessing key encryption key Url. Ex: /subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/<resource
	// group>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId.
	ResourceID *string

	// Managed identity type to use for accessing encryption key Url.
	Type *Type
}

// EndpointDependency - A domain name and connection details used to access a dependency.
type EndpointDependency struct {
	// REQUIRED; The domain name of the dependency. Domain names may be fully qualified or may contain a * wildcard.
	DomainName *string

	// Human-readable supplemental information about the dependency and when it is applicable.
	Description *string

	// The list of connection details for this endpoint.
	EndpointDetails []*EndpointDetail
}

// EndpointDetail - Details about the connection between the Batch service and the endpoint.
type EndpointDetail struct {
	// The port an endpoint is connected to.
	Port *int32
}

// LoadTestProperties - LoadTest resource properties.
type LoadTestProperties struct {
	// Resource data plane URI.
	DataPlaneURI *string

	// Description of the resource.
	Description *string

	// CMK Encryption property.
	Encryption *EncryptionProperties

	// Resource provisioning state.
	ProvisioningState *ResourceState
}

// LoadTestResource - LoadTest details.
type LoadTestResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// READ-ONLY; Load Test name
	Name *string

	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *LoadTestProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// Resource tags.
	Tags map[string]*string

	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// LoadTestResourceListResult - The response of a LoadTestResource list operation.
type LoadTestResourceListResult struct {
	// REQUIRED; The LoadTestResource items on this page
	Value []*LoadTestResource

	// The link to the next page of items
	NextLink *string
}

// LoadTestResourceUpdate - The type used for update operations of the LoadTestResource.
type LoadTestResourceUpdate struct {
	// The managed service identities assigned to this resource.
	Identity   *ManagedServiceIdentity
	Properties *LoadTestResourceUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// LoadTestResourceUpdateProperties - The updatable properties of the LoadTestResource.
type LoadTestResourceUpdateProperties struct {
	// Description of the resource.
	Description *string

	// CMK Encryption property.
	Encryption *EncryptionProperties
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; The type of managed identity assigned to this resource.
	Type *ManagedServiceIdentityType

	// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
	PrincipalID *string

	// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// Localized display information for this particular operation.
	Display *OperationDisplay

	// Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure Resource Manager/control-plane
	// operations.
	IsDataAction *bool

	// The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is
	// "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
type OperationDisplay struct {
	// The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual Machine",
	// "Restart Virtual Machine".
	Operation *string

	// The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft Compute".
	Provider *string

	// The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// OutboundEnvironmentEndpoint - A collection of related endpoints from the same service for which the Batch service requires
// outbound access.
type OutboundEnvironmentEndpoint struct {
	// The type of service that Azure Load Testing connects to.
	Category *string

	// The endpoints for this service to which the Batch service makes outbound calls.
	Endpoints []*EndpointDependency
}

// PagedOutboundEnvironmentEndpoint - Values returned by the List operation.
type PagedOutboundEnvironmentEndpoint struct {
	// REQUIRED; The OutboundEnvironmentEndpoint items on this page
	Value []*OutboundEnvironmentEndpoint

	// The link to the next page of items
	NextLink *string
}

// QuotaBucketRequest - Request object of new quota for a quota bucket.
type QuotaBucketRequest struct {
	// Request object of new quota for a quota bucket.
	Properties *QuotaBucketRequestProperties
}

// QuotaBucketRequestProperties - New quota request request properties.
type QuotaBucketRequestProperties struct {
	// Current quota limit of the quota bucket.
	CurrentQuota *int32

	// Current quota usage of the quota bucket.
	CurrentUsage *int32

	// Dimensions for new quota request.
	Dimensions *QuotaBucketRequestPropertiesDimensions

	// New quota limit of the quota bucket.
	NewQuota *int32
}

// QuotaBucketRequestPropertiesDimensions - Dimensions for new quota request.
type QuotaBucketRequestPropertiesDimensions struct {
	// Location dimension for new quota request of the quota bucket.
	Location *string

	// Subscription Id dimension for new quota request of the quota bucket.
	SubscriptionID *string
}

// QuotaResource - Quota bucket details object.
type QuotaResource struct {
	// The resource-specific properties for this resource.
	Properties *QuotaResourceProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// READ-ONLY; The quota name.
	Name *string

	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string
}

// QuotaResourceListResult - The response of a QuotaResource list operation.
type QuotaResourceListResult struct {
	// REQUIRED; The QuotaResource items on this page
	Value []*QuotaResource

	// The link to the next page of items
	NextLink *string
}

// QuotaResourceProperties - Quota bucket resource properties.
type QuotaResourceProperties struct {
	// Current quota limit of the quota bucket.
	Limit *int32

	// Resource provisioning state.
	ProvisioningState *ResourceState

	// Current quota usage of the quota bucket.
	Usage *int32
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// The client ID of the assigned identity.
	ClientID *string

	// The principal ID of the assigned identity.
	PrincipalID *string
}
