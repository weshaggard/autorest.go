// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// Client - Azure API Center Resource Provider.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClient creates a new instance of Client with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewAPIDefinitionsClient creates a new instance of [APIDefinitionsClient].
func (client *Client) NewAPIDefinitionsClient() *APIDefinitionsClient {
	return &APIDefinitionsClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewAPIVersionsClient creates a new instance of [APIVersionsClient].
func (client *Client) NewAPIVersionsClient() *APIVersionsClient {
	return &APIVersionsClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewApisClient creates a new instance of [ApisClient].
func (client *Client) NewApisClient() *ApisClient {
	return &ApisClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewDeletedServicesClient creates a new instance of [DeletedServicesClient].
func (client *Client) NewDeletedServicesClient() *DeletedServicesClient {
	return &DeletedServicesClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewDeploymentsClient creates a new instance of [DeploymentsClient].
func (client *Client) NewDeploymentsClient() *DeploymentsClient {
	return &DeploymentsClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewEnvironmentsClient creates a new instance of [EnvironmentsClient].
func (client *Client) NewEnvironmentsClient() *EnvironmentsClient {
	return &EnvironmentsClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewMetadataSchemasClient creates a new instance of [MetadataSchemasClient].
func (client *Client) NewMetadataSchemasClient() *MetadataSchemasClient {
	return &MetadataSchemasClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewOperationsClient creates a new instance of [OperationsClient].
func (client *Client) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewServicesClient creates a new instance of [ServicesClient].
func (client *Client) NewServicesClient() *ServicesClient {
	return &ServicesClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewWorkspacesClient creates a new instance of [WorkspacesClient].
func (client *Client) NewWorkspacesClient() *WorkspacesClient {
	return &WorkspacesClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}
