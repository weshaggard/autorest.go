// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter

import "time"

// API entity.
type API struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *APIProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The name of the API.
	Name *string
}

// API definition entity.
type APIDefinition struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *APIDefinitionProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The name of the API definition.
	Name *string
}

// The response of a ApiDefinition list operation.
type APIDefinitionListResult struct {
	// REQUIRED; The ApiDefinition items on this page
	Value []*APIDefinition

	// The link to the next page of items
	NextLink *string
}

// API definition properties entity.
type APIDefinitionProperties struct {
	// REQUIRED; API definition title.
	Title *string

	// API definition description.
	Description *string

	// API specification details.
	Specification *APIDefinitionPropertiesSpecification
}

// API specification details.
type APIDefinitionPropertiesSpecification struct {
	// Specification name.
	Name *string

	// Specification version.
	Version *string
}

// The API specification was successfully imported.
type APIImportSuccess struct {
}

// The response of a Api list operation.
type APIListResult struct {
	// REQUIRED; The Api items on this page
	Value []*API

	// The link to the next page of items
	NextLink *string
}

// API properties.
type APIProperties struct {
	// REQUIRED; Kind of API. For example, REST or GraphQL.
	Kind *APIKind

	// REQUIRED; API title.
	Title *string

	// The set of contacts
	Contacts []*Contact

	// The custom metadata defined for API catalog entities.
	CustomProperties *CustomProperties

	// Description of the API.
	Description *string

	// The set of external documentation
	ExternalDocumentation []*ExternalDocumentation

	// The license information for the API.
	License *License

	// Current lifecycle stage of the API.
	LifecycleStage *LifecycleStage

	// Short description of the API.
	Summary *string

	// Terms of service for the API.
	TermsOfService *TermsOfService
}

// The API specification export result.
type APISpecExportResult struct {
	// The format of exported result
	Format *APISpecExportResultFormat

	// The result of the export operation.
	Value *string
}

// The API specification source entity properties.
type APISpecImportRequest struct {
	// Format of the API specification source.
	Format *APISpecImportSourceFormat

	// API specification details.
	Specification *APISpecImportRequestSpecification

	// Value of the API specification source.
	Value *string
}

// API specification details.
type APISpecImportRequestSpecification struct {
	// Specification name.
	Name *string

	// Specification version.
	Version *string
}

// API version entity.
type APIVersion struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *APIVersionProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The name of the API version.
	Name *string
}

// The response of a ApiVersion list operation.
type APIVersionListResult struct {
	// REQUIRED; The ApiVersion items on this page
	Value []*APIVersion

	// The link to the next page of items
	NextLink *string
}

// API version properties entity.
type APIVersionProperties struct {
	// REQUIRED; Current lifecycle stage of the API.
	LifecycleStage *LifecycleStage

	// REQUIRED; API version title.
	Title *string
}

// Common properties for all Azure Resource Manager resources.
type ArmResource struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData
}

// Base class used for type definitions
type ArmResourceBase struct {
}

// Contact information
type Contact struct {
	// Email address of the contact.
	Email *string

	// Name of the contact.
	Name *string

	// URL for the contact.
	URL *string
}

// Custom Properties
type CustomProperties struct {
}

// API deployment entity.
type Deployment struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *DeploymentProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The name of the API deployment.
	Name *string
}

// The response of a Deployment list operation.
type DeploymentListResult struct {
	// REQUIRED; The Deployment items on this page
	Value []*Deployment

	// The link to the next page of items
	NextLink *string
}

// API deployment entity properties.
type DeploymentProperties struct {
	// The custom metadata defined for API catalog entities.
	CustomProperties *CustomProperties

	// API center-scoped definition resource ID.
	DefinitionID *string

	// Description of the deployment.
	Description *string

	// API center-scoped environment resource ID.
	EnvironmentID *string

	// The deployment server
	Server *DeploymentServer

	// State of API deployment.
	State *DeploymentState

	// API deployment title
	Title *string
}

// Server
type DeploymentServer struct {
	// Base runtime URLs for this deployment.
	RuntimeURI []*string
}

// Environment entity.
type Environment struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *EnvironmentProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The name of the environment.
	Name *string
}

// The response of a Environment list operation.
type EnvironmentListResult struct {
	// REQUIRED; The Environment items on this page
	Value []*Environment

	// The link to the next page of items
	NextLink *string
}

