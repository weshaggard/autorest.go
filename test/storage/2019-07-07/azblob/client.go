// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/url"
	"strings"
)

const scope = "https://storage.azure.com/.default"
const telemetryInfo = "azsdk-go-azblob/<version>"

// ClientOptions contains configuration settings for the default client's pipeline.
type clientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
	// ApplicationID is an application-specific identification string used in telemetry.
	// It has a maximum length of 24 characters and must not contain any spaces.
	ApplicationID string
}

// defaultClientOptions creates a clientOptions type initialized with default values.
func defaultClientOptions() clientOptions {
	return clientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

func (c *ClientOptions) telemetryOptions() azcore.TelemetryOptions {
	t := telemetryInfo
	if c.ApplicationID != "" {
		a := strings.ReplaceAll(c.ApplicationID, " ", "/")
		if len(a) > 24 {
			a = a[:24]
		}
		t = fmt.Sprintf("%s %s", a, telemetryInfo)
	}
	if c.Telemetry.Value == "" {
		return azcore.TelemetryOptions{Value: t}
	}
	return azcore.TelemetryOptions{Value: fmt.Sprintf("%s %s", c.Telemetry.Value, t)}
}

type client struct {
	u *url.URL
	p azcore.Pipeline
}

// newClient creates an instance of the client type with the specified endpoint.
func newClient(endpoint string, cred azcore.Credential, options *clientOptions) (*client, error) {
	if options == nil {
		o := defaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}}),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return newClientWithPipeline(endpoint, p)
}

// newClientWithPipeline creates an instance of the client type with the specified endpoint and pipeline.
func newClientWithPipeline(endpoint string, p azcore.Pipeline) (*client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		return nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)
	}
	return &client{u: u, p: p}, nil
}

// ServiceOperations returns the ServiceOperations associated with this client.
func (client *client) ServiceOperations() ServiceOperations {
	return &serviceOperations{client: client}
}

// ContainerOperations returns the ContainerOperations associated with this client.
func (client *client) ContainerOperations() ContainerOperations {
	return &containerOperations{client: client}
}

// DirectoryOperations returns the DirectoryOperations associated with this client.
func (client *client) DirectoryOperations(pathRenameMode *PathRenameMode) DirectoryOperations {
	return &directoryOperations{client: client, pathRenameMode: pathRenameMode}
}

// BlobOperations returns the BlobOperations associated with this client.
func (client *client) BlobOperations(pathRenameMode *PathRenameMode) BlobOperations {
	return &blobOperations{client: client, pathRenameMode: pathRenameMode}
}

// PageBlobOperations returns the PageBlobOperations associated with this client.
func (client *client) PageBlobOperations() PageBlobOperations {
	return &pageBlobOperations{client: client}
}

// AppendBlobOperations returns the AppendBlobOperations associated with this client.
func (client *client) AppendBlobOperations() AppendBlobOperations {
	return &appendBlobOperations{client: client}
}

// BlockBlobOperations returns the BlockBlobOperations associated with this client.
func (client *client) BlockBlobOperations() BlockBlobOperations {
	return &blockBlobOperations{client: client}
}
