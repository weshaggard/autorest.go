// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armlargeinstance

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"time"
)

// AzureLargeInstance - Azure Large Instance info on Azure (ARM properties and AzureLargeInstance
// properties)
type AzureLargeInstance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *Properties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// READ-ONLY; Name of the AzureLargeInstance.
	Name *string

	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// Resource tags.
	Tags map[string]*string
}

// AzureLargeStorageInstance info on Azure (ARM properties and
// AzureLargeStorageInstance properties)
type AzureLargeStorageInstance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *AzureLargeStorageInstanceProperties

	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// READ-ONLY; Name of the AzureLargeStorageInstance.
	Name *string

	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// Resource tags.
	Tags map[string]*string
}

// AzureLargeStorageInstanceListResult - The response of a AzureLargeStorageInstance list operation.
type AzureLargeStorageInstanceListResult struct {
	// REQUIRED; The AzureLargeStorageInstance items on this page
	Value []*AzureLargeStorageInstance

	// The link to the next page of items
	NextLink *string
}

// AzureLargeStorageInstanceProperties - Describes the properties of an AzureLargeStorageInstance.
type AzureLargeStorageInstanceProperties struct {
	// Specifies the AzureLargeStorageInstance unique ID.
	AzureLargeStorageInstanceUniqueIdentifier *string

	// Specifies the storage properties for the AzureLargeStorage instance.
	StorageProperties *StorageProperties
}

// AzureLargeStorageInstanceTagsUpdate - The type used for updating tags in AzureLargeStorageInstance resources.
type AzureLargeStorageInstanceTagsUpdate struct {
	// Resource tags.
	Tags map[string]*string
}

// Disk - Specifies the disk information fo the Azure Large Instance
type Disk struct {
	// Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB *int32

	// Specifies the logical unit number of the data disk. This value is used to
	// identify data disks within the VM and therefore must be unique for each data
	// disk attached to a VM.
	Lun *int32

	// The disk name.
	Name *string
}

// ForceState - The active state empowers the server with the ability to forcefully terminate
// and halt any existing processes that may be running on the server
type ForceState struct {
	// Whether to force restart by shutting all processes.
	ForceState *ForcePowerState
}

// HardwareProfile - Specifies the hardware settings for the Azure Large Instance.
type HardwareProfile struct {
	// Specifies the Azure Large Instance SKU.
	AzureLargeInstanceSize *SizeNamesEnum

	// Name of the hardware type (vendor and/or their product name)
	HardwareType *HardwareTypeNamesEnum
}

// IPAddress - Specifies the IP address of the network interface.
type IPAddress struct {
	// Specifies the IP address of the network interface.
	IPAddress *string
}

// ListResult - The response of a AzureLargeInstance list operation.
type ListResult struct {
	// REQUIRED; The AzureLargeInstance items on this page
	Value []*AzureLargeInstance

	// The link to the next page of items
	NextLink *string
}

// NetworkProfile - Specifies the network settings for the Azure Large Instance disks.
type NetworkProfile struct {
	// Specifies the circuit id for connecting to express route.
	CircuitID *string

	// Specifies the network interfaces for the Azure Large Instance.
	NetworkInterfaces []*IPAddress
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

// OperationStatusResult - The current status of an async operation.
type OperationStatusResult struct {
	// REQUIRED; The operations list.
	Operations []*OperationStatusResult

	// REQUIRED; Operation status.
	Status *string

	// The end time of the operation.
	EndTime *time.Time

	// If present, details of the operation error.
	Error *azcore.ResponseError

	// Fully qualified ID for the async operation.
	ID *string

	// Name of the async operation.
	Name *string

	// Percent of the operation that is complete.
	PercentComplete *int32

	// The start time of the operation.
	StartTime *time.Time
}

// OsProfile - Specifies the operating system settings for the Azure Large Instance.
type OsProfile struct {
	// Specifies the host OS name of the Azure Large Instance.
	ComputerName *string

	// This property allows you to specify the type of the OS.
	OSType *string

	// Specifies the SSH public key used to access the operating system.
	SSHPublicKey *string

	// Specifies version of operating system.
	Version *string
}

// PagedOperation - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get
// the next set of results.
type PagedOperation struct {
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// Properties - Describes the properties of an Azure Large Instance.
type Properties struct {
	// Specifies the Azure Large Instance unique ID.
	AzureLargeInstanceID *string

	// Specifies the hardware settings for the Azure Large Instance.
	HardwareProfile *HardwareProfile

	// Hardware revision of an Azure Large Instance
	HwRevision *string

	// Specifies the network settings for the Azure Large Instance.
	NetworkProfile *NetworkProfile

	// Specifies the operating system settings for the Azure Large Instance.
	OSProfile *OsProfile

	// ARM ID of another AzureLargeInstance that will share a network with this
	// AzureLargeInstance
	PartnerNodeID *string

	// Resource power state
	PowerState *PowerStateEnum

	// State of provisioning of the AzureLargeInstance
	ProvisioningState *ProvisioningStatesEnum

	// Resource proximity placement group
	ProximityPlacementGroup *string

	// Specifies the storage settings for the Azure Large Instance disks.
	StorageProfile *StorageProfile
}

// StorageBillingProperties - Describes the billing related details of the AzureLargeStorageInstance.
type StorageBillingProperties struct {
	// the billing mode for the storage instance
	BillingMode *string

	// the SKU type that is provisioned
	SKU *string
}

// StorageProfile - Specifies the storage settings for the Azure Large Instance disks.
type StorageProfile struct {
	// IP Address to connect to storage.
	NfsIPAddress *string

	// Specifies information about the operating system disk used by Azure Large
	// Instance.
	OSDisks []*Disk
}

// StorageProperties - described the storage properties of the azure large storage instance
type StorageProperties struct {
	// the kind of storage instance
	Generation *string

	// the hardware type of the storage instance
	HardwareType *HardwareTypeNamesEnum

	// the offering type for which the resource is getting provisioned
	OfferingType *string

	// State of provisioning of the AzureLargeStorageInstance
	ProvisioningState *ProvisioningState

	// the billing related information for the resource
	StorageBillingProperties *StorageBillingProperties

	// the storage protocol for which the resource is getting provisioned
	StorageType *string

	// the workload for which the resource is getting provisioned
	WorkloadType *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
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

// TagsUpdate - The type used for updating tags in AzureLargeInstance resources.
type TagsUpdate struct {
	// Resource tags.
	Tags map[string]*string
}
