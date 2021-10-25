//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// Connection - Test Infrastructure for AutoRest Swagger BAT
type Connection struct {
	u string
	p runtime.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "http://localhost:3000"

// NewDefaultConnection creates an instance of the Connection type using the DefaultEndpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func NewDefaultConnection(options *azcore.ClientOptions) *Connection {
	return NewConnection(DefaultEndpoint, options)
}

// NewConnection creates an instance of the Connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func NewConnection(endpoint string, options *azcore.ClientOptions) *Connection {
	cp := azcore.ClientOptions{}
	if options != nil {
		cp = *options
	}
	return &Connection{u: endpoint, p: runtime.NewPipeline(module, version, nil, nil, &cp)}
}

// Endpoint returns the connection's endpoint.
func (c *Connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *Connection) Pipeline() runtime.Pipeline {
	return c.p
}