// Environment properties entity.
type EnvironmentProperties struct {
	// REQUIRED; Environment kind.
	Kind *EnvironmentKind

	// REQUIRED; Environment title.
	Title *string

	// The custom metadata defined for API catalog entities.
	CustomProperties *CustomProperties

	// The environment description.
	Description *string

	// Environment onboarding information
	Onboarding *Onboarding

	// Server information of the environment.
	Server *EnvironmentServer
}

// Server information of the environment.
type EnvironmentServer struct {
	// The location of the management portal
	ManagementPortalURI []*string

	// Type of the server that represents the environment.
	Type *EnvironmentServerType
}

type ErrorAdditionalInfoInfo struct {
}

// Additional, external documentation for the API.
type ExternalDocumentation struct {
	// REQUIRED; URL pointing to the documentation.
	URL *string

	// Description of the documentation.
	Description *string

	// Title of the documentation.
	Title *string
}

// The license information for the API.
type License struct {
	// SPDX license information for the API. The identifier field is mutually
	// exclusive of the URL field.
	Identifier *string

	// Name of the license.
	Name *string

	// URL pointing to the license details. The URL field is mutually exclusive of the
	// identifier field.
	URL *string
}

// The properties of the managed service identities assigned to this resource.
type ManagedIdentityProperties struct {
	// REQUIRED; The type of managed identity assigned to this resource.
	Type *ManagedIdentityType

	// The active directory identifier of this principal.
	PrincipalID *string

	// The Active Directory tenant id of the principal.
	TenantID *string

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity
}

// Assignment metadata
type MetadataAssignment struct {
	// Deprecated assignment
	Deprecated *bool

	// The entities this metadata schema component gets applied to.
	Entity *MetadataAssignmentEntity

	// Required assignment
	Required *bool
}

// Metadata schema entity. Used to define metadata for the entities in API catalog.
type MetadataSchema struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *MetadataSchemaProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The name of the metadata schema.
	Name *string
}

// The metadata schema export request.
type MetadataSchemaExportRequest struct {
	// An entity the metadata schema is requested for.
	AssignedTo *MetadataAssignmentEntity
}

// The metadata schema export result.
type MetadataSchemaExportResult struct {
	// The export format for the schema
	Format *MetadataSchemaExportFormat

	// The result of the export operation.
	Value *string
}

// The response of a MetadataSchema list operation.
type MetadataSchemaListResult struct {
	// REQUIRED; The MetadataSchema items on this page
	Value []*MetadataSchema

	// The link to the next page of items
	NextLink *string
}

// Metadata schema properties.
type MetadataSchemaProperties struct {
	// REQUIRED; The schema defining the type.
	Schema *string

	// The assignees
	AssignedTo []*MetadataAssignment
}

// Onboarding information
type Onboarding struct {
	// The location of the development portal
	DeveloperPortalURI []*string

	// Onboarding guide.
	Instructions *string
}

// Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
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

// Localized display information for and operation.
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

// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
type PagedOperation struct {
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// The base proxy resource.
type ProxyResourceBase struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData
}

// The service entity.
type Service struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *ServiceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The name of Azure API Center service.
	Name *string

	// The managed service identities assigned to this resource.
	Identity *ManagedIdentityProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData
}

// The response of a Service list operation.
type ServiceListResult struct {
	// REQUIRED; The Service items on this page
	Value []*Service

	// The link to the next page of items
	NextLink *string
}

// The properties of the service.
type ServiceProperties struct {
	// Provisioning state of the service.
	ProvisioningState *ProvisioningState
}

// The type used for update operations of the Service.
type ServiceUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedIdentityProperties

	// Resource tags.
	Tags map[string]*string
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The type of identity that created the resource.
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

// Terms of service for the API.
type TermsOfService struct {
	// REQUIRED; URL pointing to the terms of service.
	URL *string
}

// The resource model definition for an Azure Resource Manager tracked top level resource
type TrackedResourceBase struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// Resource tags.
	Tags map[string]*string
}

// A managed identity assigned by the user.
type UserAssignedIdentity struct {
	// The active directory client identifier for this principal.
	ClientID *string

	// The active directory identifier for this principal.
	PrincipalID *string
}

// Workspace entity.
type Workspace struct {
	// REQUIRED; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// REQUIRED; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// The resource-specific properties for this resource.
	Properties *WorkspaceProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The name of the workspace.
	Name *string
}

// The response of a Workspace list operation.
type WorkspaceListResult struct {
	// REQUIRED; The Workspace items on this page
	Value []*Workspace

	// The link to the next page of items
	NextLink *string
}

// Workspace properties.
type WorkspaceProperties struct {
	// REQUIRED; Workspace title.
	Title *string

	// Workspace description.
	Description *string
}